build: docker_arm64 docker_amd64

docker_arm64:
		docker run -it -v "${PWD}:/app" -v "${PWD}/../krakend/plugins:/krakend/plugins" -w /app --platform=linux/amd64 \
		-e "CGO_ENABLED=1" \
		-e "CC=aarch64-linux-musl-gcc" \
		-e "GOARCH=arm64" \
		-e "GOHOSTARCH=amd64" \
		krakend/builder:2.7.0 \
		go build -v -ldflags='-extldflags=-fuse-ld=bfd -extld=aarch64-linux-musl-gcc' \
		-buildmode=plugin -o /krakend/plugins/pheno_ui_maps_docker_arm64.so .

watch_docker_arm64:
		docker run -it -v "${PWD}:/app" -v "${PWD}/../krakend/plugins:/krakend/plugins" -w /app --platform=linux/amd64 \
		-e "CGO_ENABLED=1" \
		-e "GOHOSTARCH=amd64" \
		krakend/builder:2.7.0 \
		sh -c 'GOPATH=/usr/local/go go install github.com/mitranim/gow@latest && \
		CC=aarch64-linux-musl-gcc GOARCH=arm64 gow build -v -ldflags="-extldflags=-fuse-ld=bfd -extld=aarch64-linux-musl-gcc" \
		-buildmode=plugin -o /krakend/plugins/pheno_ui_maps_docker_arm64.so .'

docker_amd64:
	docker run -it -v "${PWD}:/app" -v "${PWD}/../krakend/plugins:/krakend/plugins" -w /app --platform=linux/amd64 \
  krakend/builder:2.7.0 \
  go build -v -buildmode=plugin -o /krakend/plugins/pheno_ui_maps_docker_amd64.so .

watch_docker_amd64:
	docker run -it -v "${PWD}:/app" -v "${PWD}/../krakend/plugins:/krakend/plugins" -w /app --platform=linux/amd64 \
  krakend/builder:2.7.0 \
  sh -c 'GOPATH=/usr/local/go go install github.com/mitranim/gow@latest && \
  gow build -v -buildmode=plugin -o /krakend/plugins/pheno_ui_maps_docker_amd64.so .'

.PHONY: build
