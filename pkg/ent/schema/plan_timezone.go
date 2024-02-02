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
		field.String("use_alerm").Default("N").SchemaType(varchar(1)),
		field.Time("scheduled_at").Optional(),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Optional(),
	}
}

func (PlanTimeZone) Edges() []ent.Edge {
	return nil
}
