package subsystems

import (
	"fmt"
	"io/ioutil"
	"path"
	"os"
	"strconv"
)

type CpusetSubsystem struct{

}

func (s *CpusetSubsystem) Set(cgroupPath string, res *ResourceConfig) error{
	if subsystemCgroupPath, err := GetCgroupPath(s.Name(), cgroupPath, true); err == nil{
		if res.CpuSet != ""{
			// 通过cpu的编号来进行绑定
			if err := ioutil.WriteFile(path.Join(subsystemCgroupPath, "cpuset.cpus"), []byte(res.CpuSet), 0644); err != nil{
				return fmt.Errorf("set cgroup cpuset fail %v", err)
			}
		}
		return nil
	}else{
		return err
	}
}

func (s * CpusetSubsystem) Remove(cgroupPath string) error{
	if subsystemCgroupPath, err := GetCgroupPath(s.Name(), cgroupPath, false); err == nil{
		return os.RemoveAll(subsystemCgroupPath)
	}else{
		return err
	}
}

func (s *CpusetSubsystem) Apply(cgroupPath string, pid int) error{
	if subsystemCgroupPath, err := GetCgroupPath(s.Name(), cgroupPath, false); err == nil{
		if err := ioutil.WriteFile(path.Join(subsystemCgroupPath, "tasks"), []byte(strconv.Itoa(pid)), 0644); err != nil{
			return fmt.Errorf("set cgroup proc fail %v", err)
		}
		return nil
	}else{
		return fmt.Errorf("get cgroup %s error: %v", cgroupPath, err)
	}

}

func (s *CpusetSubsystem) Name()string{
	return "cpuset"
}
