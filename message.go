// Package go_pubsub
// @Title  go-pubsub
// @Description  消息数据
// @Author  zxx1224@gmail.com  2022/4/12 3:33 PM
// @Update  zxx1224@gmail.com  2022/4/12 3:33 PM
package go_pubsub

// Message The message metadata
type Message struct {
	topic     string
	payload   interface{}
	createdAt int64
}

// GetTopic Return the topic of the current message
func (m *Message) GetTopic() string {
	return m.topic
}

// GetPayload Get the payload of the current message
func (m *Message) GetPayload() interface{} {
	return m.payload
}

// GetCreatedAt Get the creation time of this message
func (m *Message) GetCreatedAt() int64 {
	return m.createdAt
}
