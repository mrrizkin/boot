package view

import (
	"bytes"
	"crypto/sha256"
	"errors"
	"fmt"
	"net/http"
	"os"
	pathpkg "path"
	"path/filepath"
	"sort"

	"github.com/nikolalohinski/gonja/v2"
	"github.com/nikolalohinski/gonja/v2/exec"
	"github.com/nikolalohinski/gonja/v2/loaders"

	"github.com/mrrizkin/boot/system/view/tag"
)

// stat returns the FileInfo structure describing file.
func stat(fs http.FileSystem, name string) (os.FileInfo, error) {
	f, err := fs.Open(name)
	if err != nil {
		return nil, err
	}
	defer f.Close() //nolint: errcheck // No need to check error
	return f.Stat()
}

func walk(fs http.FileSystem, root string, walkFn filepath.WalkFunc) error {
	info, err := stat(fs, root)
	if err != nil {
		return walkFn(root, nil, err)
	}
	return walkInternal(fs, root, info, walkFn)
}

func walkInternal(
	fs http.FileSystem,
	path string,
	info os.FileInfo,
	walkFn filepath.WalkFunc,
) error {
	err := walkFn(path, info, nil)
	if err != nil {
		if info.IsDir() && errors.Is(err, filepath.SkipDir) {
			return nil
		}
		return err
	}

	if !info.IsDir() {
		return nil
	}

	names, err := readDirNames(fs, path)
	if err != nil {
		return walkFn(path, info, err)
	}

	for _, name := range names {
		filename := pathpkg.Join(path, name)
		fileInfo, err := stat(fs, filename)
		if err != nil {
			if err := walkFn(filename, fileInfo, err); err != nil &&
				!errors.Is(err, filepath.SkipDir) {
				return err
			}
		} else {
			err = walkInternal(fs, filename, fileInfo, walkFn)
			if err != nil {
				if !fileInfo.IsDir() || !errors.Is(err, filepath.SkipDir) {
					return err
				}
			}
		}
	}
	return nil
}

func readDirNames(fs http.FileSystem, dirname string) ([]string, error) {
	fis, err := readDir(fs, dirname)
	if err != nil {
		return nil, err
	}
	names := make([]string, len(fis))
	for i := range fis {
		names[i] = fis[i].Name()
	}
	sort.Strings(names)
	return names, nil
}

func readDir(fs http.FileSystem, name string) ([]os.FileInfo, error) {
	f, err := fs.Open(name)
	if err != nil {
		return nil, err
	}
	defer f.Close() //nolint: errcheck // No need to check error
	return f.Readdir(0)
}

func fromBytes(source []byte, loader loaders.Loader) (*exec.Template, error) {
	rootID := fmt.Sprintf("root-%s", string(sha256.New().Sum(source)))
	shiftedLoader, err := loaders.NewShiftedLoader(rootID, bytes.NewReader(source), loader)
	if err != nil {
		return nil, err
	}
	environment := gonja.DefaultEnvironment
	for tag, parser := range tag.All {
		environment.ControlStructures.Register(tag, parser)
	}
	return exec.NewTemplate(rootID, gonja.DefaultConfig, shiftedLoader, gonja.DefaultEnvironment)
}
