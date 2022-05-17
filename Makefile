all:
	cd pkg && docker buildx build --platform linux/amd64,linux/arm64,linux/arm . -t tibbar/cat-cozy:latest --push
deploy:
	helm upgrade --install cat-cozy .
