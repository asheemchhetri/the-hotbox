package main

import (
	"fmt"
	"os"
	"gocv.io/x/gocv"
	"./video"
)


func main() {
	if len(os.Args) != 2 {
		fmt.Println("arg err") // check number of cli arguments
		return
	}

	file := os.Args[1] // get file name
	videoIn:= video.OpenVideo(file) // open file as video
	defer videoIn.Close()


	outfilename := video.NameOut(file)
	videoOut, _ := gocv.VideoWriterFile(outfilename,
								videoIn.CodecString(),
								videoIn.Get(gocv.VideoCaptureFPS),
								int(videoIn.Get(gocv.VideoCaptureFrameWidth)),
								int(videoIn.Get(gocv.VideoCaptureFrameHeight)),
								true)
	defer videoOut.Close()

	curr := gocv.NewMat() // reader mat
	defer curr.Close()
	
	//video.ModifyContrast(videoIn,videoOut,.8)
	video.ModifyBrightnessSync(videoIn,videoOut,50)


	// NOTE: the output doesn't have sound
}
