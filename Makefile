clean:
	rm -rf client models

generate: clean
	docker run --rm -it -e GOPATH=$(HOME)/go:/go \
		-v $(HOME):$(HOME) \
		-w $(shell pwd) \
		quay.io/goswagger/swagger \
		generate client -f ./swagger.yaml --skip-validation

download-spec:
	curl -sSLfo swagger.yaml https://api-reference.pagerduty.com/output.yaml

update-spec: download-spec patch-spec

patch-spec:
	git apply -v on_call_details.patch