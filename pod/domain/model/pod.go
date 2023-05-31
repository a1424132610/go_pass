package model

type Pod struct {
	ID int64 `gorm:"primary_key;not_null;auto_increment" json:"id"`
	PodName string `gorm:"unique_index;not_null" json:"pod_name"`
	PodNamespace string `json:"pod_namespace"`
	//pod 所属的团队
	PodTeamID int64 `json:"pod_team_id"`
	//pod 使用cpu的最小值
	PodCpuMin float32 `json:"pod_cpu_min"`
	//pod 使用的cpu最大值
	PodCpuMax float32 `json:"pod_cpu_max"`
	//pod副本数量
	PodReplicas int32 `json:"pod_replicas"`
	//pod使用的内存最小值
	PodMemoryMin float32 `json:"pod_memory_min"`
	//pod使用的内存最大值
	PodMemoryMax float32 `json:"pod_memory_max"`
	//pod开放的端口
	PodPort[] PodPorts
}
