create table whitelists
(
    id         serial                    not null
        constraint whitelists_pk
            primary key,
    subnet     cidr                      not null,
    created_at timestamptz default now() not null
);
