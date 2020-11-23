## Benefits
- Run container in seconds instead of minutes
- Uses less memory
- Does not need full Os
- Deployment
- Testing


 ## What is an image?:Image is a template for creating an environment of your Choice
 
 
 #### The command line I used on 20 Nov for practicing
 - docker --help 
 - docker ps
 - docker pull nginx
 - docker images    //List all images you downloaded
 - docker run nginx:lastest
 - docker container ls
 - docker run -d nginx:latest  // -d --detached Run container in background and print container ID
 - docker stop 696f07e28c2c  
 - docker run -d -p 8080:80   // Publish a container's port(s) to the host
 - docker run -d -p 3000:80 -p 8080:80 nginx:latest //Run a docker image and publish it to two ports
 - docker ps -a //show all containers
 - docker rm 696f07e28c2c
 - docker build -t website:latest .
 - docker tag website:latest website:1 // Create a tag TARGET_IMAGE that refers to SOURCE_IMAGE
 ##### A tip for delete all docker container with one command
 ```console
  ❯ docker ps -aq
  6fa15761ebae
  5ab789c6b71d
  696f07e28c2c
  a556ca2a8b95
  d8965e36734e
  e0a32eb23192
  ❯ docker rm $(docker ps -aq)
  ❯ docker ps -aq
  6fa15761ebae
  5ab789c6b71d
  696f07e28c2c
  a556ca2a8b95
  d8965e36734e
  e0a32eb23192
  ❯ docker rm $(docker ps -aq)
  6fa15761ebae
  5ab789c6b71d
  696f07e28c2c
  a556ca2a8b95
  d8965e36734e
  e0a32eb23192
  ❯ docker ps -a
  CONTAINER ID   IMAGE     COMMAND   CREATED   STATUS    PORTS     NAMES
 ```
 - docker rm -f a556ca2a8b95 //Force the removal of a running container
 ```console
 ❯ docker run --name website -d -p 3000:80 nginx:latest //run a image and name it
 b2293ca8b691c37f25ffe585b6e6cbecb1b44dfa75d41ef2ce3d5ebafe58328f
```
- hosting some simple static content　
- https://hub.docker.com/_/nginx
```console
❯ docker run --name website -v $(pwd):/usr/share/nginx/html:ro -d -p 8080:80 nginx
1ff11515ab4cf36e2d55d828ef94122ab7a57d81985f27f0edc6deaf6078bec0
❯ ls
apidoc.html      apidoc.html.json dock.md
```
### Docker Volumes Allows sharing of data. Volume between host and container.
### Note:To reduce the size of image,pull image with alpine tag. 
- delete docker image example:
```console
❯ docker image rm 04790c585886
Error response from daemon: conflict: unable to delete 04790c585886 (must be forced) - image is being used by stopped container 3207687a8740
❯ docker image rm  -f 04790c585886
Untagged: gnemes-v1:latest
Deleted: sha256:04790c5858868882819f61097cccf9f6eea815a0ab008e634f9a4f516a1f3e91
```
### Docker Registries private/Public for publicising and manage image 
- https://hub.docker.com/repository/docker/niconicocsc/gnemes 

### Docker inspect
```console
 ∅ /  docker inspect 6fe6395c960c      
```
### Docker logs
```console
❯ docker logs 3207687a8740
[INFO] 2020/11/21 15:06 init Memes handle
[INFO] 2020/11/21 15:06 initial db connection
[ERRO] 2020/11/21 15:06 mongodb ping failedserver selection error: server selection timeout, current topology: { Type: Unknown, Servers: [{ Addr: localhost:27017, Type: Unknown, State: Connected, Average RTT: 0, Last error: connection() : dial tcp 127.0.0.1:27017: connect: connection refused }, ] }
Now listening on: http://localhost:8080
Application started. Press CTRL+C to shut down.
2020/11/21 15:06:56 
{
  "timestamp": "2020-11-21 15:07:26",
  "latency": 30000842300,
  "code": 400,
  "method": "GET",
  "path": "/api/v1/gnemes/memes",
  "ip": "172.17.0.1",
  "bytes_sent": -1
}
```
### Login in a Docker container
```console
❯ docker ps
CONTAINER ID   IMAGE           COMMAND                  CREATED         STATUS         PORTS                  NAMES
0f51ef50e088   apidoc:latest   "/docker-entrypoint.…"   2 minutes ago   Up 2 minutes   0.0.0.0:8080->80/tcp   mApidocs
❯ docker exec -it mApidocs /bin/bash
root@0f51ef50e088:/# ls
bin  boot  dev  docker-entrypoint.d  docker-entrypoint.sh  etc  home  lib  lib64  media  mnt  opt  proc  root  run  sbin  srv  sys  tmp  usr  var
```
### Docker network 