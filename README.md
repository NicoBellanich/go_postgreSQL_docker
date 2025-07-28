# go_postgreSQL_docker
Dockerized backend app. 

- Dockerized Backend API with Golang using Fiber framework
- Dockerized PostgreSQL
- Orchestrated by Docker-Compose


## Domain 

Trivia app. The only domain entity is `Fact`that holds a `Question`which is a string, and an `Answer`that is another string

## Based on

Based on this youtuve video : https://www.youtube.com/watch?v=p08c0-99SyU


## Local Execution

1. run `docker compose build`
2. run `docker compose run`
3. Now you can request over http://localhost:3000