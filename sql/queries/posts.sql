-- name: CreatePost :execrows
INSERT INTO posts (id, title, description, published_at, url, feed_id, created_at, updated_at)
VALUES (
  $1,
  $2,
  $3,
  $4,
  $5,
  $6,
  $7,
  $8
)
ON CONFLICT (url)
  DO NOTHING;

-- name: GetPostsForUser :many
SELECT
  posts.*,
  feeds.name AS feed_name,
  feeds.url AS feed_url
FROM
  posts
  INNER JOIN feed_follows ON posts.feed_id = feed_follows.feed_id
  INNER JOIN feeds ON posts.feed_id = feeds.id
WHERE
  feed_follows.user_id = $1
ORDER BY
  posts.published_at DESC NULLS LAST,
  posts.created_at DESC
LIMIT $2;

