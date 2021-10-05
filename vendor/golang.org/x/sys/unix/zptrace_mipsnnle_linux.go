// Code generated by linux/mkall.go generatePtracePair("mipsle", "mips64le"). DO NOT EDIT.

//go:build linux && (mipsle || mips64le)
// +build linux
// +build mipsle mips64le

package unix

import "unsafe"

// PtraceRegsMipsle is the registers used by mipsle binaries.
type PtraceRegsMipsle struct ***REMOVED***
	Regs     [32]uint64
	Lo       uint64
	Hi       uint64
	Epc      uint64
	Badvaddr uint64
	Status   uint64
	Cause    uint64
***REMOVED***

// PtraceGetRegsMipsle fetches the registers used by mipsle binaries.
func PtraceGetRegsMipsle(pid int, regsout *PtraceRegsMipsle) error ***REMOVED***
	return ptrace(PTRACE_GETREGS, pid, 0, uintptr(unsafe.Pointer(regsout)))
***REMOVED***

// PtraceSetRegsMipsle sets the registers used by mipsle binaries.
func PtraceSetRegsMipsle(pid int, regs *PtraceRegsMipsle) error ***REMOVED***
	return ptrace(PTRACE_SETREGS, pid, 0, uintptr(unsafe.Pointer(regs)))
***REMOVED***

// PtraceRegsMips64le is the registers used by mips64le binaries.
type PtraceRegsMips64le struct ***REMOVED***
	Regs     [32]uint64
	Lo       uint64
	Hi       uint64
	Epc      uint64
	Badvaddr uint64
	Status   uint64
	Cause    uint64
***REMOVED***

// PtraceGetRegsMips64le fetches the registers used by mips64le binaries.
func PtraceGetRegsMips64le(pid int, regsout *PtraceRegsMips64le) error ***REMOVED***
	return ptrace(PTRACE_GETREGS, pid, 0, uintptr(unsafe.Pointer(regsout)))
***REMOVED***

// PtraceSetRegsMips64le sets the registers used by mips64le binaries.
func PtraceSetRegsMips64le(pid int, regs *PtraceRegsMips64le) error ***REMOVED***
	return ptrace(PTRACE_SETREGS, pid, 0, uintptr(unsafe.Pointer(regs)))
***REMOVED***
