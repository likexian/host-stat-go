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

// NetStat storing net stat
type NetStat struct {
	Device       string `json:"device"`
	RXBytes      uint64 `json:"rx_bytes"`
	RXPackets    uint64 `json:"rx_packets"`
	RXErrs       uint64 `json:"rx_errs"`
	RXDrop       uint64 `json:"rx_drop"`
	RXFifo       uint64 `json:"rx_fifo"`
	RXFrame      uint64 `json:"rx_frame"`
	RXCompressed uint64 `json:"rx_compressed"`
	RXMulticast  uint64 `json:"rx_multicase"`
	TXBytes      uint64 `json:"tx_bytes"`
	TXPackets    uint64 `json:"tx_packetss"`
	TXErrs       uint64 `json:"tx_errs"`
	TXDrop       uint64 `json:"tx_drop"`
	TXFifo       uint64 `json:"tx_fifo"`
	TXColls      uint64 `json:"tx_colls"`
	TXCarrier    uint64 `json:"tx_carrier"`
	TXCompressed uint64 `json:"tx_compressed"`
}

// GetNetStat returns net stat
func GetNetStat() (stat []NetStat, err error) {
	text, err := ioutil.ReadFile("/proc/net/dev")
	if err != nil {
		return
	}

	stat = []NetStat{}
	lines := strings.Split(string(text), "\n")
	for _, v := range lines {
		if v == "" || !strings.Contains(v, ":") {
			continue
		}

		netStat := NetStat{}
		fields := strings.Split(v, ":")
		netStat.Device = strings.TrimSpace(fields[0])

		fields = strings.Fields("Device " + fields[1])
		netStat.RXBytes, _ = strconv.ParseUint(fields[1], 10, strconv.IntSize)
		netStat.RXPackets, _ = strconv.ParseUint(fields[2], 10, strconv.IntSize)
		netStat.RXErrs, _ = strconv.ParseUint(fields[3], 10, strconv.IntSize)
		netStat.RXDrop, _ = strconv.ParseUint(fields[4], 10, strconv.IntSize)
		netStat.RXFifo, _ = strconv.ParseUint(fields[5], 10, strconv.IntSize)
		netStat.RXFrame, _ = strconv.ParseUint(fields[6], 10, strconv.IntSize)
		netStat.RXCompressed, _ = strconv.ParseUint(fields[7], 10, strconv.IntSize)
		netStat.RXMulticast, _ = strconv.ParseUint(fields[8], 10, strconv.IntSize)
		netStat.TXBytes, _ = strconv.ParseUint(fields[9], 10, strconv.IntSize)
		netStat.TXPackets, _ = strconv.ParseUint(fields[10], 10, strconv.IntSize)
		netStat.TXErrs, _ = strconv.ParseUint(fields[11], 10, strconv.IntSize)
		netStat.TXDrop, _ = strconv.ParseUint(fields[12], 10, strconv.IntSize)
		netStat.TXFifo, _ = strconv.ParseUint(fields[13], 10, strconv.IntSize)
		netStat.TXColls, _ = strconv.ParseUint(fields[14], 10, strconv.IntSize)
		netStat.TXCarrier, _ = strconv.ParseUint(fields[15], 10, strconv.IntSize)
		netStat.TXCompressed, _ = strconv.ParseUint(fields[16], 10, strconv.IntSize)

		stat = append(stat, netStat)
	}

	return
}
