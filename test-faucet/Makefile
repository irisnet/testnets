BINARY_NAME=faucet
all: get_tools clean get_vendor_deps build

########################################
### Build

build:
ifeq ($(OS),Windows_NT)
	go build -o build/faucet.exe ./
else
	go build -o build/faucet ./
endif

clean:
	@rm -rf vendor/
	@rm -rf build/

get_vendor_deps:
	@echo "--> Running dep ensure"
	@dep ensure -v

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v


######################################
## Tools

check_tools:
	cd tools && $(MAKE) check_tools

get_tools:
	cd tools && $(MAKE) get_tools

update_tools:
	cd tools && $(MAKE) update_tools