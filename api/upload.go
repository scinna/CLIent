package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gabriel-vasile/mimetype"
	"github.com/scinna/CLIent/serrors"
	"github.com/scinna/CLIent/utils"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

// Upload does what it sound it does, upload the picture to the server
func Upload(serverURL, token string, title, description string, visibility utils.Visibility, collection string, file *os.File) (string, error) {
	_, err := file.Seek(0, io.SeekStart)
	if err != nil {
		return "", errors.New("something went wrong rewinding the file")
	}

	mime, err := mimetype.DetectReader(file)
	if err != nil {
		return "", errors.New("can't determinate the mimetype")
	}

	if !utils.IsMimetypeAllowed(mime.String()) {
		return "", errors.New("unauthorized mimetype")
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	titleField, _ := writer.CreateFormField("title")
	_, _ = titleField.Write([]byte(title))

	desc, _ := writer.CreateFormField("description")
	_, _ = desc.Write([]byte(description))

	visibilityField, _ := writer.CreateFormField("visibility")
	_, _ = visibilityField.Write([]byte(fmt.Sprintf("%v", visibility)))

	collectionField, _ := writer.CreateFormField("collection")
	_, _ = collectionField.Write([]byte(collection))

	pict, _ := writer.CreateFormFile("picture", file.Name())

	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		return "", errors.New("something went wrong rewinding the file")
	}

	_, err = io.Copy(pict, file)
	_ = writer.Close()

	r, err := http.NewRequest("POST", serverURL+"api/upload", body)
	if err != nil {
		return "", errors.New("something went wrong creating the request: "+err.Error())
	}

	r.Header.Set("Content-Type", writer.FormDataContentType())
	r.Header.Set("Authorization", "Bearer "+token)

	resp, err := client.Do(r)
	if err != nil {
		return "", err
	}

	switch resp.StatusCode {
	case http.StatusCreated:

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", errors.New("unknown error")
		}

		var upRes struct {
			MediaID string
		}
		err = json.Unmarshal(body, &upRes)
		if err != nil {
			return "", errors.New("something went wrong creating the request: "+err.Error())
		}

		return serverURL + upRes.MediaID, nil
	case http.StatusUnauthorized:
		return "", serrors.ErrorNotAuthed
	default:
		return "", errors.New(fmt.Sprintf("something went wrong creating the request: %v\n", resp.StatusCode))
	}
}
