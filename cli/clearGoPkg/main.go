package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"golang.org/x/mod/semver"
)

func extractModuleVersion(path string, modRoot string) (string, string, bool) {
	rel, err := filepath.Rel(modRoot, path)
	if err != nil {
		return "", "", false
	}
	if !strings.Contains(rel, "@") {
		return "", "", false
	}

	parts := strings.Split(rel, "@")
	if len(parts) != 2 {
		return "", "", false
	}

	return parts[0], parts[1], true
}

func normalize(v string) string {
	if !strings.HasPrefix(v, "v") {
		return "v" + v
	}
	return v
}

func main() {
	deleteMode := false
	if len(os.Args) > 1 && os.Args[1] == "--delete" {
		deleteMode = true
		fmt.Println("⚠️ 正在执行实际删除，请谨慎！")
	} else {
		fmt.Println("💡 Dry run 模式（不会删除），使用 --delete 启用实际删除")
	}

	// 获取模块缓存路径
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		gopath = filepath.Join(os.Getenv("HOME"), "go")
	}
	modCache := filepath.Join(gopath, "pkg", "mod")

	modVersions := make(map[string][]string)
	versionDirs := make(map[string]map[string]string) // [mod][version] => fullpath

	// 扫描模块版本
	err := filepath.WalkDir(modCache, func(path string, d os.DirEntry, err error) error {
		if err != nil || !d.IsDir() {
			return nil
		}
		if strings.Contains(d.Name(), "@") {
			modPath, version, ok := extractModuleVersion(path, modCache)
			if !ok {
				return nil
			}
			modVersions[modPath] = append(modVersions[modPath], version)
			if versionDirs[modPath] == nil {
				versionDirs[modPath] = make(map[string]string)
			}
			versionDirs[modPath][version] = path
			return filepath.SkipDir
		}
		return nil
	})

	if err != nil {
		fmt.Println("扫描失败:", err)
		return
	}

	// 开始处理删除逻辑
	for mod, versions := range modVersions {
		if len(versions) <= 1 {
			continue
		}

		// 排序版本（新版本在最后）
		sort.Slice(versions, func(i, j int) bool {
			return semver.Compare(normalize(versions[i]), normalize(versions[j])) < 0
		})

		latest := versions[len(versions)-1]

		fmt.Printf("\n📦 %s:\n", mod)
		for _, v := range versions {
			if v == latest {
				fmt.Printf("  ✅ 保留: %s\n", v)
				continue
			}
			fullPath := versionDirs[mod][v]
			if deleteMode {
				err := os.RemoveAll(fullPath)
				if err != nil {
					fmt.Printf("  ❌ 删除失败: %s (%v)\n", v, err)
				} else {
					fmt.Printf("  🗑️ 已删除: %s\n", v)
				}
			} else {
				fmt.Printf("  🗑️ 将删除: %s\n", v)
			}
		}
	}
}
