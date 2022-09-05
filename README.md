# temporal-hello-world-descriptive

## Bring the Temporal cluster up
  ```
  git clone https://github.com/temporalio/docker-compose.git
  cd docker-compose
  docker-compose up
  ```

## Run the worker
In a separate terminal window.
```
go run worker/main.go
```

## Run the starter
In a separate terminal window.
```
go run starter/main.go
```
