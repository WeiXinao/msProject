package jwtx

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"time"
)

func TestJwtToken_ParseToken(t *testing.T) {
	tokenString := "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzYwODM1NDIsIlVzZXJJZCI6NX0.SaTvwTU2zBAZJjU0PK8yuJBIsT73-zBltbuPT01zzwx_Sg_2o101_2hrTqDH97WSFIDEfPyaeyYPMk81Od0-FA"
	jwter := NewJwtToken("95osj3fUD7fo0mlYdDbncXz4VD2igvf0", "", time.Hour, time.Hour)
	c := UserClaims{}
	err := jwter.ParseToken(tokenString, &c)
	assert.NoError(t, err)
	log.Print(c)
	return
}
