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


type NetStat struct {
    Device       string `json:"device"`
    RXBytes      uint64 `json:"rx_bytes"`
    RXPackets    uint64 `json:"rx_packets"`
    RXErrs       uint64 `json:"rx_errs"`
    RXDrop       uint64 `json:"rx_drop"`
    RXFifo       uint64 `json:"rx_fifo"`
    RXFrame      uint64 `json:"rx_frame"`
    RXCompressed uint64 `json:"rx_compressed"`
    RXMulticast  uint64 `json:"rx_multicase"`
    TXBytes      uint64 `json:"tx_bytes"`
    TXPackets    uint64 `json:"tx_packetss"`
    TXErrs       uint64 `json:"tx_errs"`
    TXDrop       uint64 `json:"tx_drop"`
    TXFifo       uint64 `json:"tx_fifo"`
    TXColls      uint64 `json:"tx_colls"`
    TXCarrier    uint64 `json:"tx_carrier"`
    TXCompressed uint64 `json:"tx_compressed"`
}


func GetNetStat() (stat []NetStat, err error) {
    text, err := ioutil.ReadFile("/proc/net/dev")
    if err != nil {
        return
    }

    stat = []NetStat{}
    lines := strings.Split(string(text), "\n")
    for _, v := range lines {
        if v == "" || !strings.Contains(v, ":") {
            continue
        }

        net_stat := NetStat{}
        fields := strings.Split(v, ":")
        net_stat.Device = strings.Trim(fields[0], " ")

        fields = strings.Fields("Device " + fields[1])
        net_stat.RXBytes, _ = strconv.ParseUint(fields[1], 10, strconv.IntSize)
        net_stat.RXPackets, _ = strconv.ParseUint(fields[2], 10, strconv.IntSize)
        net_stat.RXErrs, _ = strconv.ParseUint(fields[3], 10, strconv.IntSize)
        net_stat.RXDrop, _ = strconv.ParseUint(fields[4], 10, strconv.IntSize)
        net_stat.RXFifo, _ = strconv.ParseUint(fields[5], 10, strconv.IntSize)
        net_stat.RXFrame, _ = strconv.ParseUint(fields[6], 10, strconv.IntSize)
        net_stat.RXCompressed, _ = strconv.ParseUint(fields[7], 10, strconv.IntSize)
        net_stat.RXMulticast, _ = strconv.ParseUint(fields[8], 10, strconv.IntSize)
        net_stat.TXBytes, _ = strconv.ParseUint(fields[9], 10, strconv.IntSize)
        net_stat.TXPackets, _ = strconv.ParseUint(fields[10], 10, strconv.IntSize)
        net_stat.TXErrs, _ = strconv.ParseUint(fields[11], 10, strconv.IntSize)
        net_stat.TXDrop, _ = strconv.ParseUint(fields[12], 10, strconv.IntSize)
        net_stat.TXFifo, _ = strconv.ParseUint(fields[13], 10, strconv.IntSize)
        net_stat.TXColls, _ = strconv.ParseUint(fields[14], 10, strconv.IntSize)
        net_stat.TXCarrier, _ = strconv.ParseUint(fields[15], 10, strconv.IntSize)
        net_stat.TXCompressed, _ = strconv.ParseUint(fields[16], 10, strconv.IntSize)

        stat = append(stat, net_stat)
    }

    return
}
