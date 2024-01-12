package schema

import (
	"entgo.io/ent"
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
		field.UUID("user_id", uuid.UUID{}),
		field.UUID("prescription_id", uuid.UUID{}),
		field.String("medicine_name").SchemaType(varchar(50)),
		field.String("take_time_zone").Optional().SchemaType(varchar(10)),
		field.String("take_moment").Optional().SchemaType(varchar(10)),
		field.String("take_etc").Optional().SchemaType(varchar(50)),
		field.Float("take_amount").Default(0.0),
		field.String("medicine_unit").Optional().Default("개").SchemaType(varchar(3)),
		field.String("memo").Optional().SchemaType(text()),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Optional(),
	}
}
