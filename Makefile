#
# swapp / Makefile
#

#
# VARS
#

-include .env

PROJECT_NAME?=swapp

DOCKER_DEV_IMAGE?=${PROJECT_NAME}-image
DOCKER_DEV_CONTAINER?=${PROJECT_NAME}-server

# define standard colors
# https://gist.github.com/rsperl/d2dfe88a520968fbc1f49db0a29345b9
ifneq (,$(findstring xterm,${TERM}))
	BLACK        := $(shell tput -Txterm setaf 0)
	RED          := $(shell tput -Txterm setaf 1)
	GREEN        := $(shell tput -Txterm setaf 2)
	YELLOW       := $(shell tput -Txterm setaf 3)
	LIGHTPURPLE  := $(shell tput -Txterm setaf 4)
	PURPLE       := $(shell tput -Txterm setaf 5)
	BLUE         := $(shell tput -Txterm setaf 6)
	WHITE        := $(shell tput -Txterm setaf 7)
	RESET        := $(shell tput -Txterm sgr0)
else
	BLACK        := ""
	RED          := ""
	GREEN        := ""
	YELLOW       := ""
	LIGHTPURPLE  := ""
	PURPLE       := ""
	BLUE         := ""
	WHITE        := ""
	RESET        := ""
endif

export


#
# TARGETS
#

.PHONY: all info build go src make doc

all: info

info: 
	@echo -e "\n${GREEN} ${PROJECT_NAME} / Makefile ${RESET}\n"

	@echo -e "${YELLOW} make config  --- check dev environment ${RESET}"
	@echo -e "${YELLOW} make fmt     --- reformat the go source (gofmt) ${RESET}"
	@echo -e "${YELLOW} make doc     --- render documentation from code (go doc) ${RESET}\n"

	@echo -e "${YELLOW} make build   --- build project (docker image) ${RESET}"
	@echo -e "${YELLOW} make run     --- run project ${RESET}"
	@echo -e "${YELLOW} make logs    --- fetch container's logs ${RESET}"
	@echo -e "${YELLOW} make stop    --- stop and purge project (only docker containers!) ${RESET}"
	@echo -e ""

#
# swapp wasm server binary
#

#build:
#	GOARCH=wasm GOOS=js go build -o web/app.wasm
#	go build
#
#run: build
#	./hello





config:
	@exit 0

fmt:
	@echo -e "\n${YELLOW} Code reformating (gofmt)... ${RESET}\n"
	@gofmt -d .
	@find . -name "*.go" -exec gofmt {} \;

build: 
	@echo -e "\n${YELLOW} Building the project (docker-compose build)... ${RESET}\n"
	@docker compose build --no-cache

run:	build
	@echo -e "\n${YELLOW} Starting project (docker-compose up)... ${RESET}\n"
	@docker compose up --force-recreate --detach --remove-orphans

logs:
	@echo -e "\n${YELLOW} Fetching container's logs (CTRL-C to exit)... ${RESET}\n"
	@docker logs ${DOCKER_DEV_CONTAINER} -f

stop:  
	@echo -e "\n${YELLOW} Stopping and purging project (docker-compose down)... ${RESET}\n"
	@docker compose down

