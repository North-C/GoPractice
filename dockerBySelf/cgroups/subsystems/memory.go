package subsystems

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strconv"
)

type MemorySubsystem struct{}

// 实现 Set 功能
func (s *MemorySubsystem) Set(cgroupPath string, res *ResourceConfig) error{
	// GetCgroupPath获取当前subsystem在虚拟文件系统中的路径
	if subsysCgroupPath, err := GetCgroupPath(s.Name(), cgroupPath, true); err!=nil{
		if res.MemoryLimit != ""{
			// 设置内存限制, 新版在这里需要修改，不再是写入 memory.limit_in_bytes
			if err := ioutil.WriteFile(path.Join(subsysCgroupPath, "memory.limit_in_bytes"), []byte(res.MemoryLimit), 0644); err != nil{
				return fmt.Errorf("set cgroup memory fail %v", err)
			}
		}
		return nil
	}else{
		return err
	}
}

// 删除 cgroupPath对应的cgroup
func (s *MemorySubsystem) Remove(cgroupPath string) error{
	// 这里对于cgroup v2 并不适用，subsystem的方法不对
	if subsysCgroupPath, err := GetCgroupPath(s.Name(), cgroupPath, false); err != nil{
		// 删除cgroup即删除对应的目录
		return os.Remove(subsysCgroupPath)
	}else{
		return err
	}
}

// 将一个进程加入cgroupPath 对应的 cgroup 中
func (s *MemorySubsystem) Apply(cgroupPath string, pid int) error{
	if subsysCgroupPath, err := GetCgroupPath(s.Name(), cgroupPath, false); err == nil{
		// 把进程的PID写到cgroup的虚拟文件系统对应目录下的 cgroup.procs 文件中
		if err := ioutil.WriteFile(path.Join(subsysCgroupPath, "cgroup.procs"), []byte(strconv.Itoa(pid)), 0644); err!= nil{
			return fmt.Errorf("set cgroup proc fail %v", err)
		}
		return nil
	}else{
		return fmt.Errorf("get crgoup %s error", err)
	}
}

func (s *MemorySubsystem) Name() string{
	return "memory"
}
