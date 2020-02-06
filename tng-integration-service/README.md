# H5 service 

## Cài đặt
```
> make prepare              # to install dep and some tools
> make dep                  # to get the dependencies
> make build                # to build the services
> make run                  # generate document and run
> make run_without_doc      # run
> make lint 				# to check coding convention
```

## Môi trường develop
### Jetbrains Goland (khuyến cáo nên dùng)
* Download and install: https://www.jetbrains.com/go/download/
* Config GOPATH and Format (gofmt)

* Configure File Watcher with arguments run --print-issued-lines=false $FileDir$.
* Predefined File Watcher will be added in issue.


### VSCode
* Download and install: https://code.visualstudio.com/
* Edit workspace or folder setting:
```
"settings": {
	"go.formatTool": "gofmt",
	"go.gopath": <YOUR_WORKING_PATH>,
	"go.lintTool":"golangci-lint",
	"go.lintFlags": [
		"--fast"
	]
}
```

## Quy định của project
1. Đối với controller khi đặt tên thì tên struct phải là chuỗi kết thúc bằng "Controller" -> Mục đích là quy định khi đặt tên module cho mã lỗi. Ví dụ:
```
type H5ZaloPayController struct {
	BaseController
	h5ZaloPayService services.H5ZaloPayService
}
-> Với trường hợp này thì module trong file conf/errors.yml là H5ZaloPay
```

2. Migration database.
- Nếu chưa có database thì phải tạo:
```
CREATE DATABASE colossus CHARACTER SET utf8 COLLATE utf8_general_ci;
```

- Chạy makefile để tiến hành migration databse
 ```
make migrate DB_USERNAME=root DB_PASSWORD=root DB_HOST=127.0.0.1 DB_PORT=3306 DB_DATABASE=colossus
```

3. Cấu hình docker để sử dụng docker push
- Trên MacOS: thêm config vào deamon (Preferences/Daemon/)
```
{
  "insecure-registries" : [
    "vpos.asia:5000"
  ]
}
```

- Trên Linux: tạo file daemon.json trong thư mục /etc/docker với nội dung
```
{
  "insecure-registries" : [
    "vpos.asia:5000"
  ]
}
```

- Để kiểm tra xem vpos.asia:5000 đã ready chưa thì có thể dán url http://vpos.asia:5000/v2/ vào trình duyệt để kiểm tra