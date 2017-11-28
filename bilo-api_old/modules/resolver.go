package modules

import (
    "github.com/ricardorsierra/bilo-api/modules/user"
)

type Resolver struct{}

func (r *Resolver) CreateUser(args *struct { User  *user.UserInput }) *user.UserResolver {    
    User := user.CreateUser(args.User)
    return &user.UserResolver{User}
}

func (r *Resolver) User(args *struct{ ID int }) *user.UserResolver {
    User := user.FindUserByID(args.ID)
    return &user.UserResolver{User}
}
