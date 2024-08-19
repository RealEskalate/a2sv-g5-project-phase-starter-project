package custommongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	// Adjust the import path to your actual project structure
)

// MongoCursor wraps a *mongo.Cursor and implements the domain.Cursor interface.
type MongoCursor struct {
	*mongo.Cursor
}

func (m *MongoCursor) All(ctx context.Context, results interface{}) error {
	return m.Cursor.All(ctx, results)
}

func (m *MongoCursor) Next(ctx context.Context) bool {
	return m.Cursor.Next(ctx)
}

func (m *MongoCursor) Decode(val interface{}) error {
	return m.Cursor.Decode(val)
}

func (m *MongoCursor) Close(ctx context.Context) error {
	return m.Cursor.Close(ctx)
}

func (m *MongoCursor) Err() error {
	return m.Cursor.Err()
}
