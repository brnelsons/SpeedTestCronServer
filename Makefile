.PHONY: test build run build-run

build:
	@echo '>> build'
	docker build -t brnelsons/speedtest-server .

run:
	@echo '>> run'
	docker run -p "8080:8080" brnelsons/speedtest-server