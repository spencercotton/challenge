### Instructions

* Run ```make start```
* Load a gRPC client (such as bloomRPC)
* Load the `service.proto` found in the root of the directory
* Ensure the endpoint is `localhost:50051`

### Tests
* Run `go test ./...`
* I wanted to provide more test coverage but was unable to due to time constraints. I have added some tests for the repository layer.