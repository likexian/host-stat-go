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
	"regexp"
	"strconv"
	"strings"
)

// HostInfo storing host info
type HostInfo struct {
	HostName  string `json:"host_name"`
	OSType    string `json:"os_type"`
	OSRelease string `json:"os_release"`
	Version   string `json:"version"`
	Release   string `json:"release"`
	OSBit     string `json:"os_bit"`
}

// GetHostInfo returns host info
func GetHostInfo() (info HostInfo, err error) {
	info.HostName, err = xfile.ReadFirstLine("/proc/sys/kernel/hostname")
	if err != nil {
		return
	}

	info.OSType, err = xfile.ReadFirstLine("/proc/sys/kernel/ostype")
	if err != nil {
		return
	}

	info.OSRelease, err = xfile.ReadFirstLine("/proc/sys/kernel/osrelease")
	if err != nil {
		return
	}

	info.OSBit = ""
	if len(info.OSRelease) > 6 {
		if info.OSRelease[len(info.OSRelease)-6:] == "x86_64" {
			info.OSBit = "64Bit"
		} else if info.OSRelease[len(info.OSRelease)-4:] == "i686" {
			info.OSBit = "32Bit"
		}
	}

	if info.OSBit == "" {
		info.OSBit = strconv.Itoa(32<<uintptr(^uintptr(0)>>63)) + "Bit"
	}

	info.Version, err = xfile.ReadFirstLine("/proc/sys/kernel/version")
	if err != nil {
		return
	}

	info.Release, err = GetRelease()
	if err != nil {
		return
	}

	return
}

// GetRelease returns host release info
func GetRelease() (name string, err error) {
	text := ""
	if xfile.Exists("/etc/centos-release") {
		text, err = xfile.ReadFirstLine("/etc/centos-release")
		if err != nil {
			return
		}
		name = text
	} else if xfile.Exists("/etc/fedora-release") {
		text, err = xfile.ReadFirstLine("/etc/fedora-release")
		if err != nil {
			return
		}
		name = text
	} else if xfile.Exists("/etc/redhat-release") {
		text, err = xfile.ReadFirstLine("/etc/redhat-release")
		if err != nil {
			return
		}
		name = text
	} else if xfile.Exists("/etc/SuSE-release") {
		text, err = xfile.ReadFirstLine("/etc/SuSE-release")
		if err != nil {
			return
		}
		name = text
	} else if xfile.Exists("/etc/debian_version") {
		text, err = xfile.ReadFirstLine("/etc/debian_version")
		if err != nil {
			return
		}
		name = "Debian " + text
	} else if xfile.Exists("/etc/debian_release") {
		text, err = xfile.ReadFirstLine("/etc/debian_release")
		if err != nil {
			return
		}
		name = "Debian " + text
	} else if xfile.Exists("/etc/os-release") {
		text, err = xfile.ReadText("/etc/os-release")
		if err != nil {
			return
		}

		lines := strings.Split(text, "\n")
		for _, l := range lines {
			l = strings.TrimSpace(l)
			if !strings.Contains(l, "=") {
				continue
			}

			ls := strings.Split(l, "=")
			ls[1] = strings.Trim(ls[1], "\"")
			if strings.Contains(ls[1], ",") {
				ls[1] = strings.Split(ls[1], ",")[0]
			}

			if ls[0] == "NAME" {
				name = ls[1] // Fedora
			} else if ls[0] == "VERSION" {
				name = name + " " + ls[1] // "17 (Beefy Miracle)"
				break
			} else if ls[0] == "PRETTY_NAME" {
				name = ls[1] // Debian GNU/Linux wheezy/sid
				break
			}
		}
	} else if xfile.Exists("/etc/lsb-release") {
		text, err = xfile.ReadText("/etc/lsb-release")
		if err != nil {
			return
		}

		lines := strings.Split(text, "\n")
		for _, l := range lines {
			l = strings.TrimSpace(l)
			if !strings.Contains(l, "=") {
				continue
			}

			ls := strings.Split(l, "=")
			if ls[0] == "DISTRIB_ID" {
				name = ls[1] // Ubuntu
			} else if ls[0] == "DISTRIB_RELEASE" {
				name = name + " " + ls[1] // 9.04
				break
			}
		}
	}

	if strings.Contains(name, "(") {
		names := strings.Split(name, "(")
		name = names[0]
	}

	name = strings.Replace(name, "GNU/", "", -1)
	if strings.Contains(name, "/") {
		names := strings.Split(name, "/")
		name = strings.Join(names[:len(names)-1], "/")
	}

	naRe := regexp.MustCompile(`(?i)(gnu|linux|server|release)`)
	name = naRe.ReplaceAllString(name, "")

	spRe := regexp.MustCompile(`\s+`)
	name = spRe.ReplaceAllString(name, " ")

	name = strings.TrimSpace(name)

	return
}
