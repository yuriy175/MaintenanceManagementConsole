package utils

import (
	"io"
	"os"
	"log"
	"fmt"
	"archive/zip"
	"bytes"
	"io/ioutil"
	"compress/gzip"

	"../interfaces"
)

// logger implementation type
type logger struct {
	Path string
}

// LoggerNew creates an instance of logger
func LoggerNew() interfaces.ILogger {
	logImp := &logger{"server.log"}

	f, err := os.OpenFile(logImp.Path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	// defer f.Close()
	wrt := io.MultiWriter(os.Stdout, f)
	log.SetOutput(wrt)	
	
	return logImp
}

// Warning writes warning logs
func (t *logger) Warning(value string) {
	log.Println(value)
}

// Error writes error logs
func (t *logger) Error(value string) {
	log.Println(value)
}

// Info writes info logs
func (t *logger) Info(value string) {
	log.Println(value)
}

// Errorf writes formatted error logs
func (t *logger) Errorf(format string, a ...interface{}){
	t.Error(fmt.Sprintf(format, a...))
}

// Info writes formatted info logs
func (t *logger) Infof(format string, a ...interface{}){
	log := fmt.Sprintf(format, a...)
	t.Info(log)
}

// getZipContent returns zipped logs
func (t *logger) getZipContent() ([]byte, string){
	filePath := t.Path
	// Create a buffer to write our archive to.
	buf := new(bytes.Buffer)

	// Create a new zip archive.
	w := zip.NewWriter(buf)

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		t.Error("GetZipContent error")
		return nil, ""
	}

	f, err := w.Create(filePath)
	if err != nil {
		t.Error("GetZipContent error")
		return nil, ""
	}
	_, err = f.Write([]byte(data))
	if err != nil {
		t.Error("GetZipContent error")
		return nil, ""
	}
		

	// Make sure to check the error on Close.
	err = w.Close()
	if err != nil {
		t.Error("GetZipContent error")
		return nil, ""
	}

	// return buf.Bytes(), filePath
	return data, filePath
}

// WriteZipContent writes zipped logs to a writer
func (t *logger) WriteZipContent(w io.Writer) bool{
	logContent, filename := t.getZipContent()
	if logContent == nil || filename == ""{
		return false
	}
	writer, err := gzip.NewWriterLevel(w, gzip.BestCompression)
	if err != nil {
		t.Errorf("WriteZipContent error %v", err)
		return false
	}

	defer writer.Close()

	writer.Write(logContent)

	return true
}
