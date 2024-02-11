CREATE TABLE public.cities (
    id varchar not null,
    city varchar not null,
    created_at varchar not null,
    updated_at varchar null,
    deleted_at timestamp null,
    constraint cities_pk primary key (id)
)