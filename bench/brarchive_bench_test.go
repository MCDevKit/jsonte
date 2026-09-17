package bench

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/MCDevKit/jsonte/jsonte/brarchive"
	"github.com/MCDevKit/jsonte/jsonte/safeio"
)

const benchArchiveEntries = 200

// buildBenchArchive encodes entries named entry_<i>.json, each ~2 KB, in the .brarchive format.
func buildBenchArchive() []byte {
	var header, data bytes.Buffer
	header.Write([]byte{0x7d, 0x27, 0x25, 0xb1, 0xa0, 0x52, 0x70, 0x26})
	_ = binary.Write(&header, binary.LittleEndian, uint32(benchArchiveEntries))
	_ = binary.Write(&header, binary.LittleEndian, uint32(1))
	content := []byte(`{"format_version":"1.21.0","minecraft:entity":{"description":{"identifier":"minecraft:bench"},"components":{` +
		`"minecraft:health":{"value":20},"minecraft:movement":{"value":0.25},"minecraft:physics":{},"minecraft:pushable":{}` +
		`,"minecraft:collision_box":{"width":0.9,"height":1.3},"minecraft:type_family":{"family":["bench","mob"]}}}}`)
	for len(content) < 2048 {
		content = append(content, ' ')
	}
	for i := 0; i < benchArchiveEntries; i++ {
		name := fmt.Sprintf("entry_%d.json", i)
		entry := make([]byte, 256)
		entry[0] = byte(len(name))
		copy(entry[1:], name)
		binary.LittleEndian.PutUint32(entry[248:], uint32(data.Len()))
		binary.LittleEndian.PutUint32(entry[252:], uint32(len(content)))
		data.Write(content)
		header.Write(entry)
	}
	return append(header.Bytes(), data.Bytes()...)
}

func writeBenchArchive(b *testing.B) (archive string, loose string) {
	dir := b.TempDir()
	archive = filepath.Join(dir, "entities.brarchive")
	if err := os.WriteFile(archive, buildBenchArchive(), 0644); err != nil {
		b.Fatal(err)
	}
	a, err := brarchive.Open(archive)
	if err != nil {
		b.Fatal(err)
	}
	r, err := a.Open("entry_100.json")
	if err != nil {
		b.Fatal(err)
	}
	content, err := io.ReadAll(r)
	if err != nil {
		b.Fatal(err)
	}
	loose = filepath.Join(dir, "entry_100.json")
	if err := os.WriteFile(loose, content, 0644); err != nil {
		b.Fatal(err)
	}
	return archive, loose
}

// BenchmarkBrarchiveParse measures building the entry index of a 200-entry archive from memory.
func BenchmarkBrarchiveParse(b *testing.B) {
	data := buildBenchArchive()
	r := bytes.NewReader(data)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := brarchive.Parse(r, int64(len(data)), time.Time{}); err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkBrarchiveOpenCold measures mmap + parse of an archive file that is not cached yet.
func BenchmarkBrarchiveOpenCold(b *testing.B) {
	dir := b.TempDir()
	data := buildBenchArchive()
	paths := make([]string, b.N)
	for i := range paths {
		paths[i] = filepath.Join(dir, fmt.Sprintf("%d.brarchive", i))
		if err := os.WriteFile(paths[i], data, 0644); err != nil {
			b.Fatal(err)
		}
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := brarchive.Open(paths[i]); err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkBrarchiveReadEntry measures reading one ~2 KB entry through the resolver pseudo path
// from an already cached (memory-mapped) archive.
func BenchmarkBrarchiveReadEntry(b *testing.B) {
	archive, _ := writeBenchArchive(b)
	pseudo := brarchive.JoinPath(archive, "entry_100.json")
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		f, err := safeio.DefaultIOResolver.Open(pseudo)
		if err != nil {
			b.Fatal(err)
		}
		if _, err := io.ReadAll(f); err != nil {
			b.Fatal(err)
		}
		_ = f.Close()
	}
}

// BenchmarkBrarchiveReadLoose is the baseline for BenchmarkBrarchiveReadEntry: the same file read from disk.
func BenchmarkBrarchiveReadLoose(b *testing.B) {
	_, loose := writeBenchArchive(b)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		f, err := safeio.DefaultIOResolver.Open(loose)
		if err != nil {
			b.Fatal(err)
		}
		if _, err := io.ReadAll(f); err != nil {
			b.Fatal(err)
		}
		_ = f.Close()
	}
}

// BenchmarkBrarchiveStatEntry measures Stat of a pseudo path from a cached archive.
func BenchmarkBrarchiveStatEntry(b *testing.B) {
	archive, _ := writeBenchArchive(b)
	pseudo := brarchive.JoinPath(archive, "entry_100.json")
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := safeio.DefaultIOResolver.Stat(pseudo); err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkBrarchiveVanillaEntities reads every entry of the vanilla entities archive from a local
// Minecraft install. Skipped when Minecraft Bedrock is not installed.
func BenchmarkBrarchiveVanillaEntities(b *testing.B) {
	// WindowsApps cannot be listed, so the install dir has to come from PowerShell
	out, err := exec.Command("powershell", "(Get-AppxPackage -Name Microsoft.MinecraftUWP).InstallLocation").Output()
	installDir := strings.TrimSpace(string(out))
	if err != nil || installDir == "" {
		b.Skip("Minecraft Bedrock is not installed")
	}
	a, err := brarchive.Open(filepath.Join(installDir, "data", "behavior_packs", "vanilla", "__brarchive", "entities.brarchive"))
	if err != nil {
		b.Fatal(err)
	}
	names := a.Names()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, name := range names {
			r, err := a.Open(name)
			if err != nil {
				b.Fatal(err)
			}
			if _, err := io.Copy(io.Discard, r); err != nil {
				b.Fatal(err)
			}
		}
	}
}
