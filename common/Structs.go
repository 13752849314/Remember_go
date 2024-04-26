package common

import (
	"errors"
	"os"
	"strconv"
	"strings"
	"time"
)

type ChangeUserP struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

type ChangeUserI struct {
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Birthday MyDate `json:"birthday"`
}

type ChangeBillI struct {
	ConsumeType  int        `json:"consumeType"`
	ConsumeMoney float64    `json:"consumeMoney"`
	ConsumeTime  MyDateTime `json:"consumeTime"`
	Remark       string     `json:"remark"`
}

type FileInfo struct {
	Filename   string     `json:"filename"`
	UploadTime MyDateTime `json:"uploadTime"`
	Username   string     `json:"username"`
	Size       string     `json:"size"`
	Timestamp  int64      `json:"timestamp"`
}

func NewFileInfo(info os.FileInfo, username string) (*FileInfo, error) {
	s := new(FileInfo)
	filenames := info.Name()
	var (
		filename string
		times    int64
	)
	ss := strings.Split(filenames, "-")
	if len(ss) != 2 {
		return nil, errors.New("")
	}
	filename = ss[1]
	times, _ = strconv.ParseInt(ss[0], 10, 64)
	s.Timestamp = times
	s.Filename = filename
	s.Username = username
	s.Size = strconv.FormatInt(info.Size()/1024+1, 10) + "KB"
	s.UploadTime = MyDateTime(time.Unix(times, 0))
	return s, nil
}
