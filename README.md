# Pack Calculator

This repository contains a full-stack application that calculates the optimal way to fulfill an item order using predefined pack sizes. It is built in Go with a lightweight static HTML frontend and containerized using Docker.

## Getting Started

### Prerequisites

To run this project, you need:

* **Docker**
* **Docker Compose**
* (Optional) **Go 1.21+** if running outside containers

### Running the Application

To build and start the application using Docker:

```bash
make start
```

This will serve the app locally on [http://localhost:3000](http://localhost:3000).


To stop the application, run:

```bash
make stop
```

### Running Tests

To run all Go unit tests:

```bash
make test
```

## Technologies Used

* **Golang**: API and core pack logic
* **Fiber**: Fast Go HTTP framework
* **Zerolog**: Structured logging
* **HTML + CSS + JS**: Static UI
* **Docker + Compose**: Containerized app
* **Heroku**: Deployed hosting
* **Gomock + Testify**: Unit testing and mocking

## Project Structure

```
order-packs-calculator/
├── main.go
├── Dockerfile
├── docker-compose.yml
├── static/                  # HTML/CSS/JS frontend
├── internal/
│   ├── handler/             # HTTP handlers
│   ├── service/             # Business logic
│   ├── repository/          # In-memory store
│   ├── router/              # Route wiring
│   └── server/              # App startup
├── mocks/                   # Auto-generated mocks
└── Makefile
```

## Makefile Targets

| Target          | Description                  |
| --------------- | ---------------------------- |
| `make start`    | Run the app in Docker        |
| `make test`     | Run all Go unit tests        |
| `make stop`     | Stop the Docker containers   |

## API Endpoints

### Calculate Packs

* **URL**: `POST /api/calculate`
* **Request**: `{ "items": 263 }`
* **Response**:

```json
{
  "total": 300,
  "packs": [
    { "size": 250, "quantity": 1 },
    { "size": 50, "quantity": 1 }
  ]
}
```

### Get Pack Sizes

* **URL**: `GET /api/packs`
* **Description**: Returns the current pack sizes

### Update Pack Sizes

* **URL**: `POST /api/packs`
* **Body**: `[250, 500, 1000]`
* **Description**: Replaces the configured pack sizes

## Testing Strategy

* Each layer is tested independently:

  * **Handler tests** use `httptest` + mocked services
  * **Service tests** mock repositories
* Mocks are generated with `mockgen` using `go:generate`

Example:

```go
//go:generate mockgen -source=pack_calculator.go -destination=mock/pack_calculator_mock.go -package=mocks
```

## Architecture

```
[Frontend UI] → [Handler Layer (Fiber)] → [Service Layer (Business Logic)] → [Repository (In-memory)]
```

* **Handlers** handle HTTP parsing and response
* **Services** encapsulate logic + validation
* **Repositories** provide state abstraction

## UI Interaction

Visit [http://localhost:3000](http://localhost:3000). The UI allows:

* Editing available pack sizes (top form)
* Submitting an item count for calculation (bottom form)
* Viewing results in a responsive table with:

  * Total items shipped
  * Pack breakdown (min total, min packs)

## Deployment

This app is deployed on Heroku using Docker:

Visit [https://order-packs-calculator-demo-ffee6b5535a0.herokuapp.com/](https://order-packs-calculator-demo-ffee6b5535a0.herokuapp.com/)

## Author

Antonio D'aria
