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
	"strings"
	"syscall"
)

// DiskStat storing disk stat
type DiskStat struct {
	Label    string  `json:"label"`
	Mount    string  `json:"mount"`
	Total    uint64  `json:"total"`
	Used     uint64  `json:"used"`
	Free     uint64  `json:"free"`
	UsedRate float64 `json:"used_rate"`
}

// GetDiskStat returns all disk stat
func GetDiskStat() (stat []DiskStat, err error) {
	text, err := ioutil.ReadFile("/etc/mtab")
	if err != nil {
		return
	}

	stat = []DiskStat{}
	lines := strings.Split(string(text), "\n")
	for i := 0; i < len(lines); i++ {
		lines[i] = strings.TrimSpace(lines[i])
		if lines[i] == "" {
			continue
		}

		fields := strings.Fields(lines[i])
		if len(fields) < 3 {
			continue
		}

		if fields[2] != "proc" && fields[2] != "cgroup" && fields[2] != "sysfs" && fields[2] != "devpts" && fields[2] != "mqueue" {
			diskStat, _ := getStat(fields[1])
			diskStat.Label = fields[0]
			stat = append(stat, diskStat)
		}
	}

	return
}

// getStat returns the path stat
func getStat(path string) (stat DiskStat, err error) {
	fs := syscall.Statfs_t{}
	err = syscall.Statfs(path, &fs)
	if err != nil {
		return
	}

	stat = DiskStat{}
	stat.Mount = path
	stat.Total = fs.Blocks * uint64(fs.Bsize) / (1024 * 1024)
	stat.Free = fs.Bfree * uint64(fs.Bsize) / (1024 * 1024)
	stat.Used = stat.Total - stat.Free
	stat.UsedRate = xhuman.Round(float64(stat.Used)*100/float64(stat.Total), 2)

	return
}
