// Code generated for package config by go-bindata DO NOT EDIT. (@generated)
// sources:
// configs/config.ini
// configs/migrations/01_initial.up.sql
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

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _configIni = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x53\x4f\x6f\xfa\x46\x10\xbd\xef\xa7\x78\xc7\x5f\x2a\xc5\xa5\x4a\x2b\x45\x45\x1c\x9c\x40\x08\xaa\xc1\x08\x27\x69\xd5\xa8\x42\x8b\x3d\xd8\xab\x2c\xbb\x64\x67\x0d\xa5\x9f\xbe\xda\x35\xb4\x24\x6a\xd2\xc0\xc5\x1a\xbf\x3f\x33\xf3\xc6\xcf\x4c\x6e\x47\xee\x0f\xd1\x47\x5a\x55\x8e\x98\x61\x0d\xf6\x8d\x2a\x1b\xf8\x86\xd0\xbd\x86\x56\xec\xc9\x70\x22\xb2\x49\xf1\x30\x9a\x2d\xd3\xe1\x70\x31\x2a\x0a\x0c\xd0\x4b\xe2\x5f\x88\x3e\xe6\xd6\xf9\xcf\xd9\xf3\x7c\xf1\x80\x01\xae\x7b\xd7\x3d\x21\x9e\x2b\xe9\xe5\x4a\x32\x05\xf3\xe1\xf1\x19\x4c\xde\x2b\x53\x73\x82\x3b\xeb\x60\xec\x1e\xd6\xe8\x03\xb6\x96\x7d\xed\x88\x5f\x35\x14\x83\xdb\xed\xd6\x3a\x4f\x55\x22\xee\xf3\x22\x48\x6a\x5b\x4a\xdd\x58\xf6\x27\x8f\x9f\x7e\xbc\xba\x12\x8f\xc5\x68\x81\xc1\x3f\x64\x31\x4f\x8b\xe2\xd7\x7c\x31\x3c\xaf\x0d\x6f\x96\xb3\x74\x3a\x3a\x2f\x89\x3e\x8a\x22\xc3\xc6\x56\x04\x6f\xb1\x22\xb4\x4c\x15\xf6\xca\x37\x67\x8d\x24\x78\x92\x5a\x55\x11\xc6\x90\x8e\x7e\x16\x7d\x7c\x87\x4a\xb1\x5c\x69\xc2\x25\x66\x36\xc8\xc4\xa2\xa3\xd7\x56\xb9\x50\x4c\xf5\x5e\x1e\x38\xea\x7f\xe3\x17\xb5\xc5\x8e\x9c\x5a\xab\x52\x7a\x65\xcd\x45\x04\xc7\xca\xe1\xb2\x94\xef\xe0\x5d\x1d\xbe\x91\x3e\x6e\xb7\x24\xe7\x3b\x2a\x61\xeb\x88\xc9\x78\xaa\xb0\x0a\x08\x12\x7d\xe0\xb4\xfe\xbd\x64\xb0\xaa\x4d\xf7\x52\xc2\xbb\x96\x03\xf2\x36\x7d\xe3\xb7\x6e\xb5\xfe\xb2\xa3\xb2\xe6\x8d\x67\xf4\x3b\x8b\xfc\x63\x4f\x48\x53\x9d\x23\x43\x68\x30\x72\xd3\xb5\xbc\x91\xbe\x6c\x88\x23\xc0\x1a\x82\x32\xef\x47\xbd\x10\x45\x91\x2d\xa7\xf9\x30\x44\x76\x5c\xb6\x10\xcf\xda\xd6\x1c\x0e\x69\xaa\x8c\xda\xb4\x1b\x68\x5b\x43\xd3\x8e\x74\x22\xb2\x7c\xbc\xcc\x46\x4f\xa3\x2c\x9c\x6b\x0c\xb7\xb1\xad\xae\x02\x84\x43\xb8\x7b\xa7\xbc\x27\x13\xa2\x5e\x2b\x4d\x89\xb8\x9b\x64\xa3\x65\x96\x8f\xc7\x93\xd9\x18\x83\xd0\x3b\x7d\x4e\x2b\xad\x61\xab\x09\xdf\xd8\x57\xb6\xf5\x17\x89\xb8\xcd\x67\x45\xfe\x9f\x2a\x73\xe9\x9b\xc0\x09\x0d\x06\x3b\x86\x23\x2d\xbd\xda\xc5\x5b\xdb\x5b\xf7\xa2\x4c\x8d\x4a\x39\x2a\xbd\x75\x87\x44\xf4\x31\xf1\x28\xa5\x81\xd4\x1c\x8f\x91\xb7\x54\xaa\xb5\xa2\x0a\x92\x21\x57\x6c\x75\x1b\x2e\x40\xfa\xa6\x9b\x75\x9e\x3e\xdc\xc7\x4f\xa2\xe6\xef\x5b\x26\xc7\x89\xb6\x75\xb0\x1e\x7c\xf4\x0b\xc3\x1d\xbf\x3b\xac\xad\x8b\x7d\x05\x7e\xad\x4c\xfd\x3f\xbc\xa9\xfc\x13\xac\xfe\x22\xd8\x75\x4c\xea\x34\x16\x56\xb4\xb6\x8e\xa0\x3c\xc3\x59\x2f\x43\xf8\xca\x60\x7a\x93\x88\x69\xfa\xdb\xb2\x98\xfc\x1e\xf2\xfb\xa1\x77\x94\x90\xf5\x57\x15\x2a\x79\xe0\x4e\x23\x1d\x9f\x4b\xc4\xd4\x4d\xbb\x59\x91\x0b\x4a\x27\xc6\xbf\x6b\xf6\x16\x2f\x44\xdb\x8e\x7b\x93\xde\xfe\xf2\x38\x2f\x3a\xfe\xdf\x01\x00\x00\xff\xff\x86\x7b\x39\xca\x05\x05\x00\x00")

func configIniBytes() ([]byte, error) {
	return bindataRead(
		_configIni,
		"config.ini",
	)
}

func configIni() (*asset, error) {
	bytes, err := configIniBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config.ini", size: 1285, mode: os.FileMode(436), modTime: time.Unix(1603994381, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrations01_initialUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8e\xcd\xaa\xc2\x30\x10\x85\xf7\x7d\x8a\xb3\xbc\x17\x7c\x89\x2e\x02\x06\xdb\xb4\xa4\x23\x5a\x37\x12\x93\x01\x03\xb6\x86\x24\xbe\xbf\x50\x7f\x08\xd2\xb3\x1b\xe6\x7c\x87\xcf\x46\x36\x99\x91\xcd\xe5\xc6\xf0\x99\xa7\x54\xfd\x55\x00\xe0\x1d\xca\x0c\x42\xcb\xba\x41\xaf\x65\x5b\xeb\x11\x3b\x31\x6e\x96\xda\x6c\x26\x2e\x6a\x24\x8e\x84\xb5\xa8\x8e\xa0\xf6\x4d\xf3\xa2\x1c\x27\x1b\x7d\xc8\xfe\x3e\x7f\xa8\xf7\xc3\x64\x3e\x87\x47\xb4\x57\x93\xd8\x81\x64\x2b\x06\xaa\xdb\x1e\x07\x49\xdb\xe5\xc4\xa9\x53\xe2\x67\x2e\x44\x6f\x0b\x0b\xa9\xd6\x1d\xbe\x54\xf5\xff\x0c\x00\x00\xff\xff\x41\x42\x5d\xfa\xf7\x00\x00\x00")

func migrations01_initialUpSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations01_initialUpSql,
		"migrations/01_initial.up.sql",
	)
}

func migrations01_initialUpSql() (*asset, error) {
	bytes, err := migrations01_initialUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/01_initial.up.sql", size: 247, mode: os.FileMode(436), modTime: time.Unix(1603908855, 0)}
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
	"config.ini":                   configIni,
	"migrations/01_initial.up.sql": migrations01_initialUpSql,
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
	"config.ini": &bintree{configIni, map[string]*bintree{}},
	"migrations": &bintree{nil, map[string]*bintree{
		"01_initial.up.sql": &bintree{migrations01_initialUpSql, map[string]*bintree{}},
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
