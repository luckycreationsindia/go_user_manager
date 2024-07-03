package api

import (
	"fmt"
	"time"
	"user_manager/types"
)

func (s *APIServer) GetSession(sessionToken string, adminCheck bool) *types.ResponseMessage {
	cookieResult, err := s.sessionStorage.GetSession(sessionToken)
	if err != nil {
		fmt.Printf("%+v\n\n", err)
		return &types.ResponseMessage{Status: -1, Message: "Invalid Session", Redirect: "/login"}
	}
	expireTime := cookieResult.SessionExpires
	if time.Now().After(expireTime) {
		_ = s.sessionStorage.DeleteExpiredSession()
		return &types.ResponseMessage{Status: -1, Message: "Invalid Session", Redirect: "/login"}
	}
	return &types.ResponseMessage{Status: 1, Data: cookieResult}
}
