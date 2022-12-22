package images

import (
	"encoding/base64"
	"fmt"
	"io"
	"os"
)

var (
	out        *os.File
	errs, erra error
)

func ServerImage() {
	id := "ub943jbn231ee"
	types := true

	read, err := os.Open("container/image.jpg")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer read.Close()

	data, err := io.ReadAll(read)
	if err != nil {
		fmt.Println(err)
		return
	}
	encodedString := base64.StdEncoding.EncodeToString(data)
	byteslice, err := base64.StdEncoding.DecodeString(encodedString)
	if err != nil {
		fmt.Println(err)
		return
	}

	// ==============        check photo types        ================
	if types {
		out, errs = os.Create("broker/weerker_images.jpg")
	} else {
		str := fmt.Sprintf("brokerprivate/%v_worker_image.jpg", id)
		out, errs = os.Create(str)
	}
	defer out.Close()
	// ==============

	if errs != nil {
		fmt.Println(err)
		return
	}
	_, errs = out.Write(byteslice)
	if errs != nil {
		fmt.Println(err)
		return
	}
	erra = os.Mkdir("broker/admin", 0755)
	if erra != nil {
		fmt.Println(err)
		return
	}
}
