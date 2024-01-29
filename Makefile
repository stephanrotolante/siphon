run:
	go run *.go

# npm install -g nodemon@3.0.2
watch:
	nodemon

build:
	go build -o siphon.exe *.go
gen:
	cd build-tools && go run *.go -config ../siphon.config.yml