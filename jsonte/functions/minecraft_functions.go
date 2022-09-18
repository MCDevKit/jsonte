package functions

import (
	"archive/zip"
	"fmt"
	"github.com/MCDevKit/jsonte/jsonte/utils"
	"github.com/stirante/jsonc"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

func RegisterMinecraftFunctions() {
	const group = "minecraft"
	RegisterGroup(Group{
		Name:    group,
		Title:   "Minecraft functions",
		Summary: "Functions specific to Minecraft.",
	})
	RegisterFunction(JsonFunction{
		Group:    group,
		Name:     "getMinecraftInstallDir",
		Body:     getMinecraftInstallDir,
		IsUnsafe: true,
		Docs: Docs{
			Summary: "Returns a path to the folder with Minecraft app. The value is cached after the first usage.\n\n**This function works only on Windows with installed Minecraft Bedrock.**",
			Example: `
<code>
{
  "$template": {
    "$comment": "The field below will most likely be 'C:\Program Files\WindowsApps\Microsoft.MinecraftUWP_<Minecraft version>__8wekyb3d8bbwe'",
    "test": "{{getMinecraftInstallDir()}}"
  }
}
</code>`,
		},
	})
	RegisterFunction(JsonFunction{
		Group:    group,
		Name:     "getLatestBPFile",
		Body:     getLatestBPFile,
		IsUnsafe: true,
		Docs: Docs{
			Summary: "Returns a path to the latest behavior pack file.",
			Arguments: []Argument{
				{
					Name:    "path",
					Summary: "The path to the file inside behavior pack.",
				},
			},
			Example: `
<code>
{
  "$template": {
    "$comment": "The field below will most likely be 'C:\Program Files\WindowsApps\Microsoft.MinecraftUWP_<Minecraft version>__8wekyb3d8bbwe\data\behavior_packs\vanilla_1.18.10\entities\axolotl.json'",
    "test": "{{getLatestBPFile('entities/axolotl.json')}}"
  }
}
</code>`,
		},
	})
	RegisterFunction(JsonFunction{
		Group:    group,
		Name:     "getLatestRPFile",
		Body:     getLatestRPFile,
		IsUnsafe: true,
		Docs: Docs{
			Summary: "Returns a path to the latest resource pack file.",
			Arguments: []Argument{
				{
					Name:    "path",
					Summary: "The path to the file inside resource pack.",
				},
			},
			Example: `
<code>
{
  "$template": {
    "$comment": "The field below will most likely be 'C:\Program Files\WindowsApps\Microsoft.MinecraftUWP_<Minecraft version>__8wekyb3d8bbwe\data\resource_packs\vanilla_1.18.10\textures\entity\axolotl\axolotl_wild.png'",
    "test": "{{getLatestRPFile('textures/entity/axolotl/axolotl_wild.png')}}"
  }
}
</code>`,
		},
	})
	RegisterFunction(JsonFunction{
		Group: group,
		Name:  "listLatestRPFiles",
		Body:  listLatestRPFiles,
		Docs: Docs{
			Summary: "Returns an array of paths to the latest files in resource pack within given path.",
			Arguments: []Argument{
				{
					Name:    "path",
					Summary: "The path to the directory inside resource pack.",
				},
			},
			Example: `
<code>
{
  "$template": {
    "test": "{{listLatestRPFiles('entity')}}"
  }
}
</code>`,
		},
	})
	RegisterFunction(JsonFunction{
		Group: group,
		Name:  "listLatestBPFiles",
		Body:  listLatestBPFiles,
		Docs: Docs{
			Summary: "Returns an array of paths to the latest files in behavior pack within given path.",
			Arguments: []Argument{
				{
					Name:    "path",
					Summary: "The path to the directory inside behavior pack.",
				},
			},
			Example: `
<code>
{
  "$template": {
    "test": "{{listLatestBPFiles('entities')}}"
  }
}
</code>`,
		},
	})
}

var installDir = ""
var VanillaRpUUID = "0575c61f-a5da-4b7f-9961-ffda2908861e"
var VanillaBpUUID = "fe9f8597-5454-481a-8730-8d070a8e2e58"

var rpFiles = utils.NavigableMap[string, string]{}
var bpFiles = utils.NavigableMap[string, string]{}

func getMinecraftInstallDir() (string, error) {
	if runtime.GOOS != "windows" {
		return "", utils.WrappedErrorf("This function works only on Windows")
	}
	if installDir == "" {
		output, err := exec.Command("powershell", "(Get-AppxPackage -Name Microsoft.MinecraftUWP).InstallLocation").Output()
		if err != nil {
			return "", utils.WrapErrorf(err, "An error occurred while getting Minecraft install directory")
		}
		installDir = strings.Trim(string(output), "\r\n \t")
	}
	return installDir, nil
}

func getLatestBPFile(p string) (string, error) {
	if bpFiles.IsEmpty() {
		bp, err := findPackVersions(true, VanillaBpUUID)
		if err != nil {
			return "", utils.WrapErrorf(err, "An error occurred while reading behavior packs")
		}
		bpFiles = bp
	}
	return getLatestFile(p, bpFiles)
}

func getLatestRPFile(p string) (string, error) {
	if rpFiles.IsEmpty() {
		rp, err := findPackVersions(false, VanillaRpUUID)
		if err != nil {
			return "", utils.WrapErrorf(err, "An error occurred while reading resource packs")
		}
		rpFiles = rp
	}
	return getLatestFile(p, rpFiles)
}

func listLatestRPFiles(p string) (utils.JsonArray, error) {
	if rpFiles.IsEmpty() {
		rp, err := findPackVersions(false, VanillaRpUUID)
		if err != nil {
			return utils.JsonArray{}, utils.WrapErrorf(err, "An error occurred while reading resource packs")
		}
		rpFiles = rp
	}
	return listLatestFiles(p, rpFiles)
}

func listLatestBPFiles(p string) (utils.JsonArray, error) {
	if bpFiles.IsEmpty() {
		bp, err := findPackVersions(true, VanillaBpUUID)
		if err != nil {
			return utils.JsonArray{}, utils.WrapErrorf(err, "An error occurred while reading behavior packs")
		}
		bpFiles = bp
	}
	return listLatestFiles(p, bpFiles)
}

func listLatestFiles(p string, m utils.NavigableMap[string, string]) (utils.JsonArray, error) {
	result := map[string]string{}
	keys := m.Keys()
	for i := len(keys) - 1; i >= 0; i-- {
		s := path.Join(m.Get(keys[i]), p)
		_, err := os.Stat(s)
		if err != nil {
			if os.IsNotExist(err) {
				continue
			} else {
				return nil, utils.WrapErrorf(err, "An error occurred while reading file %s", p)
			}
		}
		err = filepath.Walk(s, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() {
				rel, err := filepath.Rel(s, path)
				if err != nil {
					return err
				}
				if _, ok := result[rel]; !ok {
					result[rel] = path
				}
			}
			return nil
		})
		if err != nil {
			return nil, utils.WrapErrorf(err, "An error occurred while reading file %s", p)
		}
	}
	arr := make(utils.JsonArray, len(result))
	i := 0
	for _, v := range result {
		arr[i] = v
		i++
	}
	return arr, nil
}

func getLatestFile(p string, m utils.NavigableMap[string, string]) (string, error) {
	keys := m.Keys()
	for i := len(keys) - 1; i >= 0; i-- {
		s := path.Join(m.Get(keys[i]), p)
		_, err := os.Stat(s)
		if err != nil {
			if os.IsNotExist(err) {
				continue
			} else {
				return "", utils.WrapErrorf(err, "An error occurred while reading file %s", p)
			}
		}
		return s, nil
	}
	return "", utils.WrapErrorf(os.ErrNotExist, "File '%s' does not exist", p)
}

// From https://stackoverflow.com/a/24792688/6459649
func unzip(src, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return utils.WrapErrorf(err, "An error occurred while reading zip file %s", src)
	}

	err = os.MkdirAll(dest, 0755)
	if err != nil {
		return utils.WrapErrorf(err, "An error occurred while creating directory %s", dest)
	}

	// Closure to address file descriptors issue with all the deferred .Close() methods
	extractAndWriteFile := func(f *zip.File) error {
		rc, err := f.Open()
		if err != nil {
			return utils.WrapErrorf(err, "An error occurred while opening file %s", f.Name)
		}

		path := filepath.Join(dest, f.Name)

		// Check for ZipSlip (Directory traversal)
		if !strings.HasPrefix(path, filepath.Clean(dest)+string(os.PathSeparator)) {
			return fmt.Errorf("illegal file path: %s", path)
		}

		if f.FileInfo().IsDir() {
			err := os.MkdirAll(path, f.Mode())
			if err != nil {
				return utils.WrapErrorf(err, "An error occurred while creating directory %s", path)
			}
		} else {
			err := os.MkdirAll(filepath.Dir(path), f.Mode())
			if err != nil {
				return utils.WrapErrorf(err, "An error occurred while creating directory %s", filepath.Dir(path))
			}
			f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return utils.WrapErrorf(err, "An error occurred while creating file %s", path)
			}

			_, err = io.Copy(f, rc)
			if err != nil {
				return utils.WrapErrorf(err, "An error occurred while writing file %s", path)
			}

			err = f.Close()
			if err != nil {
				return utils.WrapErrorf(err, "An error occurred while closing file %s", path)
			}

			err = rc.Close()
			if err != nil {
				return utils.WrapErrorf(err, "An error occurred while closing file %s", f.Name)
			}
		}
		return nil
	}

	for _, f := range r.File {
		err := extractAndWriteFile(f)
		if err != nil {
			return utils.WrapErrorf(err, "An error occurred while extracting file %s", f.Name)
		}
	}

	err = r.Close()
	if err != nil {
		return utils.WrapErrorf(err, "An error occurred while closing zip file %s", src)
	}

	return nil
}

func findPackVersions(isBp bool, uuid string) (utils.NavigableMap[string, string], error) {
	versions := utils.NewNavigableMap[string, string]()
	installDir, err := getMinecraftInstallDir()
	if err != nil {
		url := "https://aka.ms/resourcepacktemplate"
		outName := "resource_pack_template.zip"
		dirName := "RP"
		if isBp {
			url = "https://aka.ms/behaviorpacktemplate"
			outName = "behavior_pack_template.zip"
			dirName = "BP"
		}

		base := "packs"
		if utils.CacheDir != "" {
			base = utils.CacheDir
		}

		stat, err := os.Stat(filepath.Join(base, dirName))
		if err == nil && stat.IsDir() {
			versions.Put("1.0.0", filepath.Join(base, dirName))
			return versions, nil
		}
		utils.Logger.Infof("Downloading %s", url)

		err = os.MkdirAll(base, 0755)
		if err != nil && !os.IsExist(err) {
			return versions, utils.WrapErrorf(err, "An error occurred while creating cache directory")
		}
		out, err := os.Create(path.Join(base, outName))
		if err != nil {
			return versions, utils.WrapErrorf(err, "An error occurred while creating file %s", outName)
		}
		resp, err := http.Get(url)
		if err != nil {
			return versions, utils.WrapErrorf(err, "An error occurred while downloading %s", url)
		}
		_, err = io.Copy(out, resp.Body)
		if err != nil {
			return versions, utils.WrapErrorf(err, "An error occurred while downloading %s", url)
		}
		err = out.Close()
		if err != nil {
			return versions, utils.WrapErrorf(err, "An error occurred while downloading %s", url)
		}
		err = resp.Body.Close()
		if err != nil {
			return versions, utils.WrapErrorf(err, "An error occurred while downloading %s", url)
		}

		err = unzip(out.Name(), path.Join(base, dirName))
		if err != nil {
			return versions, utils.WrapErrorf(err, "An error occurred while extracting %s", outName)
		}
		err = os.Remove(out.Name())
		if err != nil {
			return versions, utils.WrapErrorf(err, "An error occurred while removing %s", outName)
		}

		versions.Put("0.0.0", path.Join(base, dirName))
		return versions, err
	}
	packDir := path.Join(installDir, "data", "behavior_packs")
	if !isBp {
		packDir = path.Join(installDir, "data", "resource_packs")
	}
	dir, err := os.ReadDir(packDir)
	if err != nil {
		return versions, utils.WrapErrorf(err, "Failed to read pack directory")
	}
	for _, d := range dir {
		p := path.Join(packDir, d.Name())
		if d.IsDir() {
			f, err := ioutil.ReadFile(path.Join(p, "manifest.json"))
			if err != nil {
				if os.IsNotExist(err) {
					continue
				}
				return versions, utils.WrapErrorf(err, "Failed to read manifest.json in %s", p)
			}
			var manifest map[string]interface{}
			err = jsonc.Unmarshal(f, &manifest)
			if err != nil {
				return versions, utils.WrapErrorf(err, "Failed to parse manifest.json in %s", p)
			}
			if header, ok := manifest["header"]; ok {
				if headerMap, ok := header.(map[string]interface{}); ok {
					if headerMap["uuid"] != uuid {
						continue
					}
					if version, ok := headerMap["version"]; ok {
						if versionArr, ok := version.([]interface{}); ok {
							array := utils.ParseSemverArray(versionArr)
							versions.Put(array.String(), p)
						}
					}
				}
			}
		}
	}
	versions.Sort(func(a, b string) int {
		aVer, err := utils.ParseSemverString(a)
		if err != nil {
			return 0
		}
		bVer, err := utils.ParseSemverString(b)
		if err != nil {
			return 0
		}
		return aVer.CompareTo(bVer)
	})
	return versions, nil
}
