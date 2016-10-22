package ioctl

import (
	"syscall"
	"unsafe"
)

func Termios(fd uintptr, request uint64, termios *syscall.Termios) error {
	_, _, err := syscall.Syscall(syscall.SYS_IOCTL, fd, uintptr(request), uintptr(unsafe.Pointer(termios)))
	if err != 0 {
		return err
	}
	return nil
}

func Int(fd uintptr, req uint64, i *int) error {
	_, _, err := syscall.Syscall(syscall.SYS_IOCTL, fd, uintptr(req), uintptr(unsafe.Pointer(i)))
	if err != 0 {
		return err
	}
	return nil
}

func Null(fd uintptr, req uint64) error {
	_, _, err := syscall.Syscall(syscall.SYS_IOCTL, fd, uintptr(req), uintptr(unsafe.Pointer(nil)))
	if err != 0 {
		return err
	}
	return nil
}

func Zero(fd uintptr, req uint64) error {
	_, _, err := syscall.Syscall(syscall.SYS_IOCTL, fd, uintptr(req), uintptr(0))
	if err != 0 {
		return err
	}
	return nil
}
