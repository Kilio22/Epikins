package controllers

import (
	"github.com/gofiber/fiber/v2"
)

var Cities = []string{
	"TIR",
	"BRU",
	"COT",
	"BDX",
	"RUN",
	"LIL",
	"LYN",
	"MAR",
	"MPL",
	"MLH",
	"NCY",
	"NAN",
	"NCE",
	"PAR",
	"REN",
	"STG",
	"TLS",
	"BER",
	"BAR",
}

func GetCitiesController(c *fiber.Ctx) error {
	return c.JSON(Cities)
}
