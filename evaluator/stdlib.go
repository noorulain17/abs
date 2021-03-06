package evaluator

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

var _stdlib_cli_index_abs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x54\x4d\x6f\xdb\x38\x14\xbc\xf3\x57\xcc\x4a\x0b\x58\xda\x78\x7d\xd8\x9c\x92\x85\xd1\xa6\x01\x0a\x04\x08\x7a\x68\x8f\xae\x51\xd0\xe2\x93\x4c\x98\xa2\x0c\x3e\xca\x49\x6a\xf8\xbf\x17\x94\x68\x5b\x4a\x9c\xf0\xa4\x8f\x37\xf3\x38\xf3\x86\x4c\x71\xff\xf8\x00\xb9\xdd\x8a\xc2\x68\xcc\xb1\x3f\x08\x91\xe2\xbe\xa9\x6b\x69\x15\xc3\x51\xa5\xd9\x93\x23\x85\x27\xed\xd7\xda\xc2\xaf\x35\x0f\x31\xb3\xe2\x58\x7b\x04\x7f\x6d\x6d\xe1\x75\x63\xd1\x32\x29\xf8\xe6\x44\x02\x89\x58\xdc\x03\x6b\x85\x39\xca\xcc\xca\x9a\xa6\x50\xc4\x85\xd3\xdb\x00\x9c\xa2\x34\xb2\xe2\x1c\x7b\x01\x00\x8e\x7c\xeb\x2c\xca\xac\xb4\xc7\x4f\x61\x0d\x7b\x2f\x02\xc7\xb2\xdb\xc1\xff\x1f\x14\x9c\x5a\x0e\x79\xc2\x4a\x71\xef\x48\x7a\xea\x1b\x77\x52\xa1\xa8\x94\xad\xf1\xd8\x49\xd3\x12\x8f\xca\xcb\xc6\x61\x33\xc5\x2f\x68\x1b\x11\x63\xba\xb0\x76\xa1\x91\x91\x55\xb6\xc9\xc5\x9b\x9f\xba\xc4\xee\x02\xa6\xe3\x0e\x7c\x8b\x4d\x10\xb3\x7b\x53\x70\x10\xef\xbf\xbd\x16\x24\x8d\x81\x5f\x13\x1a\xa7\x2b\x6d\xa5\x41\x51\xab\x51\x8d\x23\x0e\xfa\xe6\x28\xed\xac\x90\xc6\x64\x0b\xe9\x2a\xce\xf2\xc5\xf5\xed\x32\x8e\x60\xf9\x6a\xef\x29\x1e\xca\x40\xea\x68\xc2\x90\x47\x86\x27\xc2\xd6\x69\xeb\xa1\x3d\x9a\xd6\x8b\x57\x4a\x63\xd5\x5b\xb9\x54\xac\x9b\xac\xff\x9b\xbf\x23\xec\x20\x3e\x1a\xe7\x20\x33\x98\x0f\x13\x24\x7a\x6c\x17\xc7\xef\xad\xed\x7c\x18\x66\xd6\xb5\x76\x94\x83\x14\x77\x5f\x7e\x80\x89\x18\x89\x5c\x31\x7a\xa2\x59\x78\x7c\x7e\xf9\x9d\xc4\x1a\x6e\x3a\xa2\xb8\x09\x68\xee\x5e\xaf\x9d\x82\x74\x55\x5b\x93\xed\xa5\xf7\x21\x93\xae\xca\xfe\x8b\xf6\xa5\xf8\xd6\x78\x6c\x25\xb3\xb6\xd5\xf9\x14\x7c\xc2\x23\xf9\x09\x47\xf3\x02\xd7\x9a\xcc\x56\x44\xdb\xfe\x0a\x3c\x67\xd3\xe2\x29\x18\xb9\x30\x09\xf5\x93\x2e\xd6\x59\x2e\x06\x7e\x75\xf0\x61\x65\x51\xab\xe5\x80\x8c\x9e\xb5\xcf\x6e\x6e\xa6\x48\x8e\x62\x26\x7f\xef\x8b\x5a\x1d\x26\xb0\x8d\x47\xd9\xb4\x56\x25\x23\xc6\x4b\xed\x03\x69\xec\xdd\x39\x7d\xa7\x14\xe4\xe9\xdc\x84\xbd\x9d\xbc\xf2\x6b\xe9\x51\x48\x8b\x15\x89\x14\xcd\x8e\x9c\xd3\x4a\x91\xc5\xea\xa5\xf7\x54\x1a\x43\x4e\x7c\x8e\xf7\x42\x96\x04\x74\x32\x45\x72\xf4\x46\x73\x4f\x58\x13\xb3\xac\x28\x99\x62\x7f\xc8\x45\xd9\x7d\x3c\x8d\xb1\x0b\x54\x72\xb7\x93\xda\xc8\x95\x39\x4d\x8a\x6f\x7f\xda\x24\x8e\x22\x1c\xdd\x60\xac\x1e\x6b\x99\x6d\xe8\x85\xb3\x7c\xc6\x8d\xf3\xa3\xdb\x21\x5c\x6b\x09\xf0\x0f\x7a\x7f\x92\x73\x1e\x75\x79\xc1\x8d\x61\x22\xc7\x81\x67\x5c\x05\xa6\x7f\x91\xe0\xea\x63\xe0\xa5\xf0\x77\xca\x38\x3f\xe7\xfa\x3c\x8f\x3f\x01\x00\x00\xff\xff\x40\x8e\xce\x5c\xc2\x05\x00\x00")

func stdlib_cli_index_abs() ([]byte, error) {
	return bindata_read(
		_stdlib_cli_index_abs,
		"stdlib/cli/index.abs",
	)
}

var _stdlib_runtime_index_abs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x4a\x2d\x29\x2d\xca\x53\xa8\xe6\x52\x50\x50\x50\x50\xca\x4b\xcc\x4d\x55\xb2\x52\x50\x4a\x4c\x2a\x56\xd2\x81\x08\x95\xa5\x16\x15\x67\xe6\xe7\x29\x59\x29\x38\x3a\x05\xc7\x87\xb9\x06\x05\x7b\xfa\xfb\x71\xd5\x02\x02\x00\x00\xff\xff\x7c\x61\xb4\xcc\x38\x00\x00\x00")

func stdlib_runtime_index_abs() ([]byte, error) {
	return bindata_read(
		_stdlib_runtime_index_abs,
		"stdlib/runtime/index.abs",
	)
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
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
var _bindata = map[string]func() ([]byte, error){
	"stdlib/cli/index.abs": stdlib_cli_index_abs,
	"stdlib/runtime/index.abs": stdlib_runtime_index_abs,
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
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func func() ([]byte, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"stdlib": &_bintree_t{nil, map[string]*_bintree_t{
		"cli": &_bintree_t{nil, map[string]*_bintree_t{
			"index.abs": &_bintree_t{stdlib_cli_index_abs, map[string]*_bintree_t{
			}},
		}},
		"runtime": &_bintree_t{nil, map[string]*_bintree_t{
			"index.abs": &_bintree_t{stdlib_runtime_index_abs, map[string]*_bintree_t{
			}},
		}},
	}},
}}
