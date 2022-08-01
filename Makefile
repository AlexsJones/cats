V := $(shell cat pkg/VERSION)
all:
	cd pkg && docker buildx build --build-arg=VERSION=$V --platform linux/amd64,linux/arm64,linux/arm . -t cnskunkworks/cats:$V --push
deploy:
	helm upgrade --install cats .
undeploy:
	helm uninstall cats
