VER=v1.0
REGISTRYHOST?=vpos.asia:5000

h5-service: NAME=h5-service
h5-service: TAG=tng/$(NAME)
h5-service: SOURCE_PATH=../tng/h5-integration-service
h5-service:
	mkdir -p swagger && cp -r ../tng/h5-integration-service/swagger/ swagger/
	mkdir -p conf && cp -r ../tng/h5-integration-service/conf/ conf/
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main $(SOURCE_PATH)/main.go
	docker build -t $(TAG) -t $(TAG):$(VER) -f dockerfile .
	rm -rf main swagger/ conf/
	docker tag $(TAG) $(REGISTRYHOST)/$(NAME)
	docker push $(REGISTRYHOST)/$(NAME)

dev-tool-service: NAME=dev-tool-service
dev-tool-service: TAG=tng/$(NAME)
dev-tool-service: SOURCE_PATH=../tng/dev-tool-service
dev-tool-service:
	mkdir -p swagger && cp -r ../tng/dev-tool-service/swagger/ swagger/
	mkdir -p conf && cp -r ../tng/dev-tool-service/conf/ conf/
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main $(SOURCE_PATH)/main.go
	docker build -t $(TAG) -t $(TAG):$(VER) -f dockerfile .
	rm -rf main swagger/ conf/
	docker tag $(TAG) $(REGISTRYHOST)/$(NAME)
	docker push $(REGISTRYHOST)/$(NAME)

pos-service: NAME=pos-service
pos-service: TAG=tng/$(NAME)
pos-service: SOURCE_PATH=../tng/pos-service
pos-service:
	mkdir -p swagger && cp -r ../tng/pos-service/swagger/ swagger/
	mkdir -p conf && cp -r ../tng/pos-service/conf/ conf/
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main $(SOURCE_PATH)/main.go
	docker build -t $(TAG) -t $(TAG):$(VER) -f dockerfile .
	rm -rf main swagger/ conf/
	docker tag $(TAG) $(REGISTRYHOST)/$(NAME)
	docker push $(REGISTRYHOST)/$(NAME)

promotion-service: NAME=promotion-service
promotion-service: TAG=tng/$(NAME)
promotion-service: SOURCE_PATH=../tng/promotion-service
promotion-service:
	mkdir -p swagger && cp -r ../tng/promotion-service/swagger/ swagger/
	mkdir -p conf && cp -r ../tng/promotion-service/conf/ conf/
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main $(SOURCE_PATH)/main.go
	docker build -t $(TAG) -t $(TAG):$(VER) -f dockerfile .
	rm -rf main swagger/ conf/
	docker tag $(TAG) $(REGISTRYHOST)/$(NAME)
	docker push $(REGISTRYHOST)/$(NAME)

cron-job-service: NAME=cron-job-service
cron-job-service: TAG=tng/$(NAME)
cron-job-service: SOURCE_PATH=../tng/cron-job-service
cron-job-service:
	mkdir -p swagger && cp -r ../tng/cron-job-service/swagger/ swagger/
	mkdir -p conf && cp -r ../tng/cron-job-service/conf/ conf/
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main $(SOURCE_PATH)/main.go
	docker build -t $(TAG) -t $(TAG):$(VER) -f dockerfile .
	rm -rf main swagger/ conf/
	docker tag $(TAG) $(REGISTRYHOST)/$(NAME)
	docker push $(REGISTRYHOST)/$(NAME)

shipper-service: NAME=shipper-service
shipper-service: TAG=tng/$(NAME)
shipper-service: SOURCE_PATH=../tng/shipper-service
shipper-service:
	mkdir -p swagger && cp -r ../tng/shipper-service/swagger/ swagger/
	mkdir -p conf && cp -r ../tng/shipper-service/conf/ conf/
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main $(SOURCE_PATH)/main.go
	docker build -t $(TAG) -t $(TAG):$(VER) -f dockerfile .
	rm -rf main swagger/ conf/
	docker tag $(TAG) $(REGISTRYHOST)/$(NAME)
	docker push $(REGISTRYHOST)/$(NAME)

loyalty-service: NAME=loyalty-service
loyalty-service: TAG=tng/$(NAME)
loyalty-service: SOURCE_PATH=../tng/loyalty-service
loyalty-service:
	mkdir -p swagger && cp -r ../tng/loyalty-service/swagger/ swagger/
	mkdir -p conf && cp -r ../tng/loyalty-service/conf/ conf/
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main $(SOURCE_PATH)/main.go
	docker build -t $(TAG) -t $(TAG):$(VER) -f dockerfile .
	rm -rf main swagger/ conf/
	docker tag $(TAG) $(REGISTRYHOST)/$(NAME)
	docker push $(REGISTRYHOST)/$(NAME)

menu-service: NAME=menu-service
menu-service: TAG=tng/$(NAME)
menu-service: SOURCE_PATH=../tng/menu-service
menu-service:
	mkdir -p swagger && cp -r ../tng/menu-service/swagger/ swagger/
	mkdir -p conf && cp -r ../tng/menu-service/conf/ conf/
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main $(SOURCE_PATH)/main.go
	docker build -t $(TAG) -t $(TAG):$(VER) -f dockerfile .
	rm -rf main swagger/ conf/
	docker tag $(TAG) $(REGISTRYHOST)/$(NAME)
	docker push $(REGISTRYHOST)/$(NAME)

user-profile-service: NAME=user-profile-service
user-profile-service: TAG=tng/$(NAME)
user-profile-service: SOURCE_PATH=../tng/user-profile-service
user-profile-service:
	mkdir -p swagger && cp -r ../tng/user-profile-service/swagger/ swagger/
	mkdir -p conf && cp -r ../tng/user-profile-service/conf/ conf/
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main $(SOURCE_PATH)/main.go
	docker build -t $(TAG) -t $(TAG):$(VER) -f dockerfile .
	rm -rf main swagger/ conf/
	docker tag $(TAG) $(REGISTRYHOST)/$(NAME)
	docker push $(REGISTRYHOST)/$(NAME)
