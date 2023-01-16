package lib

import ffmpeg "github.com/u2takey/ffmpeg-go"

// Convert converts video
func (c *Client) Convert(from, to string, overwrite bool) error {
	cmd := ffmpeg.Input(from, ffmpeg.KwArgs{}).Output(to)
	if overwrite {
		cmd = cmd.OverWriteOutput()
	}
	if c.Debug {
		cmd = cmd.ErrorToStdOut()
	}
	return cmd.Run()
}
