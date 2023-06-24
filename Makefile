.PHONY: up
up:
	docker compose up -d

.PHONY: run
run:
	docker compose up

.PHONY: rerun
rerun:
	docker compose down
	docker compose up

.PHONY: reset
reset:
	docker compose down
	-docker image rm otp-getter-server
	docker compose up

.PHONY: bash
bash:
	docker compose exec -it --user root server bash 
