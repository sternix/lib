package termios

import (
	"syscall"
	"unsafe"
)

//https://github.com/freebsd/freebsd/blob/master/sys/sys/ioccom.h

const (
	IOCPARM_SHIFT = 13
	IOCPARM_MASK  = ((1 << IOCPARM_SHIFT) - 1) // parameter length mask

	IOCPARM_MAX = (1 << IOCPARM_SHIFT) // max size of ioctl
	IOC_VOID    = 0x20000000           // no parameters
	IOC_OUT     = 0x40000000           // copy out parameters
	IOC_IN      = 0x80000000           // copy in parameters
	IOC_INOUT   = (IOC_IN | IOC_OUT)
	IOC_DIRMASK = (IOC_VOID | IOC_OUT | IOC_IN)
)

func Len(x int) int {
	return (((x) >> 16) & IOCPARM_SHIFT)
}

// ((x) & ~(IOCPARM_MASK << 16))
func BaseCmd(x int) int {
	return ((x) & ^(IOCPARM_MASK << 16))
}

func Group(x int) int {
	return (((x) >> 8) & 0xff)
}

func IOC(inout int, group byte, num int, size int) uint64 {
	return uint64((inout) | (((size) & IOCPARM_MASK) << 16) | (int(group) << 8) | (num))
}

func IO(group byte, num int) uint64 {
	return IOC(IOC_VOID, group, num, 0)
}

func IOWINT(group byte, num int) uint64 {
	var a int
	return IOC(IOC_VOID, group, num, int(unsafe.Sizeof(a)))
}

func IOR(group byte, num int, size uintptr) uint64 {
	return IOC(IOC_OUT, group, num, int(size))
}

func IOW(group byte, num int, size uintptr) uint64 {
	return IOC(IOC_IN, group, num, int(size))
}

func IOWR(group byte, num int, size uintptr) uint64 {
	return IOC(IOC_INOUT, group, num, int(size))
}

func UIOCCMD(num int) uint64 {
	return IO('u', num)
}

//https://github.com/freebsd/freebsd/blob/master/sys/sys/ttycom.h
var (
	TIOCEXCL       uint64
	TIOCNXCL       uint64
	TIOCGPTN       uint64
	TIOCFLUSH      uint64
	TIOCGETA       uint64
	TIOCSETA       uint64
	TIOCSETAW      uint64
	TIOCSETAF      uint64
	TIOCGETD       uint64
	TIOCSETD       uint64
	TIOCPTMASTER   uint64
	TIOCGDRAINWAIT uint64
	TIOCSDRAINWAIT uint64
	TIOCTIMESTAMP  uint64
	TIOCMGDTRWAIT  uint64
	TIOCMSDTRWAIT  uint64
	TIOCDRAIN      uint64
	TIOCSIG        uint64
	TIOCEXT        uint64
	TIOCSCTTY      uint64
	TIOCCONS       uint64
	TIOCGSID       uint64
	TIOCSTAT       uint64
	TIOCUCNTL      uint64
	TIOCSWINSZ     uint64
	TIOCGWINSZ     uint64
	TIOCMGET       uint64
	TIOCMBIC       uint64
	TIOCMBIS       uint64
	TIOCMSET       uint64
	TIOCSTART      uint64
	TIOCSTOP       uint64
	TIOCPKT        uint64
	TIOCNOTTY      uint64
	TIOCSTI        uint64
	TIOCOUTQ       uint64
	TIOCSPGRP      uint64
	TIOCGPGRP      uint64
	TIOCCDTR       uint64
	TIOCSDTR       uint64
	TIOCCBRK       uint64
	TIOCSBRK       uint64
)

const (
	TIOCM_LE  = 0001 // line enable
	TIOCM_DTR = 0002 // data terminal ready
	TIOCM_RTS = 0004 // request to send
	TIOCM_ST  = 0010 // secondary transmit
	TIOCM_SR  = 0020 // secondary receive
	TIOCM_CTS = 0040 // clear to send
	TIOCM_DCD = 0100 // data carrier detect
	TIOCM_RI  = 0200 // ring indicate
	TIOCM_DSR = 0400 // data set ready
	TIOCM_CD  = TIOCM_DCD
	TIOCM_CAR = TIOCM_DCD
	TIOCM_RNG = TIOCM_RI

	TIOCPKT_DATA       = 0x00 // data packet
	TIOCPKT_FLUSHREAD  = 0x01 // flush packet
	TIOCPKT_FLUSHWRITE = 0x02 // flush packet
	TIOCPKT_STOP       = 0x04 // stop output
	TIOCPKT_START      = 0x08 // start output
	TIOCPKT_NOSTOP     = 0x10 // no more ^S, ^Q
	TIOCPKT_DOSTOP     = 0x20 // now do ^S ^Q
	TIOCPKT_IOCTL      = 0x40 // state change of pty driver

	TTYDISC      = 0 // termios tty line discipline
	SLIPDISC     = 4 // serial IP discipline
	PPPDISC      = 5 // PPP discipline
	NETGRAPHDISC = 6 // Netgraph tty node discipline
	H4DISC       = 7 // Netgraph Bluetooth H4 discipline
)

func init() {
	var (
		i       int
		termios syscall.Termios
		c       byte
		w       Winsize
		tv      syscall.Timeval
	)

	si := unsafe.Sizeof(i)
	st := unsafe.Sizeof(termios)
	sc := unsafe.Sizeof(c)
	sw := unsafe.Sizeof(w)
	stv := unsafe.Sizeof(tv)

	TIOCEXCL = IO('t', 13)            // set exclusive use of tty
	TIOCNXCL = IO('t', 14)            // reset exclusive use of tty
	TIOCGPTN = IOR('t', 15, si)       // Get pts number.
	TIOCFLUSH = IOW('t', 16, si)      // flush buffers
	TIOCGETA = IOR('t', 19, st)       // get termios struct
	TIOCSETA = IOW('t', 20, st)       // set termios struct
	TIOCSETAW = IOW('t', 21, st)      // drain output, set
	TIOCSETAF = IOW('t', 22, st)      // drn out, fls in, set
	TIOCGETD = IOR('t', 26, si)       // get line discipline
	TIOCSETD = IOW('t', 27, si)       // set line discipline
	TIOCPTMASTER = IO('t', 28)        // pts master validation
	TIOCGDRAINWAIT = IOR('t', 86, si) // get ttywait timeout
	TIOCSDRAINWAIT = IOW('t', 87, si) // set ttywait timeout
	TIOCTIMESTAMP = IOR('t', 89, stv) // enable/get timestamp
	TIOCMGDTRWAIT = IOR('t', 90, si)  // modem: get wait on close
	TIOCMSDTRWAIT = IOW('t', 91, si)  // modem: set wait on close
	TIOCDRAIN = IO('t', 94)           // wait till output drained
	TIOCSIG = IOWINT('t', 95)         // pty: generate signal
	TIOCEXT = IOW('t', 96, si)        // pty: external processing
	TIOCSCTTY = IO('t', 97)           // become controlling tty
	TIOCCONS = IOW('t', 98, si)       // become virtual console
	TIOCGSID = IOR('t', 99, si)       // get session id
	TIOCSTAT = IO('t', 101)           // simulate ^T status message
	TIOCUCNTL = IOW('t', 102, si)     // pty: set/clr usr cntl mode
	TIOCSWINSZ = IOW('t', 103, sw)    // set window size
	TIOCGWINSZ = IOR('t', 104, sw)    // get window size
	TIOCMGET = IOR('t', 106, si)      // get all modem bits
	TIOCMBIC = IOW('t', 107, si)      // bic modem bits
	TIOCMBIS = IOW('t', 108, si)      // bis modem bits
	TIOCMSET = IOW('t', 109, si)      // set all modem bits
	TIOCSTART = IO('t', 110)          // start output, like ^Q
	TIOCSTOP = IO('t', 111)           // stop output, like ^S
	TIOCPKT = IOW('t', 112, si)       // pty: set/clear packet mode
	TIOCNOTTY = IO('t', 113)          // void tty association
	TIOCSTI = IOW('t', 114, sc)       // simulate terminal input
	TIOCOUTQ = IOR('t', 115, si)      // output queue size
	TIOCSPGRP = IOW('t', 118, si)     // set pgrp of tty
	TIOCGPGRP = IOR('t', 119, si)     // get pgrp of tty
	TIOCCDTR = IO('t', 120)           // clear data terminal ready
	TIOCSDTR = IO('t', 121)           // set data terminal ready
	TIOCCBRK = IO('t', 122)           // clear break bit
	TIOCSBRK = IO('t', 123)           // set break bit
}
