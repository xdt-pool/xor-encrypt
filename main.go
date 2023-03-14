package main

import (
	"encoding/hex"
	"flag"
	"io/ioutil"
	"os"
	"path"

	"github.com/kpango/glg"
)

var filePath string
var key string
var d byte
var dsSlice []byte

func main() {
	flag.StringVar(&filePath, "f", "", "raw information file path")
	flag.StringVar(&key, "k", "beehivekey", "encryption key (default:beehiveKey)")
	flag.Parse()

	if filePath == "" {
		glg.Error("error filepath")
		return
	}
	glg.Info("filepath:", filePath)
	glg.Info("key:", key)
	keyByte, err := hex.DecodeString(key)
	if err != nil {
		glg.Error("error key hex decode")
		return
	}
	glg.Info("keyByte:", keyByte)
	_, err = os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			//目录不存在
			glg.Error("raw information file not exist")
			return
		} else {
			panic("get raw info file error: " + err.Error())
		}
	}
	f, err := os.Open(filePath)
	if err != nil {
		panic("get raw info file error: " + err.Error())
	}
	buf, err := ioutil.ReadAll(f)
	if err != nil {
		glg.Error(err)
	}
	glg.Info("原文: ", string(buf))
	for i, v := range buf {
		d = byte(v ^ keyByte[i%len(keyByte)])
		dsSlice = append(dsSlice, d)
	}
	glg.Info("密文: ", string(dsSlice))

	dir, file := path.Split(filePath)
	outputPath := path.Join(dir, "en_"+file)
	glg.Info("outputPath: ", outputPath)
	err = os.WriteFile(outputPath, []byte(string(dsSlice)), 0644)
	if err != nil {
		glg.Error("write pull stream config error: ", err)
		return
	}

	// for i, v := range dsSlice {
	// 	d = byte(v ^ keyByte[i%len(keyByte)])
	// 	sSlice = append(sSlice, d)
	// }
	// glg.Info("原文: ", string(sSlice))
}
