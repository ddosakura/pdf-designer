.PHONY: build
build:
	go build -o ./build/pdf-designer

.PHONY: save
save:
	go run . save -s ./assets/template/origin -p ./local -t origin

.PHONY: init
init:
	go run . init -s ./wp -t ./local/origin.pdt

.PHONY: init-buildin
init-buildin:
	go run . init -s ./wp

.PHONY: work
work:
	go run . work -s ./wp
