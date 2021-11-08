package subsystems

// 资源限制 
type ResourceConfig struct{
	MemoryLimit string		// 内存限制
	CpuShare string			// CPU时间片
	CpuSet string			// CPU核心数
}

// Subsystem 接口
// cgroup抽象成path,在hierarchy以虚拟文件系统中的虚拟路径出现
type Subsystem interface{
	Name() string			
	Set(path string, res *ResourceConfig) error
	// 将进程添加到某个cgroup中
	Apply(path string, pid int) error
	// 移除某个cgroup
	Remove(path string) error
}

// 创建实例
var(
	SubsystemIns = []Subsystem{
		&CpusetSubsystem{},
		&MemorySubsystem{},
		&CpuSubsystem{},
	}
)