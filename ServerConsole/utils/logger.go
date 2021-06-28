package utils

import (
	"io"
	"os"
	"log"
	"fmt"
	"archive/zip"
	"bytes"
	"io/ioutil"

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
	t.Error(fmt.Sprintf(format, a))
}

// Info writes formatted info logs
func (t *logger) Infof(format string, a ...interface{}){
	t.Info(fmt.Sprintf(format, a))
}

// GetZipContent returns zipped logs
func (t *logger) GetZipContent() ([]byte, string){
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

	return buf.Bytes(), filePath
}