dev:
	docker compose --env-file="./.env" up

test:
	go test
