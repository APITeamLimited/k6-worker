// Copyright 2009,2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Darwin system calls.
// This file is compiled as ordinary Go code,
// but it is also input to mksyscall,
// which parses the //sys lines and generates system call stubs.
// Note that sometimes we use a lowercase //sys name and wrap
// it in our own nicer implementation, either here or in
// syscall_bsd.go or syscall_unix.go.

package unix

import (
	"fmt"
	"runtime"
	"syscall"
	"unsafe"
)

// SockaddrDatalink implements the Sockaddr interface for AF_LINK type sockets.
type SockaddrDatalink struct ***REMOVED***
	Len    uint8
	Family uint8
	Index  uint16
	Type   uint8
	Nlen   uint8
	Alen   uint8
	Slen   uint8
	Data   [12]int8
	raw    RawSockaddrDatalink
***REMOVED***

// SockaddrCtl implements the Sockaddr interface for AF_SYSTEM type sockets.
type SockaddrCtl struct ***REMOVED***
	ID   uint32
	Unit uint32
	raw  RawSockaddrCtl
***REMOVED***

func (sa *SockaddrCtl) sockaddr() (unsafe.Pointer, _Socklen, error) ***REMOVED***
	sa.raw.Sc_len = SizeofSockaddrCtl
	sa.raw.Sc_family = AF_SYSTEM
	sa.raw.Ss_sysaddr = AF_SYS_CONTROL
	sa.raw.Sc_id = sa.ID
	sa.raw.Sc_unit = sa.Unit
	return unsafe.Pointer(&sa.raw), SizeofSockaddrCtl, nil
***REMOVED***

// SockaddrVM implements the Sockaddr interface for AF_VSOCK type sockets.
// SockaddrVM provides access to Darwin VM sockets: a mechanism that enables
// bidirectional communication between a hypervisor and its guest virtual
// machines.
type SockaddrVM struct ***REMOVED***
	// CID and Port specify a context ID and port address for a VM socket.
	// Guests have a unique CID, and hosts may have a well-known CID of:
	//  - VMADDR_CID_HYPERVISOR: refers to the hypervisor process.
	//  - VMADDR_CID_LOCAL: refers to local communication (loopback).
	//  - VMADDR_CID_HOST: refers to other processes on the host.
	CID  uint32
	Port uint32
	raw  RawSockaddrVM
***REMOVED***

func (sa *SockaddrVM) sockaddr() (unsafe.Pointer, _Socklen, error) ***REMOVED***
	sa.raw.Len = SizeofSockaddrVM
	sa.raw.Family = AF_VSOCK
	sa.raw.Port = sa.Port
	sa.raw.Cid = sa.CID

	return unsafe.Pointer(&sa.raw), SizeofSockaddrVM, nil
***REMOVED***

func anyToSockaddrGOOS(fd int, rsa *RawSockaddrAny) (Sockaddr, error) ***REMOVED***
	switch rsa.Addr.Family ***REMOVED***
	case AF_SYSTEM:
		pp := (*RawSockaddrCtl)(unsafe.Pointer(rsa))
		if pp.Ss_sysaddr == AF_SYS_CONTROL ***REMOVED***
			sa := new(SockaddrCtl)
			sa.ID = pp.Sc_id
			sa.Unit = pp.Sc_unit
			return sa, nil
		***REMOVED***
	case AF_VSOCK:
		pp := (*RawSockaddrVM)(unsafe.Pointer(rsa))
		sa := &SockaddrVM***REMOVED***
			CID:  pp.Cid,
			Port: pp.Port,
		***REMOVED***
		return sa, nil
	***REMOVED***
	return nil, EAFNOSUPPORT
***REMOVED***

// Some external packages rely on SYS___SYSCTL being defined to implement their
// own sysctl wrappers. Provide it here, even though direct syscalls are no
// longer supported on darwin.
const SYS___SYSCTL = SYS_SYSCTL

// Translate "kern.hostname" to []_C_int***REMOVED***0,1,2,3***REMOVED***.
func nametomib(name string) (mib []_C_int, err error) ***REMOVED***
	const siz = unsafe.Sizeof(mib[0])

	// NOTE(rsc): It seems strange to set the buffer to have
	// size CTL_MAXNAME+2 but use only CTL_MAXNAME
	// as the size. I don't know why the +2 is here, but the
	// kernel uses +2 for its own implementation of this function.
	// I am scared that if we don't include the +2 here, the kernel
	// will silently write 2 words farther than we specify
	// and we'll get memory corruption.
	var buf [CTL_MAXNAME + 2]_C_int
	n := uintptr(CTL_MAXNAME) * siz

	p := (*byte)(unsafe.Pointer(&buf[0]))
	bytes, err := ByteSliceFromString(name)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***

	// Magic sysctl: "setting" 0.3 to a string name
	// lets you read back the array of integers form.
	if err = sysctl([]_C_int***REMOVED***0, 3***REMOVED***, p, &n, &bytes[0], uintptr(len(name))); err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	return buf[0 : n/siz], nil
***REMOVED***

func direntIno(buf []byte) (uint64, bool) ***REMOVED***
	return readInt(buf, unsafe.Offsetof(Dirent***REMOVED******REMOVED***.Ino), unsafe.Sizeof(Dirent***REMOVED******REMOVED***.Ino))
***REMOVED***

func direntReclen(buf []byte) (uint64, bool) ***REMOVED***
	return readInt(buf, unsafe.Offsetof(Dirent***REMOVED******REMOVED***.Reclen), unsafe.Sizeof(Dirent***REMOVED******REMOVED***.Reclen))
***REMOVED***

func direntNamlen(buf []byte) (uint64, bool) ***REMOVED***
	return readInt(buf, unsafe.Offsetof(Dirent***REMOVED******REMOVED***.Namlen), unsafe.Sizeof(Dirent***REMOVED******REMOVED***.Namlen))
***REMOVED***

func PtraceAttach(pid int) (err error) ***REMOVED*** return ptrace(PT_ATTACH, pid, 0, 0) ***REMOVED***
func PtraceDetach(pid int) (err error) ***REMOVED*** return ptrace(PT_DETACH, pid, 0, 0) ***REMOVED***

type attrList struct ***REMOVED***
	bitmapCount uint16
	_           uint16
	CommonAttr  uint32
	VolAttr     uint32
	DirAttr     uint32
	FileAttr    uint32
	Forkattr    uint32
***REMOVED***

//sysnb	pipe(p *[2]int32) (err error)

func Pipe(p []int) (err error) ***REMOVED***
	if len(p) != 2 ***REMOVED***
		return EINVAL
	***REMOVED***
	var x [2]int32
	err = pipe(&x)
	if err == nil ***REMOVED***
		p[0] = int(x[0])
		p[1] = int(x[1])
	***REMOVED***
	return
***REMOVED***

func Getfsstat(buf []Statfs_t, flags int) (n int, err error) ***REMOVED***
	var _p0 unsafe.Pointer
	var bufsize uintptr
	if len(buf) > 0 ***REMOVED***
		_p0 = unsafe.Pointer(&buf[0])
		bufsize = unsafe.Sizeof(Statfs_t***REMOVED******REMOVED***) * uintptr(len(buf))
	***REMOVED***
	return getfsstat(_p0, bufsize, flags)
***REMOVED***

func xattrPointer(dest []byte) *byte ***REMOVED***
	// It's only when dest is set to NULL that the OS X implementations of
	// getxattr() and listxattr() return the current sizes of the named attributes.
	// An empty byte array is not sufficient. To maintain the same behaviour as the
	// linux implementation, we wrap around the system calls and pass in NULL when
	// dest is empty.
	var destp *byte
	if len(dest) > 0 ***REMOVED***
		destp = &dest[0]
	***REMOVED***
	return destp
***REMOVED***

//sys	getxattr(path string, attr string, dest *byte, size int, position uint32, options int) (sz int, err error)

func Getxattr(path string, attr string, dest []byte) (sz int, err error) ***REMOVED***
	return getxattr(path, attr, xattrPointer(dest), len(dest), 0, 0)
***REMOVED***

func Lgetxattr(link string, attr string, dest []byte) (sz int, err error) ***REMOVED***
	return getxattr(link, attr, xattrPointer(dest), len(dest), 0, XATTR_NOFOLLOW)
***REMOVED***

//sys	fgetxattr(fd int, attr string, dest *byte, size int, position uint32, options int) (sz int, err error)

func Fgetxattr(fd int, attr string, dest []byte) (sz int, err error) ***REMOVED***
	return fgetxattr(fd, attr, xattrPointer(dest), len(dest), 0, 0)
***REMOVED***

//sys	setxattr(path string, attr string, data *byte, size int, position uint32, options int) (err error)

func Setxattr(path string, attr string, data []byte, flags int) (err error) ***REMOVED***
	// The parameters for the OS X implementation vary slightly compared to the
	// linux system call, specifically the position parameter:
	//
	//  linux:
	//      int setxattr(
	//          const char *path,
	//          const char *name,
	//          const void *value,
	//          size_t size,
	//          int flags
	//      );
	//
	//  darwin:
	//      int setxattr(
	//          const char *path,
	//          const char *name,
	//          void *value,
	//          size_t size,
	//          u_int32_t position,
	//          int options
	//      );
	//
	// position specifies the offset within the extended attribute. In the
	// current implementation, only the resource fork extended attribute makes
	// use of this argument. For all others, position is reserved. We simply
	// default to setting it to zero.
	return setxattr(path, attr, xattrPointer(data), len(data), 0, flags)
***REMOVED***

func Lsetxattr(link string, attr string, data []byte, flags int) (err error) ***REMOVED***
	return setxattr(link, attr, xattrPointer(data), len(data), 0, flags|XATTR_NOFOLLOW)
***REMOVED***

//sys	fsetxattr(fd int, attr string, data *byte, size int, position uint32, options int) (err error)

func Fsetxattr(fd int, attr string, data []byte, flags int) (err error) ***REMOVED***
	return fsetxattr(fd, attr, xattrPointer(data), len(data), 0, 0)
***REMOVED***

//sys	removexattr(path string, attr string, options int) (err error)

func Removexattr(path string, attr string) (err error) ***REMOVED***
	// We wrap around and explicitly zero out the options provided to the OS X
	// implementation of removexattr, we do so for interoperability with the
	// linux variant.
	return removexattr(path, attr, 0)
***REMOVED***

func Lremovexattr(link string, attr string) (err error) ***REMOVED***
	return removexattr(link, attr, XATTR_NOFOLLOW)
***REMOVED***

//sys	fremovexattr(fd int, attr string, options int) (err error)

func Fremovexattr(fd int, attr string) (err error) ***REMOVED***
	return fremovexattr(fd, attr, 0)
***REMOVED***

//sys	listxattr(path string, dest *byte, size int, options int) (sz int, err error)

func Listxattr(path string, dest []byte) (sz int, err error) ***REMOVED***
	return listxattr(path, xattrPointer(dest), len(dest), 0)
***REMOVED***

func Llistxattr(link string, dest []byte) (sz int, err error) ***REMOVED***
	return listxattr(link, xattrPointer(dest), len(dest), XATTR_NOFOLLOW)
***REMOVED***

//sys	flistxattr(fd int, dest *byte, size int, options int) (sz int, err error)

func Flistxattr(fd int, dest []byte) (sz int, err error) ***REMOVED***
	return flistxattr(fd, xattrPointer(dest), len(dest), 0)
***REMOVED***

func setattrlistTimes(path string, times []Timespec, flags int) error ***REMOVED***
	_p0, err := BytePtrFromString(path)
	if err != nil ***REMOVED***
		return err
	***REMOVED***

	var attrList attrList
	attrList.bitmapCount = ATTR_BIT_MAP_COUNT
	attrList.CommonAttr = ATTR_CMN_MODTIME | ATTR_CMN_ACCTIME

	// order is mtime, atime: the opposite of Chtimes
	attributes := [2]Timespec***REMOVED***times[1], times[0]***REMOVED***
	options := 0
	if flags&AT_SYMLINK_NOFOLLOW != 0 ***REMOVED***
		options |= FSOPT_NOFOLLOW
	***REMOVED***
	return setattrlist(
		_p0,
		unsafe.Pointer(&attrList),
		unsafe.Pointer(&attributes),
		unsafe.Sizeof(attributes),
		options)
***REMOVED***

//sys	setattrlist(path *byte, list unsafe.Pointer, buf unsafe.Pointer, size uintptr, options int) (err error)

func utimensat(dirfd int, path string, times *[2]Timespec, flags int) error ***REMOVED***
	// Darwin doesn't support SYS_UTIMENSAT
	return ENOSYS
***REMOVED***

/*
 * Wrapped
 */

//sys	fcntl(fd int, cmd int, arg int) (val int, err error)

//sys	kill(pid int, signum int, posix int) (err error)

func Kill(pid int, signum syscall.Signal) (err error) ***REMOVED*** return kill(pid, int(signum), 1) ***REMOVED***

//sys	ioctl(fd int, req uint, arg uintptr) (err error)

func IoctlCtlInfo(fd int, ctlInfo *CtlInfo) error ***REMOVED***
	err := ioctl(fd, CTLIOCGINFO, uintptr(unsafe.Pointer(ctlInfo)))
	runtime.KeepAlive(ctlInfo)
	return err
***REMOVED***

// IfreqMTU is struct ifreq used to get or set a network device's MTU.
type IfreqMTU struct ***REMOVED***
	Name [IFNAMSIZ]byte
	MTU  int32
***REMOVED***

// IoctlGetIfreqMTU performs the SIOCGIFMTU ioctl operation on fd to get the MTU
// of the network device specified by ifname.
func IoctlGetIfreqMTU(fd int, ifname string) (*IfreqMTU, error) ***REMOVED***
	var ifreq IfreqMTU
	copy(ifreq.Name[:], ifname)
	err := ioctl(fd, SIOCGIFMTU, uintptr(unsafe.Pointer(&ifreq)))
	return &ifreq, err
***REMOVED***

// IoctlSetIfreqMTU performs the SIOCSIFMTU ioctl operation on fd to set the MTU
// of the network device specified by ifreq.Name.
func IoctlSetIfreqMTU(fd int, ifreq *IfreqMTU) error ***REMOVED***
	err := ioctl(fd, SIOCSIFMTU, uintptr(unsafe.Pointer(ifreq)))
	runtime.KeepAlive(ifreq)
	return err
***REMOVED***

//sys	sysctl(mib []_C_int, old *byte, oldlen *uintptr, new *byte, newlen uintptr) (err error) = SYS_SYSCTL

func Uname(uname *Utsname) error ***REMOVED***
	mib := []_C_int***REMOVED***CTL_KERN, KERN_OSTYPE***REMOVED***
	n := unsafe.Sizeof(uname.Sysname)
	if err := sysctl(mib, &uname.Sysname[0], &n, nil, 0); err != nil ***REMOVED***
		return err
	***REMOVED***

	mib = []_C_int***REMOVED***CTL_KERN, KERN_HOSTNAME***REMOVED***
	n = unsafe.Sizeof(uname.Nodename)
	if err := sysctl(mib, &uname.Nodename[0], &n, nil, 0); err != nil ***REMOVED***
		return err
	***REMOVED***

	mib = []_C_int***REMOVED***CTL_KERN, KERN_OSRELEASE***REMOVED***
	n = unsafe.Sizeof(uname.Release)
	if err := sysctl(mib, &uname.Release[0], &n, nil, 0); err != nil ***REMOVED***
		return err
	***REMOVED***

	mib = []_C_int***REMOVED***CTL_KERN, KERN_VERSION***REMOVED***
	n = unsafe.Sizeof(uname.Version)
	if err := sysctl(mib, &uname.Version[0], &n, nil, 0); err != nil ***REMOVED***
		return err
	***REMOVED***

	// The version might have newlines or tabs in it, convert them to
	// spaces.
	for i, b := range uname.Version ***REMOVED***
		if b == '\n' || b == '\t' ***REMOVED***
			if i == len(uname.Version)-1 ***REMOVED***
				uname.Version[i] = 0
			***REMOVED*** else ***REMOVED***
				uname.Version[i] = ' '
			***REMOVED***
		***REMOVED***
	***REMOVED***

	mib = []_C_int***REMOVED***CTL_HW, HW_MACHINE***REMOVED***
	n = unsafe.Sizeof(uname.Machine)
	if err := sysctl(mib, &uname.Machine[0], &n, nil, 0); err != nil ***REMOVED***
		return err
	***REMOVED***

	return nil
***REMOVED***

func Sendfile(outfd int, infd int, offset *int64, count int) (written int, err error) ***REMOVED***
	if raceenabled ***REMOVED***
		raceReleaseMerge(unsafe.Pointer(&ioSync))
	***REMOVED***
	var length = int64(count)
	err = sendfile(infd, outfd, *offset, &length, nil, 0)
	written = int(length)
	return
***REMOVED***

func GetsockoptIPMreqn(fd, level, opt int) (*IPMreqn, error) ***REMOVED***
	var value IPMreqn
	vallen := _Socklen(SizeofIPMreqn)
	errno := getsockopt(fd, level, opt, unsafe.Pointer(&value), &vallen)
	return &value, errno
***REMOVED***

func SetsockoptIPMreqn(fd, level, opt int, mreq *IPMreqn) (err error) ***REMOVED***
	return setsockopt(fd, level, opt, unsafe.Pointer(mreq), unsafe.Sizeof(*mreq))
***REMOVED***

// GetsockoptXucred is a getsockopt wrapper that returns an Xucred struct.
// The usual level and opt are SOL_LOCAL and LOCAL_PEERCRED, respectively.
func GetsockoptXucred(fd, level, opt int) (*Xucred, error) ***REMOVED***
	x := new(Xucred)
	vallen := _Socklen(SizeofXucred)
	err := getsockopt(fd, level, opt, unsafe.Pointer(x), &vallen)
	return x, err
***REMOVED***

func SysctlKinfoProc(name string, args ...int) (*KinfoProc, error) ***REMOVED***
	mib, err := sysctlmib(name, args...)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***

	var kinfo KinfoProc
	n := uintptr(SizeofKinfoProc)
	if err := sysctl(mib, (*byte)(unsafe.Pointer(&kinfo)), &n, nil, 0); err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	if n != SizeofKinfoProc ***REMOVED***
		return nil, EIO
	***REMOVED***
	return &kinfo, nil
***REMOVED***

func SysctlKinfoProcSlice(name string, args ...int) ([]KinfoProc, error) ***REMOVED***
	mib, err := sysctlmib(name, args...)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***

	// Find size.
	n := uintptr(0)
	if err := sysctl(mib, nil, &n, nil, 0); err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	if n == 0 ***REMOVED***
		return nil, nil
	***REMOVED***
	if n%SizeofKinfoProc != 0 ***REMOVED***
		return nil, fmt.Errorf("sysctl() returned a size of %d, which is not a multiple of %d", n, SizeofKinfoProc)
	***REMOVED***

	// Read into buffer of that size.
	buf := make([]KinfoProc, n/SizeofKinfoProc)
	if err := sysctl(mib, (*byte)(unsafe.Pointer(&buf[0])), &n, nil, 0); err != nil ***REMOVED***
		return nil, err
	***REMOVED***
	if n%SizeofKinfoProc != 0 ***REMOVED***
		return nil, fmt.Errorf("sysctl() returned a size of %d, which is not a multiple of %d", n, SizeofKinfoProc)
	***REMOVED***

	// The actual call may return less than the original reported required
	// size so ensure we deal with that.
	return buf[:n/SizeofKinfoProc], nil
***REMOVED***

//sys	sendfile(infd int, outfd int, offset int64, len *int64, hdtr unsafe.Pointer, flags int) (err error)

//sys	shmat(id int, addr uintptr, flag int) (ret uintptr, err error)
//sys	shmctl(id int, cmd int, buf *SysvShmDesc) (result int, err error)
//sys	shmdt(addr uintptr) (err error)
//sys	shmget(key int, size int, flag int) (id int, err error)

/*
 * Exposed directly
 */
//sys	Access(path string, mode uint32) (err error)
//sys	Adjtime(delta *Timeval, olddelta *Timeval) (err error)
//sys	Chdir(path string) (err error)
//sys	Chflags(path string, flags int) (err error)
//sys	Chmod(path string, mode uint32) (err error)
//sys	Chown(path string, uid int, gid int) (err error)
//sys	Chroot(path string) (err error)
//sys	ClockGettime(clockid int32, time *Timespec) (err error)
//sys	Close(fd int) (err error)
//sys	Clonefile(src string, dst string, flags int) (err error)
//sys	Clonefileat(srcDirfd int, src string, dstDirfd int, dst string, flags int) (err error)
//sys	Dup(fd int) (nfd int, err error)
//sys	Dup2(from int, to int) (err error)
//sys	Exchangedata(path1 string, path2 string, options int) (err error)
//sys	Exit(code int)
//sys	Faccessat(dirfd int, path string, mode uint32, flags int) (err error)
//sys	Fchdir(fd int) (err error)
//sys	Fchflags(fd int, flags int) (err error)
//sys	Fchmod(fd int, mode uint32) (err error)
//sys	Fchmodat(dirfd int, path string, mode uint32, flags int) (err error)
//sys	Fchown(fd int, uid int, gid int) (err error)
//sys	Fchownat(dirfd int, path string, uid int, gid int, flags int) (err error)
//sys	Fclonefileat(srcDirfd int, dstDirfd int, dst string, flags int) (err error)
//sys	Flock(fd int, how int) (err error)
//sys	Fpathconf(fd int, name int) (val int, err error)
//sys	Fsync(fd int) (err error)
//sys	Ftruncate(fd int, length int64) (err error)
//sys	Getcwd(buf []byte) (n int, err error)
//sys	Getdtablesize() (size int)
//sysnb	Getegid() (egid int)
//sysnb	Geteuid() (uid int)
//sysnb	Getgid() (gid int)
//sysnb	Getpgid(pid int) (pgid int, err error)
//sysnb	Getpgrp() (pgrp int)
//sysnb	Getpid() (pid int)
//sysnb	Getppid() (ppid int)
//sys	Getpriority(which int, who int) (prio int, err error)
//sysnb	Getrlimit(which int, lim *Rlimit) (err error)
//sysnb	Getrusage(who int, rusage *Rusage) (err error)
//sysnb	Getsid(pid int) (sid int, err error)
//sysnb	Gettimeofday(tp *Timeval) (err error)
//sysnb	Getuid() (uid int)
//sysnb	Issetugid() (tainted bool)
//sys	Kqueue() (fd int, err error)
//sys	Lchown(path string, uid int, gid int) (err error)
//sys	Link(path string, link string) (err error)
//sys	Linkat(pathfd int, path string, linkfd int, link string, flags int) (err error)
//sys	Listen(s int, backlog int) (err error)
//sys	Mkdir(path string, mode uint32) (err error)
//sys	Mkdirat(dirfd int, path string, mode uint32) (err error)
//sys	Mkfifo(path string, mode uint32) (err error)
//sys	Mknod(path string, mode uint32, dev int) (err error)
//sys	Open(path string, mode int, perm uint32) (fd int, err error)
//sys	Openat(dirfd int, path string, mode int, perm uint32) (fd int, err error)
//sys	Pathconf(path string, name int) (val int, err error)
//sys	Pread(fd int, p []byte, offset int64) (n int, err error)
//sys	Pwrite(fd int, p []byte, offset int64) (n int, err error)
//sys	read(fd int, p []byte) (n int, err error)
//sys	Readlink(path string, buf []byte) (n int, err error)
//sys	Readlinkat(dirfd int, path string, buf []byte) (n int, err error)
//sys	Rename(from string, to string) (err error)
//sys	Renameat(fromfd int, from string, tofd int, to string) (err error)
//sys	Revoke(path string) (err error)
//sys	Rmdir(path string) (err error)
//sys	Seek(fd int, offset int64, whence int) (newoffset int64, err error) = SYS_LSEEK
//sys	Select(nfd int, r *FdSet, w *FdSet, e *FdSet, timeout *Timeval) (n int, err error)
//sys	Setegid(egid int) (err error)
//sysnb	Seteuid(euid int) (err error)
//sysnb	Setgid(gid int) (err error)
//sys	Setlogin(name string) (err error)
//sysnb	Setpgid(pid int, pgid int) (err error)
//sys	Setpriority(which int, who int, prio int) (err error)
//sys	Setprivexec(flag int) (err error)
//sysnb	Setregid(rgid int, egid int) (err error)
//sysnb	Setreuid(ruid int, euid int) (err error)
//sysnb	Setrlimit(which int, lim *Rlimit) (err error)
//sysnb	Setsid() (pid int, err error)
//sysnb	Settimeofday(tp *Timeval) (err error)
//sysnb	Setuid(uid int) (err error)
//sys	Symlink(path string, link string) (err error)
//sys	Symlinkat(oldpath string, newdirfd int, newpath string) (err error)
//sys	Sync() (err error)
//sys	Truncate(path string, length int64) (err error)
//sys	Umask(newmask int) (oldmask int)
//sys	Undelete(path string) (err error)
//sys	Unlink(path string) (err error)
//sys	Unlinkat(dirfd int, path string, flags int) (err error)
//sys	Unmount(path string, flags int) (err error)
//sys	write(fd int, p []byte) (n int, err error)
//sys	mmap(addr uintptr, length uintptr, prot int, flag int, fd int, pos int64) (ret uintptr, err error)
//sys	munmap(addr uintptr, length uintptr) (err error)
//sys	readlen(fd int, buf *byte, nbuf int) (n int, err error) = SYS_READ
//sys	writelen(fd int, buf *byte, nbuf int) (n int, err error) = SYS_WRITE

/*
 * Unimplemented
 */
// Profil
// Sigaction
// Sigprocmask
// Getlogin
// Sigpending
// Sigaltstack
// Ioctl
// Reboot
// Execve
// Vfork
// Sbrk
// Sstk
// Ovadvise
// Mincore
// Setitimer
// Swapon
// Select
// Sigsuspend
// Readv
// Writev
// Nfssvc
// Getfh
// Quotactl
// Mount
// Csops
// Waitid
// Add_profil
// Kdebug_trace
// Sigreturn
// Atsocket
// Kqueue_from_portset_np
// Kqueue_portset
// Getattrlist
// Setattrlist
// Getdirentriesattr
// Searchfs
// Delete
// Copyfile
// Watchevent
// Waitevent
// Modwatch
// Fsctl
// Initgroups
// Posix_spawn
// Nfsclnt
// Fhopen
// Minherit
// Semsys
// Msgsys
// Shmsys
// Semctl
// Semget
// Semop
// Msgctl
// Msgget
// Msgsnd
// Msgrcv
// Shm_open
// Shm_unlink
// Sem_open
// Sem_close
// Sem_unlink
// Sem_wait
// Sem_trywait
// Sem_post
// Sem_getvalue
// Sem_init
// Sem_destroy
// Open_extended
// Umask_extended
// Stat_extended
// Lstat_extended
// Fstat_extended
// Chmod_extended
// Fchmod_extended
// Access_extended
// Settid
// Gettid
// Setsgroups
// Getsgroups
// Setwgroups
// Getwgroups
// Mkfifo_extended
// Mkdir_extended
// Identitysvc
// Shared_region_check_np
// Shared_region_map_np
// __pthread_mutex_destroy
// __pthread_mutex_init
// __pthread_mutex_lock
// __pthread_mutex_trylock
// __pthread_mutex_unlock
// __pthread_cond_init
// __pthread_cond_destroy
// __pthread_cond_broadcast
// __pthread_cond_signal
// Setsid_with_pid
// __pthread_cond_timedwait
// Aio_fsync
// Aio_return
// Aio_suspend
// Aio_cancel
// Aio_error
// Aio_read
// Aio_write
// Lio_listio
// __pthread_cond_wait
// Iopolicysys
// __pthread_kill
// __pthread_sigmask
// __sigwait
// __disable_threadsignal
// __pthread_markcancel
// __pthread_canceled
// __semwait_signal
// Proc_info
// sendfile
// Stat64_extended
// Lstat64_extended
// Fstat64_extended
// __pthread_chdir
// __pthread_fchdir
// Audit
// Auditon
// Getauid
// Setauid
// Getaudit
// Setaudit
// Getaudit_addr
// Setaudit_addr
// Auditctl
// Bsdthread_create
// Bsdthread_terminate
// Stack_snapshot
// Bsdthread_register
// Workq_open
// Workq_ops
// __mac_execve
// __mac_syscall
// __mac_get_file
// __mac_set_file
// __mac_get_link
// __mac_set_link
// __mac_get_proc
// __mac_set_proc
// __mac_get_fd
// __mac_set_fd
// __mac_get_pid
// __mac_get_lcid
// __mac_get_lctx
// __mac_set_lctx
// Setlcid
// Read_nocancel
// Write_nocancel
// Open_nocancel
// Close_nocancel
// Wait4_nocancel
// Recvmsg_nocancel
// Sendmsg_nocancel
// Recvfrom_nocancel
// Accept_nocancel
// Fcntl_nocancel
// Select_nocancel
// Fsync_nocancel
// Connect_nocancel
// Sigsuspend_nocancel
// Readv_nocancel
// Writev_nocancel
// Sendto_nocancel
// Pread_nocancel
// Pwrite_nocancel
// Waitid_nocancel
// Poll_nocancel
// Msgsnd_nocancel
// Msgrcv_nocancel
// Sem_wait_nocancel
// Aio_suspend_nocancel
// __sigwait_nocancel
// __semwait_signal_nocancel
// __mac_mount
// __mac_get_mount
// __mac_getfsstat
