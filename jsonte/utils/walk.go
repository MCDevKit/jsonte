package utils

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"
)

// WalkDirFollowSymlinks walks the directory tree rooted at root, calling fn for each file or directory.
// If a directory is a symlink, its contents are also walked.
func WalkDirFollowSymlinks(root string, fn fs.WalkDirFunc) error {
	var walk func(string, string, map[string]struct{}) error
	walk = func(logical, real string, seen map[string]struct{}) error {
		if _, ok := seen[real]; ok {
			return nil
		}
		seen[real] = struct{}{}
		info, err := os.Lstat(real)
		if err != nil {
			return fn(logical, nil, err)
		}
		var d = fs.FileInfoToDirEntry(info)
		var target string
		if info.Mode()&fs.ModeSymlink != 0 {
			target, err = filepath.EvalSymlinks(real)
			if err != nil {
				return fn(logical, d, err)
			}
			tinfo, err := os.Stat(real)
			if err != nil {
				return fn(logical, d, err)
			}
			if tinfo.IsDir() {
				d = fs.FileInfoToDirEntry(tinfo)
				info = tinfo
			}
		}
		if err := fn(logical, d, nil); err != nil {
			if errors.Is(err, fs.SkipDir) {
				return nil
			}
			return err
		}
		if info.Mode()&fs.ModeSymlink != 0 && target != "" {
			if !info.IsDir() {
				return nil
			}
			if _, ok := seen[target]; ok {
				return nil
			}
			seen[target] = struct{}{}
			entries, err := os.ReadDir(target)
			if err != nil {
				return fn(logical, d, err)
			}
			for _, e := range entries {
				lpath := filepath.Join(logical, e.Name())
				rpath := filepath.Join(target, e.Name())
				if err := walk(lpath, rpath, seen); err != nil {
					if errors.Is(err, fs.SkipDir) {
						continue
					}
					return err
				}
			}
			return nil
		}
		if !info.IsDir() {
			return nil
		}
		entries, err := os.ReadDir(real)
		if err != nil {
			return fn(logical, d, err)
		}
		for _, e := range entries {
			lpath := filepath.Join(logical, e.Name())
			rpath := filepath.Join(real, e.Name())
			if err := walk(lpath, rpath, seen); err != nil {
				if errors.Is(err, fs.SkipDir) {
					continue
				}
				return err
			}
		}
		return nil
	}
	abs, err := filepath.Abs(root)
	if err != nil {
		abs = root
	}
	return walk(root, abs, map[string]struct{}{})
}
