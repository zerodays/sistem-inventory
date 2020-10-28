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

var _configIni = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x53\x4d\x6f\xe2\x48\x10\xbd\xf7\xaf\x78\xc7\x64\xa5\x78\xd9\x5b\xb4\x88\x83\x13\x08\x41\x6b\x30\xc2\x49\x76\xb5\xd1\x08\x35\x76\x61\xb7\xd2\x74\x93\xae\x36\x0c\xf3\xeb\x47\xdd\x26\x33\x24\x9a\x64\x02\x17\xab\xfc\x3e\xaa\xea\x95\x1f\x99\xdc\x8e\xdc\x17\xd1\x47\x5a\x55\x8e\x98\x61\x0d\xf6\x8d\x2a\x1b\xf8\x86\xd0\xbd\x86\x56\xec\xc9\x70\x22\xb2\x49\x71\x37\x9a\x2d\xd3\xe1\x70\x31\x2a\x0a\x0c\xd0\x4b\xe2\x5f\x88\x3e\xe6\xd6\xf9\x8f\xd9\xf3\x7c\x71\x87\x01\x2e\x7b\x97\x3d\x21\x1e\x2b\xe9\xe5\x4a\x32\x05\xf3\xe1\xf1\x19\x4c\xde\x2b\x53\x73\x82\x1b\xeb\x60\xec\x1e\xd6\xe8\x03\xb6\x96\x7d\xed\x88\x9f\x35\x14\x83\xdb\xed\xd6\x3a\x4f\x55\x22\x6e\xf3\x22\x48\x6a\x5b\x4a\xdd\x58\xf6\xe2\xbe\x18\x2d\x30\xf8\x41\x10\xf3\xb4\x28\xfe\xcd\x17\xc3\xd3\xda\xf0\x6a\x39\x4b\xa7\xa3\xd3\x92\xe8\xa3\x28\x32\x6c\x6c\x45\xf0\x16\x2b\x42\xcb\x54\x61\xaf\x7c\x73\x62\x9e\xe0\x41\x6a\x55\x45\x18\x43\x3a\xfa\x5b\xf4\xf1\x07\x2a\xc5\x72\xa5\x09\x17\x98\xd9\x20\x13\x8b\x8e\x9e\x5b\xe5\x42\x31\xd5\x7b\x79\xe0\xa8\x7f\xc6\x4f\x6a\x8b\x1d\x39\xb5\x56\xa5\xf4\xca\x9a\xf3\x08\x8e\x95\xc3\x45\x29\xdf\xc0\xbb\x3a\x7c\x23\x7d\xdc\x68\x49\xce\x77\x54\xc2\xd6\x11\x93\xf1\x54\x61\x15\x10\x24\xfa\xc0\xcb\xca\xf7\x92\xc1\xaa\x36\xdd\x4b\x09\xef\x5a\x0e\xc8\xeb\xf4\x95\xdf\xba\xd5\xfa\xd3\x8e\xca\x9a\x57\x9e\xd1\xef\x24\xe6\xf7\x3d\x21\x4d\x75\x8a\x0c\x41\xc1\xc8\x4d\xd7\xf2\x46\xfa\xb2\x21\x8e\x00\x6b\x08\xca\xbc\x1d\xf5\x5c\x14\x45\xb6\x9c\xe6\xc3\x10\xd9\x71\xd9\x42\x3c\x6a\x5b\x73\x38\x9e\xa9\x32\x6a\xd3\x6e\xa0\x6d\x0d\x4d\x3b\xd2\x89\xc8\xf2\xf1\x32\x1b\x3d\x8c\xb2\x70\xa2\x31\xdc\xc6\xb6\xba\x0a\x10\x0e\xe1\xee\x9d\xf2\x9e\x4c\x88\x7a\xad\x34\x25\xe2\x66\x92\x8d\x96\x59\x3e\x1e\x4f\x66\x63\x0c\x42\xef\xf4\x31\xad\xb4\x86\xad\x26\x9c\xb1\xaf\x6c\xeb\xcf\x13\x71\x9d\xcf\x8a\xfc\x97\x2a\x73\xe9\x9b\xc0\x09\x0d\x06\x3b\x86\x23\x2d\xbd\xda\xc5\x5b\xdb\x5b\xf7\xa4\x4c\x8d\x4a\x39\x2a\xbd\x75\x87\x44\xf4\x31\xf1\x28\xa5\x81\xd4\x1c\x8f\x91\xb7\x54\xaa\xb5\xa2\x0a\x92\x21\x57\x6c\x75\x1b\x2e\x40\xfa\xa6\x9b\x75\x9e\xde\xdd\xc6\xcf\xa0\xe6\x3f\x5b\x26\xc7\x89\xb6\x75\xb0\x1e\xbc\xf7\x0b\xc3\x1d\xbf\x35\xac\xad\x8b\x7d\x05\x7e\xad\x4c\xfd\x1b\xde\x54\x7e\x05\xab\x6f\x04\xbb\x8e\x49\xbd\x8c\x85\x15\xad\xad\x23\x28\xcf\x70\xd6\xcb\x10\xbe\x32\x98\x5e\x25\x62\x9a\xfe\xb7\x2c\x26\xff\x87\xfc\xfe\xea\x1d\x25\x64\xfd\x59\x85\x4a\x1e\xb8\xd3\x48\xc7\xa7\x12\x31\x75\xd3\x6e\x56\xe4\x82\xd2\x0b\xe3\xe7\x9a\xbd\xc5\x13\xd1\xb6\xe3\x5e\xa5\xd7\xff\xdc\xcf\x8b\x8e\xff\x3d\x00\x00\xff\xff\xd9\xb1\x45\x5b\xf9\x04\x00\x00")

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

	info := bindataFileInfo{name: "config.ini", size: 1273, mode: os.FileMode(436), modTime: time.Unix(1603883902, 0)}
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
