# Credits to https://medium.com/google-cloud/deploy-go-application-to-kubernetes-in-30-seconds-ebff0f51d67b

.PHONY: install build serve clean pack deploy ship

TAG?=$(shell git rev-list HEAD --max-count=1 --abbrev-commit)

export TAG

OUTPUT_BINARY = shop_api

install:
	go get .

build: install
	go build -o ${OUTPUT_BINARY} -ldflags "-X main.version=$(TAG)" .

serve: build
	./${OUTPUT_BINARY}

clean:
	rm ./${OUTPUT_BINARY}

pack:
	GOOS=linux make build
	sudo docker build -t us.gcr.io/itsashopchallenge/shop_api:$(TAG) .

upload:
	gcloud docker -- us.gcr.io/itsashopchallenge/shop_api:$(TAG)

# deploy:
# 	envsubst < k8s/deployment.yml | kubectl apply -f -

ship: test pack upload deploy clean
