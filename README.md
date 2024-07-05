# Arcade_5

# GOJWT
GoJWT is a simple RESTful API built with Go and the Gin framework. It provides basic user authentication functionality, including signup, login, and token validation.

## Features
- User signup
- User login
- JWT token generation
- JWT token validation

## Tech Stack
- Go
- Gin
- PostgreSQL 
- Gorm (for database interactions)

## Getting Started
These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites
- Go (version 1.16 or later)
- PostgreSQL

## Installation
Clone the repository
```bash
  git clone https://github.com/yourusername/gopro.git
  cd gojwt
```

Install the dependencies:
```bash
go mod download
```

Set up your environment variables:
Create a `.env` file in the root directory of the project, and add the following variables:
```
PORT=3000
DB="host user=username password=password dbname=dbname sslmode=disable"
JWT="secrettoken"
```
Replace `yourusername`,`yourpassword`, and `yourdbname` with your PostgreSQL username, password, and database name.

Run the application:
```bash
air
```
API Endpoints
-  `POST /signup`: Register a new user. Body--> email & password
-  `POST /login`: Authenticate a user and return a JWT token. Body--> email & password
-  `GET /validate`: Validate a JWT token.


## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License
This project is licensed under the MIT License - see the LICENSE.md file for details.