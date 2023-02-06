generate-proto:
	docker run --volume "$(pwd):/workspace" --workdir /workspace bufbuild/buf generate


run-with-mock:
	docker run --rm --name grpc-mock -p 4770:4770 -p 4771:4771 -v /Users/yourusername/path/try-gripmock/proto:/proto -v /Users/yourusername/path/try-gripmock/stubs:/stub \
	tkpd/gripmock --stub=/stub /proto/hello.proto & \
