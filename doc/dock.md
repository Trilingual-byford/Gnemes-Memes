## Benefits
- Run container in seconds instead of minutes
- Uses less memory
- Does not need full Os
- Deployment
- Testing


 ## What is an image?:Image is a template for creating an environment of your Choice
 
 
 #### The commands i used on 20 Nov for practicing
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
### Docker Volumes Allows sharing of data.