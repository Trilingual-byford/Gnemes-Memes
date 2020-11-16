package utils

import (
	"errors"
	"github.com/corona10/goimagehash"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
)

func GetPicHash(r io.Reader, extension string) (string, error) {
	var decodedFile image.Image
	var err error
	switch extension {
	case ".jpg":
		decodedFile, err = jpeg.Decode(r)
	case ".jpeg":
		//contentType = "image/jpeg"
		decodedFile, err = jpeg.Decode(r)
	case ".gif":
		//contentType = "image/gif"
		decodedFile, err = gif.Decode(r)
	case ".png":
		//contentType = "image/png"
		decodedFile, err = png.Decode(r)
	default:
		return "", errors.New("this extension is invalid")
	}
	if err != nil {
		return "", err
	}
	hash, hashErr := goimagehash.AverageHash(decodedFile)
	if hashErr != nil {
		return "", err
	}
	return hash.ToString(), nil

}
