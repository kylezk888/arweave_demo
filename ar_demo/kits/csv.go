package kits

import (
	"bufio"
	"os"
	"strings"
)

func LoadCSV(fpath string) ([]string, error) {
	list := []string{}
	//wd, err := os.Getwd()
	//if err != nil {
	//	return nil, err
	//}
	////
	//fpath = filepath.Join(wd, fpath)
	dataCSV, err := os.Open(fpath)
	if err != nil {
		return nil, err
	}
	defer dataCSV.Close()
	//
	scan := bufio.NewScanner(dataCSV)
	for scan.Scan() {
		line := strings.TrimSpace(scan.Text())
		list = append(list, line)
	}
	return list, nil
}
