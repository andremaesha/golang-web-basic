package golangweb

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func UploadForm(w http.ResponseWriter, r *http.Request) {
	myTemplates.ExecuteTemplate(w, "upload.form.gohtml", nil)
}

func Upload(w http.ResponseWriter, r *http.Request) {
	// r.ParseMultipartForm(100 << 20) // 100 mb
	file, fileHeader, err := r.FormFile("file")

	if err != nil {
		panic(err)
	}

	fileDestination, err := os.Create("./resources/" + fileHeader.Filename)

	if err != nil {
		panic(err)
	}

	_, err = io.Copy(fileDestination, file)

	if err != nil {
		panic(err)
	}

	name := r.PostFormValue("name")

	myTemplates.ExecuteTemplate(w, "upload.success.gohtml", map[string]interface{}{
		"Name": name,
		"File": "/static/" + fileHeader.Filename,
	})
}

func TestUploadFormServer(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", UploadForm)
	mux.HandleFunc("/upload", Upload)
	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./resources"))))

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}

//go:embed resources/pasPhoto.jpg
var uploadFileTest []byte

func TestUploadFile(t *testing.T) {
	body := new(bytes.Buffer)

	writer := multipart.NewWriter(body)
	writer.WriteField("name", "andre maesha")

	file, _ := writer.CreateFormFile("file", "EXAMPLE.png")

	file.Write(uploadFileTest)
	writer.Close()

	req := httptest.NewRequest(http.MethodPost, "http://localhost:8080/upload", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	recoder := httptest.NewRecorder()

	Upload(recoder, req)

	res := recoder.Result()
	bodyResponse, _ := io.ReadAll(res.Body)
	fmt.Println(string(bodyResponse))
}
