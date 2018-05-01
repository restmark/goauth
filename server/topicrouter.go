package server

import (
	"github.com/restmark/goauth/handlers"
	"github.com/restmark/goauth/store"
	"github.com/restmark/goauth/store/mongodb"
)

// routes the received kafka message to its correct handler function
func (a *API) SetupTopicRouter() {
	topicRouter := NewTopicRouter(a)

	topicRouter.AddHandler(a.Config.GetString("kafka_topic"), handlers.HandleUserCreated)

	a.TopicRouter = topicRouter
}

type HandlerFunc func(store.Store, []byte)error //TODO: Replace *API with a context

type TopicRouter struct {
	store store.Store
	topicMap map[string]HandlerFunc
}

func NewTopicRouter(api *API) *TopicRouter {
	return &TopicRouter{
		mongodb.New(api.Database),
		make(map[string]HandlerFunc),
	}
}

func (t *TopicRouter) AddHandler(topic string, handler HandlerFunc) {
	t.topicMap[topic] = handler
}

func (t* TopicRouter) CallHandler(topic string, message []byte) error {
	err := t.topicMap[topic](t.store, message)
	if err != nil {
		return err
	}

	return nil
}