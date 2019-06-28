package app

import (
	"context"
	"testing"
	"time"

	"github.com/gmidorii/recout/backend/app/mock"
	"github.com/gmidorii/recout/backend/form"
	"github.com/gmidorii/recout/backend/infra/entity"
	"github.com/gmidorii/recout/backend/infra/pixela"
	"github.com/gmidorii/recout/backend/infra/repository"
	"github.com/golang/mock/gomock"
)

func Test_recout_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type fields struct {
		ctn           Container
		repoRecout    repository.Recout
		repoUser      repository.User
		repoContinues repository.Continues
		pixelaClient  pixela.Client
	}
	type args struct {
		ctx  context.Context
		form form.Recout
	}

	now := time.Date(2019, time.May, 1, 10, 0, 0, 0, time.UTC)
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantUid string
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				ctn: Container{
					Now:      now,
					Location: time.UTC,
				},
				repoRecout: func() repository.Recout {
					m := mock.NewMockRecout(ctrl)
					m.EXPECT().
						Put(gomock.Any(), entity.Recout{
							AccountID: "recgmidorii",
							Message:   "Hi!!",
							CreatedAt: now.In(time.UTC).Unix(),
						}).
						Return("mock_uid", nil)
					return m
				}(),
				repoUser: func() repository.User {
					m := mock.NewMockUser(ctrl)
					m.EXPECT().
						Get(gomock.Any(), "recgmidorii").
						Return("key", entity.User{
							AccountID:   "recgmidorii",
							AccessToken: "xxxxxxxx",
							Name:        "midori",
							PixelaGraph: "dev-recout",
							PixelaToken: "pixela_token",
						}, nil)
					return m
				}(),
				repoContinues: func() repository.Continues {
					m := mock.NewMockContinues(ctrl)
					m.EXPECT().
						Get(gomock.Any(), "recgmidorii").
						Return("mock_key", entity.Continues{
							AccountID: "recgmidorii",
							LastDate:  "20190430",
							Count:     4,
						}, nil)
					m.EXPECT().
						PutKey(gomock.Any(), "mock_key", entity.Continues{
							AccountID: "recgmidorii",
							LastDate:  "20190501",
							Count:     5,
						}).
						Return(nil)
					return m
				}(),
				pixelaClient: func() pixela.Client {
					m := mock.NewMockClient(ctrl)
					m.EXPECT().
						Increment("recgmidorii", "pixela_token", "dev-recout").
						Return(nil)
					return m
				}(),
			},
			args: args{
				form: form.Recout{
					AccountID: "gmidorii",
					Message:   "Hi!!",
				},
			},
			wantUid: "mock_uid",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &recout{
				ctn:           tt.fields.ctn,
				repoRecout:    tt.fields.repoRecout,
				repoUser:      tt.fields.repoUser,
				repoContinues: tt.fields.repoContinues,
				pixelaClient:  tt.fields.pixelaClient,
			}
			gotUid, err := r.Create(tt.args.ctx, tt.args.form)
			if (err != nil) != tt.wantErr {
				t.Errorf("recout.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotUid != tt.wantUid {
				t.Errorf("recout.Create() = %v, want %v", gotUid, tt.wantUid)
			}
		})
	}
}
