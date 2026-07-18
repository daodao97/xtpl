.PHONY: install clean build admin admin_dev web web_dev run dev typecheck

APP_NAME = server
BUILD_DIR = $(PWD)/build

install:
	pnpm --dir adminui install --frozen-lockfile
	pnpm --dir web install --frozen-lockfile

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

web:
	pnpm --dir web install --frozen-lockfile
	pnpm --dir web run build

web_dev:
	pnpm --dir web run dev

run: admin web
	go run ./cmd --app-env dev --bind :4001

dev: web
	DEV_MODE=1 DEV_SERVER_URL=http://127.0.0.1:3333 go run ./cmd --app-env dev --bind :4001

build: admin web
	# build the backend
	mkdir -p $(BUILD_DIR)
	go build -ldflags="-w -s" -o $(BUILD_DIR)/$(APP_NAME) ./cmd

typecheck:
	pnpm --dir web run typecheck
