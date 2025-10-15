package responses

import "github.com/caiiomp/vehicle-resale-auth/src/core/domain/entity"

type Login struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}

func LoginFromDomain(auth entity.Auth) Login {
	return Login{
		AccessToken: auth.AccessToken,
		ExpiresIn:   auth.ExpiresIn,
	}
}
