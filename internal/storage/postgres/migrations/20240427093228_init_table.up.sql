CREATE TYPE user_role as enum('student', 'teacher');
CREATE TYPE user_gender as enum('male', 'female');

CREATE EXTENSION "uuid-ossp";

CREATE TABLE users(
    id UUID DEFAULT uuid_generate_v4(),
    email VARCHAR(255) not null,
    password VARCHAR(255) not null,
    role user_role DEFAULT 'student',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    PRIMARY KEY (id)
);

CREATE TABLE profile(
    id UUID DEFAULT uuid_generate_v4(),
    name VARCHAR(255) not null,
    date Date,
    address VARCHAR(255),
    gender user_gender,
    profile_img VARCHAR(255),
    users_id UUID not null,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    FOREIGN KEY (users_id) REFERENCES users(id)
);

CREATE TABLE teacher(
    id UUID DEFAULT uuid_generate_v4(),
    users_id UUID not null,
    PRIMARY KEY (id),
    FOREIGN KEY (users_id) REFERENCES users(id)
);

CREATE TABLE courses(
    id UUID DEFAULT uuid_generate_v4(),
    course_name VARCHAR(255),
    course_code VARCHAR(10) UNIQUE,
    teacher_id UUID not null,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    FOREIGN KEY (teacher_id) REFERENCES teacher(id)
);

CREATE TABLE student(
    id UUID DEFAULT uuid_generate_v4(),
    users_id UUID not null,
    course_id UUID,
    threshold INT DEFAULT 0,
    PRIMARY KEY (id),
    FOREIGN KEY (users_id) REFERENCES users(id),
    FOREIGN KEY (course_id) REFERENCES courses(id)
);

CREATE TABLE assessment(
    id UUID DEFAULT uuid_generate_v4(),
    student_id UUID not null,
    course_id UUID not null,
    assessment_value int,
    assessment_code varchar(5),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (student_id) REFERENCES student(id),
    FOREIGN KEY (course_id) REFERENCES courses(id)
);