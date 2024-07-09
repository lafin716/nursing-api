package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

type Token struct {
	ent.Schema
}

func (Token) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("user_id", uuid.UUID{}),
		field.String("access_token").SchemaType(text()),
		field.String("refresh_token").SchemaType(text()),
		field.Time("access_token_expires"),
		field.Time("refresh_token_expires"),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Optional(),
	}
}
