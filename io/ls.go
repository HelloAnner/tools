package io_tools

import (
	"io/ioutil"
	"os"
)

// Ls 列出文件夹下的所有文件，不递归
func Ls(dirPath string) []string {
	var res []string

	if !IsDir(dirPath) {
		return res
	}

	infos, err := ioutil.ReadDir(dirPath)

	if err != nil {
		return res
	}

	for _, info := range infos {
		res = append(res, info.Name())
	}

	return res

}

func LsWithFilter(dirPath string, allow func(name string) bool) []string {
	var res []string

	allNames := Ls(dirPath)

	for _, name := range allNames {
		if allow(name) {
			res = append(res, name)
		}
	}

	return res
}

func IsDir(filePath string) bool {
	f, err := os.Open(filePath)

	if err != nil {
		return false
	}

	stat, err := f.Stat()

	if err != nil {
		return false
	}

	return stat.IsDir()
}
