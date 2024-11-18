.PHONY: build-docker-client
build-docker-client:
	docker buildx build  --ssh default --load --file client.Dockerfile --progress plain -t "wow-client" .

.PHONY: build-docker-server
build-docker-server:
	docker buildx build  --ssh default --load --file server.Dockerfile --progress plain -t "wow-server" .