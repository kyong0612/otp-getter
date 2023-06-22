.PHONY: up
up:
	docker compose up -d

.PHONY: reset
reset:
	docker compose down
	-docker image rm otp-getter-server
	docker compose up -d

.PHONY: bash
bash:
	docker compose exec -it --user root server bash 

