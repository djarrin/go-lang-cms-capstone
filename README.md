# go-lang-cms-capstone
This is the capstone project for my Golang nanodegree. It is a mock CMS backend that essentially consists of a CRUD API. I also took the liberty of setting up a circle CI CI/CD pipeline that consists of setting up and EC2 instance, installing the code and doing a quick smoke test to ensure the application is up. 

## Installation
There should be two ways to run this project
1. clone this repo into your go folder and run `go run main.go` in the base directory
2. clone this repo into your go folder and run `go build main.go` in the base directory then run `./main`

Then you should be able to access the API's home page by going to localhost:3000, the home page will describe the endpoints and how to use them

## Dependencies 
- You must have go installed
- github.com/gorilla/mux v1.8.0 (should be installed automatically when running go run)

