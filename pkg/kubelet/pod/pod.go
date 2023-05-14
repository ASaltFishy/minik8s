package pod

import (
	"fmt"
	object2 "k8s/object"
)

func CreatePod(podConfig object2.Pod) error {
	// 拉取镜像
	var images []string
	for _, configItem := range podConfig.Spec.Containers {
		images = append(images, configItem.Image)
	}
	err := ListImage()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	err = PullImages(images)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	// 创建emptyDir数据卷（pod中的各个容器共享）
	_, err = createVolumes(podConfig.Spec.Volumes)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	// 创建pod中的容器并运行
	var containerMeta []object2.ContainerMeta
	containerMeta, err = CreateContainers(podConfig.Spec.Containers)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	podConfig.Spec.ContainerMeta = containerMeta

	// 打印容器信息
	for _, it := range containerMeta {
		fmt.Println(it.Name, " id:", it.ContainerID)
	}

	return nil
}

func StartPod(podConfig object2.Pod) error {
	containerMeta := podConfig.Spec.ContainerMeta
	// 开启容器
	for _, it := range containerMeta {
		StartContainer(it.ContainerID)
	}
	return nil
}

func ClosePod(podConfig object2.Pod) error {
	containerMeta := podConfig.Spec.ContainerMeta
	// 关闭容器
	for _, it := range containerMeta {
		StopContainer(it.ContainerID)
	}
	return nil
}

func RemovePod(podConfig object2.Pod) error {
	containerMeta := podConfig.Spec.ContainerMeta
	// 关闭容器
	for _, it := range containerMeta {
		StopContainer(it.ContainerID)
	}
	// 删除容器
	for _, it := range containerMeta {
		RemoveContainer(it.ContainerID)
	}
	return nil
}
