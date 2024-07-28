# GrabBites API Backend

## Table of Contents
- [Introduction](#introduction)
- [Features](#features)
- [Technologies](#technologies)
- [Architecture](#architecture)
- [Installation](#installation)
- [Usage](#usage)
- [API Endpoints](#api-endpoints)
- [Configuration](#configuration)
  
## Introduction

GrabBites API Backend is a robust and scalable backend service designed to integrate AI-generated data and serve it seamlessly to Android applications. This project leverages the power of Golang for efficient performance, PostgreSQL for reliable data storage, and BigQuery for saving and querying AI-generated data.

## Features

- **AI Data Integration**: Seamlessly integrate AI-generated data.
- **High Performance**: Built with Golang for speed and efficiency.
- **Scalable Data Storage**: Utilizes PostgreSQL for robust data management.
- **Advanced Analytics**: Incorporates BigQuery for saving and querying AI-generated data.
- **Secure Endpoints**: Provides secure and authenticated API endpoints.

## Technologies

- **Golang**: The primary programming language used for backend development.
- **PostgreSQL**: For reliable and scalable data storage.
- **BigQuery**: For saving and querying AI-generated data.
- **Zap**: For structured and high-performance logging.

## Architecture

![Architecture Diagram](path-to-your-architecture-diagram.png)

## Installation

1. **Clone the repository**:
    ```sh
    git clone https://github.com/hackjakarta2024/backend
    cd backend
    ```

2. **Install dependencies**:
    ```sh
    go mod download
    ```

3. **Set up your environment variables** (see [Configuration](#configuration)).

## Usage

1. **Run the server**:
    ```sh
    go run main.go
    ```

2. **Access the API**:
    The API will be available at `http://localhost:3000`.

## API Endpoints

### Authentication

- **POST /api/v1/login**
    - **Description**: Authenticate user and return a token.
    - **Request**: 
      ```json
      {
        "email": "string",
        "password": "string"
      }
      ```
    - **Response**:
      ```json
      {
        "message: "string"
        "token": "string"
      }
      ```

### Data Endpoint

- **GET /api/v1/data**
    - **Description**: Send recommendation .
    - **Response**:
      ```json
      {
        "data": [
          {
            "id": "string",
            "value": "string"
          }
        ]
      }
      ```

## Configuration

Configure the necessary environment variables:

```sh
export DATABASE_HOST=your-db-host
export DATABASE_USER=your-db-user
export DATABASE_PASSWORD=your-db-password
export DATABASE_NAME=your-db-name
export DATABASE_PORT=your-db-port
export JWT_SECRETKEY=your-jwt-secretkey
export PROJECT_ID=your-project-id
