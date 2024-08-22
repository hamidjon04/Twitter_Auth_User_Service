package postgres

import (
	"auth/config"
	"auth/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	cfg := config.Load()
	db, err := ConnectToDB(cfg)
	if err != nil {
		t.Fatal(err)
	}

	repo := NewUserRepo(db)
	user := model.RegisterReq{Email: "test@test.com", Password: "123456"}
	_, err = repo.Register(&user)
	if err != nil {
		t.Fatal(err)
	}

	assert.NoError(t, err)
}

func TestGetUserByEmail(t *testing.T) {
	cfg := config.Load()
	db, err := ConnectToDB(cfg)
	if err != nil {
		t.Fatal(err)
	}

	repo := NewUserRepo(db)
	_, err = repo.GetUserByEmail("test@test.com")
	if err != nil {
		t.Fatal(err)
	}

	assert.NoError(t, err)
}

func TestSaveRefreshToken(t *testing.T) {
	cfg := config.Load()
    db, err := ConnectToDB(cfg)
    if err!= nil {
        t.Fatal(err)
    }

    repo := NewUserRepo(db)
    token := model.SaveToken{UserId: "43f1f0e7-ec9f-43d6-a42c-97145103e6ef", RefreshToken: "abc", ExpiresAt: "2022-12-31 23:59:59"}
    err = repo.SaveRefreshToken(&token)
    if err!= nil {
        t.Fatal(err)
    }

    assert.NoError(t, err)
}
