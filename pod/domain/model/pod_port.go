package model


type PodPort struct {
	//id
	ID int64 `gorm:"primary_key;not_null;auto_increment" json:"id"`
	//podid
	PodId int64 `json:"pod_id"`
	//容器端口
	ContainerPort int32 `json:"container_port"`
	//协议类型
	Protocol string `json:"protocol"`
}
