dev:
	docker compose --profile database --env-file="./.env" up

test:
	go test
