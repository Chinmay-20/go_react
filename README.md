# TaskMaster

TaskMaster is a full-stack web application that integrates a Go backend with a React frontend. This project serves as a template for building robust and efficient web applications using Go and React.

## Features

- **Backend**: Developed with Go, offering high performance and scalability.
- **Frontend**: Built using React, providing a dynamic and responsive user interface.
- **API Communication**: Utilizes RESTful APIs for seamless interaction between frontend and backend.
- **Hot Reloading**: Configured for both frontend and backend to enhance the development experience.

## Prerequisites

Ensure you have the following installed:

- [Go](https://golang.org/doc/install)
- [Node.js and npm](https://nodejs.org/)

## Installation

1. **Clone the repository**:

   ```bash
   git clone https://github.com/Chinmay-20/go_react.git
   cd go_react
   ```

2. **Backend Setup**:

   - Navigate to the backend directory:

     ```bash
     cd backend
     ```

   - Install Go dependencies:

     ```bash
     go mod download
     ```

3. **Frontend Setup**:

   - Navigate to the frontend directory:

     ```bash
     cd ../client
     ```

   - Install npm dependencies:

     ```bash
     npm install
     ```

## Running the Application

1. **Start the Backend Server**:

   - In the `backend` directory, run:

     ```bash
     go run main.go
     ```

   - The backend server will start on `http://localhost:8080`.

2. **Start the Frontend Server**:

   - In the `client` directory, run:

     ```bash
     npm start
     ```

   - The React application will start on `http://localhost:3000`.

## Project Structure

```plaintext
go_react/
├── backend/        # Go backend source code
│   ├── main.go     # Entry point for the Go application
│   └── ...
├── client/         # React frontend source code
│   ├── src/
│   ├── public/
│   └── ...
├── screenshots/    # Screenshots of the application
│   ├── screenshot1.png
│   ├── screenshot2.png
│   ├── screenshot3.png
│   ├── screenshot4.png
├── README.md       # Project documentation
└── ...
```

## Output

### Initial View
![Initial View](screenshots/Screenshot%202025-02-04%20at%201.24.37%E2%80%AFPM.png)

### Adding Tasks
![Adding Tasks](screenshots/Screenshot%202025-02-04%20at%201.25.37%E2%80%AFPM.png)

### Marking Tasks as Done
![Marking Tasks as Done](screenshots/Screenshot%202025-02-04%20at%201.26.04%E2%80%AFPM.png)

### Removing Tasks
![Removing Tasks](screenshots/Screenshot%202025-02-04%20at%201.26.41%E2%80%AFPM.png)

## Contributing

Contributions are welcome! Please fork the repository and create a pull request with your changes.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

