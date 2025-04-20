package zonotools

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestGetQueryParam(t *testing.T) {
	t.Run("normal", func(t *testing.T) {
		// Create a request with query parameters
		req, err := http.NewRequest("GET", "http://example.com?param1=value1&param2=value2", nil)
		assert.NoError(t, err)

		// Test getting a parameter that exists
		value := GetQueryParam(req, "param1")
		assert.Equal(t, "value1", value)

		// Test getting a parameter that exists
		value = GetQueryParam(req, "param2")
		assert.Equal(t, "value2", value)

		// Test getting a parameter that doesn't exist
		value = GetQueryParam(req, "param3")
		assert.Equal(t, "", value)
	})
}

func TestGetFormValue(t *testing.T) {
	t.Run("normal", func(t *testing.T) {
		// Create form data
		form := url.Values{}
		form.Add("field1", "value1")
		form.Add("field2", "value2")

		// Create a request with form data
		req, err := http.NewRequest("POST", "http://example.com", strings.NewReader(form.Encode()))
		assert.NoError(t, err)
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		req.ParseForm()

		// Test getting a field that exists
		value := GetFormValue(req, "field1")
		assert.Equal(t, "value1", value)

		// Test getting a field that exists
		value = GetFormValue(req, "field2")
		assert.Equal(t, "value2", value)

		// Test getting a field that doesn't exist
		value = GetFormValue(req, "field3")
		assert.Equal(t, "", value)
	})
}

func TestGetFormFile(t *testing.T) {
	t.Run("normal", func(t *testing.T) {
		// Create a buffer to write our multipart form to
		var b bytes.Buffer
		w := multipart.NewWriter(&b)

		// Create a form file field
		fileContents := "test file contents"
		fw, err := w.CreateFormFile("file", "test.txt")
		assert.NoError(t, err)
		_, err = fw.Write([]byte(fileContents))
		assert.NoError(t, err)

		// Close the writer
		w.Close()

		// Create a request with the form data
		req, err := http.NewRequest("POST", "http://example.com", &b)
		assert.NoError(t, err)
		req.Header.Set("Content-Type", w.FormDataContentType())

		// Parse the multipart form
		err = req.ParseMultipartForm(10 << 20) // 10 MB
		assert.NoError(t, err)

		// Test getting the file
		file, header, err := GetFormFile(req, "file")
		assert.NoError(t, err)
		assert.NotNil(t, file)
		assert.NotNil(t, header)
		assert.Equal(t, "test.txt", header.Filename)

		// Read the file contents
		data, err := io.ReadAll(file)
		assert.NoError(t, err)
		assert.Equal(t, fileContents, string(data))

		// Test getting a file that doesn't exist
		_, _, err = GetFormFile(req, "nonexistent")
		assert.Error(t, err)
	})
}

func TestGetRequestBody(t *testing.T) {
	t.Run("normal", func(t *testing.T) {
		// Create a request with a body
		body := "test request body"
		req, err := http.NewRequest("POST", "http://example.com", strings.NewReader(body))
		assert.NoError(t, err)

		// Test getting the request body
		data, err := GetRequestBody(req)
		assert.NoError(t, err)
		assert.Equal(t, body, string(data))
	})

	t.Run("empty body", func(t *testing.T) {
		// Create a request with an empty body
		req, err := http.NewRequest("GET", "http://example.com", nil)
		assert.NoError(t, err)

		// Test getting an empty request body
		data, err := GetRequestBody(req)
		assert.NoError(t, err)
		assert.Equal(t, "", string(data))
	})
}

func TestCopyFile(t *testing.T) {
	t.Run("normal", func(t *testing.T) {
		// Create a temporary directory for test files
		tempDir, err := os.MkdirTemp("", "copyfile_test")
		assert.NoError(t, err)
		defer os.RemoveAll(tempDir)

		// Create a multipart form with a file
		var b bytes.Buffer
		w := multipart.NewWriter(&b)
		fileContents := "test file contents"
		fw, err := w.CreateFormFile("file", "test.txt")
		assert.NoError(t, err)
		_, err = fw.Write([]byte(fileContents))
		assert.NoError(t, err)
		w.Close()

		// Create a request with the form data
		req, err := http.NewRequest("POST", "http://example.com", &b)
		assert.NoError(t, err)
		req.Header.Set("Content-Type", w.FormDataContentType())

		// Parse the multipart form
		err = req.ParseMultipartForm(10 << 20) // 10 MB
		assert.NoError(t, err)

		// Get the file from the form
		file, _, err := req.FormFile("file")
		assert.NoError(t, err)
		defer file.Close()

		// Define the destination path
		destPath := filepath.Join(tempDir, "copied_test.txt")

		// Test copying the file
		err = CopyFile(file, destPath)
		assert.NoError(t, err)

		// Verify the file was copied correctly
		copiedData, err := os.ReadFile(destPath)
		assert.NoError(t, err)
		assert.Equal(t, fileContents, string(copiedData))
	})

	t.Run("large file", func(t *testing.T) {
		// Create a temporary directory for test files
		tempDir, err := os.MkdirTemp("", "copyfile_test_large")
		assert.NoError(t, err)
		defer os.RemoveAll(tempDir)

		// Create a temporary source file
		srcPath := filepath.Join(tempDir, "source.dat")
		destPath := filepath.Join(tempDir, "dest.dat")
		
		// Create a file with content larger than the buffer size (8192 bytes)
		size := 10000
		data := make([]byte, size)
		for i := 0; i < size; i++ {
			data[i] = byte(i % 256)
		}
		
		err = os.WriteFile(srcPath, data, 0666)
		assert.NoError(t, err)
		
		// Open the source file
		srcFile, err := os.Open(srcPath)
		assert.NoError(t, err)
		defer srcFile.Close()
		
		// Test copying the file
		err = CopyFile(srcFile, destPath)
		assert.NoError(t, err)
		
		// Verify the file was copied correctly
		copiedData, err := os.ReadFile(destPath)
		assert.NoError(t, err)
		assert.Equal(t, data, copiedData)
	})
}
