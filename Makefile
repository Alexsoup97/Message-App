
run:
	cd frontend &&\
	npm run build &&\
	cd ../backend &&\
	go build -o ./bin ./cmd/ &&\
	./bin/cmd


start_frontend:
	cd frontend &&\
	npm run build

# MAC/LINUX
# start_backend:
# 	cd backend &&\
# 	go build -o ./bin ./cmd/ &&\
# 	./bin/cmd.exe


# Windows
start_backend:
	cd backend &&\
	go build -o ./bin ./cmd/ &&\
	start ./bin/cmd.exe

