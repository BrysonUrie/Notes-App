# Notes App

This is a simple notes app that allows users to create, read, update, and delete notes. It is built using [Air](https://github.com/aofei/air), a lightweight web framework for Go.

## Features

- User authentication: Users can sign up, log in, and log out.
- Session management: User sessions are managed using cookies.
- Database integration: The app uses a database to store user information and notes.

## Prerequisites

Before running the app, make sure you have the following installed:

- Go: [Installation Guide](https://golang.org/doc/install)

## Installation

1. Clone the repository:

    ```shell
    git clone https://github.com/BrysonUrie/notes-app.git
    ```

2. Navigate to the project directory:

    ```shell
    cd notes-app
    ```

3. Install the dependencies:

    ```shell
    go mod download
    ```

4. Build and run the app:

    ```shell
    air
    ```

5. Open your browser and visit `http://localhost:8080` to access the app.

## Usage

- Sign up for a new account or log in with your existing credentials.
- Create, view, update, and delete notes.
- Log out when you're done.

## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvements, please open an issue or submit a pull request.

## License

This project is licensed under the [MIT License](LICENSE).