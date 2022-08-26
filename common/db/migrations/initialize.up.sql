-- Initialize the database -- this is all for testing since
-- we will be sending the data to Omeka.
CREATE TABLE IF NOT EXISTS conthreads_collections (
    id text PRIMARY KEY,
    title text,
    description text,
    medium text[],
    url text,
    date text,
    country text,
    type text,
    api jsonb
);

CREATE TABLE IF NOT EXISTS conthreads_items (
    id text PRIMARY KEY,
    title text,
    description text,
    medium text[],
    url text,
    date text,
    country text,
    type text,
    timestamp bigint,
    api jsonb
);

CREATE INDEX ON conthreads_items (date);

CREATE INDEX ON conthreads_items (timestamp);

CREATE INDEX ON conthreads_items (api)
WHERE
    api IS NULL;