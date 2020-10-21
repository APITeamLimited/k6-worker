// go run mksysnum.go https://cvsweb.openbsd.org/cgi-bin/cvsweb/~checkout~/src/sys/kern/syscalls.master
// Code generated by the command above; see README.md. DO NOT EDIT.

// +build amd64,openbsd

package unix

const (
	SYS_EXIT           = 1   // ***REMOVED*** void sys_exit(int rval); ***REMOVED***
	SYS_FORK           = 2   // ***REMOVED*** int sys_fork(void); ***REMOVED***
	SYS_READ           = 3   // ***REMOVED*** ssize_t sys_read(int fd, void *buf, size_t nbyte); ***REMOVED***
	SYS_WRITE          = 4   // ***REMOVED*** ssize_t sys_write(int fd, const void *buf, size_t nbyte); ***REMOVED***
	SYS_OPEN           = 5   // ***REMOVED*** int sys_open(const char *path, int flags, ... mode_t mode); ***REMOVED***
	SYS_CLOSE          = 6   // ***REMOVED*** int sys_close(int fd); ***REMOVED***
	SYS_GETENTROPY     = 7   // ***REMOVED*** int sys_getentropy(void *buf, size_t nbyte); ***REMOVED***
	SYS___TFORK        = 8   // ***REMOVED*** int sys___tfork(const struct __tfork *param, size_t psize); ***REMOVED***
	SYS_LINK           = 9   // ***REMOVED*** int sys_link(const char *path, const char *link); ***REMOVED***
	SYS_UNLINK         = 10  // ***REMOVED*** int sys_unlink(const char *path); ***REMOVED***
	SYS_WAIT4          = 11  // ***REMOVED*** pid_t sys_wait4(pid_t pid, int *status, int options, struct rusage *rusage); ***REMOVED***
	SYS_CHDIR          = 12  // ***REMOVED*** int sys_chdir(const char *path); ***REMOVED***
	SYS_FCHDIR         = 13  // ***REMOVED*** int sys_fchdir(int fd); ***REMOVED***
	SYS_MKNOD          = 14  // ***REMOVED*** int sys_mknod(const char *path, mode_t mode, dev_t dev); ***REMOVED***
	SYS_CHMOD          = 15  // ***REMOVED*** int sys_chmod(const char *path, mode_t mode); ***REMOVED***
	SYS_CHOWN          = 16  // ***REMOVED*** int sys_chown(const char *path, uid_t uid, gid_t gid); ***REMOVED***
	SYS_OBREAK         = 17  // ***REMOVED*** int sys_obreak(char *nsize); ***REMOVED*** break
	SYS_GETDTABLECOUNT = 18  // ***REMOVED*** int sys_getdtablecount(void); ***REMOVED***
	SYS_GETRUSAGE      = 19  // ***REMOVED*** int sys_getrusage(int who, struct rusage *rusage); ***REMOVED***
	SYS_GETPID         = 20  // ***REMOVED*** pid_t sys_getpid(void); ***REMOVED***
	SYS_MOUNT          = 21  // ***REMOVED*** int sys_mount(const char *type, const char *path, int flags, void *data); ***REMOVED***
	SYS_UNMOUNT        = 22  // ***REMOVED*** int sys_unmount(const char *path, int flags); ***REMOVED***
	SYS_SETUID         = 23  // ***REMOVED*** int sys_setuid(uid_t uid); ***REMOVED***
	SYS_GETUID         = 24  // ***REMOVED*** uid_t sys_getuid(void); ***REMOVED***
	SYS_GETEUID        = 25  // ***REMOVED*** uid_t sys_geteuid(void); ***REMOVED***
	SYS_PTRACE         = 26  // ***REMOVED*** int sys_ptrace(int req, pid_t pid, caddr_t addr, int data); ***REMOVED***
	SYS_RECVMSG        = 27  // ***REMOVED*** ssize_t sys_recvmsg(int s, struct msghdr *msg, int flags); ***REMOVED***
	SYS_SENDMSG        = 28  // ***REMOVED*** ssize_t sys_sendmsg(int s, const struct msghdr *msg, int flags); ***REMOVED***
	SYS_RECVFROM       = 29  // ***REMOVED*** ssize_t sys_recvfrom(int s, void *buf, size_t len, int flags, struct sockaddr *from, socklen_t *fromlenaddr); ***REMOVED***
	SYS_ACCEPT         = 30  // ***REMOVED*** int sys_accept(int s, struct sockaddr *name, socklen_t *anamelen); ***REMOVED***
	SYS_GETPEERNAME    = 31  // ***REMOVED*** int sys_getpeername(int fdes, struct sockaddr *asa, socklen_t *alen); ***REMOVED***
	SYS_GETSOCKNAME    = 32  // ***REMOVED*** int sys_getsockname(int fdes, struct sockaddr *asa, socklen_t *alen); ***REMOVED***
	SYS_ACCESS         = 33  // ***REMOVED*** int sys_access(const char *path, int amode); ***REMOVED***
	SYS_CHFLAGS        = 34  // ***REMOVED*** int sys_chflags(const char *path, u_int flags); ***REMOVED***
	SYS_FCHFLAGS       = 35  // ***REMOVED*** int sys_fchflags(int fd, u_int flags); ***REMOVED***
	SYS_SYNC           = 36  // ***REMOVED*** void sys_sync(void); ***REMOVED***
	SYS_STAT           = 38  // ***REMOVED*** int sys_stat(const char *path, struct stat *ub); ***REMOVED***
	SYS_GETPPID        = 39  // ***REMOVED*** pid_t sys_getppid(void); ***REMOVED***
	SYS_LSTAT          = 40  // ***REMOVED*** int sys_lstat(const char *path, struct stat *ub); ***REMOVED***
	SYS_DUP            = 41  // ***REMOVED*** int sys_dup(int fd); ***REMOVED***
	SYS_FSTATAT        = 42  // ***REMOVED*** int sys_fstatat(int fd, const char *path, struct stat *buf, int flag); ***REMOVED***
	SYS_GETEGID        = 43  // ***REMOVED*** gid_t sys_getegid(void); ***REMOVED***
	SYS_PROFIL         = 44  // ***REMOVED*** int sys_profil(caddr_t samples, size_t size, u_long offset, u_int scale); ***REMOVED***
	SYS_KTRACE         = 45  // ***REMOVED*** int sys_ktrace(const char *fname, int ops, int facs, pid_t pid); ***REMOVED***
	SYS_SIGACTION      = 46  // ***REMOVED*** int sys_sigaction(int signum, const struct sigaction *nsa, struct sigaction *osa); ***REMOVED***
	SYS_GETGID         = 47  // ***REMOVED*** gid_t sys_getgid(void); ***REMOVED***
	SYS_SIGPROCMASK    = 48  // ***REMOVED*** int sys_sigprocmask(int how, sigset_t mask); ***REMOVED***
	SYS_SETLOGIN       = 50  // ***REMOVED*** int sys_setlogin(const char *namebuf); ***REMOVED***
	SYS_ACCT           = 51  // ***REMOVED*** int sys_acct(const char *path); ***REMOVED***
	SYS_SIGPENDING     = 52  // ***REMOVED*** int sys_sigpending(void); ***REMOVED***
	SYS_FSTAT          = 53  // ***REMOVED*** int sys_fstat(int fd, struct stat *sb); ***REMOVED***
	SYS_IOCTL          = 54  // ***REMOVED*** int sys_ioctl(int fd, u_long com, ... void *data); ***REMOVED***
	SYS_REBOOT         = 55  // ***REMOVED*** int sys_reboot(int opt); ***REMOVED***
	SYS_REVOKE         = 56  // ***REMOVED*** int sys_revoke(const char *path); ***REMOVED***
	SYS_SYMLINK        = 57  // ***REMOVED*** int sys_symlink(const char *path, const char *link); ***REMOVED***
	SYS_READLINK       = 58  // ***REMOVED*** ssize_t sys_readlink(const char *path, char *buf, size_t count); ***REMOVED***
	SYS_EXECVE         = 59  // ***REMOVED*** int sys_execve(const char *path, char * const *argp, char * const *envp); ***REMOVED***
	SYS_UMASK          = 60  // ***REMOVED*** mode_t sys_umask(mode_t newmask); ***REMOVED***
	SYS_CHROOT         = 61  // ***REMOVED*** int sys_chroot(const char *path); ***REMOVED***
	SYS_GETFSSTAT      = 62  // ***REMOVED*** int sys_getfsstat(struct statfs *buf, size_t bufsize, int flags); ***REMOVED***
	SYS_STATFS         = 63  // ***REMOVED*** int sys_statfs(const char *path, struct statfs *buf); ***REMOVED***
	SYS_FSTATFS        = 64  // ***REMOVED*** int sys_fstatfs(int fd, struct statfs *buf); ***REMOVED***
	SYS_FHSTATFS       = 65  // ***REMOVED*** int sys_fhstatfs(const fhandle_t *fhp, struct statfs *buf); ***REMOVED***
	SYS_VFORK          = 66  // ***REMOVED*** int sys_vfork(void); ***REMOVED***
	SYS_GETTIMEOFDAY   = 67  // ***REMOVED*** int sys_gettimeofday(struct timeval *tp, struct timezone *tzp); ***REMOVED***
	SYS_SETTIMEOFDAY   = 68  // ***REMOVED*** int sys_settimeofday(const struct timeval *tv, const struct timezone *tzp); ***REMOVED***
	SYS_SETITIMER      = 69  // ***REMOVED*** int sys_setitimer(int which, const struct itimerval *itv, struct itimerval *oitv); ***REMOVED***
	SYS_GETITIMER      = 70  // ***REMOVED*** int sys_getitimer(int which, struct itimerval *itv); ***REMOVED***
	SYS_SELECT         = 71  // ***REMOVED*** int sys_select(int nd, fd_set *in, fd_set *ou, fd_set *ex, struct timeval *tv); ***REMOVED***
	SYS_KEVENT         = 72  // ***REMOVED*** int sys_kevent(int fd, const struct kevent *changelist, int nchanges, struct kevent *eventlist, int nevents, const struct timespec *timeout); ***REMOVED***
	SYS_MUNMAP         = 73  // ***REMOVED*** int sys_munmap(void *addr, size_t len); ***REMOVED***
	SYS_MPROTECT       = 74  // ***REMOVED*** int sys_mprotect(void *addr, size_t len, int prot); ***REMOVED***
	SYS_MADVISE        = 75  // ***REMOVED*** int sys_madvise(void *addr, size_t len, int behav); ***REMOVED***
	SYS_UTIMES         = 76  // ***REMOVED*** int sys_utimes(const char *path, const struct timeval *tptr); ***REMOVED***
	SYS_FUTIMES        = 77  // ***REMOVED*** int sys_futimes(int fd, const struct timeval *tptr); ***REMOVED***
	SYS_MINCORE        = 78  // ***REMOVED*** int sys_mincore(void *addr, size_t len, char *vec); ***REMOVED***
	SYS_GETGROUPS      = 79  // ***REMOVED*** int sys_getgroups(int gidsetsize, gid_t *gidset); ***REMOVED***
	SYS_SETGROUPS      = 80  // ***REMOVED*** int sys_setgroups(int gidsetsize, const gid_t *gidset); ***REMOVED***
	SYS_GETPGRP        = 81  // ***REMOVED*** int sys_getpgrp(void); ***REMOVED***
	SYS_SETPGID        = 82  // ***REMOVED*** int sys_setpgid(pid_t pid, pid_t pgid); ***REMOVED***
	SYS_FUTEX          = 83  // ***REMOVED*** int sys_futex(uint32_t *f, int op, int val, const struct timespec *timeout, uint32_t *g); ***REMOVED***
	SYS_UTIMENSAT      = 84  // ***REMOVED*** int sys_utimensat(int fd, const char *path, const struct timespec *times, int flag); ***REMOVED***
	SYS_FUTIMENS       = 85  // ***REMOVED*** int sys_futimens(int fd, const struct timespec *times); ***REMOVED***
	SYS_KBIND          = 86  // ***REMOVED*** int sys_kbind(const struct __kbind *param, size_t psize, int64_t proc_cookie); ***REMOVED***
	SYS_CLOCK_GETTIME  = 87  // ***REMOVED*** int sys_clock_gettime(clockid_t clock_id, struct timespec *tp); ***REMOVED***
	SYS_CLOCK_SETTIME  = 88  // ***REMOVED*** int sys_clock_settime(clockid_t clock_id, const struct timespec *tp); ***REMOVED***
	SYS_CLOCK_GETRES   = 89  // ***REMOVED*** int sys_clock_getres(clockid_t clock_id, struct timespec *tp); ***REMOVED***
	SYS_DUP2           = 90  // ***REMOVED*** int sys_dup2(int from, int to); ***REMOVED***
	SYS_NANOSLEEP      = 91  // ***REMOVED*** int sys_nanosleep(const struct timespec *rqtp, struct timespec *rmtp); ***REMOVED***
	SYS_FCNTL          = 92  // ***REMOVED*** int sys_fcntl(int fd, int cmd, ... void *arg); ***REMOVED***
	SYS_ACCEPT4        = 93  // ***REMOVED*** int sys_accept4(int s, struct sockaddr *name, socklen_t *anamelen, int flags); ***REMOVED***
	SYS___THRSLEEP     = 94  // ***REMOVED*** int sys___thrsleep(const volatile void *ident, clockid_t clock_id, const struct timespec *tp, void *lock, const int *abort); ***REMOVED***
	SYS_FSYNC          = 95  // ***REMOVED*** int sys_fsync(int fd); ***REMOVED***
	SYS_SETPRIORITY    = 96  // ***REMOVED*** int sys_setpriority(int which, id_t who, int prio); ***REMOVED***
	SYS_SOCKET         = 97  // ***REMOVED*** int sys_socket(int domain, int type, int protocol); ***REMOVED***
	SYS_CONNECT        = 98  // ***REMOVED*** int sys_connect(int s, const struct sockaddr *name, socklen_t namelen); ***REMOVED***
	SYS_GETDENTS       = 99  // ***REMOVED*** int sys_getdents(int fd, void *buf, size_t buflen); ***REMOVED***
	SYS_GETPRIORITY    = 100 // ***REMOVED*** int sys_getpriority(int which, id_t who); ***REMOVED***
	SYS_PIPE2          = 101 // ***REMOVED*** int sys_pipe2(int *fdp, int flags); ***REMOVED***
	SYS_DUP3           = 102 // ***REMOVED*** int sys_dup3(int from, int to, int flags); ***REMOVED***
	SYS_SIGRETURN      = 103 // ***REMOVED*** int sys_sigreturn(struct sigcontext *sigcntxp); ***REMOVED***
	SYS_BIND           = 104 // ***REMOVED*** int sys_bind(int s, const struct sockaddr *name, socklen_t namelen); ***REMOVED***
	SYS_SETSOCKOPT     = 105 // ***REMOVED*** int sys_setsockopt(int s, int level, int name, const void *val, socklen_t valsize); ***REMOVED***
	SYS_LISTEN         = 106 // ***REMOVED*** int sys_listen(int s, int backlog); ***REMOVED***
	SYS_CHFLAGSAT      = 107 // ***REMOVED*** int sys_chflagsat(int fd, const char *path, u_int flags, int atflags); ***REMOVED***
	SYS_PLEDGE         = 108 // ***REMOVED*** int sys_pledge(const char *promises, const char *execpromises); ***REMOVED***
	SYS_PPOLL          = 109 // ***REMOVED*** int sys_ppoll(struct pollfd *fds, u_int nfds, const struct timespec *ts, const sigset_t *mask); ***REMOVED***
	SYS_PSELECT        = 110 // ***REMOVED*** int sys_pselect(int nd, fd_set *in, fd_set *ou, fd_set *ex, const struct timespec *ts, const sigset_t *mask); ***REMOVED***
	SYS_SIGSUSPEND     = 111 // ***REMOVED*** int sys_sigsuspend(int mask); ***REMOVED***
	SYS_SENDSYSLOG     = 112 // ***REMOVED*** int sys_sendsyslog(const char *buf, size_t nbyte, int flags); ***REMOVED***
	SYS_UNVEIL         = 114 // ***REMOVED*** int sys_unveil(const char *path, const char *permissions); ***REMOVED***
	SYS_GETSOCKOPT     = 118 // ***REMOVED*** int sys_getsockopt(int s, int level, int name, void *val, socklen_t *avalsize); ***REMOVED***
	SYS_THRKILL        = 119 // ***REMOVED*** int sys_thrkill(pid_t tid, int signum, void *tcb); ***REMOVED***
	SYS_READV          = 120 // ***REMOVED*** ssize_t sys_readv(int fd, const struct iovec *iovp, int iovcnt); ***REMOVED***
	SYS_WRITEV         = 121 // ***REMOVED*** ssize_t sys_writev(int fd, const struct iovec *iovp, int iovcnt); ***REMOVED***
	SYS_KILL           = 122 // ***REMOVED*** int sys_kill(int pid, int signum); ***REMOVED***
	SYS_FCHOWN         = 123 // ***REMOVED*** int sys_fchown(int fd, uid_t uid, gid_t gid); ***REMOVED***
	SYS_FCHMOD         = 124 // ***REMOVED*** int sys_fchmod(int fd, mode_t mode); ***REMOVED***
	SYS_SETREUID       = 126 // ***REMOVED*** int sys_setreuid(uid_t ruid, uid_t euid); ***REMOVED***
	SYS_SETREGID       = 127 // ***REMOVED*** int sys_setregid(gid_t rgid, gid_t egid); ***REMOVED***
	SYS_RENAME         = 128 // ***REMOVED*** int sys_rename(const char *from, const char *to); ***REMOVED***
	SYS_FLOCK          = 131 // ***REMOVED*** int sys_flock(int fd, int how); ***REMOVED***
	SYS_MKFIFO         = 132 // ***REMOVED*** int sys_mkfifo(const char *path, mode_t mode); ***REMOVED***
	SYS_SENDTO         = 133 // ***REMOVED*** ssize_t sys_sendto(int s, const void *buf, size_t len, int flags, const struct sockaddr *to, socklen_t tolen); ***REMOVED***
	SYS_SHUTDOWN       = 134 // ***REMOVED*** int sys_shutdown(int s, int how); ***REMOVED***
	SYS_SOCKETPAIR     = 135 // ***REMOVED*** int sys_socketpair(int domain, int type, int protocol, int *rsv); ***REMOVED***
	SYS_MKDIR          = 136 // ***REMOVED*** int sys_mkdir(const char *path, mode_t mode); ***REMOVED***
	SYS_RMDIR          = 137 // ***REMOVED*** int sys_rmdir(const char *path); ***REMOVED***
	SYS_ADJTIME        = 140 // ***REMOVED*** int sys_adjtime(const struct timeval *delta, struct timeval *olddelta); ***REMOVED***
	SYS_GETLOGIN_R     = 141 // ***REMOVED*** int sys_getlogin_r(char *namebuf, u_int namelen); ***REMOVED***
	SYS_SETSID         = 147 // ***REMOVED*** int sys_setsid(void); ***REMOVED***
	SYS_QUOTACTL       = 148 // ***REMOVED*** int sys_quotactl(const char *path, int cmd, int uid, char *arg); ***REMOVED***
	SYS_NFSSVC         = 155 // ***REMOVED*** int sys_nfssvc(int flag, void *argp); ***REMOVED***
	SYS_GETFH          = 161 // ***REMOVED*** int sys_getfh(const char *fname, fhandle_t *fhp); ***REMOVED***
	SYS_SYSARCH        = 165 // ***REMOVED*** int sys_sysarch(int op, void *parms); ***REMOVED***
	SYS_PREAD          = 173 // ***REMOVED*** ssize_t sys_pread(int fd, void *buf, size_t nbyte, int pad, off_t offset); ***REMOVED***
	SYS_PWRITE         = 174 // ***REMOVED*** ssize_t sys_pwrite(int fd, const void *buf, size_t nbyte, int pad, off_t offset); ***REMOVED***
	SYS_SETGID         = 181 // ***REMOVED*** int sys_setgid(gid_t gid); ***REMOVED***
	SYS_SETEGID        = 182 // ***REMOVED*** int sys_setegid(gid_t egid); ***REMOVED***
	SYS_SETEUID        = 183 // ***REMOVED*** int sys_seteuid(uid_t euid); ***REMOVED***
	SYS_PATHCONF       = 191 // ***REMOVED*** long sys_pathconf(const char *path, int name); ***REMOVED***
	SYS_FPATHCONF      = 192 // ***REMOVED*** long sys_fpathconf(int fd, int name); ***REMOVED***
	SYS_SWAPCTL        = 193 // ***REMOVED*** int sys_swapctl(int cmd, const void *arg, int misc); ***REMOVED***
	SYS_GETRLIMIT      = 194 // ***REMOVED*** int sys_getrlimit(int which, struct rlimit *rlp); ***REMOVED***
	SYS_SETRLIMIT      = 195 // ***REMOVED*** int sys_setrlimit(int which, const struct rlimit *rlp); ***REMOVED***
	SYS_MMAP           = 197 // ***REMOVED*** void *sys_mmap(void *addr, size_t len, int prot, int flags, int fd, long pad, off_t pos); ***REMOVED***
	SYS_LSEEK          = 199 // ***REMOVED*** off_t sys_lseek(int fd, int pad, off_t offset, int whence); ***REMOVED***
	SYS_TRUNCATE       = 200 // ***REMOVED*** int sys_truncate(const char *path, int pad, off_t length); ***REMOVED***
	SYS_FTRUNCATE      = 201 // ***REMOVED*** int sys_ftruncate(int fd, int pad, off_t length); ***REMOVED***
	SYS_SYSCTL         = 202 // ***REMOVED*** int sys_sysctl(const int *name, u_int namelen, void *old, size_t *oldlenp, void *new, size_t newlen); ***REMOVED***
	SYS_MLOCK          = 203 // ***REMOVED*** int sys_mlock(const void *addr, size_t len); ***REMOVED***
	SYS_MUNLOCK        = 204 // ***REMOVED*** int sys_munlock(const void *addr, size_t len); ***REMOVED***
	SYS_GETPGID        = 207 // ***REMOVED*** pid_t sys_getpgid(pid_t pid); ***REMOVED***
	SYS_UTRACE         = 209 // ***REMOVED*** int sys_utrace(const char *label, const void *addr, size_t len); ***REMOVED***
	SYS_SEMGET         = 221 // ***REMOVED*** int sys_semget(key_t key, int nsems, int semflg); ***REMOVED***
	SYS_MSGGET         = 225 // ***REMOVED*** int sys_msgget(key_t key, int msgflg); ***REMOVED***
	SYS_MSGSND         = 226 // ***REMOVED*** int sys_msgsnd(int msqid, const void *msgp, size_t msgsz, int msgflg); ***REMOVED***
	SYS_MSGRCV         = 227 // ***REMOVED*** int sys_msgrcv(int msqid, void *msgp, size_t msgsz, long msgtyp, int msgflg); ***REMOVED***
	SYS_SHMAT          = 228 // ***REMOVED*** void *sys_shmat(int shmid, const void *shmaddr, int shmflg); ***REMOVED***
	SYS_SHMDT          = 230 // ***REMOVED*** int sys_shmdt(const void *shmaddr); ***REMOVED***
	SYS_MINHERIT       = 250 // ***REMOVED*** int sys_minherit(void *addr, size_t len, int inherit); ***REMOVED***
	SYS_POLL           = 252 // ***REMOVED*** int sys_poll(struct pollfd *fds, u_int nfds, int timeout); ***REMOVED***
	SYS_ISSETUGID      = 253 // ***REMOVED*** int sys_issetugid(void); ***REMOVED***
	SYS_LCHOWN         = 254 // ***REMOVED*** int sys_lchown(const char *path, uid_t uid, gid_t gid); ***REMOVED***
	SYS_GETSID         = 255 // ***REMOVED*** pid_t sys_getsid(pid_t pid); ***REMOVED***
	SYS_MSYNC          = 256 // ***REMOVED*** int sys_msync(void *addr, size_t len, int flags); ***REMOVED***
	SYS_PIPE           = 263 // ***REMOVED*** int sys_pipe(int *fdp); ***REMOVED***
	SYS_FHOPEN         = 264 // ***REMOVED*** int sys_fhopen(const fhandle_t *fhp, int flags); ***REMOVED***
	SYS_PREADV         = 267 // ***REMOVED*** ssize_t sys_preadv(int fd, const struct iovec *iovp, int iovcnt, int pad, off_t offset); ***REMOVED***
	SYS_PWRITEV        = 268 // ***REMOVED*** ssize_t sys_pwritev(int fd, const struct iovec *iovp, int iovcnt, int pad, off_t offset); ***REMOVED***
	SYS_KQUEUE         = 269 // ***REMOVED*** int sys_kqueue(void); ***REMOVED***
	SYS_MLOCKALL       = 271 // ***REMOVED*** int sys_mlockall(int flags); ***REMOVED***
	SYS_MUNLOCKALL     = 272 // ***REMOVED*** int sys_munlockall(void); ***REMOVED***
	SYS_GETRESUID      = 281 // ***REMOVED*** int sys_getresuid(uid_t *ruid, uid_t *euid, uid_t *suid); ***REMOVED***
	SYS_SETRESUID      = 282 // ***REMOVED*** int sys_setresuid(uid_t ruid, uid_t euid, uid_t suid); ***REMOVED***
	SYS_GETRESGID      = 283 // ***REMOVED*** int sys_getresgid(gid_t *rgid, gid_t *egid, gid_t *sgid); ***REMOVED***
	SYS_SETRESGID      = 284 // ***REMOVED*** int sys_setresgid(gid_t rgid, gid_t egid, gid_t sgid); ***REMOVED***
	SYS_MQUERY         = 286 // ***REMOVED*** void *sys_mquery(void *addr, size_t len, int prot, int flags, int fd, long pad, off_t pos); ***REMOVED***
	SYS_CLOSEFROM      = 287 // ***REMOVED*** int sys_closefrom(int fd); ***REMOVED***
	SYS_SIGALTSTACK    = 288 // ***REMOVED*** int sys_sigaltstack(const struct sigaltstack *nss, struct sigaltstack *oss); ***REMOVED***
	SYS_SHMGET         = 289 // ***REMOVED*** int sys_shmget(key_t key, size_t size, int shmflg); ***REMOVED***
	SYS_SEMOP          = 290 // ***REMOVED*** int sys_semop(int semid, struct sembuf *sops, size_t nsops); ***REMOVED***
	SYS_FHSTAT         = 294 // ***REMOVED*** int sys_fhstat(const fhandle_t *fhp, struct stat *sb); ***REMOVED***
	SYS___SEMCTL       = 295 // ***REMOVED*** int sys___semctl(int semid, int semnum, int cmd, union semun *arg); ***REMOVED***
	SYS_SHMCTL         = 296 // ***REMOVED*** int sys_shmctl(int shmid, int cmd, struct shmid_ds *buf); ***REMOVED***
	SYS_MSGCTL         = 297 // ***REMOVED*** int sys_msgctl(int msqid, int cmd, struct msqid_ds *buf); ***REMOVED***
	SYS_SCHED_YIELD    = 298 // ***REMOVED*** int sys_sched_yield(void); ***REMOVED***
	SYS_GETTHRID       = 299 // ***REMOVED*** pid_t sys_getthrid(void); ***REMOVED***
	SYS___THRWAKEUP    = 301 // ***REMOVED*** int sys___thrwakeup(const volatile void *ident, int n); ***REMOVED***
	SYS___THREXIT      = 302 // ***REMOVED*** void sys___threxit(pid_t *notdead); ***REMOVED***
	SYS___THRSIGDIVERT = 303 // ***REMOVED*** int sys___thrsigdivert(sigset_t sigmask, siginfo_t *info, const struct timespec *timeout); ***REMOVED***
	SYS___GETCWD       = 304 // ***REMOVED*** int sys___getcwd(char *buf, size_t len); ***REMOVED***
	SYS_ADJFREQ        = 305 // ***REMOVED*** int sys_adjfreq(const int64_t *freq, int64_t *oldfreq); ***REMOVED***
	SYS_SETRTABLE      = 310 // ***REMOVED*** int sys_setrtable(int rtableid); ***REMOVED***
	SYS_GETRTABLE      = 311 // ***REMOVED*** int sys_getrtable(void); ***REMOVED***
	SYS_FACCESSAT      = 313 // ***REMOVED*** int sys_faccessat(int fd, const char *path, int amode, int flag); ***REMOVED***
	SYS_FCHMODAT       = 314 // ***REMOVED*** int sys_fchmodat(int fd, const char *path, mode_t mode, int flag); ***REMOVED***
	SYS_FCHOWNAT       = 315 // ***REMOVED*** int sys_fchownat(int fd, const char *path, uid_t uid, gid_t gid, int flag); ***REMOVED***
	SYS_LINKAT         = 317 // ***REMOVED*** int sys_linkat(int fd1, const char *path1, int fd2, const char *path2, int flag); ***REMOVED***
	SYS_MKDIRAT        = 318 // ***REMOVED*** int sys_mkdirat(int fd, const char *path, mode_t mode); ***REMOVED***
	SYS_MKFIFOAT       = 319 // ***REMOVED*** int sys_mkfifoat(int fd, const char *path, mode_t mode); ***REMOVED***
	SYS_MKNODAT        = 320 // ***REMOVED*** int sys_mknodat(int fd, const char *path, mode_t mode, dev_t dev); ***REMOVED***
	SYS_OPENAT         = 321 // ***REMOVED*** int sys_openat(int fd, const char *path, int flags, ... mode_t mode); ***REMOVED***
	SYS_READLINKAT     = 322 // ***REMOVED*** ssize_t sys_readlinkat(int fd, const char *path, char *buf, size_t count); ***REMOVED***
	SYS_RENAMEAT       = 323 // ***REMOVED*** int sys_renameat(int fromfd, const char *from, int tofd, const char *to); ***REMOVED***
	SYS_SYMLINKAT      = 324 // ***REMOVED*** int sys_symlinkat(const char *path, int fd, const char *link); ***REMOVED***
	SYS_UNLINKAT       = 325 // ***REMOVED*** int sys_unlinkat(int fd, const char *path, int flag); ***REMOVED***
	SYS___SET_TCB      = 329 // ***REMOVED*** void sys___set_tcb(void *tcb); ***REMOVED***
	SYS___GET_TCB      = 330 // ***REMOVED*** void *sys___get_tcb(void); ***REMOVED***
)
