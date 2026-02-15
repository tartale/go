package reflectx

import (
	"fmt"
	"path"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
)

func CallerPackagePath(skip int) string {

	if pc, _, _, ok := runtime.Caller(skip); ok {
		if fn := runtime.FuncForPC(pc); fn != nil {
			fname := fn.Name()
			dirname := filepath.Dir(fname)
			basename := filepath.Base(fname)
			pkgname := strings.SplitN(basename, ".", 2)[0]
			packagePath := path.Join(dirname, pkgname)

			return packagePath
		}
	}

	return ""
}

func ObjectTypePackagePath(obj any) string {

	return reflect.TypeOf(obj).PkgPath()
}

func ObjectTypePath(obj any) string {

	return fmt.Sprintf("%s.%s", ObjectTypePackagePath(obj), reflect.TypeOf(obj).Name())
}
