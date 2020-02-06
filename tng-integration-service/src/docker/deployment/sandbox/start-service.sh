docker-compose down && \
docker rmi vpos.asia:5000/h5-service && \
docker-compose pull && \
docker-compose up -d
#docker cp conf/app.conf h5-zalopay_h5-service_1:/conf/app.conf && \
#docker stop h5-zalopay_h5-service_1 && \
#docker start h5-zalopay_h5-service_1