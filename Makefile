build:
		GOOS=linux GOARCH=amd64 go build -C functions/ -tags lambda.norpc -o build/bootstrap main.go && \
		zip -jrm lambda.zip functions/build/bootstrap
deploy:
		make build
		cd ./infra && \
		npm i && \
		npx cdk deploy
