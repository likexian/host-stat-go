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
    "io/ioutil"
)


type LoadStat struct {
    LoadNow      float64 `json:"load_now"`
    LoadPre      float64 `json:"load_pre"`
    LoadFar      float64 `json:"load_far"`
    ProcessTotal uint64  `json:"process_total"`
    ProcessRun   uint64  `json:"process_run"`
    ProcessIdle  uint64  `json:"process_idle"`
}


func GetLoadStat() (stat LoadStat, err error) {
    text, err := ioutil.ReadFile("/proc/loadavg")
    if err != nil {
        return
    }

    lines := strings.Split(string(text), "\n")
    fields := strings.Fields(lines[0])

    stat = LoadStat{}
    stat.LoadNow, _ = strconv.ParseFloat(fields[0], strconv.IntSize)
    stat.LoadPre, _ = strconv.ParseFloat(fields[1], strconv.IntSize)
    stat.LoadFar, _ = strconv.ParseFloat(fields[2], strconv.IntSize)

    stat.LoadNow = Round(stat.LoadNow, 2)
    stat.LoadPre = Round(stat.LoadPre, 2)
    stat.LoadFar = Round(stat.LoadFar, 2)

    processes := strings.Split(fields[3], "/")
    stat.ProcessRun, _ = strconv.ParseUint(processes[0], 10, strconv.IntSize)
    stat.ProcessTotal, _ = strconv.ParseUint(processes[1], 10, strconv.IntSize)
    stat.ProcessIdle = stat.ProcessTotal - stat.ProcessRun

    return
}
