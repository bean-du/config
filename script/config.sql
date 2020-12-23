create table record
(
    id         bigserial                not null
        constraint record_pk
        primary key,
    key        varchar                  not null,
    value      text                     not null,
    version    varchar(100)             not null,
    pointer    varchar(30)              not null,
    created_at timestamp with time zone not null
);

alter table record
    owner to root;

create unique index record_id_uindex
    on record (id);

create unique index record_key_value_uindex
    on record (key, value);
