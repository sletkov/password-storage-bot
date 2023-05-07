CREATE TABLE services(
    id bigserial not null primary key,
    user_name varchar not null,
    service_name varchar not null,
    login varchar not null,
    password varchar not null
)