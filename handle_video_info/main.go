package main

// 录制的视频元数据被更改, 在vlc中播放时标题被改变, 此为重改元数据程序
import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	var filePath, handleType string
	flag.StringVar(&filePath, "path", "", "目录路径")
	flag.StringVar(&handleType, "type", "edit", "操作类型, edit 更改元信息, mv 重命名")
	flag.Parse()
	fmt.Println(filePath)
	if filePath == "" {
		log.Fatalln("处理路径path参数不能为空")
		os.Exit(1)
	}

	path, err := exec.LookPath("ffmpeg")
	if err != nil {
		log.Fatal(err, "\n请先安装ffmpeg软件,或者将ffmpeg添加到PATH中")
	} else {
		fmt.Printf("ffmpeg is available as %s\n", path)
	}

	matches, err := filepath.Glob(filePath + "*/*.mp4")
	if err != nil {
		fmt.Println(err)
	}

	// fmt.Println(matches)
	// return
	tmpAfterFix := "--golangback.mp4"

	if handleType == "edit" {
		for _, filestr := range matches {
			if strings.Index(filestr, tmpAfterFix) != -1 {
				// 已处理过的文件不再处理
				continue
			}
			target := strings.Replace(filestr, ".mp4", tmpAfterFix, 1)
			// fmt.Println(target)
			handle(filestr, target)
		}

	} else if handleType == "mv" {
		for _, filestr := range matches {
			if strings.Index(filestr, tmpAfterFix) == -1 {
				// 不需要处理的文件
				continue
			}
			target := strings.Replace(filestr, tmpAfterFix, ".mp4", 1)
			// fmt.Println(target)
			mv(filestr, target)
		}
	}

}

func handle(src, target string) {

	//ffmpeg -i '.\10-2 使用Channel等待任务结束.mp4' -c copy -metadata title='10-2 使用Channel等待任务结束' -metadata comment='golang' '10-2 使用Channel等待任务结束.mp4'
	cmd := exec.Command("ffmpeg", "-i", src, "-c", "copy", "-metadata", "title="+strings.Replace(filepath.Base(src), ".mp4", "", 1), "-metadata", "comment=golang handles", target)
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("has cp ", src, " to ", target)
}

func mv(src, target string) {
	if _, err := os.Stat(target); !os.IsNotExist(err) {
		// 目标文件已存在,先删除
		if err = os.Remove(target); err != nil {
			// 删除失败直接返回mv函数
			fmt.Println(target, " 删除失败", err)
			return
		}
		fmt.Println(target, "文件已删除")

	}

	// 重命名src -> target
	if err := os.Rename(src, target); err != nil {
		fmt.Println(src, "重命名失败", err)
		return
	}

}
