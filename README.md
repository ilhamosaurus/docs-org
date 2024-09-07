## Description

Hai! in this project, i'm trying make web application to organize my personal document using Go, Templ, HTMX, Tailwindcss and DaisyUI. So yes it is a Fullstack app and it has a long way to complete yet i hope i can colaborate with you.

## Installation

1. Please check `.env.example` file for database connection and JWT secret then delete `.example` from the filename.
2. I recommend you to install air to user watchmode like nodemon by `go install github.com/air-verse/air@latest`
3. I'm using tailwindcss and daisyui so please install them with `npm install -D tailwindcss postcss autoprefixer daisyui`

```bash
# installing dependencies
$ go mod download

# build Go app
$ go build -o main.go
```

## Running the app

```bash
# watch tailwindcss
$ make tailwind-dev

# watch templ generate
$ make templ-dev

# serve the app
$ make dev
```

<!-- ## Deployment using docker

```bash
# build an image
$ docker-compose build

# running container
$ docker-compose up

# running container on background
$ docker-compose up -d
``` -->
