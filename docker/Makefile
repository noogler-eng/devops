DATABASE_URL:=postgres://postgres:foobarbaz@localhost:5432/postgres

# make run-postgres	Starts a PostgreSQL database container.
# make run-api-node	Runs the Node.js API in development mode.
# make run-api-golang Runs the Golang API in development mode.
# make run-client-react	Starts the React frontend.

# each of this called make target

.PHONY: run-postgres
run-postgres:
	@echo Starting postgres container
	-docker run \
		-e POSTGRES_PASSWORD=foobarbaz \
		-v pgdata:/var/lib/postgresql/data \
		-p 5432:5432 \
		postgres:15.1-alpine

.PHONY: run-api-node
run-api-node:
	@echo Starting node api
	cd api-node && \
		DATABASE_URL=${DATABASE_URL} \
		npm run dev

.PHONY: run-api-golang
run-api-golang:
	@echo Starting golang api
	cd api-golang && \
		DATABASE_URL=${DATABASE_URL} \
		go run main.go

.PHONY: run-client-react
run-client-react:
	@echo Starting react client
	cd client-react && \
		npm run dev

.PHONY: say_hello
say_hello:
	@echo "Hello World"