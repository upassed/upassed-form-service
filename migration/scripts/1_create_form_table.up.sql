create table if not exists form (
    id uuid primary key,
    name varchar not null,
    teacher_username varchar not null
);
