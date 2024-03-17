watch:
	sh scripts/run.sh
watch-be:
	cd be && make watch
watch-fe:
	cd web && npm run dev
build-fe:
	cd web && npm run build
build-fe-local:
	cd web && cp .env.production .env.production.bak1 && cp .env.local .env.production && npm run build && rm -rf ../be/static && mv dist/ ../be/static && mv .env.production.bak1 .env.production
build-be:
	cd be && make build-awd

build-all: clean build-fe build-be
	cd web && mv dist/ ../be/bin/static
	cd be && tar -czvf bin.tar.gz bin/
	mkdir bin/
	mv be/bin.tar.gz ./bin/
clean:
	rm -rf bin be/bin be/tmp web/dist web/.env.production.bak bin.tar.gz