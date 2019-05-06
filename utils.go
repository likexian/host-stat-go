/*
 * Copyright 2014-2019 Li Kexian
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * Go module for collecting host stat
 * https://www.likexian.com/
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
