default: build

deps:

	go get -t -v ./...

build:

	cd cmd/hackapp; \
	go build; \
	cd ../../cli; \
	go build
	#mkdir -p packaging/root/usr/local/bin
	#go build -o packaging/root/usr/local/bin/hackappd github.com/dbongo/hackapp/cmd/hackapp

install:

	install -t /usr/local/bin packaging/root/usr/local/bin/hackappd

run:

	cd cli; \
	hackapp server
	#@go run cmd/hackapp/main.go

clean:

	rm -f cmd/hackapp/hackapp; \
	rm -f cli/cli
	#rm -f packaging/root/usr/local/bin/hackappd

docker_run_db:

	docker run -d -P --name mongodb mongo

docker_stop_db:

	docker stop mongodb && docker rm mongodb

docker_build_hackapp:

	docker build -t dbongo/hackapp .

docker_run_hackapp:

	docker run -d -p 3000:3000 --link mongodb:mongodb --name hackapp dbongo/hackapp

docker_stop_hackapp:

	docker stop hackapp && docker rm hackapp
