package handle

import (
	"fmt"
	"io"
	"net/http"
)

func Parse(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(1 << 10) // 即 1 * 2^10

	body := make([]byte, r.ContentLength)
	r.Body.Read(body)
	fmt.Println("Body:" + string(body))

	fmt.Fprintln(w, r.Form)
	if r.MultipartForm != nil {
		fmt.Fprintln(w, r.MultipartForm)
	}
	// 自己写的文件读取
	// if r.MultipartForm.File != nil {
	// 	fileHeader := r.MultipartForm.File["uploadFile"][0]
	// 	file, err := fileHeader.Open()
	// 	if err != nil {
	// 		fmt.Fprintln(w, err)
	// 	}
	// 	defer file.Close()
	// 	content := make([]byte, fileHeader.Size)
	// 	file.Read(content)
	// 	fmt.Fprintln(w, string(content))
	// }
	file, header, err := r.FormFile("uploadFile")
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
	defer file.Close()
	fmt.Fprintln(w, header.Header)
	fmt.Fprintln(w, header.Filename)
	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
	fmt.Fprintln(w, string(data))
}
