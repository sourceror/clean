# About

An implementation example of the CLEAN architecture.

    docker run --rm --user $(shell id -u):$(shell id -g) -v $(CURDIR):/defs namely/protoc-all:1.32_2 --go-source-relative -i proto/ -f books.proto -l go -o ./proto/stubs