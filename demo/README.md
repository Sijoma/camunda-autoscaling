# Demo

The demo will deploy the [sign-up-process.bpmn](worker/src/main/resources/bpmn/demo/sign-up-process.bpmn) process,
modeling a user sign up flow for the
premiere German cooking streaming service, Mettflix. This project provides deployment files to set up and operate the
demo application - _not the C8 cluster!_

![sign-up-process](./sign-up-process.png)

## Pre-requisites

The demo application reuses building blocks from the [camunda/camunda](https://github.com/camunda/camunda) repository,
as well as the publicly available [benchmark Helm chart](https://github.com/zeebe-io/benchmark-helm.git).

You will need:

- Docker 27+
- Java 21.0.3+

Helm is used to deploy the demo application, and Docker is used to build custom application images.

## Building the exporter

You can build the standalone exporter using `make demo-exporter`. This will
build the exporter JAR, as well as a Docker image which is to be used as an init container
to copy the exporter over to the broker.

## Building the starter/worker applications

The starter/worker applications are based on the Zeebe team's benchmark applications.
We build custom Docker images as we want to ship custom BPMN processes and payload for our
own application.

To build the customer applications, run `make demo-app`, which will build the Docker images
locally.