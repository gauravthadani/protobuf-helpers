package equals

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.temporal.io/api/cloud/identity/v1"
	"google.golang.org/protobuf/encoding/protojson"
)

func TestEquals(t *testing.T) {
	a := &identity.AccountAccess{
		RoleDeprecated: "random Value",
		Role:           identity.AccountAccess_ROLE_ADMIN,
	}
	b := &identity.AccountAccess{
		RoleDeprecated: "random Value",
		Role:           identity.AccountAccess_ROLE_ADMIN,
	}

	t.Run("Equals", func(t *testing.T) {
		assert.True(t, Equals(a, b, false))

	})

	t.Run("Different Role Deprecated Field", func(t *testing.T) {
		a.RoleDeprecated = "different Value"
		assert.False(t, Equals(a, b, false))

	})

	t.Run("Different Role", func(t *testing.T) {
		a.Role = identity.AccountAccess_ROLE_DEVELOPER
		assert.False(t, Equals(a, b, false))
	})
}

func TestEqualsIgnoreDeprecated(t *testing.T) {
	a := &identity.AccountAccess{
		RoleDeprecated: "random Value",
		Role:           identity.AccountAccess_ROLE_ADMIN,
	}
	b := &identity.AccountAccess{
		RoleDeprecated: "random Value",
		Role:           identity.AccountAccess_ROLE_ADMIN,
	}

	t.Run("Equals", func(t *testing.T) {
		assert.True(t, Equals(a, b, true))
	})

	t.Run("Different Role Deprecated Field", func(t *testing.T) {
		a.RoleDeprecated = "different Value"
		assert.True(t, Equals(a, b, true))
	})

	t.Run("Different Role", func(t *testing.T) {
		a.Role = identity.AccountAccess_ROLE_DEVELOPER
		assert.False(t, Equals(a, b, true))
	})

}

func TestCustomerExample(t *testing.T) {
	aData := []byte(`{"role_deprecated":"admin" ,"role":2}`)
	bData := []byte(`{"role":2}`)

	// Create a new protobuf message
	a := &identity.AccountAccess{}
	b := &identity.AccountAccess{}
	// Unmarshal JSON into the protobuf message
	err := protojson.Unmarshal(aData, a)
	if err != nil {
		fmt.Printf("Error unmarshaling JSON: %v\n", err)
		t.FailNow()
	}

	err = protojson.Unmarshal(bData, b)
	if err != nil {
		fmt.Printf("Error unmarshaling JSON: %v\n", err)
		t.FailNow()
	}

	assert.True(t, Equals(a, b, true))

}

func TestCustomerExampleSpec(t *testing.T) {
	aData := []byte(`{
    "access": {
        "accountAccess": {
            "role": "ROLE_READ",
            "roleDeprecated": "read"
        },
        "namespaceAccesses": {
            "f2f-dev.67e09": {
                "permission": "PERMISSION_WRITE",
                "permissionDeprecated": "write"
            },
            "maintenance-non-dev.68e09": {
                "permission": "PERMISSION_WRITE",
                "permissionDeprecated": "write"
            }
        }
    },
    "email": "john.travolta@pulpfiction.com"
}`)

	bData := []byte(`{
    "access": {
        "accountAccess": {
            "role": "ROLE_READ"
        },
        "namespaceAccesses": {
            "f2f-dev.67e09": {
                "permission": "PERMISSION_WRITE"
            },
            "maintenance-non-dev.68e09": {
                "permission": "PERMISSION_WRITE"
            }
        }
    },
    "email": "john.travolta@pulpfiction.com"
}`)

	// Create a new protobuf message
	a := &identity.UserSpec{}
	b := &identity.UserSpec{}
	// Unmarshal JSON into the protobuf message
	err := protojson.Unmarshal(aData, a)
	if err != nil {
		fmt.Printf("Error unmarshaling JSON: %v\n", err)
		t.FailNow()
	}

	err = protojson.Unmarshal(bData, b)
	if err != nil {
		fmt.Printf("Error unmarshaling JSON: %v\n", err)
		t.FailNow()
	}
	assert.True(t, Equals(a, b, true))
}
