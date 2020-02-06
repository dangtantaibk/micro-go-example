* Yêu cầu: server đã cài sẵn docker, docker-compose.
* Copy thư mục sandbox lên server sau đó cd vào thư mục sandbox, chạy command

```
./start-service.sh
```

* Ghi chú: 

```
- Sandbox: 
(*) File app.conf: để cấu hình RunMode = "sandbox"
```

```
- Prodution:
(*) File app.conf: để cấu hình RunMode = "prod"
```

```
Khi docker cp phải sửa lại tên file cho phù hợp với môi trường sandbox hay prod
```
