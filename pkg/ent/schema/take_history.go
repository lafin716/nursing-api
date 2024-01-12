package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

type TakeHistory struct {
	ent.Schema
}

func (TakeHistory) Field() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.UUID("user_id", uuid.UUID{}),
		field.UUID("prescription_id", uuid.UUID{}),
		field.Time("take_date").Optional(),
		field.String("take_status").Optional().Default("NEVER").SchemaType(varchar(10)),
		field.String("memo").Optional().SchemaType(text()),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Optional(),
	}
}
