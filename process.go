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
	"strconv"
	"strings"
)

// ProcessStat storing process stat
type ProcessStat struct {
	Pid   int    `json:"pid"`
	Name  string `json:"name"`
	State string `json:"state"`
}

// StateMap is state map
var StateMap = map[string]string{
	"R": "running",
	"S": "sleeping",
	"D": "disk sleep",
	"T": "stopped",
	"Z": "zombie",
	"X": "dead",
}

// GetProcessStat returns process stat
func GetProcessStat() (stat []ProcessStat, err error) {
	ls, err := xfile.ListDir("/proc/", xfile.TypeDir, -1)
	if err != nil {
		return
	}

	for _, f := range ls {
		pid, err := strconv.Atoi(f.Name)
		if err != nil {
			continue
		}
		text, err := xfile.ReadFirstLine(f.Path + "/stat")
		if err != nil {
			continue
		}
		fields := strings.Fields(text)
		if len(fields) > 3 {
			name := strings.TrimSpace(strings.Trim(strings.Trim(fields[1], "("), ")"))
			state := fields[2]
			if _, ok := StateMap[state]; ok {
				state = StateMap[state]
			}
			stat = append(stat, ProcessStat{pid, name, state})
		}
	}

	return
}
