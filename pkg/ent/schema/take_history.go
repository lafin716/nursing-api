package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

type TakeHistory struct {
	ent.Schema
}

func (TakeHistory) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.UUID("user_id", uuid.UUID{}),
		field.UUID("timezone_id", uuid.UUID{}).Optional(),
		field.Time("take_date"),
		field.String("take_status").Optional().Default("NEVER").SchemaType(varchar(10)),
		field.String("memo").Optional().SchemaType(text()),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Optional(),
	}
}

func (TakeHistory) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("timezone", TimeZone.Type).
			Ref("take_history").
			Unique().
			Field("timezone_id"),
	}
}
