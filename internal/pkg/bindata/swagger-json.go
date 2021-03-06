// Code generated by go-bindata. DO NOT EDIT.
// sources:
// api/api.swagger.json
package bindata

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

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

var _apiApiSwaggerJson = []byte(`{
  "swagger": "2.0",
  "info": {
    "title": "cookie",
    "version": "version not set"
  },
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "paths": {
    "/pet/delete": {
      "delete": {
        "operationId": "PetStore_DeletePet",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/v1DeletePetResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "petID",
            "in": "query",
            "required": false,
            "type": "string"
          }
        ],
        "tags": [
          "PetStore"
        ]
      }
    },
    "/pet/get": {
      "get": {
        "operationId": "PetStore_GetPet",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/v1GetPetResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "pet_id",
            "in": "query",
            "required": false,
            "type": "string"
          }
        ],
        "tags": [
          "PetStore"
        ]
      }
    },
    "/pet/put": {
      "post": {
        "operationId": "PetStore_PutPet",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/v1PutPetResponse"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/v1PutPetRequest"
            }
          }
        ],
        "tags": [
          "PetStore"
        ]
      }
    }
  },
  "definitions": {
    "v1DeletePetResponse": {
      "type": "object"
    },
    "v1GetPetResponse": {
      "type": "object",
      "properties": {
        "pet": {
          "$ref": "#/definitions/v1Pet"
        }
      }
    },
    "v1Pet": {
      "type": "object",
      "properties": {
        "pet_type": {
          "$ref": "#/definitions/v1PetType"
        },
        "pet_id": {
          "type": "string"
        },
        "name": {
          "type": "string"
        }
      },
      "description": "Pet represents a pet in the pet store."
    },
    "v1PetType": {
      "type": "string",
      "enum": [
        "PET_TYPE_UNSPECIFIED",
        "PET_TYPE_CAT",
        "PET_TYPE_DOG",
        "PET_TYPE_SNAKE",
        "PET_TYPE_HAMSTER"
      ],
      "default": "PET_TYPE_UNSPECIFIED",
      "description": "PetType represents the different types of pets in the pet store."
    },
    "v1PutPetRequest": {
      "type": "object",
      "properties": {
        "pet_type": {
          "$ref": "#/definitions/v1PetType"
        },
        "id": {
          "type": "string"
        },
        "name": {
          "type": "string"
        }
      }
    },
    "v1PutPetResponse": {
      "type": "object",
      "properties": {
        "pet": {
          "$ref": "#/definitions/v1Pet"
        }
      }
    }
  }
}
`)

func apiApiSwaggerJsonBytes() ([]byte, error) {
	return _apiApiSwaggerJson, nil
}

func apiApiSwaggerJson() (*asset, error) {
	bytes, err := apiApiSwaggerJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "api/api.swagger.json", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"api/api.swagger.json": apiApiSwaggerJson,
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
	"api": &bintree{nil, map[string]*bintree{
		"api.swagger.json": &bintree{apiApiSwaggerJson, map[string]*bintree{}},
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
