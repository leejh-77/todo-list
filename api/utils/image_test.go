package utils

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestCreateDir(t *testing.T) {
	err := CreateImageDir()
	if err != nil {
		t.Fatal(err)
	}
	_, err = ioutil.ReadDir(ImageDir)

	assert.Nil(t, err)
}

func TestReadImage(t *testing.T) {
	image, err := ReadImage(int64(4))
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(t, image)
}
