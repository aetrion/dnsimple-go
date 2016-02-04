package dnsimple

import (
	"fmt"
	"testing"
)

func testCredentials(t *testing.T, credentials Credentials, expectedName, expectedValue string) {
	headerName, headerValue := credentials.HttpHeader()
	if headerName != expectedName {
		t.Errorf("Header name: %v, want %v", headerName, expectedName)
	}

	if headerValue != expectedValue {
		t.Errorf("Header value: %v, want %v", headerValue, expectedValue)
	}
}

func TestDomainTokenCredentialsHttpHeader(t *testing.T) {
	domainToken := "domain-token"
	credentials := NewDomainTokenCredentials(domainToken)
	testCredentials(t, credentials, httpHeaderDomainToken, domainToken)
}

func TestHttpBasicCredentialsHttpHeader(t *testing.T) {
	email, password := "email", "password"
	credentials := NewHttpBasicCredentials(email, password)
	expectedHeaderValue := "Basic ZW1haWw6cGFzc3dvcmQ="
	testCredentials(t, credentials, httpHeaderAuthorization, expectedHeaderValue)
}

func TestOauthTokenCredentialsHttpHeader(t *testing.T) {
	oauthToken := "oauth-token"
	credentials := NewOauthTokenCredentials(oauthToken)
	expectedHeaderValue := fmt.Sprintf("Bearer %v", oauthToken)
	testCredentials(t, credentials, httpHeaderAuthorization, expectedHeaderValue)
}
