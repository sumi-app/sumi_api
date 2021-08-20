CREATE TABLE sumireviews
(
    id bigserial not null primary key,
    blogger_id bigserial not null,
    press_tour_id bigserial not null,
    post_link varchar,
    mark bigserial
);