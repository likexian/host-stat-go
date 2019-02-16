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
	"io/ioutil"
	"math"
	"os"
	"strings"
)

// Round returns math round
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

// IsFileExists returns is file exists
func IsFileExists(fname string) bool {
	if _, err := os.Stat(fname); err == nil {
		return true
	}
	return false
}

// ReadFirstLine returns the first line of file
func ReadFirstLine(fname string) (line string, err error) {
	text, err := ReadFile(fname)
	if err != nil {
		return
	}

	lines := strings.Split(string(text), "\n")
	line = strings.TrimSpace(lines[0])

	return
}

// ReadFile returns text of file
func ReadFile(fname string) (result string, err error) {
	text, err := ioutil.ReadFile(fname)
	if err != nil {
		return
	}

	result = string(text)

	return
}
