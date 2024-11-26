# Portradis

**Portradis (PORtfolio TRAjano DIctionary Server)** is a simplified remake of Redis, developed in Go. This project aims to recreate some of the basic features of Redis to study and demonstrate skills in backend development, code structuring and the use of good practices.

## About the project

Redis is a key-value structure database that is widely used due to its high performance and versatility. Portradis was created as an educational and simplified implementation, allowing you to understand the fundamental concepts behind a database server similar to Redis.

## Features

- Functional integration with redis-cli.
- HTTP server for basic data storage operations.
- Handling of key-value structures.
- Modular and organized code for easy extension.

## Requirements

- Go (version 1.20+ recommended)
- Any Go-compatible operating system.

## Installation

1. Clone this repository:
  ```bash

   git clone https://github.com/everton-trajano/portradis.git
```
2. Navigate to the project directory:
  ```bash

  cd portradis
  ```
3. Install the dependencies:
  ```bash

  go mod tidy
  ```
4. To start the server, use:
  ```bash

  go run main.go
  ```

## Project structure

   * app/: Main application code.
   * commands/: Initialization and configuration scripts.
   * tools/: Auxiliary tools and utilities.

## Contribution

Feel free to open issues and pull requests. Feedback is always welcome!
