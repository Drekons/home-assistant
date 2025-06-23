package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/Drekons/home-assistant/backend/cmd/app"
	"github.com/Drekons/home-assistant/backend/internal/handler/auth"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: cli <command> [arguments]")
		fmt.Println("Available commands:")
		fmt.Println("  register --username <username> --password <password> --email <email>")
		os.Exit(1)
	}

	ctx := context.Background()
	application := app.NewApp(ctx)
	defer application.Shutdown(ctx)

	command := os.Args[1]

	switch command {
	case "register":
		registerUser(ctx, application.Deps)
	default:
		fmt.Printf("Unknown command: %s\n", command)
		os.Exit(1)
	}
}

func registerUser(ctx context.Context, deps *app.Deps) {
	registerCmd := flag.NewFlagSet("register", flag.ExitOnError)
	username := registerCmd.String("username", "", "Username for the new user")
	password := registerCmd.String("password", "", "Password for the new user")
	email := registerCmd.String("email", "", "Email for the new user")

	err := registerCmd.Parse(os.Args[2:])
	if err != nil {
		fmt.Println("Error parsing command-line arguments:", err)
		os.Exit(1)
	}

	if *username == "" || *password == "" || *email == "" {
		fmt.Println("Username, password, and email are required")
		registerCmd.PrintDefaults()
		os.Exit(1)
	}

	handler := auth.NewRegistry(deps)

	user, err := handler.Register(ctx, *username, *password, *email)
	if err != nil {
		fmt.Printf("Error registering user: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("User created: %+v\n", user.ID)

	fmt.Println("User registered successfully")
}
