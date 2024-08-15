.PHONY: test build dist dist-docker


dist-docker:
	docker build -t docker.io/brnelsons/speedtest-server .