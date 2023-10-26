# Backend Assignement
## Project Name: Vessel tracking System
### Question Link
```
https://shippioinc.notion.site/Backend-Coding-Assignment-72da02b75f2943b598d4c0245b0fb3f0
```
##  Problem statement
### Vessel tracking

Because ocean transportation is prone to delays due to a variety of causes, our customers always think about when their cargo will arrive. Therefore our operators track vessel statuses so that customers can easily locate the whereabouts of their shipments and know when a vessel will depart or arrive at a port.

## Requirements
### Minimum requirements

Write a simple API that allows consumers to manage vessels. The APIs should allow the consumer to create, update, and retrieve a list of vessels from the server. A vessel should have the following information:

- Name (string)
- Owner ID (string)
- NACCS code (string) which is unique

### Bonus
Write a simple API that allows consumers to manage vessel voyages. A voyage is the transit of a vessel that takes it from one location to another between specified times. Traveling from LA to NYC on January 1st from 10:00 am to 1:00 pm is a single voyage for an airline. The API should allow the consumer to create, update, and retrieve the current voyage for a single vessel.

### Methods handled in this project
make sure to read vessel.proto file to see messages for each.
```
service VesselService {
//vessel related rpc
    rpc CreateVessel(CreateVesselRequest) returns (CreateVesselResponse) {}
    rpc GetVessels(GetVesselsRequest) returns (GetVesselsResponse) {}
    rpc GetVessel(GetVesselRequest) returns (GetVesselResponse) {}
    rpc UpdateVessel(UpdateVesselRequest)returns(UpdateVesselResponse){}

// voyage related rpc
    rpc CreateVoyage(CreateVoyageRequest) returns (CreateVoyageResponse) {}
    rpc GetVoyage(GetVoyageRequest) returns (GetVoyageResponse) {}
    rpc UpdateVoyage(UpdateVoyageRequest)returns(UpdateVoyageResponse){}
}
```

## Prerequisites
- golang installed
- Docker
- Grpc libraries
- protoc

## Tech used
- Golang
- Grpc for serving Api

## Installation
Follow the steps to run the project with ease
1) unzip the project
2) Go inside project Vessel_tracking
3) Run command to make database

```
make up-dev
```
 Connect the database with any software i use Sequel Ace  .Follow the .envrc file to get to know about the Database_port,Server address etc

 4) After connecting copy the query and run the queries below to make database and tables
 ```
 create database vessel_tracking;
 ```
 ```

 CREATE TABLE Vessel (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    owner_id VARCHAR(255) NOT NULL,
    naccs_code VARCHAR(255) NOT NULL UNIQUE
);
CREATE TABLE Voyages (
    voyage_id INT AUTO_INCREMENT PRIMARY KEY,
    vessel_id INT NOT NULL,
    departure_location VARCHAR(255) NOT NULL,
    arrival_location VARCHAR(255) NOT NULL,
    departure_time DATETIME NOT NULL,
    arrival_time DATETIME NOT NULL,
    details TEXT,
    FOREIGN KEY (vessel_id) REFERENCES Vessel(id)
);
 ```
 Database setup is done

5) Now go to server folder
```
cd cmd/app/server
```
6) run the main.go file to run the program
```
go run main.go
```

7) (Optional) If by any chance you are facing problem with dependency just run
```
go mod tidy
```

> [!IMPORTANT]
> Make sure you have all the environment variables present so that your program can read your configurations . I use .envrc file to store my comfigs .
### my .envrc file
```
export DEBUG=true
export SERVER_ADDRESS=0.0.0.0:8080
export DATABASE_HOST=0.0.0.0
export DATABASE_PORT=3306
export DATABASE_USER=root
export DATABASE_PASSWORD=123456
export DATABASE_NAME=vessel_tracking
```
> [!IMPORTANT]
> Make sure to run ```source .envrc``` to load configs  .


### Some commands to ease things out
For test
```
make test
```
For generating proto files
```
make codegen-proto
```

## Folder structure
```
├── Makefile
├── Readme.md
├── app.env
├── buf.gen.yaml
├── buf.work.yaml
├── cmd
│   ├── app
│   │   └── server
│   │       └── main.go   file to run program
│   └── vesselpb
│       ├── buf.yaml
│       ├── vessel.pb.go
│       ├── vessel.proto   Proto  file to  declare messages
│       └── vessel_grpc.pb.go
├── docker-compose.yaml
├── generate.sh
├── go.mod
├── go.sum
├── internal
│   ├── domain
│   │   ├── model
│   │   │   ├── vessel.go
│   │   │   └── voyage.go
│   │   └── service   main business logic lies here
│   │       ├── service.go
│   │       ├── service_test.go
│   │       └── usecase.go
│   ├── errors
│   │   └── errors.go
│   ├── handler
│   │   ├── handler.go
│   │   └── helper.go
│   └── infrastructure   layer to talk to  database
│       └── datastore
│           ├── helper.go
│           ├── repository.go
│           ├── vessel.go
│           └── vessel_test.go
├── mocks  // For mocking purpose for testing
│   ├── repository.go
│   └── usecase.go
└── utils
    ├── config.go
    └── logger.go

```