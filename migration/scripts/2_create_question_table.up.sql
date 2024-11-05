create table if not exists question (
    id uuid primary key,
    text varchar not null,
    form_id uuid not null,

    constraint fk_question_form foreign key (form_id) references form (id)
);
