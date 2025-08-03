# Go, Gin & GORM CRUD API with Docker & MySQL

This project is a complete, containerized CRUD (Create, Read, Update, Delete) API built with Golang, the Gin web framework, and the GORM library. It follows a clean architecture using a repository pattern to decouple the application logic from the database, enabling easy testing with mocks.

### Features

* RESTful API: Standard CRUD operations for a `books` resource.

* Go & Gin: Fast, efficient, and lightweight backend.

* GORM & Repository Pattern: Clean separation of concerns with a repository interface for database operations.

* MySQL: Popular and reliable open-source relational database.

* Dockerized: The application and database are fully containerized for easy setup and deployment.

* Unit Tests with Mocks: The test suite uses mocks (`testify/mock`) to test controller logic without needing a database connection, making tests fast and reliable.

* CI/CD Ready: Includes a GitHub Actions workflow for automated testing.

### Prerequisites

Before you begin, ensure you have the following installed on your local machine:

**Go (1.24+)**

**Docker**

**Docker Compose (usually included with Docker Desktop)**

**MySQL (optional)**: If you want to run the database locally without Docker, ensure MySQL is installed and running.

### How to Run the Application

There are two primary ways to run this application locally.

#### Method 1: Running Fully Containerized (Recommended for Quick Start)

This method runs both the API and the MySQL database inside Docker containers. It's the simplest way to get the application running without installing anything other than Docker.

1. Clone the Repository:

```
git clone <repository-url>
cd go-gin-gorm-crud
```

2. Build and Run with Docker Compose:<br/>
Open a terminal in the project's root directory and run the following command. This will build the Go application's Docker image and start both the API and database containers.

```
docker-compose up --build
```

(You can add the ```-d``` flag to run it in the background: ```docker-compose up --build -d```)

3. That's it! The API is now running and accessible at ```http://localhost:8080```.

#### Method 2: Running Locally for Development (Hybrid Mode)

This method is ideal for active development. You will run the database inside a Docker container, but you'll run the Go application directly on your machine. This allows you to make code changes and see them instantly without rebuilding a Docker image.

1. Start the Database Container Only:<br/>
In your terminal, run the following command to start just the MySQL service from your `docker-compose.yml` file.

```
docker-compose up -d db
```

2. Create the Local Configuration File:<br/>
Your Go application needs to know how to connect to the database.

* In the root of the project, create a new file named ```.env```.

* Copy and paste the following content into it. This tells your app to connect to the database running on ```localhost```.

```
# Local Development Database Configuration
DB_HOST=localhost
DB_PORT=3306
DB_USER=myuser
DB_PASSWORD=mypassword
DB_NAME=mydatabase
```

3. Run the Go Application:<br/>
Now, run the Go application directly from your terminal.

```
go run main.go
```

You will see logs in your terminal indicating the server has started. The API is now running and accessible at `http://localhost:8080`.

### Testing the API

Once the application is running (using either method), you can test the endpoints with a tool like Postman or `curl`.

#### Example: Create a new book
```
curl -X POST -H "Content-Type: application/json" \
-d '{"title":"The Pragmatic Programmer","author":"Andy Hunt"}' \
http://localhost:8080/api/v1/books
```
#### Example: Get all books

```
curl http://localhost:8080/api/v1/books
```

#### Running Unit Tests

The unit tests use mocks and do not require a database connection. You can run them at any time with the following command:

```
go test -v ./...
```
