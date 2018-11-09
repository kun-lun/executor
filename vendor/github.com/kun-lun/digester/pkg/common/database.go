package common

import (
	"fmt"
)

type Database struct {
	Driver         string `name:"Driver" question:"What's the driver?" allow:"{mysql}"`
	Version        string `name:"Version" question:"What's the version?" allow:"{5.6, 5.7}"`
	Storage        string `name:"Storage in GB" question:"What's the storage in GB?"`
	OriginHost     string `name:"Origin Host" question:"What's the host of the origin database?"`
	OriginName     string `name:"Origin Name" question:"What's the name of the origin database?"`
	OriginUsername string `name:"Origin Username" question:"What's the username of the origin database?"`
	OriginPassword string `name:"Origin Password" question:"What's the password of the origin database?"`
	/*
	EnvVarHost     string `name:"The Environment Variable for Host" question:"What's the environment variable for the host?"`
	EnvVarDatabase string `name:"The Environment Variable for Database" question:"What's the environment variable for the database name?"`
	EnvVarUsername string `name:"The Environment Variable for Username" question:"What's the environment variable for the username?"`
	EnvVarPassword string `name:"The Environment Variable for Password" question:"What's the environment variable for the password?"`
	*/
}

func (d Database) ValidateField(field string, input string) error {
	switch field {
	case "Driver":
		switch input {
		case "mysql":
			return nil
		default:
			return fmt.Errorf("Not support %s", input)
		}
	case "Version":
		switch input {
		case "5.6":
		case "5.7":
			return nil
		default:
			return fmt.Errorf("Not support %s", input)
		}
	}
	return nil
}
