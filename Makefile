.PHONY: compose
compose:
	docker-compose up

.PHONY: compose-down
compose-down:
	docker-compose down --remove-orphans
