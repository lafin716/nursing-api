package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("user_name").Optional().SchemaType(varchar(50)),
		field.String("user_email").SchemaType(varchar(100)),
		field.String("user_password").SchemaType(varchar(255)),
		field.String("user_status").Default("INACTIVE").SchemaType(varchar(20)),
		field.String("user_type").Default("EMAIL").SchemaType(varchar(20)),
		field.Int("fail_count").Default(0),
		field.Time("last_signed_in").Optional(),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Optional(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
