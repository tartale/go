package filez

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/tartale/go/pkg/errorz"
)

// IsDir is a simple convenience function for the
// os.IsDir() function, which ignores the error and
// returns false.
func IsDir(name string) bool {

	stat, err := os.Stat(name)
	if err != nil {
		return false
	}

	return stat.IsDir()
}

// Exists returns true if the given file or directory
// exists and is reachable.
// See https://stackoverflow.com/a/12518877/1258206
func Exists(name string) bool {

	if _, err := os.Stat(name); err == nil {
		return true
	}

	return false
}

// Exist returns a list of files or directories
// that do not exist
func Exist(paths ...string) (missingPaths []string) {

	for _, p := range paths {
		if !Exists(p) {
			missingPaths = append(missingPaths, p)
		}
	}

	return missingPaths
}

func PathWithoutExtension(path string) string {

	if path == "" {
		return ""
	}
	ext := filepath.Ext(path)
	path = strings.TrimSuffix(path, ext)

	return path
}

func NameWithoutExtension(path string) string {

	if path == "" {
		return ""
	}
	newPath := PathWithoutExtension(path)
	name := filepath.Base(newPath)

	return name
}

func MustOpenFile(name string, flag int, perm fs.FileMode) *os.File {

	file, err1 := os.OpenFile(name, os.O_CREATE|os.O_WRONLY, 0644)
	if errors.Is(err1, os.ErrNotExist) {
		dir := path.Dir(name)
		err2 := os.MkdirAll(dir, os.FileMode(0755))
		if err2 != nil {
			err := fmt.Errorf("%w: %w: %w", errorz.ErrFatal, err1, err2)
			panic(err)
		}
	}

	return file
}

// MkdirAllParents creates the parent directories for all the
// given file paths.
func MkdirAllParents(paths ...string) error {

	for _, p := range paths {
		dir := filepath.Dir(p)
		err := os.MkdirAll(dir, os.FileMode(0755))
		if err != nil {
			return err
		}
	}

	return nil
}

// MkdirAll creates directories for all the given paths.
func MkdirAll(paths ...string) error {

	for _, p := range paths {
		err := os.MkdirAll(p, os.FileMode(0755))
		if err != nil {
			return err
		}
	}

	return nil
}

func MustMkdirAll(dir string) {

	err := os.MkdirAll(dir, os.FileMode(0755))
	if err != nil {
		err := fmt.Errorf("%w: %w", errorz.ErrFatal, err)
		panic(err)
	}
}

func MustRename(oldpath, newpath string) {

	var (
		renameErr error
		mkdirErr  error
	)

	if !Exists(oldpath) {
		panic(fmt.Errorf("%w: %s", os.ErrNotExist, oldpath))
	}
	renameErr = os.Rename(oldpath, newpath)
	if renameErr == nil {
		return
	}
	if renameErr != nil && !errors.Is(renameErr, os.ErrNotExist) {
		panic(fmt.Errorf("%w: %w", errorz.ErrFatal, renameErr))
	}
	dir := path.Dir(newpath)
	mkdirErr = os.MkdirAll(dir, os.FileMode(0755))
	if mkdirErr != nil {
		panic(fmt.Errorf("%w: %w", errorz.ErrFatal, mkdirErr))
	}
	renameErr = os.Rename(oldpath, newpath)
	if renameErr != nil {
		panic(fmt.Errorf("%w: %w", errorz.ErrFatal, renameErr))
	}
}

func MustReadAll(path string) []byte {

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	return bytes
}

func GetRootDirForCaller(caller int) (rootDir string, found bool) {

	var (
		file string
		ok   bool
	)
	if _, file, _, ok = runtime.Caller(caller); !ok {
		return "", false
	}

	return GetRootDirForFile(file)
}

func GetRootDirForFile(file string) (rootDir string, found bool) {

	rootDir = path.Dir(file)
	found = false
	for rootDir != "/" {
		gitDir := path.Join(rootDir, ".git")
		if !IsDir(gitDir) {
			rootDir = path.Dir(rootDir)
			continue
		}
		found = true
		break
	}

	return
}
