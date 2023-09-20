build: ui-build go-build

ui-build:
	cd ui && npm run build

go-build:
	go build
