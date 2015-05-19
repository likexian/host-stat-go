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
    "fmt"
    "os"
    "testing"
    "github.com/bmizerany/assert"
)


func TestHostStat(t *testing.T) {
    host_info, err := GetHostInfo()
    assert.Equal(t, nil, err)
    fmt.Println(host_info)

    cpu_info, err := GetCPUInfo()
    assert.Equal(t, nil, err)
    fmt.Println(cpu_info)

    cpu_stat, err := GetCPUStat()
    assert.Equal(t, nil, err)
    fmt.Println(cpu_stat)

    mem_stat, err := GetMemStat()
    assert.Equal(t, nil, err)
    fmt.Println(mem_stat)

    disk_stat, err := GetDiskStat()
    assert.Equal(t, nil, err)
    fmt.Println(disk_stat)

    io_stat, err := GetIOStat()
    if err != nil {
        if e, ok := err.(*os.PathError); ok {
            fmt.Println(e)
        } else {
            assert.Equal(t, nil, err)
        }
    }
    fmt.Println(io_stat)

    net_stat, err := GetNetStat()
    assert.Equal(t, nil, err)
    fmt.Println(net_stat)

    uptime_stat, err := GetUptimeStat()
    assert.Equal(t, nil, err)
    fmt.Println(uptime_stat)

    load_stat, err := GetLoadStat()
    assert.Equal(t, nil, err)
    fmt.Println(load_stat)
}
