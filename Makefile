build:
	go generate ./...
	goimports -w -l -local github.com/ripta/unihan ./pkg

download:
	mkdir -p tmp
	curl -fsSL -o tmp/Unihan_16.0.0.zip https://unicode.org/Public/16.0.0/ucd/Unihan.zip
