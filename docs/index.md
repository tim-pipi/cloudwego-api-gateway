# CloudWeGo API Gateway Documentation

Welcome to the documentation for CloudWeGo API Gateway!

## Introduction

The CloudWeGo API Gateway is a powerful and lightweight API gateway, designed to deliver high performance and scalability. It is built on top of
[Kitex](https://github.com/cloudwego/kitex) and [Hertz](https://github.com/cloudwego/hertz), providing a seamless integration of their capabilities.

## Key Features

1. **HTTP Request Handling:** The API Gateway is designed to accept HTTP requests with JSON-encoded bodies, making it easy for clients to communicate with the backend services.

2. **Generic-Call with Kitex:** Leveraging the power of Kitex's Generic-Call feature, the API Gateway seamlessly translates JSON requests into Thrift binary format. This ensures efficient communication and compatibility between the client and backend services.

3. **Load Balancing:** To achieve optimal performance and reliability, the API Gateway integrates a robust load balancing mechanism.
   It efficiently distributes incoming requests among multiple backend RPC servers, ensuring even utilisation of resources.

4. **Service Registry and Discovery:** The API Gateway and RPC servers are seamlessly integrated with a service registry and discovery mechanism, using [etcd](https://github.com/etcd-io/etcd/releases/).

5. **Kitex-based RPC Servers:** As part of the project requirements, we have developed backend RPC servers using Kitex.

6. **Observability**: The CloudWeGo API Gateway is equipped with observability features,
   allowing developers to gain valuable insights into the system's behavior and performance.

The CloudWeGo API Gateway project is built with a focus on ease of use. With these key features in place, developers can confidently build and deploy robust microservices architectures, enhancing the overall efficiency of their applications.

## Installation

To get started with CloudWeGo API Gateway, follow the setup instructions in the [Setup Guide](setup.md)

## Feedback and Support

We value your feedback! If you have any questions or encounter any issues while using CloudWeGo API Gateway,
please [file an issue](https://github.com/tim-pipi/cloudwego-api-gateway/issues) on our GitHub page.
