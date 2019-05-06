# host-stat

[![License](https://img.shields.io/badge/license-Apache%202.0-blue.svg)](LICENSE)
[![GoDoc](https://godoc.org/github.com/likexian/host-stat-go?status.svg)](https://godoc.org/github.com/likexian/host-stat-go)
[![Build Status](https://travis-ci.org/likexian/host-stat-go.svg?branch=master)](https://travis-ci.org/likexian/host-stat-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/likexian/host-stat-go)](https://goreportcard.com/report/github.com/likexian/host-stat-go)
[![Code Cover](https://codecov.io/gh/likexian/host-stat-go/graph/badge.svg)](https://codecov.io/gh/likexian/host-stat-go)

host-stat-go is a Go module for collecting host stat.

## Overview

This module provided functions to collect cpu/mem/disk/io/load/uptime/kernel info of the host.

*Work for popular LINUX distributions ONLY*

## Installation

    go get -u github.com/likexian/host-stat-go

## Importing

    import (
        "github.com/likexian/host-stat-go"
    )

## Documentation

Visit the docs on [GoDoc](https://godoc.org/github.com/likexian/host-stat-go)

## Example

Get the memory stat

```go
memStat, err := hoststat.GetMemStat()
if err == nil {
    // print total memory of host in KB
    fmt.Println(memStat.MemTotal)

    // print used memory of host in KB
    fmt.Println(memStat.MemUsed)

    // print free memory of host in KB
    fmt.Println(memStat.MemFree)

    // print used memory rate of host in percent
    fmt.Println(memStat.MemRate)
}
```

## LICENSE

Copyright 2014-2019 Li Kexian

Licensed under the Apache License 2.0

## About

- [Li Kexian](https://www.likexian.com/)

## DONATE

- [Help me make perfect](https://www.likexian.com/donate/)
