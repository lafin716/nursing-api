package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

// User holds the schema definition for the User entity.
type Medicine struct {
	ent.Schema
}

// Fields of the User.
func (Medicine) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("medicine_name").SchemaType(varchar(50)),
		field.String("item_seq").SchemaType(varchar(20)),
		field.String("company_name").Optional().SchemaType(varchar(30)),
		field.String("description").Optional().SchemaType(varchar(255)),
		field.String("usage").Optional().SchemaType(text()),
		field.String("effect").Optional().SchemaType(text()),
		field.String("side_effect").Optional().SchemaType(text()),
		field.String("caution").Optional().SchemaType(text()),
		field.String("warning").Optional().SchemaType(text()),
		field.String("interaction").Optional().SchemaType(text()),
		field.String("keep_method").Optional().SchemaType(text()),
		field.String("appearance").Optional().SchemaType(varchar(20)),
		field.String("color_class1").Optional().SchemaType(varchar(20)),
		field.String("color_class2").Optional().SchemaType(varchar(20)),
		field.String("pill_image").Optional().SchemaType(varchar(255)),
		field.String("class_name").Optional().SchemaType(varchar(255)),
		field.String("otc_name").Optional().SchemaType(varchar(255)),
		field.String("form_code_name").Optional().SchemaType(varchar(255)),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Optional(),
	}
}

// Edges of the User.
func (Medicine) Edges() []ent.Edge {
	return nil
}
