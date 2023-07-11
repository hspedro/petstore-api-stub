OUT_PATH=api
SERVER_LANG=go-echo-server
OPENAPI_SCHEMA="https://raw.githubusercontent.com/OAI/OpenAPI-Specification/main/examples/v3.0/petstore.yaml"

ifeq ($(DEBUG),1)
Q=
else
Q=@
endif

.PHONY: install
install:
	$(Q)go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest

.PHONY: gen
gen:
	$(Q)oapi-codegen \
		-package main \
		-generate server \
		petstore-expanded.yaml > impl.go
