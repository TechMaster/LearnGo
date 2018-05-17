package encode_server

import (
	"fmt"
	"os/exec"
	"strconv"
	"bytes"
	"os"
	"bufio"
	"strings"
	"path/filepath"
	"path"
	si "github.com/TechMaster/LearnGo/UploadConvertFeedback/socket_interface"
	"math/rand"
	"time"
)

type VideoEncoder struct {
	outputFolder string
	//socketServer *socket_server.SocketServer //this code cause circular dependency
	socketServer si.ISocketNotify
}

func NewVideoEncoder(outputFolder string, ss si.ISocketNotify) VideoEncoder {
	return VideoEncoder{
		outputFolder: outputFolder,
		socketServer: ss,
	}
}

func (vd VideoEncoder) encodeVideoToHLS(i int, job job) {
	fmt.Printf("Start encoding video %s\n", job.fileName)
	if totalFrame, err := vd.getTotalFrame(job.fileName); err == nil {

		//Send totalFrame to WebSocket Server
		vd.socketServer.NotifyTotalFrame(si.TotalFrameMessage{
			WSConnID:job.wsConID,
			TotalFrame:totalFrame,
		})

		vd.encodeVideo(job, totalFrame)
	}
}

func (vd VideoEncoder) mockEncodeVideoHLS(i int, job job) {
	rand.Seed(time.Now().Unix())
	totalFrame := rand.Intn(100) + 30
	//Send totalFrame to WebSocket Server
	vd.socketServer.NotifyTotalFrame(si.TotalFrameMessage{
		WSConnID:job.wsConID,
		TotalFrame:totalFrame,
	})
	vd.mockEncode(job, totalFrame)
}

func (vd VideoEncoder) mockEncode(job job, totalFrame int) {
	for i:=0; i <= totalFrame; i++ {
		duration, _ := time.ParseDuration("500ms")
		time.Sleep(duration)
		vd.socketServer.NotifyCurrentFrame(si.CurrentFrameMessage{
			WSConnID: job.wsConID,
			CurrentFrame: i,
		})
	}

}
func (vd VideoEncoder) getTotalFrame(videoFile string) (totalFrame int, error error) {
	ffprobeArgs := []string {
		"-select_streams", "v:0", "-show_entries", "stream=nb_frames", "-of",
		"csv=s=x:p=0", "-v", "quiet", videoFile,
	}

	ffprobeCmd := exec.Command("ffprobe", ffprobeArgs...)
	var stdout bytes.Buffer
	ffprobeCmd.Stdout = &stdout
	err := ffprobeCmd.Run()
	if err != nil {
		fmt.Println("Error getTotalFrame")
		return 0, err
	}
	outStr := string(stdout.Bytes())[0 : len(string(stdout.Bytes())) - 1]
	outInt, err := strconv.Atoi(outStr)
	return outInt, err
}

func (vd VideoEncoder) encodeVideo(job job, totalFrame int) {
	fileName := filepath.Base(job.fileName)
	mp3u8 := vd.outputFolder + "/" + strings.TrimSuffix(fileName, path.Ext(fileName)) + ".m3u8"

	fmt.Println(mp3u8)
	cmdArgs := []string{
		"-i", job.fileName,
		"-vf", "scale=w=1920:h=1080:force_original_aspect_ratio=decrease", "-c:v", "h264", "-c:a", "aac", "-ar", "48000", "-profile:v", "main", "-crf", "20", "-sc_threshold", "0", "-g", "48", "-keyint_min", "48", "-maxrate", "5350k", "-bufsize", "7500k", "-b:a", "192k", "-hls_playlist_type", "vod", "-s",
		"1920x1080", mp3u8,
		"-hide_banner", "-loglevel", "level+quiet", "-progress", "/dev/stdout",
	}
	cmd := exec.Command("ffmpeg",cmdArgs...)
	cmdReader, err := cmd.StdoutPipe()

	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating StdoutPipe for Cmd", err)
	}

	scanner := bufio.NewScanner(cmdReader)
	go func() {
		for scanner.Scan(){
			tmp := scanner.Text()
			if strings.HasPrefix(tmp, "frame=") {
				currentFrame, _ := strconv.Atoi(tmp[6:])

				//Notify currentFrame to WebSocket Server
				vd.socketServer.NotifyCurrentFrame(si.CurrentFrameMessage{
					WSConnID: job.wsConID,
					CurrentFrame: currentFrame,
				})
			}

		}
	}()

	err = cmd.Start()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error starting Cmd", err)
	}

	err = cmd.Wait()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error waiting for Cmd", err)
	}
}
