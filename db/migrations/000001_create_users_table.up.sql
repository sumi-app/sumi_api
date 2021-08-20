CREATE TABLE sumibloggers
(
    id             bigserial not null primary key,
    name           varchar,
    login          varchar   not null,
    type           varchar,
    description    varchar,
    subs_count     bigserial,
    avatar         varchar,
    social_network varchar
);