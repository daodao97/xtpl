.PHONY: clean build admin run

APP_NAME = server
BUILD_DIR = $(PWD)/build

up:
	docker compose up -d

down:
	docker compose down

clean:
	rm -rf ./build
	rm -rf ./dist
	rm -rf ./web/dist

admin:
	cd adminui && pnpm i && pnpm build

admin_dev:
	pnpm --dir adminui run dev	

run: admin
	go run ./cmd --app-env dev --bind :4001

build:
	# build the backend
	go build -ldflags="-w -s" -o $(BUILD_DIR)/$(APP_NAME) main.go

typecheck:
	pnpm --dir web run typecheck