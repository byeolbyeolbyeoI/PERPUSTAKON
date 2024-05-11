# PERPUSTAKON

Golang web-app to help library management.

## Tech Stack

Backend :

1. Gofiber
2. MySQL

Frontend :

1. HTML
2. CSS
3. Javascript

## Overview

This library management web application is high in features. It uses JWT token to authenticate users. There are 3 user roles in this web-app, admin, librarian and users. Admin deals with adding and deleting users, librarian deals with adding, deleting, borrowing and returning books, and users can see their status, the book list and it's availability. All that functionality is written magnificiently with a beatiful architecture so it is high in code-readability.

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
