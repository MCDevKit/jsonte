package functions

import (
	"archive/zip"
	"encoding/json"
	"github.com/Bedrock-OSS/go-burrito/burrito"
	"github.com/MCDevKit/jsonte/jsonte/safeio"
	"github.com/MCDevKit/jsonte/jsonte/utils"
	"github.com/sahilm/fuzzy"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"time"
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
}

var installDir = ""
var VanillaRpUUID = "0575c61f-a5da-4b7f-9961-ffda2908861e"
var VanillaBpUUID = "fe9f8597-5454-481a-8730-8d070a8e2e58"

var rpFiles = utils.NavigableMap[string, string]{}
var bpFiles = utils.NavigableMap[string, string]{}
var itemInfos = utils.NavigableMap[string, map[int]interface{}]{}
var itemInfosByName = utils.NavigableMap[string, interface{}]{}

func getMinecraftInstallDir() (string, error) {
	if runtime.GOOS != "windows" {
		return "", burrito.WrappedErrorf("This function works only on Windows")
	}
	if installDir == "" {
		output, err := safeio.Resolver.ExecCommand("powershell", "(Get-AppxPackage -Name Microsoft.MinecraftUWP).InstallLocation")
		if err != nil {
			return "", burrito.WrapErrorf(err, "An error occurred while getting Minecraft install directory")
		}
		installDir = strings.Trim(string(output), "\r\n \t")
	}
	return installDir, nil
}

func getLatestBPFile(p string) (string, error) {
	if bpFiles.IsEmpty() {
		bp, err := findPackVersions(true, VanillaBpUUID)
		if err != nil {
			return "", burrito.WrapErrorf(err, "An error occurred while reading behavior packs")
		}
		bpFiles = bp
	}
	return getLatestFile(p, bpFiles)
}

func getLatestRPFile(p string) (string, error) {
	if rpFiles.IsEmpty() {
		rp, err := findPackVersions(false, VanillaRpUUID)
		if err != nil {
			return "", burrito.WrapErrorf(err, "An error occurred while reading resource packs")
		}
		rpFiles = rp
	}
	return getLatestFile(p, rpFiles)
}

func listLatestRPFiles(p string) ([]interface{}, error) {
	if rpFiles.IsEmpty() {
		rp, err := findPackVersions(false, VanillaRpUUID)
		if err != nil {
			return []interface{}{}, burrito.WrapErrorf(err, "An error occurred while reading resource packs")
		}
		rpFiles = rp
	}
	return listLatestFiles(p, rpFiles)
}

func listLatestBPFiles(p string) ([]interface{}, error) {
	if bpFiles.IsEmpty() {
		bp, err := findPackVersions(true, VanillaBpUUID)
		if err != nil {
			return []interface{}{}, burrito.WrapErrorf(err, "An error occurred while reading behavior packs")
		}
		bpFiles = bp
	}
	return listLatestFiles(p, bpFiles)
}

func listLatestFiles(p string, m utils.NavigableMap[string, string]) ([]interface{}, error) {
	result := map[string]string{}
	keys := m.Keys()
	for i := len(keys) - 1; i >= 0; i-- {
		s := path.Join(m.Get(keys[i]), p)
		_, err := safeio.Resolver.Stat(s)
		if err != nil {
			if os.IsNotExist(err) {
				continue
			} else {
				return nil, burrito.WrapErrorf(err, "An error occurred while reading file %s", p)
			}
		}
		recursive, err := safeio.Resolver.OpenDirRecursive(s)
		if err != nil {
			return nil, burrito.WrapErrorf(err, "An error occurred while reading file %s", p)
		}
		for _, f := range recursive {
			rel, err := filepath.Rel(s, f)
			if err != nil {
				return nil, err
			}
			if _, ok := result[rel]; !ok {
				result[rel] = f
			}
		}
	}
	arr := make([]interface{}, len(result))
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
		_, err := safeio.Resolver.Stat(s)
		if err != nil {
			if os.IsNotExist(err) {
				continue
			} else {
				return "", burrito.WrapErrorf(err, "An error occurred while reading file %s", p)
			}
		}
		return s, nil
	}
	return "", burrito.WrapErrorf(os.ErrNotExist, "File '%s' does not exist", p)
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

func findPackVersions(isBp bool, uuid string) (utils.NavigableMap[string, string], error) {
	versions := utils.NewNavigableMap[string, string]()
	installDir, err := getMinecraftInstallDir()
	if err != nil {
		url := "https://api.github.com/repos/Mojang/bedrock-samples/releases/latest"
		outName := "packs.zip"
		dirName := "RP"
		if isBp {
			dirName = "BP"
		}

		base := "packs"
		if utils.CacheDir != "" {
			base = utils.CacheDir
		}

		stat, err := safeio.Resolver.Stat(path.Join(base, dirName))
		if err == nil && stat.IsDir() {
			versions.Put("1.0.0", path.Join(base, dirName))
			return versions, nil
		}
		utils.Logger.Infof("Resolving %s", url)
		resp, err := safeio.Resolver.HttpGet(url)
		if err != nil {
			return versions, burrito.WrapErrorf(err, "An error occurred while resolving %s", url)
		}
		var release map[string]interface{}
		err = json.NewDecoder(resp).Decode(&release)
		if err != nil {
			return versions, burrito.WrapErrorf(err, "An error occurred while parsing %s", url)
		}
		if release["zipball_url"] == nil {
			return versions, burrito.WrapErrorf(err, "Couldn't find zipball_url in %s", url)
		}
		url, ok := release["zipball_url"].(string)
		if !ok {
			return versions, burrito.WrapErrorf(err, "zipball_url is not a string in %s", url)
		}
		err = resp.Close()
		if err != nil {
			return versions, burrito.WrapErrorf(err, "An error occurred while closing %s", url)
		}

		err = safeio.Resolver.MkdirAll(base)
		if err != nil && !os.IsExist(err) {
			return versions, burrito.WrapErrorf(err, "An error occurred while creating cache directory")
		}
		out, err := safeio.Resolver.Create(path.Join(base, outName))
		if err != nil {
			return versions, burrito.WrapErrorf(err, "An error occurred while creating file %s", outName)
		}
		utils.Logger.Infof("Downloading %s", url)
		resp, err = safeio.Resolver.HttpGet(url)
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

		err = unzip(path.Join(base, outName), base)
		if err != nil {
			return versions, burrito.WrapErrorf(err, "An error occurred while extracting %s", outName)
		}
		err = safeio.Resolver.Remove(path.Join(base, outName))
		if err != nil {
			return versions, burrito.WrapErrorf(err, "An error occurred while removing %s", outName)
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
		return versions, burrito.WrapErrorf(err, "Failed to read pack directory")
	}
	for _, d := range dir {
		p := path.Join(packDir, d.Name())
		if d.IsDir() {
			f, err := ioutil.ReadFile(path.Join(p, "manifest.json"))
			if err != nil {
				if os.IsNotExist(err) {
					continue
				}
				return versions, burrito.WrapErrorf(err, "Failed to read manifest.json in %s", p)
			}
			var manifest utils.NavigableMap[string, interface{}]
			manifest, err = utils.ParseJsonObject(f)
			if err != nil {
				return versions, burrito.WrapErrorf(err, "Failed to parse manifest.json in %s", p)
			}
			if manifest.ContainsKey("header") {
				header := manifest.Get("header").(utils.NavigableMap[string, interface{}])
				if header.Get("uuid") != uuid {
					continue
				}
				if header.ContainsKey("version") {
					version := utils.UnwrapContainers(header.Get("version")).([]interface{})
					array, err := utils.ParseSemverArray(version)
					if err != nil {
						return versions, burrito.WrapErrorf(err, "Failed to parse version in %s", p)
					}
					versions.Put(array.String(), p)
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

func getItemInfo(id string, metadata utils.JsonNumber) (interface{}, error) {
	if id == "" {
		return nil, nil
	}
	if itemInfos.IsEmpty() {
		err := fetchItemInfos()
		if err != nil {
			return nil, burrito.WrapErrorf(err, "Failed to fetch item infos")
		}
	}
	item := itemInfos.Get(id)
	if item == nil {
		return nil, nil
	}
	return item[int(metadata.IntValue())], nil
}

func getItemInfo1(id string) (interface{}, error) {
	return getItemInfo(id, utils.JsonNumber{Value: 0})
}

func findItemInfoById(id string, metadata utils.JsonNumber) (interface{}, error) {
	if id == "" {
		return nil, nil
	}
	if itemInfos.IsEmpty() {
		err := fetchItemInfos()
		if err != nil {
			return nil, burrito.WrapErrorf(err, "Failed to fetch item infos")
		}
	}
	if itemInfos.ContainsKey(id) {
		return itemInfos.Get(id), nil
	}
	find := fuzzy.Find(id, itemInfos.Keys())
	if len(find) == 0 {
		return nil, nil
	} else {
		return itemInfos.Get(find[0].Str)[int(metadata.IntValue())], nil
	}
}

func findItemInfoById1(id string) (interface{}, error) {
	return findItemInfoById(id, utils.JsonNumber{Value: 0})
}

func findItemInfoByName(name string) (interface{}, error) {
	if name == "" {
		return nil, nil
	}
	if itemInfos.IsEmpty() {
		err := fetchItemInfos()
		if err != nil {
			return nil, burrito.WrapErrorf(err, "Failed to fetch item infos")
		}
	}
	if itemInfosByName.ContainsKey(name) {
		return itemInfosByName.Get(name), nil
	}
	find := fuzzy.Find(name, itemInfosByName.Keys())
	if len(find) == 0 {
		return nil, nil
	} else {
		return itemInfosByName.Get(find[0].Str), nil
	}
}

func fetchItemInfos() error {
	url := "https://raw.githubusercontent.com/stirante/minecraft-item-db/main/items.json"
	outName := "items.json"

	base := "packs"
	if utils.CacheDir != "" {
		base = utils.CacheDir
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
		resp, err := safeio.Resolver.HttpGet(url)
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
	readAll, err := ioutil.ReadAll(open)
	if err != nil {
		return burrito.WrapErrorf(err, "An error occurred while reading %s", outName)
	}
	arr, err := utils.ParseJsonArray(readAll)
	if err != nil {
		return burrito.WrapErrorf(err, "An error occurred while parsing %s", outName)
	}
	itemInfos = utils.NewNavigableMap[string, map[int]interface{}]()
	itemInfosByName = utils.NewNavigableMap[string, interface{}]()
	for _, v := range arr {
		item := v.(utils.NavigableMap[string, interface{}])
		id := strings.TrimPrefix(item.Get("id").(string), "minecraft:")
		if !itemInfos.ContainsKey(id) {
			itemInfos.Put(id, make(map[int]interface{}))
		}
		itemInfos.Get(id)[int((item.Get("metadata").(utils.JsonNumber)).IntValue())] = item
		itemInfosByName.Put(item.Get("langName").(string), item)
	}
	return nil
}

func FetchCache() error {
	_, err := findPackVersions(true, VanillaBpUUID)
	if err != nil {
		return burrito.WrapErrorf(err, "Failed to cache vanilla behavior pack")
	}
	_, err = findPackVersions(false, VanillaRpUUID)
	if err != nil {
		return burrito.WrapErrorf(err, "Failed to cache vanilla resource pack")
	}
	err = fetchItemInfos()
	if err != nil {
		return burrito.WrapErrorf(err, "Failed to cache item infos")
	}
	return nil
}
