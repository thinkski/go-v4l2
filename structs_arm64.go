// +build linux
// +build arm64

package v4l2

import "unsafe"

const (
	maxSizeBufferDotM         = 8
	maxSizeExtControlDotValue = 8
	maxSizeFormatDotFmt       = 200
	sizePixFormat             = 48
)

type v4l2_capability struct {
	driver       [16]uint8
	card         [32]uint8
	bus_info     [32]uint8
	version      uint32
	capabilities uint32
	device_caps  uint32
	reserved     [3]uint32
}

type v4l2_pix_format struct {
	width        uint32
	height       uint32
	pixelformat  uint32
	field        uint32
	bytesperline uint32
	sizeimage    uint32
	colorspace   uint32
	priv         uint32
	flags        uint32
	xx_enc       uint32
	quantization uint32
	xfer_func    uint32
}

type v4l2_format struct {
	typ uint64
	fmt [maxSizeFormatDotFmt]byte // union
}

type v4l2_control struct {
	id    uint32
	value int32
}

type v4l2_requestbuffers struct {
	count        uint32
	typ          uint32
	memory       uint32
	capabilities uint32
	flags        uint8
	reserved     [3]uint8
}

type v4l2_timecode struct {
	typ      uint32
	flags    uint32
	frames   uint8
	seconds  uint8
	minutes  uint8
	hours    uint8
	userbits [4]uint8
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

type v4l2_ext_control struct {
	id        uint32
	size      uint32
	reserved2 [1]uint32
	value     [maxSizeExtControlDotValue]byte // union
}

type v4l2_ext_controls struct {
	ctrl_class uint32
	count      uint32
	error_idx  uint32
	reserved   [2]uint32
	controls   unsafe.Pointer
}

// marshals v4l2_pix_format struct into v4l2_format.fmt union
func (pfmt *v4l2_pix_format) marshal() [maxSizeFormatDotFmt]byte {
	var b [maxSizeFormatDotFmt]byte

	copy(b[0:sizePixFormat], (*[sizePixFormat]byte)(unsafe.Pointer(pfmt))[:])

	return b
}
