-- Return the database to empty.
-- This is destructive, so please be careful.
BEGIN;
DROP TABLE IF EXISTS conthreads_collections CASCADE;
DROP TABLE IF EXISTS conthreads_items CASCADE;
END;