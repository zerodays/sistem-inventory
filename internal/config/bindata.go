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

var _configIni = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x54\x5f\x6f\xe2\x46\x10\x7f\xdf\x4f\x31\x8f\x97\x4a\x67\x88\xae\x95\xa2\x20\x1e\x9c\xc0\x71\xa8\x06\x23\x9c\x5c\xab\x46\x15\x5a\xec\xc1\x5e\x65\xd9\xf1\xed\x8c\xa1\xf4\xd3\x57\xbb\x0e\x57\x72\xea\x5d\x03\x2f\xd6\xf8\xf7\x67\x76\x7e\xb3\x7e\x62\xf4\x07\xf4\x7f\xaa\x11\xa4\x55\xe5\x91\x19\xc8\xc1\xb1\x31\x65\x03\xd2\x20\xf4\xaf\xc1\x1a\x16\x74\x9c\xa8\x6c\x5e\x3c\x4c\x97\x9b\x74\x32\x59\x4f\x8b\x02\xc6\x30\x4c\xe2\x5f\xa9\x11\xac\xc8\xcb\x8f\xd9\xab\x7c\xfd\x00\x63\xb8\x19\xde\x5c\x2b\xf5\x54\x69\xd1\x5b\xcd\x18\xcc\x27\x2f\xcf\xc0\x28\x62\x5c\xcd\x09\x7c\x24\x0f\x8e\x8e\x40\xce\x9e\xa0\x25\x96\xda\x23\x7f\xb1\x60\x18\xb8\x6b\x5b\xf2\x82\x55\xa2\x3e\xe5\x45\x90\xb4\x54\x6a\xdb\x10\xcb\xd9\xe3\x97\x9f\x3f\x7c\x50\x8f\xc5\x74\x0d\xe3\xaf\x64\xb5\x4a\x8b\xe2\xb7\x7c\x3d\xb9\xac\x4d\xee\x36\xcb\x74\x31\xbd\x2c\xa9\x11\x14\x45\x06\x7b\xaa\x10\x84\x60\x8b\xd0\x31\x56\x70\x34\xd2\x5c\x34\x92\xc0\x67\x6d\x4d\x15\x61\x0c\xda\xe3\xad\x1a\xc1\x4f\x50\x19\xd6\x5b\x8b\xf0\x1e\x96\x14\x64\x62\xd1\xe3\x97\xce\xf8\x50\x4c\xed\x51\x9f\x38\xea\xbf\xe3\x67\xd3\xc2\x01\xbd\xd9\x99\x52\x8b\x21\x77\x15\xc1\xb1\x72\x7a\x5f\xea\x6f\xe0\x7d\x1d\xa4\xd1\x12\xa7\x5b\xa2\x97\x9e\x8a\xd0\x7a\x64\x74\x82\x15\x6c\x03\x02\xd5\x08\xe0\x3c\xfe\xa3\x66\x60\x53\xbb\xfe\xa5\x06\xf1\x1d\x07\xe4\x7d\xfa\xca\x6f\xd7\x59\xfb\x66\x47\x43\xee\x95\x67\xf4\xbb\x88\xfc\xfb\x9e\xa0\x5d\x75\x89\x0c\xa1\x81\xd3\xfb\xbe\xe5\xbd\x96\xb2\x41\x8e\x00\x72\x08\xc6\x7d\x7b\xd4\x2b\x55\x14\xd9\x66\x91\x4f\x42\x64\x2f\xc3\x56\xea\xc9\x52\xcd\x61\x91\x16\xc6\x99\x7d\xb7\x07\x4b\x35\x58\x3c\xa0\x4d\x54\x96\xcf\x36\xd9\xf4\xf3\x34\x0b\xeb\x1a\xc3\x6d\xa8\xb3\x55\x80\x70\x08\xf7\xe8\x8d\x08\xba\x10\xf5\xce\x58\x4c\xd4\xc7\x79\x36\xdd\x64\xf9\x6c\x36\x5f\xce\x60\x1c\x7a\xc7\x1f\xd3\x4a\x72\x4c\x16\xe1\x1d\x4b\x45\x9d\x5c\x25\xea\x3e\x5f\x16\xf9\x7f\xaa\xac\xb4\x34\x81\x13\x1a\x0c\x76\x0c\x1e\xad\x16\x73\x88\xbb\x76\x24\xff\x6c\x5c\x0d\x95\xf1\x58\x0a\xf9\x53\xa2\x46\x30\x17\x28\xb5\x03\x6d\x39\x2e\x23\xb7\x58\x9a\x9d\xc1\x0a\x34\x83\xde\x32\xd9\x2e\x6c\x80\x96\xa6\x3f\xeb\x2a\x7d\xf8\x14\xaf\x44\xcd\x83\x8e\xd1\x73\x62\xa9\x0e\xd6\xe3\xef\xfd\xc2\xe1\x5e\xee\x1d\xec\xc8\xc7\xbe\x02\xbf\x36\xae\xfe\x1f\xde\x42\xff\x05\x6c\xfe\x46\xa0\x5d\x4c\xea\x7c\x2c\xd8\xe2\x8e\x3c\x82\x11\x06\x4f\xa2\x43\xf8\xc6\xc1\xe2\x2e\x51\x8b\xf4\xf7\x4d\x31\xff\x23\xe4\x77\x3d\x7c\x91\xd0\xf5\x5b\x15\x2a\x7d\xe2\x5e\x23\x9d\x5d\x4a\xc4\xd4\x5d\xb7\xdf\xa2\x0f\x4a\x67\xc6\xbf\x63\x16\x82\x67\xc4\xb6\xe7\xde\xa5\xf7\xbf\x3e\xae\x8a\x9e\xaf\x9e\xf6\xa6\xf4\x14\x16\xd2\x94\x18\xd7\xe8\xd1\xdb\x7e\x14\x24\x0d\xfa\xfe\xfe\xbf\x02\xc5\x6f\x4b\xb1\x79\x5c\x87\xad\x6a\x44\xda\xdb\xc1\xe0\xeb\x47\xe8\xf6\x66\x78\x33\x1c\xe8\xd6\x0c\x0e\xd7\xff\x04\x00\x00\xff\xff\xf2\xb4\x5c\x0b\x62\x05\x00\x00")

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

	info := bindataFileInfo{name: "config.ini", size: 1378, mode: os.FileMode(436), modTime: time.Unix(1607011280, 0)}
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
