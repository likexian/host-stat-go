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


type IOStat struct {
    Major        uint64 `json:"major"`         // major device number
    Minor        uint64 `json:"minor"`         // minor device number
    Name         string `json:"name"`          // device name
    ReadIOs      uint64 `json:"read_ios"`      // r/s total number of reads
    ReadMerges   uint64 `json:"read_merges"`   // rrqm/s number of read I/Os merged with in-queue I/O
    ReadSectors  uint64 `json:"read_sectors"`  // rsec/s number of 512 byte sectors read
    ReadTicks    uint64 `json:"read_ticks"`    // number of milliseconds spent reading
    WriteIOs     uint64 `json:"write_ios"`     // w/s total number of writes
    WriteMerges  uint64 `json:"write_merges"`  // wrqm/s number of write I/Os merged with in-queue I/O
    WriteSectors uint64 `json:"write_sectors"` // wsec/s number of 512 byte sectors written
    WriteTicks   uint64 `json:"write_ticks"`   // number of milliseconds spent writing
    InFlight     uint64 `json:"in_flight"`     // number of I/Os currently in progress
    IOTicks      uint64 `json:"io_ticks"`      // number of milliseconds spent doing I/Os(in flight)
    TimeInQueue  uint64 `json:"time_in_queue"` // total wait time for all requests in milliseconds
    ReadBytes    uint64 `json:"read_bytes"`    // rB/s
    WriteBytes   uint64 `json:"write_bytes"`   // wB/s
}


func GetIOStat() (stat []IOStat, err error) {
    text, err := ioutil.ReadFile("/proc/diskstats")
    if err != nil {
        return
    }

    stat = []IOStat{}
    lines := strings.Split(string(text), "\n")
    for _, v := range lines {
        if v == "" {
            continue
        }

        fields := strings.Fields(v)
        if fields[3] == "0" {
            continue
        }

        io_stat := IOStat{}
        io_stat.Major, _ = strconv.ParseUint(fields[0], 10, strconv.IntSize)
        io_stat.Minor, _ = strconv.ParseUint(fields[1], 10, strconv.IntSize)
        io_stat.Name = fields[2]
        io_stat.ReadIOs, _ = strconv.ParseUint(fields[3], 10, strconv.IntSize)
        io_stat.ReadMerges, _ = strconv.ParseUint(fields[4], 10, strconv.IntSize)
        io_stat.ReadSectors, _ = strconv.ParseUint(fields[5], 10, strconv.IntSize)
        io_stat.ReadTicks, _ = strconv.ParseUint(fields[6], 10, strconv.IntSize)
        io_stat.WriteIOs, _ = strconv.ParseUint(fields[7], 10, strconv.IntSize)
        io_stat.WriteMerges, _ = strconv.ParseUint(fields[8], 10, strconv.IntSize)
        io_stat.WriteSectors, _ = strconv.ParseUint(fields[9], 10, strconv.IntSize)
        io_stat.WriteTicks, _ = strconv.ParseUint(fields[10], 10, strconv.IntSize)
        io_stat.InFlight, _ = strconv.ParseUint(fields[11], 10, strconv.IntSize)
        io_stat.IOTicks, _ = strconv.ParseUint(fields[12], 10, strconv.IntSize)
        io_stat.TimeInQueue, _ = strconv.ParseUint(fields[13], 10, strconv.IntSize)

        io_stat.ReadBytes = io_stat.ReadSectors * 512
        io_stat.WriteBytes = io_stat.WriteSectors * 512

        stat = append(stat, io_stat)
    }

    return
}
