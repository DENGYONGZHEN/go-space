package tools

//Ignore the Warning in Your Editor:
//Since the file is only used to track tool dependencies (and will not be part of
//your regular build), you can safely ignore the warning if everything else works as expected.
// Many developers using this pattern have reported that despite the warning,
//code generation and builds work correctly

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2"
	_ "github.com/rakyll/statik"
	_ "google.golang.org/grpc/cmd/protoc-gen-go-grpc"
	_ "google.golang.org/protobuf/cmd/protoc-gen-go"
)
