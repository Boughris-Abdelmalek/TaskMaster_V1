run:
	docker-compose up --build postgres api

run-fe:
	cd ./ui && yarn dev