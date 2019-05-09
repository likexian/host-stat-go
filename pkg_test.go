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
	"github.com/likexian/gokit/assert"
	"os"
	"testing"
)

func TestHostStat(t *testing.T) {
	hostInfo, err := GetHostInfo()
	assert.Nil(t, err)
	t.Log(hostInfo)

	cpuInfo, err := GetCPUInfo()
	assert.Nil(t, err)
	t.Log(cpuInfo)

	cpuStat, err := GetCPUStat()
	assert.Nil(t, err)
	t.Log(cpuStat)

	memStat, err := GetMemStat()
	assert.Nil(t, err)
	t.Log(memStat)

	diskStat, err := GetDiskStat()
	assert.Nil(t, err)
	for _, v := range diskStat {
		t.Log(v)
	}

	ioStat, err := GetIOStat()
	if err != nil {
		if e, ok := err.(*os.PathError); ok {
			t.Log(e)
		} else {
			assert.Nil(t, err)
		}
	}
	t.Log(ioStat)

	netStat, err := GetNetStat()
	assert.Nil(t, err)
	t.Log(netStat)

	uptimeStat, err := GetUptimeStat()
	assert.Nil(t, err)
	t.Log(uptimeStat)

	loadStat, err := GetLoadStat()
	assert.Nil(t, err)
	t.Log(loadStat)
}
