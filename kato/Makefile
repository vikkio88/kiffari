watch-be:
	cd be && make watch
watch-fe:
	cd web && npm run dev
build-fe:
	cd web && npm run build
build-be:
	cd be && make build-awd

build-all: clean build-fe build-be
	cd web && mv dist/ ../be/bin/static
	cd be && tar -czvf bin.tar.gz bin/
	mkdir bin/
	mv be/bin.tar.gz ./bin/
clean:
	rm -rf bin be/bin be/tmp be/testing.db web/dist web/.env.production.bak bin.tar.gz