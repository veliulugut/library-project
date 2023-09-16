package user

import (
	"github.com/gin-gonic/gin"
	"library/service/user"
	"net/http"
	"strconv"
)

var _ Handler = (*User)(nil)

type User struct {
	userService user.Service
}

func NewUser(us user.Service) *User {
	return &User{
		userService: us,
	}
}

// CreateUser
// @Summary create user
// @Schemes
// @Description create user
// @Tags user
// @Accept json
// @Produce json
// @Param user body user.CreateUserModel true "Create a new user"
// @Success 201
// @Router /user [post]
func (u *User) CreateUser(c *gin.Context) {
	var (
		userdb user.CreateUserModel
		err    error
	)

	if err = c.ShouldBindJSON(&userdb); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = u.userService.CreateUser(c.Request.Context(), &userdb); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusCreated)
}

// GetUserByID godoc
// @Summary get user by id
// @Schemes
// @Description
// @Tags user
// @Accept  json
// @Produce  json
// @Param id   path int true "id"
// @Success 200 {object} user.GetUserModel "ok"
// @Router /user/{id} [get]
func (u *User) GetUserByID(c *gin.Context) {
	idParam := c.Param("id")

	var (
		id  int64
		err error
	)

	if id, err = strconv.ParseInt(idParam, 10, 64); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var um *user.GetUserModel

	if um, err = u.userService.GetUserByID(c.Request.Context(), int(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, um)

}

// DeleteUser godoc
// @Summary delete user by id
// @Schemes
// @Description
// @Tags user
// @Param id   path int true "id"
// @Success 204
// @Router /user/{id} [delete]
func (u *User) DeleteUser(c *gin.Context) {
	idParam := c.Param("id")

	var (
		id  uint64
		err error
	)

	if id, err = strconv.ParseUint(idParam, 10, 64); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = u.userService.DeleteUser(c.Request.Context(), int(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// UpdateUser godoc
// @Summary update user by id
// @Schemes
// @Description
// @Tags user
// @Accept  json
// @Produce  json
// @Param user body user.UpdateUserModel  true "Update a  user"
// @Param id   path int true "id"
// @Success 204
// @Router /user/{id} [post]
func (u *User) UpdateUser(c *gin.Context) {
	idParam := c.Param("id")

	var (
		id  uint64
		err error
	)

	if id, err = strconv.ParseUint(idParam, 10, 64); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var um *user.UpdateUserModel

	if err = c.ShouldBindJSON(&um); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = u.userService.UpdateUser(c.Request.Context(), int(id), um); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
