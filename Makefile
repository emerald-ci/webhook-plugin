build:
	gb build all
run:
	echo \{\"job\":\{\"state\":\"passed\"\},\"config\":\{\"url\":\"https://morning-citadel-6326.herokuapp.com/\"\}\} | docker run -i -a stdin -a stdout emeraldci/webhook-plugin
docker:
	docker build -t emeraldci/webhook-plugin .
