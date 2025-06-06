package types

import v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type ResourceSliceInfo struct {
	Name              string                `json:"name"`
	NodeName          string                `json:"node_name"`
	CreationTimestamp v1.Time               `json:"creation_timestamp"`
	Driver            string                `json:"driver"`
	Devices           []ResourceSliceDevice `json:"devices"`
	Pool              string                `json:"pool"`
}

type ResourceSliceDevice struct {
	Name string `json:"name"`
	UUID string `json:"uuid"`
}
