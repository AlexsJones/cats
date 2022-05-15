all:
	cd pkg && docker build . -t tibbar/cat-cozy:latest && docker push tibbar/cat-cozy:latest
deploy:
	helm upgrade --install cat-cozy .
