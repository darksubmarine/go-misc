all:

PROJECT_ROOT := $(shell pwd)

testify:
	cd _codegen && go build -o _codegen main.go
	cd ..

	cd test/bdd && ../../_codegen/_codegen -output-package=bdd -template=assertion_testify.go.tmpl -out="$(PROJECT_ROOT)/test/bdd/assertion_testify.go"