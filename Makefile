lint:
	golangci-lint run --timeout=5m

lint-fix:
	gci write --skip-generated .
	golangci-lint run --timeout=5m --fix

build-docker:
	DOCKER_BUILDKIT=0 docker build --no-cache -t "jsonbeautifier:latest" .

run-docker:
	docker stop jsonbeautifier || true
	docker rm jsonbeautifier || true
	docker run -d -p 8080:8080 --name jsonbeautifier jsonbeautifier:latest

deploy:
	rm deploy.tar || true
	tar -cvf ./deploy.tar  ./*
	caprover deploy -t ./deploy.tar --host https://captain.app.tinygreencat.dev --caproverPassword ${CAPROVER_PASSWORD} --appName jsonbeautifier
	rm deploy.tar
