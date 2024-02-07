.PHONY: build clean deploy

LAMBDA_FUNCTIONS = level-1

build:
	make clean
	mkdir -p ./bin
	
	for func in ${LAMBDA_FUNCTIONS} ; do \
		GOOS=linux GOARCH=amd64 go build -tags lambda.norpc -buildmode pie \
		-o bin/bootstrap functions/$$func/*.go ; \
		cd bin && zip $$func.zip bootstrap && cd .. ; \
	done

clean:
	rm -rf ./bin

dry-deploy:
	make clean
	make build
	sls deploy --verbose --noDeploy

deploy: 
	make clean
	make build
	sls deploy --verbose
