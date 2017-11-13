package config

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"io/ioutil"
	"log"
)

// Load settings from disk
func Load() map[string]string {
	var config map[string]string
	b := loadFromDisk("settings")
	jsonText := decompress(b)
	err := json.Unmarshal(jsonText, &config)
	if err != nil {
		log.Println(err)
	}

	return config
}

// LoadJSON settings from disk as JSON
func LoadJSON() string {
	b := loadFromDisk("settings")
	jsonText := decompress(b)

	return string(jsonText)
}

// Save settings to disk
func Save(config map[string]string) {
	b, err := json.Marshal(config)
	if err != nil {
		log.Println(err)
	}
	b2 := compress(b)

	saveToDisk("settings", b2.Bytes())
}

func compress(json []byte) bytes.Buffer {
	var b bytes.Buffer
	gz := gzip.NewWriter(&b)
	if _, err := gz.Write(json); err != nil {
		log.Panic(err)
	}
	if err := gz.Flush(); err != nil {
		log.Panic(err)
	}
	if err := gz.Close(); err != nil {
		log.Panic(err)
	}
	return b
}

func decompress(data []byte) []byte {
	if data == nil {
		return nil
	}
	rdata := bytes.NewReader(data)
	r, _ := gzip.NewReader(rdata)
	s, _ := ioutil.ReadAll(r)

	return s
}

func saveToDisk(filename string, b []byte) {
	err := ioutil.WriteFile(filename, b, 0644)
	if err != nil {
		log.Panic(err)
	}
}

func loadFromDisk(filename string) []byte {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil
	}

	return data
}
