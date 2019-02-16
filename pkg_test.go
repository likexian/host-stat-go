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
    hostInfo, err := GetHostInfo()
    assertNotError(t, err)
    t.Log(hostInfo)

    cpuInfo, err := GetCPUInfo()
    assertNotError(t, err)
    t.Log(cpuInfo)

    cpuStat, err := GetCPUStat()
    assertNotError(t, err)
    t.Log(cpuStat)

    memStat, err := GetMemStat()
    assertNotError(t, err)
    t.Log(memStat)

    diskStat, err := GetDiskStat()
    assertNotError(t, err)
    t.Log(diskStat)

    ioStat, err := GetIOStat()
    if err != nil {
        if e, ok := err.(*os.PathError); ok {
            t.Log(e)
        } else {
            assertNotError(t, err)
        }
    }
    t.Log(ioStat)

    netStat, err := GetNetStat()
    assertNotError(t, err)
    t.Log(netStat)

    uptimeStat, err := GetUptimeStat()
    assertNotError(t, err)
    t.Log(uptimeStat)

    loadStat, err := GetLoadStat()
    assertNotError(t, err)
    t.Log(loadStat)
}
