package instabot

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserProfile(t *testing.T) {
	pageAccessToken := "page_access_token"
	instagramUserID := "4576841382327552"

	type args struct {
		ctx             context.Context
		instagramUserID string
	}

	type fields struct {
		wantRequestQuery   url.Values
		returnResponse     string
		returnResponseCode int
	}

	type test struct {
		args    args
		fields  fields
		want    *GetUserProfileResponse
		wantErr error
	}

	tests := map[string]func(t *testing.T) test{
		"get user profile success": func(t *testing.T) test {
			args := args{
				ctx:             context.Background(),
				instagramUserID: instagramUserID,
			}

			q := url.Values{}
			q.Add("fields", "name,profile_pic,is_verified_user,follower_count,is_user_follow_business,is_business_follow_user")
			q.Add("access_token", pageAccessToken)

			fields := fields{
				wantRequestQuery: q,
				returnResponse: fmt.Sprintf(`{
					"name": "Shahin Mahmud",
					"profile_pic": "https://scontent.fdac14-1.fna.fbcdn.net/v/t51.2885-15/s.jpg",
					"id": "4576841382327552",
					"is_verified_user": true,
					"follower_count": 20,
					"is_user_follow_business": true,
					"is_business_follow_user": true
				  }`),
				returnResponseCode: 200,
			}

			want := &GetUserProfileResponse{
				Name:                    "Shahin Mahmud",
				ProfilePic:              "https://scontent.fdac14-1.fna.fbcdn.net/v/t51.2885-15/s.jpg",
				ID:                      "4576841382327552",
				IsVerifiedUser:          true,
				FollowerCount:           20,
				IsUserFollowingBusiness: true,
				IsBusinessFollowingUser: true,
			}

			return test{
				args:    args,
				fields:  fields,
				want:    want,
				wantErr: nil,
			}
		},
		"get user profile error": func(t *testing.T) test {
			args := args{
				ctx:             context.Background(),
				instagramUserID: instagramUserID,
			}

			q := url.Values{}
			q.Add("fields", "name,profile_pic,is_verified_user,follower_count,is_user_follow_business,is_business_follow_user")
			q.Add("access_token", pageAccessToken)

			fields := fields{
				wantRequestQuery: q,
				returnResponse: fmt.Sprintf(`{
					"error": {
						"message": "error",
						"type": "not found",
						"code": 100,
						"error_subcode": 23434,
						"fbtrace_id": "fbtrace_id"
					}
				}`),
				returnResponseCode: 400,
			}

			return test{
				args:   args,
				fields: fields,
				want:   nil,
				wantErr: &ErrorResponse{
					StatusCode: 400,
					APIError: APIError{
						Message:   "error",
						Type:      "not found",
						Code:      100,
						SubCode:   23434,
						FbTraceID: "fbtrace_id",
					},
				},
			}
		},
	}

	var currentTest string
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tc := tests[currentTest](t)

		assert.Equal(t, http.MethodGet, r.Method)

		assert.Equal(t, GetAPIEndpointUserProfile(instagramUserID), r.URL.Path)

		query := r.URL.Query()
		assert.Equal(t, tc.fields.wantRequestQuery, query)

		w.WriteHeader(tc.fields.returnResponseCode)
		w.Write([]byte(tc.fields.returnResponse))
	}))
	defer mockServer.Close()

	for name, fn := range tests {
		currentTest = name
		tt := fn(t)

		t.Run(name, func(t *testing.T) {
			client, err := New(pageAccessToken, WithEndpointBase(mockServer.URL))
			assert.NoError(t, err)

			res, err := client.GetUserProfile(tt.args.ctx, instagramUserID)
			if tt.wantErr != nil {
				assert.EqualError(t, tt.wantErr, err.Error())
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tt.want, res)
		})
	}
}
