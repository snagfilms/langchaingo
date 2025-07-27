package mongo

import (
	"errors"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

const (
	mongoDefaultDBName         = "chat_history"
	mongoDefaultCollectionName = "message_store"
)

var (
	errMongoInvalidURL         = errors.New("invalid mongo url option")
	errMongoInvalidURLOrClient = errors.New("invalid mongo url or mongo client option")
	errMongoInvalidSessionID   = errors.New("invalid mongo session id option")
)

type ChatMessageHistoryOption func(m *ChatMessageHistory)

func applyMongoDBChatOptions(options ...ChatMessageHistoryOption) (*ChatMessageHistory, error) {
	h := &ChatMessageHistory{
		databaseName:   mongoDefaultDBName,
		collectionName: mongoDefaultCollectionName,
	}

	for _, option := range options {
		option(h)
	}

	if h.url == "" && h.client == nil {
		return nil, errMongoInvalidURLOrClient
	}
	if h.sessionID == "" {
		return nil, errMongoInvalidSessionID
	}

	return h, nil
}

// WithConnectionURL is an option for specifying the MongoDB connection URL. Must be set.
func WithConnectionURL(connectionURL string) ChatMessageHistoryOption {
	return func(p *ChatMessageHistory) {
		p.url = connectionURL
	}
}

// WithClient is an option for specifying the MongoDB existing client. Must be set.
func WithClient(client *mongo.Client) ChatMessageHistoryOption {
	return func(p *ChatMessageHistory) {
		p.client = client
	}
}

// WithCollection is an option for specifying the MongoDB existing client. Must be set.
func WithCollection(collection *mongo.Collection) ChatMessageHistoryOption {
	return func(p *ChatMessageHistory) {
		p.collection = collection
	}
}

// WithSessionID is an arbitrary key that is used to store the messages of a single chat session,
// like user name, email, chat id etc. Must be set.
func WithSessionID(sessionID string) ChatMessageHistoryOption {
	return func(p *ChatMessageHistory) {
		p.sessionID = sessionID
	}
}

// WithCollectionName is an option for specifying the collection name.
func WithCollectionName(name string) ChatMessageHistoryOption {
	return func(p *ChatMessageHistory) {
		p.collectionName = name
	}
}

// WithDataBaseName is an option for specifying the database name.
func WithDataBaseName(name string) ChatMessageHistoryOption {
	return func(p *ChatMessageHistory) {
		p.databaseName = name
	}
}
