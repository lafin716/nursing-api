package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

type Prescription struct {
	ent.Schema
}

func (Prescription) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.UUID("user_id", uuid.UUID{}),
		field.String("prescription_name").Optional().SchemaType(varchar(30)),
		field.String("hospital_name").Optional().SchemaType(varchar(50)),
		field.Int("take_days").Default(0),
		field.Time("started_at").Optional().Default(time.Now).SchemaType(date()),
		field.Time("finished_at").Optional().Default(time.Now).SchemaType(date()),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Optional(),
	}
}

func (Prescription) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("prescription_items", PrescriptionItem.Type),
	}
}
