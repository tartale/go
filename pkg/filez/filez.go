package filez

import "os"

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
