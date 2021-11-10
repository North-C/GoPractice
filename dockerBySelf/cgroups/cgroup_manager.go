package cgroups

import (
	"dockerBySelf/cgroups/subsystems"
	"github.com/sirupsen/logrus"
)

type CgroupManager struct{
	// cgroup 在hierarchy当中的路径
	Path string
	// 资源配置
	Resource *subsystems.ResourceConfig
}

func NewCgroupManager(path string) *CgroupManager{
	return &CgroupManager{
		Path: path,
	}
}

// 加入 pid 到 cgroup中
func (c *CgroupManager) Apply(pid int) error{
	// SubsystemIns 是 subsystem的 实例
	for _, subSysIns := range(subsystems.SubsystemIns){
		subSysIns.Apply(c.Path, pid)
	}
	return nil
}

// 设置各个 subsystem 挂载中的 cgroup 资源限制
func (c *CgroupManager) Set(res *subsystems.ResourceConfig) error{
	for _, subsystems := range(subsystems.SubsystemIns){
		subsystems.Set(c.Path, res)
	}
	return nil
}

// 释放各个subsystem挂载中的cgroup
func (c *CgroupManager) Destroy() error{
	for _, subSysIns := range(subsystems.SubsystemIns){
		if err := subSysIns.Remove(c.Path); err != nil{
			logrus.Warnf("remove cgroup fail %v", err)
		}
	}
	return nil
}



