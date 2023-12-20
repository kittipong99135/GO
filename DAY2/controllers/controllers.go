package controllers

import (
	m "go-day2/models"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func FactParams(c *fiber.Ctx) error {
	result := 1
	paramsStr := c.Params("number")
	paramsInt, _ := strconv.Atoi(paramsStr)
	for i := paramsInt; i > 0; i-- {
		result *= i
	}

	return c.JSON(result)
}

func AsciiConv(c *fiber.Ctx) error {
	var resultStr string
	strname := c.Query("tax_id")
	for _, v := range strname {
		s := strconv.Itoa(int(v))
		resultStr = resultStr + s + " "
	}
	return c.JSON(resultStr)
}

func Register(c *fiber.Ctx) error {
	user := new(m.RegisBody)
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	validate := validator.New()
	errors := validate.Struct(user)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors.Error())
	}
	if !containUsername(user.Username) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Username is Invalid.",
		})
	}
	if !containLineIid(user.LineId) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "LineID is Invalid.",
		})
	}
	if !containPhone(user.PhoneNumber) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "PhoneNumber is Invalid.",
		})
	}
	if !containPhone(user.PhoneNumber) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "PhoneNumber is Invalid.",
		})
	}
	if !containURL(user.WUrl) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "URL is Invalid.",
		})
	}
	return c.JSON(user)
}

func containUsername(str string) bool {
	condiCha := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
		"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
		"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "-", "_"}
	return containText(str, condiCha)
}

func containLineIid(str string) bool {
	condiCha := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
		"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
		"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "-", "_", "@", "."}
	return containText(str, condiCha)
}

func containURL(str string) bool {
	strSplit := strings.Split(str, ".")
	condiCha := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
		"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	return containText(strSplit[0], condiCha)
}

func containPhone(str string) bool {
	condiCha := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	if len(str) != 10 {
		return false
	} else {
		for _, valueStr := range str {
			count := 0
			for _, valueCondiCha := range condiCha {
				if string(valueStr) == string(valueCondiCha) {
					count++
				}
			}
			if count < 1 {
				return false
			}
		}
		return true
	}
}

func containText(str string, arr []string) bool {
	for _, valueStr := range str {
		count := 0
		for _, valueCondiCha := range arr {
			if string(valueStr) == string(valueCondiCha) {
				count++
			}
		}
		if count < 1 {
			return false
		}
	}
	return true
}
