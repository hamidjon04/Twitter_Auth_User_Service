package postgres

import (
	"auth/config"
	"auth/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	cfg := config.LoadConfig()
	db, err := ConnectToDB(cfg)
	if err != nil {
		t.Fatal(err)
	}

	repo := NewUserRepo(db)
	user := model.RegisterReq{Email: "test@tests.com", Password: "123456",Username: "test"}
	_, err = repo.Register(&user)
	if err != nil {
		t.Fatal(err)
	}

	assert.NoError(t, err)
}

func TestGetUserByEmail(t *testing.T) {
	cfg := config.LoadConfig()
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
	cfg := config.LoadConfig()
	db, err := ConnectToDB(cfg)
	if err != nil {
		t.Fatal(err)
	}

	repo := NewUserRepo(db)
	token := model.SaveToken{UserId: "43f1f0e7-ec9f-43d6-a42c-97145103e6ef", RefreshToken: "abc", ExpiresAt: 1000000}
	err = repo.SaveRefreshToken(&token)
	if err != nil {
		t.Fatal(err)
	}

	assert.NoError(t, err)
}

func TestResetPass(t *testing.T) {
	cfg := config.LoadConfig()
	db, err := ConnectToDB(cfg)
	if err != nil {
		t.Fatal(err)
	}

	repo := NewUserRepo(db)
	user := model.ResetPassReq{Email: "test@test.com", Password: "123456"}
	_, err = repo.ResetPass(&user)
	if err != nil {
		t.Fatal(err)
	}

	assert.NoError(t, err)
}

func TestChangePass(t *testing.T) {
	cfg := config.LoadConfig()
	db, err := ConnectToDB(cfg)
	if err != nil {
		t.Fatal(err)
	}

	repo := NewUserRepo(db)
	user := model.ChangePassReq{UserId: "43f1f0e7-ec9f-43d6-a42c-97145103e6ef", NowPassword: "123456", NewPassword: "654321"}
	resp, err := repo.ChangePass(&user)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, resp.Message, "Password changed successfully")
}

func TestInvalidateRefreshToken(t *testing.T) {
	cfg := config.LoadConfig()
	db, err := ConnectToDB(cfg)
	if err != nil {
		t.Fatal(err)
	}

	repo := NewUserRepo(db)
	err = repo.InvalidateRefreshToken("43f1f0e7-ec9f-43d6-a42c-97145103e6ef")
	if err != nil {
		t.Fatal(err)
	}

	assert.NoError(t, err)
}


func TestIsRefreshTokenValid(t *testing.T) {
	cfg := config.LoadConfig()
    db, err := ConnectToDB(cfg)
    if err != nil {
        t.Fatal(err)
    }

    repo := NewUserRepo(db)
    valid,err := repo.IsRefreshTokenValid("43f1f0e7-ec9f-43d6-a42c-97145103e6ef")
	if err !=nil{
		t.Fatal(err)
	}
    if !valid {
        t.Fatal("Expected valid token, got invalid")
    }
}
