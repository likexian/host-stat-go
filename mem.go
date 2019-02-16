/*
 * Go module for collecting host stat
 * https://www.likexian.com/
 *
 * Copyright 2014-2019, Li Kexian
 * Released under the Apache License, Version 2.0
 *
 */

package hoststat

import (
	"io/ioutil"
	"strconv"
	"strings"
)

// MemStat storing memory stat
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

// GetMemStat returns memory stat
func GetMemStat() (stat MemStat, err error) {
	text, err := ioutil.ReadFile("/proc/meminfo")
	if err != nil {
		return
	}

	stat = MemStat{}
	lines := strings.Split(string(text), "\n")
	for i := 0; i < len(lines); i++ {
		if !strings.Contains(lines[i], ":") {
			continue
		}

		maps := strings.Split(lines[i], ":")
		key := strings.TrimSpace(maps[0])
		if key == "MemTotal" {
			stat.MemTotal = parseMemValue(maps[1])
		} else if key == "MemFree" {
			stat.MemFree = parseMemValue(maps[1])
		} else if key == "Buffers" {
			stat.Buffers = parseMemValue(maps[1])
		} else if key == "Cached" {
			stat.Cached = parseMemValue(maps[1])
		} else if key == "SwapTotal" {
			stat.SwapTotal = parseMemValue(maps[1])
		} else if key == "SwapFree" {
			stat.SwapFree = parseMemValue(maps[1])
		}
	}

	stat.MemUsed = stat.MemTotal - stat.MemFree
	stat.SwapUsed = stat.SwapTotal - stat.SwapFree
	if stat.MemTotal > 0 {
		stat.MemRate = Round(float64(stat.MemUsed)*100/float64(stat.MemTotal), 2)
	}
	if stat.SwapTotal > 0 {
		stat.SwapRate = Round(float64(stat.SwapUsed)*100/float64(stat.SwapTotal), 2)
	}

	return
}

func parseMemValue(value string) (mem uint64) {
	data := strings.Fields(value)

	mem, _ = strconv.ParseUint(strings.TrimSpace(data[0]), 10, strconv.IntSize)
	mem = mem / 1024

	return
}
