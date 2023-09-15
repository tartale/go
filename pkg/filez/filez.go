package filez

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path"

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

// Exists returns true if the file exists and
// is reachable.
// See https://stackoverflow.com/a/12518877/1258206
func Exists(name string) bool {

	if _, err := os.Stat(name); err == nil {
		return true
	}

	return false
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
