create table public.stores (
    id varchar not null ,
    user_id varchar not null,
    name varchar not null,
    points int null,
    created_at varchar not null,
    updated_at varchar null,
    deleted_at timestamp null,
    constraint stores_pk primary key (id),
    constraint stores_users_fk foreign key (user_id) references public.users(id)
)