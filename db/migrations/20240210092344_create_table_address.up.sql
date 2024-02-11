create table public.address (
    id varchar not null,
    name varchar not null,
    city_id varchar not null,
    postal_code int,
    created_at varchar not null,
    updated_at varchar null,
    deleted_at timestamp null,
    constraint address_pk primary key (id),
    constraint address_city_fk
    foreign key (city_id) references public.cities(id)
)