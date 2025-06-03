up:
	docker-compose up --build
down:
	docker-compose down
ps:
	docker-compose ps
logs:
	docker-compose logs -f
restart:
	docker-compose restart
db:
	docker-compose exec -it db psql -U postgres taskmanager
help:
	@echo "Available commands:"
	@echo "  up      - Start the containers"
	@echo "  down    - Stop the containers"
	@echo "  ps      - List the containers"
	@echo "  db      - Connect to the PostgreSQL database"
	@echo "  logs    - Show logs for the containers"
	@echo "  restart - Restart the containers"
	@echo "  help    - Show this help message"