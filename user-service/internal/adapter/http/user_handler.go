package http

import (
	"context"
	"net/http"
	"time"
	"user-service/internal/domain/model"

	"user-service/internal/utils"

	"github.com/gin-gonic/gin"
)

func (h *UserHandler) SignUp(ctx *gin.Context) {
	// Create a new context with a timeout
	timeoutCtx, cancel := context.WithTimeout(ctx.Request.Context(), 5*time.Second)
	defer cancel() // Ensure the cancel function is called to release resources

	var user model.CreateRequest

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	//validate the value for the name,email,address
	if user.Name == "" || user.Email == "" || user.Address == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "name,email,address are need to fill"})
		return
	}
	//check password and repassword match
	if user.Password != user.Repassword {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Passwords do not match"})
		return
	}
	if len(user.Name) < 3 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Name must be atleast 3 character"})
	}

	CreateResponse, err := h.hand.User().CreateUser(timeoutCtx, &user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, CreateResponse)

}

func (h *UserHandler) GetAllUsers(ctx *gin.Context) {
	// Create a new context with a timeout
	timeoutCtx, cancel := context.WithTimeout(ctx.Request.Context(), 5*time.Second)
	defer cancel() // Ensure the cancel function is called to release resources

	// Call the service to get all users
	Users, err := h.hand.User().GetAllUsers(timeoutCtx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	var userresponse []model.CreateResponse
	for _, user := range Users {
		userresponse = append(userresponse, model.CreateResponse{ID: user.ID, Name: user.Name, Email: user.Email, Address: user.Address})
	}
	ctx.JSON(http.StatusOK, userresponse)
}

func (h *UserHandler) Login(ctx *gin.Context) {
	// Create a new context with a timeout
	timeoutCtx, cancel := context.WithTimeout(ctx.Request.Context(), 10*time.Second)
	defer cancel()

	// Bind the request body JSON to the user object
	var user model.LoginRequest
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	// Validate that email and password are not empty
	if user.Email == "" || user.Password == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Email and password are required"})
		return
	}

	// Fetch the user by email from the database
	signupValue, err := h.hand.User().GetUserByEmail(timeoutCtx, user.Email)
	if err != nil || signupValue == nil {
		// Handle case where the user is not found or other errors
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Compare the stored password hash with the provided password
	err = utils.ComparePassword(signupValue.Password, user.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// // Convert string to int (parsing user ID)
	// parsedInt, err := strconv.Atoi(signupValue.ID)
	// if err != nil {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	// // Cast to uint
	// u := uint(parsedInt)

	// Generate JWT token
	token, err := utils.GenerateJWT(signupValue.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// Respond with the JWT token
	ctx.JSON(http.StatusOK, gin.H{"token": token})
}

func (h *UserHandler) GetProfile(ctx *gin.Context) {
	// Create a new context with a timeout
	timeoutCtx, cancel := context.WithTimeout(ctx.Request.Context(), 10*time.Second)
	defer cancel()
	// Get the JWT token from the Authorization header
	id, exist := ctx.Get("userID")
	if !exist {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User Id not found"})
		return
	}
	// Perform type assertion to ensure it's the correct type
	userIDString, ok := id.(string)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID type"})
		return
	}
	// userIDInt := int(userIDUint) // Convert uint to int if needed

	user, err := h.hand.User().GetUserById(timeoutCtx, userIDString)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user profile"})
		return
	}
	ctx.JSON(http.StatusOK, user)

}

func (h *UserHandler) UpdateProfile(ctx *gin.Context) {
	//set timout
	timeoutCtx, cancel := context.WithTimeout(ctx.Request.Context(), 5*time.Second)
	defer cancel()
	id, exist := ctx.Get("userID")
	if !exist {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User Id not found"})
		return
	}
	// Perform type assertion to ensure it's the correct type
	userIDString, ok := id.(string)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID type"})
		return
	}
	// Get the updated user data from the request body
	user, err := h.hand.User().GetUserById(timeoutCtx, userIDString)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User is not exist"})
		return
	}
	var updateuser model.UpdateUser
	if err := ctx.ShouldBindJSON(&updateuser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Json"})
	}
	if updateuser.Email == "" && updateuser.Name == "" && updateuser.Address == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}
	//update the data
	user.Name = updateuser.Name
	user.Email = updateuser.Email
	user.Address = updateuser.Address

	updateduser, err := h.hand.User().UpdateUser(timeoutCtx, user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user profile"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Profile updated successfully", "user": updateduser})

}
