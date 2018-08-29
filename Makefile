FUNCTION_NAME=palitra
BINPATH=bin
REGION=eu-west-1

all: clean format build zip

install:
	dep ensure -v
	go get -u github.com/aws/aws-lambda-go/cmd/build-lambda-zip

format:
	gofmt -w palitra

build:
	env GOOS=linux GOARCH=amd64 \
	go build -ldflags="-d -s -w" -a -o ${BINPATH}/${FUNCTION_NAME} ${FUNCTION_NAME}/*.go
	chmod +x -R ${BINPATH}

zip:
	build-lambda-zip -o ${FUNCTION_NAME}.zip bin/${FUNCTION_NAME}

clean:
	rm -rf ./bin ${FUNCTION_NAME}.zip

clean-all: clean
	rm -rf ./vendor

deploy:
	aws lambda create-function --function-name ${FUNCTION_NAME} --region ${REGION} --zip-file fileb://./${FUNCTION_NAME}.zip --runtime go1.x --tags media=image --role role-arn --handler Handler

test:
	go test ./... -v
