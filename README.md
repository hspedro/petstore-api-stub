# petstore-api-stub

This project uses [oapi-codegen](https://github.com/deepmap/oapi-codegen)
to generate a useful stub server implementation of Pestore API based on its
OpenAPI YAML.

## Installing dependencies

```sh
make install
```

## Running the server

```sh
go mod tidy
go run .
```


## (Re)-Generating the server

```sh
make gen
```
