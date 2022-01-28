package c2file

import (
	"bufio"
	"encoding/binary"
	"io/ioutil"
	"os"
)

type FileUtils struct {
	file    *os.File
	bufior  *bufio.Reader
	bufiow  *bufio.Writer
	bufiorw *bufio.ReadWriter
}

func CreateFileUtil(filePath string) *FileUtils {

	file, err := os.OpenFile(filePath, os.O_RDWR, 666)
	if err != nil {
		// log.Println(err)
		return nil
	}
	reader := bufio.NewReader(file)
	writer := bufio.NewWriter(file)
	readWriter := bufio.NewReadWriter(reader, writer)
	return &FileUtils{
		file:    file,
		bufior:  reader,
		bufiow:  writer,
		bufiorw: readWriter,
	}
}

func (fileIoUtil *FileUtils) WriteBytesByBufio(bytes []byte) (error, int) {
	numBytes, err := fileIoUtil.bufiorw.Write(bytes)
	if err != nil {
		// log.Println(err)
		return err, 0
	}
	return nil, numBytes
}

func (fileIoUtil *FileUtils) ReadBytesByBufio(numBytes int) (*[]byte, error) {
	bytes := make([]byte, numBytes)
	fileIoUtil.bufiorw.Read(bytes)
	return &bytes, nil
}

func (fileIoUtil *FileUtils) WriteStringByBufio(pString string) (error, int) {
	nn, err := fileIoUtil.bufiorw.WriteString(pString)
	if err != nil {
		// log.Println(err)
		return err, 0
	}
	return nil, nn
}

func (fileIoUtil *FileUtils) ReadStringByBufio(delim byte) (string, error) {
	pString, err := fileIoUtil.bufiorw.ReadString(delim)
	if err != nil {
		return "", err
	}
	return pString, nil
}

func (fileIoUtil *FileUtils) ReadStringByBufioScaner(processFunc func(string)) {
	scanner := bufio.NewScanner(fileIoUtil.file)
	for scanner.Scan() {
		// fmt.Println(scanner.Text())
		processFunc(scanner.Text())
	}
}

func (fileIoUtil *FileUtils) Write(bytes []byte) (error, int) {
	numBytes, err := fileIoUtil.file.Write(bytes)
	if err != nil {
		// log.Println(err)
		return err, 0
	}
	return nil, numBytes
}

func (fileIoUtil *FileUtils) Read(numBytes int) (*[]byte, error) {
	bytes := make([]byte, numBytes)
	fileIoUtil.file.Read(bytes)
	return &bytes, nil
}

func (fileIoUtil *FileUtils) WriteString(bytes []byte) (error, int) {
	numBytes, err := fileIoUtil.file.Write(bytes)
	if err != nil {
		// log.Println(err)
		return err, 0
	}
	return nil, numBytes
}

func (fileIoUtil *FileUtils) ReadByIoUtil() (*[]byte, error) {
	bytes, err := ioutil.ReadAll(fileIoUtil.file)
	if err != nil {
		// log.Println(err)
		return nil, err
	}
	return &bytes, nil
}

func (fileIoUtil *FileUtils) ReadByBinary(data interface{}) error {
	err := binary.Read(fileIoUtil.file, binary.LittleEndian, data)
	if err != nil {
		return err
	}
	return nil
}

func (fileIoUtil *FileUtils) WriteByBinary(data interface{}) error {
	err := binary.Write(fileIoUtil.file, binary.LittleEndian, data)
	if err != nil {
		return err
	}
	return nil
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

/**
 * 判断文件是否存在  存在返回 true 不存在返回false
 */
func CheckFileExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}
