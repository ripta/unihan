build:
	go generate ./...
	goimports -w -l -local github.com/ripta/unihan ./pkg

download:
	mkdir -p tmp
	curl -fsSL -o tmp/Unihan_17.0.0.zip https://unicode.org/Public/17.0.0/ucd/Unihan.zip
