package user

import (
	"fmt"
	"net/mail"
	"strings"
)

type Email struct {
	local  string
	domain string
	full   string
}

func NewEmail(raw string) (Email, error) {
	s := strings.TrimSpace(raw)
	if s == "" {
		return Email{}, fmt.Errorf("email is required")
	}

	addr, err := mail.ParseAddress(s)
	if err != nil {
		return Email{}, fmt.Errorf("invalid email: %w", err)
	}

	parts := strings.Split(addr.Address, "@")
	if len(parts) != 2 {
		return Email{}, fmt.Errorf("invalid email")
	}
	local, domain := parts[0], strings.ToLower(parts[1])

	return Email{
		local:  local,
		domain: domain,
		full:   local + "@" + domain,
	}, nil
}

func (e Email) String() string { return e.full }
func (e Email) Domain() string { return e.domain }
func (e Email) Local() string  { return e.local }
