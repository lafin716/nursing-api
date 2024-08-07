package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
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
		field.UUID("prescription_id", uuid.UUID{}).Optional(),
		field.UUID("prescription_item_id", uuid.UUID{}).Optional(),
		field.UUID("timezone_id", uuid.UUID{}).Optional(),
		field.UUID("medicine_id", uuid.UUID{}).Optional(),
		field.String("medicine_name").SchemaType(varchar(50)),
		field.String("timezone_name").Optional().SchemaType(varchar(50)),
		field.String("midday").SchemaType(varchar(2)),
		field.String("hour").SchemaType(varchar(2)),
		field.String("minute").SchemaType(varchar(2)),
		field.Bool("take_status").Default(false),
		field.Float("total_amount").Default(0.0),
		field.Float("remain_amount").Default(0.0),
		field.Float("take_amount").Default(0.0),
		field.String("take_unit").SchemaType(varchar(2)),
		field.String("take_date").SchemaType(varchar(10)),
		field.String("take_time").SchemaType(varchar(8)),
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

func (TakeHistoryItem) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("take_date", "prescription_item_id"),
	}
}
