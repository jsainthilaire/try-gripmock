## Simple example using [GripMock](https://github.com/tokopedia/gripmock)

### How to use
    - Go to the makefile and set the path to the project.
    - Run make generate-proto.
    - Run make run-with-mock.
    - When you see the message "Serving gRPC on tcp://:4770" run go run main.go

#### Notes
 - The server.go file shows an implementation of the server.
 - This shows how to mock a simple endpoint, we can use matchers and other params to decide the returned values.
