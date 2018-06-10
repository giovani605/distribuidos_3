package cliente

const endUrl string = "https://www.myservice.com"

type Cliente struct{
	Username string
	Password string
}

func Autenticacao(user, pass string) *Cliente{
	return &Cliente{
		Username: user,
		Password: pass,
	}
}

