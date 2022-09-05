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
```![Screen Shot 2022-09-05 at 4 29 19 PM](https://user-images.githubusercontent.com/43081882/188516253-af765804-90b5-49e9-b313-85c27252dff6.png)



## Run the starter
In a separate terminal window.
```
go run starter/main.go
```
![Screen Shot 2022-09-05 at 4 29 07 PM](https://user-images.githubusercontent.com/43081882/188516259-956567e1-4ff0-4100-8910-906202739605.png)
