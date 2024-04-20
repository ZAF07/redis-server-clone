## Redis Server Mock (Golang)

This project is a mock implementation of a Redis server written in Golang. It aims to replicate the behavior of a real Redis server for learning purposes.

### Why This Project?

Understanding the inner workings of Redis is crucial for developers who work with it extensively. This mock server provides a controlled environment to:

- Experiment with Redis commands and their functionalities.
- Test applications that interact with Redis without relying on a real server instance.
- Deepen my understanding of the Redis protocol (RESP).
- Understand and gain expertise in Golang's concurrency pattern

### Features

- Simulates core Redis commands (SET, GET, ECHO, etc. with TTL feature).
- In-memory data storage for testing purposes.
- TCP server implementation for client connections.

### Getting Started

**Prerequisites:**

- Golang (version 1.19 or later) installed on your system.

**Running the Mock Server:**

1. Clone this repository.
2. Open a terminal in the project directory.
3. Run `go run app/server.go` to start the mock server on the default Redis port (6379).

**Using the Mock Server:**

## Any TCP Tool (Netcat)

Open a new terminal and run:

```bash
 echo '*2\r\n$4\r\nping\r\n' | nc localhost 6379
```

## Redis-cli

You can use any Redis client tool to interact with the mock server on port 6379. For example, with the `redis-cli` tool:

```bash
redis-cli localhost 6379
SET key value
GET key
```
