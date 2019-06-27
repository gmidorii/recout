package app

import (
	"context"
	"testing"

	"github.com/gmidorii/recout/backend/app/mock"
	"github.com/gmidorii/recout/backend/form"
	"github.com/gmidorii/recout/backend/infra/entity"
	"github.com/gmidorii/recout/backend/infra/pixela"
	"github.com/gmidorii/recout/backend/infra/repository"
	"github.com/golang/mock/gomock"
)

func Test_user_Save(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type fields struct {
		ctn          Container
		repoUser     repository.User
		pixelaClient pixela.Client
	}
	type args struct {
		ctx  context.Context
		form form.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "success",
			fields: fields{
				ctn: Container{
					Generator: func() RandomGenerator {
						m := mock.NewMockRandomGenerator(ctrl)
						m.EXPECT().
							Do(gomock.Eq(pixelaTokenLen)).
							Return("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
						m.EXPECT().
							Do(gomock.Eq(pixelaGraphIDLen)).
							Return("aaaaaaaa")
						return m
					}(),
				},
				repoUser: func() repository.User {
					m := mock.NewMockUser(ctrl)
					m.EXPECT().
						Put(gomock.Any(), entity.User{
							AccountID:   "recgmidorii",
							PixelaGraph: "aaaaaaaa",
							PixelaToken: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
						})
					return m
				}(),
				pixelaClient: func() pixela.Client {
					m := mock.NewMockClient(ctrl)
					m.EXPECT().
						CreateUser(pixela.User{
							Token:               "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
							UserName:            "recgmidorii",
							AgreeTermsOfService: pixela.Yes,
							NotMinor:            pixela.Yes,
						}).
						Return(nil)
					m.EXPECT().
						CreateGraph("aaaaaaaa", "recgmidorii", "recgmidorii", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa").
						Return(nil)
					return m
				}(),
			},
			args: args{
				ctx: context.Background(),
				form: form.User{
					AccountID: "gmidorii",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &user{
				ctn:          tt.fields.ctn,
				repoUser:     tt.fields.repoUser,
				pixelaClient: tt.fields.pixelaClient,
			}
			if err := p.Save(tt.args.ctx, tt.args.form); (err != nil) != tt.wantErr {
				t.Errorf("user.Save() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
