package cacheService

func Cache() error {
	if err := Permission(); err != nil {
		return err
	}
	return nil
}
