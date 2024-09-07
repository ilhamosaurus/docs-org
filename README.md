# Personal Document Organizer

This is a fullstack web application designed to help you organize personal documents effectively. It is built using **Go**, **Templ** (for templating), **HTMX** (for modern interactive web features without JavaScript), **Tailwind CSS**, and **DaisyUI** (for styling). Although the project is still under development, it aims to provide a simple, responsive, and user-friendly interface to manage and store personal documents.

This app is in an early stage, and collaboration and ideas are always welcome!

## Features

- **Document management**: Easily organize, store, and retrieve personal documents.
- **Interactive UI**: Uses HTMX to deliver interactive UI components without the need for complex JavaScript.
- **Responsive Design**: Styled with Tailwind CSS and DaisyUI to ensure a smooth user experience on any device.
- **Templ-based Templating**: Templ, a Go templating engine, is used to generate fast and efficient HTML templates.
- **Backend in Go**: The backend is written in Go, leveraging its performance and simplicity for serving web applications.

## Installation

1. **Environment Setup**:

   - Rename `.env.example` to `.env`, and configure it with your database connection details and JWT secret.

   Example `.env` content:

   ```bash
   DATABASE_URL=postgres://user:password@localhost:5432/mydb
   JWT_SECRET=my-secret
   ```

2. **Install Go and Dependencies**:

   - Install the Go dependencies listed in `go.mod`:
     ```bash
     $ go mod download
     ```

3. **Install Air for Hot Reloading** (optional but recommended):

   - Use Air for live-reloading during development (similar to Nodemon in Node.js):
     ```bash
     $ go install github.com/air-verse/air@latest
     ```

4. **Install Tailwind CSS and DaisyUI**:

   - TailwindCSS and DaisyUI are used for the frontend UI. Install them by running:
     ```bash
     $ npm install -D tailwindcss postcss autoprefixer daisyui
     ```

5. **Build the Go Application**:
   - Once the dependencies are installed, build the Go app:
     ```bash
     $ go build -o main.go
     ```

## Running the App

There are several commands to run the application and its various development tools:

1. **Watch TailwindCSS for changes**:

   - Use this command to watch for changes in your CSS:
     ```bash
     $ make tailwind-dev
     ```

2. **Watch Templ for Template Generation**:

   - Automatically generate your HTML templates from Go:
     ```bash
     $ make templ-dev
     ```

3. **Serve the Application**:
   - Finally, to serve the app in development mode:
     ```bash
     $ make dev
     ```

## Development Tools

- **Go**: The core programming language for the backend.
- **Templ**: For rendering HTML templates in Go.
- **HTMX**: Enhances user experience by allowing dynamic updates to the UI without full page reloads.
- **Tailwind CSS**: A utility-first CSS framework for styling.
- **DaisyUI**: A TailwindCSS plugin that provides pre-built, customizable UI components.

<!-- ## Deployment using Docker (Optional)

You can also deploy this application using Docker. Here’s how you can build and run the app in containers:

### Build the Docker Image

```bash
$ docker-compose build
```

### Run the Container

```bash
$ docker-compose up
```

To run the container in the background:

```bash
$ docker-compose up -d
``` -->

## Contributing

Contributions are welcome! Feel free to open issues, submit pull requests, or provide feedback.
