db:
	docker compose up -d

ui:	
	cd web && npm run dev

start:
	go run main.go

.PHONY: db ui start