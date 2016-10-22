package termios

import (
	"syscall"
)

import (
	"github.com/sternix/lib/termios/ioctl"
)

func Isatty(fd uintptr) bool {
	var t syscall.Termios
	return (Tcgetattr(fd,&t) == nil)
}

func GetWinsize(fd uintptr,ws *Winsize) error {
	return ioctl.Winsize(fd,TIOCGWINSZ,ws)
}

func Tcgetattr(fd uintptr, termios *syscall.Termios) error {
	return ioctl.Termios(fd, TIOCGETA, termios)
}

func Tcsetattr(fd uintptr, opt int, termios *syscall.Termios) error {
	if (opt & TCSASOFT) == 1 {
		termios.Cflag |= CIGNORE
	}

	switch opt & ^TCSASOFT {
	case TCSANOW:
		return ioctl.Termios(fd, TIOCSETA, termios)
	case TCSADRAIN:
		return ioctl.Termios(fd, TIOCSETAW, termios)
	case TCSAFLUSH:
		return ioctl.Termios(fd, TIOCSETAF, termios)
	default:
		return syscall.EINVAL
	}
}

func Tcsetpgrp(fd uintptr, pid int) error {
	return ioctl.Int(fd, TIOCSPGRP, &pid)
}

func Tcgetpgrp(fd uintptr) (int, error) {
	var s int
	err := ioctl.Int(fd, TIOCGPGRP, &s)
	if err != nil {
		return -1, err
	}
	return s, nil
}

func Tcgetsid(fd uintptr) (int, error) {
	var s int
	err := ioctl.Int(fd, TIOCGSID, &s)
	if err != nil {
		return -1, err
	}
	return s, nil
}

func Tcsetsid(fd uintptr, pid int) error {
	sid, _ := syscall.Getsid(0)
	if pid != sid {
		return syscall.EINVAL
	}
	return ioctl.Null(fd, TIOCSCTTY)
}

func Cfgetospeed(termios *syscall.Termios) uint32 {
	return termios.Ospeed
}

func Cfgetispeed(termios *syscall.Termios) uint32 {
	return termios.Ispeed
}

func Cfsetospeed(termios *syscall.Termios, speed uint32) {
	termios.Ospeed = speed
}

func Cfsetispeed(termios *syscall.Termios, speed uint32) {
	termios.Ispeed = speed
}

func Cfsetspeed(termios *syscall.Termios, speed uint32) {
	termios.Ispeed, termios.Ospeed = speed, speed
}

func Cfmakeraw(termios *syscall.Termios) {
	termios.Iflag &^= (IMAXBEL | IXOFF | INPCK | BRKINT | PARMRK | ISTRIP | INLCR | IGNCR | ICRNL | IXON | IGNPAR)
	termios.Iflag |= IGNBRK
	termios.Oflag &^= OPOST
	termios.Lflag &^= (ECHO | ECHOE | ECHOK | ECHONL | ICANON | ISIG | IEXTEN | NOFLSH | TOSTOP | PENDIN)
	termios.Cflag &^= (CSIZE | PARENB)
	termios.Cflag |= CS8 | CREAD
	termios.Cc[VMIN] = 1
	termios.Cc[VTIME] = 0
}

func Cfmakesane(termios *syscall.Termios) {
	termios.Cflag = TTYDEF_CFLAG
	termios.Iflag = TTYDEF_IFLAG
	termios.Lflag = TTYDEF_LFLAG
	termios.Oflag = TTYDEF_OFLAG
	termios.Ispeed = TTYDEF_SPEED
	termios.Ospeed = TTYDEF_SPEED

	copy(termios.Cc[:], Ttydefchars)
}

func Tcsendbreak(fd uintptr) error {
	var sleepyTime syscall.Timeval
	sleepyTime.Sec = 0
	sleepyTime.Usec = 400000
	err := ioctl.Zero(fd, TIOCSBRK)
	if err != nil {
		return err
	}
	syscall.Select(0, nil, nil, nil, &sleepyTime)
	err = ioctl.Zero(fd, TIOCCBRK)
	if err != nil {
		return err
	}
	return nil
}

func Tcdrain(fd uintptr) error {
	return ioctl.Zero(fd, TIOCDRAIN)
}

func Tcflush(fd uintptr, which int) error {
	var com int

	switch which {
	case TCIFLUSH:
		com = FREAD
	case TCOFLUSH:
		com = FWRITE
	case TCIOFLUSH:
		com = FREAD | FWRITE
	default:
		return syscall.EINVAL
	}
	return ioctl.Int(fd, TIOCFLUSH, &com)
}

func Tcflow(fd uintptr, action int) error {
	var (
		term syscall.Termios
		c    uint8
	)

	switch action {
	case TCOOFF:
		return ioctl.Zero(fd, TIOCSTOP)
	case TCOON:
		return ioctl.Zero(fd, TIOCSTART)
	case TCION:
	case TCIOFF:
		if err := Tcgetattr(fd, &term); err != nil {
			return err
		}

		if action == TCIOFF {
			c = term.Cc[VSTOP]
		} else {
			c = term.Cc[VSTART]
		}

		if c != _POSIX_VDISABLE {
			bslc := []byte{c}
			n, err := syscall.Write(int(fd), bslc)
			if err != nil {
				return err
			}
			if n != 0 {
				return syscall.EINVAL
			}
		}
		return nil
	}
	return syscall.EINVAL
}
