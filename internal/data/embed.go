package data

import (
	"embed"
	"encoding/json"
	"io"
	"io/fs"
)

//go:embed embedded/*
var internalData embed.FS

func GetBaseFS() (fs.FS, error) {
	return fs.Sub(internalData, "embedded")
}

func GetTexturesFS() (fs.FS, error) {
	return fs.Sub(internalData, "embedded/textures")
}

func Parse[T any](path string) (T, error) {
	var d T

	fileSystem, err := GetBaseFS()

	if err != nil {
		return d, err
	}

	file, err := fileSystem.Open(path)

	if err != nil {
		return d, err
	}

	bs, err := io.ReadAll(file)

	if err != nil {
		return d, err
	}

	err = json.Unmarshal(bs, &d)

	if err != nil {
		return d, err
	}

	return d, nil
}
