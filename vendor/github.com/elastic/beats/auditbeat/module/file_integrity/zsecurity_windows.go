// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// MACHINE GENERATED BY 'go generate' COMMAND; DO NOT EDIT

package file_integrity

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

var _ unsafe.Pointer

var (
	modadvapi32 = windows.NewLazySystemDLL("advapi32.dll")

	procGetSecurityInfo = modadvapi32.NewProc("GetSecurityInfo")
)

func GetSecurityInfo(handle syscall.Handle, objectType ObjectType, securityInformation SecurityInformation, ppsidOwner **syscall.SID, ppsidGroup **syscall.SID, ppDacl **ACL, ppSacl **ACL, ppSecurityDescriptor **SecurityDescriptor) (err error) {
	r1, _, e1 := syscall.Syscall9(procGetSecurityInfo.Addr(), 8, uintptr(handle), uintptr(objectType), uintptr(securityInformation), uintptr(unsafe.Pointer(ppsidOwner)), uintptr(unsafe.Pointer(ppsidGroup)), uintptr(unsafe.Pointer(ppDacl)), uintptr(unsafe.Pointer(ppSacl)), uintptr(unsafe.Pointer(ppSecurityDescriptor)), 0)
	if r1 != 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}
