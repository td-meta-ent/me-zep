package models

import (
	"context"
)

// MemoryStore interface
type MemoryStore[T any] interface {
	// GetMemory returns the most recent Summary and a list of messages for a given sessionID.
	// GetMemory returns:
	//   - the most recent Summary, if one exists
	//   - the lastNMessages messages, if lastNMessages > 0
	//   - all messages since the last SummaryPoint, if lastNMessages == 0
	//   - if no Summary (and no SummaryPoint) exists and lastNMessages == 0, returns
	//     all undeleted messages
	GetMemory(ctx context.Context,
		appState *AppState,
		sessionID string,
		lastNMessages int) (*Memory, error)
	// GetSummary retrieves the most recent Summary for a given sessionID. The Summary return includes the UUID of the
	// SummaryPoint, which the most recent Message in the collection of messages that was used to generate the Summary.
	GetSummary(ctx context.Context,
		appState *AppState,
		sessionID string) (*Summary, error)
	// PutMemory stores a Memory for a given sessionID. If the SessionID doesn't exist, a new one is created.
	PutMemory(ctx context.Context,
		appState *AppState,
		sessionID string,
		memoryMessages *Memory,
		skipNotify bool) error // skipNotify is used to prevent loops when calling NotifyExtractors.
	// PutSummary stores a new Summary for a given sessionID.
	PutSummary(ctx context.Context,
		appState *AppState,
		sessionID string,
		summary *Summary) error
	// PutMessageMetadata creates, updates, or deletes metadata for a given message, and does not
	// update the message itself.
	// isPrivileged indicates whether the caller is privileged to add or update system metadata.
	PutMessageMetadata(ctx context.Context,
		appState *AppState,
		sessionID string,
		messageMetaSet []MessageMetadata,
		isPrivileged bool) error
	// PutMessageVectors stores a collection of Embeddings for a given sessionID.
	PutMessageVectors(ctx context.Context,
		appState *AppState,
		sessionID string,
		embeddings []Embeddings) error
	// GetMessageVectors retrieves a collection of Embeddings for a given sessionID.
	GetMessageVectors(ctx context.Context,
		appState *AppState,
		sessionID string) ([]Embeddings, error)
	// SearchMemory retrieves a collection of SearchResults for a given sessionID and query. Currently, the query
	// is a simple string, but this could be extended to support more complex queries in the future. The SearchResult
	// structure can include both Messages and Summaries. Currently, we only search Messages.
	SearchMemory(
		ctx context.Context,
		appState *AppState,
		sessionID string,
		query *SearchPayload,
		limit int) ([]SearchResult, error)
	// DeleteSession deletes all records for a given sessionID. This is a soft delete. Hard deletes will be handled
	// by a separate process or left to the implementation.
	DeleteSession(ctx context.Context, sessionID string) error
	// OnStart is called when the application starts. This is a good place to initialize any resources or configs that
	// are required by the MemoryStore implementation.
	OnStart(ctx context.Context, appState *AppState) error
	// Attach is used by Extractors to register themselves with the MemoryStore. This allows the MemoryStore to notify
	// the Extractors when new occur.
	Attach(observer Extractor)
	// NotifyExtractors notifies all registered Extractors of a new MessageEvent.
	NotifyExtractors(
		ctx context.Context,
		appState *AppState,
		eventData *MessageEvent,
	)
	// Close is called when the application is shutting down. This is a good place to clean up any resources used by
	// the MemoryStore implementation.
	Close() error
}
