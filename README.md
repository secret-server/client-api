# client-api
client-api to access secret server must be running.

- Built on a subset of the the secret server.
- To see what is supported, look at the swagger.yaml in this project.
- See the client-demo at https://github.com/secret-server/client-demo for examples of how to use the API.

- The following is how the project was built using go-swagger
```
// Validate the swagger file.  Clean up any issues
swagger validate ./swagger.yaml

swagger generate client  --principal jwt.MapClaims -f swagger.yaml

	For this generation to compile you need to have some packages in your go.mod:
		* github.com/go-openapi/errors
		* github.com/go-openapi/runtime
		* github.com/go-openapi/runtime/client
		* github.com/go-openapi/strfmt

go get -u ./...

go get github.com/go-openapi/errors
go get github.com/go-openapi/runtime
go get github.com/go-openapi/runtime/client
go get github.com/go-openapi/strfmt

go mod tidy

go build ./...
```
