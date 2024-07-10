package prescription

import (
	"github.com/google/uuid"
	_ "github.com/joho/godotenv/autoload"
	"nursing_api/internal/config"
	"nursing_api/pkg/database"
	"nursing_api/pkg/jwt"
	"nursing_api/pkg/mono"
	"reflect"
	"testing"
)

func TestNewService(t *testing.T) {
	type args struct {
		repo      Repository
		jwtClient *jwt.JwtClient
	}
	tests := []struct {
		name string
		args args
		want UseCase
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewService(tt.args.repo, tt.args.jwtClient); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_prescriptionService_AddItem(t *testing.T) {
	type fields struct {
		repo      Repository
		mono      *mono.Client
		jwtClient *jwt.JwtClient
	}
	type args struct {
		req *AddItemRequest
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *AddItemResponse
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := prescriptionService{
				repo:      tt.fields.repo,
				mono:      tt.fields.mono,
				jwtClient: tt.fields.jwtClient,
			}
			if got := p.AddItem(tt.args.req); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddItem() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_prescriptionService_Delete(t *testing.T) {
	type fields struct {
		repo      Repository
		mono      *mono.Client
		jwtClient *jwt.JwtClient
	}
	type args struct {
		req *DeleteRequest
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *DeleteResponse
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := prescriptionService{
				repo:      tt.fields.repo,
				mono:      tt.fields.mono,
				jwtClient: tt.fields.jwtClient,
			}
			if got := p.Delete(tt.args.req); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_prescriptionService_DeleteItem(t *testing.T) {
	type fields struct {
		repo      Repository
		mono      *mono.Client
		jwtClient *jwt.JwtClient
	}
	type args struct {
		req *DeleteItemRequest
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *DeleteItemResponse
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := prescriptionService{
				repo:      tt.fields.repo,
				mono:      tt.fields.mono,
				jwtClient: tt.fields.jwtClient,
			}
			if got := p.DeleteItem(tt.args.req); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteItem() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_prescriptionService_GetById(t *testing.T) {
	type fields struct {
		repo      Repository
		mono      *mono.Client
		jwtClient *jwt.JwtClient
	}
	type args struct {
		req *GetByIdRequest
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *GetByIdResponse
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := prescriptionService{
				repo:      tt.fields.repo,
				mono:      tt.fields.mono,
				jwtClient: tt.fields.jwtClient,
			}
			if got := p.GetById(tt.args.req); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_prescriptionService_GetItemById(t *testing.T) {
	type fields struct {
		repo      Repository
		mono      *mono.Client
		jwtClient *jwt.JwtClient
	}
	type args struct {
		req *GetItemByIdRequest
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *GetItemByIdResponse
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := prescriptionService{
				repo:      tt.fields.repo,
				mono:      tt.fields.mono,
				jwtClient: tt.fields.jwtClient,
			}
			if got := p.GetItemById(tt.args.req); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetItemById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_prescriptionService_GetItemList(t *testing.T) {
	type fields struct {
		repo      Repository
		mono      *mono.Client
		jwtClient *jwt.JwtClient
	}
	type args struct {
		req *GetItemListRequest
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *GetItemListResponse
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := prescriptionService{
				repo:      tt.fields.repo,
				mono:      tt.fields.mono,
				jwtClient: tt.fields.jwtClient,
			}
			if got := p.GetItemList(tt.args.req); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetItemList() = %v, want %v", got, tt.want)
			}
		})
	}
}

type fields struct {
	repo      Repository
	mono      *mono.Client
	jwtClient *jwt.JwtClient
}

func defaultFields() fields {
	return fields{
		repo:      NewRepository(database.NewPostgresClient(config.NewDatabaseConfig())),
		mono:      mono.NewMono(),
		jwtClient: jwt.NewJwtClient(config.NewJwtConfig()),
	}
}

func Test_prescriptionService_GetList(t *testing.T) {
	userId, err := uuid.FromBytes([]byte(""))
	if err != nil {
		t.Errorf("uuid.FromBytes() = %v", err)
	}
	type args struct {
		req *GetListRequest
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *GetListResponse
	}{
		{name: "정상요청", fields: defaultFields(), args: args{req: &GetListRequest{UserId: userId}}, want: &GetListResponse{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := prescriptionService{
				repo:      tt.fields.repo,
				mono:      tt.fields.mono,
				jwtClient: tt.fields.jwtClient,
			}
			if got := p.GetList(tt.args.req); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_prescriptionService_Register(t *testing.T) {
	type fields struct {
		repo      Repository
		mono      *mono.Client
		jwtClient *jwt.JwtClient
	}
	type args struct {
		req *RegisterRequest
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *RegisterResponse
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := prescriptionService{
				repo:      tt.fields.repo,
				mono:      tt.fields.mono,
				jwtClient: tt.fields.jwtClient,
			}
			if got := p.Register(tt.args.req); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Register() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_prescriptionService_Update(t *testing.T) {
	type fields struct {
		repo      Repository
		mono      *mono.Client
		jwtClient *jwt.JwtClient
	}
	type args struct {
		req *UpdateRequest
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *UpdateResponse
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := prescriptionService{
				repo:      tt.fields.repo,
				mono:      tt.fields.mono,
				jwtClient: tt.fields.jwtClient,
			}
			if got := p.Update(tt.args.req); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Update() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_prescriptionService_UpdateItem(t *testing.T) {
	type fields struct {
		repo      Repository
		mono      *mono.Client
		jwtClient *jwt.JwtClient
	}
	type args struct {
		req *UpdateItemRequest
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *UpdateItemResponse
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := prescriptionService{
				repo:      tt.fields.repo,
				mono:      tt.fields.mono,
				jwtClient: tt.fields.jwtClient,
			}
			if got := p.UpdateItem(tt.args.req); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateItem() = %v, want %v", got, tt.want)
			}
		})
	}
}
