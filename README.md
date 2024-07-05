# User Manager in Go with MongoDB

## Prerequisites (Tested with)
- Go 1.22 or later
- MongoDB 7.0 or later

## Installation
1. Clone the repository:

`git clone https://github.com/luckycreationsindia/go_user_manager.git`

2. Install dependencies:

`go install github.com/air-verse/air@latest`

`go mod download`

3. Create a .env file in the root directory and add variables from .env-example file.

## Usage
1. Start the MongoDB server.

2. Run the application:

make run

3. The application will start, and you can interact with the user management API.

## API Endpoints
- `POST /register`: Create a new user
- `POST /login`: Login a user
- `GET /profile`: Get a user profile after login
- `GET /admin-profile`: Get a user profile with admin permission middleware
- `GET /permission-profile`: Get a user profile with permission list containing 1

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.
