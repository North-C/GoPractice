package subsystems

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strings"
)

func FindCgroupMountPoint(subsystem string) string {
	f, err := os.Open("/proc/self/mountinfo")
	if err != nil {
		return ""
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() { // 读取到下一个token处
		txt := scanner.Text()             // 以string形式返回 Scan 读取到的token
		fields := strings.Split(txt, " ") // 按照空格分割
		// 对最后一项进行选择，以 , 再分割
		for _, opt := range strings.Split(fields[len(fields)-1], ",") {
			if opt == subsystem {
				return fields[4]
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return ""
	}
	return ""
}

func GetCgroupPath(subsystem string, cgroupPath string, autoCreate bool) (string, error) {
	// 找到 cgroup 的 root
	cgroupRoot := FindCgroupMountPoint(subsystem)
	// 
	if _, err := os.Stat(path.Join(cgroupRoot, cgroupPath)); err == nil || (autoCreate && os.IsNotExist(err)){
		if os.IsNotExist(err){
			// 不存在则 创建新的目录
			if err := os.Mkdir(path.Join(cgroupRoot, cgroupPath), 0755); err == nil{
			}else{
				return path.Join(cgroupRoot, cgroupPath), nil
			}
		}
		return path.Join(cgroupRoot, cgroupPath), nil
	}else{
		return "", fmt.Errorf("cgroup path error %v", err)
	}
}


