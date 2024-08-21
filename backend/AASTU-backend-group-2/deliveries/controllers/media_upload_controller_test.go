package controllers

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestFileUpload(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("success", func(t *testing.T) {
		// Setup
		recorder := httptest.NewRecorder()
		context, _ := gin.CreateTestContext(recorder)

		// Create a file to upload
		body := new(bytes.Buffer)
		writer := multipart.NewWriter(body)
		part, _ := writer.CreateFormFile("file", "test.jpg")
		io.Copy(part, bytes.NewReader([]byte("dummy content")))
		writer.Close()

		request, _ := http.NewRequest(http.MethodPost, "/upload", body)
		request.Header.Set("Content-Type", writer.FormDataContentType())
		context.Request = request

		// Execute
		FileUpload(context)

		// Assert
		assert.Equal(t, http.StatusOK, recorder.Code)

		// For this example, replace "<url_of_uploaded_file>" with the actual expected URL.
		assert.Contains(t, recorder.Body.String(), `"status_code":200`)
		assert.Contains(t, recorder.Body.String(), `"message":"success"`)
		assert.Contains(t, recorder.Body.String(), `"data":{"data":"http://`)
	})

	t.Run("failure - no file selected", func(t *testing.T) {
		// Setup
		recorder := httptest.NewRecorder()
		context, _ := gin.CreateTestContext(recorder)

		request, _ := http.NewRequest(http.MethodPost, "/upload", nil)
		context.Request = request

		// Execute
		FileUpload(context)

		// Assert
		assert.Equal(t, http.StatusInternalServerError, recorder.Code)
		expectedResponse := `{"status_code":500,"message":"error","data":{"data":"Select a file to upload"}}`
		assert.JSONEq(t, expectedResponse, recorder.Body.String())
	})

	t.Run("failure - file upload error", func(t *testing.T) {
		// Since we are not mocking, this test will only work if the infrastructure's FileUpload
		// can fail in a controlled manner in your environment.

		// Setup
		recorder := httptest.NewRecorder()
		context, _ := gin.CreateTestContext(recorder)

		// Create a file to upload
		body := new(bytes.Buffer)
		writer := multipart.NewWriter(body)
		part, _ := writer.CreateFormFile("file", "test.jpg")
		io.Copy(part, bytes.NewReader([]byte("dummy content")))
		writer.Close()

		request, _ := http.NewRequest(http.MethodPost, "/upload", body)
		request.Header.Set("Content-Type", writer.FormDataContentType())
		context.Request = request

		// Execute
		FileUpload(context)

		// Assert
		// This test assumes the upload function may fail in your infrastructure.
		assert.Equal(t, http.StatusInternalServerError, recorder.Code)
		assert.Contains(t, recorder.Body.String(), `"status_code":500`)
		assert.Contains(t, recorder.Body.String(), `"message":"error"`)
	})
}
