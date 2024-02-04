package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

type PlanTimeZone struct {
	ent.Schema
}

func (PlanTimeZone) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.UUID("user_id", uuid.UUID{}),
		field.String("timezone_name").Optional().SchemaType(varchar(50)),
		field.Bool("is_default").Default(false),
		field.Bool("use_alert").Default(false),
		field.String("meridiem").SchemaType(varchar(2)),
		field.String("hour").SchemaType(varchar(2)),
		field.String("minute").SchemaType(varchar(2)),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Optional(),
	}
}

func (PlanTimeZone) Edges() []ent.Edge {
	return nil
}
