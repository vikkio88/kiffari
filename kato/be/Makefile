build:
	sh scripts/update_version.sh
	go build -o bin/ -tags prod
run:
	go run .
watch:
	air .
clean:
	rm -rf *.db bin/
	sh scripts/restore_version.sh