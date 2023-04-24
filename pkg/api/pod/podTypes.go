package pod

type ContainerPort struct {
	Port int32 `yaml:"containerPort"`
}

type ContainerVolumeMountConfig struct {
	Name      string `yaml:"name"`
	MountPath string `yaml:"mountPath"`
	ReadOnly  bool   `yaml:"readOnly"`
}

type ContainerResources struct {
	Cpu    string `yaml:"cpu"`
	Memory string `yaml:"memory"`
}

type ContainerResourcesConfig struct {
	Limits ContainerResources `yaml:"limits"`
}

type Container struct {
	Name         string                       `yaml:"name"`
	Image        string                       `yaml:"image"`
	Ports        []ContainerPort              `yaml:"ports"`
	Command      []string                     `yaml:"command"`
	Args         []string                     `yaml:"args"`
	Resources    ContainerResourcesConfig     `yaml:"resources"`
	VolumeMounts []ContainerVolumeMountConfig `yaml:"volumeMounts"`
}

type VolumeEmptyDirConfig struct {
}

type VolumeHostPathConfig struct {
	Path string `yaml:"path"`
}

type VolumeConfig struct {
	Name     string               `yaml:"name"`
	EmptyDir VolumeEmptyDirConfig `yaml:"emptyDir"`
	HostPath VolumeHostPathConfig `yaml:"hostPath"`
}

type Spec struct {
	Containers []Container    `yaml:"containers"`
	Volumes    []VolumeConfig `yaml:"volumes"`
}

type Metadata struct {
	Name   string `yaml:"name"`
	Labels Labels `yaml:"labels"`
}

type Labels struct {
	App string `yaml:"app"`
	Env string `yaml:"env"`
}

type ContainerMeta struct {
	Name        string
	ContainerID string
}

type Pod struct {
	ApiVersion string   `yaml:"apiVersion"`
	Kind       string   `yaml:"kind"`
	Metadata   Metadata `yaml:"metadata"`
	Spec       Spec     `yaml:"spec"`
}