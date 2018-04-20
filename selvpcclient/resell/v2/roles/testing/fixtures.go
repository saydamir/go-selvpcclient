package testing

import "github.com/selectel/go-selvpcclient/selvpcclient/resell/v2/roles"

// TestListProjectResponseRaw represents a raw response from ListProject requests.
const TestListProjectResponseRaw = `
{
    "roles": [
        {
            "project_id": "49338ac045f448e294b25d013f890317",
            "user_id": "b006a55e3a904472824061d64d61be75"
        },
        {
            "project_id": "49338ac045f448e294b25d013f890317",
            "user_id": "a699c05690ec44ad8cb2a6fe80b29e70"
        },
        {
            "project_id": "49338ac045f448e294b25d013f890317",
            "user_id": "763eecfaeb0c8e9b76ab12a82eb4c11"
        }
    ]
}
`

// TestListResponseSingleRaw represents a raw response with a single role from List requests.
const TestListResponseSingleRaw = `
{
    "roles": [
        {
            "project_id": "49338ac045f448e294b25d013f890317",
            "user_id": "763eecfaeb0c8e9b76ab12a82eb4c11"
        }
    ]
}
`

// TestListResponseSingle represents the unmarshalled TestListResponseSingleRaw response.
var TestListResponseSingle = []*roles.Role{
	{
		ProjectID: "49338ac045f448e294b25d013f890317",
		UserID:    "763eecfaeb0c8e9b76ab12a82eb4c11",
	},
}

// TestListUserResponseRaw represents a raw response from ListUser requests.
const TestListUserResponseRaw = `
{
    "roles": [
        {
            "project_id": "81800a8ec3fc49fca2cf00857de3ae9d",
            "user_id": "763eecfaeb0c8e9b76ab12a82eb4c11"
        },
        {
            "project_id": "d7452adc9769422a908edfd2281d7c55",
            "user_id": "763eecfaeb0c8e9b76ab12a82eb4c11"
        },
        {
            "project_id": "49338ac045f448e294b25d013f890317",
            "user_id": "763eecfaeb0c8e9b76ab12a82eb4c11"
        }
    ]
}
`

// TestManyRolesInvalidResponseRaw represents a raw invalid response with several roles.
const TestManyRolesInvalidResponseRaw = `
{
    "roles": [
        {
            "project_id": 123
        }
    ]
}
`
