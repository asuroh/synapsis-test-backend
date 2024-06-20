package usecase

import (
	"errors"
	"synapsis-test-backend/helper"
	"synapsis-test-backend/model"
	"synapsis-test-backend/pkg/bcrypt"
	"synapsis-test-backend/pkg/logruslogger"
	"synapsis-test-backend/pkg/str"
	"synapsis-test-backend/server/request"
	"synapsis-test-backend/usecase/viewmodel"
	"time"

	uuid "github.com/satori/go.uuid"
)

// UserUC ...
type UserUC struct {
	*ContractUC
}

// BuildBody ...
func (uc UserUC) BuildBody(data *model.UserEntity, res *viewmodel.UserVM, isShowPassword bool) {

	res.ID = data.ID
	res.Name = data.Name.String
	res.Email = data.Email
	res.Password = str.ShowString(isShowPassword, data.Password)
	if data.ImagePath.String != "" {
		res.ImagePath = uc.EnvConfig["APP_IMAGE_URL"] + uc.EnvConfig["FILE_PATH"] + data.ImagePath.String
	}
	res.CreatedAt = data.CreatedAt
	res.UpdatedAt = data.UpdatedAt
	res.DeletedAt = data.DeletedAt.String
}

// Login ...
func (uc UserUC) Login(data request.UserLoginRequest) (res viewmodel.JwtVM, err error) {
	ctx := "UserUC.Login"

	if len(data.Password) < 8 {
		logruslogger.Log(logruslogger.WarnLevel, "", ctx, "password_length", uc.ReqID)
		return res, errors.New(helper.InvalidCredentials)
	}

	user, err := uc.FindByEmail(data.Email, true)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "find_by_email", uc.ReqID)
		return res, errors.New(helper.InvalidCredentials)
	}

	isMatch := bcrypt.CheckPasswordHash(data.Password, user.Password)
	if !isMatch {
		logruslogger.Log(logruslogger.WarnLevel, "", ctx, "invalid_password", uc.ReqID)
		return res, errors.New(helper.InvalidCredentials)
	}

	// Jwe the payload & Generate jwt token
	payload := map[string]interface{}{
		"id": user.ID,
	}
	jwtUc := JwtUC{ContractUC: uc.ContractUC}
	err = jwtUc.GenerateToken(payload, &res)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "jwt", uc.ReqID)
		return res, errors.New(helper.InternalServer)
	}

	return res, err
}

// Create ...
func (uc UserUC) Register(data *request.UserRequest) (res viewmodel.RegisterUserVM, err error) {
	ctx := "UserUC.Create"

	err = uc.CheckDetails(data, &res.User)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "check_details", uc.ReqID)
		return res, err
	}

	now := time.Now().UTC()
	res.User = viewmodel.UserVM{
		ID:        uuid.NewV4().String(),
		Name:      data.Name,
		Email:     data.Email,
		Password:  data.Password,
		CreatedAt: now.Format(time.RFC3339),
		UpdatedAt: now.Format(time.RFC3339),
	}
	m := model.NewUserModel(uc.DB)
	err = m.Store(res.User.ID, res.User, now)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "query", uc.ReqID)
		return res, err
	}

	res.User.Password = ""
	payload := map[string]interface{}{
		"id": res.User.ID,
	}
	jwtUc := JwtUC{ContractUC: uc.ContractUC}
	err = jwtUc.GenerateToken(payload, &res.Jwt)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "jwt", uc.ReqID)
		return res, errors.New(helper.InternalServer)
	}

	return res, err
}

// FindByID ...
func (uc UserUC) FindByID(id string, isShowPassword bool) (res viewmodel.UserVM, err error) {
	ctx := "UserUC.FindByID"

	m := model.NewUserModel(uc.DB)
	data, err := m.FindByID(id)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "query", uc.ReqID)
		return res, err
	}
	uc.BuildBody(&data, &res, isShowPassword)

	return res, err
}

// FindByEmail ...
func (uc UserUC) FindByEmail(Email string, isShowPassword bool) (res viewmodel.UserVM, err error) {
	ctx := "UserUC.FindByEmail"

	m := model.NewUserModel(uc.DB)
	data, err := m.FindByEmail(Email)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "query", uc.ReqID)
		return res, err
	}

	uc.BuildBody(&data, &res, isShowPassword)

	return res, err
}

// CheckDetails ...
func (uc UserUC) CheckDetails(data *request.UserRequest, oldData *viewmodel.UserVM) (err error) {
	ctx := "UserUC.CheckDetails"

	user, _ := uc.FindByEmail(data.Email, false)
	if user.ID != "" && user.ID != oldData.ID {
		logruslogger.Log(logruslogger.WarnLevel, data.Email, ctx, helper.DuplicateEmail, uc.ReqID)
		return errors.New(helper.DuplicateEmail)
	}

	if data.Password == "" && oldData.Password == "" {
		logruslogger.Log(logruslogger.WarnLevel, data.Email, ctx, helper.InvalidPassword, uc.ReqID)
		return errors.New(helper.InvalidPassword)
	}

	// Decrypt password input
	if data.Password == "" {
		data.Password = oldData.Password
	}

	// Encrypt password
	data.Password, err = bcrypt.HashPassword(data.Password)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "encrypt_password", uc.ReqID)
		return err
	}

	return err
}
