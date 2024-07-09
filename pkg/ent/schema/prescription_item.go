package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

type PrescriptionItem struct {
	ent.Schema
}

func (PrescriptionItem) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.UUID("prescription_id", uuid.UUID{}),
		field.UUID("timezone_id", uuid.UUID{}),
		field.UUID("medicine_id", uuid.UUID{}),
		field.String("medicine_name").SchemaType(varchar(50)),
		field.String("timezone_name").Optional().SchemaType(varchar(50)),
		field.String("midday").SchemaType(varchar(2)),
		field.String("hour").SchemaType(varchar(2)),
		field.String("minute").SchemaType(varchar(2)),
		field.Float("total_amount").Default(0.0),
		field.Float("remain_amount").Default(0.0),
		field.Float("take_amount").Default(0.0),
		field.String("medicine_unit").Optional().Default("개").SchemaType(varchar(3)),
		field.String("memo").Optional().SchemaType(text()),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Optional(),
	}
}

func (PrescriptionItem) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("prescription", Prescription.Type).
			Ref("prescription_items").
			Unique().
			Required().
			Field("prescription_id"),
		edge.To("take_history_item", TakeHistoryItem.Type).
			Unique(),
	}
}
