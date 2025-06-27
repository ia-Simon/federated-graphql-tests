# Go GraphQL Example

> ⚠️ WIP: This repository is a Work In Progress, beware the rough edges.

This repository contains a simple distributed GraphQL system, composed of several federated subgraphs written in Go using the [GQLGen library](https://gqlgen.com). The composition and serving of the complete supergraph is provided by the Apollo Router.

## Components

| Component                              | Description                                                                             |
| -------------------------------------- | --------------------------------------------------------------------------------------- |
| [Router](./router)                     | Apollo Router, a powerful GraphQL federated planner written in Rust                     |
| [User Subgraph](./user-subgraph)       | Subgraph dedicated to serve user information                                            |
| [Todo Subgraph](./todo-subgraph)       | Subgraph dedicated to serve todo information                                            |
| [Product Subgraph](./product-subgraph) | Subgraph dedicated to serve product information                                         |
| [Review Subgraph](./review-subgraph)   | Subgraph dedicated to serve review information                                          |
| [Sys Headers Library](./sys-headers)   | Example library containing useful helpers to retrieve and provide common system headers |
