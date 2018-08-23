package global

import (
	"apiserver/handler"
	"apiserver/pkg/global/dir"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"os"
	"io"
	"bytes"
	"time"
	"path"
	"mime/multipart"
	"sync"
)

func Uploads(c *gin.Context) {
	if err := c.Request.ParseMultipartForm(2000); err != nil {
		handler.SendResponse(c, err, nil)
		return
	}
	formData := c.Request.MultipartForm
	files := formData.File["file"]
	fileNames := make([]string,0)
	wg := sync.WaitGroup{}
	errChan := make(chan error, 1)
	finished := make(chan bool, 1)
	for _, v := range files {
		wg.Add(1)
		go func(v *multipart.FileHeader) {
			defer wg.Done()
			file, err := v.Open()
			defer file.Close()
			if err != nil {
				errChan <- err
			}
			var currentPath bytes.Buffer

			currentTime := time.Now().Format("2006/01")

			currentPath.WriteString("uploads/")
			currentPath.WriteString(currentTime)

			if _, err := dir.IsDir(currentPath.String(), true); err != nil {
				errChan <- err
			}

			cUuid, err := uuid.NewV1()
			if err != nil {
				errChan <- err
			}
			currentPath.WriteString("/")
			currentPath.WriteString(cUuid.String())
			currentPath.WriteString(path.Ext(v.Filename))
			name := currentPath.String()
			out, err := os.Create(name)
			defer out.Close()

			if err != nil {
				errChan <- err
			}
			if _, err := io.Copy(out, file); err != nil {
				errChan <- err
			}
			fileNames = append(fileNames,name)
		}(v)

	}

	go func() {
		wg.Wait()
		close(finished)
	}()

	select {
	case <-finished:
	case err := <-errChan:
		handler.SendResponse(c, err, nil)
	}

	handler.SendResponse(c, nil, fileNames)
}
