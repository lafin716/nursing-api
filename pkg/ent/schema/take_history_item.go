package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
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
		field.UUID("prescription_item_id", uuid.UUID{}).Optional(),
		field.Bool("take_status").Default(false),
		field.Float("take_amount").Default(0.0),
		field.Float("remain_amount").Default(0.0),
		field.Float("total_amount").Default(0.0),
		field.String("take_unit").SchemaType(varchar(2)),
		field.Text("memo").Optional(),
		field.Time("take_date"),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Optional(),
	}
}

func (TakeHistoryItem) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("prescription_item", PrescriptionItem.Type).
			Ref("take_history_item").
			Unique().
			Field("prescription_item_id"),
	}
}
