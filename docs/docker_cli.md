`docker build . -t ecom.pawtopia.vn` - Build Docker image
`docker network connect bridge redis-6381` - Connect redis container to Docker network
`docker network connect bridge mysql-container` - Connect mysql container to Docker network
`docker run --link mysql-container:mysql-container --link redis-6381:redis-6381 -p 8002:8000 ecom.pawtopia.vn` - run Docker image