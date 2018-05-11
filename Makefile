clean:
	rm -rf client models

generate: clean
	docker run --rm -it -e GOPATH=$(HOME)/go:/go \
		-v $(HOME):$(HOME) \
		-w $(shell pwd) \
		quay.io/goswagger/swagger \
		generate client -f ./swagger.yaml --skip-validation

patch:
	git apply -v on_call_details.patch