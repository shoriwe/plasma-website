package static

import (
	"embed"
	"github.com/spf13/afero"
	"io"
	"io/fs"
	"path/filepath"
)

var (
	//go:embed css
	css embed.FS
	//go:embed vendoring
	vendoring embed.FS
	//go:embed wasm
	wasm embed.FS
	//go:embed js
	js embed.FS
	//go:embed articles
	articles embed.FS
)

const (
	Path = "static"
)

func walkFunc(output afero.Fs, embeddedFs embed.FS) fs.WalkDirFunc {
	return func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}
		directory, _ := filepath.Split(path)
		mkdirError := output.MkdirAll(directory, 0777)
		if mkdirError != nil {
			return mkdirError
		}
		file, createError := output.Create(path)
		if createError != nil {
			return createError
		}
		defer file.Close()
		embeddedFile, _ := embeddedFs.Open(path)
		defer embeddedFile.Close()
		_, copyError := io.Copy(file, embeddedFile)
		return copyError
	}
}

func Write(output afero.Fs) error {
	writeToOutError := fs.WalkDir(css, ".", walkFunc(afero.NewBasePathFs(output, Path), css))
	if writeToOutError != nil {
		return writeToOutError
	}
	writeToOutError = fs.WalkDir(vendoring, ".", walkFunc(afero.NewBasePathFs(output, Path), vendoring))
	if writeToOutError != nil {
		return writeToOutError
	}
	writeToOutError = fs.WalkDir(js, ".", walkFunc(afero.NewBasePathFs(output, Path), js))
	if writeToOutError != nil {
		return writeToOutError
	}
	writeToOutError = fs.WalkDir(articles, ".", walkFunc(afero.NewBasePathFs(output, Path), articles))
	if writeToOutError != nil {
		return writeToOutError
	}
	return fs.WalkDir(wasm, ".", walkFunc(afero.NewBasePathFs(output, Path), wasm))
}
