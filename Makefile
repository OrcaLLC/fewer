clean:
	rm -rf target/

build: clean build-osx build-linux

build-osx:
	GOOS=darwin go build -o target/osx/fewer
	zip target/fewer-osx.zip target/osx/fewer

build-linux:
	GOOS=linux go build -o target/linux/fewer
	zip target/fewer-linux.zip target/linux/fewer