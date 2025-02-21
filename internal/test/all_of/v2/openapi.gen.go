// Package v2 provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.0.0-00010101000000-000000000000 DO NOT EDIT.
package v2

import (
	"bytes"
	"compress/flate"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"

	"github.com/getkin/kin-openapi/openapi3"
)

// Person defines model for Person.
type Person struct {
	FirstName          string `json:"FirstName"`
	GovernmentIDNumber *int64 `json:"GovernmentIDNumber,omitempty"`
	LastName           string `json:"LastName"`
}

// PersonProperties These are fields that specify a person. They are all optional, and
// would be used by an `Edit` style API endpoint, where each is optional.
type PersonProperties struct {
	FirstName          *string `json:"FirstName,omitempty"`
	GovernmentIDNumber *int64  `json:"GovernmentIDNumber,omitempty"`
	LastName           *string `json:"LastName,omitempty"`
}

// PersonWithID defines model for PersonWithID.
type PersonWithID struct {
	FirstName          string `json:"FirstName"`
	GovernmentIDNumber *int64 `json:"GovernmentIDNumber,omitempty"`
	ID                 int64  `json:"ID"`
	LastName           string `json:"LastName"`
}

// Base64 encoded, compressed with deflate, json marshaled Swagger object
const swaggerSpec = "" +
	"lJRPb9s4EMW/yoC7R0FOsIs96Bas20BAkRpo0hziABlLI4spNWTJUQzB8HcvSMn/4BZpfRrA5OO8N7/R" +
	"VlW2c5aJJahiq0LVUoepXJAPlmOFxnxuVPG0VX97alSh/podb82mK7Px/MJbR140BbXLtsrT9157qlXx" +
	"pD5qH+QOO1KZ+oRT+bx7zlRNofLaiY7vqftWB9ABEFySzGCjpYUOuUaxfoAmCgFyDQaDAGNHGax6AZsk" +
	"0EA5XzL33Yp8DkluY3tTw4rAk/SeqYbVAAgvtyQvEGQwBDeLModHgo78mkBamp5fsjt4GjtBttKShy/J" +
	"OWxaXbVg2QzgvH3TNQXY+4ZGk6lDvmSVKRkcqULZ1StVonaZuois2F5kQYEAPU1CIC0KBEeVboZDQtEk" +
	"DekYGnOIIYsZLfngvQ+Tb4aXD7U+dQ7EtbOaJYNNS56AsGrjEPZaowN31upxoMV2by6I17yO5m7tG3nu" +
	"iKWc36VZxGON9R2KKpRm+e/fYyiahdbk48UDG5equ1+G+KilLed/Smti9NzUKPJum7vsjO1y/jskg6fK" +
	"+howHDlsvO0A4X9PKHQYQw6lQGVZUHNYcpxqJHKCwDaAsDhdDmTAutYT/p6C7X1F8PBQzn/KXuxfc2NT" +
	"xlpM/O+eggS4ifFBSiwkPZWpN/JhdHSdX+VXMXXriNFpVah/8qv8OrKB0qYEZ85gRa019TjyNckl2F/R" +
	"6LTOATbIAihgKG6zZYIolUGwIDHADr9RBJ86aNG5YTQUZ4ZRrKxVoRYnT8bJBGc57Beqwd6kFmKgxKlE" +
	"54yuksDsdfrOjWzE6n1yJt5SkOfOTt3v0u9HAAAA//8="

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	compressed, err := base64.StdEncoding.DecodeString(swaggerSpec)
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr := flate.NewReader(bytes.NewReader(compressed))
	var buf bytes.Buffer
	if _, err := buf.ReadFrom(zr); err != nil {
		return nil, fmt.Errorf("read flate: %w", err)
	}
	if err := zr.Close(); err != nil {
		return nil, fmt.Errorf("close flate reader: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
