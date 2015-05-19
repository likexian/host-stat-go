/*
 * Go module for collecting host stat
 * https://www.likexian.com/
 *
 * Copyright 2014-2015, Li Kexian
 * Released under the Apache License, Version 2.0
 *
 */

package host_stat


import(
    "os"
    "math"
    "strings"
    "io/ioutil"
)


func Round(data float64, precision int) (result float64) {
    pow := math.Pow(10, float64(precision))
    digit := pow * data
    _, div := math.Modf(digit)

    if div >= 0.5 {
        result = math.Ceil(digit)
    } else {
        result = math.Floor(digit)
    }
    result = result / pow

    return
}


func IsFileExists(fname string) (bool) {
    if _, err := os.Stat(fname); err == nil {
        return true
    }
    return false
}


func ReadFirstLine(fname string) (line string, err error) {
    text, err := ReadFile(fname)
    if err != nil {
        return
    }

    lines := strings.Split(string(text), "\n")
    line = strings.Trim(lines[0], " ")

    return
}


func ReadFile(fname string) (result string, err error) {
    text, err := ioutil.ReadFile(fname)
    if err != nil {
        return
    }
    result = string(text)

    return
}
