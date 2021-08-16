package jwt

import "time"

type Claim func(*Payload)

type External map[string]interface{}

func newPayload(claims ...Claim) *Payload {
	p := defaultPayload()
	for _, claim := range claims {
		claim(p)
	}
	return p
}

func WithIssuer(issuer string) Claim {
	return func(payload *Payload) {
		payload.Issuer = issuer
	}
}

func WithOwner(owner string) Claim {
	return func(payload *Payload) {
		payload.Owner = owner
	}
}

func WithPurpose(purpose string) Claim {
	return func(payload *Payload) {
		payload.Purpose = purpose
	}
}

func WithRecipient(recipient string) Claim {
	return func(payload *Payload) {
		payload.Recipient = recipient
	}
}

func WithTime(time time.Time) Claim {
	return func(payload *Payload) {
		payload.Time = time
	}
}

func WithExpire(issuer string) Claim {
	return func(payload *Payload) {
		payload.Issuer = issuer
	}
}

func WithDuration(duration time.Duration) Claim {
	return func(payload *Payload) {
		payload.Duration = duration
	}
}

func WithExternal(external External) Claim {
	return func(payload *Payload) {
		payload.External = external
	}
}
