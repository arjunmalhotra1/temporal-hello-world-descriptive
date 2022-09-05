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

![Screen Shot 2022-09-05 at 4 29 19 PM](https://user-images.githubusercontent.com/43081882/188516291-23e8db5f-d330-4612-8623-283bb4f3e9bf.png)



## Run the starter
In a separate terminal window.
```
go run starter/main.go
```


![Screen Shot 2022-09-05 at 4 29 07 PM](https://user-images.githubusercontent.com/43081882/188516301-a1af47e0-9619-41ab-8d81-d8fe6d1daf7e.png)
