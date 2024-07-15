run:
	go run main.go
	
run_db_migrate_up:
	go run database/migrate/up/up.go

run_db_migrate_down:
	go run database/migrate/down/down.go

run_db_seed_user:
	go run database/seed/user/user.go