CREATE TABLE users(
    id UUID not null,
    email VARCHAR(255) not null,
    password VARCHAR(255) not null,
    name VARCHAR(255) not null,
    date Date,
    address VARCHAR(255),
    gender VARCHAR(255),
    profile_img VARCHAR(255),
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
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
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    PRIMARY KEY (id),
    FOREIGN KEY (users_id) REFERENCES users(id)
);

CREATE TABLE course_enrollment(
    id UUID not null,
    users_id UUID not null,
    course_id UUID,
    progress INT DEFAULT 0,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    PRIMARY KEY (id),
    FOREIGN KEY (users_id) REFERENCES users(id),
    FOREIGN KEY (course_id) REFERENCES courses(id)
);

CREATE TABLE assessment(
    id UUID not null,
    users_id UUID not null,
    course_id UUID not null,
    assessment_value int,
    assessment_code varchar(10),
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    FOREIGN KEY (users_id) REFERENCES users(id),
    FOREIGN KEY (course_id) REFERENCES courses(id)
);