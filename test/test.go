package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	const (
		A = iota
		B = iota
		C
		D = 6
		E = iota
	)
	fmt.Println(A, B, C, D, E)
	// time.Sleep(100)
	end := time.Now()
	fmt.Println(end.Sub(start))

	overSeasJsonFile := "wuwa-overseas-md5.json"
	filesHash("E:\\games\\wuwa-overseas\\Wuthering Waves Game", overSeasJsonFile)

	cnJsonFile := "wuwa-cn-bili-md5.json"
	filesHash("E:\\games\\Wuthering Waves\\Wuthering Waves Game", cnJsonFile)

	// saveName := "./wwSources.json"
	// results := map[string]string{}
	// if err := readResultFromJSON(saveName, &results); err != nil {
	// 	slog.Error("读取json文件失败", "err", err)
	// 	return
	// }
	// slog.Info("results", "results", results)

}

// FileMD5 存储文件名和对应的MD5值
type FileMD5 struct {
	Path string `json:"path"`
	Hash string `json:"hash"`
}

func filesHash(root, saveName string) {
	files, err := getAllFiles(root)
	if err != nil {
		fmt.Printf("Error getting files: %v\n", err)
		return
	}

	// 创建通道用于传递结果
	resultChan := make(chan FileMD5, len(files))
	var wg sync.WaitGroup

	// 限制并发goroutine数量
	maxWorkers := runtime.NumCPU()
	// 利用通道满时写入会阻塞来控制并发数量量
	semaphore := make(chan struct{}, maxWorkers)

	// 为每个文件启动一个goroutine计算MD5
	for i, file := range files {
		wg.Add(1)
		fmt.Printf("正在处理%v/%v\n", i, len(files))
		semaphore <- struct{}{} // 获取信号量
		go func(f string, index int) {
			defer wg.Done()
			defer func() { <-semaphore }() // 释放信号量

			hash, err := calculateMD5(f)
			if err != nil {
				fmt.Printf("Error calculating MD5 for %s: %v\n", f, err)
				return
			}
			fmt.Printf("已处理第%v个文件\n", index)
			resultChan <- FileMD5{Path: f, Hash: hash}
		}(file, i)
	}

	// 等待所有goroutine完成
	wg.Wait()
	close(resultChan)

	// 收集并打印结
	results := make([]FileMD5, len(files))

	index := 0
	for result := range resultChan {
		// fmt.Printf("%s  %s\n", result.Hash, result.Path)
		result.Path, _ = filepath.Rel(root, result.Path)
		results[index] = result
		index++
	}
	writeResultsToJSON(saveName, results)
}

// getAllFiles 递归获取目录下所有文件路径
func getAllFiles(root string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}

// calculateMD5 计算文件的MD5值
func calculateMD5(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	return hex.EncodeToString(hash.Sum(nil)), nil
}

// writeResultsToJSON 将结果写入JSON文件
func writeResultsToJSON(filename string, results any) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // 美化输出，缩进两个空格
	return encoder.Encode(results)
}

func readResultFromJSON(filename string, results *[]FileMD5) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}

	defer file.Close()
	decoder := json.NewDecoder(file)
	return decoder.Decode(results)
}
