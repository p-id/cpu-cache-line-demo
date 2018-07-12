# CPU Cache line invalidation using Golang primitives

[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/p-id/cpu-cache-line-demo/master/LICENSE)

# Introduction
[False sharing](http://en.wikipedia.org/wiki/False_sharing) has been a pitfall for parallelism.
The difference becomes apparent with a performance penality for applications where simulataneous read/write updates are done on shared contexts allocated within a cache-line boundary.

Modern CPUs' tend to utilize cache for frequently updated memory locations. The minimize unit in CPU’ cache’ is a cache line, most CPUs' have a cache line of 64 byte thus when CPU read a variable from memory, it would read all variables nearby that variable. The problem of false sharing arises if one variable exists in two cache lines in different CPU cores. A load-update operation on one core would force other cores to update cache. Updating cache in multiple cores along with updating the actual location location in main memory as well leads to a freeze of few cycles across all participating cores.

The common way to solve above problem is to identify and align the frequently updated data-blobs (variables/structures) with the CPU cache line of 64 byte boundaries. That would force one variable to occupy a core’s cache line alone, so when other cores update a single variable other variables would not make that core reload the variable from memory.

# Demo application
To demonstrate false-sharing and cpu-cache invalidation penalty imposed, consider a common use-case of a common shared-context which is updated by one or more go-routines and read by another set of go-routines. This could be a regular case of maintaining and publishing counters.

```go
// SharedContextNoPadding vanilla container with no CPU cache-line padding
type SharedContextNoPadding struct {
	counterA uint64
	counterB uint64
}

// The above Golang struct could be updated to include CPU-cache line padding
// SharedContextWithPadding container with CPU cache-line padding
type SharedContextWithPadding struct {
	counterA uint64
	_p1      [8]uint64
	counterB uint64
}

// Ofcourse, this comes with added requirement of more amount of memory required in CPU caches (i.e 64 bytes) per shared variable.
// On a side note, Golang does use a pool on common-size objects so in some-cases the actual impact might be a bit less in terms of memory utilization 
```

The sample application is built using Golang. The application would require a golang development environment [setup](https://golang.org/doc/install)
To run the code/benchmark use the following command

$ go test -bench=.

# About
This project was created by [Piyush Dewnani](mailto:piyush@dewnani.net)
