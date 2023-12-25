CREATE TYPE users_role as enum('student', 'teacher');
CREATE TYPE users_gender as enum('Laki-Laki', 'Perempuan');

CREATE EXTENSION "uuid-ossp";

CREATE TABLE users(
    id UUID DEFAULT uuid_generate_v4(),
    email VARCHAR(255) not null,
    password VARCHAR(255) not null,
    role users_role DEFAULT 'student',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    PRIMARY KEY (id)
);

CREATE TABLE profile(
    id UUID DEFAULT uuid_generate_v4(),
    name VARCHAR(255) not null,
    gender users_gender not null,
    users_id UUID not null,
    profile_img VARCHAR(255),
    PRIMARY KEY (id),
    FOREIGN KEY (users_id) REFERENCES users(id)
);

CREATE TABLE teacher(
    id UUID DEFAULT uuid_generate_v4(),
    users_id UUID not null,
    PRIMARY KEY (id),
    FOREIGN KEY (users_id) REFERENCES users(id)
);

CREATE TABLE class(
    id UUID DEFAULT uuid_generate_v4(),
    teacher_id UUID not null,
    class_name VARCHAR(50),
    PRIMARY KEY (id),
    FOREIGN KEY (teacher_id) REFERENCES teacher(id)
);

CREATE TABLE student(
    id UUID DEFAULT uuid_generate_v4(),
    users_id UUID not null,
    class_id UUID,
    threshold INT DEFAULT 0,
    PRIMARY KEY (id),
    FOREIGN KEY (users_id) REFERENCES users(id),
    FOREIGN KEY (class_id) REFERENCES class(id)
);

CREATE TABLE assessment(
    id UUID DEFAULT uuid_generate_v4(),
    student_id UUID not null,
    assessment_value int,
    code_assessment varchar(20),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (student_id ) REFERENCES student(id)
);