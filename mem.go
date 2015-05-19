/*
 * Go module for collecting host stat
 * http://www.likexian.com/
 *
 * Copyright 2014, Kexian Li
 * Released under the Apache License, Version 2.0
 *
 */

package host_stat


import (
    "strings"
    "strconv"
    "io/ioutil"
)


type MemStat struct {
    MemTotal  uint64  `json:"mem_total"`
    MemUsed   uint64  `json:"mem_used"`
    MemFree   uint64  `json:"mem_free"`
    Buffers   uint64  `json:"buffers"`
    Cached    uint64  `json:"cached"`
    SwapTotal uint64  `json:"swap_total"`
    SwapUsed  uint64  `json:"swap_used"`
    SwapFree  uint64  `json:"swap_free"`
    MemRate   float64 `json:"mem_rate"`
    SwapRate  float64 `json:"swap_rate"`
}


func GetMemStat() (stat MemStat, err error) {
    text, err := ioutil.ReadFile("/proc/meminfo")
    if err != nil {
        return
    }

    stat = MemStat{}
    lines := strings.Split(string(text), "\n")
    for i:=0; i<len(lines); i++ {
        if !strings.Contains(lines[i], ":") {
            continue
        }

        maps := strings.Split(lines[i], ":")
        key := strings.Trim(maps[0], " ")
        if key == "MemTotal" {
            stat.MemTotal = parse_mem_value(maps[1])
        } else if key == "MemFree" {
            stat.MemFree = parse_mem_value(maps[1])
        } else if key == "Buffers" {
            stat.Buffers = parse_mem_value(maps[1])
        } else if key == "Cached" {
            stat.Cached = parse_mem_value(maps[1])
        } else if key == "SwapTotal" {
            stat.SwapTotal = parse_mem_value(maps[1])
        } else if key == "SwapFree" {
            stat.SwapFree = parse_mem_value(maps[1])
        }
    }

    stat.MemUsed = stat.MemTotal - stat.MemFree
    stat.SwapUsed = stat.SwapTotal - stat.SwapFree
    if stat.MemTotal > 0 {
        stat.MemRate = Round(float64(stat.MemUsed) * 100 / float64(stat.MemTotal), 2)
    }
    if stat.SwapTotal > 0 {
        stat.SwapRate = Round(float64(stat.SwapUsed) * 100 / float64(stat.SwapTotal), 2)
    }

    return
}


func parse_mem_value(value string) (mem uint64) {
    data := strings.Fields(value)

    mem, _ = strconv.ParseUint(strings.Trim(data[0], " "), 10, strconv.IntSize)
    mem = mem / 1024

    return
}
