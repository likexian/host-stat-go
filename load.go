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
	"github.com/likexian/gokit/xhuman"
	"io/ioutil"
	"strconv"
	"strings"
)

// LoadStat storing load stat
type LoadStat struct {
	LoadNow      float64 `json:"load_now"`
	LoadPre      float64 `json:"load_pre"`
	LoadFar      float64 `json:"load_far"`
	ProcessTotal uint64  `json:"process_total"`
	ProcessRun   uint64  `json:"process_run"`
	ProcessIdle  uint64  `json:"process_idle"`
}

// GetLoadStat returns load stat
func GetLoadStat() (stat LoadStat, err error) {
	text, err := ioutil.ReadFile("/proc/loadavg")
	if err != nil {
		return
	}

	lines := strings.Split(string(text), "\n")
	fields := strings.Fields(lines[0])

	stat = LoadStat{}
	stat.LoadNow, _ = strconv.ParseFloat(fields[0], strconv.IntSize)
	stat.LoadPre, _ = strconv.ParseFloat(fields[1], strconv.IntSize)
	stat.LoadFar, _ = strconv.ParseFloat(fields[2], strconv.IntSize)

	stat.LoadNow = xhuman.Round(stat.LoadNow, 2)
	stat.LoadPre = xhuman.Round(stat.LoadPre, 2)
	stat.LoadFar = xhuman.Round(stat.LoadFar, 2)

	processes := strings.Split(fields[3], "/")
	stat.ProcessRun, _ = strconv.ParseUint(processes[0], 10, strconv.IntSize)
	stat.ProcessTotal, _ = strconv.ParseUint(processes[1], 10, strconv.IntSize)
	stat.ProcessIdle = stat.ProcessTotal - stat.ProcessRun

	return
}
