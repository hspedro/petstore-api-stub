// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.13.0 DO NOT EDIT.
package api

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	. "github.com/deepmap/oapi-codegen/examples/petstore-expanded/echo/api/models"
	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (GET /pets)
	FindPets(ctx echo.Context, params FindPetsParams) error

	// (POST /pets)
	AddPet(ctx echo.Context) error

	// (DELETE /pets/{id})
	DeletePet(ctx echo.Context, id int64) error

	// (GET /pets/{id})
	FindPetById(ctx echo.Context, id int64) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// FindPets converts echo context to params.
func (w *ServerInterfaceWrapper) FindPets(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params FindPetsParams
	// ------------- Optional query parameter "tags" -------------

	err = runtime.BindQueryParameter("form", true, false, "tags", ctx.QueryParams(), &params.Tags)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter tags: %s", err))
	}

	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", ctx.QueryParams(), &params.Limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter limit: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.FindPets(ctx, params)
	return err
}

// AddPet converts echo context to params.
func (w *ServerInterfaceWrapper) AddPet(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.AddPet(ctx)
	return err
}

// DeletePet converts echo context to params.
func (w *ServerInterfaceWrapper) DeletePet(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id int64

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.DeletePet(ctx, id)
	return err
}

// FindPetById converts echo context to params.
func (w *ServerInterfaceWrapper) FindPetById(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id int64

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.FindPetById(ctx, id)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/pets", wrapper.FindPets)
	router.POST(baseURL+"/pets", wrapper.AddPet)
	router.DELETE(baseURL+"/pets/:id", wrapper.DeletePet)
	router.GET(baseURL+"/pets/:id", wrapper.FindPetById)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
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
	var res = make(map[string]func() ([]byte, error))
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
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
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
