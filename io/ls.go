package io_tools

import (
	errors "github.com/HelloAnner/tools/error"
	"io/ioutil"
	"os"
)

// Ls 列出文件夹下的所有文件，不递归
func Ls(dirPath string) ([]string, error) {
	var res []string

	if !IsDir(dirPath) {
		return res, errors.New(dirPath + " is not dir")
	}

	infos, err := ioutil.ReadDir(dirPath)

	if err != nil {
		return res, err
	}

	for _, info := range infos {
		res = append(res, info.Name())
	}

	return res, nil

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
