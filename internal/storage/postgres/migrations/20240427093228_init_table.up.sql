CREATE TABLE users(
    id UUID not null,
    email VARCHAR(255) not null,
    password VARCHAR(255) not null,
    name VARCHAR(255) not null,
    date BIGINT,
    address VARCHAR(255),
    gender VARCHAR(255),
    profile_img VARCHAR(255),
    created_at BIGINT,
    updated_at BIGINT,
    deleted_at BIGINT,
    PRIMARY KEY (id)
);

CREATE TABLE roles(
    users_id UUID not null,
    role VARCHAR(255) not null,
    PRIMARY KEY (users_id)
);

CREATE TABLE courses(
    id UUID not null,
    users_id UUID not null,
    course_name VARCHAR(255),
    course_code VARCHAR(10) UNIQUE,
    created_at BIGINT,
    updated_at BIGINT,
    PRIMARY KEY (id),
    FOREIGN KEY (users_id) REFERENCES users(id)
);

CREATE TABLE course_enrollment(
    id UUID not null,
    users_id UUID not null,
    courses_id UUID,
    progress INT DEFAULT 0,
    created_at BIGINT,
    updated_at BIGINT,
    PRIMARY KEY (id),
    FOREIGN KEY (users_id) REFERENCES users(id),
    FOREIGN KEY (courses_id) REFERENCES courses(id)
);

CREATE TABLE users_assessment(
    id UUID not null,
    users_id UUID not null,
    assessment_code varchar(10) not null,
    random_array_id integer[],
    status VARCHAR(255),
    created_at BIGINT,
    updated_at BIGINT,
    FOREIGN KEY (users_id) REFERENCES users(id)
);

CREATE TABLE users_assessment_result(
    id UUID not null,
    users_id UUID not null,
    courses_id UUID not null,
    assessment_value FLOAT,
    assessment_code varchar(10) not null,
    status INT,
    created_at BIGINT,
    updated_at BIGINT,
    FOREIGN KEY (users_id) REFERENCES users(id),
    FOREIGN KEY (courses_id) REFERENCES courses(id)
);

