package functions

import (
	"github.com/MCDevKit/jsonte/jsonte/utils"
	"github.com/stirante/jsonc"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
)

func RegisterMinecraftFunctions() {
	RegisterFunction(JsonFunction{
		Name:     "getMinecraftInstallDir",
		Body:     getMinecraftInstallDir,
		IsUnsafe: true,
	})
	RegisterFunction(JsonFunction{
		Name:     "getLatestBPFile",
		Body:     getLatestBPFile,
		IsUnsafe: true,
	})
	RegisterFunction(JsonFunction{
		Name:     "getLatestRPFile",
		Body:     getLatestRPFile,
		IsUnsafe: true,
	})
	RegisterFunction(JsonFunction{
		Name: "listLatestRPFiles",
		Body: listLatestRPFiles,
	})
	RegisterFunction(JsonFunction{
		Name: "listLatestBPFiles",
		Body: listLatestBPFiles,
	})
}

var installDir = ""
var VanillaRpUUID = "0575c61f-a5da-4b7f-9961-ffda2908861e"
var VanillaBpUUID = "fe9f8597-5454-481a-8730-8d070a8e2e58"

var rpFiles = utils.NavigableMap[string, string]{}
var bpFiles = utils.NavigableMap[string, string]{}

func getMinecraftInstallDir() (string, error) {
	if installDir == "" {
		output, err := exec.Command("powershell", "(Get-AppxPackage -Name Microsoft.MinecraftUWP).InstallLocation").Output()
		if err != nil {
			return "", err
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
	if bpFiles.IsEmpty() {
		rp, err := findPackVersions(false, VanillaRpUUID)
		if err != nil {
			return "", utils.WrapErrorf(err, "An error occurred while reading resource packs")
		}
		rpFiles = rp
	}
	return getLatestFile(p, rpFiles)
}

func listLatestRPFiles(p string) (utils.JsonArray, error) {
	if bpFiles.IsEmpty() {
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
	return "", utils.WrapErrorf(os.ErrNotExist, "File %s does not exist", p)
}

func findPackVersions(isBp bool, uuid string) (utils.NavigableMap[string, string], error) {
	versions := utils.NewNavigableMap[string, string]()
	installDir, err := getMinecraftInstallDir()
	if err != nil {
		return versions, utils.WrapErrorf(err, "Failed to get Minecraft install directory")
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
							array, err := utils.ParseSemverArray(versionArr)
							if err != nil {
								continue
							}
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
