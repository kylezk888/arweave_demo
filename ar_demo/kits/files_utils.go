package kits

import (
	"bytes"
	"crypto/rand"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

type File struct {
	Reader    *bytes.Reader
	Path      string
	Size      int64
	Name      string
	Extension string
}

func GetFile(path string) (*File, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	stat, err := file.Stat()
	if err != nil {
		return nil, err
	}
	split := strings.Split(file.Name(), ".")
	readAll, err := io.ReadAll(file)
	reader := bytes.NewReader(readAll)
	return &File{reader, path, stat.Size(), stat.Name(), split[1]}, nil
}
func GetFileByBuffer(size int) bytes.Buffer {
	var buffer bytes.Buffer
	line := `1234567890,1234567890,1234567890,1234567890,1234567890,1234567890,1234567890,1234567890,1234567890,
	1234567890,1234567890,1234567890,1234567890,1234567890,1234567890,1234567890,1234567890,1234567890,1234567890,
	1234567890,1234567890,1234567890,1234567890,1234567890,1234567890,1234567890,1234567890,1234567890,1234567890,
	1234567890,1234567890,1234567890,1234567890,1234567890,1234567890,1234567890,1234567890,1234567890,1234567890,
	1234567890,1234567890,1234567890,1234567890,1234567890,1234567890,1234567890,1234567890,1234567890,1234567890,
	1234567890,1234567890,1234567890,1234567890,1234567890,1234567890,1234567890,1234567890,1234567890,1234567890,
	1234567890,1234567890,1234567890,1234567890,1234567890,1234567890,1234567890,1234567890,1234567890,1234567890,
	1234567890,1234567890,1234567890,1234567890,1234567890,1234567890,1234567890,1234567890,1234567890,1234567890,
	1234567890,1234567890,1234567890,1234567890,1234567890,1234567890,1234567890,1234567890,1234567890,1234567890,
	1234567890,1234567890,12345`
	// Create 1MiB content where each line contains 1024 characters.
	for i := 0; i < size; i++ {
		buffer.WriteString(fmt.Sprintf("%s", line))
	}
	return buffer
}

func GetFileByBufferAllsize(size int64, rands ...bool) bytes.Buffer {
	if len(rands) > 0 {
		b := make([]byte, size)
		_, err := rand.Read(b)
		if err != nil {
			panic(err)
		}
		var buffer bytes.Buffer
		buffer.Write(b)
		return buffer
	} else {
		kb := size / 1024
		b := size % 1024
		var buffer bytes.Buffer
		buffer = GetFileByBuffer(int(kb))
		line := "1"
		// Create 1MiB content where each line contains 1024 characters.
		for i := 0; i < int(b); i++ {
			buffer.WriteString(fmt.Sprintf("%s", line))
		}
		return buffer
	}
}

func AppendFile(buf bytes.Buffer, pre string) *bytes.Buffer {
	if buf.Len() < len(pre) {
		return bytes.NewBuffer([]byte(pre[0:buf.Len()]))
	} else {
		return bytes.NewBuffer(append(buf.Bytes()[0:buf.Len()-len(pre)], []byte(pre)...))
	}
}
func Getfilepath(path string) string {
	wd, _ := os.Getwd()
	wd, _, _ = strings.Cut(wd, "greenfield-integration-test")
	const dir = "file"
	result := filepath.Join(wd, "greenfield-integration-test", dir, path)
	return result
}
func GetSizeByString(size string) int64 {
	len := len(size)
	mul := int64(1)
	switch size[len-1] {
	case 'g':
		{
			mul = 1024 * 1024 * 1024
		}
	case 'm':
		{
			mul = 1024 * 1024
		}
	case 'k':
		{
			mul = 1024
		}
	case 'b':
		{
			mul = 1
		}
	}
	i, _ := strconv.Atoi(size[0 : len-1])
	return int64(i) * mul
}
func GetParamFromCSV(path string, skipFirstRow bool) [][]string {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	record, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}
	if skipFirstRow {
		record = record[1:]
	}
	return record
}
func WriteToCSV(data [][]string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	for _, d := range data {
		err := writer.Write(d)
		if err != nil {
			return err
		}
	}
	writer.Flush()
	return err
}
func CreateFile(filePath string) error {
	// 判断文件是否存在
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		// 如果文件不存在，获取文件所在目录
		dir := filepath.Dir(filePath)
		// 判断目录是否存在，不存在则创建目录
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			if err := os.MkdirAll(dir, 0755); err != nil {
				return err
			}
		}
		// 创建文件
		if _, err := os.Create(filePath); err != nil {
			return err
		}
	} else if err != nil {
		// 如果发生其他错误，则返回错误信息
		return err
	}
	return nil
}

func GetStats(timestamps []int) {
	sort.Ints(timestamps)

	n := len(timestamps)
	p90Index := int(float64(n) * 0.9)
	p90 := timestamps[p90Index]

	var sum int
	var max, min int

	for _, value := range timestamps {
		sum += value
		if value > max {
			max = value
		}
		if value < min || min == 0 {
			min = value
		}
	}

	avg := sum / n

	fmt.Println("p90:", p90)
	fmt.Println("Avg:", avg)
	fmt.Println("Max:", max)
	fmt.Println("Min:", min)
}
