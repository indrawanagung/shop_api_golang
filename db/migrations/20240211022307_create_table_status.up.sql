create table public.status (
    id varchar not null,
    category_status_id varchar not null,
    name varchar,
    created_at varchar not null,
    updated_at varchar null,
    deleted_at timestamp null,
    constraint status_pk primary key (id),
    constraint status_category_status_fk foreign key (category_status_id) references public.category_status(id)
)