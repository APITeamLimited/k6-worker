// mkerrors.sh -Wall -Werror -static -I/tmp/include
// Code generated by the command above; see README.md. DO NOT EDIT.

// +build mips64,linux

// Code generated by cmd/cgo -godefs; DO NOT EDIT.
// cgo -godefs -- -Wall -Werror -static -I/tmp/include _const.go

package unix

import "syscall"

const (
	B1000000                         = 0x1008
	B115200                          = 0x1002
	B1152000                         = 0x1009
	B1500000                         = 0x100a
	B2000000                         = 0x100b
	B230400                          = 0x1003
	B2500000                         = 0x100c
	B3000000                         = 0x100d
	B3500000                         = 0x100e
	B4000000                         = 0x100f
	B460800                          = 0x1004
	B500000                          = 0x1005
	B57600                           = 0x1001
	B576000                          = 0x1006
	B921600                          = 0x1007
	BLKBSZGET                        = 0x40081270
	BLKBSZSET                        = 0x80081271
	BLKFLSBUF                        = 0x20001261
	BLKFRAGET                        = 0x20001265
	BLKFRASET                        = 0x20001264
	BLKGETSIZE                       = 0x20001260
	BLKGETSIZE64                     = 0x40081272
	BLKPBSZGET                       = 0x2000127b
	BLKRAGET                         = 0x20001263
	BLKRASET                         = 0x20001262
	BLKROGET                         = 0x2000125e
	BLKROSET                         = 0x2000125d
	BLKRRPART                        = 0x2000125f
	BLKSECTGET                       = 0x20001267
	BLKSECTSET                       = 0x20001266
	BLKSSZGET                        = 0x20001268
	BOTHER                           = 0x1000
	BS1                              = 0x2000
	BSDLY                            = 0x2000
	CBAUD                            = 0x100f
	CBAUDEX                          = 0x1000
	CIBAUD                           = 0x100f0000
	CLOCAL                           = 0x800
	CR1                              = 0x200
	CR2                              = 0x400
	CR3                              = 0x600
	CRDLY                            = 0x600
	CREAD                            = 0x80
	CS6                              = 0x10
	CS7                              = 0x20
	CS8                              = 0x30
	CSIZE                            = 0x30
	CSTOPB                           = 0x40
	ECHOCTL                          = 0x200
	ECHOE                            = 0x10
	ECHOK                            = 0x20
	ECHOKE                           = 0x800
	ECHONL                           = 0x40
	ECHOPRT                          = 0x400
	EFD_CLOEXEC                      = 0x80000
	EFD_NONBLOCK                     = 0x80
	EPOLL_CLOEXEC                    = 0x80000
	EXTPROC                          = 0x10000
	FF1                              = 0x8000
	FFDLY                            = 0x8000
	FLUSHO                           = 0x2000
	FS_IOC_ENABLE_VERITY             = 0x80806685
	FS_IOC_GETFLAGS                  = 0x40086601
	FS_IOC_GET_ENCRYPTION_POLICY     = 0x800c6615
	FS_IOC_GET_ENCRYPTION_PWSALT     = 0x80106614
	FS_IOC_SET_ENCRYPTION_POLICY     = 0x400c6613
	F_GETLK                          = 0xe
	F_GETLK64                        = 0xe
	F_GETOWN                         = 0x17
	F_RDLCK                          = 0x0
	F_SETLK                          = 0x6
	F_SETLK64                        = 0x6
	F_SETLKW                         = 0x7
	F_SETLKW64                       = 0x7
	F_SETOWN                         = 0x18
	F_UNLCK                          = 0x2
	F_WRLCK                          = 0x1
	HUPCL                            = 0x400
	ICANON                           = 0x2
	IEXTEN                           = 0x100
	IN_CLOEXEC                       = 0x80000
	IN_NONBLOCK                      = 0x80
	IOCTL_VM_SOCKETS_GET_LOCAL_CID   = 0x200007b9
	ISIG                             = 0x1
	IUCLC                            = 0x200
	IXOFF                            = 0x1000
	IXON                             = 0x400
	MAP_ANON                         = 0x800
	MAP_ANONYMOUS                    = 0x800
	MAP_DENYWRITE                    = 0x2000
	MAP_EXECUTABLE                   = 0x4000
	MAP_GROWSDOWN                    = 0x1000
	MAP_HUGETLB                      = 0x80000
	MAP_LOCKED                       = 0x8000
	MAP_NONBLOCK                     = 0x20000
	MAP_NORESERVE                    = 0x400
	MAP_POPULATE                     = 0x10000
	MAP_RENAME                       = 0x800
	MAP_STACK                        = 0x40000
	MCL_CURRENT                      = 0x1
	MCL_FUTURE                       = 0x2
	MCL_ONFAULT                      = 0x4
	NFDBITS                          = 0x40
	NLDLY                            = 0x100
	NOFLSH                           = 0x80
	NS_GET_NSTYPE                    = 0x2000b703
	NS_GET_OWNER_UID                 = 0x2000b704
	NS_GET_PARENT                    = 0x2000b702
	NS_GET_USERNS                    = 0x2000b701
	OLCUC                            = 0x2
	ONLCR                            = 0x4
	O_APPEND                         = 0x8
	O_ASYNC                          = 0x1000
	O_CLOEXEC                        = 0x80000
	O_CREAT                          = 0x100
	O_DIRECT                         = 0x8000
	O_DIRECTORY                      = 0x10000
	O_DSYNC                          = 0x10
	O_EXCL                           = 0x400
	O_FSYNC                          = 0x4010
	O_LARGEFILE                      = 0x0
	O_NDELAY                         = 0x80
	O_NOATIME                        = 0x40000
	O_NOCTTY                         = 0x800
	O_NOFOLLOW                       = 0x20000
	O_NONBLOCK                       = 0x80
	O_PATH                           = 0x200000
	O_RSYNC                          = 0x4010
	O_SYNC                           = 0x4010
	O_TMPFILE                        = 0x410000
	O_TRUNC                          = 0x200
	PARENB                           = 0x100
	PARODD                           = 0x200
	PENDIN                           = 0x4000
	PERF_EVENT_IOC_DISABLE           = 0x20002401
	PERF_EVENT_IOC_ENABLE            = 0x20002400
	PERF_EVENT_IOC_ID                = 0x40082407
	PERF_EVENT_IOC_MODIFY_ATTRIBUTES = 0x8008240b
	PERF_EVENT_IOC_PAUSE_OUTPUT      = 0x80042409
	PERF_EVENT_IOC_PERIOD            = 0x80082404
	PERF_EVENT_IOC_QUERY_BPF         = 0xc008240a
	PERF_EVENT_IOC_REFRESH           = 0x20002402
	PERF_EVENT_IOC_RESET             = 0x20002403
	PERF_EVENT_IOC_SET_BPF           = 0x80042408
	PERF_EVENT_IOC_SET_FILTER        = 0x80082406
	PERF_EVENT_IOC_SET_OUTPUT        = 0x20002405
	PPPIOCATTACH                     = 0x8004743d
	PPPIOCATTCHAN                    = 0x80047438
	PPPIOCCONNECT                    = 0x8004743a
	PPPIOCDETACH                     = 0x8004743c
	PPPIOCDISCONN                    = 0x20007439
	PPPIOCGASYNCMAP                  = 0x40047458
	PPPIOCGCHAN                      = 0x40047437
	PPPIOCGDEBUG                     = 0x40047441
	PPPIOCGFLAGS                     = 0x4004745a
	PPPIOCGIDLE                      = 0x4010743f
	PPPIOCGIDLE32                    = 0x4008743f
	PPPIOCGIDLE64                    = 0x4010743f
	PPPIOCGL2TPSTATS                 = 0x40487436
	PPPIOCGMRU                       = 0x40047453
	PPPIOCGRASYNCMAP                 = 0x40047455
	PPPIOCGUNIT                      = 0x40047456
	PPPIOCGXASYNCMAP                 = 0x40207450
	PPPIOCSACTIVE                    = 0x80107446
	PPPIOCSASYNCMAP                  = 0x80047457
	PPPIOCSCOMPRESS                  = 0x8010744d
	PPPIOCSDEBUG                     = 0x80047440
	PPPIOCSFLAGS                     = 0x80047459
	PPPIOCSMAXCID                    = 0x80047451
	PPPIOCSMRRU                      = 0x8004743b
	PPPIOCSMRU                       = 0x80047452
	PPPIOCSNPMODE                    = 0x8008744b
	PPPIOCSPASS                      = 0x80107447
	PPPIOCSRASYNCMAP                 = 0x80047454
	PPPIOCSXASYNCMAP                 = 0x8020744f
	PPPIOCXFERUNIT                   = 0x2000744e
	PR_SET_PTRACER_ANY               = 0xffffffffffffffff
	PTRACE_GETFPREGS                 = 0xe
	PTRACE_GET_THREAD_AREA           = 0x19
	PTRACE_GET_THREAD_AREA_3264      = 0xc4
	PTRACE_GET_WATCH_REGS            = 0xd0
	PTRACE_OLDSETOPTIONS             = 0x15
	PTRACE_PEEKDATA_3264             = 0xc1
	PTRACE_PEEKTEXT_3264             = 0xc0
	PTRACE_POKEDATA_3264             = 0xc3
	PTRACE_POKETEXT_3264             = 0xc2
	PTRACE_SETFPREGS                 = 0xf
	PTRACE_SET_THREAD_AREA           = 0x1a
	PTRACE_SET_WATCH_REGS            = 0xd1
	RLIMIT_AS                        = 0x6
	RLIMIT_MEMLOCK                   = 0x9
	RLIMIT_NOFILE                    = 0x5
	RLIMIT_NPROC                     = 0x8
	RLIMIT_RSS                       = 0x7
	RNDADDENTROPY                    = 0x80085203
	RNDADDTOENTCNT                   = 0x80045201
	RNDCLEARPOOL                     = 0x20005206
	RNDGETENTCNT                     = 0x40045200
	RNDGETPOOL                       = 0x40085202
	RNDRESEEDCRNG                    = 0x20005207
	RNDZAPENTCNT                     = 0x20005204
	RTC_AIE_OFF                      = 0x20007002
	RTC_AIE_ON                       = 0x20007001
	RTC_ALM_READ                     = 0x40247008
	RTC_ALM_SET                      = 0x80247007
	RTC_EPOCH_READ                   = 0x4008700d
	RTC_EPOCH_SET                    = 0x8008700e
	RTC_IRQP_READ                    = 0x4008700b
	RTC_IRQP_SET                     = 0x8008700c
	RTC_PIE_OFF                      = 0x20007006
	RTC_PIE_ON                       = 0x20007005
	RTC_PLL_GET                      = 0x40207011
	RTC_PLL_SET                      = 0x80207012
	RTC_RD_TIME                      = 0x40247009
	RTC_SET_TIME                     = 0x8024700a
	RTC_UIE_OFF                      = 0x20007004
	RTC_UIE_ON                       = 0x20007003
	RTC_VL_CLR                       = 0x20007014
	RTC_VL_READ                      = 0x40047013
	RTC_WIE_OFF                      = 0x20007010
	RTC_WIE_ON                       = 0x2000700f
	RTC_WKALM_RD                     = 0x40287010
	RTC_WKALM_SET                    = 0x8028700f
	SCM_TIMESTAMPING                 = 0x25
	SCM_TIMESTAMPING_OPT_STATS       = 0x36
	SCM_TIMESTAMPING_PKTINFO         = 0x3a
	SCM_TIMESTAMPNS                  = 0x23
	SCM_TXTIME                       = 0x3d
	SCM_WIFI_STATUS                  = 0x29
	SFD_CLOEXEC                      = 0x80000
	SFD_NONBLOCK                     = 0x80
	SIOCATMARK                       = 0x40047307
	SIOCGPGRP                        = 0x40047309
	SIOCGSTAMPNS_NEW                 = 0x40108907
	SIOCGSTAMP_NEW                   = 0x40108906
	SIOCINQ                          = 0x467f
	SIOCOUTQ                         = 0x7472
	SIOCSPGRP                        = 0x80047308
	SOCK_CLOEXEC                     = 0x80000
	SOCK_DGRAM                       = 0x1
	SOCK_NONBLOCK                    = 0x80
	SOCK_STREAM                      = 0x2
	SOL_SOCKET                       = 0xffff
	SO_ACCEPTCONN                    = 0x1009
	SO_ATTACH_BPF                    = 0x32
	SO_ATTACH_REUSEPORT_CBPF         = 0x33
	SO_ATTACH_REUSEPORT_EBPF         = 0x34
	SO_BINDTODEVICE                  = 0x19
	SO_BINDTOIFINDEX                 = 0x3e
	SO_BPF_EXTENSIONS                = 0x30
	SO_BROADCAST                     = 0x20
	SO_BSDCOMPAT                     = 0xe
	SO_BUSY_POLL                     = 0x2e
	SO_CNX_ADVICE                    = 0x35
	SO_COOKIE                        = 0x39
	SO_DETACH_REUSEPORT_BPF          = 0x44
	SO_DOMAIN                        = 0x1029
	SO_DONTROUTE                     = 0x10
	SO_ERROR                         = 0x1007
	SO_INCOMING_CPU                  = 0x31
	SO_INCOMING_NAPI_ID              = 0x38
	SO_KEEPALIVE                     = 0x8
	SO_LINGER                        = 0x80
	SO_LOCK_FILTER                   = 0x2c
	SO_MARK                          = 0x24
	SO_MAX_PACING_RATE               = 0x2f
	SO_MEMINFO                       = 0x37
	SO_NOFCS                         = 0x2b
	SO_OOBINLINE                     = 0x100
	SO_PASSCRED                      = 0x11
	SO_PASSSEC                       = 0x22
	SO_PEEK_OFF                      = 0x2a
	SO_PEERCRED                      = 0x12
	SO_PEERGROUPS                    = 0x3b
	SO_PEERSEC                       = 0x1e
	SO_PROTOCOL                      = 0x1028
	SO_RCVBUF                        = 0x1002
	SO_RCVBUFFORCE                   = 0x21
	SO_RCVLOWAT                      = 0x1004
	SO_RCVTIMEO                      = 0x1006
	SO_RCVTIMEO_NEW                  = 0x42
	SO_RCVTIMEO_OLD                  = 0x1006
	SO_REUSEADDR                     = 0x4
	SO_REUSEPORT                     = 0x200
	SO_RXQ_OVFL                      = 0x28
	SO_SECURITY_AUTHENTICATION       = 0x16
	SO_SECURITY_ENCRYPTION_NETWORK   = 0x18
	SO_SECURITY_ENCRYPTION_TRANSPORT = 0x17
	SO_SELECT_ERR_QUEUE              = 0x2d
	SO_SNDBUF                        = 0x1001
	SO_SNDBUFFORCE                   = 0x1f
	SO_SNDLOWAT                      = 0x1003
	SO_SNDTIMEO                      = 0x1005
	SO_SNDTIMEO_NEW                  = 0x43
	SO_SNDTIMEO_OLD                  = 0x1005
	SO_STYLE                         = 0x1008
	SO_TIMESTAMPING                  = 0x25
	SO_TIMESTAMPING_NEW              = 0x41
	SO_TIMESTAMPING_OLD              = 0x25
	SO_TIMESTAMPNS                   = 0x23
	SO_TIMESTAMPNS_NEW               = 0x40
	SO_TIMESTAMPNS_OLD               = 0x23
	SO_TIMESTAMP_NEW                 = 0x3f
	SO_TXTIME                        = 0x3d
	SO_TYPE                          = 0x1008
	SO_WIFI_STATUS                   = 0x29
	SO_ZEROCOPY                      = 0x3c
	TAB1                             = 0x800
	TAB2                             = 0x1000
	TAB3                             = 0x1800
	TABDLY                           = 0x1800
	TCFLSH                           = 0x5407
	TCGETA                           = 0x5401
	TCGETS                           = 0x540d
	TCGETS2                          = 0x4030542a
	TCSAFLUSH                        = 0x5410
	TCSBRK                           = 0x5405
	TCSBRKP                          = 0x5486
	TCSETA                           = 0x5402
	TCSETAF                          = 0x5404
	TCSETAW                          = 0x5403
	TCSETS                           = 0x540e
	TCSETS2                          = 0x8030542b
	TCSETSF                          = 0x5410
	TCSETSF2                         = 0x8030542d
	TCSETSW                          = 0x540f
	TCSETSW2                         = 0x8030542c
	TCXONC                           = 0x5406
	TFD_CLOEXEC                      = 0x80000
	TFD_NONBLOCK                     = 0x80
	TIOCCBRK                         = 0x5428
	TIOCCONS                         = 0x80047478
	TIOCEXCL                         = 0x740d
	TIOCGDEV                         = 0x40045432
	TIOCGETD                         = 0x7400
	TIOCGETP                         = 0x7408
	TIOCGEXCL                        = 0x40045440
	TIOCGICOUNT                      = 0x5492
	TIOCGISO7816                     = 0x40285442
	TIOCGLCKTRMIOS                   = 0x548b
	TIOCGLTC                         = 0x7474
	TIOCGPGRP                        = 0x40047477
	TIOCGPKT                         = 0x40045438
	TIOCGPTLCK                       = 0x40045439
	TIOCGPTN                         = 0x40045430
	TIOCGPTPEER                      = 0x20005441
	TIOCGRS485                       = 0x4020542e
	TIOCGSERIAL                      = 0x5484
	TIOCGSID                         = 0x7416
	TIOCGSOFTCAR                     = 0x5481
	TIOCGWINSZ                       = 0x40087468
	TIOCINQ                          = 0x467f
	TIOCLINUX                        = 0x5483
	TIOCMBIC                         = 0x741c
	TIOCMBIS                         = 0x741b
	TIOCMGET                         = 0x741d
	TIOCMIWAIT                       = 0x5491
	TIOCMSET                         = 0x741a
	TIOCM_CAR                        = 0x100
	TIOCM_CD                         = 0x100
	TIOCM_CTS                        = 0x40
	TIOCM_DSR                        = 0x400
	TIOCM_RI                         = 0x200
	TIOCM_RNG                        = 0x200
	TIOCM_SR                         = 0x20
	TIOCM_ST                         = 0x10
	TIOCNOTTY                        = 0x5471
	TIOCNXCL                         = 0x740e
	TIOCOUTQ                         = 0x7472
	TIOCPKT                          = 0x5470
	TIOCSBRK                         = 0x5427
	TIOCSCTTY                        = 0x5480
	TIOCSERCONFIG                    = 0x5488
	TIOCSERGETLSR                    = 0x548e
	TIOCSERGETMULTI                  = 0x548f
	TIOCSERGSTRUCT                   = 0x548d
	TIOCSERGWILD                     = 0x5489
	TIOCSERSETMULTI                  = 0x5490
	TIOCSERSWILD                     = 0x548a
	TIOCSER_TEMT                     = 0x1
	TIOCSETD                         = 0x7401
	TIOCSETN                         = 0x740a
	TIOCSETP                         = 0x7409
	TIOCSIG                          = 0x80045436
	TIOCSISO7816                     = 0xc0285443
	TIOCSLCKTRMIOS                   = 0x548c
	TIOCSLTC                         = 0x7475
	TIOCSPGRP                        = 0x80047476
	TIOCSPTLCK                       = 0x80045431
	TIOCSRS485                       = 0xc020542f
	TIOCSSERIAL                      = 0x5485
	TIOCSSOFTCAR                     = 0x5482
	TIOCSTI                          = 0x5472
	TIOCSWINSZ                       = 0x80087467
	TIOCVHANGUP                      = 0x5437
	TOSTOP                           = 0x8000
	TUNATTACHFILTER                  = 0x801054d5
	TUNDETACHFILTER                  = 0x801054d6
	TUNGETDEVNETNS                   = 0x200054e3
	TUNGETFEATURES                   = 0x400454cf
	TUNGETFILTER                     = 0x401054db
	TUNGETIFF                        = 0x400454d2
	TUNGETSNDBUF                     = 0x400454d3
	TUNGETVNETBE                     = 0x400454df
	TUNGETVNETHDRSZ                  = 0x400454d7
	TUNGETVNETLE                     = 0x400454dd
	TUNSETCARRIER                    = 0x800454e2
	TUNSETDEBUG                      = 0x800454c9
	TUNSETFILTEREBPF                 = 0x400454e1
	TUNSETGROUP                      = 0x800454ce
	TUNSETIFF                        = 0x800454ca
	TUNSETIFINDEX                    = 0x800454da
	TUNSETLINK                       = 0x800454cd
	TUNSETNOCSUM                     = 0x800454c8
	TUNSETOFFLOAD                    = 0x800454d0
	TUNSETOWNER                      = 0x800454cc
	TUNSETPERSIST                    = 0x800454cb
	TUNSETQUEUE                      = 0x800454d9
	TUNSETSNDBUF                     = 0x800454d4
	TUNSETSTEERINGEBPF               = 0x400454e0
	TUNSETTXFILTER                   = 0x800454d1
	TUNSETVNETBE                     = 0x800454de
	TUNSETVNETHDRSZ                  = 0x800454d8
	TUNSETVNETLE                     = 0x800454dc
	UBI_IOCATT                       = 0x80186f40
	UBI_IOCDET                       = 0x80046f41
	UBI_IOCEBCH                      = 0x80044f02
	UBI_IOCEBER                      = 0x80044f01
	UBI_IOCEBISMAP                   = 0x40044f05
	UBI_IOCEBMAP                     = 0x80084f03
	UBI_IOCEBUNMAP                   = 0x80044f04
	UBI_IOCMKVOL                     = 0x80986f00
	UBI_IOCRMVOL                     = 0x80046f01
	UBI_IOCRNVOL                     = 0x91106f03
	UBI_IOCRPEB                      = 0x80046f04
	UBI_IOCRSVOL                     = 0x800c6f02
	UBI_IOCSETVOLPROP                = 0x80104f06
	UBI_IOCSPEB                      = 0x80046f05
	UBI_IOCVOLCRBLK                  = 0x80804f07
	UBI_IOCVOLRMBLK                  = 0x20004f08
	UBI_IOCVOLUP                     = 0x80084f00
	VDISCARD                         = 0xd
	VEOF                             = 0x10
	VEOL                             = 0x11
	VEOL2                            = 0x6
	VMIN                             = 0x4
	VREPRINT                         = 0xc
	VSTART                           = 0x8
	VSTOP                            = 0x9
	VSUSP                            = 0xa
	VSWTC                            = 0x7
	VSWTCH                           = 0x7
	VT1                              = 0x4000
	VTDLY                            = 0x4000
	VTIME                            = 0x5
	VWERASE                          = 0xe
	WDIOC_GETBOOTSTATUS              = 0x40045702
	WDIOC_GETPRETIMEOUT              = 0x40045709
	WDIOC_GETSTATUS                  = 0x40045701
	WDIOC_GETSUPPORT                 = 0x40285700
	WDIOC_GETTEMP                    = 0x40045703
	WDIOC_GETTIMELEFT                = 0x4004570a
	WDIOC_GETTIMEOUT                 = 0x40045707
	WDIOC_KEEPALIVE                  = 0x40045705
	WDIOC_SETOPTIONS                 = 0x40045704
	WORDSIZE                         = 0x40
	XCASE                            = 0x4
	XTABS                            = 0x1800
)

// Errors
const (
	EADDRINUSE      = syscall.Errno(0x7d)
	EADDRNOTAVAIL   = syscall.Errno(0x7e)
	EADV            = syscall.Errno(0x44)
	EAFNOSUPPORT    = syscall.Errno(0x7c)
	EALREADY        = syscall.Errno(0x95)
	EBADE           = syscall.Errno(0x32)
	EBADFD          = syscall.Errno(0x51)
	EBADMSG         = syscall.Errno(0x4d)
	EBADR           = syscall.Errno(0x33)
	EBADRQC         = syscall.Errno(0x36)
	EBADSLT         = syscall.Errno(0x37)
	EBFONT          = syscall.Errno(0x3b)
	ECANCELED       = syscall.Errno(0x9e)
	ECHRNG          = syscall.Errno(0x25)
	ECOMM           = syscall.Errno(0x46)
	ECONNABORTED    = syscall.Errno(0x82)
	ECONNREFUSED    = syscall.Errno(0x92)
	ECONNRESET      = syscall.Errno(0x83)
	EDEADLK         = syscall.Errno(0x2d)
	EDEADLOCK       = syscall.Errno(0x38)
	EDESTADDRREQ    = syscall.Errno(0x60)
	EDOTDOT         = syscall.Errno(0x49)
	EDQUOT          = syscall.Errno(0x46d)
	EHOSTDOWN       = syscall.Errno(0x93)
	EHOSTUNREACH    = syscall.Errno(0x94)
	EHWPOISON       = syscall.Errno(0xa8)
	EIDRM           = syscall.Errno(0x24)
	EILSEQ          = syscall.Errno(0x58)
	EINIT           = syscall.Errno(0x8d)
	EINPROGRESS     = syscall.Errno(0x96)
	EISCONN         = syscall.Errno(0x85)
	EISNAM          = syscall.Errno(0x8b)
	EKEYEXPIRED     = syscall.Errno(0xa2)
	EKEYREJECTED    = syscall.Errno(0xa4)
	EKEYREVOKED     = syscall.Errno(0xa3)
	EL2HLT          = syscall.Errno(0x2c)
	EL2NSYNC        = syscall.Errno(0x26)
	EL3HLT          = syscall.Errno(0x27)
	EL3RST          = syscall.Errno(0x28)
	ELIBACC         = syscall.Errno(0x53)
	ELIBBAD         = syscall.Errno(0x54)
	ELIBEXEC        = syscall.Errno(0x57)
	ELIBMAX         = syscall.Errno(0x56)
	ELIBSCN         = syscall.Errno(0x55)
	ELNRNG          = syscall.Errno(0x29)
	ELOOP           = syscall.Errno(0x5a)
	EMEDIUMTYPE     = syscall.Errno(0xa0)
	EMSGSIZE        = syscall.Errno(0x61)
	EMULTIHOP       = syscall.Errno(0x4a)
	ENAMETOOLONG    = syscall.Errno(0x4e)
	ENAVAIL         = syscall.Errno(0x8a)
	ENETDOWN        = syscall.Errno(0x7f)
	ENETRESET       = syscall.Errno(0x81)
	ENETUNREACH     = syscall.Errno(0x80)
	ENOANO          = syscall.Errno(0x35)
	ENOBUFS         = syscall.Errno(0x84)
	ENOCSI          = syscall.Errno(0x2b)
	ENODATA         = syscall.Errno(0x3d)
	ENOKEY          = syscall.Errno(0xa1)
	ENOLCK          = syscall.Errno(0x2e)
	ENOLINK         = syscall.Errno(0x43)
	ENOMEDIUM       = syscall.Errno(0x9f)
	ENOMSG          = syscall.Errno(0x23)
	ENONET          = syscall.Errno(0x40)
	ENOPKG          = syscall.Errno(0x41)
	ENOPROTOOPT     = syscall.Errno(0x63)
	ENOSR           = syscall.Errno(0x3f)
	ENOSTR          = syscall.Errno(0x3c)
	ENOSYS          = syscall.Errno(0x59)
	ENOTCONN        = syscall.Errno(0x86)
	ENOTEMPTY       = syscall.Errno(0x5d)
	ENOTNAM         = syscall.Errno(0x89)
	ENOTRECOVERABLE = syscall.Errno(0xa6)
	ENOTSOCK        = syscall.Errno(0x5f)
	ENOTSUP         = syscall.Errno(0x7a)
	ENOTUNIQ        = syscall.Errno(0x50)
	EOPNOTSUPP      = syscall.Errno(0x7a)
	EOVERFLOW       = syscall.Errno(0x4f)
	EOWNERDEAD      = syscall.Errno(0xa5)
	EPFNOSUPPORT    = syscall.Errno(0x7b)
	EPROTO          = syscall.Errno(0x47)
	EPROTONOSUPPORT = syscall.Errno(0x78)
	EPROTOTYPE      = syscall.Errno(0x62)
	EREMCHG         = syscall.Errno(0x52)
	EREMDEV         = syscall.Errno(0x8e)
	EREMOTE         = syscall.Errno(0x42)
	EREMOTEIO       = syscall.Errno(0x8c)
	ERESTART        = syscall.Errno(0x5b)
	ERFKILL         = syscall.Errno(0xa7)
	ESHUTDOWN       = syscall.Errno(0x8f)
	ESOCKTNOSUPPORT = syscall.Errno(0x79)
	ESRMNT          = syscall.Errno(0x45)
	ESTALE          = syscall.Errno(0x97)
	ESTRPIPE        = syscall.Errno(0x5c)
	ETIME           = syscall.Errno(0x3e)
	ETIMEDOUT       = syscall.Errno(0x91)
	ETOOMANYREFS    = syscall.Errno(0x90)
	EUCLEAN         = syscall.Errno(0x87)
	EUNATCH         = syscall.Errno(0x2a)
	EUSERS          = syscall.Errno(0x5e)
	EXFULL          = syscall.Errno(0x34)
)

// Signals
const (
	SIGBUS    = syscall.Signal(0xa)
	SIGCHLD   = syscall.Signal(0x12)
	SIGCLD    = syscall.Signal(0x12)
	SIGCONT   = syscall.Signal(0x19)
	SIGEMT    = syscall.Signal(0x7)
	SIGIO     = syscall.Signal(0x16)
	SIGPOLL   = syscall.Signal(0x16)
	SIGPROF   = syscall.Signal(0x1d)
	SIGPWR    = syscall.Signal(0x13)
	SIGSTOP   = syscall.Signal(0x17)
	SIGSYS    = syscall.Signal(0xc)
	SIGTSTP   = syscall.Signal(0x18)
	SIGTTIN   = syscall.Signal(0x1a)
	SIGTTOU   = syscall.Signal(0x1b)
	SIGURG    = syscall.Signal(0x15)
	SIGUSR1   = syscall.Signal(0x10)
	SIGUSR2   = syscall.Signal(0x11)
	SIGVTALRM = syscall.Signal(0x1c)
	SIGWINCH  = syscall.Signal(0x14)
	SIGXCPU   = syscall.Signal(0x1e)
	SIGXFSZ   = syscall.Signal(0x1f)
)

// Error table
var errorList = [...]struct ***REMOVED***
	num  syscall.Errno
	name string
	desc string
***REMOVED******REMOVED***
	***REMOVED***1, "EPERM", "operation not permitted"***REMOVED***,
	***REMOVED***2, "ENOENT", "no such file or directory"***REMOVED***,
	***REMOVED***3, "ESRCH", "no such process"***REMOVED***,
	***REMOVED***4, "EINTR", "interrupted system call"***REMOVED***,
	***REMOVED***5, "EIO", "input/output error"***REMOVED***,
	***REMOVED***6, "ENXIO", "no such device or address"***REMOVED***,
	***REMOVED***7, "E2BIG", "argument list too long"***REMOVED***,
	***REMOVED***8, "ENOEXEC", "exec format error"***REMOVED***,
	***REMOVED***9, "EBADF", "bad file descriptor"***REMOVED***,
	***REMOVED***10, "ECHILD", "no child processes"***REMOVED***,
	***REMOVED***11, "EAGAIN", "resource temporarily unavailable"***REMOVED***,
	***REMOVED***12, "ENOMEM", "cannot allocate memory"***REMOVED***,
	***REMOVED***13, "EACCES", "permission denied"***REMOVED***,
	***REMOVED***14, "EFAULT", "bad address"***REMOVED***,
	***REMOVED***15, "ENOTBLK", "block device required"***REMOVED***,
	***REMOVED***16, "EBUSY", "device or resource busy"***REMOVED***,
	***REMOVED***17, "EEXIST", "file exists"***REMOVED***,
	***REMOVED***18, "EXDEV", "invalid cross-device link"***REMOVED***,
	***REMOVED***19, "ENODEV", "no such device"***REMOVED***,
	***REMOVED***20, "ENOTDIR", "not a directory"***REMOVED***,
	***REMOVED***21, "EISDIR", "is a directory"***REMOVED***,
	***REMOVED***22, "EINVAL", "invalid argument"***REMOVED***,
	***REMOVED***23, "ENFILE", "too many open files in system"***REMOVED***,
	***REMOVED***24, "EMFILE", "too many open files"***REMOVED***,
	***REMOVED***25, "ENOTTY", "inappropriate ioctl for device"***REMOVED***,
	***REMOVED***26, "ETXTBSY", "text file busy"***REMOVED***,
	***REMOVED***27, "EFBIG", "file too large"***REMOVED***,
	***REMOVED***28, "ENOSPC", "no space left on device"***REMOVED***,
	***REMOVED***29, "ESPIPE", "illegal seek"***REMOVED***,
	***REMOVED***30, "EROFS", "read-only file system"***REMOVED***,
	***REMOVED***31, "EMLINK", "too many links"***REMOVED***,
	***REMOVED***32, "EPIPE", "broken pipe"***REMOVED***,
	***REMOVED***33, "EDOM", "numerical argument out of domain"***REMOVED***,
	***REMOVED***34, "ERANGE", "numerical result out of range"***REMOVED***,
	***REMOVED***35, "ENOMSG", "no message of desired type"***REMOVED***,
	***REMOVED***36, "EIDRM", "identifier removed"***REMOVED***,
	***REMOVED***37, "ECHRNG", "channel number out of range"***REMOVED***,
	***REMOVED***38, "EL2NSYNC", "level 2 not synchronized"***REMOVED***,
	***REMOVED***39, "EL3HLT", "level 3 halted"***REMOVED***,
	***REMOVED***40, "EL3RST", "level 3 reset"***REMOVED***,
	***REMOVED***41, "ELNRNG", "link number out of range"***REMOVED***,
	***REMOVED***42, "EUNATCH", "protocol driver not attached"***REMOVED***,
	***REMOVED***43, "ENOCSI", "no CSI structure available"***REMOVED***,
	***REMOVED***44, "EL2HLT", "level 2 halted"***REMOVED***,
	***REMOVED***45, "EDEADLK", "resource deadlock avoided"***REMOVED***,
	***REMOVED***46, "ENOLCK", "no locks available"***REMOVED***,
	***REMOVED***50, "EBADE", "invalid exchange"***REMOVED***,
	***REMOVED***51, "EBADR", "invalid request descriptor"***REMOVED***,
	***REMOVED***52, "EXFULL", "exchange full"***REMOVED***,
	***REMOVED***53, "ENOANO", "no anode"***REMOVED***,
	***REMOVED***54, "EBADRQC", "invalid request code"***REMOVED***,
	***REMOVED***55, "EBADSLT", "invalid slot"***REMOVED***,
	***REMOVED***56, "EDEADLOCK", "file locking deadlock error"***REMOVED***,
	***REMOVED***59, "EBFONT", "bad font file format"***REMOVED***,
	***REMOVED***60, "ENOSTR", "device not a stream"***REMOVED***,
	***REMOVED***61, "ENODATA", "no data available"***REMOVED***,
	***REMOVED***62, "ETIME", "timer expired"***REMOVED***,
	***REMOVED***63, "ENOSR", "out of streams resources"***REMOVED***,
	***REMOVED***64, "ENONET", "machine is not on the network"***REMOVED***,
	***REMOVED***65, "ENOPKG", "package not installed"***REMOVED***,
	***REMOVED***66, "EREMOTE", "object is remote"***REMOVED***,
	***REMOVED***67, "ENOLINK", "link has been severed"***REMOVED***,
	***REMOVED***68, "EADV", "advertise error"***REMOVED***,
	***REMOVED***69, "ESRMNT", "srmount error"***REMOVED***,
	***REMOVED***70, "ECOMM", "communication error on send"***REMOVED***,
	***REMOVED***71, "EPROTO", "protocol error"***REMOVED***,
	***REMOVED***73, "EDOTDOT", "RFS specific error"***REMOVED***,
	***REMOVED***74, "EMULTIHOP", "multihop attempted"***REMOVED***,
	***REMOVED***77, "EBADMSG", "bad message"***REMOVED***,
	***REMOVED***78, "ENAMETOOLONG", "file name too long"***REMOVED***,
	***REMOVED***79, "EOVERFLOW", "value too large for defined data type"***REMOVED***,
	***REMOVED***80, "ENOTUNIQ", "name not unique on network"***REMOVED***,
	***REMOVED***81, "EBADFD", "file descriptor in bad state"***REMOVED***,
	***REMOVED***82, "EREMCHG", "remote address changed"***REMOVED***,
	***REMOVED***83, "ELIBACC", "can not access a needed shared library"***REMOVED***,
	***REMOVED***84, "ELIBBAD", "accessing a corrupted shared library"***REMOVED***,
	***REMOVED***85, "ELIBSCN", ".lib section in a.out corrupted"***REMOVED***,
	***REMOVED***86, "ELIBMAX", "attempting to link in too many shared libraries"***REMOVED***,
	***REMOVED***87, "ELIBEXEC", "cannot exec a shared library directly"***REMOVED***,
	***REMOVED***88, "EILSEQ", "invalid or incomplete multibyte or wide character"***REMOVED***,
	***REMOVED***89, "ENOSYS", "function not implemented"***REMOVED***,
	***REMOVED***90, "ELOOP", "too many levels of symbolic links"***REMOVED***,
	***REMOVED***91, "ERESTART", "interrupted system call should be restarted"***REMOVED***,
	***REMOVED***92, "ESTRPIPE", "streams pipe error"***REMOVED***,
	***REMOVED***93, "ENOTEMPTY", "directory not empty"***REMOVED***,
	***REMOVED***94, "EUSERS", "too many users"***REMOVED***,
	***REMOVED***95, "ENOTSOCK", "socket operation on non-socket"***REMOVED***,
	***REMOVED***96, "EDESTADDRREQ", "destination address required"***REMOVED***,
	***REMOVED***97, "EMSGSIZE", "message too long"***REMOVED***,
	***REMOVED***98, "EPROTOTYPE", "protocol wrong type for socket"***REMOVED***,
	***REMOVED***99, "ENOPROTOOPT", "protocol not available"***REMOVED***,
	***REMOVED***120, "EPROTONOSUPPORT", "protocol not supported"***REMOVED***,
	***REMOVED***121, "ESOCKTNOSUPPORT", "socket type not supported"***REMOVED***,
	***REMOVED***122, "ENOTSUP", "operation not supported"***REMOVED***,
	***REMOVED***123, "EPFNOSUPPORT", "protocol family not supported"***REMOVED***,
	***REMOVED***124, "EAFNOSUPPORT", "address family not supported by protocol"***REMOVED***,
	***REMOVED***125, "EADDRINUSE", "address already in use"***REMOVED***,
	***REMOVED***126, "EADDRNOTAVAIL", "cannot assign requested address"***REMOVED***,
	***REMOVED***127, "ENETDOWN", "network is down"***REMOVED***,
	***REMOVED***128, "ENETUNREACH", "network is unreachable"***REMOVED***,
	***REMOVED***129, "ENETRESET", "network dropped connection on reset"***REMOVED***,
	***REMOVED***130, "ECONNABORTED", "software caused connection abort"***REMOVED***,
	***REMOVED***131, "ECONNRESET", "connection reset by peer"***REMOVED***,
	***REMOVED***132, "ENOBUFS", "no buffer space available"***REMOVED***,
	***REMOVED***133, "EISCONN", "transport endpoint is already connected"***REMOVED***,
	***REMOVED***134, "ENOTCONN", "transport endpoint is not connected"***REMOVED***,
	***REMOVED***135, "EUCLEAN", "structure needs cleaning"***REMOVED***,
	***REMOVED***137, "ENOTNAM", "not a XENIX named type file"***REMOVED***,
	***REMOVED***138, "ENAVAIL", "no XENIX semaphores available"***REMOVED***,
	***REMOVED***139, "EISNAM", "is a named type file"***REMOVED***,
	***REMOVED***140, "EREMOTEIO", "remote I/O error"***REMOVED***,
	***REMOVED***141, "EINIT", "unknown error 141"***REMOVED***,
	***REMOVED***142, "EREMDEV", "unknown error 142"***REMOVED***,
	***REMOVED***143, "ESHUTDOWN", "cannot send after transport endpoint shutdown"***REMOVED***,
	***REMOVED***144, "ETOOMANYREFS", "too many references: cannot splice"***REMOVED***,
	***REMOVED***145, "ETIMEDOUT", "connection timed out"***REMOVED***,
	***REMOVED***146, "ECONNREFUSED", "connection refused"***REMOVED***,
	***REMOVED***147, "EHOSTDOWN", "host is down"***REMOVED***,
	***REMOVED***148, "EHOSTUNREACH", "no route to host"***REMOVED***,
	***REMOVED***149, "EALREADY", "operation already in progress"***REMOVED***,
	***REMOVED***150, "EINPROGRESS", "operation now in progress"***REMOVED***,
	***REMOVED***151, "ESTALE", "stale file handle"***REMOVED***,
	***REMOVED***158, "ECANCELED", "operation canceled"***REMOVED***,
	***REMOVED***159, "ENOMEDIUM", "no medium found"***REMOVED***,
	***REMOVED***160, "EMEDIUMTYPE", "wrong medium type"***REMOVED***,
	***REMOVED***161, "ENOKEY", "required key not available"***REMOVED***,
	***REMOVED***162, "EKEYEXPIRED", "key has expired"***REMOVED***,
	***REMOVED***163, "EKEYREVOKED", "key has been revoked"***REMOVED***,
	***REMOVED***164, "EKEYREJECTED", "key was rejected by service"***REMOVED***,
	***REMOVED***165, "EOWNERDEAD", "owner died"***REMOVED***,
	***REMOVED***166, "ENOTRECOVERABLE", "state not recoverable"***REMOVED***,
	***REMOVED***167, "ERFKILL", "operation not possible due to RF-kill"***REMOVED***,
	***REMOVED***168, "EHWPOISON", "memory page has hardware error"***REMOVED***,
	***REMOVED***1133, "EDQUOT", "disk quota exceeded"***REMOVED***,
***REMOVED***

// Signal table
var signalList = [...]struct ***REMOVED***
	num  syscall.Signal
	name string
	desc string
***REMOVED******REMOVED***
	***REMOVED***1, "SIGHUP", "hangup"***REMOVED***,
	***REMOVED***2, "SIGINT", "interrupt"***REMOVED***,
	***REMOVED***3, "SIGQUIT", "quit"***REMOVED***,
	***REMOVED***4, "SIGILL", "illegal instruction"***REMOVED***,
	***REMOVED***5, "SIGTRAP", "trace/breakpoint trap"***REMOVED***,
	***REMOVED***6, "SIGABRT", "aborted"***REMOVED***,
	***REMOVED***7, "SIGEMT", "EMT trap"***REMOVED***,
	***REMOVED***8, "SIGFPE", "floating point exception"***REMOVED***,
	***REMOVED***9, "SIGKILL", "killed"***REMOVED***,
	***REMOVED***10, "SIGBUS", "bus error"***REMOVED***,
	***REMOVED***11, "SIGSEGV", "segmentation fault"***REMOVED***,
	***REMOVED***12, "SIGSYS", "bad system call"***REMOVED***,
	***REMOVED***13, "SIGPIPE", "broken pipe"***REMOVED***,
	***REMOVED***14, "SIGALRM", "alarm clock"***REMOVED***,
	***REMOVED***15, "SIGTERM", "terminated"***REMOVED***,
	***REMOVED***16, "SIGUSR1", "user defined signal 1"***REMOVED***,
	***REMOVED***17, "SIGUSR2", "user defined signal 2"***REMOVED***,
	***REMOVED***18, "SIGCHLD", "child exited"***REMOVED***,
	***REMOVED***19, "SIGPWR", "power failure"***REMOVED***,
	***REMOVED***20, "SIGWINCH", "window changed"***REMOVED***,
	***REMOVED***21, "SIGURG", "urgent I/O condition"***REMOVED***,
	***REMOVED***22, "SIGIO", "I/O possible"***REMOVED***,
	***REMOVED***23, "SIGSTOP", "stopped (signal)"***REMOVED***,
	***REMOVED***24, "SIGTSTP", "stopped"***REMOVED***,
	***REMOVED***25, "SIGCONT", "continued"***REMOVED***,
	***REMOVED***26, "SIGTTIN", "stopped (tty input)"***REMOVED***,
	***REMOVED***27, "SIGTTOU", "stopped (tty output)"***REMOVED***,
	***REMOVED***28, "SIGVTALRM", "virtual timer expired"***REMOVED***,
	***REMOVED***29, "SIGPROF", "profiling timer expired"***REMOVED***,
	***REMOVED***30, "SIGXCPU", "CPU time limit exceeded"***REMOVED***,
	***REMOVED***31, "SIGXFSZ", "file size limit exceeded"***REMOVED***,
***REMOVED***
