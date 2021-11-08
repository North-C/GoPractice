package subsystems


type CpuSubsystem struct{

}


func (s *CpuSubsystem) Set(cgroupPath string, res *ResourceConfig) error{
	if 
}

func (s *CpuSubsystem) Remove(path string) error{
	
}

func (s *CpuSubsystem) Apply(path string, pid int) error{

}


func (s *CpuSubsystem) Name() string{
	return "cpu"
}
