package termios

import (
	"syscall"
	"testing"
)

func TestEquality(t *testing.T) {

	if TIOCEXCL != syscall.TIOCEXCL {
		t.Logf("TIOCEXCL != syscall.TIOCEXCL")
	}

	if TIOCNXCL != syscall.TIOCNXCL {
		t.Logf("TIOCNXCL != syscall.TIOCNXCL")
	}

	if TIOCGPTN != syscall.TIOCGPTN {
		t.Logf("TIOCGPTN != syscall.TIOCGPTN")
	}

	if TIOCFLUSH != syscall.TIOCFLUSH {
		t.Logf("TIOCFLUSH != syscall.TIOCFLUSH")
	}

	if TIOCGETA != syscall.TIOCGETA {
		t.Logf("TIOCGETA != syscall.TIOCGETA")
	}

	if TIOCSETA != syscall.TIOCSETA {
		t.Logf("TIOCSETA != syscall.TIOCSETA")
	}

	if TIOCSETAW != syscall.TIOCSETAW {
		t.Logf("TIOCSETAW != syscall.TIOCSETAW")
	}

	if TIOCSETAF != syscall.TIOCSETAF {
		t.Logf("TIOCSETAF != syscall.TIOCSETAF")
	}

	if TIOCGETD != syscall.TIOCGETD {
		t.Logf("TIOCGETD != syscall.TIOCGETD")
	}

	if TIOCSETD != syscall.TIOCSETD {
		t.Logf("TIOCSETD != syscall.TIOCSETD")
	}

	if TIOCPTMASTER != syscall.TIOCPTMASTER {
		t.Logf("TIOCPTMASTER != syscall.TIOCPTMASTER")
	}

	if TIOCGDRAINWAIT != syscall.TIOCGDRAINWAIT {
		t.Logf("TIOCGDRAINWAIT != syscall.TIOCGDRAINWAIT")
	}

	if TIOCSDRAINWAIT != syscall.TIOCSDRAINWAIT {
		t.Logf("TIOCSDRAINWAIT != syscall.TIOCSDRAINWAIT")
	}

	if TIOCTIMESTAMP != syscall.TIOCTIMESTAMP {
		t.Logf("TIOCTIMESTAMP != syscall.TIOCTIMESTAMP")
	}

	if TIOCMGDTRWAIT != syscall.TIOCMGDTRWAIT {
		t.Logf("TIOCMGDTRWAIT != syscall.TIOCMGDTRWAIT")
	}

	if TIOCMSDTRWAIT != syscall.TIOCMSDTRWAIT {
		t.Logf("TIOCMSDTRWAIT != syscall.TIOCMSDTRWAIT")
	}

	if TIOCDRAIN != syscall.TIOCDRAIN {
		t.Logf("TIOCDRAIN != syscall.TIOCDRAIN")
	}

	if TIOCSIG != syscall.TIOCSIG {
		t.Logf("TIOCSIG != syscall.TIOCSIG")
	}

	if TIOCEXT != syscall.TIOCEXT {
		t.Logf("TIOCEXT != syscall.TIOCEXT")
	}

	if TIOCSCTTY != syscall.TIOCSCTTY {
		t.Logf("TIOCSCTTY != syscall.TIOCSCTTY")
	}

	if TIOCCONS != syscall.TIOCCONS {
		t.Logf("TIOCCONS != syscall.TIOCCONS")
	}

	if TIOCGSID != syscall.TIOCGSID {
		t.Logf("TIOCGSID != syscall.TIOCGSID")
	}

	if TIOCSTAT != syscall.TIOCSTAT {
		t.Logf("TIOCSTAT != syscall.TIOCSTAT")
	}

	if TIOCUCNTL != syscall.TIOCUCNTL {
		t.Logf("TIOCUCNTL != syscall.TIOCUCNTL")
	}

	if TIOCSWINSZ != syscall.TIOCSWINSZ {
		t.Logf("TIOCSWINSZ != syscall.TIOCSWINSZ")
	}

	if TIOCGWINSZ != syscall.TIOCGWINSZ {
		t.Logf("TIOCGWINSZ != syscall.TIOCGWINSZ")
	}

	if TIOCMGET != syscall.TIOCMGET {
		t.Logf("TIOCMGET != syscall.TIOCMGET")
	}

	if TIOCMBIC != syscall.TIOCMBIC {
		t.Logf("TIOCMBIC != syscall.TIOCMBIC")
	}

	if TIOCMBIS != syscall.TIOCMBIS {
		t.Logf("TIOCMBIS != syscall.TIOCMBIS")
	}

	if TIOCMSET != syscall.TIOCMSET {
		t.Logf("TIOCMSET != syscall.TIOCMSET")
	}

	if TIOCSTART != syscall.TIOCSTART {
		t.Logf("TIOCSTART != syscall.TIOCSTART")
	}

	if TIOCSTOP != syscall.TIOCSTOP {
		t.Logf("TIOCSTOP != syscall.TIOCSTOP")
	}

	if TIOCPKT != syscall.TIOCPKT {
		t.Logf("TIOCPKT != syscall.TIOCPKT")
	}

	if TIOCNOTTY != syscall.TIOCNOTTY {
		t.Logf("TIOCNOTTY != syscall.TIOCNOTTY")
	}

	if TIOCSTI != syscall.TIOCSTI {
		t.Logf("TIOCSTI != syscall.TIOCSTI")
	}

	if TIOCOUTQ != syscall.TIOCOUTQ {
		t.Logf("TIOCOUTQ != syscall.TIOCOUTQ")
	}

	if TIOCSPGRP != syscall.TIOCSPGRP {
		t.Logf("TIOCSPGRP != syscall.TIOCSPGRP")
	}

	if TIOCGPGRP != syscall.TIOCGPGRP {
		t.Logf("TIOCGPGRP != syscall.TIOCGPGRP")
	}

	if TIOCCDTR != syscall.TIOCCDTR {
		t.Logf("TIOCCDTR != syscall.TIOCCDTR")
	}

	if TIOCSDTR != syscall.TIOCSDTR {
		t.Logf("TIOCSDTR != syscall.TIOCSDTR")
	}

	if TIOCCBRK != syscall.TIOCCBRK {
		t.Logf("TIOCCBRK != syscall.TIOCCBRK")
	}

	if TIOCSBRK != syscall.TIOCSBRK {
		t.Logf("TIOCSBRK != syscall.TIOCSBRK")
	}

}
