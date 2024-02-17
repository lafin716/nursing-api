package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

type TimeZoneLink struct {
	ent.Schema
}

func (TimeZoneLink) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.UUID("prescription_id", uuid.UUID{}),
		field.UUID("timezone_id", uuid.UUID{}),
		field.String("timezone_name").Optional().SchemaType(varchar(50)),
		field.Bool("use_alert").Default(false),
		field.String("midday").SchemaType(varchar(2)),
		field.String("hour").SchemaType(varchar(2)),
		field.String("minute").SchemaType(varchar(2)),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Optional(),
	}
}

func (TimeZoneLink) Edges() []ent.Edge {
	return nil
}
