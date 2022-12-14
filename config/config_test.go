package config

import (
	"fmt"
	"testing"
)

func TestSetEnvNew(t *testing.T) {
	wantPort := 3333
	t.Setenv("PORT", fmt.Sprint(wantPort))
	wantEnv := "TestEnv"
	t.Setenv("TODO_ENV", fmt.Sprint(wantEnv))
	wantDbHost := "0.0.0.0"
	t.Setenv("TODO_DB_HOST", fmt.Sprint(wantDbHost))
	wantDbPort := 3334
	t.Setenv("TODO_DB_PORT", fmt.Sprint(wantDbPort))
	wantDbUser := "test"
	t.Setenv("TODO_DB_USER", fmt.Sprint(wantDbUser))
	wantDbPassword := "test"
	t.Setenv("TODO_DB_PASSWORD", fmt.Sprint(wantDbPassword))
	wantDbName := "test"
	t.Setenv("TODO_DB_NAME", fmt.Sprint(wantDbName))
	wantRedisHost := "127.0.0.9"
	t.Setenv("TODO_REDIS_HOST", fmt.Sprint(wantRedisHost))
	wantRedisPort := 3335
	t.Setenv("TODO_REDIS_PORT", fmt.Sprint(wantRedisPort))

	got, err := New()
	if err != nil {
		t.Fatalf("cannot create config: %v", err)
	}
	if got.Port != wantPort {
		t.Errorf("want %d, but %d", wantPort, got.Port)
	}
	if got.Env != wantEnv {
		t.Errorf("want %s, but %s", wantEnv, got.Env)
	}
	if got.DBHost != wantDbHost {
		t.Errorf("want %s, but %s", wantDbHost, got.DBHost)
	}
	if got.DBPort != wantDbPort {
		t.Errorf("want %d, but %d", wantDbPort, got.DBPort)
	}
	if got.DBUser != wantDbUser {
		t.Errorf("want %s, but %s", wantDbUser, got.DBUser)
	}
	if got.DBPassword != wantDbPassword {
		t.Errorf("want %s, but %s", wantDbPassword, got.DBPassword)
	}
	if got.DBName != wantDbName {
		t.Errorf("want %s, but %s", wantDbName, got.DBName)
	}
	if got.RedisHost != wantRedisHost {
		t.Errorf("want %s, but %s", wantRedisHost, got.RedisHost)
	}
	if got.RedisPort != wantRedisPort {
		t.Errorf("want %d, but %d", wantRedisPort, got.RedisPort)
	}
}

func TestDefaultEnvNew(t *testing.T) {
	wantPort := 8080
	wantEnv := "dev"
	wantDbHost := "127.0.0.1"
	wantDbPort := 33306
	wantDbUser := "todo"
	wantDbPassword := "todo"
	wantDbName := "todo"
	wantRedisHost := "127.0.0.1"
	wantRedisPort := 6379

	got, err := New()
	if err != nil {
		t.Fatalf("cannot create config: %v", err)
	}
	if got.Port != wantPort {
		t.Errorf("want %d, but %d", wantPort, got.Port)
	}
	if got.Env != wantEnv {
		t.Errorf("want %s, but %s", wantEnv, got.Env)
	}
	if got.DBHost != wantDbHost {
		t.Errorf("want %s, but %s", wantDbHost, got.DBHost)
	}
	if got.DBPort != wantDbPort {
		t.Errorf("want %d, but %d", wantDbPort, got.DBPort)
	}
	if got.DBUser != wantDbUser {
		t.Errorf("want %s, but %s", wantDbUser, got.DBUser)
	}
	if got.DBPassword != wantDbPassword {
		t.Errorf("want %s, but %s", wantDbPassword, got.DBPassword)
	}
	if got.DBName != wantDbName {
		t.Errorf("want %s, but %s", wantDbName, got.DBName)
	}
	if got.RedisHost != wantRedisHost {
		t.Errorf("want %s, but %s", wantRedisHost, got.RedisHost)
	}
	if got.RedisPort != wantRedisPort {
		t.Errorf("want %d, but %d", wantRedisPort, got.RedisPort)
	}
}
