<p align="center">
  <img src="images/logo.png" width="600">
</p>

Building a distributed log based on Travis Jeffery's [Distributed Services with Go](https://pragprog.com/book/tjgo/distributed-services-with-go)

## Overview

- *Record* — the data stored in our log
- *Store* — the file we store records in
- *Index* — the file we store index entries in
- *Segment* — the abstraction that ties a store and an index together
- *Log* — the abstraction that ties all the segments together


More to come
