// This file was autogenerated using go run mkcode.go -- file.go
// DO NOT EDIT.

package ntdll

import "unsafe"

// The FileInformationClass constants have been derived from the FILE_INFORMATION_CLASS enum definition.
type FileInformationClass uint32

const (
	FileDirectoryInformation                FileInformationClass = 1
	FileFullDirectoryInformation                                 = 2
	FileBothDirectoryInformation                                 = 3
	FileBasicInformation                                         = 4
	FileStandardInformation                                      = 5
	FileInternalInformation                                      = 6
	FileEaInformation                                            = 7
	FileAccessInformation                                        = 8
	FileNameInformation                                          = 9
	FileRenameInformation                                        = 10
	FileLinkInformation                                          = 11
	FileNamesInformation                                         = 12
	FileDispositionInformation                                   = 13
	FilePositionInformation                                      = 14
	FileFullEaInformation                                        = 15
	FileModeInformation                                          = 16
	FileAlignmentInformation                                     = 17
	FileAllInformation                                           = 18
	FileAllocationInformation                                    = 19
	FileEndOfFileInformation                                     = 20
	FileAlternateNameInformation                                 = 21
	FileStreamInformation                                        = 22
	FilePipeInformation                                          = 23
	FilePipeLocalInformation                                     = 24
	FilePipeRemoteInformation                                    = 25
	FileMailslotQueryInformation                                 = 26
	FileMailslotSetInformation                                   = 27
	FileCompressionInformation                                   = 28
	FileObjectIdInformation                                      = 29
	FileCompletionInformation                                    = 30
	FileMoveClusterInformation                                   = 31
	FileQuotaInformation                                         = 32
	FileReparsePointInformation                                  = 33
	FileNetworkOpenInformation                                   = 34
	FileAttributeTagInformation                                  = 35
	FileTrackingInformation                                      = 36
	FileIdBothDirectoryInformation                               = 37
	FileIdFullDirectoryInformation                               = 38
	FileValidDataLengthInformation                               = 39
	FileShortNameInformation                                     = 40
	FileIoCompletionNotificationInformation                      = 41
	FileIoStatusBlockRangeInformation                            = 42
	FileIoPriorityHintInformation                                = 43
	FileSfioReserveInformation                                   = 44
	FileSfioVolumeInformation                                    = 45
	FileHardLinkInformation                                      = 46
	FileProcessIdsUsingFileInformation                           = 47
	FileNormalizedNameInformation                                = 48
	FileNetworkPhysicalNameInformation                           = 49
	FileIdGlobalTxDirectoryInformation                           = 50
	FileIsRemoteDeviceInformation                                = 51
	FileUnusedInformation                                        = 52
	FileNumaNodeInformation                                      = 53
	FileStandardLinkInformation                                  = 54
	FileRemoteProtocolInformation                                = 55
	FileRenameInformationBypassAccessCheck                       = 56
	FileLinkInformationBypassAccessCheck                         = 57
	FileVolumeNameInformation                                    = 58
	FileIdInformation                                            = 59
	FileIdExtdDirectoryInformation                               = 60
	FileReplaceCompletionInformation                             = 61
	FileHardLinkFullIdInformation                                = 62
	FileIdExtdBothDirectoryInformation                           = 63
	FileMaximumInformation                                       = 64
)

var (
	procNtCreateFile           = modntdll.NewProc("NtCreateFile")
	procNtDeviceIoControlFile  = modntdll.NewProc("NtDeviceIoControlFile")
	procNtOpenFile             = modntdll.NewProc("NtOpenFile")
	procNtQueryInformationFile = modntdll.NewProc("NtQueryInformationFile")
	procNtReadFile             = modntdll.NewProc("NtReadFile")
	procNtSetInformationFile2  = modntdll.NewProc("NtSetInformationFile2")
	procNtWriteFile            = modntdll.NewProc("NtWriteFile")
)

// FileDirectoryInformationT has been derived from the FILE_DIRECTORY_INFORMATION struct definition.
type FileDirectoryInformationT struct {
	NextEntryOffset uint32
	FileIndex       uint32
	CreationTime    int64
	LastAccessTime  int64
	LastWriteTime   int64
	ChangeTime      int64
	EndOfFile       int64
	AllocationSize  int64
	FileAttributes  uint32
	FileNameLength  uint32
	FileName        [1]uint16
}

// FileFullDirInformationT has been derived from the FILE_FULL_DIR_INFORMATION struct definition.
type FileFullDirInformationT struct {
	NextEntryOffset uint32
	FileIndex       uint32
	CreationTime    int64
	LastAccessTime  int64
	LastWriteTime   int64
	ChangeTime      int64
	EndOfFile       int64
	AllocationSize  int64
	FileAttributes  uint32
	FileNameLength  uint32
	EaSize          uint32
	FileName        [1]uint16
}

// FileBothDirInformationT has been derived from the FILE_BOTH_DIR_INFORMATION struct definition.
type FileBothDirInformationT struct {
	NextEntryOffset uint32
	FileIndex       uint32
	CreationTime    int64
	LastAccessTime  int64
	LastWriteTime   int64
	ChangeTime      int64
	EndOfFile       int64
	AllocationSize  int64
	FileAttributes  uint32
	FileNameLength  uint32
	EaSize          uint32
	ShortNameLength byte
	ShortName       [12]uint16
	FileName        [1]uint16
}

// FileBasicInformationT has been derived from the FILE_BASIC_INFORMATION struct definition.
type FileBasicInformationT struct {
	CreationTime   int64
	LastAccessTime int64
	LastWriteTime  int64
	ChangeTime     int64
	FileAttributes uint32
}

// FileStandardInformationT has been derived from the FILE_STANDARD_INFORMATION struct definition.
type FileStandardInformationT struct {
	AllocationSize int64
	EndOfFile      int64
	NumberOfLinks  uint32
	DeletePending  bool
	Directory      bool
}

// FileInternalInformationT has been derived from the FILE_INTERNAL_INFORMATION struct definition.
type FileInternalInformationT struct {
	IndexNumber int64
}

// FileEaInformationT has been derived from the FILE_EA_INFORMATION struct definition.
type FileEaInformationT struct {
	EaSize uint32
}

// FileAccessInformationT has been derived from the FILE_ACCESS_INFORMATION struct definition.
type FileAccessInformationT struct {
	AccessFlags AccessMask
}

// FileNameInformationT has been derived from the FILE_NAME_INFORMATION struct definition.
type FileNameInformationT struct {
	FileNameLength uint32
	FileName       [1]uint16
}

// FileRenameInformationT has been derived from the FILE_RENAME_INFORMATION struct definition.
type FileRenameInformationT struct {
	ReplaceIfExists bool
	RootDirectory   Handle
	FileNameLength  uint32
	FileName        [1]uint16
}

// FileLinkInformationT has been derived from the FILE_LINK_INFORMATION struct definition.
type FileLinkInformationT struct {
	ReplaceIfExists bool
	RootDirectory   Handle
	FileNameLength  uint32
	FileName        [1]uint16
}

// FileNamesInformationT has been derived from the FILE_NAMES_INFORMATION struct definition.
type FileNamesInformationT struct {
	NextEntryOffset uint32
	FileIndex       uint32
	FileNameLength  uint32
	FileName        [1]uint16
}

// FileDispositionInformationT has been derived from the FILE_DISPOSITION_INFORMATION struct definition.
type FileDispositionInformationT struct {
	DeleteFile bool
}

// FilePositionInformationT has been derived from the FILE_POSITION_INFORMATION struct definition.
type FilePositionInformationT struct {
	CurrentByteOffset int64
}

// FileFullEaInformationT has been derived from the FILE_FULL_EA_INFORMATION struct definition.
type FileFullEaInformationT struct {
	NextEntryOffset uint32
	Flags           byte
	EaNameLength    byte
	EaValueLength   uint16
	EaName          [1]byte
}

// FileModeInformationT has been derived from the FILE_MODE_INFORMATION struct definition.
type FileModeInformationT struct {
	Mode uint32
}

// FileAlignmentInformationT has been derived from the FILE_ALIGNMENT_INFORMATION struct definition.
type FileAlignmentInformationT struct {
	AlignmentRequirement uint32
}

// FileAllInformationT has been derived from the FILE_ALL_INFORMATION struct definition.
type FileAllInformationT struct {
	BasicInformation     FileBasicInformationT
	StandardInformation  FileStandardInformationT
	InternalInformation  FileInternalInformationT
	EaInformation        FileEaInformationT
	AccessInformation    FileAccessInformationT
	PositionInformation  FilePositionInformationT
	ModeInformation      FileModeInformationT
	AlignmentInformation FileAlignmentInformationT
	NameInformation      FileNameInformationT
}

// FileAllocationInformationT has been derived from the FILE_ALLOCATION_INFORMATION struct definition.
type FileAllocationInformationT struct {
	AllocationSize int64
}

// FileEndOfFileInformationT has been derived from the FILE_END_OF_FILE_INFORMATION struct definition.
type FileEndOfFileInformationT struct {
	EndOfFile int64
}

// FileStreamInformationT has been derived from the FILE_STREAM_INFORMATION struct definition.
type FileStreamInformationT struct {
	NextEntryOffset      uint32
	StreamNameLength     uint32
	StreamSize           int64
	StreamAllocationSize int64
	StreamName           [1]uint16
}

// FilePipeInformationT has been derived from the FILE_PIPE_INFORMATION struct definition.
type FilePipeInformationT struct {
	ReadMode       uint32
	CompletionMode uint32
}

// FilePipeLocalInformationT has been derived from the FILE_PIPE_LOCAL_INFORMATION struct definition.
type FilePipeLocalInformationT struct {
	NamedPipeType          uint32
	NamedPipeConfiguration uint32
	MaximumInstances       uint32
	CurrentInstances       uint32
	InboundQuota           uint32
	ReadDataAvailable      uint32
	OutboundQuota          uint32
	WriteQuotaAvailable    uint32
	NamedPipeState         uint32
	NamedPipeEnd           uint32
}

// FilePipeRemoteInformationT has been derived from the FILE_PIPE_REMOTE_INFORMATION struct definition.
type FilePipeRemoteInformationT struct {
	CollectDataTime        int64
	MaximumCollectionCount uint32
}

// FileMailslotQueryInformationT has been derived from the FILE_MAILSLOT_QUERY_INFORMATION struct definition.
type FileMailslotQueryInformationT struct {
	MaximumMessageSize uint32
	MailslotQuota      uint32
	NextMessageSize    uint32
	MessagesAvailable  uint32
	ReadTimeout        int64
}

// FileMailslotSetInformationT has been derived from the FILE_MAILSLOT_SET_INFORMATION struct definition.
type FileMailslotSetInformationT struct {
	ReadTimeout *int64
}

// FileCompressionInformationT has been derived from the FILE_COMPRESSION_INFORMATION struct definition.
type FileCompressionInformationT struct {
	CompressedFileSize   int64
	CompressionFormat    uint16
	CompressionUnitShift byte
	ChunkShift           byte
	ClusterShift         byte
	Reserved             [3]byte
}

// FileQuotaInformationT has been derived from the FILE_QUOTA_INFORMATION struct definition.
type FileQuotaInformationT struct {
	NextEntryOffset uint32
	SidLength       uint32
	ChangeTime      int64
	QuotaUsed       int64
	QuotaThreshold  int64
	QuotaLimit      int64
	Sid             Sid
}

// FileReparsePointInformationT has been derived from the FILE_REPARSE_POINT_INFORMATION struct definition.
type FileReparsePointInformationT struct {
	FileReference int64
	Tag           uint32
}

// FileNetworkOpenInformationT has been derived from the FILE_NETWORK_OPEN_INFORMATION struct definition.
type FileNetworkOpenInformationT struct {
	CreationTime   int64
	LastAccessTime int64
	LastWriteTime  int64
	ChangeTime     int64
	AllocationSize int64
	EndOfFile      int64
	FileAttributes uint32
}

// FileAttributeTagInformationT has been derived from the FILE_ATTRIBUTE_TAG_INFORMATION struct definition.
type FileAttributeTagInformationT struct {
	FileAttributes uint32
	ReparseTag     uint32
}

// FileIdBothDirInformationT has been derived from the FILE_ID_BOTH_DIR_INFORMATION struct definition.
type FileIdBothDirInformationT struct {
	NextEntryOffset uint32
	FileIndex       uint32
	CreationTime    int64
	LastAccessTime  int64
	LastWriteTime   int64
	ChangeTime      int64
	EndOfFile       int64
	AllocationSize  int64
	FileAttributes  uint32
	FileNameLength  uint32
	EaSize          uint32
	ShortNameLength byte
	ShortName       [12]uint16
	FileId          int64
	FileName        [1]uint16
}

// FileIdFullDirInformationT has been derived from the FILE_ID_FULL_DIR_INFORMATION struct definition.
type FileIdFullDirInformationT struct {
	NextEntryOffset uint32
	FileIndex       uint32
	CreationTime    int64
	LastAccessTime  int64
	LastWriteTime   int64
	ChangeTime      int64
	EndOfFile       int64
	AllocationSize  int64
	FileAttributes  uint32
	FileNameLength  uint32
	EaSize          uint32
	FileId          int64
	FileName        [1]uint16
}

// FileValidDataLengthInformationT has been derived from the FILE_VALID_DATA_LENGTH_INFORMATION struct definition.
type FileValidDataLengthInformationT struct {
	ValidDataLength int64
}

// FileIoPriorityHintInformationT has been derived from the FILE_IO_PRIORITY_HINT_INFORMATION struct definition.
type FileIoPriorityHintInformationT struct {
	PriorityHint IoPriorityHint
}

// FileLinkEntryInformationT has been derived from the FILE_LINK_ENTRY_INFORMATION struct definition.
type FileLinkEntryInformationT struct {
	NextEntryOffset uint32
	ParentFileId    int64
	FileNameLength  uint32
	FileName        uint16
}

// FileLinksInformationT has been derived from the FILE_LINKS_INFORMATION struct definition.
type FileLinksInformationT struct {
	BytesNeeded     uint32
	EntriesReturned uint32
	Entry           FileLinkEntryInformationT
}

// FileNetworkPhysicalNameInformationT has been derived from the FILE_NETWORK_PHYSICAL_NAME_INFORMATION struct definition.
type FileNetworkPhysicalNameInformationT struct {
	FileNameLength uint32
	FileName       [1]uint16
}

// FileIdGlobalTxDirInformationT has been derived from the FILE_ID_GLOBAL_TX_DIR_INFORMATION struct definition.
type FileIdGlobalTxDirInformationT struct {
	NextEntryOffset      uint32
	FileIndex            uint32
	CreationTime         int64
	LastAccessTime       int64
	LastWriteTime        int64
	ChangeTime           int64
	EndOfFile            int64
	AllocationSize       int64
	FileAttributes       uint32
	FileNameLength       uint32
	FileId               int64
	LockingTransactionId Guid
	TxInfoFlags          uint32
	FileName             [1]uint16
}

// FileCompletionInformationT has been derived from the FILE_COMPLETION_INFORMATION struct definition.
type FileCompletionInformationT struct {
	Port Handle
	Key  *byte
}

// OUT-parameter: FileHandle, IoStatusBlock.
// *OPT-parameter: AllocationSize.
func NtCreateFile(
	FileHandle *Handle,
	DesiredAccess AccessMask,
	ObjectAttributes *ObjectAttributes,
	IoStatusBlock *IoStatusBlock,
	AllocationSize *int64,
	FileAttributes uint32,
	ShareAccess uint32,
	CreateDisposition uint32,
	CreateOptions uint32,
	EaBuffer *byte,
	EaLength uint32,
) NtStatus {
	r0, _, _ := procNtCreateFile.Call(uintptr(unsafe.Pointer(FileHandle)),
		uintptr(DesiredAccess),
		uintptr(unsafe.Pointer(ObjectAttributes)),
		uintptr(unsafe.Pointer(IoStatusBlock)),
		uintptr(unsafe.Pointer(AllocationSize)),
		uintptr(FileAttributes),
		uintptr(ShareAccess),
		uintptr(CreateDisposition),
		uintptr(CreateOptions),
		uintptr(unsafe.Pointer(EaBuffer)),
		uintptr(EaLength))
	return NtStatus(r0)
}

// OUT-parameter: IoStatusBlock, OutputBuffer.
func NtDeviceIoControlFile(
	FileHandle Handle,
	Event Handle,
	ApcRoutine *IoApcRoutine,
	ApcContext *byte,
	IoStatusBlock *IoStatusBlock,
	IoControlCode uint32,
	InputBuffer *byte,
	InputBufferLength uint32,
	OutputBuffer *byte,
	OutputBufferLength uint32,
) NtStatus {
	r0, _, _ := procNtDeviceIoControlFile.Call(uintptr(FileHandle),
		uintptr(Event),
		uintptr(unsafe.Pointer(ApcRoutine)),
		uintptr(unsafe.Pointer(ApcContext)),
		uintptr(unsafe.Pointer(IoStatusBlock)),
		uintptr(IoControlCode),
		uintptr(unsafe.Pointer(InputBuffer)),
		uintptr(InputBufferLength),
		uintptr(unsafe.Pointer(OutputBuffer)),
		uintptr(OutputBufferLength))
	return NtStatus(r0)
}

// OUT-parameter: FileHandle, IoStatusBlock.
func NtOpenFile(
	FileHandle *Handle,
	DesiredAccess AccessMask,
	ObjectAttributes *ObjectAttributes,
	IoStatusBlock *IoStatusBlock,
	ShareAccess uint32,
	OpenOptions uint32,
) NtStatus {
	r0, _, _ := procNtOpenFile.Call(uintptr(unsafe.Pointer(FileHandle)),
		uintptr(DesiredAccess),
		uintptr(unsafe.Pointer(ObjectAttributes)),
		uintptr(unsafe.Pointer(IoStatusBlock)),
		uintptr(ShareAccess),
		uintptr(OpenOptions))
	return NtStatus(r0)
}

// OUT-parameter: IoStatusBlock, FileInformation.
func NtQueryInformationFile(
	FileHandle Handle,
	IoStatusBlock *IoStatusBlock,
	FileInformation *byte,
	Length uint32,
	FileInformationClass FileInformationClass,
) NtStatus {
	r0, _, _ := procNtQueryInformationFile.Call(uintptr(FileHandle),
		uintptr(unsafe.Pointer(IoStatusBlock)),
		uintptr(unsafe.Pointer(FileInformation)),
		uintptr(Length),
		uintptr(FileInformationClass))
	return NtStatus(r0)
}

// OUT-parameter: IoStatusBlock, Buffer.
// *OPT-parameter: Event, ApcRoutine, ApcContext, ByteOffset, Key.
func NtReadFile(
	FileHandle Handle,
	Event Handle,
	ApcRoutine *IoApcRoutine,
	ApcContext *byte,
	IoStatusBlock *IoStatusBlock,
	Buffer *byte,
	Length uint32,
	ByteOffset *int64,
	Key *uint32,
) NtStatus {
	r0, _, _ := procNtReadFile.Call(uintptr(FileHandle),
		uintptr(Event),
		uintptr(unsafe.Pointer(ApcRoutine)),
		uintptr(unsafe.Pointer(ApcContext)),
		uintptr(unsafe.Pointer(IoStatusBlock)),
		uintptr(unsafe.Pointer(Buffer)),
		uintptr(Length),
		uintptr(unsafe.Pointer(ByteOffset)),
		uintptr(unsafe.Pointer(Key)))
	return NtStatus(r0)
}

// OUT-parameter: IoStatusBlock.
func NtSetInformationFile2(
	FileHandle Handle,
	IoStatusBlock *IoStatusBlock,
	FileInformation *byte,
	Length uint32,
	FileInformationClass FileInformationClass,
) NtStatus {
	r0, _, _ := procNtSetInformationFile2.Call(uintptr(FileHandle),
		uintptr(unsafe.Pointer(IoStatusBlock)),
		uintptr(unsafe.Pointer(FileInformation)),
		uintptr(Length),
		uintptr(FileInformationClass))
	return NtStatus(r0)
}

// OUT-parameter: IoStatusBlock.
// *OPT-parameter: Event, ApcRoutine, ApcContext, ByteOffset, Key.
func NtWriteFile(
	FileHandle Handle,
	Event Handle,
	ApcRoutine *IoApcRoutine,
	ApcContext *byte,
	IoStatusBlock *IoStatusBlock,
	Buffer *byte,
	Length uint32,
	ByteOffset *int64,
	Key *uint32,
) NtStatus {
	r0, _, _ := procNtWriteFile.Call(uintptr(FileHandle),
		uintptr(Event),
		uintptr(unsafe.Pointer(ApcRoutine)),
		uintptr(unsafe.Pointer(ApcContext)),
		uintptr(unsafe.Pointer(IoStatusBlock)),
		uintptr(unsafe.Pointer(Buffer)),
		uintptr(Length),
		uintptr(unsafe.Pointer(ByteOffset)),
		uintptr(unsafe.Pointer(Key)))
	return NtStatus(r0)
}
