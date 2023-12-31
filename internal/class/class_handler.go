package class

import (
	"database/sql"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/tnnz20/godemy-be/internal/student"
	"github.com/tnnz20/godemy-be/internal/teacher"
	"github.com/tnnz20/godemy-be/util"
)

type Handler struct {
	ClassService   Service
	TeacherService teacher.Service
	StudentService student.Service
	Validate       *validator.Validate
}

func NewHandler(s Service, teacherSvc teacher.Service,
	studentSvc student.Service, validate *validator.Validate) *Handler {
	return &Handler{
		ClassService:   s,
		TeacherService: teacherSvc,
		StudentService: studentSvc,
		Validate:       validate,
	}
}

const (
	unauthorizedMsg = "Unauthorized not a "
	notFoundMsg     = " not found."
)

func (h *Handler) CreateClass(c *fiber.Ctx) error {
	var req CreateClassRequest

	id, role := util.JwtClaim(c)

	if role != "teacher" {
		return util.ErrorResponse(c, fiber.StatusUnauthorized, unauthorizedMsg+"Teacher")
	}

	teacherId := &teacher.GetTeacherIdByUserIdRequest{
		UserId: id,
	}

	if err := h.Validate.Struct(teacherId); err != nil {
		return util.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	teacher, err := h.TeacherService.GetTeacherIdByUserId(c.Context(), teacherId)
	if err != nil {
		if err == sql.ErrNoRows {
			return util.ErrorResponse(c, fiber.StatusNotFound, "Teacher"+notFoundMsg)
		}
	}

	if err := c.BodyParser(&req); err != nil {
		return util.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	req.TeacherId = teacher.ID

	if err := h.Validate.Struct(req); err != nil {
		return util.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	res, err := h.ClassService.CreateClass(c.Context(), &req)
	if err != nil {
		util.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return util.SuccessResponse(c, fiber.StatusCreated, "Class successfully created.", res)
}

func (h *Handler) GetAllClass(c *fiber.Ctx) error {
	res, err := h.ClassService.GetAllClass(c.Context())
	if err != nil {
		if err.Error() == "null" {
			return util.ErrorResponse(c, fiber.StatusNotFound, "Class still empty.")
		}
		return util.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return util.SuccessResponse(c, fiber.StatusOK, "Classes successfully retrieved.", res)
}

func (h *Handler) UpdateStudentClass(c *fiber.Ctx) error {
	var req UpdateStudentClassRequest

	id, role := util.JwtClaim(c)

	if role != "student" {
		return util.ErrorResponse(c, fiber.StatusUnauthorized, unauthorizedMsg+"Student")
	}

	studentReq := &student.GetStudentByUserIdRequest{
		UsersId: id,
	}

	student, err := h.StudentService.GetStudentByUserId(c.Context(), studentReq)
	if err != nil {
		if err == sql.ErrNoRows {
			return util.ErrorResponse(c, fiber.StatusNotFound, "Student"+notFoundMsg)
		}
		return util.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	req.StudentId = student.ID

	if err := c.BodyParser(&req); err != nil {
		return util.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	if err := h.Validate.Struct(req); err != nil {
		return util.ErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}

	if err := h.ClassService.UpdateStudentClass(c.Context(), &req); err != nil {
		util.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return util.SuccessResponse(c, fiber.StatusAccepted, "Successfully join a class.", nil)
}
