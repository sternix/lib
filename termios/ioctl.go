package termios

import (
	"syscall"
	//	"unsafe"
)

//int ioctl(int fd, unsigned long request, ...);

// ioctl for tty
func ioctl(fd uintptr, request uint64, ptr uintptr) error {
	_, _, err := syscall.Syscall(syscall.SYS_IOCTL, fd, uintptr(request), ptr)
	if err != 0 {
		return err
	}
	return nil
}

/*
func ioctlTermios(fd uintptr, request uint64, termios *syscall.Termios) error {
	_, _, err := syscall.Syscall(syscall.SYS_IOCTL, fd, uintptr(request), uintptr(unsafe.Pointer(termios)))
	if err != 0 {
		return err
	}
	return nil
}

func ioctlInt(fd uintptr, request uint64, p *int) error {
	_, _, err := syscall.Syscall(syscall.SYS_IOCTL, fd, uintptr(request), uintptr(pid))
	if err != 0 {
		return err
	}
	return nil
}
*/

/*
func ioctl(fd uintptr, request uint64, termios *syscall.Termios) error {
	_, _, err := syscall.Syscall(syscall.SYS_IOCTL, fd, uintptr(request), uintptr(unsafe.Pointer(termios)))
	if err != 0 {
		return err
	}
	return nil
}

func pidIoctl(fd uintptr, request uint64, pid int) error {
	_, _, err := syscall.Syscall(syscall.SYS_IOCTL, fd, uintptr(request), uintptr(pid))
	if err != 0 {
		return err
	}
	return nil
}
*/
