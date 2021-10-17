package utils

import (
	base642 "encoding/base64"
	"errors"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	ImageDir = "../profile"
)

type Base64Image struct {
	Data string `json:"data"`
	Type string `json:"type"`
}

func CreateImageDir() error {
	if _, err := os.Stat(ImageDir); os.IsNotExist(err) {
		return os.Mkdir(ImageDir, 0755)
	}
	return nil
}

func WriteImage(uid int64, image *Base64Image) error {
	decoded, err := base642.StdEncoding.DecodeString(image.Data)
	if err != nil {
		return err
	}
	var ext string
	switch image.Type {
	case "image/png":
		ext = "png"
		break
	case "image/jpg":
		ext = "jpg"
		break
	case "image/jpeg":
		ext = "jpeg"
		break
	}
	name := ImageDir + "/" + strconv.FormatInt(uid, 16) + "." + ext
	return ioutil.WriteFile(name, decoded, 0644)
}

func ReadImage(uid int64) (*Base64Image, error) {
	pre := strconv.FormatInt(uid, 16) + "."
	var imagePath string

	err := filepath.WalkDir(ImageDir + "/", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if strings.HasPrefix(d.Name(), pre) {
			imagePath = path
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	if len(imagePath) == 0 {
		return nil, nil
	}

	data, err := os.ReadFile(imagePath)
	if err != nil {
		return nil, err
	}

	image := &Base64Image{}

	ext := filepath.Ext(imagePath)
	switch ext {
	case ".png":
		image.Type = "image/png"
		break
	case ".jpeg":
		image.Type = "image/jpeg"
		break
	case ".jpg":
		image.Type = "image/jpg"
		break
	default:
		return nil, errors.New("invalid image format")
	}

	base64 := base642.StdEncoding.EncodeToString(data)
	image.Data = base64
	return image, nil
}
