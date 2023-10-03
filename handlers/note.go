package handlers

import (
	"fmt"
	"go-solid/database"
	"go-solid/models"

	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New(validator.WithRequiredStructEnabled())

func CreateNote(c *fiber.Ctx) error {
	db := database.DB
	var note models.Note

	if err := c.BodyParser(&note); err != nil {
		log.Printf("Couldn't parse the response body. \n{%s}", err.Error())
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err.Error()})
	}

	if err := validate.Struct(&note); err != nil {
		log.Printf("Validation Error %s", err.Error())
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err.Error()})
	}

	if err := db.Create(&note).Error; err != nil {
		log.Printf("Couldn't create note \n{%s}\n", err.Error())
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create note", "data": err.Error()})
	}

	log.Printf("Created Note {%d} Successfully.\n", note.ID)

	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "Created todo", "data": note})
}

func GetNote(c *fiber.Ctx) error {
	db := database.DB
	var note models.Note

	id, err := c.ParamsInt("id")

	if err != nil {
		log.Printf("Couldn't parse the id from params\n {%s}\n", err.Error())
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Couldn't parse the id", "data": err.Error()})
	}

	if err = db.First(&note, id).Error; err != nil {
		log.Printf("Couldn't get note \n{%s}\n", err.Error())
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't get note", "data": err.Error()})
	}

	return c.Status(200).JSON(note)
}

func GetAllNotes(c *fiber.Ctx) error {
	db := database.DB
	var notes []models.Note

	if err := db.Find(&notes).Error; err != nil {
		log.Printf("Couldn't get notes \n{%s}\n", err.Error())
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't get notes", "data": err.Error()})
	}

	return c.Status(200).JSON(notes)
}

func UpdateNote(c *fiber.Ctx) error {
	db := database.DB
	var note models.Note

	id, err := c.ParamsInt("id")

	if err != nil {
		log.Printf("Couldn't parse the id from params\n {%s}\n", err.Error())
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Couldn't parse the id", "data": err.Error()})
	}

	if err := db.First(&note, id).Error; err != nil {
		log.Printf("Couldn't get the note with id %d\n %s", id, err.Error())
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Couldn't find the note with id " + fmt.Sprint(id), "data": err.Error()})
	}

	if err := c.BodyParser(&note); err != nil {
		log.Printf("Couldn't parse the response body. \n{%s}", err.Error())
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err.Error()})
	}

	if err := validate.Struct(&note); err != nil {
		log.Printf("Validation Error %s", err.Error())
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err.Error()})
	}

	if err := db.Save(&note).Error; err != nil {
		log.Printf("Couldn't update the note.\n{%s}\n", err.Error())
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't update the note", "data": err.Error()})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Updated note successfully"})
}

func DeleteNote(c *fiber.Ctx) error {
	db := database.DB

	id, err := c.ParamsInt("id")

	if err != nil {
		log.Printf("Couldn't parse the id from params\n {%s}\n", err.Error())
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Couldn't parse the id", "data": err.Error()})
	}

	log.Printf("Deleting Note {%d}...\n", id)
	if err = db.Unscoped().Delete(&models.Note{}, id).Error; err != nil {
		log.Printf("Couldn't delete the note with %d from the database.\n {%s}\n", id, err)
	}
	log.Printf("Deleted Note {%d} Successfully.\n", id)

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Deleted note successfully"})
}
