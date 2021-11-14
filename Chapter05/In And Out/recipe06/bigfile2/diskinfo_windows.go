package main

import (
	"errors"
	"os"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

const ERROR_SUCCESS syscall.Errno = 0

func diskInfo() (total, free int64, err error) {
	var wd string
	wd, err = os.Getwd()
	if err != nil {
		return
	}

	h := windows.MustLoadDLL("kernel32.dll")
	c := h.MustFindProc("GetDiskFreeSpaceExW")

	// https://pkg.go.dev/golang.org/x/sys/windows?utm_source=godoc#GetDiskFreeSpaceEx
	_, _, err = c.Call(uintptr(unsafe.Pointer(windows.StringToUTF16Ptr(wd))),
		uintptr(unsafe.Pointer(&free)), uintptr(unsafe.Pointer(&total)), 0)

	var lastErr syscall.Errno
	if errors.As(err, &lastErr) && lastErr == ERROR_SUCCESS {
		err = nil
	}
	return
}
