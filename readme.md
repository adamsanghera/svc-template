# Microservice

This is an example microservice application.  

## Architecture

We're now going to go over the directory structure.

### src

`src` contains all of the code used to "power" the microservice.

The directory contains `bus.go`, which is a package designed to connect a service to rabbitmq, a pub/sub service used to buffer 1-1, 1-N, and N-1 message passing.

It also contains `listeners` and `publishers` which are described later on in detail.

Finally, it contains `queues`, which is a directory containing 

## Usage

As a user, you want easy access to two functions.  Both of these functions touch queues, which need to be instantiated in your main method.

Before calling init, you need to declare some queues.  After the Bus is initiated, publishes and subscribes may be called without restriction.

One can access these by importing `bus.go`, which gives access to `DeclareQueue`, `Publish` and `Subscribe`.



### Listens

Inside `src`, you'll find a `listeners` directory.  To create a new listener handler, simply create a new file in here, following the structure of the example provided.  

### Publishes

