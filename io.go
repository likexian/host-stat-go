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
	"strconv"
	"strings"
)

// IOStat storing io stat
type IOStat struct {
	Major        uint64 `json:"major"`         // major device number
	Minor        uint64 `json:"minor"`         // minor device number
	Name         string `json:"name"`          // device name
	ReadIOs      uint64 `json:"read_ios"`      // r/s total number of reads
	ReadMerges   uint64 `json:"read_merges"`   // rrqm/s number of read I/Os merged with in-queue I/O
	ReadSectors  uint64 `json:"read_sectors"`  // rsec/s number of 512 byte sectors read
	ReadTicks    uint64 `json:"read_ticks"`    // number of milliseconds spent reading
	WriteIOs     uint64 `json:"write_ios"`     // w/s total number of writes
	WriteMerges  uint64 `json:"write_merges"`  // wrqm/s number of write I/Os merged with in-queue I/O
	WriteSectors uint64 `json:"write_sectors"` // wsec/s number of 512 byte sectors written
	WriteTicks   uint64 `json:"write_ticks"`   // number of milliseconds spent writing
	InFlight     uint64 `json:"in_flight"`     // number of I/Os currently in progress
	IOTicks      uint64 `json:"io_ticks"`      // number of milliseconds spent doing I/Os(in flight)
	TimeInQueue  uint64 `json:"time_in_queue"` // total wait time for all requests in milliseconds
	ReadBytes    uint64 `json:"read_bytes"`    // rB/s
	WriteBytes   uint64 `json:"write_bytes"`   // wB/s
}

// GetIOStat returns io stat
func GetIOStat() (stat []IOStat, err error) {
	text, err := ioutil.ReadFile("/proc/diskstats")
	if err != nil {
		return
	}

	stat = []IOStat{}
	lines := strings.Split(string(text), "\n")
	for _, v := range lines {
		if v == "" {
			continue
		}

		fields := strings.Fields(v)
		if fields[3] == "0" {
			continue
		}

		ioStat := IOStat{}
		ioStat.Major, _ = strconv.ParseUint(fields[0], 10, strconv.IntSize)
		ioStat.Minor, _ = strconv.ParseUint(fields[1], 10, strconv.IntSize)
		ioStat.Name = fields[2]
		ioStat.ReadIOs, _ = strconv.ParseUint(fields[3], 10, strconv.IntSize)
		ioStat.ReadMerges, _ = strconv.ParseUint(fields[4], 10, strconv.IntSize)
		ioStat.ReadSectors, _ = strconv.ParseUint(fields[5], 10, strconv.IntSize)
		ioStat.ReadTicks, _ = strconv.ParseUint(fields[6], 10, strconv.IntSize)
		ioStat.WriteIOs, _ = strconv.ParseUint(fields[7], 10, strconv.IntSize)
		ioStat.WriteMerges, _ = strconv.ParseUint(fields[8], 10, strconv.IntSize)
		ioStat.WriteSectors, _ = strconv.ParseUint(fields[9], 10, strconv.IntSize)
		ioStat.WriteTicks, _ = strconv.ParseUint(fields[10], 10, strconv.IntSize)
		ioStat.InFlight, _ = strconv.ParseUint(fields[11], 10, strconv.IntSize)
		ioStat.IOTicks, _ = strconv.ParseUint(fields[12], 10, strconv.IntSize)
		ioStat.TimeInQueue, _ = strconv.ParseUint(fields[13], 10, strconv.IntSize)

		ioStat.ReadBytes = ioStat.ReadSectors * 512
		ioStat.WriteBytes = ioStat.WriteSectors * 512

		stat = append(stat, ioStat)
	}

	return
}
