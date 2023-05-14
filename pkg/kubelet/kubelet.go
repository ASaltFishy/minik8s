package kubelet

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"k8s/object"
	"k8s/pkg/global"
	"k8s/pkg/util/HTTPClient"
	"k8s/pkg/util/msgQueue/subscriber"
)

const serverHost = "http://127.0.0.1:8080"

type Kubelet struct {
	client        *HTTPClient.Client
	node          object.Node
	podSubscriber *subscriber.Subscriber
	podQueue      string
}

type podHandler struct {
}

func (h podHandler) Handle(pod []byte) {
	// TODO: 监听到集群状态变化的处理函数
	fmt.Printf(string(pod))
}

// Run kubelet运行的入口函数
func (kl *Kubelet) Run() {
	// TODO 发送请求获取pod列表，可能需要缓存在本地？

	// 开始监听消息队列中pod的增量信息
	handler := podHandler{}
	err := kl.podSubscriber.Subscribe(kl.podQueue, subscriber.Handler(handler))
	if err != nil {
		fmt.Printf(err.Error())
		kl.podSubscriber.CloseConnection()
	}
}

// NewKubelet kubelet对象的构造函数
func NewKubelet(name string) (*Kubelet, error) {
	// 使用HTTP，构建node对象传递到APIServer处
	client := HTTPClient.CreateHTTPClient(serverHost)
	id, _ := uuid.NewUUID()
	nodeInfo := object.Node{
		Metadata: object.Metadata{
			Name:      name,
			Namespace: "default",
			Uid:       id.String(),
		},
		// TODO 根据需要填入或生成IP
		IP: "127.0.0.1",
	}
	info, _ := json.Marshal(nodeInfo)
	podQueue := client.Post("/nodes/create", info)
	fmt.Println("get response from APIServer" + podQueue)

	// 建立消息监听队列
	sub, _ := subscriber.NewSubscriber(global.MQHost)

	// 创建kubelet监听队列
	kub := Kubelet{
		client:        client,
		node:          nodeInfo,
		podSubscriber: sub,
		podQueue:      podQueue,
	}
	return &kub, nil
}