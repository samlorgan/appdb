# Installs
Openapi
go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
Chi
go get -u github.com/go-chi/chi/v5


-generate flag. It defaults to types,client,server,spec
	
	types: generate all type definitions for all types in the OpenAPI spec. This will be everything under #components, as well as request parameter, request body, and response type objects.
    server: generate the Echo server boilerplate. server requires the types in the same package to compile.
    chi-server: generate the Chi server boilerplate. This code is dependent on that produced by the types target.
    fiber: generate the Fiber server boilerplate. This code is dependent on that produced by the types target.
    client: generate the client boilerplate. It, too, requires the types to be present in its package.
    spec: embed the OpenAPI spec into the generated code as a gzipped blob. This is then usable with the OapiRequestValidator, or to be used by other methods that need access to the parsed OpenAPI specification
    skip-fmt: skip running goimports on the generated code. This is useful for debugging the generated file in case the spec contains weird strings.
    skip-prune: skip pruning unused components from the spec prior to generating the code.
    import-mapping: specifies a map of references external OpenAPI specs to go Go include paths. Please see below.




oapi-codegen -package petfulldefault ./config/petstore-expanded.yaml > ./generated/fulldefault/petstore.gen.go

oapi-codegen -generate server -package petsplitserver ./config/petstore-expanded.yaml > ./generated/splitdefault/petstoreserver.gen.go
oapi-codegen -generate chi-server -package petsplitchiserver ./config/petstore-expanded.yaml > ./generated/splitdefault/petstorechiserver.gen.go
oapi-codegen -generate client -package petsplitclient ./config/petstore-expanded.yaml > ./generated/splitdefault/petstoreclient.gen.go
oapi-codegen -generate types -package petsplittypes ./config/petstore-expanded.yaml > ./generated/splitdefault/petstoretypes.gen.go
oapi-codegen -generate spec -package petsplitspec ./config/petstore-expanded.yaml > ./generated/splitdefault/petstorespec.gen.go