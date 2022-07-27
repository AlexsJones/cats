all:
	cd pkg && docker buildx build --platform linux/amd64,linux/arm64,linux/arm . -t cnskunkworks/cats:bad-cat --push
deploy:
	helm upgrade --install cats .
undeploy:
	helm uninstall cats
