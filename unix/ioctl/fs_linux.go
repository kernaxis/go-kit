package ioctl

import (
	"golang.org/x/sys/unix"
)

var (
	BLKGETSIZE   = uintptr(unix.BLKGETSIZE)
	BLKGETSIZE64 = uintptr(unix.BLKGETSIZE64)
)
