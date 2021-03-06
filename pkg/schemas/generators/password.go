package generators

import (
	"github.com/jenkins-x/jx-secret/pkg/schemas/secrets"
	"github.com/pkg/errors"
)

// Password generates a password value
func Password(args *Arguments) (string, error) {
	length := args.Property.MaxLength
	if length == 0 {
		length = args.Property.MinLength
		if length == 0 {
			length = 20
		}
	}
	value, err := secrets.DefaultGenerateSecret(length)
	if err != nil {
		return value, errors.WithStack(err)
	}
	return value, nil
}
