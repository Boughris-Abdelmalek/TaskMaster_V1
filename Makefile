run:
	make -j 2 stop-port run-build

run-build:
	docker-compose up --build postgres api

stop:
	docker-compose down -v

run-fe:
	cd ./ui && yarn dev

stop-port:
	kill $( lsof -i :5432 -t )
