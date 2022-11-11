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

func getMinecraftInstallDir() (types.JsonString, error) {
	if runtime.GOOS != "windows" {
		return types.EmptyString, burrito.WrappedErrorf("This function works only on Windows")
	}
	if installDir == "" {
		output, err := safeio.Resolver.ExecCommand("powershell", "(Get-AppxPackage -Name Microsoft.MinecraftUWP).InstallLocation")
		if err != nil {
			return types.EmptyString, burrito.WrapErrorf(err, "An error occurred while getting Minecraft install directory")
		}
		installDir = strings.Trim(string(output), "\r\n \t")
	}
	return types.NewString(installDir), nil
}

func getLatestBPFile(p types.JsonString) (types.JsonString, error) {
	if bpFiles.IsEmpty() {
		bp, err := findPackVersions(true, VanillaBpUUID)
		if err != nil {
			return types.EmptyString, burrito.WrapErrorf(err, "An error occurred while reading behavior packs")
		}
		bpFiles = bp
	}
	return getLatestFile(p.StringValue(), bpFiles)
}

func getLatestRPFile(p types.JsonString) (types.JsonString, error) {
	if rpFiles.IsEmpty() {
		rp, err := findPackVersions(false, VanillaRpUUID)
		if err != nil {
			return types.EmptyString, burrito.WrapErrorf(err, "An error occurred while reading resource packs")
		}
		rpFiles = rp
	}
	return getLatestFile(p.StringValue(), rpFiles)
}

func listLatestRPFiles(p types.JsonString) (types.JsonArray, error) {
	if rpFiles.IsEmpty() {
		rp, err := findPackVersions(false, VanillaRpUUID)
		if err != nil {
			return types.NewJsonArray(), burrito.WrapErrorf(err, "An error occurred while reading resource packs")
		}
		rpFiles = rp
	}
	return listLatestFiles(p.StringValue(), rpFiles)
}

func listLatestBPFiles(p types.JsonString) (types.JsonArray, error) {
	if bpFiles.IsEmpty() {
		bp, err := findPackVersions(true, VanillaBpUUID)
		if err != nil {
			return types.NewJsonArray(), burrito.WrapErrorf(err, "An error occurred while reading behavior packs")
		}
		bpFiles = bp
	}
	return listLatestFiles(p.StringValue(), bpFiles)
}

func listLatestFiles(p string, m utils.NavigableMap[string, string]) (types.JsonArray, error) {
	result := map[string]string{}
	keys := m.Keys()
	for i := len(keys) - 1; i >= 0; i-- {
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
			rel, err := filepath.Rel(s, f)
			if err != nil {
				return types.NewJsonArray(), err
			}
			if _, ok := result[rel]; !ok {
				result[rel] = f
			}
		}
	}
	arr := make([]types.JsonType, len(result))
	i := 0
	for _, v := range result {
		arr[i] = types.NewString(v)
		i++
	}
	return types.JsonArray{Value: arr}, nil
}

func getLatestFile(p string, m utils.NavigableMap[string, string]) (types.JsonString, error) {
	keys := m.Keys()
	for i := len(keys) - 1; i >= 0; i-- {
		s := path.Join(m.Get(keys[i]), p)
		_, err := safeio.Resolver.Stat(s)
		if err != nil {
			if os.IsNotExist(err) {
				continue
			} else {
				return types.EmptyString, burrito.WrapErrorf(err, "An error occurred while reading file %s", p)
			}
		}
		return types.NewString(s), nil
	}
	return types.EmptyString, burrito.WrapErrorf(os.ErrNotExist, "File '%s' does not exist", p)
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
		if json2.CacheDir != "" {
			base = json2.CacheDir
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
			f, err := ioutil.ReadFile(path.Join(p, "manifest.json"))
			if err != nil {
				if os.IsNotExist(err) {
					continue
				}
				return versions, burrito.WrapErrorf(err, "Failed to read manifest.json in %s", p)
			}
			var manifest types.JsonObject
			manifest, err = types.ParseJsonObject(f)
			if err != nil {
				return versions, burrito.WrapErrorf(err, "Failed to parse manifest.json in %s", p)
			}
			if manifest.ContainsKey("header") {
				header := manifest.Get("header").(types.JsonObject)
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

func getItemInfo(id types.JsonString, metadata types.JsonNumber) (types.JsonObject, error) {
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

func getItemInfo1(id types.JsonString) (types.JsonObject, error) {
	return getItemInfo(id, types.JsonNumber{Value: 0})
}

func findItemInfoById(id types.JsonString, metadata types.JsonNumber) (types.JsonObject, error) {
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

func findItemInfoById1(id types.JsonString) (types.JsonObject, error) {
	return findItemInfoById(id, types.JsonNumber{Value: 0})
}

func findItemInfoByName(name types.JsonString) (types.JsonObject, error) {
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
	arr, err := types.ParseJsonArray(readAll)
	if err != nil {
		return burrito.WrapErrorf(err, "An error occurred while parsing %s", outName)
	}
	itemInfos = utils.NewNavigableMap[string, map[int]interface{}]()
	itemInfosByName = utils.NewNavigableMap[string, interface{}]()
	for _, v := range arr.Value {
		item := v.(types.JsonObject)
		id := strings.TrimPrefix(item.Get("id").StringValue(), "minecraft:")
		if !itemInfos.ContainsKey(id) {
			itemInfos.Put(id, make(map[int]interface{}))
		}
		itemInfos.Get(id)[int((item.Get("metadata").(types.JsonNumber)).IntValue())] = item
		itemInfosByName.Put(item.Get("langName").StringValue(), item)
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
