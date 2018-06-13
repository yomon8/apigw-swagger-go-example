STAGE	:=dev
STACK_NAME :=apigw-swagger-go-example-$(STAGE)
SWAGGER_FILE := swagger.$(STAGE).yml
S3_STAGING_BUCKET	:=your-bucket-name
DIR_BIN := ./bin
DIR_HANDLER :=./src/handler

build:
	rm -f ./bin/*
	GOOS=linux GOARCH=amd64 go build -o $(DIR_BIN)/root $(DIR_HANDLER)/root 
	GOOS=linux GOARCH=amd64 go build -o $(DIR_BIN)/hello $(DIR_HANDLER)/hello 

package: build
	sam package --s3-bucket $(S3_STAGING_BUCKET) \
		--template-file ./template.yml --output-template-file  ./.packaged.yml

deploy: package
	aws s3 cp ./src/api/$(SWAGGER_FILE) s3://$(S3_STAGING_BUCKET)/
	sam deploy --template-file ./.packaged.yml --stack-name $(STACK_NAME) --capabilities CAPABILITY_IAM \
	--parameter-overrides \
	Stage=$(STAGE) \
	ArtifactBucket=$(S3_STAGING_BUCKET) \
	SwaggerFile=$(SWAGGER_FILE)
	aws cloudformation describe-stacks --stack-name $(STACK_NAME) --query "Stacks[0].Outputs" --output table

local: build
	sam local start-api

.PHONY: build package deploy local