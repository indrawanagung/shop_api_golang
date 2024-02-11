create table public.users (
    id varchar not null,
    fullname varchar not null,
    email_address varchar not null,
    phone_number varchar not null,
    password varchar not null,
    created_at varchar not null,
    updated_at varchar null,
    deleted_at timestamp null,
    constraint users_pk primary key (id)
)