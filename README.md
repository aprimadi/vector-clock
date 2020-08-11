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

```
package main

import (
  "github.com/aprimadi/vector-clock"
)

func main() {
  // Initialize empty vector clock
  v1 := vclock.VClock{}

  // Increment vector clock for a given process id
  v1.tick("pid1")     // v1 = VClock{"pid1": 1}

  // Create a copy of v1 and advance clock for pid2
  v2 := v1.Copy()
  v2.tick("pid2")     // v2 = VClock{"pid1": 1, "pid2": 1}

  // Create a copy of v1 and advance clock for pid3
  v3 := v1.Copy()
  v3.tick("pid3")     // v3 = VClock{"pid1": 1, "pid3": 1}

  // Use relation to find out the relation of two vector clocks
  r1 := v2.Relation(v3) // Conflict
  r2 := v2.Relation(v1) // Descendant

  // Merge two vector clocks
  v4 := v2.Merge(v3)  // v4 = VClock{"pid1": 1, "pid2": 1, "pid3": 1}
}
```

## Serializing/deserializing

Since VClock uses map[string]uint64 as its underlying implementation, it can be
serialized directly over the network using any serialization protocol you
choose such as JSON, msgpack, etc.
