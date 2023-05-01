# go-sensor

A simple golang backend that generate a mock up sensor data, and sends it to a gRPC server that collect the data. It also serve an API to configure the frequency of data generation.

## Getting Started

How to setup in local machine

### Prerequisites

- [Go](https://golang.org/doc/install) (version 1.19 or higher)
- [Protobuf Compiler](https://grpc.io/docs/protoc-installation/)

### Installation

1. Clone the repository:

```sh
git clone https://github.com/muhammadtaufan/go-sensor.git
```

2. go to the directory:

```sh
cd go-sensor
```

3. setup project:

```sh
make setup
```

4. build and run the service:

```sh
make run
```

### Setting up GoSensor with Docker Compose/Swarm

This guide will help you set up GoSensor, with Docker Compose or Docker Swarm.

Prerequisites
Before proceeding, make sure you have Docker installed on your system. You can download it from the official Docker website: https://www.docker.com/products/docker-desktop

##### Building the Docker image

```sh
docker build -t gosensor:latest .
```

##### Using Docker Swarm
To run GoSensor with Docker Swarm, use the following steps:

Create a network for GoSensor so we can connect with GoSensor Collector:

```sh
docker network create --driver overlay --scope swarm --attachable gosensor_network
```

Deploy the GoSensor service:

```sh
docker stack deploy -c docker-compose.yml gosensor-client
```

Check the logs:

```sh
docker service logs --follow gosensor-client_sensor_type
```
