package generate

// For ogen options, see https://github.com/ogen-go/ogen/blob/main/cmd/ogen/main.go
//go:generate go run github.com/ogen-go/ogen/cmd/ogen@latest --target senzingrestapi --package senzingrestapi  --generate-tests --debug.noerr --clean senzing-rest-api.yaml
