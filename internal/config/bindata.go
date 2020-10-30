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

var _configIni = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x54\x5f\x6f\xe2\x46\x10\x7f\xdf\x4f\x31\x8f\x97\x4a\x67\x20\xd7\x56\x34\x88\x07\x27\x70\x1c\xaa\xc1\x08\x87\x6b\x55\x54\xa1\xc5\x1e\xec\x55\x96\x5d\xdf\xce\x18\x4a\x3f\x7d\xb5\x6b\x68\x49\xd4\xe4\x02\x2f\xd6\xf8\xf7\x67\x76\x7e\xe3\x5d\x13\xba\x03\xba\x3f\xc5\x00\xe2\xa2\x70\x48\x04\xd6\xc0\xb1\x52\x79\x05\x5c\x21\xb4\xaf\x41\x2b\x62\x34\x14\x89\x64\x9a\x3d\x8e\xe7\x9b\x78\x34\x5a\x8e\xb3\x0c\x86\xd0\x8d\xc2\x5f\x88\x01\x2c\xac\xe3\xb7\xd9\x8b\x74\xf9\x08\x43\xe8\x77\xfb\x3d\x21\xd6\x85\x64\xb9\x95\x84\xde\x7c\x74\x7e\x06\x42\x66\x65\x4a\x8a\xe0\xb3\x75\x60\xec\x11\xac\xd1\x27\xa8\x2d\x71\xe9\x90\xbe\x69\x50\x04\xd4\xd4\xb5\x75\x8c\x45\x24\xbe\xa4\x99\x97\xec\xfd\x72\x1b\xf5\x7e\xee\x47\xbd\xe8\xf6\xf6\x62\xf3\xd3\x8f\x9f\x3e\x89\x55\x36\x5e\xc2\xf0\x5f\xbe\x58\xc4\x59\xf6\x5b\xba\x1c\x5d\xd7\x46\xf7\x9b\x79\x3c\x1b\x5f\x97\xc4\x00\xb2\x2c\x81\xbd\x2d\x10\xd8\xc2\x16\xa1\x21\x2c\xe0\xa8\xb8\xba\xea\x25\x82\xaf\x52\xab\x22\xc0\x08\xa4\xc3\x3b\x31\x80\x1f\xa0\x50\x24\xb7\x1a\xe1\x23\xcc\xad\x97\x09\x45\x87\xdf\x1a\xe5\x7c\x31\xd6\x47\x79\xa2\xa0\xff\x81\x9e\x54\x0d\x07\x74\x6a\xa7\x72\xc9\xca\x9a\x9b\x00\x0e\x95\xd3\xc7\x5c\xbe\x80\xb7\x75\xe0\x4a\x72\x18\x70\x8e\x8e\x5b\x2a\x42\xed\x90\xd0\x30\x16\xb0\xf5\x08\x14\x03\x80\x4b\x02\x47\x49\x40\xaa\x34\xed\x4b\x09\xec\x1a\xf2\xc8\x87\xf8\x99\xdf\xae\xd1\xfa\xdd\x8e\xca\x9a\x67\x9e\xc1\xef\x2a\xf5\xd7\x3d\x41\x9a\xe2\x1a\x59\x59\x62\x30\x72\xdf\xb6\xbc\x97\x9c\x57\x48\x01\x60\x0d\x82\x32\x2f\x8f\x7a\x23\xb2\x2c\xd9\xcc\xd2\x91\x8f\xec\x3c\x6c\x21\xd6\xda\x96\xe4\x77\x69\xa6\x8c\xda\x37\x7b\xd0\xb6\x04\x8d\x07\xd4\x91\x48\xd2\xc9\x26\x19\x7f\x1d\x27\x7e\x63\x43\xb8\x95\x6d\x74\xe1\x21\xe4\xc3\x3d\x3a\xc5\x8c\xc6\x47\xbd\x53\x1a\x23\xf1\x79\x9a\x8c\x37\x49\x3a\x99\x4c\xe7\x13\x18\xfa\xde\xf1\x6d\x5a\x6e\x0d\x59\x8d\xf0\x81\xb8\xb0\x0d\xdf\x44\xe2\x21\x9d\x67\xe9\xff\xaa\x2c\x24\x57\x9e\xe3\x1b\xf4\x76\x04\x0e\xb5\x64\x75\x08\xbb\x76\xb4\xee\x49\x99\x12\x0a\xe5\x30\x67\xeb\x4e\x91\x18\xc0\x94\x21\x97\x06\xa4\xa6\xb0\x8c\x54\x63\xae\x76\x0a\x0b\x90\x04\x72\x4b\x56\x37\x7e\x03\x24\x57\xed\x59\x17\xf1\xe3\x17\x18\x86\x3e\x3b\x0d\xa1\xa3\x48\xdb\xd2\x5b\x0f\x5f\xfb\xf9\xc3\x9d\x3f\x3d\xd8\x59\x17\xfa\xf2\xfc\x52\x99\xf2\x3b\xbc\x99\xfc\x0b\x48\xfd\x8d\x60\x77\x21\xa9\xcb\xb1\x60\x8b\x3b\xeb\x10\x14\x13\x38\xcb\xd2\x87\xaf\x0c\xcc\xee\x23\x31\x8b\x7f\xdf\x64\xd3\x3f\x7c\x7e\xbd\xee\x59\x42\x96\xef\x55\x28\xe4\x89\x5a\x8d\x78\x72\x2d\x11\x52\x37\xcd\x7e\x8b\xce\x2b\x5d\x18\xff\x8d\x99\x2d\x3c\x21\xd6\x2d\xf7\x3e\x7e\xf8\x75\xb5\xc8\x5a\xbe\x58\xef\x55\xee\xac\x5f\x48\x95\x63\x58\xa3\x95\xd3\xed\x28\x2c\x57\xe8\xda\xef\xff\x19\x28\xdc\x2d\xd9\x66\xb5\xf4\x5b\x55\x31\xd7\x77\x9d\xce\xf5\x3d\x74\xd7\xef\xf6\xbb\x1d\x59\xab\xce\xa1\xf7\x4f\x00\x00\x00\xff\xff\x48\x08\x64\x79\x68\x05\x00\x00")

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

	info := bindataFileInfo{name: "config.ini", size: 1384, mode: os.FileMode(436), modTime: time.Unix(1604060782, 0)}
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
