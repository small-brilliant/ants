package jwt
import "github.com/golang-jwt/jwt/v4"

const (
	JWT_KEY_ISS= "iss"
	JWT_KEY_EXP="exp"
	JWT_KEY_IAT="iat"
)

func CreateToken(){
	token := jwt.NewWithClaims(jwt.SigningMethodPS256,jwt.MapClaims{
		"iss": 
	})
}