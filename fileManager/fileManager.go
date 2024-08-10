package fileManager

import (
	"encoding/json"
	"errors"
	"os"
)

type FileManagerType struct {
	outPutPath string
}

func (receiver FileManagerType) WriteJson(data any) error {
	file, err := os.Create(receiver.outPutPath)
	if err != nil {
		return errors.New("failed to create file")
	}

	encoderVal := json.NewEncoder(file)
	err = encoderVal.Encode(data)
	if err != nil {
		return errors.New("data passed cannot be converted to JSON")
		file.Close()
	}

	file.Close()
	return nil
}

func New(outPutPath string) FileManagerType {
	return FileManagerType{outPutPath: outPutPath}
}
