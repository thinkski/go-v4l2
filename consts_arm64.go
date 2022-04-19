// go:build linux && arm64

package v4l2

const (
	VIDIOC_DQBUF       = 0xc0585611
	VIDIOC_QBUF        = 0xc058560f
	VIDIOC_QUERYBUF    = 0xc0585609
	VIDIOC_G_EXT_CTRLS = 0xc0205647
	VIDIOC_S_EXT_CTRLS = 0xc0205648
	VIDIOC_S_FMT       = 0xc0d05605
)
