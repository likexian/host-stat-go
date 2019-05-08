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

// CPUInfo storing cpu info
type CPUInfo struct {
	ModelName string `json:"model_name"`
	CoreCount uint64 `json:"core_count"`
}

// CPUStat storing cpu stat
type CPUStat struct {
	User       uint64  `json:"user"`
	Nice       uint64  `json:"nice"`
	System     uint64  `json:"system"`
	Idle       uint64  `json:"idle"`
	IOWait     uint64  `json:"io_wait"`
	IRQ        uint64  `json:"irq"`
	SoftIRQS   uint64  `json:"soft_irqs"`
	Steal      uint64  `json:"steal"`
	Guest      uint64  `json:"guest"`
	GuestNice  uint64  `json:"guest_nice"`
	UserRate   float64 `json:"user_rate"`
	SystemRate float64 `json:"system_rate"`
	IdleRate   float64 `json:"idle_rate"`
	IOWaitRate float64 `json:"io_wait_rate"`
}

// GetCPUInfo returns cpu info
func GetCPUInfo() (info CPUInfo, err error) {
	text, err := ioutil.ReadFile("/proc/cpuinfo")
	if err != nil {
		return
	}

	info = CPUInfo{}
	info.CoreCount = 0
	lines := strings.Split(string(text), "\n")
	for _, l := range lines {
		if !strings.Contains(l, ":") {
			continue
		}

		maps := strings.Split(l, ":")
		maps[0] = strings.TrimSpace(maps[0])
		maps[1] = strings.TrimSpace(maps[1])

		if maps[0] == "model name" {
			if info.ModelName == "" {
				info.ModelName = strings.Join(strings.Fields(maps[1]), " ")
			}
			info.CoreCount += 1
		}
	}

	return
}

// GetCPUStat returns cpu stat
func GetCPUStat() (stat CPUStat, err error) {
	text, err := ioutil.ReadFile("/proc/stat")
	if err != nil {
		return
	}

	lines := strings.Split(string(text), "\n")
	fields := strings.Fields(lines[0])

	stat = CPUStat{}
	stat.User, _ = strconv.ParseUint(fields[1], 10, strconv.IntSize)
	stat.Nice, _ = strconv.ParseUint(fields[2], 10, strconv.IntSize)
	stat.System, _ = strconv.ParseUint(fields[3], 10, strconv.IntSize)
	stat.Idle, _ = strconv.ParseUint(fields[4], 10, strconv.IntSize)
	if len(fields) > 5 { // 2.5.41+
		stat.IOWait, _ = strconv.ParseUint(fields[5], 10, strconv.IntSize)
	}
	if len(fields) > 6 { // 2.6.0-test4+
		stat.IRQ, _ = strconv.ParseUint(fields[6], 10, strconv.IntSize)
	}
	if len(fields) > 7 { // 2.6.0-test4+
		stat.SoftIRQS, _ = strconv.ParseUint(fields[7], 10, strconv.IntSize)
	}
	if len(fields) > 8 { // 2.6.11+
		stat.Steal, _ = strconv.ParseUint(fields[8], 10, strconv.IntSize)
	}
	if len(fields) > 9 { // 2.6.24+
		stat.Guest, _ = strconv.ParseUint(fields[9], 10, strconv.IntSize)
	}
	if len(fields) > 10 { // 2.6.33+
		stat.GuestNice, _ = strconv.ParseUint(fields[10], 10, strconv.IntSize)
	}

	total := stat.User + stat.Nice + stat.System + stat.Idle + stat.IOWait +
		stat.IRQ + stat.SoftIRQS + stat.Steal + stat.Guest + stat.GuestNice
	stat.UserRate = xhuman.Round(float64(stat.User+stat.Nice)*100/float64(total), 2)
	stat.SystemRate = xhuman.Round(float64(stat.System+stat.IRQ+stat.SoftIRQS)*100/float64(total), 2)
	stat.IdleRate = xhuman.Round(float64(stat.Idle)*100/float64(total), 2)
	stat.IOWaitRate = xhuman.Round(float64(stat.IOWait)*100/float64(total), 2)

	return
}
