package functions

import (
	"archive/zip"
	"encoding/json"
	"github.com/Bedrock-OSS/go-burrito/burrito"
	json2 "github.com/MCDevKit/jsonte/jsonte/json"
	"github.com/MCDevKit/jsonte/jsonte/safeio"
	"github.com/MCDevKit/jsonte/jsonte/types"
	"github.com/MCDevKit/jsonte/jsonte/utils"
	"github.com/paul-mannino/go-fuzzywuzzy"
	"io"
	"math"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"slices"
	"strings"
	"time"
)

var releases []map[string]interface{}

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
				{
					Name:     "shouldFail",
					Optional: true,
					Summary:  "Whether to throw an error and stop compilation or return null when the file is not found.",
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
		Name:     "getLatestBPFile",
		Body:     getLatestBPFileOrFail,
		IsUnsafe: true,
	})
	RegisterFunction(JsonFunction{
		Group:    group,
		Name:     "getBPFile",
		Body:     getBPFile,
		IsUnsafe: true,
		Docs: Docs{
			Summary: "Returns a path to the latest behavior pack file.",
			Arguments: []Argument{
				{
					Name:    "path",
					Summary: "The path to the file inside behavior pack.",
				},
				{
					Name:    "version",
					Summary: "The version of Minecraft vanilla files to search for. If not specified, the latest version will be used.",
				},
				{
					Name:     "shouldFail",
					Optional: true,
					Summary:  "Whether to throw an error and stop compilation or return null when the file is not found.",
				},
			},
			Example: `
<code>
{
  "$template": {
    "$comment": "The field below will most likely be 'C:\Program Files\WindowsApps\Microsoft.MinecraftUWP_<Minecraft version>__8wekyb3d8bbwe\data\behavior_packs\vanilla_1.18.10\entities\axolotl.json'",
    "test": "{{getBPFile('entities/axolotl.json', semver('1.18.10'))}}"
  }
}
</code>`,
		},
	})
	RegisterFunction(JsonFunction{
		Group:    group,
		Name:     "getBPFile",
		Body:     getBPFileOrFail,
		IsUnsafe: true,
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
				{
					Name:     "shouldFail",
					Optional: true,
					Summary:  "Whether to throw an error and stop compilation or return null when the file is not found.",
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
		Group:    group,
		Name:     "getLatestRPFile",
		Body:     getLatestRPFileOrFail,
		IsUnsafe: true,
	})
	RegisterFunction(JsonFunction{
		Group:    group,
		Name:     "getRPFile",
		Body:     getRPFile,
		IsUnsafe: true,
		Docs: Docs{
			Summary: "Returns a path to the latest resource pack file.",
			Arguments: []Argument{
				{
					Name:    "path",
					Summary: "The path to the file inside resource pack.",
				},
				{
					Name:    "version",
					Summary: "The version of Minecraft vanilla files to search for.",
				},
				{
					Name:     "shouldFail",
					Optional: true,
					Summary:  "Whether to throw an error and stop compilation or return null when the file is not found.",
				},
			},
			Example: `
<code>
{
  "$template": {
    "$comment": "The field below will most likely be 'C:\Program Files\WindowsApps\Microsoft.MinecraftUWP_<Minecraft version>__8wekyb3d8bbwe\data\resource_packs\vanilla_1.18.10\textures\entity\axolotl\axolotl_wild.png'",
    "test": "{{getRPFile('textures/entity/axolotl/axolotl_wild.png', semver('1.18.10'))}}"
  }
}
</code>`,
		},
	})
	RegisterFunction(JsonFunction{
		Group:    group,
		Name:     "getRPFile",
		Body:     getRPFileOrFail,
		IsUnsafe: true,
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
	RegisterFunction(JsonFunction{
		Group: group,
		Name:  "listRPFiles",
		Body:  listRPFiles,
		Docs: Docs{
			Summary: "Returns an array of paths to the latest files in resource pack within given path.",
			Arguments: []Argument{
				{
					Name:    "path",
					Summary: "The path to the directory inside resource pack.",
				},
				{
					Name:    "version",
					Summary: "The version of Minecraft vanilla files to search for.",
				},
			},
			Example: `
<code>
{
  "$template": {
    "test": "{{listRPFiles('entity', semver('1.17.30'))}}"
  }
}
</code>`,
		},
	})
	RegisterFunction(JsonFunction{
		Group: group,
		Name:  "listBPFiles",
		Body:  listBPFiles,
		Docs: Docs{
			Summary: "Returns an array of paths to the latest files in behavior pack within given path.",
			Arguments: []Argument{
				{
					Name:    "path",
					Summary: "The path to the directory inside behavior pack.",
				},
				{
					Name:    "version",
					Summary: "The version of Minecraft vanilla files to search for.",
				},
			},
			Example: `
<code>
{
  "$template": {
    "test": "{{listBPFiles('entities', semver('1.17.30'))}}"
  }
}
</code>`,
		},
	})
	RegisterFunction(JsonFunction{
		Group: group,
		Name:  "getItemInfo",
		Body:  getItemInfo,
		Docs: Docs{
			Summary: "Returns info in an object based on item ID. Uses https://github.com/stirante/minecraft-item-db/blob/main/items.json",
			Arguments: []Argument{
				{
					Name:    "id",
					Summary: "The item ID.",
				},
				{
					Name:     "metadata",
					Optional: true,
					Summary:  "The item data value.",
				},
			},
			Example: `
<code>
{
  "$template": {
    "test": "{{getItemInfo('stone', 0)}}"
  }
}
</code>`,
		},
	})
	RegisterFunction(JsonFunction{
		Group: group,
		Name:  "getItemInfo",
		Body:  getItemInfo1,
	})
	RegisterFunction(JsonFunction{
		Group: group,
		Name:  "findItemInfoById",
		Body:  findItemInfoById,
		Docs: Docs{
			Summary: "Returns info in an object based on fuzzy matching item ID. Uses https://github.com/stirante/minecraft-item-db/blob/main/items.json",
			Arguments: []Argument{
				{
					Name:    "id",
					Summary: "The item ID.",
				},
				{
					Name:     "metadata",
					Optional: true,
					Summary:  "The item data value.",
				},
			},
			Example: `
<code>
{
  "$template": {
// {"id":"minecraft:blue_glazed_terracotta","legacyId":231,"metadata":0,"maxDurability":0,"damage":0,"armor":0,"maxStackSize":64,"tags":[],"category":"construction","nutrition":0,"fuelDuration":0,"aliases":["minecraft:glazedTerracotta.blue"],"nameKey":"tile.glazedTerracotta.blue.name","langName":"Blue Glazed Terracotta"}
    "test": "{{findItemInfoById('blue_terracotta', 0)}}"
  }
}
</code>`,
		},
	})
	RegisterFunction(JsonFunction{
		Group: group,
		Name:  "findItemInfoById",
		Body:  findItemInfoById1,
	})
	RegisterFunction(JsonFunction{
		Group: group,
		Name:  "findItemInfoByName",
		Body:  findItemInfoByName,
		Docs: Docs{
			Summary: "Returns info in an object based on fuzzy matching item English name. Uses https://github.com/stirante/minecraft-item-db/blob/main/items.json",
			Arguments: []Argument{
				{
					Name:    "name",
					Summary: "The item English name.",
				},
			},
			Example: `
<code>
{
  "$template": {
// {"id":"minecraft:blue_glazed_terracotta","legacyId":231,"metadata":0,"maxDurability":0,"damage":0,"armor":0,"maxStackSize":64,"tags":[],"category":"construction","nutrition":0,"fuelDuration":0,"aliases":["minecraft:glazedTerracotta.blue"],"nameKey":"tile.glazedTerracotta.blue.name","langName":"Blue Glazed Terracotta"}
    "test": "{{findItemInfoByName('blue terracotta')}}"
  }
}
</code>`,
		},
	})
	RegisterFunction(JsonFunction{
		Group: group,
		Name:  "getAllItems",
		Body:  getAllItems,
		Docs: Docs{
			Summary: "Returns a list of vanilla item infos. Uses https://github.com/stirante/minecraft-item-db/blob/main/items.json",
			Example: `
<code>
{
  "$template": {
	// This will be a list of all vanilla item infos
    "test": "{{getAllItems()}}"
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
var itemInfos = utils.NavigableMap[string, map[int]interface{}]{}
var itemInfosByName = utils.NavigableMap[string, interface{}]{}

type CachedRequest struct {
	Response map[string]interface{}
	Next     string
}

var requestCache = map[string]CachedRequest{}

func getMinecraftInstallDir() (*types.JsonString, error) {
	if ServerMode {
		return nil, burrito.WrappedErrorf("This function works only in client mode")
	}
	if runtime.GOOS != "windows" {
		return nil, burrito.WrappedErrorf("This function works only on Windows")
	}
	if installDir == "" {
		output, err := safeio.Resolver.ExecCommand("powershell", "(Get-AppxPackage -Name Microsoft.MinecraftUWP).InstallLocation")
		if err != nil {
			return nil, burrito.WrapErrorf(err, "An error occurred while getting Minecraft install directory")
		}
		installDir = strings.Trim(string(output), "\r\n \t")
	}
	return types.NewString(installDir), nil
}

func getLatestBPFileOrFail(p *types.JsonString) (types.JsonType, error) {
	return getBPFile(p, nil, types.True())
}

func getLatestBPFile(p *types.JsonString, shouldFail *types.JsonBool) (types.JsonType, error) {
	return getBPFile(p, nil, shouldFail)
}

func getBPFileOrFail(p *types.JsonString, version *types.Semver) (types.JsonType, error) {
	return getBPFile(p, version, types.True())
}

func getBPFile(p *types.JsonString, version *types.Semver, shouldFail *types.JsonBool) (types.JsonType, error) {
	if bpFiles.IsEmpty() {
		bp, err := findPackVersions(true, VanillaBpUUID, version)
		if err != nil {
			return nil, burrito.WrapErrorf(err, "An error occurred while reading behavior packs")
		}
		bpFiles = bp
	}
	file, err := getLatestFile(p.StringValue(), bpFiles, version)
	if !shouldFail.BoolValue() && err != nil && burrito.AsBurritoError(err).HasTag(os.ErrNotExist.Error()) {
		return types.Null, nil
	}
	return file, err
}

func getLatestRPFileOrFail(p *types.JsonString) (types.JsonType, error) {
	return getRPFile(p, nil, types.True())
}

func getLatestRPFile(p *types.JsonString, shouldFail *types.JsonBool) (types.JsonType, error) {
	return getRPFile(p, nil, shouldFail)
}

func getRPFileOrFail(p *types.JsonString, version *types.Semver) (types.JsonType, error) {
	return getRPFile(p, version, types.True())
}

func getRPFile(p *types.JsonString, version *types.Semver, shouldFail *types.JsonBool) (types.JsonType, error) {
	if rpFiles.IsEmpty() {
		rp, err := findPackVersions(false, VanillaRpUUID, version)
		if err != nil {
			return nil, burrito.WrapErrorf(err, "An error occurred while reading resource packs")
		}
		rpFiles = rp
	}
	file, err := getLatestFile(p.StringValue(), rpFiles, version)
	if !shouldFail.BoolValue() && err != nil && burrito.AsBurritoError(err).HasTag(os.ErrNotExist.Error()) {
		return types.Null, nil
	}
	return file, err
}

func listLatestRPFiles(p *types.JsonString) (*types.JsonArray, error) {
	return listRPFiles(p, nil)
}

func listRPFiles(p *types.JsonString, version *types.Semver) (*types.JsonArray, error) {
	if rpFiles.IsEmpty() {
		rp, err := findPackVersions(false, VanillaRpUUID, version)
		if err != nil {
			return types.NewJsonArray(), burrito.WrapErrorf(err, "An error occurred while reading resource packs")
		}
		rpFiles = rp
	}
	return listLatestFiles(p.StringValue(), rpFiles, version)
}

func listLatestBPFiles(p *types.JsonString) (*types.JsonArray, error) {
	return listBPFiles(p, nil)
}

func listBPFiles(p *types.JsonString, version *types.Semver) (*types.JsonArray, error) {
	if bpFiles.IsEmpty() {
		bp, err := findPackVersions(true, VanillaBpUUID, version)
		if err != nil {
			return types.NewJsonArray(), burrito.WrapErrorf(err, "An error occurred while reading behavior packs")
		}
		bpFiles = bp
	}
	return listLatestFiles(p.StringValue(), bpFiles, version)
}

func listLatestFiles(p string, m utils.NavigableMap[string, string], version *types.Semver) (*types.JsonArray, error) {
	result := map[string]string{}
	keys := m.Keys()
	for i := len(keys) - 1; i >= 0; i-- {
		if version != nil && !version.IsEmpty() {
			ver, err := types.ParseSemverString(keys[i])
			if err != nil {
				continue
			}
			than, _ := ver.LessThan(version)
			if !than && !ver.Equals(version) {
				continue
			}
		}
		s := path.Join(m.Get(keys[i]), p)
		_, err := safeio.Resolver.Stat(s)
		if err != nil {
			if os.IsNotExist(err) {
				continue
			} else {
				return types.NewJsonArray(), burrito.WrapErrorf(err, "An error occurred while reading file %s", p)
			}
		}
		recursive, err := safeio.Resolver.OpenDirRecursive(s)
		if err != nil {
			return types.NewJsonArray(), burrito.WrapErrorf(err, "An error occurred while reading file %s", p)
		}
		for _, f := range recursive {
			joined := filepath.Join(s, f)
			rel, err := filepath.Rel(s, joined)
			if err != nil {
				return types.NewJsonArray(), err
			}
			if _, ok := result[rel]; !ok {
				result[rel] = joined
			}
		}
	}
	arr := make([]types.JsonType, len(result))
	i := 0
	for _, v := range result {
		arr[i] = types.NewString(v)
		i++
	}
	return &types.JsonArray{Value: arr}, nil
}

func getLatestFile(p string, m utils.NavigableMap[string, string], version *types.Semver) (*types.JsonString, error) {
	keys := m.Keys()
	for i := len(keys) - 1; i >= 0; i-- {
		if version != nil && !version.IsEmpty() {
			ver, err := types.ParseSemverString(keys[i])
			if err != nil {
				continue
			}
			than, _ := ver.LessThan(version)
			if !than && !ver.Equals(version) {
				continue
			}
		}
		s := path.Join(m.Get(keys[i]), p)
		_, err := safeio.Resolver.Stat(s)
		if err != nil {
			if os.IsNotExist(err) {
				continue
			} else {
				return nil, burrito.WrapErrorf(err, "An error occurred while reading file %s", p)
			}
		}
		return types.NewString(s), nil
	}
	err := burrito.AsBurritoError(burrito.WrapErrorf(os.ErrNotExist, "File '%s' does not exist", p))
	err.AddTag(os.ErrNotExist.Error())
	return nil, err
}

// From https://stackoverflow.com/a/24792688/6459649
func unzip(src, dest string) error {
	stat, err := safeio.Resolver.Stat(src)
	if err != nil {
		return burrito.WrapErrorf(err, "An error occurred while reading file %s", src)
	}
	open, err := safeio.Resolver.Open(src)
	if err != nil {
		return burrito.WrapErrorf(err, "An error occurred while opening file %s", src)
	}
	r, err := zip.NewReader(open, stat.Size())
	if err != nil {
		return burrito.WrapErrorf(err, "An error occurred while reading zip file %s", src)
	}

	err = safeio.Resolver.MkdirAll(dest)
	if err != nil {
		return burrito.WrapErrorf(err, "An error occurred while creating directory %s", dest)
	}

	// Closure to address file descriptors issue with all the deferred .Close() methods
	extractAndWriteFile := func(f *zip.File) error {
		rc, err := f.Open()
		if err != nil {
			return burrito.WrapErrorf(err, "An error occurred while opening file %s", f.Name)
		}

		name := f.Name[strings.Index(f.Name, "/")+1:]
		if strings.HasPrefix(name, "behavior_pack") {
			name = strings.Replace(name, "behavior_pack", "BP", 1)
		} else if strings.HasPrefix(name, "resource_pack") {
			name = strings.Replace(name, "resource_pack", "RP", 1)
		} else {
			return nil
		}

		path := filepath.Join(dest, name)

		// Check for ZipSlip (Directory traversal)
		if !strings.HasPrefix(path, filepath.Clean(dest)+string(os.PathSeparator)) {
			return burrito.WrappedErrorf("Illegal file path: %s", path)
		}

		if f.FileInfo().IsDir() {
			err := safeio.Resolver.MkdirAll(path)
			if err != nil {
				return burrito.WrapErrorf(err, "An error occurred while creating directory %s", path)
			}
		} else {
			err := safeio.Resolver.MkdirAll(filepath.Dir(path))
			if err != nil {
				return burrito.WrapErrorf(err, "An error occurred while creating directory %s", filepath.Dir(path))
			}
			f, err := safeio.Resolver.Create(path)
			if err != nil {
				return burrito.WrapErrorf(err, "An error occurred while creating file %s", path)
			}

			_, err = io.Copy(f, rc)
			if err != nil {
				return burrito.WrapErrorf(err, "An error occurred while writing file %s", path)
			}

			err = f.Close()
			if err != nil {
				return burrito.WrapErrorf(err, "An error occurred while closing file %s", path)
			}

			err = rc.Close()
			if err != nil {
				return burrito.WrapErrorf(err, "An error occurred while closing file %s", path)
			}
		}
		return nil
	}

	for _, f := range r.File {
		err := extractAndWriteFile(f)
		if err != nil {
			return burrito.WrapErrorf(err, "An error occurred while extracting file %s", f.Name)
		}
	}

	err = open.Close()
	if err != nil {
		return burrito.WrapErrorf(err, "An error occurred while closing zip file %s", src)
	}

	return nil
}

func findPackVersions(isBp bool, uuid string, version *types.Semver) (utils.NavigableMap[string, string], error) {
	versions := utils.NewNavigableMap[string, string]()
	installDir, err := getMinecraftInstallDir()
	if err != nil {
		outName := "packs.zip"
		dirName := "RP"
		if isBp {
			dirName = "BP"
		}

		base := "packs"
		if json2.CacheDir != "" {
			base = json2.CacheDir
		}

		stat, err := safeio.Resolver.Stat(base)
		if err == nil && stat.IsDir() {
			// list folders in base
			dir, err := safeio.Resolver.OpenDir(base)
			if err != nil {
				return versions, burrito.WrapErrorf(err, "An error occurred while opening directory %s", base)
			}
			for _, d := range dir {
				stat, err = safeio.Resolver.Stat(path.Join(base, d))
				if err == nil && stat.IsDir() {
					stat, err = safeio.Resolver.Stat(path.Join(base, d, dirName))
					if err == nil && stat.IsDir() {
						_, err = types.ParseSemverString(d)
						if err != nil {
							continue
						}
						versions.Put(d, path.Join(base, d, dirName))
					}
				}
			}
			if !versions.IsEmpty() && (version == nil || version.IsEmpty() || versions.ContainsMatchingKey(func(s string) bool {
				ver, err2 := types.ParseSemverString(s)
				if err2 != nil {
					return false
				}
				return ver.Major == version.Major && ver.Minor == version.Minor
			})) {
				return versions, nil
			}
		}

		var release map[string]interface{}
		if version == nil || version.IsEmpty() {
			url := "https://api.github.com/repos/Mojang/bedrock-samples/releases/latest"
			utils.Logger.Infof("Resolving %s", url)
			if _, ok := requestCache[url]; !ok {
				resp, _, err := safeio.Resolver.HttpGet(url)
				if err != nil {
					return versions, burrito.WrapErrorf(err, "An error occurred while resolving %s", url)
				}
				err = json.NewDecoder(resp).Decode(&release)
				if err != nil {
					return versions, burrito.WrapErrorf(err, "An error occurred while parsing %s", url)
				}
				err = resp.Close()
				if err != nil {
					return versions, burrito.WrapErrorf(err, "An error occurred while closing %s", url)
				}
				requestCache[url] = CachedRequest{Response: release, Next: ""}
			} else {
				release = requestCache[url].Response
			}
		} else {
			url := "https://api.github.com/repos/Mojang/bedrock-samples/releases"
			if len(releases) == 0 {
				var fetchedURLs []string
				counter := 0
				var fetchedReleases []map[string]interface{}
			outer:
				for {
					counter++
					if counter > 10 {
						utils.Logger.Error("Too many redirects")
						break
					}
					fetchedURLs = append(fetchedURLs, url)
					utils.Logger.Infof("Resolving %s", url)
					resp, header, err := safeio.Resolver.HttpGet(url)
					if err != nil {
						return versions, burrito.WrapErrorf(err, "An error occurred while resolving %s", url)
					}
					err = json.NewDecoder(resp).Decode(&fetchedReleases)
					if err != nil {
						return versions, burrito.WrapErrorf(err, "An error occurred while parsing %s", url)
					}
					err = resp.Close()
					if err != nil {
						return versions, burrito.WrapErrorf(err, "An error occurred while closing %s", url)
					}
					releases = append(releases, fetchedReleases...)
					if header.Get("link") == "" {
						break
					}
					linkHeader := header.Get("link")
					links := strings.Split(linkHeader, ",")
					for _, l := range links {
						if strings.Contains(l, `rel="next"`) {
							url = strings.Trim(strings.Split(strings.Split(l, `rel="next"`)[0], ";")[0], "<> ")
							if slices.Contains(fetchedURLs, url) {
								continue
							}
							continue outer
						}
					}
					break
				}
			}
			var closest *types.Semver = nil
			vs := utils.NewNavigableMap[string, *types.Semver]()
			vs1 := utils.NewNavigableMap[string, map[string]interface{}]()
			for _, r := range releases {
				if r["tag_name"] == nil {
					return versions, burrito.WrapErrorf(err, "Couldn't find tag_name in %s", url)
				}
				tag, ok := r["tag_name"].(string)
				if !ok {
					return versions, burrito.WrapErrorf(err, "tag_name is not a string in %s", url)
				}
				ver, err := types.ParseSemverString(strings.TrimPrefix(strings.Split(tag, "-")[0], "v"))
				if err != nil {
					return versions, burrito.WrapErrorf(err, "tag_name %s is not a valid semver", tag)
				}
				vs.Put(tag, ver)
				vs1.Put(tag, r)
			}
			for _, tag := range vs.Keys() {
				ver := vs.Get(tag)
				if ver.Major != version.Major {
					continue
				}
				if ver.Minor == version.Minor {
					closest = ver
					release = vs1.Get(tag)
					break
				}
				if closest == nil || closest.IsEmpty() || math.Abs(float64(ver.Minor-version.Minor)) < math.Abs(float64(closest.Minor-version.Minor)) {
					closest = ver
					release = vs1.Get(tag)
				}
			}
			for _, tag := range vs.Keys() {
				ver := vs.Get(tag)
				if ver.Major != version.Major {
					continue
				}
				if closest != nil && ver.Minor != closest.Minor {
					continue
				}
				if ver.Patch == version.Patch {
					closest = ver
					release = vs1.Get(tag)
					break
				}
				if closest == nil || math.Abs(float64(ver.Patch-version.Patch)) < math.Abs(float64(closest.Patch-version.Patch)) {
					closest = ver
					release = vs1.Get(tag)
				}
			}
			if release == nil {
				return versions, burrito.WrapErrorf(err, "Couldn't find a release for %s", version.StringValue())
			}
		}
		if release["zipball_url"] == nil {
			return versions, burrito.WrapErrorf(err, "Couldn't find zipball_url")
		}
		url, ok := release["zipball_url"].(string)
		if !ok {
			return versions, burrito.WrapErrorf(err, "zipball_url is not a string in %s", url)
		}

		if release["tag_name"] == nil {
			return versions, burrito.WrapErrorf(err, "Couldn't find tag_name in %s", url)
		}
		tag, ok := release["tag_name"].(string)
		if !ok {
			return versions, burrito.WrapErrorf(err, "tag_name is not a string in %s", url)
		}
		cleanVer := strings.TrimPrefix(strings.Split(tag, "-")[0], "v")
		err = safeio.Resolver.MkdirAll(path.Join(base, cleanVer))
		if err != nil && !os.IsExist(err) {
			return versions, burrito.WrapErrorf(err, "An error occurred while creating cache directory")
		}
		out, err := safeio.Resolver.Create(path.Join(base, cleanVer, outName))
		if err != nil {
			return versions, burrito.WrapErrorf(err, "An error occurred while creating file %s", outName)
		}
		utils.Logger.Infof("Downloading %s", url)
		resp, _, err := safeio.Resolver.HttpGet(url)
		if err != nil {
			return versions, burrito.WrapErrorf(err, "An error occurred while downloading %s", url)
		}
		_, err = io.Copy(out, resp)
		if err != nil {
			return versions, burrito.WrapErrorf(err, "An error occurred while downloading %s", url)
		}
		err = out.Close()
		if err != nil {
			return versions, burrito.WrapErrorf(err, "An error occurred while downloading %s", url)
		}
		err = resp.Close()
		if err != nil {
			return versions, burrito.WrapErrorf(err, "An error occurred while downloading %s", url)
		}

		err = unzip(path.Join(base, cleanVer, outName), filepath.Join(base, cleanVer))
		if err != nil {
			return versions, burrito.WrapErrorf(err, "An error occurred while extracting %s", outName)
		}
		err = safeio.Resolver.Remove(path.Join(base, cleanVer, outName))
		if err != nil {
			return versions, burrito.WrapErrorf(err, "An error occurred while removing %s", outName)
		}

		versions.Put(cleanVer, path.Join(base, cleanVer, dirName))
		return versions, err
	}
	packDir := path.Join(installDir.StringValue(), "data", "behavior_packs")
	if !isBp {
		packDir = path.Join(installDir.StringValue(), "data", "resource_packs")
	}
	dir, err := os.ReadDir(packDir)
	if err != nil {
		return versions, burrito.WrapErrorf(err, "Failed to read pack directory")
	}
	for _, d := range dir {
		p := path.Join(packDir, d.Name())
		if d.IsDir() {
			f, err := os.ReadFile(path.Join(p, "manifest.json"))
			if err != nil {
				if os.IsNotExist(err) {
					continue
				}
				return versions, burrito.WrapErrorf(err, "Failed to read manifest.json in %s", p)
			}
			var manifest *types.JsonObject
			manifest, err = types.ParseJsonObject(f)
			if err != nil {
				return versions, burrito.WrapErrorf(err, "Failed to parse manifest.json in %s", p)
			}
			if manifest.ContainsKey("header") {
				header := manifest.Get("header").(*types.JsonObject)
				if header.Get("uuid").StringValue() != uuid {
					continue
				}
				if header.ContainsKey("version") {
					version := header.Get("version").Unbox().([]interface{})
					array, err := types.ParseSemverArray(version)
					if err != nil {
						return versions, burrito.WrapErrorf(err, "Failed to parse version in %s", p)
					}
					versions.Put(array.StringValue(), p)
				}
			}
		}
	}
	versions.Sort(func(a, b string) int {
		aVer, err := types.ParseSemverString(a)
		if err != nil {
			return 0
		}
		bVer, err := types.ParseSemverString(b)
		if err != nil {
			return 0
		}
		return aVer.CompareTo(bVer)
	})
	return versions, nil
}

func getItemInfo(id *types.JsonString, metadata *types.JsonNumber) (*types.JsonObject, error) {
	if id.StringValue() == "" {
		return types.NewJsonObject(), nil
	}
	if itemInfos.IsEmpty() {
		err := fetchItemInfos()
		if err != nil {
			return types.NewJsonObject(), burrito.WrapErrorf(err, "Failed to fetch item infos")
		}
	}
	item := itemInfos.Get(id.StringValue())
	if item == nil {
		return types.NewJsonObject(), nil
	}
	return types.AsObject(item[int(metadata.IntValue())]), nil
}

func getItemInfo1(id *types.JsonString) (*types.JsonObject, error) {
	return getItemInfo(id, &types.JsonNumber{Value: 0})
}

func getAllItems() *types.JsonArray {
	if itemInfos.IsEmpty() {
		err := fetchItemInfos()
		if err != nil {
			utils.Logger.Error(err)
			return types.NewJsonArray()
		}
	}
	var items = make([]*types.JsonObject, 0)
	for _, item := range itemInfosByName.Values() {
		items = append(items, types.AsObject(item))
	}
	return types.AsArray(items)
}

func findItemInfoById(id *types.JsonString, metadata *types.JsonNumber) (*types.JsonObject, error) {
	if id.StringValue() == "" {
		return types.NewJsonObject(), nil
	}
	if itemInfos.IsEmpty() {
		err := fetchItemInfos()
		if err != nil {
			return types.NewJsonObject(), burrito.WrapErrorf(err, "Failed to fetch item infos")
		}
	}
	if itemInfos.ContainsKey(id.StringValue()) {
		return types.AsObject(itemInfos.Get(id.StringValue())), nil
	}
	find, err := fuzzy.ExtractOne(id.StringValue(), itemInfos.Keys())
	if err != nil {
		return types.NewJsonObject(), burrito.WrapErrorf(err, "Failed to find item info")
	}
	if find == nil {
		return types.NewJsonObject(), nil
	} else {
		return types.AsObject(itemInfos.Get(find.Match)[int(metadata.IntValue())]), nil
	}
}

func findItemInfoById1(id *types.JsonString) (*types.JsonObject, error) {
	return findItemInfoById(id, &types.JsonNumber{Value: 0})
}

func findItemInfoByName(name *types.JsonString) (*types.JsonObject, error) {
	if name.StringValue() == "" {
		return types.NewJsonObject(), nil
	}
	if itemInfos.IsEmpty() {
		err := fetchItemInfos()
		if err != nil {
			return types.NewJsonObject(), burrito.WrapErrorf(err, "Failed to fetch item infos")
		}
	}
	if itemInfosByName.ContainsKey(name.StringValue()) {
		return types.AsObject(itemInfosByName.Get(name.StringValue())), nil
	}
	find, err := fuzzy.ExtractOne(name.StringValue(), itemInfosByName.Keys())
	if err != nil {
		return types.NewJsonObject(), burrito.WrapErrorf(err, "Failed to find item info")
	}
	if find == nil {
		return types.NewJsonObject(), nil
	} else {
		return types.AsObject(itemInfosByName.Get(find.Match)), nil
	}
}

func fetchItemInfos() error {
	url := "https://raw.githubusercontent.com/stirante/minecraft-item-db/main/items.json"
	outName := "items.json"

	base := "packs"
	if json2.CacheDir != "" {
		base = json2.CacheDir
	}

	stat, err := safeio.Resolver.Stat(path.Join(base, outName))
	if os.IsNotExist(err) || time.Since(stat.ModTime()) > 168*time.Hour {
		err = safeio.Resolver.MkdirAll(base)
		if err != nil && !os.IsExist(err) {
			return burrito.WrapErrorf(err, "An error occurred while creating cache directory")
		}
		out, err := safeio.Resolver.Create(path.Join(base, outName))
		if err != nil {
			return burrito.WrapErrorf(err, "An error occurred while creating file %s", outName)
		}
		utils.Logger.Infof("Downloading %s", url)
		resp, _, err := safeio.Resolver.HttpGet(url)
		if err != nil {
			return burrito.WrapErrorf(err, "An error occurred while downloading %s", url)
		}
		_, err = io.Copy(out, resp)
		if err != nil {
			return burrito.WrapErrorf(err, "An error occurred while downloading %s", url)
		}
		err = out.Close()
		if err != nil {
			return burrito.WrapErrorf(err, "An error occurred while downloading %s", url)
		}
		err = resp.Close()
		if err != nil {
			return burrito.WrapErrorf(err, "An error occurred while downloading %s", url)
		}
	}
	open, err := safeio.Resolver.Open(path.Join(base, outName))
	if err != nil {
		return burrito.WrapErrorf(err, "An error occurred while opening %s", outName)
	}
	readAll, err := io.ReadAll(open)
	if err != nil {
		return burrito.WrapErrorf(err, "An error occurred while reading %s", outName)
	}
	arr, err := types.ParseJsonArray(readAll)
	if err != nil {
		return burrito.WrapErrorf(err, "An error occurred while parsing %s", outName)
	}
	itemInfos = utils.NewNavigableMap[string, map[int]interface{}]()
	itemInfosByName = utils.NewNavigableMap[string, interface{}]()
	for _, v := range arr.Value {
		item := v.(*types.JsonObject)
		id := strings.TrimPrefix(item.Get("id").StringValue(), "minecraft:")
		if !itemInfos.ContainsKey(id) {
			itemInfos.Put(id, make(map[int]interface{}))
		}
		itemInfos.Get(id)[int((item.Get("metadata").(*types.JsonNumber)).IntValue())] = item
		itemInfosByName.Put(item.Get("langName").StringValue(), item)
	}
	return nil
}

func FetchCache() error {
	_, err := findPackVersions(true, VanillaBpUUID, nil)
	if err != nil {
		return burrito.WrapErrorf(err, "Failed to cache vanilla behavior pack")
	}
	_, err = findPackVersions(false, VanillaRpUUID, nil)
	if err != nil {
		return burrito.WrapErrorf(err, "Failed to cache vanilla resource pack")
	}
	err = fetchItemInfos()
	if err != nil {
		return burrito.WrapErrorf(err, "Failed to cache item infos")
	}
	return nil
}
