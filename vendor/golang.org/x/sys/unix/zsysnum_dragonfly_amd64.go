// go run mksysnum.go https://gitweb.dragonflybsd.org/dragonfly.git/blob_plain/HEAD:/sys/kern/syscalls.master
// Code generated by the command above; see README.md. DO NOT EDIT.

//go:build amd64 && dragonfly
// +build amd64,dragonfly

package unix

const (
	SYS_EXIT  = 1 // ***REMOVED*** void exit(int rval); ***REMOVED***
	SYS_FORK  = 2 // ***REMOVED*** int fork(void); ***REMOVED***
	SYS_READ  = 3 // ***REMOVED*** ssize_t read(int fd, void *buf, size_t nbyte); ***REMOVED***
	SYS_WRITE = 4 // ***REMOVED*** ssize_t write(int fd, const void *buf, size_t nbyte); ***REMOVED***
	SYS_OPEN  = 5 // ***REMOVED*** int open(char *path, int flags, int mode); ***REMOVED***
	SYS_CLOSE = 6 // ***REMOVED*** int close(int fd); ***REMOVED***
	SYS_WAIT4 = 7 // ***REMOVED*** int wait4(int pid, int *status, int options, struct rusage *rusage); ***REMOVED*** wait4 wait_args int
	// SYS_NOSYS = 8;  // ***REMOVED*** int nosys(void); ***REMOVED*** __nosys nosys_args int
	SYS_LINK                   = 9   // ***REMOVED*** int link(char *path, char *link); ***REMOVED***
	SYS_UNLINK                 = 10  // ***REMOVED*** int unlink(char *path); ***REMOVED***
	SYS_CHDIR                  = 12  // ***REMOVED*** int chdir(char *path); ***REMOVED***
	SYS_FCHDIR                 = 13  // ***REMOVED*** int fchdir(int fd); ***REMOVED***
	SYS_MKNOD                  = 14  // ***REMOVED*** int mknod(char *path, int mode, int dev); ***REMOVED***
	SYS_CHMOD                  = 15  // ***REMOVED*** int chmod(char *path, int mode); ***REMOVED***
	SYS_CHOWN                  = 16  // ***REMOVED*** int chown(char *path, int uid, int gid); ***REMOVED***
	SYS_OBREAK                 = 17  // ***REMOVED*** int obreak(char *nsize); ***REMOVED*** break obreak_args int
	SYS_GETFSSTAT              = 18  // ***REMOVED*** int getfsstat(struct statfs *buf, long bufsize, int flags); ***REMOVED***
	SYS_GETPID                 = 20  // ***REMOVED*** pid_t getpid(void); ***REMOVED***
	SYS_MOUNT                  = 21  // ***REMOVED*** int mount(char *type, char *path, int flags, caddr_t data); ***REMOVED***
	SYS_UNMOUNT                = 22  // ***REMOVED*** int unmount(char *path, int flags); ***REMOVED***
	SYS_SETUID                 = 23  // ***REMOVED*** int setuid(uid_t uid); ***REMOVED***
	SYS_GETUID                 = 24  // ***REMOVED*** uid_t getuid(void); ***REMOVED***
	SYS_GETEUID                = 25  // ***REMOVED*** uid_t geteuid(void); ***REMOVED***
	SYS_PTRACE                 = 26  // ***REMOVED*** int ptrace(int req, pid_t pid, caddr_t addr, int data); ***REMOVED***
	SYS_RECVMSG                = 27  // ***REMOVED*** int recvmsg(int s, struct msghdr *msg, int flags); ***REMOVED***
	SYS_SENDMSG                = 28  // ***REMOVED*** int sendmsg(int s, caddr_t msg, int flags); ***REMOVED***
	SYS_RECVFROM               = 29  // ***REMOVED*** int recvfrom(int s, caddr_t buf, size_t len, int flags, caddr_t from, int *fromlenaddr); ***REMOVED***
	SYS_ACCEPT                 = 30  // ***REMOVED*** int accept(int s, caddr_t name, int *anamelen); ***REMOVED***
	SYS_GETPEERNAME            = 31  // ***REMOVED*** int getpeername(int fdes, caddr_t asa, int *alen); ***REMOVED***
	SYS_GETSOCKNAME            = 32  // ***REMOVED*** int getsockname(int fdes, caddr_t asa, int *alen); ***REMOVED***
	SYS_ACCESS                 = 33  // ***REMOVED*** int access(char *path, int flags); ***REMOVED***
	SYS_CHFLAGS                = 34  // ***REMOVED*** int chflags(const char *path, u_long flags); ***REMOVED***
	SYS_FCHFLAGS               = 35  // ***REMOVED*** int fchflags(int fd, u_long flags); ***REMOVED***
	SYS_SYNC                   = 36  // ***REMOVED*** int sync(void); ***REMOVED***
	SYS_KILL                   = 37  // ***REMOVED*** int kill(int pid, int signum); ***REMOVED***
	SYS_GETPPID                = 39  // ***REMOVED*** pid_t getppid(void); ***REMOVED***
	SYS_DUP                    = 41  // ***REMOVED*** int dup(int fd); ***REMOVED***
	SYS_PIPE                   = 42  // ***REMOVED*** int pipe(void); ***REMOVED***
	SYS_GETEGID                = 43  // ***REMOVED*** gid_t getegid(void); ***REMOVED***
	SYS_PROFIL                 = 44  // ***REMOVED*** int profil(caddr_t samples, size_t size, u_long offset, u_int scale); ***REMOVED***
	SYS_KTRACE                 = 45  // ***REMOVED*** int ktrace(const char *fname, int ops, int facs, int pid); ***REMOVED***
	SYS_GETGID                 = 47  // ***REMOVED*** gid_t getgid(void); ***REMOVED***
	SYS_GETLOGIN               = 49  // ***REMOVED*** int getlogin(char *namebuf, size_t namelen); ***REMOVED***
	SYS_SETLOGIN               = 50  // ***REMOVED*** int setlogin(char *namebuf); ***REMOVED***
	SYS_ACCT                   = 51  // ***REMOVED*** int acct(char *path); ***REMOVED***
	SYS_SIGALTSTACK            = 53  // ***REMOVED*** int sigaltstack(stack_t *ss, stack_t *oss); ***REMOVED***
	SYS_IOCTL                  = 54  // ***REMOVED*** int ioctl(int fd, u_long com, caddr_t data); ***REMOVED***
	SYS_REBOOT                 = 55  // ***REMOVED*** int reboot(int opt); ***REMOVED***
	SYS_REVOKE                 = 56  // ***REMOVED*** int revoke(char *path); ***REMOVED***
	SYS_SYMLINK                = 57  // ***REMOVED*** int symlink(char *path, char *link); ***REMOVED***
	SYS_READLINK               = 58  // ***REMOVED*** int readlink(char *path, char *buf, int count); ***REMOVED***
	SYS_EXECVE                 = 59  // ***REMOVED*** int execve(char *fname, char **argv, char **envv); ***REMOVED***
	SYS_UMASK                  = 60  // ***REMOVED*** int umask(int newmask); ***REMOVED*** umask umask_args int
	SYS_CHROOT                 = 61  // ***REMOVED*** int chroot(char *path); ***REMOVED***
	SYS_MSYNC                  = 65  // ***REMOVED*** int msync(void *addr, size_t len, int flags); ***REMOVED***
	SYS_VFORK                  = 66  // ***REMOVED*** pid_t vfork(void); ***REMOVED***
	SYS_SBRK                   = 69  // ***REMOVED*** caddr_t sbrk(size_t incr); ***REMOVED***
	SYS_SSTK                   = 70  // ***REMOVED*** int sstk(size_t incr); ***REMOVED***
	SYS_MUNMAP                 = 73  // ***REMOVED*** int munmap(void *addr, size_t len); ***REMOVED***
	SYS_MPROTECT               = 74  // ***REMOVED*** int mprotect(void *addr, size_t len, int prot); ***REMOVED***
	SYS_MADVISE                = 75  // ***REMOVED*** int madvise(void *addr, size_t len, int behav); ***REMOVED***
	SYS_MINCORE                = 78  // ***REMOVED*** int mincore(const void *addr, size_t len, char *vec); ***REMOVED***
	SYS_GETGROUPS              = 79  // ***REMOVED*** int getgroups(u_int gidsetsize, gid_t *gidset); ***REMOVED***
	SYS_SETGROUPS              = 80  // ***REMOVED*** int setgroups(u_int gidsetsize, gid_t *gidset); ***REMOVED***
	SYS_GETPGRP                = 81  // ***REMOVED*** int getpgrp(void); ***REMOVED***
	SYS_SETPGID                = 82  // ***REMOVED*** int setpgid(int pid, int pgid); ***REMOVED***
	SYS_SETITIMER              = 83  // ***REMOVED*** int setitimer(u_int which, struct itimerval *itv, struct itimerval *oitv); ***REMOVED***
	SYS_SWAPON                 = 85  // ***REMOVED*** int swapon(char *name); ***REMOVED***
	SYS_GETITIMER              = 86  // ***REMOVED*** int getitimer(u_int which, struct itimerval *itv); ***REMOVED***
	SYS_GETDTABLESIZE          = 89  // ***REMOVED*** int getdtablesize(void); ***REMOVED***
	SYS_DUP2                   = 90  // ***REMOVED*** int dup2(int from, int to); ***REMOVED***
	SYS_FCNTL                  = 92  // ***REMOVED*** int fcntl(int fd, int cmd, long arg); ***REMOVED***
	SYS_SELECT                 = 93  // ***REMOVED*** int select(int nd, fd_set *in, fd_set *ou, fd_set *ex, struct timeval *tv); ***REMOVED***
	SYS_FSYNC                  = 95  // ***REMOVED*** int fsync(int fd); ***REMOVED***
	SYS_SETPRIORITY            = 96  // ***REMOVED*** int setpriority(int which, int who, int prio); ***REMOVED***
	SYS_SOCKET                 = 97  // ***REMOVED*** int socket(int domain, int type, int protocol); ***REMOVED***
	SYS_CONNECT                = 98  // ***REMOVED*** int connect(int s, caddr_t name, int namelen); ***REMOVED***
	SYS_GETPRIORITY            = 100 // ***REMOVED*** int getpriority(int which, int who); ***REMOVED***
	SYS_BIND                   = 104 // ***REMOVED*** int bind(int s, caddr_t name, int namelen); ***REMOVED***
	SYS_SETSOCKOPT             = 105 // ***REMOVED*** int setsockopt(int s, int level, int name, caddr_t val, int valsize); ***REMOVED***
	SYS_LISTEN                 = 106 // ***REMOVED*** int listen(int s, int backlog); ***REMOVED***
	SYS_GETTIMEOFDAY           = 116 // ***REMOVED*** int gettimeofday(struct timeval *tp, struct timezone *tzp); ***REMOVED***
	SYS_GETRUSAGE              = 117 // ***REMOVED*** int getrusage(int who, struct rusage *rusage); ***REMOVED***
	SYS_GETSOCKOPT             = 118 // ***REMOVED*** int getsockopt(int s, int level, int name, caddr_t val, int *avalsize); ***REMOVED***
	SYS_READV                  = 120 // ***REMOVED*** int readv(int fd, struct iovec *iovp, u_int iovcnt); ***REMOVED***
	SYS_WRITEV                 = 121 // ***REMOVED*** int writev(int fd, struct iovec *iovp, u_int iovcnt); ***REMOVED***
	SYS_SETTIMEOFDAY           = 122 // ***REMOVED*** int settimeofday(struct timeval *tv, struct timezone *tzp); ***REMOVED***
	SYS_FCHOWN                 = 123 // ***REMOVED*** int fchown(int fd, int uid, int gid); ***REMOVED***
	SYS_FCHMOD                 = 124 // ***REMOVED*** int fchmod(int fd, int mode); ***REMOVED***
	SYS_SETREUID               = 126 // ***REMOVED*** int setreuid(int ruid, int euid); ***REMOVED***
	SYS_SETREGID               = 127 // ***REMOVED*** int setregid(int rgid, int egid); ***REMOVED***
	SYS_RENAME                 = 128 // ***REMOVED*** int rename(char *from, char *to); ***REMOVED***
	SYS_FLOCK                  = 131 // ***REMOVED*** int flock(int fd, int how); ***REMOVED***
	SYS_MKFIFO                 = 132 // ***REMOVED*** int mkfifo(char *path, int mode); ***REMOVED***
	SYS_SENDTO                 = 133 // ***REMOVED*** int sendto(int s, caddr_t buf, size_t len, int flags, caddr_t to, int tolen); ***REMOVED***
	SYS_SHUTDOWN               = 134 // ***REMOVED*** int shutdown(int s, int how); ***REMOVED***
	SYS_SOCKETPAIR             = 135 // ***REMOVED*** int socketpair(int domain, int type, int protocol, int *rsv); ***REMOVED***
	SYS_MKDIR                  = 136 // ***REMOVED*** int mkdir(char *path, int mode); ***REMOVED***
	SYS_RMDIR                  = 137 // ***REMOVED*** int rmdir(char *path); ***REMOVED***
	SYS_UTIMES                 = 138 // ***REMOVED*** int utimes(char *path, struct timeval *tptr); ***REMOVED***
	SYS_ADJTIME                = 140 // ***REMOVED*** int adjtime(struct timeval *delta, struct timeval *olddelta); ***REMOVED***
	SYS_SETSID                 = 147 // ***REMOVED*** int setsid(void); ***REMOVED***
	SYS_QUOTACTL               = 148 // ***REMOVED*** int quotactl(char *path, int cmd, int uid, caddr_t arg); ***REMOVED***
	SYS_STATFS                 = 157 // ***REMOVED*** int statfs(char *path, struct statfs *buf); ***REMOVED***
	SYS_FSTATFS                = 158 // ***REMOVED*** int fstatfs(int fd, struct statfs *buf); ***REMOVED***
	SYS_GETFH                  = 161 // ***REMOVED*** int getfh(char *fname, struct fhandle *fhp); ***REMOVED***
	SYS_SYSARCH                = 165 // ***REMOVED*** int sysarch(int op, char *parms); ***REMOVED***
	SYS_RTPRIO                 = 166 // ***REMOVED*** int rtprio(int function, pid_t pid, struct rtprio *rtp); ***REMOVED***
	SYS_EXTPREAD               = 173 // ***REMOVED*** ssize_t extpread(int fd, void *buf, size_t nbyte, int flags, off_t offset); ***REMOVED***
	SYS_EXTPWRITE              = 174 // ***REMOVED*** ssize_t extpwrite(int fd, const void *buf, size_t nbyte, int flags, off_t offset); ***REMOVED***
	SYS_NTP_ADJTIME            = 176 // ***REMOVED*** int ntp_adjtime(struct timex *tp); ***REMOVED***
	SYS_SETGID                 = 181 // ***REMOVED*** int setgid(gid_t gid); ***REMOVED***
	SYS_SETEGID                = 182 // ***REMOVED*** int setegid(gid_t egid); ***REMOVED***
	SYS_SETEUID                = 183 // ***REMOVED*** int seteuid(uid_t euid); ***REMOVED***
	SYS_PATHCONF               = 191 // ***REMOVED*** int pathconf(char *path, int name); ***REMOVED***
	SYS_FPATHCONF              = 192 // ***REMOVED*** int fpathconf(int fd, int name); ***REMOVED***
	SYS_GETRLIMIT              = 194 // ***REMOVED*** int getrlimit(u_int which, struct rlimit *rlp); ***REMOVED*** getrlimit __getrlimit_args int
	SYS_SETRLIMIT              = 195 // ***REMOVED*** int setrlimit(u_int which, struct rlimit *rlp); ***REMOVED*** setrlimit __setrlimit_args int
	SYS_MMAP                   = 197 // ***REMOVED*** caddr_t mmap(caddr_t addr, size_t len, int prot, int flags, int fd, int pad, off_t pos); ***REMOVED***
	SYS_LSEEK                  = 199 // ***REMOVED*** off_t lseek(int fd, int pad, off_t offset, int whence); ***REMOVED***
	SYS_TRUNCATE               = 200 // ***REMOVED*** int truncate(char *path, int pad, off_t length); ***REMOVED***
	SYS_FTRUNCATE              = 201 // ***REMOVED*** int ftruncate(int fd, int pad, off_t length); ***REMOVED***
	SYS___SYSCTL               = 202 // ***REMOVED*** int __sysctl(int *name, u_int namelen, void *old, size_t *oldlenp, void *new, size_t newlen); ***REMOVED*** __sysctl sysctl_args int
	SYS_MLOCK                  = 203 // ***REMOVED*** int mlock(const void *addr, size_t len); ***REMOVED***
	SYS_MUNLOCK                = 204 // ***REMOVED*** int munlock(const void *addr, size_t len); ***REMOVED***
	SYS_UNDELETE               = 205 // ***REMOVED*** int undelete(char *path); ***REMOVED***
	SYS_FUTIMES                = 206 // ***REMOVED*** int futimes(int fd, struct timeval *tptr); ***REMOVED***
	SYS_GETPGID                = 207 // ***REMOVED*** int getpgid(pid_t pid); ***REMOVED***
	SYS_POLL                   = 209 // ***REMOVED*** int poll(struct pollfd *fds, u_int nfds, int timeout); ***REMOVED***
	SYS___SEMCTL               = 220 // ***REMOVED*** int __semctl(int semid, int semnum, int cmd, union semun *arg); ***REMOVED***
	SYS_SEMGET                 = 221 // ***REMOVED*** int semget(key_t key, int nsems, int semflg); ***REMOVED***
	SYS_SEMOP                  = 222 // ***REMOVED*** int semop(int semid, struct sembuf *sops, u_int nsops); ***REMOVED***
	SYS_MSGCTL                 = 224 // ***REMOVED*** int msgctl(int msqid, int cmd, struct msqid_ds *buf); ***REMOVED***
	SYS_MSGGET                 = 225 // ***REMOVED*** int msgget(key_t key, int msgflg); ***REMOVED***
	SYS_MSGSND                 = 226 // ***REMOVED*** int msgsnd(int msqid, const void *msgp, size_t msgsz, int msgflg); ***REMOVED***
	SYS_MSGRCV                 = 227 // ***REMOVED*** int msgrcv(int msqid, void *msgp, size_t msgsz, long msgtyp, int msgflg); ***REMOVED***
	SYS_SHMAT                  = 228 // ***REMOVED*** caddr_t shmat(int shmid, const void *shmaddr, int shmflg); ***REMOVED***
	SYS_SHMCTL                 = 229 // ***REMOVED*** int shmctl(int shmid, int cmd, struct shmid_ds *buf); ***REMOVED***
	SYS_SHMDT                  = 230 // ***REMOVED*** int shmdt(const void *shmaddr); ***REMOVED***
	SYS_SHMGET                 = 231 // ***REMOVED*** int shmget(key_t key, size_t size, int shmflg); ***REMOVED***
	SYS_CLOCK_GETTIME          = 232 // ***REMOVED*** int clock_gettime(clockid_t clock_id, struct timespec *tp); ***REMOVED***
	SYS_CLOCK_SETTIME          = 233 // ***REMOVED*** int clock_settime(clockid_t clock_id, const struct timespec *tp); ***REMOVED***
	SYS_CLOCK_GETRES           = 234 // ***REMOVED*** int clock_getres(clockid_t clock_id, struct timespec *tp); ***REMOVED***
	SYS_NANOSLEEP              = 240 // ***REMOVED*** int nanosleep(const struct timespec *rqtp, struct timespec *rmtp); ***REMOVED***
	SYS_MINHERIT               = 250 // ***REMOVED*** int minherit(void *addr, size_t len, int inherit); ***REMOVED***
	SYS_RFORK                  = 251 // ***REMOVED*** int rfork(int flags); ***REMOVED***
	SYS_OPENBSD_POLL           = 252 // ***REMOVED*** int openbsd_poll(struct pollfd *fds, u_int nfds, int timeout); ***REMOVED***
	SYS_ISSETUGID              = 253 // ***REMOVED*** int issetugid(void); ***REMOVED***
	SYS_LCHOWN                 = 254 // ***REMOVED*** int lchown(char *path, int uid, int gid); ***REMOVED***
	SYS_LCHMOD                 = 274 // ***REMOVED*** int lchmod(char *path, mode_t mode); ***REMOVED***
	SYS_LUTIMES                = 276 // ***REMOVED*** int lutimes(char *path, struct timeval *tptr); ***REMOVED***
	SYS_EXTPREADV              = 289 // ***REMOVED*** ssize_t extpreadv(int fd, const struct iovec *iovp, int iovcnt, int flags, off_t offset); ***REMOVED***
	SYS_EXTPWRITEV             = 290 // ***REMOVED*** ssize_t extpwritev(int fd, const struct iovec *iovp, int iovcnt, int flags, off_t offset); ***REMOVED***
	SYS_FHSTATFS               = 297 // ***REMOVED*** int fhstatfs(const struct fhandle *u_fhp, struct statfs *buf); ***REMOVED***
	SYS_FHOPEN                 = 298 // ***REMOVED*** int fhopen(const struct fhandle *u_fhp, int flags); ***REMOVED***
	SYS_MODNEXT                = 300 // ***REMOVED*** int modnext(int modid); ***REMOVED***
	SYS_MODSTAT                = 301 // ***REMOVED*** int modstat(int modid, struct module_stat* stat); ***REMOVED***
	SYS_MODFNEXT               = 302 // ***REMOVED*** int modfnext(int modid); ***REMOVED***
	SYS_MODFIND                = 303 // ***REMOVED*** int modfind(const char *name); ***REMOVED***
	SYS_KLDLOAD                = 304 // ***REMOVED*** int kldload(const char *file); ***REMOVED***
	SYS_KLDUNLOAD              = 305 // ***REMOVED*** int kldunload(int fileid); ***REMOVED***
	SYS_KLDFIND                = 306 // ***REMOVED*** int kldfind(const char *file); ***REMOVED***
	SYS_KLDNEXT                = 307 // ***REMOVED*** int kldnext(int fileid); ***REMOVED***
	SYS_KLDSTAT                = 308 // ***REMOVED*** int kldstat(int fileid, struct kld_file_stat* stat); ***REMOVED***
	SYS_KLDFIRSTMOD            = 309 // ***REMOVED*** int kldfirstmod(int fileid); ***REMOVED***
	SYS_GETSID                 = 310 // ***REMOVED*** int getsid(pid_t pid); ***REMOVED***
	SYS_SETRESUID              = 311 // ***REMOVED*** int setresuid(uid_t ruid, uid_t euid, uid_t suid); ***REMOVED***
	SYS_SETRESGID              = 312 // ***REMOVED*** int setresgid(gid_t rgid, gid_t egid, gid_t sgid); ***REMOVED***
	SYS_AIO_RETURN             = 314 // ***REMOVED*** int aio_return(struct aiocb *aiocbp); ***REMOVED***
	SYS_AIO_SUSPEND            = 315 // ***REMOVED*** int aio_suspend(struct aiocb * const * aiocbp, int nent, const struct timespec *timeout); ***REMOVED***
	SYS_AIO_CANCEL             = 316 // ***REMOVED*** int aio_cancel(int fd, struct aiocb *aiocbp); ***REMOVED***
	SYS_AIO_ERROR              = 317 // ***REMOVED*** int aio_error(struct aiocb *aiocbp); ***REMOVED***
	SYS_AIO_READ               = 318 // ***REMOVED*** int aio_read(struct aiocb *aiocbp); ***REMOVED***
	SYS_AIO_WRITE              = 319 // ***REMOVED*** int aio_write(struct aiocb *aiocbp); ***REMOVED***
	SYS_LIO_LISTIO             = 320 // ***REMOVED*** int lio_listio(int mode, struct aiocb * const *acb_list, int nent, struct sigevent *sig); ***REMOVED***
	SYS_YIELD                  = 321 // ***REMOVED*** int yield(void); ***REMOVED***
	SYS_MLOCKALL               = 324 // ***REMOVED*** int mlockall(int how); ***REMOVED***
	SYS_MUNLOCKALL             = 325 // ***REMOVED*** int munlockall(void); ***REMOVED***
	SYS___GETCWD               = 326 // ***REMOVED*** int __getcwd(u_char *buf, u_int buflen); ***REMOVED***
	SYS_SCHED_SETPARAM         = 327 // ***REMOVED*** int sched_setparam (pid_t pid, const struct sched_param *param); ***REMOVED***
	SYS_SCHED_GETPARAM         = 328 // ***REMOVED*** int sched_getparam (pid_t pid, struct sched_param *param); ***REMOVED***
	SYS_SCHED_SETSCHEDULER     = 329 // ***REMOVED*** int sched_setscheduler (pid_t pid, int policy, const struct sched_param *param); ***REMOVED***
	SYS_SCHED_GETSCHEDULER     = 330 // ***REMOVED*** int sched_getscheduler (pid_t pid); ***REMOVED***
	SYS_SCHED_YIELD            = 331 // ***REMOVED*** int sched_yield (void); ***REMOVED***
	SYS_SCHED_GET_PRIORITY_MAX = 332 // ***REMOVED*** int sched_get_priority_max (int policy); ***REMOVED***
	SYS_SCHED_GET_PRIORITY_MIN = 333 // ***REMOVED*** int sched_get_priority_min (int policy); ***REMOVED***
	SYS_SCHED_RR_GET_INTERVAL  = 334 // ***REMOVED*** int sched_rr_get_interval (pid_t pid, struct timespec *interval); ***REMOVED***
	SYS_UTRACE                 = 335 // ***REMOVED*** int utrace(const void *addr, size_t len); ***REMOVED***
	SYS_KLDSYM                 = 337 // ***REMOVED*** int kldsym(int fileid, int cmd, void *data); ***REMOVED***
	SYS_JAIL                   = 338 // ***REMOVED*** int jail(struct jail *jail); ***REMOVED***
	SYS_SIGPROCMASK            = 340 // ***REMOVED*** int sigprocmask(int how, const sigset_t *set, sigset_t *oset); ***REMOVED***
	SYS_SIGSUSPEND             = 341 // ***REMOVED*** int sigsuspend(const sigset_t *sigmask); ***REMOVED***
	SYS_SIGACTION              = 342 // ***REMOVED*** int sigaction(int sig, const struct sigaction *act, struct sigaction *oact); ***REMOVED***
	SYS_SIGPENDING             = 343 // ***REMOVED*** int sigpending(sigset_t *set); ***REMOVED***
	SYS_SIGRETURN              = 344 // ***REMOVED*** int sigreturn(ucontext_t *sigcntxp); ***REMOVED***
	SYS_SIGTIMEDWAIT           = 345 // ***REMOVED*** int sigtimedwait(const sigset_t *set,siginfo_t *info, const struct timespec *timeout); ***REMOVED***
	SYS_SIGWAITINFO            = 346 // ***REMOVED*** int sigwaitinfo(const sigset_t *set,siginfo_t *info); ***REMOVED***
	SYS___ACL_GET_FILE         = 347 // ***REMOVED*** int __acl_get_file(const char *path, acl_type_t type, struct acl *aclp); ***REMOVED***
	SYS___ACL_SET_FILE         = 348 // ***REMOVED*** int __acl_set_file(const char *path, acl_type_t type, struct acl *aclp); ***REMOVED***
	SYS___ACL_GET_FD           = 349 // ***REMOVED*** int __acl_get_fd(int filedes, acl_type_t type, struct acl *aclp); ***REMOVED***
	SYS___ACL_SET_FD           = 350 // ***REMOVED*** int __acl_set_fd(int filedes, acl_type_t type, struct acl *aclp); ***REMOVED***
	SYS___ACL_DELETE_FILE      = 351 // ***REMOVED*** int __acl_delete_file(const char *path, acl_type_t type); ***REMOVED***
	SYS___ACL_DELETE_FD        = 352 // ***REMOVED*** int __acl_delete_fd(int filedes, acl_type_t type); ***REMOVED***
	SYS___ACL_ACLCHECK_FILE    = 353 // ***REMOVED*** int __acl_aclcheck_file(const char *path, acl_type_t type, struct acl *aclp); ***REMOVED***
	SYS___ACL_ACLCHECK_FD      = 354 // ***REMOVED*** int __acl_aclcheck_fd(int filedes, acl_type_t type, struct acl *aclp); ***REMOVED***
	SYS_EXTATTRCTL             = 355 // ***REMOVED*** int extattrctl(const char *path, int cmd, const char *filename, int attrnamespace, const char *attrname); ***REMOVED***
	SYS_EXTATTR_SET_FILE       = 356 // ***REMOVED*** int extattr_set_file(const char *path, int attrnamespace, const char *attrname, void *data, size_t nbytes); ***REMOVED***
	SYS_EXTATTR_GET_FILE       = 357 // ***REMOVED*** int extattr_get_file(const char *path, int attrnamespace, const char *attrname, void *data, size_t nbytes); ***REMOVED***
	SYS_EXTATTR_DELETE_FILE    = 358 // ***REMOVED*** int extattr_delete_file(const char *path, int attrnamespace, const char *attrname); ***REMOVED***
	SYS_AIO_WAITCOMPLETE       = 359 // ***REMOVED*** int aio_waitcomplete(struct aiocb **aiocbp, struct timespec *timeout); ***REMOVED***
	SYS_GETRESUID              = 360 // ***REMOVED*** int getresuid(uid_t *ruid, uid_t *euid, uid_t *suid); ***REMOVED***
	SYS_GETRESGID              = 361 // ***REMOVED*** int getresgid(gid_t *rgid, gid_t *egid, gid_t *sgid); ***REMOVED***
	SYS_KQUEUE                 = 362 // ***REMOVED*** int kqueue(void); ***REMOVED***
	SYS_KEVENT                 = 363 // ***REMOVED*** int kevent(int fd, const struct kevent *changelist, int nchanges, struct kevent *eventlist, int nevents, const struct timespec *timeout); ***REMOVED***
	SYS_KENV                   = 390 // ***REMOVED*** int kenv(int what, const char *name, char *value, int len); ***REMOVED***
	SYS_LCHFLAGS               = 391 // ***REMOVED*** int lchflags(const char *path, u_long flags); ***REMOVED***
	SYS_UUIDGEN                = 392 // ***REMOVED*** int uuidgen(struct uuid *store, int count); ***REMOVED***
	SYS_SENDFILE               = 393 // ***REMOVED*** int sendfile(int fd, int s, off_t offset, size_t nbytes, struct sf_hdtr *hdtr, off_t *sbytes, int flags); ***REMOVED***
	SYS_VARSYM_SET             = 450 // ***REMOVED*** int varsym_set(int level, const char *name, const char *data); ***REMOVED***
	SYS_VARSYM_GET             = 451 // ***REMOVED*** int varsym_get(int mask, const char *wild, char *buf, int bufsize); ***REMOVED***
	SYS_VARSYM_LIST            = 452 // ***REMOVED*** int varsym_list(int level, char *buf, int maxsize, int *marker); ***REMOVED***
	SYS_EXEC_SYS_REGISTER      = 465 // ***REMOVED*** int exec_sys_register(void *entry); ***REMOVED***
	SYS_EXEC_SYS_UNREGISTER    = 466 // ***REMOVED*** int exec_sys_unregister(int id); ***REMOVED***
	SYS_SYS_CHECKPOINT         = 467 // ***REMOVED*** int sys_checkpoint(int type, int fd, pid_t pid, int retval); ***REMOVED***
	SYS_MOUNTCTL               = 468 // ***REMOVED*** int mountctl(const char *path, int op, int fd, const void *ctl, int ctllen, void *buf, int buflen); ***REMOVED***
	SYS_UMTX_SLEEP             = 469 // ***REMOVED*** int umtx_sleep(volatile const int *ptr, int value, int timeout); ***REMOVED***
	SYS_UMTX_WAKEUP            = 470 // ***REMOVED*** int umtx_wakeup(volatile const int *ptr, int count); ***REMOVED***
	SYS_JAIL_ATTACH            = 471 // ***REMOVED*** int jail_attach(int jid); ***REMOVED***
	SYS_SET_TLS_AREA           = 472 // ***REMOVED*** int set_tls_area(int which, struct tls_info *info, size_t infosize); ***REMOVED***
	SYS_GET_TLS_AREA           = 473 // ***REMOVED*** int get_tls_area(int which, struct tls_info *info, size_t infosize); ***REMOVED***
	SYS_CLOSEFROM              = 474 // ***REMOVED*** int closefrom(int fd); ***REMOVED***
	SYS_STAT                   = 475 // ***REMOVED*** int stat(const char *path, struct stat *ub); ***REMOVED***
	SYS_FSTAT                  = 476 // ***REMOVED*** int fstat(int fd, struct stat *sb); ***REMOVED***
	SYS_LSTAT                  = 477 // ***REMOVED*** int lstat(const char *path, struct stat *ub); ***REMOVED***
	SYS_FHSTAT                 = 478 // ***REMOVED*** int fhstat(const struct fhandle *u_fhp, struct stat *sb); ***REMOVED***
	SYS_GETDIRENTRIES          = 479 // ***REMOVED*** int getdirentries(int fd, char *buf, u_int count, long *basep); ***REMOVED***
	SYS_GETDENTS               = 480 // ***REMOVED*** int getdents(int fd, char *buf, size_t count); ***REMOVED***
	SYS_USCHED_SET             = 481 // ***REMOVED*** int usched_set(pid_t pid, int cmd, void *data, int bytes); ***REMOVED***
	SYS_EXTACCEPT              = 482 // ***REMOVED*** int extaccept(int s, int flags, caddr_t name, int *anamelen); ***REMOVED***
	SYS_EXTCONNECT             = 483 // ***REMOVED*** int extconnect(int s, int flags, caddr_t name, int namelen); ***REMOVED***
	SYS_MCONTROL               = 485 // ***REMOVED*** int mcontrol(void *addr, size_t len, int behav, off_t value); ***REMOVED***
	SYS_VMSPACE_CREATE         = 486 // ***REMOVED*** int vmspace_create(void *id, int type, void *data); ***REMOVED***
	SYS_VMSPACE_DESTROY        = 487 // ***REMOVED*** int vmspace_destroy(void *id); ***REMOVED***
	SYS_VMSPACE_CTL            = 488 // ***REMOVED*** int vmspace_ctl(void *id, int cmd, 		struct trapframe *tframe,	struct vextframe *vframe); ***REMOVED***
	SYS_VMSPACE_MMAP           = 489 // ***REMOVED*** int vmspace_mmap(void *id, void *addr, size_t len, int prot, int flags, int fd, off_t offset); ***REMOVED***
	SYS_VMSPACE_MUNMAP         = 490 // ***REMOVED*** int vmspace_munmap(void *id, void *addr,	size_t len); ***REMOVED***
	SYS_VMSPACE_MCONTROL       = 491 // ***REMOVED*** int vmspace_mcontrol(void *id, void *addr, 	size_t len, int behav, off_t value); ***REMOVED***
	SYS_VMSPACE_PREAD          = 492 // ***REMOVED*** ssize_t vmspace_pread(void *id, void *buf, size_t nbyte, int flags, off_t offset); ***REMOVED***
	SYS_VMSPACE_PWRITE         = 493 // ***REMOVED*** ssize_t vmspace_pwrite(void *id, const void *buf, size_t nbyte, int flags, off_t offset); ***REMOVED***
	SYS_EXTEXIT                = 494 // ***REMOVED*** void extexit(int how, int status, void *addr); ***REMOVED***
	SYS_LWP_CREATE             = 495 // ***REMOVED*** int lwp_create(struct lwp_params *params); ***REMOVED***
	SYS_LWP_GETTID             = 496 // ***REMOVED*** lwpid_t lwp_gettid(void); ***REMOVED***
	SYS_LWP_KILL               = 497 // ***REMOVED*** int lwp_kill(pid_t pid, lwpid_t tid, int signum); ***REMOVED***
	SYS_LWP_RTPRIO             = 498 // ***REMOVED*** int lwp_rtprio(int function, pid_t pid, lwpid_t tid, struct rtprio *rtp); ***REMOVED***
	SYS_PSELECT                = 499 // ***REMOVED*** int pselect(int nd, fd_set *in, fd_set *ou, fd_set *ex, const struct timespec *ts,    const sigset_t *sigmask); ***REMOVED***
	SYS_STATVFS                = 500 // ***REMOVED*** int statvfs(const char *path, struct statvfs *buf); ***REMOVED***
	SYS_FSTATVFS               = 501 // ***REMOVED*** int fstatvfs(int fd, struct statvfs *buf); ***REMOVED***
	SYS_FHSTATVFS              = 502 // ***REMOVED*** int fhstatvfs(const struct fhandle *u_fhp, struct statvfs *buf); ***REMOVED***
	SYS_GETVFSSTAT             = 503 // ***REMOVED*** int getvfsstat(struct statfs *buf,          struct statvfs *vbuf, long vbufsize, int flags); ***REMOVED***
	SYS_OPENAT                 = 504 // ***REMOVED*** int openat(int fd, char *path, int flags, int mode); ***REMOVED***
	SYS_FSTATAT                = 505 // ***REMOVED*** int fstatat(int fd, char *path, 	struct stat *sb, int flags); ***REMOVED***
	SYS_FCHMODAT               = 506 // ***REMOVED*** int fchmodat(int fd, char *path, int mode, int flags); ***REMOVED***
	SYS_FCHOWNAT               = 507 // ***REMOVED*** int fchownat(int fd, char *path, int uid, int gid, int flags); ***REMOVED***
	SYS_UNLINKAT               = 508 // ***REMOVED*** int unlinkat(int fd, char *path, int flags); ***REMOVED***
	SYS_FACCESSAT              = 509 // ***REMOVED*** int faccessat(int fd, char *path, int amode, int flags); ***REMOVED***
	SYS_MQ_OPEN                = 510 // ***REMOVED*** mqd_t mq_open(const char * name, int oflag, mode_t mode, struct mq_attr *attr); ***REMOVED***
	SYS_MQ_CLOSE               = 511 // ***REMOVED*** int mq_close(mqd_t mqdes); ***REMOVED***
	SYS_MQ_UNLINK              = 512 // ***REMOVED*** int mq_unlink(const char *name); ***REMOVED***
	SYS_MQ_GETATTR             = 513 // ***REMOVED*** int mq_getattr(mqd_t mqdes, struct mq_attr *mqstat); ***REMOVED***
	SYS_MQ_SETATTR             = 514 // ***REMOVED*** int mq_setattr(mqd_t mqdes, const struct mq_attr *mqstat, struct mq_attr *omqstat); ***REMOVED***
	SYS_MQ_NOTIFY              = 515 // ***REMOVED*** int mq_notify(mqd_t mqdes, const struct sigevent *notification); ***REMOVED***
	SYS_MQ_SEND                = 516 // ***REMOVED*** int mq_send(mqd_t mqdes, const char *msg_ptr, size_t msg_len, unsigned msg_prio); ***REMOVED***
	SYS_MQ_RECEIVE             = 517 // ***REMOVED*** ssize_t mq_receive(mqd_t mqdes, char *msg_ptr, size_t msg_len, unsigned *msg_prio); ***REMOVED***
	SYS_MQ_TIMEDSEND           = 518 // ***REMOVED*** int mq_timedsend(mqd_t mqdes, const char *msg_ptr, size_t msg_len, unsigned msg_prio, const struct timespec *abs_timeout); ***REMOVED***
	SYS_MQ_TIMEDRECEIVE        = 519 // ***REMOVED*** ssize_t mq_timedreceive(mqd_t mqdes, char *msg_ptr, size_t msg_len, unsigned *msg_prio, const struct timespec *abs_timeout); ***REMOVED***
	SYS_IOPRIO_SET             = 520 // ***REMOVED*** int ioprio_set(int which, int who, int prio); ***REMOVED***
	SYS_IOPRIO_GET             = 521 // ***REMOVED*** int ioprio_get(int which, int who); ***REMOVED***
	SYS_CHROOT_KERNEL          = 522 // ***REMOVED*** int chroot_kernel(char *path); ***REMOVED***
	SYS_RENAMEAT               = 523 // ***REMOVED*** int renameat(int oldfd, char *old, int newfd, char *new); ***REMOVED***
	SYS_MKDIRAT                = 524 // ***REMOVED*** int mkdirat(int fd, char *path, mode_t mode); ***REMOVED***
	SYS_MKFIFOAT               = 525 // ***REMOVED*** int mkfifoat(int fd, char *path, mode_t mode); ***REMOVED***
	SYS_MKNODAT                = 526 // ***REMOVED*** int mknodat(int fd, char *path, mode_t mode, dev_t dev); ***REMOVED***
	SYS_READLINKAT             = 527 // ***REMOVED*** int readlinkat(int fd, char *path, char *buf, size_t bufsize); ***REMOVED***
	SYS_SYMLINKAT              = 528 // ***REMOVED*** int symlinkat(char *path1, int fd, char *path2); ***REMOVED***
	SYS_SWAPOFF                = 529 // ***REMOVED*** int swapoff(char *name); ***REMOVED***
	SYS_VQUOTACTL              = 530 // ***REMOVED*** int vquotactl(const char *path, struct plistref *pref); ***REMOVED***
	SYS_LINKAT                 = 531 // ***REMOVED*** int linkat(int fd1, char *path1, int fd2, char *path2, int flags); ***REMOVED***
	SYS_EACCESS                = 532 // ***REMOVED*** int eaccess(char *path, int flags); ***REMOVED***
	SYS_LPATHCONF              = 533 // ***REMOVED*** int lpathconf(char *path, int name); ***REMOVED***
	SYS_VMM_GUEST_CTL          = 534 // ***REMOVED*** int vmm_guest_ctl(int op, struct vmm_guest_options *options); ***REMOVED***
	SYS_VMM_GUEST_SYNC_ADDR    = 535 // ***REMOVED*** int vmm_guest_sync_addr(long *dstaddr, long *srcaddr); ***REMOVED***
	SYS_PROCCTL                = 536 // ***REMOVED*** int procctl(idtype_t idtype, id_t id, int cmd, void *data); ***REMOVED***
	SYS_CHFLAGSAT              = 537 // ***REMOVED*** int chflagsat(int fd, const char *path, u_long flags, int atflags);***REMOVED***
	SYS_PIPE2                  = 538 // ***REMOVED*** int pipe2(int *fildes, int flags); ***REMOVED***
	SYS_UTIMENSAT              = 539 // ***REMOVED*** int utimensat(int fd, const char *path, const struct timespec *ts, int flags); ***REMOVED***
	SYS_FUTIMENS               = 540 // ***REMOVED*** int futimens(int fd, const struct timespec *ts); ***REMOVED***
	SYS_ACCEPT4                = 541 // ***REMOVED*** int accept4(int s, caddr_t name, int *anamelen, int flags); ***REMOVED***
	SYS_LWP_SETNAME            = 542 // ***REMOVED*** int lwp_setname(lwpid_t tid, const char *name); ***REMOVED***
	SYS_PPOLL                  = 543 // ***REMOVED*** int ppoll(struct pollfd *fds, u_int nfds, const struct timespec *ts, const sigset_t *sigmask); ***REMOVED***
	SYS_LWP_SETAFFINITY        = 544 // ***REMOVED*** int lwp_setaffinity(pid_t pid, lwpid_t tid, const cpumask_t *mask); ***REMOVED***
	SYS_LWP_GETAFFINITY        = 545 // ***REMOVED*** int lwp_getaffinity(pid_t pid, lwpid_t tid, cpumask_t *mask); ***REMOVED***
	SYS_LWP_CREATE2            = 546 // ***REMOVED*** int lwp_create2(struct lwp_params *params, const cpumask_t *mask); ***REMOVED***
	SYS_GETCPUCLOCKID          = 547 // ***REMOVED*** int getcpuclockid(pid_t pid, lwpid_t lwp_id, clockid_t *clock_id); ***REMOVED***
	SYS_WAIT6                  = 548 // ***REMOVED*** int wait6(idtype_t idtype, id_t id, int *status, int options, struct __wrusage *wrusage, siginfo_t *info); ***REMOVED***
	SYS_LWP_GETNAME            = 549 // ***REMOVED*** int lwp_getname(lwpid_t tid, char *name, size_t len); ***REMOVED***
	SYS_GETRANDOM              = 550 // ***REMOVED*** ssize_t getrandom(void *buf, size_t len, unsigned flags); ***REMOVED***
	SYS___REALPATH             = 551 // ***REMOVED*** ssize_t __realpath(const char *path, char *buf, size_t len); ***REMOVED***
)
