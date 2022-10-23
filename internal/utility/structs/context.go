package structs

type Context struct {
	User *ContextUser
}

type ContextUser struct {
	UserGuid string
	Username string
}
