# go-v4l2

A pure Go implementation of Video4Linux2 stream capture with a simple channel
based interface:

* No C code. No separate cross-compiler required.
* Zero copy. Memory-mapped double-buffer scheme makes kernel memory reference
  available via Go channel.

## Quickstart

See `examples/record.go` for a more detailed example.

```
device, err := v4l2.Open("/dev/video0")

device.SetPixelFormat(1280, 720, v4l2.V4L2_PIX_FMT_H264)

device.SetBitrate(1000000)

device.Start()

for {
	select {
	case sample := <-dev.C:
		// Each read is a H.264 NAL unit
		fmt.Printf("Read %d byte sample\n", len(sample.Data))
		sample.Release()
	}
}
```

Build example:

`GOARCH=arm GOOS=linux go build examples/record.go`
