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
