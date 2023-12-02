
run:
	cd frontend &&\
	npm run build &&\
	cd ../backend &&\
	go build -o ./bin ./cmd/ &&\
	./bin/cmd


start_frontend:
	cd frontend &&\
	npm run build


start_backend:
	cd backend &&\
	go build -o ./bin ./cmd/ &&\
	./bin/cmd
