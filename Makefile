all:
	cd pkg && docker buildx build --platform linux/amd64,linux/arm64,linux/arm . -t tibbar/cat-cozy:latest && docker push tibbar/cat-cozy:latest
deploy:
	helm upgrade --install cat-cozy .
