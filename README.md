# Generated Go gRPC libraries for CBWS

This repository contains the Go module `go.cbws.cloud/api`: the messages and the gRPC clients of the public CBWS APIs. A process generates it from the definitions in [cbws/protobufs](https://github.com/cbws/protobufs). Do not change the generated files by hand.

## Installation

To add the module to your project, run this command:

```
go get go.cbws.cloud/api
```

Each API version is a package, for example:

```go
import vpcpb "go.cbws.cloud/api/vpc/v1alpha1"
```

## Versions

The tags of this repository have the same version numbers as the tags of [cbws/protobufs](https://github.com/cbws/protobufs). Each commit names the commit of cbws/protobufs it was generated from.

## License

The module is available under the Apache License, Version 2.0. Refer to the `LICENSE` file.
