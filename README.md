Vector Clock
============

Vector clock is a building block for allowing multiple versions of data in a
distributed systems. It's being used on a distributed key value store such as
[Dynamo](https://www.allthingsdistributed.com/files/amazon-dynamo-sosp2007.pdf).
This library is an implementation of a vector clock in Golang.

## Installation

```
go get -u github.com/aprimadi/vector-clock
```

## Usage

```go
package main

import (
  "github.com/aprimadi/vector-clock"
)

func main() {
  // Initialize empty vector clock
  v1 := vclock.VClock{}

  // Increment vector clock for a given process id
  v1.Tick("pid1")     // v1 = VClock{"pid1": 1}

  // Create a copy of v1 and advance clock for pid2
  v2 := v1.Copy()
  v2.Tick("pid2")     // v2 = VClock{"pid1": 1, "pid2": 1}

  // Create a copy of v1 and advance clock for pid3
  v3 := v1.Copy()
  v3.Tick("pid3")     // v3 = VClock{"pid1": 1, "pid3": 1}

  // Use relation to find out the relation of two vector clocks
  v2.Relation(v3)     // Conflict
  v2.Relation(v1)     // Descendant

  // Merge two vector clocks
  v2.Merge(v3)        // VClock{"pid1": 1, "pid2": 1, "pid3": 1}
}
```

## Serializing/deserializing

Since VClock uses map[string]uint64 as its underlying implementation, it can be
serialized directly over the network using any serialization protocol you
choose such as JSON, msgpack, etc.
