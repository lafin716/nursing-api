package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

type TakeHistoryMemo struct {
	ent.Schema
}

func (TakeHistoryMemo) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.UUID("user_id", uuid.UUID{}),
		field.UUID("timezone_id", uuid.UUID{}),
		field.String("take_date").SchemaType(varchar(10)),
		field.String("memo").Optional().SchemaType(text()),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Optional(),
	}
}

func (TakeHistoryMemo) Edges() []ent.Edge {
	return nil
}
