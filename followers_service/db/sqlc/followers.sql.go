// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: followers.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const followUser = `-- name: FollowUser :one
INSERT INTO followers (leader_unique_id, follower_unique_id)
VALUES ($1, $2)
ON CONFLICT (leader_unique_id, follower_unique_id) DO NOTHING
RETURNING id, leader_unique_id, follower_unique_id, created_at, updated_at
`

type FollowUserParams struct {
	LeaderUniqueID   uuid.UUID `json:"leader_unique_id"`
	FollowerUniqueID uuid.UUID `json:"follower_unique_id"`
}

func (q *Queries) FollowUser(ctx context.Context, arg FollowUserParams) (Follower, error) {
	row := q.db.QueryRowContext(ctx, followUser, arg.LeaderUniqueID, arg.FollowerUniqueID)
	var i Follower
	err := row.Scan(
		&i.ID,
		&i.LeaderUniqueID,
		&i.FollowerUniqueID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getFollowers = `-- name: GetFollowers :many
SELECT "follower_unique_id" FROM followers
WHERE leader_unique_id = $1
ORDER BY created_at DESC
LIMIT $2
OFFSET $3
`

type GetFollowersParams struct {
	LeaderUniqueID uuid.UUID `json:"leader_unique_id"`
	Limit          int32     `json:"limit"`
	Offset         int32     `json:"offset"`
}

func (q *Queries) GetFollowers(ctx context.Context, arg GetFollowersParams) ([]uuid.UUID, error) {
	rows, err := q.db.QueryContext(ctx, getFollowers, arg.LeaderUniqueID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []uuid.UUID{}
	for rows.Next() {
		var follower_unique_id uuid.UUID
		if err := rows.Scan(&follower_unique_id); err != nil {
			return nil, err
		}
		items = append(items, follower_unique_id)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getFollowersCount = `-- name: GetFollowersCount :one
SELECT COUNT(*) FROM followers
WHERE leader_unique_id = $1
`

func (q *Queries) GetFollowersCount(ctx context.Context, leaderUniqueID uuid.UUID) (int64, error) {
	row := q.db.QueryRowContext(ctx, getFollowersCount, leaderUniqueID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const getFollowing = `-- name: GetFollowing :many
SELECT "leader_unique_id" FROM followers
WHERE follower_unique_id = $1
ORDER BY created_at DESC
LIMIT $2
OFFSET $3
`

type GetFollowingParams struct {
	FollowerUniqueID uuid.UUID `json:"follower_unique_id"`
	Limit            int32     `json:"limit"`
	Offset           int32     `json:"offset"`
}

func (q *Queries) GetFollowing(ctx context.Context, arg GetFollowingParams) ([]uuid.UUID, error) {
	rows, err := q.db.QueryContext(ctx, getFollowing, arg.FollowerUniqueID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []uuid.UUID{}
	for rows.Next() {
		var leader_unique_id uuid.UUID
		if err := rows.Scan(&leader_unique_id); err != nil {
			return nil, err
		}
		items = append(items, leader_unique_id)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getFollowingCount = `-- name: GetFollowingCount :one
SELECT COUNT(*) FROM followers
WHERE follower_unique_id = $1
`

func (q *Queries) GetFollowingCount(ctx context.Context, followerUniqueID uuid.UUID) (int64, error) {
	row := q.db.QueryRowContext(ctx, getFollowingCount, followerUniqueID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const unfollowUser = `-- name: UnfollowUser :one
DELETE FROM followers
WHERE leader_unique_id = $1 AND follower_unique_id = $2
RETURNING id, leader_unique_id, follower_unique_id, created_at, updated_at
`

type UnfollowUserParams struct {
	LeaderUniqueID   uuid.UUID `json:"leader_unique_id"`
	FollowerUniqueID uuid.UUID `json:"follower_unique_id"`
}

func (q *Queries) UnfollowUser(ctx context.Context, arg UnfollowUserParams) (Follower, error) {
	row := q.db.QueryRowContext(ctx, unfollowUser, arg.LeaderUniqueID, arg.FollowerUniqueID)
	var i Follower
	err := row.Scan(
		&i.ID,
		&i.LeaderUniqueID,
		&i.FollowerUniqueID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
