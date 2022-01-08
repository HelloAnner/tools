package compression_tools

import (
	"archive/zip"
	errortools "github.com/HelloAnner/tools/error"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func Zip(srcDir string) error {
	zipFileName := srcDir + ".zip"

	err := os.RemoveAll(zipFileName)
	if err != nil {
		return errortools.WithStack(err)
	}

	zipfile, err := os.Create(zipFileName)

	defer zipfile.Close()
	archive := zip.NewWriter(zipfile)
	defer archive.Close()

	err = filepath.Walk(srcDir, func(path string, info os.FileInfo, _ error) error {
		if path == srcDir {
			return nil
		}
		header, _ := zip.FileInfoHeader(info)
		header.Name = strings.TrimPrefix(path, srcDir+`\`)
		if info.IsDir() {
			header.Name += `/`
		} else {
			header.Method = zip.Deflate
		}
		writer, _ := archive.CreateHeader(header)
		if !info.IsDir() {
			file, _ := os.Open(path)
			defer file.Close()

			_, err := io.Copy(writer, file)
			if err != nil {
				return errortools.WithMessagef(err, "zip dir %s fail", srcDir)
			}
		}
		return nil
	})

	return errortools.Cause(err)
}
