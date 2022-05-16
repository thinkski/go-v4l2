//go:build linux && arm64
// +build linux,arm64

package v4l2

const maxSizeBufferDotM = 8

type v4l2_format struct {
	typ uint64
	fmt [maxSizeFormatDotFmt]byte // union
}

type v4l2_requestbuffers struct {
	count        uint32
	typ          uint32
	memory       uint32
	capabilities uint32
	flags        uint8
	reserved     [3]uint8
}

type timeval struct {
	tv_sec  uint64
	tv_usec uint64
}

type v4l2_buffer struct {
	index     uint32
	typ       uint32
	bytesused uint32
	flags     uint32
	field     uint32
	timestamp timeval
	timecode  v4l2_timecode
	sequence  uint32
	memory    uint32
	m         [maxSizeBufferDotM]byte // union
	length    uint32
	reserved2 uint32
	reserved  uint32
}
