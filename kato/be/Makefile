build:
	sh scripts/update_version.sh
	go build -o bin/ -tags prod
	sh scripts/restore_version.sh
run:
	go run .
watch:
	air .
clean:
	rm -rf *.db bin/
	sh scripts/restore_version.sh
build-awd:
	sh scripts/update_version.sh
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bin/ -tags prod
	sh scripts/restore_version.sh