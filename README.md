<p align="center">
  <img src="images/logo-2.png" width="600">
</p>

[![Actions Status](https://github.com/wbrowne/chronicle/workflows/Go/badge.svg)](https://github.com/wbrowne/chronicle/actions)

Building a distributed log based on Travis Jeffery's [Distributed Services with Go](https://pragprog.com/book/tjgo/distributed-services-with-go)

## Build

- Install go (tested with 1.13)
    - https://golang.org/dl/
- Install protoc binary (tested with 3.11.4)
    - https://github.com/protocolbuffers/protobuf/releases
- Install protoc-gen-go (plugin for protobuf compiler to generate go code)
    - `go get github.com/gogo/protobuf/protoc-gen-gogo`
- Install Cloudflare cert CLI tools
    - `go get github.com/cloudflare/cfssl/cmd/cfssl@v1.4.1`
    - `go get github.com/cloudflare/cfssl/cmd/cfssljson@v1.4.1`
- Execute
    - `make init`
    - `make compile`
    - `make gencert`
    - `make test`

## Overview

- *Record* — the data stored in our log
- *Store* — the file we store records in
- *Index* — the file we store index entries in
- *Segment* — the abstraction that ties a store and an index together
- *Log* — the abstraction that ties all the segments together

More to come
