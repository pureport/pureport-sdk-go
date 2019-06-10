package credentials

import (
	"errors"
	"testing"
	"time"
)

type providerStub struct {
	credentials Value
	expired     bool
	err         error
}

func (p *providerStub) Retrieve() (Value, error) {
	p.expired = false
	p.credentials.ProviderName = "providerStub"
	return p.credentials, p.err
}

func (p *providerStub) IsExpired() bool {
	return p.expired
}

func TestGetCredentials(t *testing.T) {
	c := NewCredentials(&providerStub{
		credentials: Value{
			APIKey:       "SomeKey",
			Secret:       "SomeSecret",
			SessionToken: "",
		}, expired: true,
	})

	credentials, err := c.Get()
	if err != nil {
		t.Errorf("Did not expect to get error from this test. %+v", err)
	}

	if e, a := "SomeKey", credentials.APIKey; e != a {
		t.Errorf("Invalid API Key: expected=%v, was=%v", e, a)
	}

	if e, a := "SomeSecret", credentials.Secret; e != a {
		t.Errorf("Invalid API Secret: expected=%v, was=%v", e, a)
	}

	if e, a := "", credentials.SessionToken; e != a {
		t.Errorf("Invalid Session Key: expected=%v, was=%v", e, a)
	}

	if e, a := "providerStub", credentials.ProviderName; e != a {
		t.Errorf("Invalid Provider Name: expected=%v, was=%v", e, a)
	}
}

func TestGetCredentialsWithError(t *testing.T) {

	c := NewCredentials(&providerStub{err: errors.New("some error")})

	_, err := c.Get()
	if err == nil {
		t.Error("Expected to receive error")
	}

	if err.Error() != "some error" {
		t.Error("Invalid Error values specified")
	}
}

func TestGetCredentialWithExpired(t *testing.T) {
	stub := providerStub{expired: false}
	c := NewCredentials(&stub)

	if !c.IsExpired() {
		t.Errorf("Expected credential to be expired")
	}

	c.forceRefresh = false
	if c.IsExpired() {
		t.Errorf("Expected credential to not be expired")
	}

	c.Expire()
	if !c.IsExpired() {
		t.Errorf("Expected credential to be expired")
	}

	c.forceRefresh = false
	stub.expired = true
	if !c.IsExpired() {
		t.Errorf("Expected credential to be expired")
	}
}

type mockProvider struct {
	Expiry
}

func (*mockProvider) Retrieve() (Value, error) {
	return Value{}, nil
}

func (p *mockProvider) IsExpired() bool {
	return p.Expiry.IsExpired()
}

func TestCredentialExpiry(t *testing.T) {
	m := mockProvider{}
	m.Expiry.SetExpiration(time.Now().Add(5*time.Second), 0)

	c := NewCredentials(&m)
	_, _ = c.Get()

	t.Logf("Expires at: %v: %v", m.Expiry.ExpiresAt(), time.Now())
	t.Logf("Expired?: %v", c.IsExpired())

	start := make(chan struct{})
	for i := 0; i < 10; i++ {

		go func(count int) {
			<-start
			for {
				if c.IsExpired() {
					t.Logf("Expired: %v: %v", count, time.Now())
					break
				}
			}
		}(i)
	}
	close(start)

	time.Sleep(6 * time.Second)
}
