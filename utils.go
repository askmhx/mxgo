package mxgo

import "os"

func PathExist(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}
