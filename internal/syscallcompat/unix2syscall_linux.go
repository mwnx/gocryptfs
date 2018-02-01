package syscallcompat

import (
	"syscall"

	"golang.org/x/sys/unix"
)

// Unix2syscall converts a unix.Stat_t struct to a syscall.Stat_t struct.
// A direct cast does not work because the padding is named differently in
// unix.Stat_t for some reason ("X__unused" in syscall, "_" in unix).
func Unix2syscall(u unix.Stat_t) syscall.Stat_t {
	s := syscall.Stat_t{
		Dev:     u.Dev,
		Ino:     u.Ino,
		Nlink:   u.Nlink,
		Mode:    u.Mode,
		Uid:     u.Uid,
		Gid:     u.Gid,
		Rdev:    u.Rdev,
		Size:    u.Size,
		Blksize: u.Blksize,
		Blocks:  u.Blocks,
		// Casting unix.Timespec to syscall.Timespec does not work on gccgo:
		// > cannot use type int64 as type syscall.Timespec_sec_t
		// So do it manually, element-wise.
		//
		//Atim:   syscall.Timespec(u.Atim),
		//Mtim:   syscall.Timespec(u.Mtim),
		//Ctim:   syscall.Timespec(u.Ctim),
	}
	s.Atim.Sec = u.Atim.Sec
	s.Atim.Nsec = u.Atim.Nsec

	s.Mtim.Sec = u.Mtim.Sec
	s.Mtim.Nsec = u.Mtim.Nsec

	s.Ctim.Sec = u.Ctim.Sec
	s.Ctim.Nsec = u.Ctim.Nsec

	return s
}
