# Arithmetic Operations API

The Arithmetic Operations API is a simple Go (Golang) application that provides an HTTP API for performing arithmetic calculations. It supports basic addition and subtraction operations and requires a "superuser" access token for authentication.

## Getting Started

### Prerequisites

- Go (Golang) must be installed. You can download it from the official website: https://golang.org/
- cURL (optional) or any other HTTP client to test the API.

### Installation

1. Clone the repository to your local machine:
``` bash
git clone https://github.com/orekhovskiy/soft
```
2. Change into the project directory:

``` bash
cd soft
```


3. Build and run the application:

``` go
go run main.go
```

The API will start running on `http://localhost:8080`.

## Usage

### Authentication

The API requires a "User-Access" header with the value "superuser" for authentication. Make sure to include this header in your requests.

### Endpoint

#### Calculate

- **URL:** `/calculate`
- **Method:** POST
- **Headers:**
    - `Content-Type: application/json`
    - `User-Access: superuser`
- **Request Body:**
    - `expression` (string): The arithmetic expression to calculate. Use `+` for addition and `-` for subtraction.

#### Example Request

``` bash
curl -X POST http://localhost:8080/calculate
-H "Content-Type: application/json"
-H "User-Access: superuser"
-d '{"expression": "2+2-3-5+1"}'
```
#### Example Response

```json
{
  "result": -3
}
```
