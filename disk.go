// +build !windows

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
    "syscall"
    "strings"
    "io/ioutil"
)


type DiskStat struct {
    Label    string  `json:"label"`
    Mount    string  `json:"mount"`
    Total    uint64  `json:"total"`
    Used     uint64  `json:"used"`
    Free     uint64  `json:"free"`
    UsedRate float64 `json:"used_rate"`
}


func GetDiskStat() (stat []DiskStat, err error) {
    text, err := ioutil.ReadFile("/etc/mtab")
    if err != nil {
        return
    }

    stat = []DiskStat{}
    lines := strings.Split(string(text), "\n")
    for i:=0; i<len(lines); i++ {
        lines[i] = strings.Trim(lines[i], " ")
        if lines[i] == "" {
            continue
        }

        fields := strings.Fields(lines[i])
        if len(fields) < 3 {
            continue
        }

        if fields[0] != "none" && fields[0] != "proc" && fields[0] != "sysfs" && fields[0] != "devpts" {
            disk_stat, _ := GetStat(fields[1])
            disk_stat.Label = fields[0]
            stat = append(stat, disk_stat)
        }
    }

    return
}


func GetStat(path string) (stat DiskStat, err error) {
    fs := syscall.Statfs_t{}
    err = syscall.Statfs(path, &fs)
    if err != nil {
        return
    }

    stat = DiskStat{}
    stat.Mount = path
    stat.Total = fs.Blocks * uint64(fs.Bsize) / (1024 * 1024)
    stat.Free = fs.Bfree * uint64(fs.Bsize) / (1024 * 1024)
    stat.Used = stat.Total - stat.Free
    stat.UsedRate = Round(float64(stat.Used) * 100 / float64(stat.Total), 2)

    return
}
