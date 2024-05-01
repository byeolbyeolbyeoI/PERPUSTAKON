# PERPUSTAKON 
Golang API that does shit ton of works
## Tech Stack
1. Golang Fiber
2. MySQL
3. JWT Token
## Overview
This API deals with shits like managing users, books and borrowing books. All that complexity is written magnificiently using Clean Code Architecture so it is high in code readability
## How to run
### Install Go Programming Language
You can do so by visiting [https://go.dev/doc/install]
### Clone this repository
Run this following command :  
`git clone https://github.com/chaaaeeee/PERPUSTAKON.git`
### Download the dependencies
Change the directory  
`cd PERPUSTAKON`  
then download it's dependencies   
`go mod download`
### Create and set .env file
Create .env file
`touch .env`
Set .env file using this format below
```
SECRET=secret_key
DRIVER=driver
USER=username
PASSWORD=password
PROTOCOL=protocol
PATH=sock_path
DBNAME=dbname
```
### All Set!
Run the code by executing 
`go run *.go`
