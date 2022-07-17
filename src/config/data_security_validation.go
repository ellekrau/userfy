package config

import "errors"

func validateDataSecurityKey() error {
	if len(DataSecurity.SecurityDataKey) != 16 {
		return errors.New("data security key must have 16 chars")
	}

	return nil
}
