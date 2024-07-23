run:
	@echo "Start application..."
	docker-compose up --build

clean:
	@echo "Cleaning up"
	docker-compose down -v