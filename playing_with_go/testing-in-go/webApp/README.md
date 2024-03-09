# WebApp

A web app and an API that allow uploading an image to a user's profile.

### Requirements

* Docker
  * Install Docker Desktop [here](https://docs.docker.com/get-docker/)

### Running

#### Web

* Start the Docker Desktop app locally
* In one terminal, run:
  ```bash
  docker-compose up
  ```
* In another terminal, run:
  ```bash
  go run ./cmd/web
  ```
* Navigate to `http://localhost:8080`
* Log in with the username `admin@example.com`, and the password `secret`
* You can now upload an image to your profile!

#### API

* Start the Docker Desktop app locally
* In one terminal, run:
  ```bash
  docker-compose up
  ```
* In another terminal, run:
  ```bash
  go run ./cmd/api
  ```
* In another terminal, run:
  ```bash
  curl http://localhost:8090/test
  ```
* You should receive something like this:
  ```
  {"message":"hello, world"}%
  ```
* Different API options:
  * Authenticate
    ```bash
    curl http://localhost:8090/auth -X POST -H "Content-Type: application/json" -d '{"email":"admin@example.com","password":"secret"}'
    ```

#### Tests

* Start the Docker Desktop app locally
* In a terminal, run:
  ```bash
  go test ./... -v
  ```

\*If you watch Docker Desktop as the tests run, you'll see the tests actually create a docker container to test the postgres connection, and then it cleans it up when the tests are done... cool!
