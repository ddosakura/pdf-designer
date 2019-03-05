.PHONY: build
build:
	go build -o ./build/pdf-designer

.PHONY: xbuild
xbuild:
	./build.sh

.PHONY: save-buildin
save-buildin:
	go run . save -s ./assets/template/origin -p ./local -t origin

.PHONY: save
save:
	go run . save -s ./wp -p ./local -t origin-wp

.PHONY: savet
savet:
	go run . save -s ./a-demo-template -p ./local -t ddosakura-template

.PHONY: pkg
pkg:
	go run . pkg -s ./wp -p ./local

.PHONY: init
init:
	go run . init -s ./wp -t ./local/origin-wp.pdt

.PHONY: init-buildin
init-buildin:
	go run . init -s ./wp

.PHONY: work
work:
	go run . work -s ./wp

.PHONY: workt
workt:
	go run . work -s ./a-demo-template

.PHONY: initt
initt:
	go run . init -s ./a-demo-template -t ./local/ddosakura-template.pdt
