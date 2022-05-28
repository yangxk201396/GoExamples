
### 1. 运行hello_server
    go run main.go server

### 2. 运行hello_client
    go run client.go

### 3. http调用server
    curl -X POST -k https://localhost:50052/hello_world -d '{"referer": "yxk"}'   
    curl -X POST -k https://localhost:50052/http/hello_world -H "Content-Type:application/json" -d '{"referer": "yxk"}'

### 4. 访问swagger.json
    https://127.0.0.1:50052/swagger/hello.swagger.json
    https://127.0.0.1:50052/openapiv2/hello.swagger.json