# file for building testing and deploying the xApp 


export GO111MODULE=on



.PHONY: build


build:  		## build the xApp and generate the executable in /build/out 
	go build -o build/out/rmet-xapp ./main.go

run: 			## run the xApp
	./build/out/rmet-xapp


clean: 			## clean the build dir
	rm -fr build/*


help:           ## Help page for rmet-xapp Makefile.
	fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'
