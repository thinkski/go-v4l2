//go:build linux && !arm64
// +build linux,!arm64

package v4l2

const (
	VIDIOC_DQBUF       = 0xc0445611
	VIDIOC_QBUF        = 0xc044560f
	VIDIOC_QUERYBUF    = 0xc0445609
	VIDIOC_G_EXT_CTRLS = 0xc0185647
	VIDIOC_S_EXT_CTRLS = 0xc0185648
	VIDIOC_S_FMT       = 0xc0cc5605
)
