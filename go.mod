module go.cbws.cloud/api

go 1.26.0

// The direct dependencies of the generated code, pinned so that generating the same protos twice
// gives the same go.mod and go.sum. go mod tidy adds the rest from these.
require (
	cloud.google.com/go/longrunning v0.6.4
	google.golang.org/genproto/googleapis/api v0.0.0-20261019180522-ce908d29299b
	google.golang.org/grpc v1.83.2
	google.golang.org/protobuf v1.36.12
)

require (
	golang.org/x/net v0.58.0 // indirect
	golang.org/x/sys v0.47.0 // indirect
	golang.org/x/text v0.41.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20261012220534-4341edf3f941 // indirect
)
