package ioctl

var (
	DKIOCGETBLOCKSIZE  = uintptr(0x40046418) // ioctl to get block size
	DKIOCGETBLOCKCOUNT = uintptr(0x40086419) // ioctl to get block count
)
