# Go FileHub

This project is a simple backend API for a Google Drive clone, built with Golang. It serves as a portfolio piece demonstrating best practices in Go API development, including a layered architecture, JWT authentication, testing, and Docker-based deployment.

---

## Features

-   **User Management**: Secure user registration and login.
-   **Authentication**: Protected routes using JWT (JSON Web Tokens).
-   **File Management**:
    -   Upload files (stores locally).
    -   List all personal files.
    -   Download files securely.
    -   Delete files.
-   **Database**: Uses PostgreSQL for data persistence.
-   **Deployment**: Fully containerized with Docker and Docker Compose.
-   **API Documentation**: Interactive Swagger/OpenAPI documentation.

---

## Tech Stack

-   **Language**: [Golang](https://golang.org/)
-   **Framework**: [Gin](https://github.com/gin-gonic/gin)
-   **Database**: [PostgreSQL](https://www.postgresql.org/)
-   **ORM**: [GORM](https://gorm.io/)
-   **Authentication**: [JWT](https://github.com/golang-jwt/jwt)
-   **Deployment**: [Docker](https://www.docker.com/) & [Docker Compose](https://docs.docker.com/compose/)
-   **API Docs**: [Swaggo](https://github.com/swaggo/swag)

---

## Getting Started

### Prerequisites

-   [Docker](https://www.docker.com/get-started) and [Docker Compose](https://docs.docker.com/compose/install/) must be installed.
-   [Go](https://go.dev/doc/install) (v1.21+) for running locally without Docker.

### Running the Application with Docker (Recommended)

1.  **Clone the repository:**
    ```bash
    git clone https://github.com/lskeey/go-filehub.git
    cd go-filehub
    ```

2.  **Setup Environment Variables:**
    Create a `.env` file by copying the example and update the values if needed.
    ```bash
    cp .env.example .env
    ```
    **Important**: In the `.env` file, make sure `DB_HOST` is set to `db` for Docker Compose networking.
    ```env
    DB_HOST=db
    ```

3.  **Build and Run with Docker Compose:**
    This command will build the Go application image, pull the PostgreSQL image, and start both containers.
    ```bash
    docker-compose up --build
    ```

The API will be available at `http://localhost:8080`.

---

## API Documentation

Once the application is running, you can access the interactive Swagger API documentation at:

[**http://localhost:8080/swagger/index.html**](http://localhost:8080/swagger/index.html)

This UI allows you to explore all endpoints, view models, and execute API calls directly from your browser.