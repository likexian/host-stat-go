/*
 * Go module for collecting host stat
 * https://www.likexian.com/
 *
 * Copyright 2014-2019, Li Kexian
 * Released under the Apache License, Version 2.0
 *
 */

package host_stat


import (
    "os"
    "runtime"
    "testing"
)


func assertNotError(t *testing.T, err error) {
    if err != nil {
        _, file, line, _ := runtime.Caller(1)
        t.Errorf("%s:%d", file, line)
        t.Errorf(err.Error())
        t.FailNow()
    }
}


func TestHostStat(t *testing.T) {
    host_info, err := GetHostInfo()
    assertNotError(t, err)
    t.Log(host_info)

    cpu_info, err := GetCPUInfo()
    assertNotError(t, err)
    t.Log(cpu_info)

    cpu_stat, err := GetCPUStat()
    assertNotError(t, err)
    t.Log(cpu_stat)

    mem_stat, err := GetMemStat()
    assertNotError(t, err)
    t.Log(mem_stat)

    disk_stat, err := GetDiskStat()
    assertNotError(t, err)
    t.Log(disk_stat)

    io_stat, err := GetIOStat()
    if err != nil {
        if e, ok := err.(*os.PathError); ok {
            t.Log(e)
        } else {
            assertNotError(t, err)
        }
    }
    t.Log(io_stat)

    net_stat, err := GetNetStat()
    assertNotError(t, err)
    t.Log(net_stat)

    uptime_stat, err := GetUptimeStat()
    assertNotError(t, err)
    t.Log(uptime_stat)

    load_stat, err := GetLoadStat()
    assertNotError(t, err)
    t.Log(load_stat)
}
