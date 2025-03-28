package seeds

import (
	"log"

	"github.com/victorlabussiere/go-echo-gorm-example/internal/model"
	"gorm.io/gorm"
)

type InsertInitialProducts struct{}

func (c *InsertInitialProducts) Up(db *gorm.DB) error {
	log.Println("Inserindo seeds 003-insert-products")
	products := []model.Product{
		{Name: "Telefone", Value: 3000, CreatedAt: nil, UpdatedAt: nil, CategoryId: 2},
		{Name: "Relogio", Value: 700, CreatedAt: nil, UpdatedAt: nil, CategoryId: 1},
		{Name: "Tênis", Value: 300, CreatedAt: nil, UpdatedAt: nil, CategoryId: 1},
	}

	for _, product := range products {
		if err := db.Where("name = ?", product.Name).First(&product).Error; err != nil {
			if err == gorm.ErrRecordNotFound { // Se o produto não foi encontrado, cria um novo
				if err := db.Create(&product).Error; err != nil {
					log.Fatalln("Erro na inserção dos dados:", err.Error())
					return err
				}
				log.Printf("Produto %s inserido com sucesso", product.Name)
			} else {
				log.Fatalln("Erro ao buscar o produto:", err.Error())
				return err
			}
		} else {
			log.Printf("Produto %s já existe, não inserido", product.Name)
		}
	}
	return nil
}

func (c *InsertInitialProducts) Down(db *gorm.DB) error {
	log.Println("Revertendo seed 003")
	return db.Where("name IN ?", []string{"Telefone", "Relogio", "Tênis"}).Delete(&model.Product{}).Error
}
