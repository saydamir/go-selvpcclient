package testing

import "github.com/selectel/go-selvpcclient/selvpcclient/resell/v2/users"

// TestListUsersResponseRaw represents a raw response from the List request.
const TestListUsersResponseRaw = `
{
    "users": [
        {
            "enabled": false,
            "id": "c4b7e0581b964c52a1597fe0931eccdf",
            "name": "User2"
        },
        {
            "enabled": true,
            "id": "4b2e452ed4c940bd87a88499eaf14c4f",
            "name": "User1"
        }
    ]
}
`

// TestListUsersSingleUserResponseRaw represents a raw response with a single user from the List request.
const TestListUsersSingleUserResponseRaw = `
{
    "users": [
        {
            "enabled": true,
            "id": "4b2e452ed4c940bd87a88499eaf14c4f",
            "name": "User1"
        }
    ]
}
`

// TestListUsersSingleUserResponse represents the unmarshalled TestListUsersSingleUserResponseRaw response.
var TestListUsersSingleUserResponse = []*users.User{
	{
		Enabled: true,
		Name:    "User1",
		ID:      "4b2e452ed4c940bd87a88499eaf14c4f",
	},
}

// TestCreateUserOptsRaw represents marshalled options for the Create request.
const TestCreateUserOptsRaw = `
{
    "user": {
        "name": "NewUser1",
        "password":"verysecret"
    }
}
`

// TestCreateUserOpts represent options for the Create request.
var TestCreateUserOpts = users.UserOpts{
	Name:     "NewUser1",
	Password: "verysecret",
}

// TestCreateUserResponseRaw represents a raw response from the Create request.
const TestCreateUserResponseRaw = `
{
    "user": {
        "enabled": true,
        "name": "NewUser1",
        "id": "4b2e452ed4c940bd87a88499eaf14c4f"
    }
}
`

// TestCreateUserResponse represents the unmarshalled TestCreateUserResponseRaw response.
var TestCreateUserResponse = &users.User{
	ID:      "4b2e452ed4c940bd87a88499eaf14c4f",
	Name:    "NewUser1",
	Enabled: true,
}
