// Copyright 2020 Chris Hiszpanski. All rights reserved.

package main

import (
	"flag"
	"log"
	"os"
	"time"

	"github.com/thinkski/go-v4l2"
)

var flagBitrate int
var flagHeight int
var flagWidth int
var flagOutput string

func init() {
	const (
		defaultBitrate = 3000000
		defaultHeight  = 720
		defaultWidth   = 1280
		defaultOutput  = ""
	)
	flag.IntVar(&flagBitrate, "b", defaultBitrate, "Bitrate")
	flag.IntVar(&flagHeight, "h", defaultHeight, "Height")
	flag.IntVar(&flagWidth, "w", defaultWidth, "Width")
	flag.StringVar(&flagOutput, "o", defaultOutput, "Output file")
}

func main() {
	flag.Parse()

	// Open device
	dev, err := v4l2.Open("/dev/video0")
	if nil != err {
		log.Fatal(err)
	}
	defer dev.Close()

	// Set pixel format
	if err := dev.SetPixelFormat(
		flagWidth,
		flagHeight,
		v4l2.V4L2_PIX_FMT_H264,
	); nil != err {
		log.Fatal(err)
	}

	// Set bitrate
	if err := dev.SetBitrate(int32(flagBitrate)); nil != err {
		log.Fatal(err)
	}

	// Set timer to stop stream after ten seconds
	timer := time.NewTimer(10 * time.Second)

	// Start stream
	if err := dev.Start(); nil != err {
		log.Fatal(err)
	}

	// Open file for writing
	f := os.Stdout
	if flagOutput != "" {
		var err error
		if f, err = os.Create(flagOutput); nil != err {
			log.Fatal(err)
		}
	}
	defer f.Close()

ReadLoop:
	for {
		select {
		case b := <-dev.C:
			f.Write(b.Data)
			b.Release()
		case <-timer.C:
			break ReadLoop
		}
	}

	// Stop stream
	if err := dev.Stop(); nil != err {
		log.Fatal()
	}
}
