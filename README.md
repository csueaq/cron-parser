# cron-parser

## Clone this repo

## Run with Docker

### Install docker, see [here](https://www.docker.com/products/docker-desktop)
### then run the following commend

    docker build -t cron-parser .
    docker run -it cron-parser "*/15 0 1,15 * 1-5 /usr/bin/find"
    
    
## Run locally, assuming [Golang 1.16 or above](https://golang.org/dl/) is installed

    make build
    ./app "*/15 0 1,15 * 1-5 /usr/bin/find"

