package termios

//https://github.com/freebsd/freebsd/blob/master/sys/sys/ttydefaults.h

var (
	TTYDEF_IFLAG        uint32 = (BRKINT | ICRNL | IMAXBEL | IXON | IXANY)
	TTYDEF_OFLAG        uint32 = (OPOST | ONLCR)
	TTYDEF_LFLAG_NOECHO uint32 = (ICANON | ISIG | IEXTEN)
	TTYDEF_LFLAG_ECHO   uint32 = (TTYDEF_LFLAG_NOECHO | ECHO | ECHOE | ECHOKE | ECHOCTL)
	TTYDEF_LFLAG        uint32 = TTYDEF_LFLAG_ECHO
	TTYDEF_CFLAG        uint32 = (CREAD | CS8 | HUPCL)
	TTYDEF_SPEED        uint32 = (B9600)

	Ttydefchars []uint8 = []uint8{
		CEOF, CEOL, CEOL, CERASE, CWERASE, CKILL, CREPRINT, CERASE2, CINTR,
		CQUIT, CSUSP, CDSUSP, CSTART, CSTOP, CLNEXT, CDISCARD, CMIN, CTIME,
		CSTATUS, _POSIX_VDISABLE,
	}
)

func CTRL(x uint8) uint8 {
	if x >= 'a' && x <= 'z' {
		return (x - 'a' + 1)
	} else {
		return ((x - 'A' + 1) & 0x7f)
	}
}

var (
	CEOF     uint8
	CERASE   uint8
	CERASE2  uint8
	CINTR    uint8
	CSTATUS  uint8
	CKILL    uint8
	CQUIT    uint8
	CSUSP    uint8
	CDSUSP   uint8
	CSTART   uint8
	CSTOP    uint8
	CLNEXT   uint8
	CDISCARD uint8
	CWERASE  uint8
	CREPRINT uint8
	CEOT     uint8
	CBRK     uint8
	CRPRNT   uint8
	CFLUSH   uint8
)

const (
	CEOL  uint8 = 0xff
	CMIN  uint8 = 1
	CTIME uint8 = 0
)

func init() {

	CEOF = CTRL('D')
	CERASE = CTRL('?')
	CERASE2 = CTRL('H')
	CINTR = CTRL('C')
	CSTATUS = CTRL('T')
	CKILL = CTRL('U')
	CQUIT = CTRL('\\')
	CSUSP = CTRL('Z')
	CDSUSP = CTRL('Y')
	CSTART = CTRL('Q')
	CSTOP = CTRL('S')
	CLNEXT = CTRL('V')
	CDISCARD = CTRL('O')
	CWERASE = CTRL('W')
	CREPRINT = CTRL('R')
	CEOT = CEOF
	CBRK = CEOL
	CRPRNT = CREPRINT
	CFLUSH = CDISCARD
}
