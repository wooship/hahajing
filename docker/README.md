build image yourself with Dockerfile:
  docker build . -t hahajing
  docker run -itd -p 80:80  hahajing

download from docker hub and run with docker-compose:
  docker-compose up -d
