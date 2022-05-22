create table if not exists games
(
    id   UUID primary key,
    game jsonb not null
);