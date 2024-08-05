# Blogging Platform API

Welcome to the Blogging Platform API project! This project aims to create a simple RESTful API for a personal blog. With this API, you can perform basic CRUD (Create, Read, Update, Delete) operations on articles, allowing you to manage your blog content efficiently.

## Overview

This project involves creating an API that handles the following responsibilities:

1. **List Articles**: Return a list of articles, optionally filtered by parameters such as publishing date or tags.
2. **Get Single Article**: Retrieve a single article based on its ID.
3. **Create Article**: Add a new article to be published.
4. **Delete Article**: Remove a single article specified by its ID.
5. **Update Article**: Modify an existing article, identified by its ID.

## Technologies Used

This project utilizes the following skills and technologies:

- **Go**: The primary programming language used to develop the API.
- **gorilla/mux**: A powerful URL router and dispatcher for matching incoming requests to their respective handler.
- **PostgreSQL**: PostgreSQL will serve as the database management system for storing article data.
- **sqlx**: A library that provides a set of extensions on Go's standard database/sql library for handling SQL operations.
- **godotenv**: A library to manage environment variables.

## Getting Started

To get started with this project, follow these steps:

1. **Clone the Repository**:
    ```sh
    git clone https://github.com/yourusername/blogging-platform.git
    cd blogging-platform
    ```

2. **Set Up Environment**:
    - Install Go: [Download and install Go](https://golang.org/dl/).
    - Install PostgreSQL: [Download and install PostgreSQL](https://www.postgresql.org/download/).
    - Set up your environment variables in a `.env` file.
    ```plaintext
    DB_HOST=your_db_host
    DB_PORT=your_db_port
    DB_USER=your_db_user
    DB_PASSWORD=your_db_password
    DB_NAME=your_db_name
    ```

3. **Install Dependencies**:
    ```sh
    go mod tidy
    ```

4. **Run the Application**:
    ```sh
    go run cmd/main.go
    ```

5. **Test Your API**: Test your API endpoints using tools like Postman or Insomnia to ensure they work as expected.

## Example API Endpoints

- **GET /articles**: Retrieve a list of articles.
- **GET /articles/{id}**: Retrieve a specific article by its ID.
- **POST /articles**: Create a new article.
- **DELETE /articles/{id}**: Delete an article by its ID.
- **PUT /articles/{id}**: Update an existing article by its ID.

## Contribution

Contributions to this project are welcome! If you have any ideas for improvements or would like to contribute code, feel free to submit a pull request.
