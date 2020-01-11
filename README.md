# Go Challenge

A backend service code challenge using [Golang](https://golang.org/).

The main idea of the task is to build an API that read the payment transaction data from files (`JSON` format) and return them in the API response as `JSON` format. There are two payment providers `flypayA` and `flypayB`.

You can find the full description for the task [here](./task.md).

***NOTE***: This my ***first time*** code in Golang so there may be a lot to improve.

## Prerequisites:

The application was built using [Golang](https://golang.org/dl/) so you should have it installed on your machine to be able to run it.

Or you can run it through [Docker](https://www.docker.com/).

## Usage:
  
1. Download or clone this repository to your machine.
2. Build the Docker image:

    ```bash
    docker build -t go-challenge .
    ```

3. Create a container from the generated image:

    ```bash
    docker run --rm --name go-challenge -p 5000:5000 go-challenge
    ```

4. You can run the service directly using:

    ```bash
    go run main.go
    ```

5. Unit tests can be run through:

    ```bash
    go test -v
    ```

***NOTE***: A Postman Collection for the API endpoints can be found [here](https://www.getpostman.com/collections/3be97a194d06646a929e).

## License:
This software is licensed under the [MIT License](https://opensource.org/licenses/MIT).