package svc

import (
	"context"
	"github.com/kutty-kumar/charminder/pkg"
	"github.com/kutty-kumar/ho_oh/pikachu_v1"
	"pikachu/pkg/domain"
	"reflect"
	"testing"
)

func TestUserService_CreateUserAttribute(t *testing.T) {
	type fields struct {
		BaseSvc              pkg.BaseSvc
		IdentityService      IdentityService
		UserAttributeService UserAttributeService
	}
	type args struct {
		ctx context.Context
		req *pikachu_v1.CreateUserAttributeRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pikachu_v1.CreateUserAttributeResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserService{
				BaseSvc:              tt.fields.BaseSvc,
				IdentityService:      tt.fields.IdentityService,
				UserAttributeService: tt.fields.UserAttributeService,
			}
			got, err := u.CreateUserAttribute(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserService.CreateUserAttribute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserService.CreateUserAttribute() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserService_UpdateUserAttribute(t *testing.T) {
	type fields struct {
		BaseSvc              pkg.BaseSvc
		IdentityService      IdentityService
		UserAttributeService UserAttributeService
	}
	type args struct {
		ctx context.Context
		req *pikachu_v1.UpdateUserAttributeRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pikachu_v1.UpdateUserAttributeResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserService{
				BaseSvc:              tt.fields.BaseSvc,
				IdentityService:      tt.fields.IdentityService,
				UserAttributeService: tt.fields.UserAttributeService,
			}
			got, err := u.UpdateUserAttribute(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserService.UpdateUserAttribute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserService.UpdateUserAttribute() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserService_GetUserAttributesByKey(t *testing.T) {
	type fields struct {
		BaseSvc              pkg.BaseSvc
		IdentityService      IdentityService
		UserAttributeService UserAttributeService
	}
	type args struct {
		ctx context.Context
		req *pikachu_v1.GetUserAttributeByKeyRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pikachu_v1.GetUserAttributeByKeyResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserService{
				BaseSvc:              tt.fields.BaseSvc,
				IdentityService:      tt.fields.IdentityService,
				UserAttributeService: tt.fields.UserAttributeService,
			}
			got, err := u.GetUserAttributesByKey(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserService.GetUserAttributesByKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserService.GetUserAttributesByKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserService_GetUserAttributes(t *testing.T) {
	type fields struct {
		BaseSvc              pkg.BaseSvc
		IdentityService      IdentityService
		UserAttributeService UserAttributeService
	}
	type args struct {
		ctx context.Context
		req *pikachu_v1.GetUserAttributesRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pikachu_v1.GetUserAttributesResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserService{
				BaseSvc:              tt.fields.BaseSvc,
				IdentityService:      tt.fields.IdentityService,
				UserAttributeService: tt.fields.UserAttributeService,
			}
			got, err := u.GetUserAttributes(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserService.GetUserAttributes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserService.GetUserAttributes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewUserService(t *testing.T) {
	type args struct {
		base             pkg.BaseSvc
		identitySvc      IdentityService
		userAttributeSvc UserAttributeService
	}
	tests := []struct {
		name string
		args args
		want UserService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserService(tt.args.base, tt.args.identitySvc, tt.args.userAttributeSvc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userOperationResponseMapper(t *testing.T) {
	type args struct {
		dto *pikachu_v1.UserDto
	}
	tests := []struct {
		name string
		args args
		want *pikachu_v1.UserOperationResponse
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := userOperationResponseMapper(tt.args.dto); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userOperationResponseMapper() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserService_handleError(t *testing.T) {
	type fields struct {
		BaseSvc              pkg.BaseSvc
		IdentityService      IdentityService
		UserAttributeService UserAttributeService
	}
	type args struct {
		err            error
		base           pkg.Base
		responseMapper func(dto *pikachu_v1.UserDto) *pikachu_v1.UserOperationResponse
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pikachu_v1.UserOperationResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserService{
				BaseSvc:              tt.fields.BaseSvc,
				IdentityService:      tt.fields.IdentityService,
				UserAttributeService: tt.fields.UserAttributeService,
			}
			got, err := u.handleError(tt.args.err, tt.args.base, tt.args.responseMapper)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserService.handleError() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserService.handleError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserService_getUser(t *testing.T) {
	type fields struct {
		BaseSvc              pkg.BaseSvc
		IdentityService      IdentityService
		UserAttributeService UserAttributeService
	}
	type args struct {
		dto pikachu_v1.UserDto
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *domain.User
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserService{
				BaseSvc:              tt.fields.BaseSvc,
				IdentityService:      tt.fields.IdentityService,
				UserAttributeService: tt.fields.UserAttributeService,
			}
			if got := u.getUser(tt.args.dto); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserService.getUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserService_CreateUser(t *testing.T) {
	type fields struct {
		BaseSvc              pkg.BaseSvc
		IdentityService      IdentityService
		UserAttributeService UserAttributeService
	}
	type args struct {
		ctx context.Context
		req *pikachu_v1.CreateUserRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pikachu_v1.UserOperationResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserService{
				BaseSvc:              tt.fields.BaseSvc,
				IdentityService:      tt.fields.IdentityService,
				UserAttributeService: tt.fields.UserAttributeService,
			}
			got, err := u.CreateUser(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserService.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserService.CreateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserService_UpdateUser(t *testing.T) {
	type fields struct {
		BaseSvc              pkg.BaseSvc
		IdentityService      IdentityService
		UserAttributeService UserAttributeService
	}
	type args struct {
		ctx context.Context
		req *pikachu_v1.UpdateUserRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pikachu_v1.UserOperationResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserService{
				BaseSvc:              tt.fields.BaseSvc,
				IdentityService:      tt.fields.IdentityService,
				UserAttributeService: tt.fields.UserAttributeService,
			}
			got, err := u.UpdateUser(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserService.UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserService.UpdateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserService_GetUserByExternalId(t *testing.T) {
	type fields struct {
		BaseSvc              pkg.BaseSvc
		IdentityService      IdentityService
		UserAttributeService UserAttributeService
	}
	type args struct {
		ctx context.Context
		req *pikachu_v1.GetUserByExternalIdRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pikachu_v1.UserOperationResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserService{
				BaseSvc:              tt.fields.BaseSvc,
				IdentityService:      tt.fields.IdentityService,
				UserAttributeService: tt.fields.UserAttributeService,
			}
			got, err := u.GetUserByExternalId(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserService.GetUserByExternalId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserService.GetUserByExternalId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserService_MultiGetUsersByExternalId(t *testing.T) {
	type fields struct {
		BaseSvc              pkg.BaseSvc
		IdentityService      IdentityService
		UserAttributeService UserAttributeService
	}
	type args struct {
		ctx context.Context
		req *pikachu_v1.MultiGetUsersByExternalIdRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pikachu_v1.MultiGetUsersResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserService{
				BaseSvc:              tt.fields.BaseSvc,
				IdentityService:      tt.fields.IdentityService,
				UserAttributeService: tt.fields.UserAttributeService,
			}
			got, err := u.MultiGetUsersByExternalId(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserService.MultiGetUsersByExternalId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserService.MultiGetUsersByExternalId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserService_CreateUserIdentity(t *testing.T) {
	type fields struct {
		BaseSvc              pkg.BaseSvc
		IdentityService      IdentityService
		UserAttributeService UserAttributeService
	}
	type args struct {
		ctx context.Context
		req *pikachu_v1.CreateUserIdentityRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pikachu_v1.CreateUserIdentityResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserService{
				BaseSvc:              tt.fields.BaseSvc,
				IdentityService:      tt.fields.IdentityService,
				UserAttributeService: tt.fields.UserAttributeService,
			}
			got, err := u.CreateUserIdentity(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserService.CreateUserIdentity() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserService.CreateUserIdentity() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserService_GetUserIdentities(t *testing.T) {
	type fields struct {
		BaseSvc              pkg.BaseSvc
		IdentityService      IdentityService
		UserAttributeService UserAttributeService
	}
	type args struct {
		ctx context.Context
		req *pikachu_v1.GetUserIdentitiesRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pikachu_v1.GetUserIdentitiesResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserService{
				BaseSvc:              tt.fields.BaseSvc,
				IdentityService:      tt.fields.IdentityService,
				UserAttributeService: tt.fields.UserAttributeService,
			}
			got, err := u.GetUserIdentities(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserService.GetUserIdentities() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserService.GetUserIdentities() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserService_UpdateUserIdentity(t *testing.T) {
	type fields struct {
		BaseSvc              pkg.BaseSvc
		IdentityService      IdentityService
		UserAttributeService UserAttributeService
	}
	type args struct {
		ctx context.Context
		req *pikachu_v1.UpdateUserIdentityRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pikachu_v1.UpdateUserIdentityResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserService{
				BaseSvc:              tt.fields.BaseSvc,
				IdentityService:      tt.fields.IdentityService,
				UserAttributeService: tt.fields.UserAttributeService,
			}
			got, err := u.UpdateUserIdentity(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserService.UpdateUserIdentity() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserService.UpdateUserIdentity() = %v, want %v", got, tt.want)
			}
		})
	}
}
