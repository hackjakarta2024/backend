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
    - **Description**: Get recommendation food.
    - **Response**:
      ```json
      {
        "message": string
        "data": {
          "user_id": "290fbc73-84f1-4a09-aa1f-1b09bbc2539e",
          "promo": {
              "id": "03990610-f220-4cf2-9e2e-e9b5f63f90ab",
              "name": "Promo Merdeka Diskon 50%"
          },
          "food": [
            {
                "id": "a8e9bce3-3444-4904-a7f5-a89b7d30735c",
                "name": "bakso telur bakso urat campur lontong",
                "restaurant_name": "Bakso Goyang Lidah Mas Alim - Gambir",
                "desc": "Nasi Goreng",
                "fake_price": 30000,
                "real_price": 25000,
                "image": "image url",
                "rating_total": 4,
                "user_review": [
                    {
                        "name": "Samantha Smith",
                        "review": "Bakso telur dan urat yang lezat, lontongnya juga pas. Kuahnya mantap!",
                        "rating": 4
                    },
                    {
                        "name": "Emily Davis",
                        "review": "Bakso telur bakso urat campur lontong ini lezat! Variasi baksonya enak semua.",
                        "rating": 5
                    }
                ]
            }
        }
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
