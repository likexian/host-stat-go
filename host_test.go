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
	"github.com/likexian/gokit/xfile"
	"os"
	"testing"
)

func TestGetRelease(t *testing.T) {
	releaseFiles := []string{
		"/etc/centos-release",
		"/etc/fedora-release",
		"/etc/redhat-release",
		"/etc/SuSE-release",
		"/etc/debian_version",
		"/etc/debian_release",
		"/etc/os-release",
		"/etc/lsb-release",
	}

	for _, v := range releaseFiles {
		os.Rename(v, v+".bak")
	}

	defer func() {
		for _, v := range releaseFiles {
			os.Rename(v+".bak", v)
		}
	}()

	// /etc/lsb-release
	releaseText := `
	DISTRIB_ID=Ubuntu
	DISTRIB_RELEASE=16.04
	DISTRIB_CODENAME=xenial
	DISTRIB_DESCRIPTION="Ubuntu 16.04.3 LTS"
	`
	err := xfile.WriteText("/etc/lsb-release", releaseText)
	assert.Nil(t, err)

	name, err := GetRelease()
	assert.Nil(t, err)
	assert.Equal(t, name, "Ubuntu 16.04")

	// /etc/os-release
	releaseText = `
	PRETTY_NAME="Debian GNU/Linux 8 (jessie)"
	NAME="Debian GNU/Linux"
	VERSION_ID="8"
	VERSION="8 (jessie)"
	ID=debian
	HOME_URL="http://www.debian.org/"
	SUPPORT_URL="http://www.debian.org/support"
	BUG_REPORT_URL="https://bugs.debian.org/"
	`
	err = xfile.WriteText("/etc/os-release", releaseText)
	assert.Nil(t, err)

	name, err = GetRelease()
	assert.Nil(t, err)
	assert.Equal(t, name, "Debian 8")

	// /etc/debian_release
	releaseText = "lenny/sid"
	err = xfile.WriteText("/etc/debian_release", releaseText)
	assert.Nil(t, err)

	name, err = GetRelease()
	assert.Nil(t, err)
	assert.Equal(t, name, "Debian lenny")

	// /etc/debian_version
	releaseText = "9.1"
	err = xfile.WriteText("/etc/debian_version", releaseText)
	assert.Nil(t, err)

	name, err = GetRelease()
	assert.Nil(t, err)
	assert.Equal(t, name, "Debian 9.1")

	// /etc/SuSE-release
	releaseText = `
	openSUSE 13.1 (x86_64)
	VERSION = 13.1
	CODENAME = Bottle
	`
	err = xfile.WriteText("/etc/SuSE-release", releaseText)
	assert.Nil(t, err)

	name, err = GetRelease()
	assert.Nil(t, err)
	assert.Equal(t, name, "openSUSE 13.1")

	// /etc/redhat-release
	releaseText = `Red Hat Enterprise Linux Server release 5.11 (Tikanga)`
	err = xfile.WriteText("/etc/redhat-release", releaseText)
	assert.Nil(t, err)

	name, err = GetRelease()
	assert.Nil(t, err)
	assert.Equal(t, name, "Red Hat Enterprise 5.11")

	// /etc/centos-release
	releaseText = `Fedora release 27 (Twenty Seven)`
	err = xfile.WriteText("/etc/fedora-release", releaseText)
	assert.Nil(t, err)

	name, err = GetRelease()
	assert.Nil(t, err)
	assert.Equal(t, name, "Fedora 27")

	// /etc/centos-release
	releaseText = `CentOS release 7.5 (Final)`
	err = xfile.WriteText("/etc/centos-release", releaseText)
	assert.Nil(t, err)

	name, err = GetRelease()
	assert.Nil(t, err)
	assert.Equal(t, name, "CentOS 7.5")
}
