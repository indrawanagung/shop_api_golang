create table public.user_address (
    user_id varchar not null,
    address_id varchar not null,
    is_default boolean,
    constraint user_address_pk primary key (user_id, address_id),
    constraint users_user_address_fk foreign key (user_id) references public.users(id)
);