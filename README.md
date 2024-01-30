# Case Study Fullstack Developer

## Logical Test

Go to logical_test folder

```cmd
cd logical_test
```

Run app

```go
go run main.go
```

Run unit testing

```go
cd grading/ && go test -v
```

## System Flow

Go to system flow folder

```cmd
cd system_flow
```

### Backend

Go to backend folder and install package

```go
go mod tidy
```

Copy .env and adjust with your environment, we are using <b>MySQL</b> database

```cmd
cp .env.example .env
```

Do migration

```go
go run migration/migration.go
```

Run app

```cmd
go run cmd/api/main.go
```

### Frontend

Go to frontend folder and install package

```cmd
npm i
```

Copy .env and adjust with your environment

```cmd
cp .env.example .env
```

Run App

```cmd
npm run dev
```
