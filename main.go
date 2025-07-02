//go:build !wasm && !js

package main

import (
	"bufio"
	"fmt"
	"github.com/Bedrock-OSS/go-burrito/burrito"
	"github.com/MCDevKit/jsonte/jsonte"
	"github.com/MCDevKit/jsonte/jsonte/functions"
	"github.com/MCDevKit/jsonte/jsonte/json"
	"github.com/MCDevKit/jsonte/jsonte/types"
	"github.com/MCDevKit/jsonte/jsonte/utils"
	"github.com/gammazero/deque"
	"github.com/gobwas/glob"
	"go.uber.org/multierr"
	"go.uber.org/zap"
	"math/rand"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"sort"
	"strings"
	"sync"
	"time"
)

var (
	commit      string
	version     = "0.0.0"
	date        string
	buildSource = "DEV"
)

func main() {
	types.Init()
	functions.Init()
	env, b := os.LookupEnv("DEBUG")
	debug := false
	if b && (env == "true" || env == "1") {
		debug = true
	}
	silent := false
	removeSrc := false
	minify := false
	scope := make([]string, 0)
	out := ""
	include := make([]string, 0)
	exclude := make([]string, 0)
	ipcName := "jsonte"
	seed := time.Now().UnixNano()
	workers := int64(1)
	cacheAll := false
	app := NewApp()
	app.BoolFlag(Flag{
		Name:  "debug",
		Usage: "Enable debug mode",
	}, &debug)
	app.BoolFlag(Flag{
		Name:  "safe-mode",
		Usage: "Enable safe mode",
	}, &functions.SafeMode)
	app.BoolFlag(Flag{
		Name:  "server-mode",
		Usage: "Enable server mode",
	}, &functions.ServerMode)
	app.BoolFlag(Flag{
		Name:  "silent",
		Usage: "Enable silent mode",
	}, &silent)
	app.BoolFlag(Flag{
		Name:  "remove-src",
		Usage: "Remove source files after compilation",
	}, &removeSrc)
	app.BoolFlag(Flag{
		Name:  "minify",
		Usage: "Minify the output",
	}, &minify)
	app.StringSliceFlag(Flag{
		Name:  "scope",
		Usage: "Scope of the output",
	}, &scope)
	app.StringFlag(Flag{
		Name:  "out",
		Usage: "Output file",
	}, &out)
	app.StringSliceFlag(Flag{
		Name:  "include",
		Usage: "Include files",
	}, &include)
	app.StringSliceFlag(Flag{
		Name:  "exclude",
		Usage: "Exclude files",
	}, &exclude)
	app.StringFlag(Flag{
		Name:  "cache-dir",
		Usage: "Directory for the cache",
	}, &json.CacheDir)
	app.StringFlag(Flag{
		Name:  "ipc-name",
		Usage: "Name for the IPC named pipe",
	}, &ipcName)
	app.IntFlag(Flag{
		Name:  "seed",
		Usage: "Seed for the random number generator",
	}, &seed)
	app.IntFlag(Flag{
		Name:  "workers",
		Usage: "Maximum number of concurrent workers",
	}, &workers)
	app.BoolFlag(Flag{
		Name:  "cache-all",
		Usage: "Cache all function calls",
	}, &cacheAll)
	app.Action(Action{
		Name: "compile",
		Function: func(args []string) error {
			functions.SetCacheAll(cacheAll)
			outFile := ""
			if out != "" {
				stat, err := os.Stat(out)
				if err != nil {
					if os.IsNotExist(err) {
						err := os.MkdirAll(out, 0755)
						if err != nil {
							return burrito.WrapError(err, "An error occurred while creating the output directory")
						}
					} else {
						return burrito.WrapError(err, "An error occurred while reading the output file")
					}
				} else if !stat.IsDir() {
					return burrito.WrappedErrorf("The output file '%s' is not a directory", out)
				}
				outFile, err = filepath.Abs(out)
			}
			object, err := getScope(scope, -1)
			if err != nil {
				return burrito.WrapError(err, "An error occurred while reading the scope")
			}
			fileSets, err := getFileList(args, include, exclude)
			if err != nil {
				return err
			}
			w := int(workers)
			if w <= 0 {
				w = runtime.NumCPU()
			}
			utils.Logger.Debugf("Using %d workers", w)
			sem := make(chan struct{}, w)
			var errMu, moduleMu sync.Mutex
			var errs error
			appendErr := func(e error) {
				errMu.Lock()
				defer errMu.Unlock()
				errs = multierr.Append(errs, e)
			}
			modules := map[string]jsonte.JsonModule{}
			var wg sync.WaitGroup
			for base, files := range fileSets {
				for _, file := range files {
					if strings.HasSuffix(file, ".modl") {
						wg.Add(1)
						go func(base, file string) {
							defer wg.Done()
							sem <- struct{}{}
							defer func() { <-sem }()
							bytes, err := os.ReadFile(file)
							if err != nil {
								appendErr(burrito.WrapErrorf(err, "An error occurred while reading the module file %s", file))
								return
							}
							module, err := jsonte.LoadModule(string(bytes), object, -1)
							if err != nil {
								appendErr(burrito.WrapErrorf(err, "An error occurred while loading the module file %s", file))
								return
							}
							moduleMu.Lock()
							modules[module.Name.StringValue()] = module
							moduleMu.Unlock()
							rel, err := filepath.Rel(base, file)
							if err != nil {
								appendErr(burrito.WrapErrorf(err, "An error occurred while relativizing the output file name"))
								return
							}
							utils.Logger.Infof("Loaded module %s", rel)
							if removeSrc {
								if err := os.Remove(file); err != nil {
									appendErr(burrito.WrapErrorf(err, "An error occurred while removing the module file %s", file))
								}
							}
						}(base, file)
					}
				}
			}
			wg.Wait()
			if errs != nil {
				return errs
			}

			type templRes struct {
				base     string
				file     string
				template utils.NavigableMap[string, types.JsonType]
			}

			templCh := make(chan templRes)
			wg = sync.WaitGroup{}
			for base, files := range fileSets {
				for _, file := range files {
					if strings.HasSuffix(file, ".templ") {
						wg.Add(1)
						go func(base, file string) {
							defer wg.Done()
							sem <- struct{}{}
							defer func() { <-sem }()
							utils.Logger.Infof("Templating file %s", file)
							bytes, err := os.ReadFile(file)
							if err != nil {
								appendErr(burrito.WrapErrorf(err, "An error occurred while reading the template file %s", file))
								return
							}
							template, err := jsonte.Process(strings.TrimSuffix(filepath.Base(file), filepath.Ext(file)), string(bytes), object, modules, -1)
							if err != nil {
								appendErr(burrito.WrapErrorf(err, "An error occurred while processing the template file %s", file))
								return
							}
							templCh <- templRes{base: base, file: file, template: template}
							if removeSrc {
								if err := os.Remove(file); err != nil {
									appendErr(burrito.WrapErrorf(err, "An error occurred while removing the template file %s", file))
								}
							}
						}(base, file)
					}
				}
			}
			go func() {
				wg.Wait()
				close(templCh)
			}()
			templates := make([]templRes, 0)
			for t := range templCh {
				templates = append(templates, t)
			}
			sort.SliceStable(templates, func(i, j int) bool {
				if templates[i].base == templates[j].base {
					return templates[i].file < templates[j].file
				}
				return templates[i].base < templates[j].base
			})
			toString := types.ToPrettyString
			if minify {
				toString = types.ToString
			}
			for _, t := range templates {
				for _, fileName := range t.template.Keys() {
					content := t.template.Get(fileName)
					filename := filepath.Dir(t.file) + "/" + fileName + ".json"
					if outFile != "" {
						var err error
						filename, err = filepath.Rel(t.base, filename)
						if err != nil {
							appendErr(burrito.WrapErrorf(err, "An error occurred while creating the output file name"))
							continue
						}
						filename = filepath.Join(outFile, t.base, filename)
						rel, err := filepath.Rel(outFile, filename)
						if err != nil {
							appendErr(burrito.WrapErrorf(err, "An error occurred while relativizing the output file name"))
							continue
						}
						utils.Logger.Infof("Writing file %s", filepath.Clean(rel))
					} else {
						utils.Logger.Infof("Writing file %s", filepath.Clean(filename))
					}
					if err := os.MkdirAll(filepath.Dir(filename), 0755); err != nil {
						appendErr(burrito.WrapErrorf(err, "An error occurred while creating the output directory %s", filepath.Dir(filename)))
						continue
					}
					if err := os.WriteFile(filename, []byte(toString(content)), 0644); err != nil {
						appendErr(burrito.WrapErrorf(err, "An error occurred while writing the output file %s", filename))
					}
				}
			}
			if errs != nil {
				return errs
			}

			type fnRes struct {
				base string
				file string
				out  string
				data string
			}

			fnCh := make(chan fnRes)
			wg = sync.WaitGroup{}
			for base, files := range fileSets {
				for _, file := range files {
					if strings.HasSuffix(file, ".mcfunction") {
						wg.Add(1)
						go func(base, file string) {
							defer wg.Done()
							sem <- struct{}{}
							defer func() { <-sem }()
							bytes, err := os.ReadFile(file)
							if err != nil {
								appendErr(burrito.WrapErrorf(err, "An error occurred while reading the mcfunction file %s", file))
								return
							}
							output, err := jsonte.ProcessMCFunction(string(bytes), object)
							if err != nil {
								appendErr(burrito.WrapErrorf(err, "An error occurred while processing the mcfunction file %s", file))
								return
							}
							fileName := strings.TrimSuffix(filepath.Base(file), filepath.Ext(file))
							filename := filepath.Dir(file) + "/" + fileName + ".mcfunction"
							fnCh <- fnRes{base: base, file: filename, data: output}
							if removeSrc {
								if err := os.Remove(file); err != nil {
									appendErr(burrito.WrapErrorf(err, "An error occurred while removing the mcfunction file %s", file))
								}
							}
						}(base, file)
					}
				}
			}
			go func() {
				wg.Wait()
				close(fnCh)
			}()
			fnResults := make([]fnRes, 0)
			for r := range fnCh {
				fnResults = append(fnResults, r)
			}
			sort.SliceStable(fnResults, func(i, j int) bool {
				if fnResults[i].base == fnResults[j].base {
					return fnResults[i].file < fnResults[j].file
				}
				return fnResults[i].base < fnResults[j].base
			})
			for _, r := range fnResults {
				filename := r.file
				if outFile != "" {
					var err error
					filename, err = filepath.Rel(r.base, filename)
					if err != nil {
						appendErr(burrito.WrapErrorf(err, "An error occurred while creating the output file name"))
						continue
					}
					filename = filepath.Join(outFile, r.base, filename)
					rel, err := filepath.Rel(outFile, filename)
					if err != nil {
						appendErr(burrito.WrapErrorf(err, "An error occurred while relativizing the output file name"))
						continue
					}
					utils.Logger.Infof("Writing file %s", filepath.Clean(rel))
				} else {
					utils.Logger.Infof("Writing file %s", filepath.Clean(filename))
				}
				if err := os.MkdirAll(filepath.Dir(filename), 0755); err != nil {
					appendErr(burrito.WrapErrorf(err, "An error occurred while creating the output directory %s", filepath.Dir(filename)))
					continue
				}
				if err := os.WriteFile(filename, []byte(r.data), 0644); err != nil {
					appendErr(burrito.WrapErrorf(err, "An error occurred while writing the output file %s", filename))
				}
			}
			if errs != nil {
				return errs
			}

			type langRes struct {
				base string
				file string
				data string
			}

			langCh := make(chan langRes)
			wg = sync.WaitGroup{}
			for base, files := range fileSets {
				for _, file := range files {
					if strings.HasSuffix(file, ".lang") {
						wg.Add(1)
						go func(base, file string) {
							defer wg.Done()
							sem <- struct{}{}
							defer func() { <-sem }()
							bytes, err := os.ReadFile(file)
							if err != nil {
								appendErr(burrito.WrapErrorf(err, "An error occurred while reading the mcfunction file %s", file))
								return
							}
							output, err := jsonte.ProcessLangFile(string(bytes), object)
							if err != nil {
								appendErr(burrito.WrapErrorf(err, "An error occurred while processing the mcfunction file %s", file))
								return
							}
							fileName := strings.TrimSuffix(filepath.Base(file), filepath.Ext(file))
							filename := filepath.Dir(file) + "/" + fileName + ".lang"
							langCh <- langRes{base: base, file: filename, data: output}
							if removeSrc {
								if err := os.Remove(file); err != nil {
									appendErr(burrito.WrapErrorf(err, "An error occurred while removing the mcfunction file %s", file))
								}
							}
						}(base, file)
					}
				}
			}
			go func() {
				wg.Wait()
				close(langCh)
			}()
			langResults := make([]langRes, 0)
			for r := range langCh {
				langResults = append(langResults, r)
			}
			sort.SliceStable(langResults, func(i, j int) bool {
				if langResults[i].base == langResults[j].base {
					return langResults[i].file < langResults[j].file
				}
				return langResults[i].base < langResults[j].base
			})
			for _, r := range langResults {
				filename := r.file
				if outFile != "" {
					var err error
					filename, err = filepath.Rel(r.base, filename)
					if err != nil {
						appendErr(burrito.WrapErrorf(err, "An error occurred while creating the output file name"))
						continue
					}
					filename = filepath.Join(outFile, r.base, filename)
					rel, err := filepath.Rel(outFile, filename)
					if err != nil {
						appendErr(burrito.WrapErrorf(err, "An error occurred while relativizing the output file name"))
						continue
					}
					utils.Logger.Infof("Writing file %s", filepath.Clean(rel))
				} else {
					utils.Logger.Infof("Writing file %s", filepath.Clean(filename))
				}
				if err := os.MkdirAll(filepath.Dir(filename), 0755); err != nil {
					appendErr(burrito.WrapErrorf(err, "An error occurred while creating the output directory %s", filepath.Dir(filename)))
					continue
				}
				if err := os.WriteFile(filename, []byte(r.data), 0644); err != nil {
					appendErr(burrito.WrapErrorf(err, "An error occurred while writing the output file %s", filename))
				}
			}
			if errs != nil {
				return errs
			}
			return nil
		},
	})
	app.Action(Action{
		Name:  "eval",
		Usage: "Evaluate a JSON expression or run a REPL",
		Function: func(args []string) error {
			functions.SetCacheAll(cacheAll)
			object, err := getScope(scope, -1)
			if err != nil {
				return burrito.WrapError(err, "An error occurred while reading the scope")
			}
			if len(args) == 0 {
				repl(object)
			} else {
				expression := strings.Join(args, " ")
				s := deque.Deque[*types.JsonObject]{}
				s.PushBack(object)
				value, err := jsonte.Eval(expression, s, "#")
				if err != nil {
					return burrito.WrapErrorf(err, "An error occurred while evaluating the expression")
				}
				fmt.Println(types.ToPrettyString(value.Value))
			}
			return nil
		},
	})
	app.Action(Action{
		Name:  "script",
		Usage: "Evaluate a JSON expression or run a REPL",
		Function: func(args []string) error {
			functions.SetCacheAll(cacheAll)
			object, err := getScope(scope, -1)
			if err != nil {
				return burrito.WrapError(err, "An error occurred while reading the scope")
			}
			if len(args) == 0 {
				return burrito.WrapErrorf(err, "No script file specified")
			} else {
				file, err := os.ReadFile(args[0])
				if err != nil {
					return burrito.WrapErrorf(err, "An error occurred while reading the script file")
				}
				file, err = json.ConvertToUTF8(file)
				if err != nil {
					return burrito.WrapErrorf(err, "An error occurred while reading the script file")
				}
				s := deque.Deque[*types.JsonObject]{}
				s.PushBack(object)
				value, err := jsonte.EvalScript(string(file), s, "#")
				if err != nil {
					return burrito.WrapErrorf(err, "An error occurred while running the script")
				}
				fmt.Println(types.ToPrettyString(value.Value))
			}
			return nil
		},
	})
	app.Action(Action{
		Name:  "version",
		Usage: "Print the version info",
		Function: func(args []string) error {
			fmt.Println("jsonte version " + version)
			if buildSource == "DEV" {
				fmt.Println("Development build")
			}
			if commit != "" {
				fmt.Println("Commit: " + commit)
			}
			if date != "" {
				fmt.Println("Built at " + date)
			}
			return nil
		},
	})
	app.Action(Action{
		Name:  "ipc",
		Usage: "Start an IPC server",
		Function: func(args []string) error {
			functions.SetCacheAll(cacheAll)
			object, err := getScope(scope, -1)
			if err != nil {
				return burrito.WrapError(err, "An error occurred while reading the scope")
			}
			return StartIPC(ipcName, object)
		},
	})
	app.Action(Action{
		Name:   "docgen",
		Hidden: true,
		Usage:  "Generate documentation",
		Function: func(args []string) error {
			return functions.GenerateDocs()
		},
	})
	err := app.Run(os.Args, func() {
		if debug && silent {
			utils.InitLogging(zap.DebugLevel)
			utils.Logger.Warn("--debug and --silent are mutually exclusive")
		} else if debug {
			utils.InitLogging(zap.DebugLevel)
		} else if silent {
			utils.InitLogging(zap.WarnLevel)
		} else {
			utils.InitLogging(zap.InfoLevel)
		}
		rand.Seed(seed)
	})
	if err != nil {
		if utils.Logger == nil {
			utils.InitLogging(zap.DebugLevel)
		}
		utils.Logger.Error(err)
		os.Exit(1)
	}
}

func getScope(scope []string, timeout int64) (*types.JsonObject, error) {
	assertionFiles := map[string]string{}
	result := types.NewJsonObject()
	for _, path := range scope {
		if strings.HasPrefix(path, "{") {
			json, err := types.ParseJsonObject([]byte(path))
			if err != nil {
				return types.NewJsonObject(), burrito.WrapErrorf(err, "An error occurred while parsing inline scope '%s'", path)
			}
			result = types.MergeObject(result, json, false, "#")
			continue
		}
		err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				if os.IsNotExist(err) {
					utils.Logger.Warnf("Skipping non-existent scope file '%s'", path)
					return nil
				}
				return burrito.WrapErrorf(err, "An error occurred while reading the scope file '%s'", path)
			}
			if !info.IsDir() {
				if strings.HasSuffix(path, ".json") {
					file, err := os.ReadFile(path)
					if err != nil {
						return burrito.WrapErrorf(err, "An error occurred while reading the scope file '%s'", path)
					}
					json, err := types.ParseJsonObject(file)
					if err != nil {
						return burrito.WrapErrorf(err, "An error occurred while parsing the scope file '%s'", path)
					}
					err = jsonte.VerifyReservedNames(json, path+"#/")
					result = types.MergeObject(result, json, false, "#")
				} else if strings.HasSuffix(path, ".assert") {
					file, err := os.ReadFile(path)
					if err != nil {
						return burrito.WrapErrorf(err, "An error occurred while reading the assertion file '%s'", path)
					}
					assertionFiles[path] = string(file)
				} else {
					utils.Logger.Debugf("Skipping non-scope file '%s'", path)
				}
			}
			return nil
		})
		if err != nil {
			return types.NewJsonObject(), burrito.WrapError(err, "An error occurred while reading the scope files")
		}
	}
	for path, file := range assertionFiles {
		err := jsonte.ProcessAssertionsFile(path, file, result, timeout)
		if err != nil {
			return types.NewJsonObject(), burrito.PassError(err)
		}
	}
	return result, nil
}

func getFileList(paths, include, exclude []string) (map[string][]string, error) {
	result := map[string][]string{}
	includes := make([]glob.Glob, 0)
	excludes := make([]glob.Glob, 0)
	pathRegex, err := regexp.Compile("[/\\\\]")
	if err != nil {
		// Should never happen as it's constant
		return nil, burrito.WrapError(err, "An error occurred while compiling the path separator regex")
	}
	for _, i := range include {
		g, err := glob.Compile(pathRegex.ReplaceAllString(i, "/"))
		if err != nil {
			return nil, burrito.WrapErrorf(err, "An error occurred while compiling the include pattern %s", i)
		}
		includes = append(includes, g)
	}
	for _, e := range exclude {
		g, err := glob.Compile(pathRegex.ReplaceAllString(e, "/"))
		if err != nil {
			return nil, burrito.WrapErrorf(err, "An error occurred while compiling the exclude pattern %s", e)
		}
		excludes = append(excludes, g)
	}
	for _, p := range paths {
		files := make([]string, 0)
		_, err := os.Stat(p)
		if err != nil {
			if os.IsNotExist(err) {
				return nil, burrito.WrappedErrorf("The path '%s' does not exist", p)
			}
			return nil, burrito.WrapErrorf(err, "An error occurred while reading the path %s", p)
		}
		err = filepath.Walk(p, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return burrito.WrapErrorf(err, "An error occurred while reading the path %s", p)
			}
			if !info.IsDir() {
				if !strings.HasSuffix(path, ".templ") && !strings.HasSuffix(path, ".modl") && !strings.HasSuffix(path, ".mcfunction") && !strings.HasSuffix(path, ".lang") {
					return nil
				}
				for _, g := range includes {
					if g.Match(pathRegex.ReplaceAllString(path, "/")) {
						files = append(files, path)
						return nil
					}
				}
				for _, g := range excludes {
					if g.Match(pathRegex.ReplaceAllString(path, "/")) {
						return nil
					}
				}
				if len(include) == 0 {
					files = append(files, path)
				}
			}
			return nil
		})
		if err != nil {
			return nil, burrito.WrapErrorf(err, "An error occurred while reading input files from %s", p)
		}
		result[p] = files
	}
	return result, nil
}

func repl(scope *types.JsonObject) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("> ")
	s := deque.Deque[*types.JsonObject]{}
	s.PushBack(scope)
	for {
		read, _ := reader.ReadString('\n')
		text := strings.TrimRight(read, "\n\r")
		if text == "exit" {
			break
		}
		eval, err := jsonte.Eval(text, s, "#/")
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(types.ToString(eval.Value))
			if !eval.VariableScope.IsEmpty() {
				s.PushBack(eval.VariableScope)
			}
		}
		fmt.Print("> ")
	}
}
