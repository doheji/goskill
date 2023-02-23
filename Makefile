postgres:
	docker run --name go_skill_db_container -e POSTGRES_DB=go_skill_db -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -p 1234:5432 -d postgres

migrateup:
	migrate -path ./db/migrations -database "postgresql://root:secret@localhost:1234/go_skill_db?sslmode=disable" up

migratedown:
	migrate -path ./db/migrations -database "postgresql://root:secret@localhost:1234/go_skill_db?sslmode=disable" down

.PHONY: postgres migrateup migratedown