package io

import (
	"fmt"
	"io"
	"log"
	"os"
)

// Copy 拷贝文件
func Copy(from, to string) (int64, error) {
	sourceFileStat, err := os.Stat(from)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", from)
	}

	source, err := os.Open(from)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(to)
	if err != nil {
		return 0, err
	}

	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	log.Printf("copy file %s --> %s\n", from, to)
	return nBytes, err
}
