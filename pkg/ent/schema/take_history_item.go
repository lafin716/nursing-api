package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

type TakeHistoryItem struct {
	ent.Schema
}

func (TakeHistoryItem) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.UUID("user_id", uuid.UUID{}),
		field.UUID("take_history_id", uuid.UUID{}),
		field.UUID("prescription_item_id", uuid.UUID{}),
		field.String("take_status").Default("N").SchemaType(varchar(1)),
		field.Float("take_amount").Default(0.0),
		field.String("take_unit").SchemaType(varchar(2)),
		field.Text("memo").Optional(),
		field.Time("take_date"),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Optional(),
	}
}
