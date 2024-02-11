create table payment_methods (
    id varchar not null,
    payment_type_id varchar not null,
    account_number varchar not null,
    expiry_date varchar not null,
    status_id varchar not null,
    created_at varchar not null,
    updated_at varchar null,
    constraint paymentMethods_pk primary key (id),
    constraint paymenMethods_paymentTypes_fk foreign key (payment_type_id) references payment_types(id),
    constraint paymenMethods_status_fk foreign key (status_id) references status(id)

)