-- name: CreateFeedFollow :one
WITH inserted_feeds_follow AS (
INSERT INTO feed_follows (id, user_id, feed_id, created_at, updated_at)
VALUES (
  $1,
  $2,
  $3,
  $4,
  $5
)
RETURNING
  *
)
SELECT
  inserted_feeds_follow.*,
  users.name AS user_name,
  feeds.name AS feed_name
FROM inserted_feeds_follow
  INNER JOIN users ON inserted_feeds_follow.user_id = users.id
  INNER JOIN feeds ON inserted_feeds_follow.feed_id = feeds.id;

-- name: GetFeedFollowsForUser :many
SELECT
  users.name AS user_name,
  feeds.name AS feed_name,
  feeds.url AS feed_url
FROM
  feed_follows
  INNER JOIN users ON feed_follows.user_id = users.id
  INNER JOIN feeds ON feed_follows.feed_id = feeds.id
WHERE
  users.id = $1;

-- name: DeleteFeedFollow :execrows
DELETE FROM feed_follows
WHERE
  user_id = $1
  AND feed_id = $2;
