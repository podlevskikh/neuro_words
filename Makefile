PROJECT?=coursera
APP?=neuro_word
PORT?=80
PORT_APP?=7784

CONTAINER_IMAGE?=$(PROJECT)/${APP}
RELEASE?=0.0.1

clean:
	rm -f bin/${APP}

gorun: clean
	go build -o bin/${APP} -tags "dev load_envs" ./cmd/ && bin/${APP}

cover: test
	go tool cover -html=coverage.out
	