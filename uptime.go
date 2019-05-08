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
	"github.com/likexian/gokit/xfile"
	"github.com/likexian/gokit/xhuman"
	"strconv"
	"strings"
)

// UptimeStat storing uptime stat
type UptimeStat struct {
	Uptime   float64 `json:"uptime"`
	IdleTime float64 `json:"idle_time"`
	IdleRate float64 `json:"idle_rate"`
}

// GetUptimeStat returns uptime stat
func GetUptimeStat() (stat UptimeStat, err error) {
	text, err := xfile.ReadFirstLine("/proc/uptime")
	if err != nil {
		return
	}

	lines := strings.Split(text, "\n")
	fields := strings.Fields(lines[0])

	stat = UptimeStat{}
	stat.Uptime, _ = strconv.ParseFloat(fields[0], strconv.IntSize)
	stat.IdleTime, _ = strconv.ParseFloat(fields[1], strconv.IntSize)
	stat.IdleRate = xhuman.Round(stat.IdleTime*100/stat.Uptime, 2)

	stat.Uptime = xhuman.Round(stat.Uptime, 2)
	stat.IdleTime = xhuman.Round(stat.IdleTime, 2)

	return
}
