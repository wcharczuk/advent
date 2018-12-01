package fileutil

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/blend/go-sdk/assert"
)

func TestRemoveMany(t *testing.T) {
	assert := assert.New(t)

	// create 3 files
	file0, err := ioutil.TempFile("", "")
	assert.Nil(err)
	assert.Nil(file0.Close())
	file1, err := ioutil.TempFile("", "")
	assert.Nil(err)
	assert.Nil(file1.Close())
	file2, err := ioutil.TempFile("", "")
	assert.Nil(err)
	assert.Nil(file2.Close())

	_, err = os.Stat(file0.Name())
	assert.Nil(err)
	_, err = os.Stat(file1.Name())
	assert.Nil(err)
	_, err = os.Stat(file2.Name())
	assert.Nil(err)

	assert.Nil(RemoveMany(file0.Name(), file1.Name(), file2.Name()))

	_, err = os.Stat(file0.Name())
	assert.NotNil(err)
	_, err = os.Stat(file1.Name())
	assert.NotNil(err)
	_, err = os.Stat(file2.Name())
	assert.NotNil(err)
}
