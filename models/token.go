package models

import (
	"fmt"
	"time"
	"todo/database"
)

func SaveTokens(userID int, accessToken, refreshToken string, accessExpiry, refreshExpiry time.Time) error {
	query := `INSERT INTO user_tokens (user_id, access_token, refresh_token, access_token_expiry, refresh_token_expiry)
	VALUES ($1, $2, $3, $4, $5)
	ON CONFLICT (user_id) DO UPDATE
	SET access_token = $2, refresh_token = $3, access_token_expiry = $4, refresh_token_expiry = $5, updated_at = CURRENT_TIMESTAMP`

	_, err := database.DB.Exec(query, userID, accessToken, refreshToken, accessExpiry, refreshExpiry)
	if err != nil {
		return err
	}
	return nil
}

func UpdateToken(userID int64, accessToken, refreshToken string, accessExpiry, refreshExpiry time.Time) error {
	query := `
		INSERT INTO user_tokens (
			user_id, access_token, refresh_token, access_token_expiry, refresh_token_expiry
		) VALUES ($1, $2, $3, $4, $5)
		ON CONFLICT (user_id)
		DO UPDATE SET
			access_token = EXCLUDED.access_token,
			refresh_token = EXCLUDED.refresh_token,
			access_token_expiry = EXCLUDED.access_token_expiry,
			refresh_token_expiry = EXCLUDED.refresh_token_expiry,
			updated_at = CURRENT_TIMESTAMP`

	_, err := database.DB.Exec(query, userID, accessToken, refreshToken, accessExpiry, refreshExpiry)
	if err != nil {
		return fmt.Errorf("failed to update tokens for user_id %d: %w", userID, err)
	}

	return nil
}
