// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// MedicinesColumns holds the columns for the "medicines" table.
	MedicinesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "medicine_name", Type: field.TypeString, SchemaType: map[string]string{"postgres": "varchar(50)"}},
		{Name: "item_seq", Type: field.TypeString, SchemaType: map[string]string{"postgres": "varchar(20)"}},
		{Name: "company_name", Type: field.TypeString, Nullable: true, SchemaType: map[string]string{"postgres": "varchar(30)"}},
		{Name: "description", Type: field.TypeString, Nullable: true, SchemaType: map[string]string{"postgres": "varchar(255)"}},
		{Name: "usage", Type: field.TypeString, Nullable: true, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "effect", Type: field.TypeString, Nullable: true, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "side_effect", Type: field.TypeString, Nullable: true, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "caution", Type: field.TypeString, Nullable: true, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "warning", Type: field.TypeString, Nullable: true, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "interaction", Type: field.TypeString, Nullable: true, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "keep_method", Type: field.TypeString, Nullable: true, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "appearance", Type: field.TypeString, Nullable: true, SchemaType: map[string]string{"postgres": "varchar(20)"}},
		{Name: "color_class1", Type: field.TypeString, Nullable: true, SchemaType: map[string]string{"postgres": "varchar(20)"}},
		{Name: "color_class2", Type: field.TypeString, Nullable: true, SchemaType: map[string]string{"postgres": "varchar(20)"}},
		{Name: "pill_image", Type: field.TypeString, Nullable: true, SchemaType: map[string]string{"postgres": "varchar(255)"}},
		{Name: "class_name", Type: field.TypeString, Nullable: true, SchemaType: map[string]string{"postgres": "varchar(255)"}},
		{Name: "otc_name", Type: field.TypeString, Nullable: true, SchemaType: map[string]string{"postgres": "varchar(255)"}},
		{Name: "form_code_name", Type: field.TypeString, Nullable: true, SchemaType: map[string]string{"postgres": "varchar(255)"}},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
	}
	// MedicinesTable holds the schema information for the "medicines" table.
	MedicinesTable = &schema.Table{
		Name:       "medicines",
		Columns:    MedicinesColumns,
		PrimaryKey: []*schema.Column{MedicinesColumns[0]},
	}
	// PlanTimeZonesColumns holds the columns for the "plan_time_zones" table.
	PlanTimeZonesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "user_id", Type: field.TypeUUID},
		{Name: "timezone_name", Type: field.TypeString, Nullable: true, SchemaType: map[string]string{"postgres": "varchar(50)"}},
		{Name: "is_default", Type: field.TypeBool, Default: false},
		{Name: "midday", Type: field.TypeString, SchemaType: map[string]string{"postgres": "varchar(2)"}},
		{Name: "hour", Type: field.TypeString, SchemaType: map[string]string{"postgres": "varchar(2)"}},
		{Name: "minute", Type: field.TypeString, SchemaType: map[string]string{"postgres": "varchar(2)"}},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
	}
	// PlanTimeZonesTable holds the schema information for the "plan_time_zones" table.
	PlanTimeZonesTable = &schema.Table{
		Name:       "plan_time_zones",
		Columns:    PlanTimeZonesColumns,
		PrimaryKey: []*schema.Column{PlanTimeZonesColumns[0]},
	}
	// PlanTimeZoneLinksColumns holds the columns for the "plan_time_zone_links" table.
	PlanTimeZoneLinksColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "prescription_id", Type: field.TypeUUID},
		{Name: "timezone_id", Type: field.TypeUUID},
		{Name: "timezone_name", Type: field.TypeString, Nullable: true, SchemaType: map[string]string{"postgres": "varchar(50)"}},
		{Name: "use_alert", Type: field.TypeBool, Default: false},
		{Name: "midday", Type: field.TypeString, SchemaType: map[string]string{"postgres": "varchar(2)"}},
		{Name: "hour", Type: field.TypeString, SchemaType: map[string]string{"postgres": "varchar(2)"}},
		{Name: "minute", Type: field.TypeString, SchemaType: map[string]string{"postgres": "varchar(2)"}},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
	}
	// PlanTimeZoneLinksTable holds the schema information for the "plan_time_zone_links" table.
	PlanTimeZoneLinksTable = &schema.Table{
		Name:       "plan_time_zone_links",
		Columns:    PlanTimeZoneLinksColumns,
		PrimaryKey: []*schema.Column{PlanTimeZoneLinksColumns[0]},
	}
	// PrescriptionsColumns holds the columns for the "prescriptions" table.
	PrescriptionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "user_id", Type: field.TypeUUID},
		{Name: "prescription_name", Type: field.TypeString, Nullable: true, SchemaType: map[string]string{"postgres": "varchar(30)"}},
		{Name: "hospital_name", Type: field.TypeString, Nullable: true, SchemaType: map[string]string{"postgres": "varchar(50)"}},
		{Name: "take_days", Type: field.TypeInt, Default: 0},
		{Name: "started_at", Type: field.TypeTime, Nullable: true, SchemaType: map[string]string{"postgres": "date"}},
		{Name: "finished_at", Type: field.TypeTime, Nullable: true, SchemaType: map[string]string{"postgres": "date"}},
		{Name: "memo", Type: field.TypeString, Nullable: true, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
	}
	// PrescriptionsTable holds the schema information for the "prescriptions" table.
	PrescriptionsTable = &schema.Table{
		Name:       "prescriptions",
		Columns:    PrescriptionsColumns,
		PrimaryKey: []*schema.Column{PrescriptionsColumns[0]},
	}
	// PrescriptionItemsColumns holds the columns for the "prescription_items" table.
	PrescriptionItemsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "timezone_link_id", Type: field.TypeUUID},
		{Name: "medicine_id", Type: field.TypeUUID},
		{Name: "medicine_name", Type: field.TypeString, SchemaType: map[string]string{"postgres": "varchar(50)"}},
		{Name: "take_amount", Type: field.TypeFloat64, Default: 0},
		{Name: "medicine_unit", Type: field.TypeString, Nullable: true, Default: "개", SchemaType: map[string]string{"postgres": "varchar(3)"}},
		{Name: "memo", Type: field.TypeString, Nullable: true, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
	}
	// PrescriptionItemsTable holds the schema information for the "prescription_items" table.
	PrescriptionItemsTable = &schema.Table{
		Name:       "prescription_items",
		Columns:    PrescriptionItemsColumns,
		PrimaryKey: []*schema.Column{PrescriptionItemsColumns[0]},
	}
	// TakeHistoriesColumns holds the columns for the "take_histories" table.
	TakeHistoriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "user_id", Type: field.TypeUUID},
		{Name: "prescription_id", Type: field.TypeUUID},
		{Name: "take_date", Type: field.TypeTime, Nullable: true},
		{Name: "take_status", Type: field.TypeString, Nullable: true, Default: "NEVER", SchemaType: map[string]string{"postgres": "varchar(10)"}},
		{Name: "memo", Type: field.TypeString, Nullable: true, SchemaType: map[string]string{"postgres": "text"}},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
	}
	// TakeHistoriesTable holds the schema information for the "take_histories" table.
	TakeHistoriesTable = &schema.Table{
		Name:       "take_histories",
		Columns:    TakeHistoriesColumns,
		PrimaryKey: []*schema.Column{TakeHistoriesColumns[0]},
	}
	// TakeHistoryItemsColumns holds the columns for the "take_history_items" table.
	TakeHistoryItemsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "user_id", Type: field.TypeUUID},
		{Name: "take_history_id", Type: field.TypeUUID},
		{Name: "prescription_item_id", Type: field.TypeUUID},
		{Name: "take_status", Type: field.TypeString, Default: "N", SchemaType: map[string]string{"postgres": "varchar(1)"}},
		{Name: "take_amount", Type: field.TypeFloat64, Default: 0},
		{Name: "take_time_zone", Type: field.TypeString, Nullable: true, SchemaType: map[string]string{"postgres": "varchar(10)"}},
		{Name: "take_moment", Type: field.TypeString, Nullable: true, SchemaType: map[string]string{"postgres": "varchar(10)"}},
		{Name: "take_etc", Type: field.TypeString, Nullable: true, SchemaType: map[string]string{"postgres": "varchar(50)"}},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
	}
	// TakeHistoryItemsTable holds the schema information for the "take_history_items" table.
	TakeHistoryItemsTable = &schema.Table{
		Name:       "take_history_items",
		Columns:    TakeHistoryItemsColumns,
		PrimaryKey: []*schema.Column{TakeHistoryItemsColumns[0]},
	}
	// TokensColumns holds the columns for the "tokens" table.
	TokensColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "user_id", Type: field.TypeUUID},
		{Name: "access_token", Type: field.TypeString, SchemaType: map[string]string{"postgres": "varchar(500)"}},
		{Name: "refresh_token", Type: field.TypeString, SchemaType: map[string]string{"postgres": "varchar(200)"}},
		{Name: "access_token_expires", Type: field.TypeTime},
		{Name: "refresh_token_expires", Type: field.TypeTime},
		{Name: "auto_login", Type: field.TypeBool},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
	}
	// TokensTable holds the schema information for the "tokens" table.
	TokensTable = &schema.Table{
		Name:       "tokens",
		Columns:    TokensColumns,
		PrimaryKey: []*schema.Column{TokensColumns[0]},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "user_name", Type: field.TypeString, Nullable: true, SchemaType: map[string]string{"postgres": "varchar(50)"}},
		{Name: "user_email", Type: field.TypeString, SchemaType: map[string]string{"postgres": "varchar(100)"}},
		{Name: "user_password", Type: field.TypeString, SchemaType: map[string]string{"postgres": "varchar(255)"}},
		{Name: "user_status", Type: field.TypeString, Default: "INACTIVE", SchemaType: map[string]string{"postgres": "varchar(20)"}},
		{Name: "user_type", Type: field.TypeString, Default: "EMAIL", SchemaType: map[string]string{"postgres": "varchar(20)"}},
		{Name: "fail_count", Type: field.TypeInt, Default: 0},
		{Name: "last_signed_in", Type: field.TypeTime, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		MedicinesTable,
		PlanTimeZonesTable,
		PlanTimeZoneLinksTable,
		PrescriptionsTable,
		PrescriptionItemsTable,
		TakeHistoriesTable,
		TakeHistoryItemsTable,
		TokensTable,
		UsersTable,
	}
)

func init() {
}
