package schema

import (
	"entgo.io/ent/dialect"
	"fmt"
)

func varchar(size int) map[string]string {
	return map[string]string{
		dialect.Postgres: fmt.Sprintf("varchar(%d)", size),
	}
}

func text() map[string]string {
	return map[string]string{
		dialect.Postgres: "text",
	}
}

func date() map[string]string {
	return map[string]string{
		dialect.Postgres: "date",
	}
}
