create table blacklists
(
    id         serial                    not null
        constraint blacklists_pk
            primary key,
    subnet     cidr                      not null,
    created_at timestamptz default now() not null
);