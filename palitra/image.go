package main

import (
	"image"
	_ "image/jpeg"
	_ "image/png"
	"net/http"

	"io"
)

func getImageFromURL(URL string) (image.Image, error) {
	var imageReader io.Reader
	res, err := http.Get(URL)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	imageReader = res.Body

	var img image.Image
	if img, _, err = image.Decode(imageReader); err != nil {
		return nil, err
	}
	return img, nil
}
