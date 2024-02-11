create table product_weights(
                                id varchar not null,
                                product_id varchar not null,
                                weight_unit_id varchar not null,
                                value int not null,
                                created_at varchar not null,
                                updated_at varchar null,
                                deleted_at timestamp null,
                                constraint productWeights_pk primary key (id),
                                constraint productWeights_product_fk foreign key (product_id) references products(id),
                                constraint productWeights_weightUnits_fk foreign key (weight_unit_id) references weight_units(id)
)