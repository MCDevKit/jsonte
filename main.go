package main

import (
	"bufio"
	"fmt"
	"github.com/MCDevKit/jsonte/jsonte"
	"github.com/MCDevKit/jsonte/jsonte/functions"
	"github.com/MCDevKit/jsonte/jsonte/utils"
	"github.com/gammazero/deque"
	"github.com/gobwas/glob"
	"go.uber.org/zap"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var (
	commit      string
	version     = "0.0.0"
	date        string
	buildSource = "DEV"
)

func main() {
	functions.Init()
	debug := false
	silent := false
	removeSrc := false
	minify := false
	scope := make([]string, 0)
	out := ""
	include := make([]string, 0)
	exclude := make([]string, 0)
	app := NewApp()
	app.BoolFlag(Flag{
		Name:  "debug",
		Usage: "Enable debug mode",
	}, &debug)
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
	app.Action(Action{
		Name: "compile",
		Function: func(args []string) error {
			outFile := ""
			if out != "" {
				stat, err := os.Stat(out)
				if err != nil {
					if os.IsNotExist(err) {
						err := os.MkdirAll(out, 0755)
						if err != nil {
							return utils.WrapError(err, "An error occurred while creating the output directory")
						}
					} else {
						return utils.WrapError(err, "An error occurred while reading the output file")
					}
				} else if !stat.IsDir() {
					return utils.WrappedErrorf("The output file %s is not a directory", out)
				}
				outFile, err = filepath.Abs(out)
			}
			object, err := getScope(scope)
			if err != nil {
				return utils.WrapError(err, "An error occurred while reading the scope")
			}
			fileSets, err := getFileList(args, include, exclude)
			if err != nil {
				return err
			}
			// Process modules
			modules := map[string]jsonte.JsonModule{}
			for base, files := range fileSets {
				for _, file := range files {
					if strings.HasSuffix(file, ".modl") {
						bytes, err := ioutil.ReadFile(file)
						if err != nil {
							return utils.WrapErrorf(err, "An error occurred while reading the module file %s", file)
						}
						module, err := jsonte.LoadModule(string(bytes))
						if err != nil {
							return utils.WrapErrorf(err, "An error occurred while loading the module file %s", file)
						}
						modules[module.Name] = module
						rel, err := filepath.Rel(base, file)
						if err != nil {
							return utils.WrapErrorf(err, "An error occurred while relativizing the output file name")
						}
						utils.Logger.Infof("Loaded module %s", rel)
						if removeSrc {
							err = os.Remove(file)
							if err != nil {
								return utils.WrapErrorf(err, "An error occurred while removing the module file %s", file)
							}
						}
					}
				}
			}
			// Process templates
			for base, files := range fileSets {
				for _, file := range files {
					if strings.HasSuffix(file, ".templ") {
						bytes, err := ioutil.ReadFile(file)
						if err != nil {
							return utils.WrapErrorf(err, "An error occurred while reading the template file %s", file)
						}
						template, err := jsonte.Process(strings.TrimSuffix(filepath.Base(file), filepath.Ext(file)), string(bytes), object, modules, -1)
						if err != nil {
							return utils.WrapErrorf(err, "An error occurred while processing the template file %s", file)
						}
						toString := utils.ToPrettyString
						if minify {
							toString = utils.ToString
						}
						for fileName, content := range template {
							filename := filepath.Dir(file) + "/" + fileName + ".json"
							if outFile != "" {
								filename, err = filepath.Rel(base, filename)
								if err != nil {
									return utils.WrapErrorf(err, "An error occurred while creating the output file name")
								}
								filename = filepath.Join(outFile, base, filename)
								rel, err := filepath.Rel(outFile, filename)
								if err != nil {
									return utils.WrapErrorf(err, "An error occurred while relativizing the output file name")
								}
								utils.Logger.Infof("Writing file %s", filepath.Clean(rel))
							} else {
								utils.Logger.Infof("Writing file %s", filepath.Clean(filename))
							}
							err = os.MkdirAll(filepath.Dir(filename), 0755)
							if err != nil {
								return utils.WrapErrorf(err, "An error occurred while creating the output directory %s", filepath.Dir(filename))
							}
							err = ioutil.WriteFile(filename, []byte(toString(content)), 0644)
							if err != nil {
								return utils.WrapErrorf(err, "An error occurred while writing the output file %s", filename)
							}
						}
						if removeSrc {
							err = os.Remove(file)
							if err != nil {
								return utils.WrapErrorf(err, "An error occurred while removing the template file %s", file)
							}
						}
					}
				}
			}
			//Process functions
			for base, files := range fileSets {
				for _, file := range files {
					if strings.HasSuffix(file, ".mcfunction") {
						bytes, err := ioutil.ReadFile(file)
						if err != nil {
							return utils.WrapErrorf(err, "An error occurred while reading the mcfunction file %s", file)
						}
						fileName := strings.TrimSuffix(filepath.Base(file), filepath.Ext(file))
						output, err := jsonte.ProcessMCFunction(string(bytes), object)
						if err != nil {
							return utils.WrapErrorf(err, "An error occurred while processing the mcfunction file %s", file)
						}
						filename := filepath.Dir(file) + "/" + fileName + ".mcfunction"
						if outFile != "" {
							filename, err = filepath.Rel(base, filename)
							if err != nil {
								return utils.WrapErrorf(err, "An error occurred while creating the output file name")
							}
							filename = filepath.Join(outFile, base, filename)
							rel, err := filepath.Rel(outFile, filename)
							if err != nil {
								return utils.WrapErrorf(err, "An error occurred while relativizing the output file name")
							}
							utils.Logger.Infof("Writing file %s", filepath.Clean(rel))
						} else {
							utils.Logger.Infof("Writing file %s", filepath.Clean(filename))
						}
						err = os.MkdirAll(filepath.Dir(filename), 0755)
						if err != nil {
							return utils.WrapErrorf(err, "An error occurred while creating the output directory %s", filepath.Dir(filename))
						}
						err = ioutil.WriteFile(filename, []byte(output), 0644)
						if err != nil {
							return utils.WrapErrorf(err, "An error occurred while writing the output file %s", filename)
						}
					}
				}
			}
			return nil
		},
	})
	app.Action(Action{
		Name:  "eval",
		Usage: "Evaluate a JSON expression or run a REPL",
		Function: func(args []string) error {
			object, err := getScope(scope)
			if err != nil {
				return utils.WrapError(err, "An error occurred while reading the scope")
			}
			if len(args) == 0 {
				repl(object)
			} else {
				expression := strings.Join(args, " ")
				s := deque.Deque[interface{}]{}
				s.PushBack(object)
				value, err := jsonte.Eval(expression, s, "#")
				if err != nil {
					return utils.WrapErrorf(err, "An error occurred while evaluating the expression")
				}
				fmt.Println(utils.ToPrettyString(value.Value))
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
	})
	if err != nil {
		utils.Logger.Error(err)
		os.Exit(1)
	}
}

func getScope(scope []string) (utils.JsonObject, error) {
	result := utils.JsonObject{}
	for _, path := range scope {
		err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				if os.IsNotExist(err) {
					utils.Logger.Warnf("Skipping non-existent scope file '%s'", path)
					return nil
				}
				return utils.WrapErrorf(err, "An error occurred while reading the scope file '%s'", path)
			}
			if !info.IsDir() && strings.HasSuffix(path, ".json") {
				file, err := ioutil.ReadFile(path)
				if err != nil {
					return utils.WrapErrorf(err, "An error occurred while reading the scope file '%s'", path)
				}
				json, err := utils.ParseJson(file)
				if err != nil {
					return utils.WrapErrorf(err, "An error occurred while parsing the scope file '%s'", path)
				}
				result = utils.MergeObject(result, json)
			}
			return nil
		})
		if err != nil {
			return nil, utils.WrapError(err, "An error occurred while reading the scope files")
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
		return nil, utils.WrapError(err, "An error occurred while compiling the path separator regex")
	}
	for _, i := range include {
		g, err := glob.Compile(pathRegex.ReplaceAllString(i, "/"))
		if err != nil {
			return nil, utils.WrapErrorf(err, "An error occurred while compiling the include pattern %s", i)
		}
		includes = append(includes, g)
	}
	for _, e := range exclude {
		g, err := glob.Compile(pathRegex.ReplaceAllString(e, "/"))
		if err != nil {
			return nil, utils.WrapErrorf(err, "An error occurred while compiling the exclude pattern %s", e)
		}
		excludes = append(excludes, g)
	}
	for _, p := range paths {
		files := make([]string, 0)
		_, err := os.Stat(p)
		if err != nil {
			if os.IsNotExist(err) {
				return nil, utils.WrappedErrorf("The path %s does not exist", p)
			}
			return nil, utils.WrapErrorf(err, "An error occurred while reading the path %s", p)
		}
		err = filepath.Walk(p, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return utils.WrapErrorf(err, "An error occurred while reading the path %s", p)
			}
			if !info.IsDir() {
				if !strings.HasSuffix(path, ".templ") && !strings.HasSuffix(path, ".modl") && !strings.HasSuffix(path, ".mcfunction") {
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
			return nil, utils.WrapErrorf(err, "An error occurred while reading input files from %s", p)
		}
		result[p] = files
	}
	return result, nil
}

func repl(scope utils.JsonObject) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("> ")
	for true {
		read, _ := reader.ReadString('\n')
		text := strings.TrimRight(read, "\n\r")
		if text == "exit" {
			break
		}
		s := deque.Deque[interface{}]{}
		s.PushBack(scope)
		eval, err := jsonte.Eval(text, s, "#/")
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(utils.ToString(eval.Value))
		}
		fmt.Print("> ")
	}
}
