package usecase

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"kel15/models"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"

	"kel15/config"
	utils "kel15/utils"
)

func (usecase *Usecase) UserList(c *gin.Context) ([]models.User, error) {
	pagination := utils.GetPagination(c)

	users, err := usecase.repository.UserList(c, pagination)

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (usecase *Usecase) UserLogin(c *gin.Context) (*models.UserLoginResponse, error) {
	var payload models.UserLoginRequest
	validate := utils.NewValidator()

	err := json.NewDecoder(c.Request.Body).Decode(&payload)
	if err != nil {
		return nil, err
	}

	err = validate.Struct(&payload)

	if err != nil {
		return nil, errors.New(utils.MessageErrorByValidation(err))
	}

	user, err := usecase.repository.GetUserByEmail(c, payload.Email)

	if err != nil {
		return nil, err
	}

	err = utils.ComparePassword(payload.Password, user.Password)

	if err != nil {
		return nil, err
	}

	tokenString, err := utils.GenerateToken(user)

	if err != nil {
		return nil, err
	}

	return &models.UserLoginResponse{
		User: &models.User{
			ID:       user.ID,
			Email:    user.Email,
			Username: user.Username,
			Name:     user.Name,
		},
		Token: tokenString,
	}, nil
}

func (usecase *Usecase) UserRegister(c *gin.Context) (*models.UserRegisterResponse, error) {
	var payload models.UserRegisterRequest
	validate := utils.NewValidator()
	err := json.NewDecoder(c.Request.Body).Decode(&payload)
	if err != nil {
		return nil, err
	}
	err = validate.Struct(&payload)

	if err != nil {
		return nil, errors.New(utils.MessageErrorByValidation(err))
	}

	// err = utils.PasswordValidator(payload.Password)

	// if err != nil {
	// 	return nil, err
	// }

	user, err := usecase.repository.GetUserByEmailOrUsername(c, payload)

	if err != nil {
		return nil, err
	}

	if user != nil {
		return nil, errors.New("User already exists")
	}

	err = usecase.repository.CreateUser(c, payload)

	if err != nil {
		return nil, err
	}

	return &models.UserRegisterResponse{
		Message: "User created successfully",
	}, nil
}

func (usecase *Usecase) UserLoginByProvider(c *gin.Context) (string, error) {
	provider := c.Params.ByName("provider")
	if provider == "google" {
		oauthConf := config.NewGoogle()
		URL, err := url.Parse(oauthConf.Endpoint.AuthURL)

		if err != nil {
			return "", err
		}

		parameters := url.Values{}
		parameters.Add("client_id", oauthConf.ClientID)
		parameters.Add("scope", strings.Join(oauthConf.Scopes, " "))
		parameters.Add("redirect_uri", oauthConf.RedirectURL)
		parameters.Add("response_type", "code")
		// parameters.Add("state", oauthStateString)
		URL.RawQuery = parameters.Encode()
		url := URL.String()
		return url, nil
	}

	return "", errors.New("Provider not found")
}

func (usecase *Usecase) UserLoginByProviderCallback(c *gin.Context) (*models.UserLoginResponse, error) {
	oauthConf := config.NewGoogle()
	code := c.Query("code")

	token, err := oauthConf.Exchange(oauth2.NoContext, code)
	if err != nil {
		return nil, err
	}

	client := oauthConf.Client(oauth2.NoContext, token)
	response, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	contents, err := ioutil.ReadAll(response.Body)

	var userProvider models.UserGoogleResponse
	err = json.Unmarshal([]byte(string(contents)), &userProvider)

	if err != nil {
		return nil, err
	}

	userByEmail, err := usecase.repository.GetUserByEmail(c, userProvider.Email)

	if err != nil && userByEmail == nil {
		return nil, err
	}

	user, err := usecase.repository.CreateUserGoogle(c, userProvider)

	if err != nil {
		return nil, err
	}

	tokenString, err := utils.GenerateToken(user)

	return &models.UserLoginResponse{
		User: &models.User{
			ID:       user.ID,
			Email:    user.Email,
			Username: user.Username,
		},
		Token: tokenString,
	}, nil
}

func (usecase *Usecase) GetUserByUsername(c *gin.Context) (*models.User, error) {
	username := c.Request.FormValue("username")
	if username == "" {
		return nil, errors.New("Username is required")
	}
	user, err := usecase.repository.GetUserByUsername(c, username)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (usecase *Usecase) UserProfileUpdate(c *gin.Context, fileName string) (*models.User, error) {
	var payload models.UserUpdateProfileRequest

	payload.Username = c.Request.FormValue("username")
	payload.FileName = &fileName
	validate := utils.NewValidator()
	err := validate.Struct(&payload)

	if err != nil {
		return nil, errors.New(utils.MessageErrorByValidation(err))
	}

	user, err := usecase.repository.UserProfileUpdate(c, payload)

	if err != nil {
		return nil, err
	}

	return user, nil
}
