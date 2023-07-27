# CloudWeGo API Gateway

This repository contains the code for the implementation of an API Gateway for Orbital 2023.

CloudWeGo API Gateway is an API Gateway based on `CloudWeGo` projects, using technologies, such as `Kitex` - RPC framework,
and `Hertz` - HTTP framework. This API Gateway can handle JSON-encoded HTTP requests and utilize `Kitex`'s Generic-Call feature.
to convert these requests into `Thrift` binary format. The API Gateway will then route the requests to one of the backend.
RPC servers obtained from the registry center.

## CloudWeGo

[CloudWeGo](https://www.cloudwego.io/) is an open-source middleware set launched by ByteDance that can be used to
quickly build enterprise-class cloud native architectures. It contains many components,
including the RPC framework Kitex, the HTTP framework Hertz, the basic network library
Netpoll, thrfitgo, etc. By combining the community's excellent open source products,
developers can quickly build a complete set of microservices systems.

## Hertz

[Hertz](https://www.cloudwego.io/docs/hertz/) [həːts] is a high-performance, high-usability, extensible HTTP framework for Go. It’s
designed to make it easy for developers to build microservices.
Inspired by other open source frameworks, combined with the unique challenges we met in
ByteDance, Hertz has become production-ready and has powered ByteDance’s internal
services over the years.

## Kitex

[Kitex](https://www.cloudwego.io/docs/kitex/) [kaɪt’eks] is a high-performance and strong-extensibility Golang RPC framework that
helps developers build microservices. If performance and extensibility are the main concerns
when you develop microservices, Kitex can be a good choice.

## Useful Links

- [Usage Instructions](https://tim-pipi.github.io/cloudwego-api-gateway/)
- [Milestone I Submission](https://drive.google.com/drive/u/0/folders/1mm--TjLNb5FZXAquGjFT_0S7Nf_3PMf1)
- [Milestone II Submission](https://drive.google.com/drive/folders/1ZqQKP6_HXSqQ5CiKRAptCXUhe7ADz-Yu?usp=drive_link)
- [Milestone III Submission](https://drive.google.com/drive/folders/1hPAVGZ5VUchy4TIrkm3r4U3CB7aM3K1Z?usp=sharing)
- [System Design Document](https://docs.google.com/document/d/1ZIIul1IiEUxCzst-V_idwqwq4sc_vt4DPmeqCSBlbmY/edit?usp=sharing)

## API Gateway Diagram

![API Gateway Diagram](gateway.png)
