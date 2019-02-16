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
    "strings"
    "strconv"
)


type HostInfo struct {
    HostName  string `json:"host_name"`
    OSType    string `json:"os_type"`
    OSRelease string `json:"os_release"`
    Version   string `json:"version"`
    Release   string `json:"release"`
    OSBit     string `json:"os_bit"`
}


func GetHostInfo() (info HostInfo, err error) {
    info.HostName, err = ReadFirstLine("/proc/sys/kernel/hostname")
    if err != nil {
        return
    }

    info.OSType, err = ReadFirstLine("/proc/sys/kernel/ostype")
    if err != nil {
        return
    }

    info.OSRelease, err = ReadFirstLine("/proc/sys/kernel/osrelease")
    if err != nil {
        return
    }

    info.OSBit = ""
    if len(info.OSRelease) > 6 {
        if info.OSRelease[len(info.OSRelease) - 6:] == "x86_64" {
            info.OSBit = "64Bit"
        } else if info.OSRelease[len(info.OSRelease) - 4:] == "i686" {
            info.OSBit = "32Bit"
        }
    }

    if info.OSBit == "" {
        info.OSBit = strconv.Itoa(32 << uintptr(^uintptr(0) >> 63)) + "Bit"
    }

    info.Version, err = ReadFirstLine("/proc/sys/kernel/version")
    if err != nil {
        return
    }

    info.Release, err = GetRelease()
    if err != nil {
        return
    }

    return
}


func GetRelease() (name string, err error) {
    text := ""
    if IsFileExists("/etc/os-release") {
        text, err = ReadFile("/etc/os-release")
        if err != nil {
            return
        }

        lines := strings.Split(text, "\n")
        for _, l := range lines {
            if !strings.Contains(l, "=") {
                continue
            }

            ls := strings.Split(l, "=")
            ls[1] = strings.Trim(ls[1], "\"")
            if strings.Contains(ls[1], ",") {
                ls[1] = strings.Split(ls[1], ",")[0]
            }

            if ls[0] == "NAME" {
                name = ls[1]  // Fedora
            } else if ls[0] == "VERSION" {
                name = name + " " + ls[1]  // "17 (Beefy Miracle)"
                break
            } else if ls[0] == "PRETTY_NAME" {
                name = ls[1] // Debian GNU/Linux wheezy/sid
                break
            }
        }
    } else if IsFileExists("/etc/redhat-release") {
        text, err = ReadFirstLine("/etc/redhat-release")
        if err != nil {
            return
        }
        name = text
    } else if IsFileExists("/etc/SuSE-release") {
        text, err = ReadFirstLine("/etc/SuSE-release")
        if err != nil {
            return
        }
        name = text
    } else if IsFileExists("/etc/debian_version") {
        text, err = ReadFirstLine("/etc/debian_version")
        if err != nil {
            return
        }
        name = "Debian " + text
    } else if IsFileExists("/etc/debian_release") {
        text, err = ReadFirstLine("/etc/debian_release")
        if err != nil {
            return
        }
        name = "Debian " + text
    } else if IsFileExists("/etc/lsb-release") {
        text, err = ReadFile("/etc/lsb-release")
        if err != nil {
            return
        }

        lines := strings.Split(text, "\n")
        for _, l := range lines {
            if !strings.Contains(l, "=") {
                continue
            }

            ls := strings.Split(l, "=")
            if ls[0] == "DISTRIB_ID" {
                name = ls[1]  // Ubuntu
            } else if ls[0] == "DISTRIB_RELEASE" {
                name = name + " " + ls[1]  // 9.04
                break
            }
        }
    }

    if strings.Contains(name, "(") {
        names := strings.Split(name, "(")
        name = names[0]
    }

    if strings.Contains(name, "/") {
        names := strings.Split(name, "/")
        name = strings.Join(names[:len(names) - 1], "/")
    }
    name = strings.Trim(name, " ")

    return
}
