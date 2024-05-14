# PERPUSTAKON

Golang web-app to help library management.

## Tech Stack

### Backend :

1. Gofiber
2. MySQL

### Frontend :

1. HTML
2. CSS
3. Javascript

## Overview

## Endpoints

### Frontend

1. `/user/dashboard`
2. `/user/bookList`
3. `/admin/dashboard`
4. `/admin/userList`
5. `/admin/addUser`
6. `/admin/deleteUser`
7. `/librarian/dashboard`
8. `/librarian/bookList`
9. `/librarian/userList`
10. `/librarian/addBook`
11. `/librarian/deleteBook`
12. `/librarian/borrowBook`
13. `/librarian/returnBook`

### API

1. `/signupHandler`
2. `/loginHandler`
3. `/getUsers`
4. `/getUserById/:id`
5. `/getUserById/:id`
6. `/addUser`
7. `/deleteUser`
8. `/getBooks`
9. `/getBookById/:id`
10. `/getBookByTitle/:title`
11. `/addBook`
12. `/deleteBook`
13. `/borrowBook`
14. `/returnBook`

## How to run

### Install Go Programming Language

You can do so by visiting [https://go.dev/doc/install]

### Clone this repository

Run this following command : `git clone https://github.com/chaaaeeee/PERPUSTAKON.git`

### Download the dependencies

Change the directory `cd PERPUSTAKON`  
then download it's dependencies `go mod download`

### Create and set .env file

Create .env file `touch .env` then set the variables using this format below

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
