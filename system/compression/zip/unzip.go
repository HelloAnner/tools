package compression

import (
	"archive/zip"
	"fmt"
	errortools "github.com/HelloAnner/tools/error"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func Unzip(src, dest string) error {

	r, err := zip.OpenReader(src)

	if err != nil {
		return err
	}

	defer r.Close()

	for _, f := range r.File {
		fPath := filepath.Join(dest, f.Name)

		if !strings.HasPrefix(fPath, filepath.Clean(dest)+string(os.PathSeparator)) {
			return fmt.Errorf("%s: illegal file path", fPath)
		}

		if f.FileInfo().IsDir() {
			_ = os.MkdirAll(fPath, os.ModePerm)
			continue
		}

		if err = os.MkdirAll(filepath.Dir(fPath), os.ModePerm); err != nil {
			return err
		}

		outFile, err := os.OpenFile(fPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return errortools.Cause(err)
		}

		rc, err := f.Open()
		if err != nil {
			return errortools.Cause(err)
		}

		_, err = io.Copy(outFile, rc)

		_ = outFile.Close()
		_ = rc.Close()

		if err != nil {
			return errortools.Cause(err)
		}
	}

	return nil
}
