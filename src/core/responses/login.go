package responses

import "github.com/caiiomp/vehicle-resale-auth/src/core/domain/entity"

type LoginResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}

func LoginResponseFromDomain(auth entity.Auth) LoginResponse {
	return LoginResponse{
		AccessToken: auth.AccessToken,
		ExpiresIn:   auth.ExpiresIn,
	}
}
