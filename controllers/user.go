package controllers

import (
	"math/rand"
	"strconv"

	"mylovepp/ent"
	"mylovepp/services"

	"github.com/gofiber/fiber/v2"
)

type UserController struct{}

var userService = new(services.UserService)

func (UserController) GetUsers(context *fiber.Ctx) error {
	users := userService.GetUsers()

	if users != nil {
		return context.Status(fiber.StatusOK).JSON(users)
	}
	return context.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"message": "Not Found",
	})
}

func (UserController) GetUser(context *fiber.Ctx) error {

	id := context.Params("id")
	userId, err := strconv.Atoi(id)
	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad Request",
		})
	}

	user := userService.GetUser(userId)
	if user != nil {
		return context.Status(fiber.StatusOK).JSON(user)
	}

	return context.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"message": "Not Found",
	})
}

func (UserController) AddUser(context *fiber.Ctx) error {
	user := new(ent.User)
	if err := context.BodyParser(user); err != nil {
		println(err.Error())
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad Request",
		})
	}

	userService.AddUser(user)
	return context.Status(fiber.StatusCreated).JSON(user)
}

func (UserController) UpdateUser(context *fiber.Ctx) error {
	user := new(ent.User)
	if err := context.BodyParser(user); err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad Request",
		})
	}

	id := context.Params("id")
	userId, err := strconv.Atoi(id)
	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad Request",
		})
	}
	user.ID = userId
	success := userService.UpdateUser(user)
	if !success {
		return context.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Not Found",
		})
	}

	return context.Status(fiber.StatusOK).JSON(user)
}

func (UserController) AddUserWithAccounts(context *fiber.Ctx) error {
	user := new(ent.User)
	if err := context.BodyParser(user); err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad Request",
		})
	}

	randomNumber := func() string {
		return strconv.Itoa(rand.Intn(10000000000))
	}
	accounts := []*ent.Account{}
	accounts = append(accounts, &ent.Account{AssetID: 1, AccountNo: randomNumber()})
	accounts = append(accounts, &ent.Account{AssetID: 2, AccountNo: randomNumber()})
	accounts = append(accounts, &ent.Account{AssetID: 3, AccountNo: randomNumber()})

	err := userService.AddUserWithAccounts(user, accounts)
	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad Request",
		})
	}

	return context.Status(fiber.StatusCreated).JSON(fiber.Map{
		"user":     user,
		"accounts": accounts,
	})
}

func (UserController) DeleteUser(context *fiber.Ctx) error {
	id := context.Params("id")
	userId, err := strconv.Atoi(id)
	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad Request",
		})
	}
	success := userService.DeleteUser(userId)
	if !success {
		return context.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Not Found",
		})
	}
	return context.SendStatus(fiber.StatusOK)
}
