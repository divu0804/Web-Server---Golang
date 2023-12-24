# Simple Go Web Server

This is a basic Go web server that handles different routes, including a form submission and a simple "hello" endpoint.

## Prerequisites

Before running this server, ensure you have Go installed on your system.

## Usage

1. Clone this repository to your local machine:
2. Navigate to the project directory:
3. Run the Go server:
4. Access the server in your web browser:
    Open a web browser and go to [http://localhost:8089/](http://localhost:8089/) to view the served content.

## Routes

- `/` : Renders the `index.html` template.
- `/form` : Handles a form submission.
- `/hello` : Displays a "Hello World!" message.

## File Structure

- `main.go` : Main Go source file containing server logic.
- `templates/` : Directory containing HTML templates.
  - `index.html` : HTML template for the root route.

