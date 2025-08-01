// cgo -godefs -objdir=/tmp/ppc/cgo -- -Wall -Werror -static -I/tmp/ppc/include linux/types.go | go run mkpost.go
// Code generated by the command above; see README.md. DO NOT EDIT.

//go:build ppc && linux

package unix

const (
	SizeofPtr  = 0x4
	SizeofLong = 0x4
)

type (
	_C_long int32
)

type Timespec struct {
	Sec  int32
	Nsec int32
}

type Timeval struct {
	Sec  int32
	Usec int32
}

type Timex struct {
	Modes     uint32
	Offset    int32
	Freq      int32
	Maxerror  int32
	Esterror  int32
	Status    int32
	Constant  int32
	Precision int32
	Tolerance int32
	Time      Timeval
	Tick      int32
	Ppsfreq   int32
	Jitter    int32
	Shift     int32
	Stabil    int32
	Jitcnt    int32
	Calcnt    int32
	Errcnt    int32
	Stbcnt    int32
	Tai       int32
	_         [44]byte
}

type Time_t int32

type Tms struct {
	Utime  int32
	Stime  int32
	Cutime int32
	Cstime int32
}

type Utimbuf struct {
	Actime  int32
	Modtime int32
}

type Rusage struct {
	Utime    Timeval
	Stime    Timeval
	Maxrss   int32
	Ixrss    int32
	Idrss    int32
	Isrss    int32
	Minflt   int32
	Majflt   int32
	Nswap    int32
	Inblock  int32
	Oublock  int32
	Msgsnd   int32
	Msgrcv   int32
	Nsignals int32
	Nvcsw    int32
	Nivcsw   int32
}

type Stat_t struct {
	Dev     uint64
	Ino     uint64
	Mode    uint32
	Nlink   uint32
	Uid     uint32
	Gid     uint32
	Rdev    uint64
	_       uint16
	_       [4]byte
	Size    int64
	Blksize int32
	_       [4]byte
	Blocks  int64
	Atim    Timespec
	Mtim    Timespec
	Ctim    Timespec
	_       uint32
	_       uint32
}

type Dirent struct {
	Ino    uint64
	Off    int64
	Reclen uint16
	Type   uint8
	Name   [256]uint8
	_      [5]byte
}

type Flock_t struct {
	Type   int16
	Whence int16
	_      [4]byte
	Start  int64
	Len    int64
	Pid    int32
	_      [4]byte
}

type DmNameList struct {
	Dev  uint64
	Next uint32
	Name [0]byte
	_    [4]byte
}

const (
	FADV_DONTNEED = 0x4
	FADV_NOREUSE  = 0x5
)

type RawSockaddrNFCLLCP struct {
	Sa_family        uint16
	Dev_idx          uint32
	Target_idx       uint32
	Nfc_protocol     uint32
	Dsap             uint8
	Ssap             uint8
	Service_name     [63]uint8
	Service_name_len uint32
}

type RawSockaddr struct {
	Family uint16
	Data   [14]uint8
}

type RawSockaddrAny struct {
	Addr RawSockaddr
	Pad  [96]uint8
}

type Iovec struct {
	Base *byte
	Len  uint32
}

type Msghdr struct {
	Name       *byte
	Namelen    uint32
	Iov        *Iovec
	Iovlen     uint32
	Control    *byte
	Controllen uint32
	Flags      int32
}

type Cmsghdr struct {
	Len   uint32
	Level int32
	Type  int32
}

type ifreq struct {
	Ifrn [16]byte
	Ifru [16]byte
}

const (
	SizeofSockaddrNFCLLCP = 0x58
	SizeofIovec           = 0x8
	SizeofMsghdr          = 0x1c
	SizeofCmsghdr         = 0xc
)

const (
	SizeofSockFprog = 0x8
)

type PtraceRegs struct {
	Gpr       [32]uint32
	Nip       uint32
	Msr       uint32
	Orig_gpr3 uint32
	Ctr       uint32
	Link      uint32
	Xer       uint32
	Ccr       uint32
	Mq        uint32
	Trap      uint32
	Dar       uint32
	Dsisr     uint32
	Result    uint32
}

type FdSet struct {
	Bits [32]int32
}

type Sysinfo_t struct {
	Uptime    int32
	Loads     [3]uint32
	Totalram  uint32
	Freeram   uint32
	Sharedram uint32
	Bufferram uint32
	Totalswap uint32
	Freeswap  uint32
	Procs     uint16
	Pad       uint16
	Totalhigh uint32
	Freehigh  uint32
	Unit      uint32
	_         [8]uint8
}

type Ustat_t struct {
	Tfree  int32
	Tinode uint32
	Fname  [6]uint8
	Fpack  [6]uint8
}

type EpollEvent struct {
	Events uint32
	_      int32
	Fd     int32
	Pad    int32
}

const (
	OPEN_TREE_CLOEXEC = 0x80000
)

const (
	POLLRDHUP = 0x2000
)

type Sigset_t struct {
	Val [32]uint32
}

const _C__NSIG = 0x41

const (
	SIG_BLOCK   = 0x0
	SIG_UNBLOCK = 0x1
	SIG_SETMASK = 0x2
)

type Siginfo struct {
	Signo int32
	Errno int32
	Code  int32
	_     [116]byte
}

type Termios struct {
	Iflag  uint32
	Oflag  uint32
	Cflag  uint32
	Lflag  uint32
	Cc     [19]uint8
	Line   uint8
	Ispeed uint32
	Ospeed uint32
}

type Taskstats struct {
	Version                   uint16
	Ac_exitcode               uint32
	Ac_flag                   uint8
	Ac_nice                   uint8
	_                         [4]byte
	Cpu_count                 uint64
	Cpu_delay_total           uint64
	Cpu_delay_max             uint64
	Cpu_delay_min             uint64
	Blkio_count               uint64
	Blkio_delay_total         uint64
	Blkio_delay_max           uint64
	Blkio_delay_min           uint64
	Swapin_count              uint64
	Swapin_delay_total        uint64
	Swapin_delay_max          uint64
	Swapin_delay_min          uint64
	Cpu_run_real_total        uint64
	Cpu_run_virtual_total     uint64
	Ac_comm                   [32]uint8
	Ac_sched                  uint8
	Ac_pad                    [3]uint8
	_                         [4]byte
	Ac_uid                    uint32
	Ac_gid                    uint32
	Ac_pid                    uint32
	Ac_ppid                   uint32
	Ac_btime                  uint32
	_                         [4]byte
	Ac_etime                  uint64
	Ac_utime                  uint64
	Ac_stime                  uint64
	Ac_minflt                 uint64
	Ac_majflt                 uint64
	Coremem                   uint64
	Virtmem                   uint64
	Hiwater_rss               uint64
	Hiwater_vm                uint64
	Read_char                 uint64
	Write_char                uint64
	Read_syscalls             uint64
	Write_syscalls            uint64
	Read_bytes                uint64
	Write_bytes               uint64
	Cancelled_write_bytes     uint64
	Nvcsw                     uint64
	Nivcsw                    uint64
	Ac_utimescaled            uint64
	Ac_stimescaled            uint64
	Cpu_scaled_run_real_total uint64
	Freepages_count           uint64
	Freepages_delay_total     uint64
	Freepages_delay_max       uint64
	Freepages_delay_min       uint64
	Thrashing_count           uint64
	Thrashing_delay_total     uint64
	Thrashing_delay_max       uint64
	Thrashing_delay_min       uint64
	Ac_btime64                uint64
	Compact_count             uint64
	Compact_delay_total       uint64
	Compact_delay_max         uint64
	Compact_delay_min         uint64
	Ac_tgid                   uint32
	_                         [4]byte
	Ac_tgetime                uint64
	Ac_exe_dev                uint64
	Ac_exe_inode              uint64
	Wpcopy_count              uint64
	Wpcopy_delay_total        uint64
	Wpcopy_delay_max          uint64
	Wpcopy_delay_min          uint64
	Irq_count                 uint64
	Irq_delay_total           uint64
	Irq_delay_max             uint64
	Irq_delay_min             uint64
}

type cpuMask uint32

const (
	_NCPUBITS = 0x20
)

const (
	CBitFieldMaskBit0  = 0x8000000000000000
	CBitFieldMaskBit1  = 0x4000000000000000
	CBitFieldMaskBit2  = 0x2000000000000000
	CBitFieldMaskBit3  = 0x1000000000000000
	CBitFieldMaskBit4  = 0x800000000000000
	CBitFieldMaskBit5  = 0x400000000000000
	CBitFieldMaskBit6  = 0x200000000000000
	CBitFieldMaskBit7  = 0x100000000000000
	CBitFieldMaskBit8  = 0x80000000000000
	CBitFieldMaskBit9  = 0x40000000000000
	CBitFieldMaskBit10 = 0x20000000000000
	CBitFieldMaskBit11 = 0x10000000000000
	CBitFieldMaskBit12 = 0x8000000000000
	CBitFieldMaskBit13 = 0x4000000000000
	CBitFieldMaskBit14 = 0x2000000000000
	CBitFieldMaskBit15 = 0x1000000000000
	CBitFieldMaskBit16 = 0x800000000000
	CBitFieldMaskBit17 = 0x400000000000
	CBitFieldMaskBit18 = 0x200000000000
	CBitFieldMaskBit19 = 0x100000000000
	CBitFieldMaskBit20 = 0x80000000000
	CBitFieldMaskBit21 = 0x40000000000
	CBitFieldMaskBit22 = 0x20000000000
	CBitFieldMaskBit23 = 0x10000000000
	CBitFieldMaskBit24 = 0x8000000000
	CBitFieldMaskBit25 = 0x4000000000
	CBitFieldMaskBit26 = 0x2000000000
	CBitFieldMaskBit27 = 0x1000000000
	CBitFieldMaskBit28 = 0x800000000
	CBitFieldMaskBit29 = 0x400000000
	CBitFieldMaskBit30 = 0x200000000
	CBitFieldMaskBit31 = 0x100000000
	CBitFieldMaskBit32 = 0x80000000
	CBitFieldMaskBit33 = 0x40000000
	CBitFieldMaskBit34 = 0x20000000
	CBitFieldMaskBit35 = 0x10000000
	CBitFieldMaskBit36 = 0x8000000
	CBitFieldMaskBit37 = 0x4000000
	CBitFieldMaskBit38 = 0x2000000
	CBitFieldMaskBit39 = 0x1000000
	CBitFieldMaskBit40 = 0x800000
	CBitFieldMaskBit41 = 0x400000
	CBitFieldMaskBit42 = 0x200000
	CBitFieldMaskBit43 = 0x100000
	CBitFieldMaskBit44 = 0x80000
	CBitFieldMaskBit45 = 0x40000
	CBitFieldMaskBit46 = 0x20000
	CBitFieldMaskBit47 = 0x10000
	CBitFieldMaskBit48 = 0x8000
	CBitFieldMaskBit49 = 0x4000
	CBitFieldMaskBit50 = 0x2000
	CBitFieldMaskBit51 = 0x1000
	CBitFieldMaskBit52 = 0x800
	CBitFieldMaskBit53 = 0x400
	CBitFieldMaskBit54 = 0x200
	CBitFieldMaskBit55 = 0x100
	CBitFieldMaskBit56 = 0x80
	CBitFieldMaskBit57 = 0x40
	CBitFieldMaskBit58 = 0x20
	CBitFieldMaskBit59 = 0x10
	CBitFieldMaskBit60 = 0x8
	CBitFieldMaskBit61 = 0x4
	CBitFieldMaskBit62 = 0x2
	CBitFieldMaskBit63 = 0x1
)

type SockaddrStorage struct {
	Family uint16
	Data   [122]byte
	_      uint32
}

type HDGeometry struct {
	Heads     uint8
	Sectors   uint8
	Cylinders uint16
	Start     uint32
}

type Statfs_t struct {
	Type    int32
	Bsize   int32
	Blocks  uint64
	Bfree   uint64
	Bavail  uint64
	Files   uint64
	Ffree   uint64
	Fsid    Fsid
	Namelen int32
	Frsize  int32
	Flags   int32
	Spare   [4]int32
	_       [4]byte
}

type TpacketHdr struct {
	Status  uint32
	Len     uint32
	Snaplen uint32
	Mac     uint16
	Net     uint16
	Sec     uint32
	Usec    uint32
}

const (
	SizeofTpacketHdr = 0x18
)

type RTCPLLInfo struct {
	Ctrl    int32
	Value   int32
	Max     int32
	Min     int32
	Posmult int32
	Negmult int32
	Clock   int32
}

type BlkpgPartition struct {
	Start   int64
	Length  int64
	Pno     int32
	Devname [64]uint8
	Volname [64]uint8
	_       [4]byte
}

const (
	BLKPG = 0x20001269
)

type CryptoUserAlg struct {
	Name        [64]uint8
	Driver_name [64]uint8
	Module_name [64]uint8
	Type        uint32
	Mask        uint32
	Refcnt      uint32
	Flags       uint32
}

type CryptoStatAEAD struct {
	Type         [64]uint8
	Encrypt_cnt  uint64
	Encrypt_tlen uint64
	Decrypt_cnt  uint64
	Decrypt_tlen uint64
	Err_cnt      uint64
}

type CryptoStatAKCipher struct {
	Type         [64]uint8
	Encrypt_cnt  uint64
	Encrypt_tlen uint64
	Decrypt_cnt  uint64
	Decrypt_tlen uint64
	Verify_cnt   uint64
	Sign_cnt     uint64
	Err_cnt      uint64
}

type CryptoStatCipher struct {
	Type         [64]uint8
	Encrypt_cnt  uint64
	Encrypt_tlen uint64
	Decrypt_cnt  uint64
	Decrypt_tlen uint64
	Err_cnt      uint64
}

type CryptoStatCompress struct {
	Type            [64]uint8
	Compress_cnt    uint64
	Compress_tlen   uint64
	Decompress_cnt  uint64
	Decompress_tlen uint64
	Err_cnt         uint64
}

type CryptoStatHash struct {
	Type      [64]uint8
	Hash_cnt  uint64
	Hash_tlen uint64
	Err_cnt   uint64
}

type CryptoStatKPP struct {
	Type                      [64]uint8
	Setsecret_cnt             uint64
	Generate_public_key_cnt   uint64
	Compute_shared_secret_cnt uint64
	Err_cnt                   uint64
}

type CryptoStatRNG struct {
	Type          [64]uint8
	Generate_cnt  uint64
	Generate_tlen uint64
	Seed_cnt      uint64
	Err_cnt       uint64
}

type CryptoStatLarval struct {
	Type [64]uint8
}

type CryptoReportLarval struct {
	Type [64]uint8
}

type CryptoReportHash struct {
	Type       [64]uint8
	Blocksize  uint32
	Digestsize uint32
}

type CryptoReportCipher struct {
	Type        [64]uint8
	Blocksize   uint32
	Min_keysize uint32
	Max_keysize uint32
}

type CryptoReportBlkCipher struct {
	Type        [64]uint8
	Geniv       [64]uint8
	Blocksize   uint32
	Min_keysize uint32
	Max_keysize uint32
	Ivsize      uint32
}

type CryptoReportAEAD struct {
	Type        [64]uint8
	Geniv       [64]uint8
	Blocksize   uint32
	Maxauthsize uint32
	Ivsize      uint32
}

type CryptoReportComp struct {
	Type [64]uint8
}

type CryptoReportRNG struct {
	Type     [64]uint8
	Seedsize uint32
}

type CryptoReportAKCipher struct {
	Type [64]uint8
}

type CryptoReportKPP struct {
	Type [64]uint8
}

type CryptoReportAcomp struct {
	Type [64]uint8
}

type LoopInfo struct {
	Number           int32
	Device           uint32
	Inode            uint32
	Rdevice          uint32
	Offset           int32
	Encrypt_type     int32
	Encrypt_key_size int32
	Flags            int32
	Name             [64]uint8
	Encrypt_key      [32]uint8
	Init             [2]uint32
	Reserved         [4]uint8
}

type TIPCSubscr struct {
	Seq     TIPCServiceRange
	Timeout uint32
	Filter  uint32
	Handle  [8]uint8
}

type TIPCSIOCLNReq struct {
	Peer     uint32
	Id       uint32
	Linkname [68]uint8
}

type TIPCSIOCNodeIDReq struct {
	Peer uint32
	Id   [16]uint8
}

type PPSKInfo struct {
	Assert_sequence uint32
	Clear_sequence  uint32
	Assert_tu       PPSKTime
	Clear_tu        PPSKTime
	Current_mode    int32
	_               [4]byte
}

const (
	PPS_GETPARAMS = 0x400470a1
	PPS_SETPARAMS = 0x800470a2
	PPS_GETCAP    = 0x400470a3
	PPS_FETCH     = 0xc00470a4
)

const (
	PIDFD_NONBLOCK = 0x800
)

type SysvIpcPerm struct {
	Key  int32
	Uid  uint32
	Gid  uint32
	Cuid uint32
	Cgid uint32
	Mode uint32
	Seq  uint32
	_    uint32
	_    uint64
	_    uint64
}
type SysvShmDesc struct {
	Perm       SysvIpcPerm
	Atime_high uint32
	Atime      uint32
	Dtime_high uint32
	Dtime      uint32
	Ctime_high uint32
	Ctime      uint32
	_          uint32
	Segsz      uint32
	Cpid       int32
	Lpid       int32
	Nattch     uint32
	_          uint32
	_          uint32
	_          [4]byte
}
