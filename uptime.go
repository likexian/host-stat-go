/*
 * Go module for collecting host stat
 * https://www.likexian.com/
 *
 * Copyright 2014-2015, Li Kexian
 * Released under the Apache License, Version 2.0
 *
 */

package host_stat


import (
    "strings"
    "strconv"
)


type UptimeStat struct {
    Uptime   float64 `json:"uptime"`
    IdleTime float64 `json:"idle_time"`
    IdleRate float64 `json:"idle_rate"`
}


func GetUptimeStat() (stat UptimeStat, err error) {
    text, err := ReadFirstLine("/proc/uptime")
    if err != nil {
        return
    }

    lines := strings.Split(text, "\n")
    fields := strings.Fields(lines[0])

    stat = UptimeStat{}
    stat.Uptime, _ = strconv.ParseFloat(fields[0], strconv.IntSize)
    stat.IdleTime, _ = strconv.ParseFloat(fields[1], strconv.IntSize)
    stat.IdleRate = Round(stat.IdleTime * 100 / stat.Uptime, 2)

    stat.Uptime = Round(stat.Uptime, 2)
    stat.IdleTime = Round(stat.IdleTime, 2)

    return
}
