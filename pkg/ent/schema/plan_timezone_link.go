package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

type PlanTimeZoneLink struct {
	ent.Schema
}

func (PlanTimeZoneLink) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.UUID("prescription_id", uuid.UUID{}),
		field.UUID("timezone_id", uuid.UUID{}),
		field.String("timezone_name").Optional().SchemaType(varchar(50)),
		field.Bool("use_alert").Default(false),
		field.String("meridiem").SchemaType(varchar(2)),
		field.String("hour").SchemaType(varchar(2)),
		field.String("minute").SchemaType(varchar(2)),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Optional(),
	}
}

func (PlanTimeZoneLink) Edges() []ent.Edge {
	return nil
}
