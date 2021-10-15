// This file was autogenerated using go run mkcode.go -- process.go
// DO NOT EDIT.

package ntdll

import "unsafe"

// The Processinfoclass constants have been derived from the PROCESSINFOCLASS enum definition.
type Processinfoclass uint32

const (
	ProcessBasicInformation           Processinfoclass = 0
	ProcessQuotaLimits                                 = 1
	ProcessIoCounters                                  = 2
	ProcessVmCounters                                  = 3
	ProcessTimes                                       = 4
	ProcessBasePriority                                = 5
	ProcessRaisePriority                               = 6
	ProcessDebugPort                                   = 7
	ProcessExceptionPort                               = 8
	ProcessAccessToken                                 = 9
	ProcessLdtInformation                              = 10
	ProcessLdtSize                                     = 11
	ProcessDefaultHardErrorMode                        = 12
	ProcessIoPortHandlers                              = 13
	ProcessPooledUsageAndLimits                        = 14
	ProcessWorkingSetWatch                             = 15
	ProcessUserModeIOPL                                = 16
	ProcessEnableAlignmentFaultFixup                   = 17
	ProcessPriorityClass                               = 18
	ProcessWx86Information                             = 19
	ProcessHandleCount                                 = 20
	ProcessAffinityMask                                = 21
	ProcessPriorityBoost                               = 22
	ProcessDeviceMap                                   = 23
	ProcessSessionInformation                          = 24
	ProcessForegroundInformation                       = 25
	ProcessWow64Information                            = 26
	ProcessImageFileName                               = 27
	ProcessLUIDDeviceMapsEnabled                       = 28
	ProcessBreakOnTermination                          = 29
	ProcessDebugObjectHandle                           = 30
	ProcessDebugFlags                                  = 31
	ProcessHandleTracing                               = 32
	ProcessIoPriority                                  = 33
	ProcessExecuteFlags                                = 34
	ProcessTlsInformation                              = 35
	ProcessCookie                                      = 36
	ProcessImageInformation                            = 37
	ProcessCycleTime                                   = 38
	ProcessPagePriority                                = 39
	ProcessInstrumentationCallback                     = 40
	ProcessThreadStackAllocation                       = 41
	ProcessWorkingSetWatchEx                           = 42
	ProcessImageFileNameWin32                          = 43
	ProcessImageFileMapping                            = 44
	ProcessAffinityUpdateMode                          = 45
	ProcessMemoryAllocationMode                        = 46
	ProcessGroupInformation                            = 47
	ProcessTokenVirtualizationEnabled                  = 48
	ProcessConsoleHostProcess                          = 49
	ProcessWindowInformation                           = 50
)

// The Threadinfoclass constants have been derived from the THREADINFOCLASS enum definition.
type Threadinfoclass uint32

const (
	ThreadBasicInformation          Threadinfoclass = 0
	ThreadTimes                                     = 1
	ThreadPriority                                  = 2
	ThreadBasePriority                              = 3
	ThreadAffinityMask                              = 4
	ThreadImpersonationToken                        = 5
	ThreadDescriptorTableEntry                      = 6
	ThreadEnableAlignmentFaultFixup                 = 7
	ThreadEventPair_Reusable                        = 8
	ThreadQuerySetWin32StartAddress                 = 9
	ThreadZeroTlsCell                               = 10
	ThreadPerformanceCount                          = 11
	ThreadAmILastThread                             = 12
	ThreadIdealProcessor                            = 13
	ThreadPriorityBoost                             = 14
	ThreadSetTlsArrayAddress                        = 15
	ThreadIsIoPending                               = 16
	ThreadHideFromDebugger                          = 17
	ThreadBreakOnTermination                        = 18
	ThreadSwitchLegacyState                         = 19
	ThreadIsTerminated                              = 20
	ThreadLastSystemCall                            = 21
	ThreadIoPriority                                = 22
	ThreadCycleTime                                 = 23
	ThreadPagePriority                              = 24
	ThreadActualBasePriority                        = 25
	ThreadTebInformation                            = 26
	ThreadCSwitchMon                                = 27
	ThreadCSwitchPmu                                = 28
	ThreadWow64Context                              = 29
	ThreadGroupInformation                          = 30
	ThreadUmsInformation                            = 31
	ThreadCounterProfiling                          = 32
	ThreadIdealProcessorEx                          = 33
	MaxThreadInfoClass                              = 34
)

var (
	procNtQueryInformationProcess = modntdll.NewProc("NtQueryInformationProcess")
	procNtOpenProcess             = modntdll.NewProc("NtOpenProcess")
	procNtOpenThread              = modntdll.NewProc("NtOpenThread")
	procNtQueryInformationThread  = modntdll.NewProc("NtQueryInformationThread")
	procNtSetInformationThread    = modntdll.NewProc("NtSetInformationThread")
	procNtSetInformationProcess   = modntdll.NewProc("NtSetInformationProcess")
)

// unknown-parameter: ProcessHandle, ProcessInformationClass, ProcessInformation, ProcessInformationLength, ReturnLength.
func NtQueryInformationProcess(
	ProcessHandle Handle,
	ProcessInformationClass Processinfoclass,
	ProcessInformation *byte,
	ProcessInformationLength uint32,
	ReturnLength *uint32,
) NtStatus {
	r0, _, _ := procNtQueryInformationProcess.Call(uintptr(ProcessHandle),
		uintptr(ProcessInformationClass),
		uintptr(unsafe.Pointer(ProcessInformation)),
		uintptr(ProcessInformationLength),
		uintptr(unsafe.Pointer(ReturnLength)))
	return NtStatus(r0)
}

// OUT-parameter: ProcessHandle.
// *OPT-parameter: ClientId.
func NtOpenProcess(
	ProcessHandle *Handle,
	DesiredAccess AccessMask,
	ObjectAttributes *ObjectAttributes,
	ClientId *ClientId,
) NtStatus {
	r0, _, _ := procNtOpenProcess.Call(uintptr(unsafe.Pointer(ProcessHandle)),
		uintptr(DesiredAccess),
		uintptr(unsafe.Pointer(ObjectAttributes)),
		uintptr(unsafe.Pointer(ClientId)))
	return NtStatus(r0)
}

// OUT-parameter: ThreadHandle.
// *OPT-parameter: ClientId.
func NtOpenThread(
	ThreadHandle *Handle,
	DesiredAccess AccessMask,
	ObjectAttributes *ObjectAttributes,
	ClientId *ClientId,
) NtStatus {
	r0, _, _ := procNtOpenThread.Call(uintptr(unsafe.Pointer(ThreadHandle)),
		uintptr(DesiredAccess),
		uintptr(unsafe.Pointer(ObjectAttributes)),
		uintptr(unsafe.Pointer(ClientId)))
	return NtStatus(r0)
}

// unknown-parameter: ThreadHandle, ThreadInformationClass, ThreadInformation, ThreadInformationLength, ReturnLength.
func NtQueryInformationThread(
	ThreadHandle Handle,
	ThreadInformationClass Threadinfoclass,
	ThreadInformation *byte,
	ThreadInformationLength uint32,
	ReturnLength *uint32,
) NtStatus {
	r0, _, _ := procNtQueryInformationThread.Call(uintptr(ThreadHandle),
		uintptr(ThreadInformationClass),
		uintptr(unsafe.Pointer(ThreadInformation)),
		uintptr(ThreadInformationLength),
		uintptr(unsafe.Pointer(ReturnLength)))
	return NtStatus(r0)
}

// unknown-parameter: ThreadHandle, ThreadInformationClass, ThreadInformation, ThreadInformationLength.
func NtSetInformationThread(
	ThreadHandle Handle,
	ThreadInformationClass Threadinfoclass,
	ThreadInformation *byte,
	ThreadInformationLength uint32,
) NtStatus {
	r0, _, _ := procNtSetInformationThread.Call(uintptr(ThreadHandle),
		uintptr(ThreadInformationClass),
		uintptr(unsafe.Pointer(ThreadInformation)),
		uintptr(ThreadInformationLength))
	return NtStatus(r0)
}

// unknown-parameter: ProcessHandle, ProcessInformationClass, ProcessInformation, ProcessInformationLength.
func NtSetInformationProcess(
	ProcessHandle Handle,
	ProcessInformationClass Processinfoclass,
	ProcessInformation *byte,
	ProcessInformationLength uint32,
) NtStatus {
	r0, _, _ := procNtSetInformationProcess.Call(uintptr(ProcessHandle),
		uintptr(ProcessInformationClass),
		uintptr(unsafe.Pointer(ProcessInformation)),
		uintptr(ProcessInformationLength))
	return NtStatus(r0)
}
