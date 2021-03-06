// Code generated by go-bindata.
// sources:
// config/config.toml
// DO NOT EDIT!

package config

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _configConfigToml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x53\x4b\x6f\xdb\x3c\x10\xbc\xf3\x57\x2c\xe4\xcb\xf7\x01\x8d\x2c\x4b\x4e\xe2\x0a\xf0\x21\x08\x72\x48\xd1\xb4\x40\x72\x34\x82\x62\x25\xad\x25\xc2\x7c\x08\x24\xe5\x3c\x7e\x7d\xb1\xb4\xe5\x44\x6d\x2e\x05\x22\x1f\x68\x72\xf6\x31\x33\x5c\xce\xe0\xda\xf6\x2f\x4e\xb6\x5d\x80\xff\xea\xff\x21\xcf\x16\x05\x9c\xf1\xb2\x82\x4a\x61\xbd\x0b\xb6\x87\x6f\xd6\x77\x03\xc2\x1d\x4a\x43\x5f\xe0\x4a\x29\xb8\xe7\x04\x0f\xf7\xe4\xc9\xed\xa9\x49\xc5\x0c\x1e\x88\xe0\xfb\xed\xf5\xcd\x8f\x87\x1b\xd8\x5a\x07\x4a\xd6\x64\x3c\x81\x34\x5b\xeb\x34\x06\x69\x4d\x2a\xc4\xec\x73\x3e\x31\x83\xbb\x2b\xee\x06\xd7\xd6\x6c\x65\x3b\xb8\xd8\x00\xfe\xbd\xce\x27\xf1\x11\x41\x06\x45\xb0\x86\xe4\x0e\x59\x39\xdc\x0f\x26\x48\x4d\x53\x7e\x89\xd8\x93\xf3\x4c\x74\x0d\xc9\x3e\x4b\x8b\x34\x5f\x26\x42\x6c\x70\x08\x9d\x75\x8f\x02\xc0\xa0\x8e\x55\x46\xef\x13\x01\x60\x5d\x8b\x46\xbe\x1e\x14\x9e\x3a\xdc\xfe\xe4\xcc\x27\xaa\x38\x6d\x70\x8a\x91\x2c\x8d\xbf\x72\x95\x71\x1e\x36\x5a\x9a\x5f\x47\x68\x91\x5f\x46\x70\x51\x16\x45\x51\x70\x2a\x69\x94\x8a\x93\x3b\xeb\x03\x87\x78\x1d\xfa\x94\x9e\x51\xf7\x8a\xd2\xda\x6a\xae\xd1\x5b\xc7\x58\x7e\xce\x4d\x3c\x39\x8e\xe3\x95\x79\x46\x1c\xbd\xe7\x33\x5e\x9f\xac\x6b\xb8\x70\x83\x01\x2b\xf4\xf4\x5e\x8f\x8e\x9c\xcf\x48\xa1\x0f\xb2\xe6\x4c\xa9\xb1\x7d\x07\xcd\x8f\x90\x27\x74\x75\x57\x5e\xa4\xcb\xe4\x4d\x57\x17\x42\x5f\xce\xe7\xca\xd6\xa8\x98\x6d\xf9\x35\xcf\xa2\xc4\xd9\x1f\x11\xd3\x22\x63\xd4\x48\x98\x03\x47\xd2\x4c\xf6\xb4\xb7\x2e\xb0\x8a\x0d\x27\x30\x6b\xbe\x39\x3b\x44\xe1\x99\x00\x20\x83\x95\x22\x0e\x0f\x6e\x20\x21\x36\x83\xfc\x40\xdb\x4e\x56\x68\xf0\x23\x69\x07\x64\xd4\x14\x9f\x4c\x34\xf2\xa4\x67\x42\x62\xb9\x2c\x1e\x3f\x6a\x4a\x66\x2f\x9d\x35\x9a\x4c\x60\xdc\x0d\x71\x18\x1a\xda\x93\xb2\x3d\x9f\x46\xef\x6d\xbd\xa3\x38\x49\x1a\xeb\x4e\x1a\x3a\x9b\xb2\x4c\x62\xe5\xa6\xb7\xd2\xc4\x3b\x0f\xf5\xd4\xd8\xbc\xb8\xbc\x48\x26\x0e\x2c\xa2\x05\x95\x34\x8d\x7f\x2b\x53\xce\x35\xaa\x27\x74\x54\x3a\xcb\xe1\x4a\x9a\x9d\xff\xfb\x9e\xcb\xc9\x7d\x70\x60\xdd\x0f\xb0\x86\xf3\xec\xf8\x31\x4f\xd2\xd6\xbd\xf0\x61\xbe\xcc\x57\x2b\x3e\x14\x1b\x65\xdb\xf6\x20\x63\x2b\x15\x4d\x25\xa4\xca\xb6\x49\x14\xf8\xec\xe5\x2b\x03\x8b\xec\xb0\x3d\xb8\x5e\x1c\x77\x15\xd6\xbb\xa1\x67\x56\x97\xcc\x90\x25\xc6\x17\xb9\x86\x2d\x2a\xcf\x8e\xf6\xce\x3e\xbf\xbc\x79\x7d\x42\x00\x78\x9c\xc6\xe9\xe0\xff\xfe\xb0\xf9\x1d\x00\x00\xff\xff\x43\x6c\x81\x8d\x2e\x05\x00\x00")

func configConfigTomlBytes() ([]byte, error) {
	return bindataRead(
		_configConfigToml,
		"config/config.toml",
	)
}

func configConfigToml() (*asset, error) {
	bytes, err := configConfigTomlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/config.toml", size: 1326, mode: os.FileMode(420), modTime: time.Unix(1536524850, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"config/config.toml": configConfigToml,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"config": &bintree{nil, map[string]*bintree{
		"config.toml": &bintree{configConfigToml, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

