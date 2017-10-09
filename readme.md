# svc-template
A generic microservice written in golang

# What does it do?

This package provides a simple, extensible framework for creating microservices.

# But Microservice is such a generic term...

That's right!

A few assumptions have been made about microservices.

Namely, all microservices

1. Listen for TCP requests on a given port.
2. Handle those requests with some internal logic.
3. Send TCP requests to other microservices.

The goal of this project is to simplify steps 1 and 2 as much as possible, by making use of sensible defaults, 

However, this project has been designed in such a way as to make its composite pieces as modular as possible.

# Handling inter-service Communication

RabbitMQ is the communication method of choice, because of the nice guarantees it makes about message delivery.

However, you could just as easily replace the