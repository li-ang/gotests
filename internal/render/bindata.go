// Code generated by go-bindata.
// sources:
// templates/call.tmpl
// templates/function.tmpl
// templates/header.tmpl
// templates/inline.tmpl
// templates/inputs.tmpl
// templates/message.tmpl
// templates/results.tmpl
// DO NOT EDIT!

package render

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

var _templatesCallTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x4c\x8e\x51\xca\xc2\x40\x0c\x84\xaf\x12\x4a\x1f\xfe\x1f\x4a\x0e\x20\x78\x00\x5f\x44\x54\xf4\x79\xd9\xa6\x1a\xa8\xab\xa4\x51\x91\x90\xbb\xdb\x2d\xd5\xf5\x29\x30\xf3\x4d\x66\xcc\x5a\xea\x38\x11\x54\x31\xf4\x7d\xe5\x6e\xf6\x64\x3d\x03\x6e\x29\x12\x3f\x48\xb2\xc2\x1d\xa4\xab\x02\xae\x86\x9d\xca\x3d\xaa\xbb\x2a\x9a\x51\x6a\xb3\xfb\x21\x01\xdd\x8b\x8a\xeb\x70\x21\xf7\x3f\x33\x09\xe9\x44\x50\x73\x03\x35\xf5\xb0\x58\x02\x6e\x82\x8c\xa6\x92\x0c\xf3\xf7\x9a\xdd\x1b\xf8\x66\x4b\xdf\x51\x58\xf3\x86\xdf\xbe\x29\x9d\xcb\x26\x10\xf7\xaf\x1b\x8d\xe4\x21\x08\x87\x96\xe3\xb8\x01\x0b\x3b\x9d\xff\xf9\xbe\x03\x00\x00\xff\xff\x9d\x9f\x57\x19\xec\x00\x00\x00")

func templatesCallTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesCallTmpl,
		"templates/call.tmpl",
	)
}

func templatesCallTmpl() (*asset, error) {
	bytes, err := templatesCallTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/call.tmpl", size: 236, mode: os.FileMode(420), modTime: time.Unix(1455084885, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesFunctionTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x95\xc1\x6e\xdb\x30\x0c\x86\xcf\xce\x53\xb0\x41\x5b\xd8\x83\xa1\xde\x5b\xf4\xb0\xa2\xdd\xd0\xc3\x96\x22\x0d\xd6\xc3\x30\x0c\xaa\x4d\x67\x02\x14\x39\x93\xe4\x14\x81\xe1\x77\x1f\x25\x5b\x8e\x9b\xc4\xcd\x4e\x8e\x28\xf2\x13\xc9\x9f\x52\xea\x3a\xc7\x42\x28\x84\x69\x51\xa9\xcc\x8a\x52\x4d\x9b\x06\xea\xfa\xbc\x80\xeb\x5b\x60\x4d\x33\x99\xb8\x0d\xb2\xb0\x05\x1a\xfb\x9d\xaf\xb0\x69\x62\x0b\x9f\x2c\xad\x84\x5a\xb2\x45\x02\xf5\x24\x72\x2b\xe3\x22\x7e\xfe\x32\x56\x57\x99\x85\x9a\x42\x44\x01\xaa\xb4\x70\x5e\xb0\x27\x2d\x94\x7d\x54\xeb\xca\x1a\x42\x46\x91\x22\x0e\x90\x27\x11\xc8\x0f\x55\xee\x0f\x7d\x13\xf6\x0f\xb0\x39\x66\x28\x36\xa8\xbd\x89\x10\xec\xd1\x3c\x7b\xa6\x37\x68\xae\x96\x08\xec\x8b\x40\x99\xb7\xac\xba\xf6\x0b\x97\xac\x4f\x73\xbb\x46\xff\x2b\x50\x51\x1a\xec\x1c\x03\x7a\xdc\xb7\xfb\x76\xa7\xb8\x92\x9f\xb8\xa6\x64\x2d\xea\x70\x9a\x37\x8c\x12\x06\x91\x73\x34\x95\xb4\x21\xec\x85\x2b\x3b\x1a\xe5\xca\x9c\xa3\xad\xb4\x32\x0f\x5a\x97\xda\xc7\xbc\x51\x04\xad\xe0\xb5\x2c\x65\xf0\x9d\x44\x0d\xb5\x3b\xba\xba\x82\xc5\xec\x7e\x76\x0d\x9f\xf3\x1c\x5c\xf3\x21\xe3\x06\x0d\xa3\xed\x49\x54\x94\x1a\x7e\xa7\x60\xad\x13\xa4\xcd\xa7\xd5\xa7\x3e\xd2\xe3\xfd\x16\x1f\xf6\x89\x20\xad\x93\xcb\xda\x7b\x72\x8a\xbb\xec\x12\xea\xaa\x61\x3f\xb8\xac\xa8\xa6\x7a\x44\x22\xd6\x4e\xce\x35\x65\xc5\x06\x82\xa5\xbb\xba\xa2\x0f\x74\x18\x6a\x10\x12\x7e\xd1\xc2\xa2\x3e\xd0\x84\xb2\xbd\x7c\xdd\x52\xbd\xec\xae\x2a\x0a\xd4\xf5\x31\x2c\x11\xb8\xca\x21\x76\xd3\xc9\x66\x4a\x6e\x87\xbd\x4f\x0e\xed\x33\x85\xbe\xbe\x04\xba\xf3\x2c\xae\xd6\x92\x5b\xba\x36\xba\x15\x79\x4a\x53\xee\xd1\xbb\x9d\x8c\x4b\xd9\x9b\x3f\x14\x9a\x8c\xed\xce\x7e\x2a\x14\x82\x34\x00\x5e\x81\x63\xe0\x9b\x9e\x1c\x3b\xbf\xb3\x5b\x50\x42\x26\xee\x4b\x6d\x0e\xe3\xe3\xe6\x25\xb2\xcc\x03\x8b\x78\x3a\x24\xad\xd0\x18\xbe\xc4\x2e\x4b\x74\x1e\x70\x0b\x17\x9b\x14\x42\xf0\xc5\x66\x9a\xbe\x3b\x5c\xf8\x4b\xbc\x8b\x48\x07\x47\x25\x61\x50\xf6\x86\x3f\xca\x4a\x45\x8f\x45\x85\x23\x6a\x1f\xbb\x33\x70\x44\x67\xdf\xa7\xaf\xa5\xdd\x8d\x65\xaf\x3b\x7b\xf6\x8f\x49\x9c\xdc\x0c\x5c\xda\x3e\x0c\xaf\x5e\xf7\x1c\x40\x8b\xbe\xe3\x46\x64\xed\x5d\x0c\xf4\x61\xa1\x92\xde\xc5\x5e\xc0\x93\xcc\x93\x8c\x33\x8d\x85\xc4\xcc\xb2\x7b\xc4\xf5\xc3\xdf\x8a\xcb\xb8\xc7\xa6\xef\xa1\x49\x4b\xed\x5a\xf5\x5f\xe2\xf9\x76\xd1\x4b\xdb\x8d\xcf\x37\xea\xa2\x58\x4b\x77\x23\xfb\xd4\x3b\xe2\x4e\xe0\x13\xea\x8e\x66\x37\x94\xcf\xbd\x38\xf4\x27\x11\x56\xff\x02\x00\x00\xff\xff\xdf\x98\x64\x59\x50\x06\x00\x00")

func templatesFunctionTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesFunctionTmpl,
		"templates/function.tmpl",
	)
}

func templatesFunctionTmpl() (*asset, error) {
	bytes, err := templatesFunctionTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/function.tmpl", size: 1616, mode: os.FileMode(420), modTime: time.Unix(1455478039, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesHeaderTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xaa\xae\x4e\x49\x4d\xcb\xcc\x4b\x55\x50\xca\x48\x4d\x4c\x49\x2d\x52\xaa\xad\xe5\x2a\x48\x4c\xce\x4e\x4c\x4f\x55\xa8\xae\xd6\x0b\x80\x30\x81\x82\x5c\x99\xb9\x05\xf9\x45\x25\x0a\x1a\x5c\xd5\xd5\x45\x89\x79\x40\x69\x3d\x4f\xb0\x48\x71\x6d\x2d\x50\xa1\x5f\x62\x2e\x50\x15\x44\x4b\x49\x06\x50\x7d\x75\x75\x6a\x5e\x0a\x90\xd6\x84\xb3\x00\x01\x00\x00\xff\xff\x81\x22\x53\x6f\x6b\x00\x00\x00")

func templatesHeaderTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesHeaderTmpl,
		"templates/header.tmpl",
	)
}

func templatesHeaderTmpl() (*asset, error) {
	bytes, err := templatesHeaderTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/header.tmpl", size: 107, mode: os.FileMode(420), modTime: time.Unix(1454818382, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesInlineTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xaa\xae\x4e\x49\x4d\xcb\xcc\x4b\x55\x50\xca\xcc\xcb\x01\xd2\x4a\xb5\xb5\xd5\xd5\x99\x69\x0a\x7a\xfe\x79\x39\x95\x41\xa9\x25\xa5\x45\x79\xc5\xfe\x79\xa9\x61\x89\x39\xa5\xa9\xb5\xb5\xe9\xf9\x25\x0a\x56\xb6\x0a\xd5\xd5\x25\xa9\xb9\x05\x39\x89\x25\x40\x6d\xc9\x89\x39\x39\x4a\x0a\x7a\xb5\xb5\xd6\x40\xe1\xd4\xbc\x14\x90\x7e\x30\x05\x08\x00\x00\xff\xff\xfc\x48\x08\x64\x5a\x00\x00\x00")

func templatesInlineTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesInlineTmpl,
		"templates/inline.tmpl",
	)
}

func templatesInlineTmpl() (*asset, error) {
	bytes, err := templatesInlineTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/inline.tmpl", size: 90, mode: os.FileMode(420), modTime: time.Unix(1454818382, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesInputsTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x4c\xcb\x41\x0a\x02\x51\x08\xc6\xf1\xab\xc8\x30\xcb\x78\x07\x08\x3a\x40\xbb\xae\xf0\x60\x3e\x43\x28\x09\xc7\x56\x1f\x73\xf7\x46\x57\xad\x94\x9f\x7f\xc9\x0d\x6a\x0e\x59\xcc\x3f\xdf\xdc\x97\xe3\x20\x57\x95\xeb\x4d\x46\xad\xa6\xb2\xea\x78\x84\x79\xde\x3b\x28\x8c\xe9\x4f\xb4\xcf\x98\x6f\x24\xe2\xe4\xcc\x41\x36\xd4\xe7\x45\x48\xf8\x56\x35\x5e\x3b\xfa\xec\x67\xfb\xef\x35\x7e\x01\x00\x00\xff\xff\x43\x89\x5c\xae\x80\x00\x00\x00")

func templatesInputsTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesInputsTmpl,
		"templates/inputs.tmpl",
	)
}

func templatesInputsTmpl() (*asset, error) {
	bytes, err := templatesInputsTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/inputs.tmpl", size: 128, mode: os.FileMode(420), modTime: time.Unix(1454818382, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesMessageTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x54\x8e\xd1\xca\x82\x40\x10\x46\x5f\x65\x10\x85\xff\x07\x99\x07\x08\x7a\x80\x6e\x22\x22\xba\x5f\xf2\xd3\x06\x74\xb3\xdd\xd5\x88\x61\xde\x3d\x15\x13\xba\xfa\xe0\x70\x38\x33\xaa\x15\x6a\xf1\xa0\xac\x43\x8c\xae\x41\x66\xa6\x2a\x35\xf9\x47\x22\x3e\x05\xf1\xe9\xe0\xfb\x21\x45\xb3\xe2\xc9\xa4\x0a\x5f\xcd\xc6\x4b\xd2\x9d\xf8\x8c\x1b\x64\x44\x98\x09\x5f\xde\x3d\xf8\xea\xda\x01\x66\xbc\x89\x7c\x74\xdd\x04\xfe\x96\xe8\x6f\x50\x35\x38\xdf\x80\x72\x29\x29\x47\x4b\xbb\xfd\x24\xb8\x30\xf9\x09\x21\xae\x7f\xe4\x62\x56\x7e\xef\x16\xe3\xd6\x5d\xe6\x7f\xdd\x4f\x00\x00\x00\xff\xff\xeb\x6d\x22\x24\xc6\x00\x00\x00")

func templatesMessageTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesMessageTmpl,
		"templates/message.tmpl",
	)
}

func templatesMessageTmpl() (*asset, error) {
	bytes, err := templatesMessageTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/message.tmpl", size: 198, mode: os.FileMode(420), modTime: time.Unix(1455085158, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesResultsTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x5c\xcd\x41\x0a\xc2\x40\x0c\x05\xd0\xab\x7c\x4a\x97\xa5\x07\x10\x5c\x8a\x7b\x6f\x20\x34\x95\x81\x92\x81\x3f\xd3\x55\xc8\xdd\x4d\x6a\x51\x70\x35\x93\xfc\x97\xc4\x6c\x91\xb5\xa8\x60\xa0\xb4\x7d\xeb\x6d\x70\x87\x19\x9f\xfa\x12\x8c\x65\xc2\x28\x1b\x2e\x57\xcc\x8f\x4f\xec\x6e\x56\xd6\x48\xdc\xa7\x70\xa2\x4b\x76\xee\xb5\x63\xce\xcf\x59\x87\x88\x81\xbe\x53\xdb\x8d\xac\x4c\x2c\xe4\x99\xe3\x00\x95\xdf\xa5\xff\x38\x0f\xfe\xec\xf1\xbe\x03\x00\x00\xff\xff\xb0\x4f\xcf\x61\xa8\x00\x00\x00")

func templatesResultsTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesResultsTmpl,
		"templates/results.tmpl",
	)
}

func templatesResultsTmpl() (*asset, error) {
	bytes, err := templatesResultsTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/results.tmpl", size: 168, mode: os.FileMode(420), modTime: time.Unix(1454818382, 0)}
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
	"templates/call.tmpl": templatesCallTmpl,
	"templates/function.tmpl": templatesFunctionTmpl,
	"templates/header.tmpl": templatesHeaderTmpl,
	"templates/inline.tmpl": templatesInlineTmpl,
	"templates/inputs.tmpl": templatesInputsTmpl,
	"templates/message.tmpl": templatesMessageTmpl,
	"templates/results.tmpl": templatesResultsTmpl,
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
	"templates": &bintree{nil, map[string]*bintree{
		"call.tmpl": &bintree{templatesCallTmpl, map[string]*bintree{}},
		"function.tmpl": &bintree{templatesFunctionTmpl, map[string]*bintree{}},
		"header.tmpl": &bintree{templatesHeaderTmpl, map[string]*bintree{}},
		"inline.tmpl": &bintree{templatesInlineTmpl, map[string]*bintree{}},
		"inputs.tmpl": &bintree{templatesInputsTmpl, map[string]*bintree{}},
		"message.tmpl": &bintree{templatesMessageTmpl, map[string]*bintree{}},
		"results.tmpl": &bintree{templatesResultsTmpl, map[string]*bintree{}},
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
