# Godemy Backend

## About Godemy

Godemy is learning platform for my skripsi, and this repository will hold anything about the backend.

## ERD

## API Endpoint

> Note

- ** optionals query
- \* required query

<details>
<summary>Users</summary>

| Method | Endpoint           | Param / query    | JWT Token | Description                  |
| ------ | ------------------ | ---------------- | --------- | ---------------------------- |
| POST   | /api/users/sign-up | ** ?role=teacher | NO        | Registration user            |
| GET    | /api/users/profile | -                | YES       | Get user profile from userId |

</details>

<details>
<summary>Auth</summary>

| Method | Endpoint          | Param / query | JWT Token | Description |
| ------ | ----------------- | ------------- | --------- | ----------- |
| GET    | /api/auth/sign-in | -             | NO        | Login user  |

</details>

<details>
<summary>Teacher</summary>

| Method | Endpoint                            | Param / query | JWT Token | Description                         |
| ------ | ----------------------------------- | ------------- | --------- | ----------------------------------- |
| GET    | /api/teachers/teacher               | -             | YES       | Get Teacher from userId             |
| GET    | /api/teachers/teacher/classes       | -             | YES       | Get all class from teacherId        |
| GET    | /api/teachers/teacher/classes/class | -             | YES       | Get list student who has same class |

</details>

<details>
<summary>Class</summary>

| Method | Endpoint                   | Param / query | JWT Token | Description                  |
| ------ | -------------------------- | ------------- | --------- | ---------------------------- |
| GET    | /api/classes               | -             | NO        | Get all class                |
| POST   | /api/classes               | -             | YES       | Create class using teacherId |
| PATCH  | /api/classes/class/student | -             | YES       | Update Class Student         |

</details>

<details>
<summary>Student</summary>

| Method | Endpoint                         | Param / query | JWT Token | Description                             |
| ------ | -------------------------------- | ------------- | --------- | --------------------------------------- |
| GET    | /api/students/student            | -             | YES       | Get student from userId                 |
| PATCH  | /api/students/student/threshold  | -             | YES       | Increment threshold student             |
| POST   | /api/students/student/assessment | -             | YES       | Assign value from assessment to student |

</details>
