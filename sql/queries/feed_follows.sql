-- name: CreateFeedFollows :many
WITH inserted_feed_follows AS (
    INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
    VALUES ($1, $2, $3, $4, $5)
    RETURNING *
)
SELECT
    inserted_feed_follows.*,
    feeds.name AS feed_name,
    users.name AS user_name
FROM inserted_feed_follows
INNER JOIN feeds ON inserted_feed_follows.feed_id = feeds.id
INNER JOIN users ON inserted_feed_follows.user_id = users.id;

-- name: GetFeedFollowsForUser :many
SELECT
    feed_follows.*,
    feeds.name AS feed_name,
    users.name AS user_name
FROM feed_follows
INNER JOIN feeds ON feed_follows.feed_id = feeds.id
INNER JOIN users ON feed_follows.user_id = users.id
WHERE feed_follows.user_id = $1;

-- name: DeleteFeedFollow :exec
DELETE FROM feed_follows
USING feeds
WHERE feed_follows.feed_id = feeds.id
    AND feed_follows.user_id = $1
    AND feeds.url = $2;

-- name: MarkFeedFetched :exec
UPDATE feed_follows
SET 
    last_fetched_at = $1, 
    updated_at = $2
WHERE feed_id = $3;

-- name: GetNextFeedToFetch :one
SELECT 
    f.id AS feed_id,
    f.url AS feed_url
FROM 
    feeds f
INNER JOIN 
    feed_follows ff 
ON 
    f.id = ff.feed_id
ORDER BY 
    ff.last_fetched_at NULLS FIRST, 
    ff.updated_at ASC
LIMIT 1;
