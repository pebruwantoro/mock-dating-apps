package repository

import "context"

func (r *Repository) CreateUser(ctx context.Context, input User) (result User, err error) {
	var id string

	err = r.Db.QueryRowContext(ctx, `
		INSERT INTO users (uuid, email, username, password, is_premium, created_at)
		VALUES ($1, $2, $3, $4, $5, NOW())
		returning uuid;
	`,
		input.UUID,
		input.Email,
		input.Username,
		input.Password,
		input.IsPremium,
	).Scan(&id)
	if err != nil {
		return
	}

	result.UUID = id

	return
}

func (r *Repository) UpdateUser(ctx context.Context, input User) (result User, err error) {
	_, err = r.Db.ExecContext(ctx, `
		UPDATE users 
		SET is_premium = $1,
		updated_at = NOW()
		WHERE uuid = $2;
	`,
		input.IsPremium,
		input.UUID,
	)
	if err != nil {
		return
	}

	result = input

	return
}

func (r *Repository) GetUserByID(ctx context.Context, id string) (result User, err error) {
	err = r.Db.QueryRowContext(ctx, `
        SELECT uuid, email, username, password, is_premium, created_at
        FROM users
        WHERE uuid = $1;
    `, id).Scan(
		&result.UUID,
		&result.Email,
		&result.Username,
		&result.Password,
		&result.IsPremium,
		&result.CreatedAt,
	)
	if err != nil {
		return
	}

	return
}

func (r *Repository) GetUserByEmail(ctx context.Context, email string) (result User, err error) {
	err = r.Db.QueryRowContext(ctx, `
        SELECT uuid, email, username, password, is_premium, created_at
        FROM users
        WHERE email = $1;
    `, email).Scan(
		&result.UUID,
		&result.Email,
		&result.Username,
		&result.Password,
		&result.IsPremium,
		&result.CreatedAt,
	)
	if err != nil {
		return
	}

	return
}

func (r *Repository) GetUserByUsername(ctx context.Context, username string) (result User, err error) {
	err = r.Db.QueryRowContext(ctx, `
        SELECT uuid, email, username, password, is_premium, created_at
        FROM users
        WHERE username = $1;
    `, username).Scan(
		&result.UUID,
		&result.Email,
		&result.Username,
		&result.Password,
		&result.IsPremium,
		&result.CreatedAt,
	)
	if err != nil {
		return
	}

	return
}

func (r *Repository) CreateSwipe(ctx context.Context, input Swipe) (result Swipe, err error) {
	var id string

	err = r.Db.QueryRowContext(ctx, `
        INSERT INTO swipes (uuid, user_id, target_id, direction, created_at, updated_at)
        VALUES ($1, $2, $3, $4, NOW(), NOW())
        RETURNING uuid;
    `,
		input.UUID,
		input.UserID,
		input.TargetID,
		input.Direction,
	).Scan(&id)
	if err != nil {
		return
	}

	result.UUID = id

	return
}

func (r *Repository) UpdateSwipe(ctx context.Context, input Swipe) (result Swipe, err error) {
	_, err = r.Db.ExecContext(ctx, `
	    UPDATE swipes
	    SET direction = $1,
		updated_at = NOW()
	    WHERE uuid = $2;
	`,
		input.Direction,
		input.UUID,
	)
	if err != nil {
		return
	}

	result = input

	return result, nil
}

func (r *Repository) GetSwipeByUserIdAndTargetId(ctx context.Context, userId, targetId string) (result Swipe, err error) {
	err = r.Db.QueryRowContext(ctx, `
        SELECT uuid, user_id, target_id, direction, created_at, updated_at
        FROM swipes
        WHERE user_id = $1
		AND target_id = $2;
    `, userId, targetId).Scan(
		&result.UUID,
		&result.UserID,
		&result.TargetID,
		&result.Direction,
		&result.CreatedAt,
		&result.UpdatedAt,
	)
	if err != nil {
		return
	}

	return
}
