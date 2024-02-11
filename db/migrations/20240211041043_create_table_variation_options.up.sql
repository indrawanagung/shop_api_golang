create table variation_options (
    id varchar not null ,
    variation_id varchar not null,
    option_name varchar not null,
    description varchar null,
    created_at varchar not null,
    updated_at varchar null,
    deleted_at timestamp null,
    constraint variationOptions_pk primary key (id),
    constraint variationOptions_variation_fk foreign key (variation_id) references variations(id)
)