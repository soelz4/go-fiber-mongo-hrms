# go-fiber-mongo-hrms

⚠️ **NOTE!!!!!** This Goalng App Written for Me - So some Things Might not Work on Your PC or Laptop as this was Never Intended to be a Usable Full Fledged Application, in that Case, Please Try if You Can Fix that Up as Much as Possible, or you can Open an Issue for Help :) - You Need to Change Something Related to Database (in Makefile, docker-compose.yml and ...) - So Please Check Everything Before Running the Database and Server.

<img src="diagram0.png" width="850px">

## Requirements

- make
- [MongoDB](https://www.mongodb.com/)
- [MongoDB Docker Image](https://hub.docker.com/r/mongodb/mongodb-community-server)
- Create a Docker Network to Connect Golang Server Container and MongoDB Container with Each other

## Makefile

A Standard GNU Make File is Provided to Help with Running and Building Locally.

```text
help                 💬 This Help Message
lint                 🔎 Lint & Format, will not Fix but Sets Exit Code on Error
lint-fix             📜 Lint & Format, will Try to Fix Errors and Modify Code
build                🔨 Build Binary File
run                  >_ Run the Web Server Locally at PORT 9010
init                 📥 Download Dependencies From go.mod File
clean                🧹 Clean up Project
mongodb              🌱 Pull MongoDB Docker Image from Docker Hub Registry
docker-network       🪡 Create Docker Network
image                📦 Build Docker Container Image from Dockerfile
push                 📤 Push Container Image to Registry
compose-up           🧷 Create and Start Containers
compose-down         🧼 Stop and Remove Containers, Networks
```

Makefile Variables

| Makefile Variable | Default                                 |
| ----------------- | --------------------------------------- |
| SRC_DIR           | ./src/                                  |
| DEFAULT_GOAL      | help                                    |
| BINARY_NAME       | main                                    |
| BINARY_DIR        | ./bin/                                  |
| IMAGE_REPO        | soelz/go-fiber-mongo-hrms               |
| IMAGE_TAG         | 0.1                                     |
| MongoDB_IMAGE     | mongodb/mongodb-community-server:latest |

## DataBase (NoSQL MongoDB)

MongoDB Go Driver Configurations ([The Official Golang driver for MongoDB](https://github.com/mongodb/mongo-go-driver)) Available in [./src/db/db.go](src/db/db.go).

## Containers

RUN Server in a Container

```bash
docker run --name mongodb --network backend -p 27017:27017 -d mongodb/mongodb-community-server:latest
```

```bash
docker run --name gfmh --network backend -p 9010:9010 soelz/go-fiber-mongo-hrms:0.1
```

OR

Just Execute this Command ~>

```bash
make compose-up
```

<br>![cat](https://github-production-user-asset-6210df.s3.amazonaws.com/62666332/253642655-0438a9e1-d47d-4570-873c-5ddd59f46e9e.svg)</br>
