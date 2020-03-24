CREATE KEYSPACE IF NOT EXISTS example WITH replication =
{'class':'SimpleStrategy','replication_factor':'1'};

CREATE TABLE IF NOT EXISTS example.users (
    id UUID,
    name text,
    gender text,
    age int,
    created_at timestamp,
    updated_at timestamp,
    PRIMARY KEY(id)
);