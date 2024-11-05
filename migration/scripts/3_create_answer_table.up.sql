create table if not exists answer (
    id uuid primary key,
    text varchar not null,
    question_id uuid not null,
    is_correct bool not null,

    constraint fk_answer_question foreign key (question_id) references question (id)
);
