create table if not exists form (
    id uuid primary key,
    name varchar not null,
    description varchar,
    testing_begin_date timestamp not null,
    testing_end_date timestamp not null,
    created_at timestamp not null,
    teacher_username varchar not null
);
