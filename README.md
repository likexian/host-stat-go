# host-stat

host-stat-go is a Go module for collecting host stat.

[![Build Status](https://secure.travis-ci.org/likexian/host-stat-go.png)](https://secure.travis-ci.org/likexian/host-stat-go)

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

CPU information

    func GetCPUInfo() (info CPUInfo, err error)

CPU usage stat

    func GetCPUStat() (stat CPUStat, err error)

Memory information and usage stat

    func GetMemStat() (stat MemStat, err error)

Disk information

    func GetDiskStat() (stat []DiskStat, err error)

Disk IO stat

    func GetIOStat() (stat []IOStat, err error)

Network stat

    func GetNetStat() (stat []NetStat, err error)

Host load stat

    func GetLoadStat() (stat LoadStat, err error)

Host uptime stat

    func GetUptimeStat() (stat UptimeStat, err error)

Host and Kernel information

    func GetHostInfo() (info HostInfo, err error)

## Example

Get the memory stat

    mem_stat, err := host_stat.GetMemStat()
    if err != nil {
        // print total memory of host in KB
        fmt.Println(mem_stat.MemTotal)

        // print used memory of host in KB
        fmt.Println(mem_stat.MemUsed)

        // print free memory of host in KB
        fmt.Println(mem_stat.MemFree)

        // print used memory rate of host in percent
        fmt.Println(mem_stat.MemRate)
    }

## LICENSE

Copyright 2014-2019, Li Kexian

Apache License, Version 2.0

## About

- [Li Kexian](https://www.likexian.com/)

## DONATE

- [Help me make perfect](https://www.likexian.com/donate/)
