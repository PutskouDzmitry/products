package data

import (
	"fmt"
	_ "github.com/lib/pq"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

type ProductData struct {
	db *gorm.DB
}

func NewProductData(db *gorm.DB) *ProductData {
	return &ProductData{db: db}
}

type Products struct {
	IdProduct int `gorm:"id_product"`
	Name string	`gorm:"name"`
	GroupOfProducts int `gorm:"group_of_products"`
	GroupOfProductsStruct GroupOfProducts `gorm:"ForeignKey:GroupOfProducts;References:IdGroupProduct"`
	Description string `gorm:"description"`
	ReleaseDate time.Time `gorm:"release_date"`
	ParametersId int	`gorm:"parameters_id"`
	ParametersIdStruct Parameters `gorm:"ForeignKey:ParametersId;References:IdParameter"`
}

type GroupOfProducts struct {
	IdGroupProduct int	`gorm:"id_group_product"`
	Name string
	GroupParameters int `gorm:"group_parameters"`
	GroupParametersStruct GroupOfParameters `gorm:"foreignKey:GroupParameters;references:IdGroupParameter"`
}

type GroupOfParameters struct {
	IdGroupParameter int `gorm:"id_group_parameter"`
	Name string
	Parameters int `gorm:"parameters"`
	ParametersStruct Parameters `gorm:"foreignKey:Parameters;references:IdParameter"`
}

type Parameters struct {
	IdParameter int `gorm:"id_parameter"`
	Name string
	UnitOfProduct string
}

func (p ProductData) ShowParametersWithSpecificGroup(action string) ([]Parameters, error) {
	var groupOfProduct []GroupOfProducts
	var groupOfParameters []GroupOfParameters
	result := p.db.Preload(clause.Associations).Where("group_of_products.name=?", action).Find(&groupOfProduct)
	if result.Error != nil {
		return nil, result.Error
	}
	result = p.db.Preload(clause.Associations).Find(&groupOfParameters)
	if result.Error != nil {
		return nil, result.Error
	}
	for i := 0; i < len(groupOfProduct); i++ {
		for j := 0; j < len(groupOfParameters); j++ {
			if groupOfProduct[i].GroupParametersStruct.Parameters == groupOfParameters[j].Parameters {
				groupOfProduct[i].GroupParametersStruct.ParametersStruct = groupOfParameters[j].ParametersStruct
			}
		}
	}
	var newParameters []Parameters
	for _, value := range groupOfProduct {
		newParameters = append(newParameters, value.GroupParametersStruct.ParametersStruct)
	}
	if len(newParameters) == 0 {
		return nil, fmt.Errorf("or you set incorrect parameter or we cannot find data with this parameter")
	}
	return newParameters, nil
}

func (p ProductData) ShowParametersWithoutSpecificGroup(action string) ([]Products, error) {
	var product []Products
	var parameter Parameters
	result := p.db.
		Preload(clause.Associations).
		Where("name=?", action).
		Find(&parameter)
	result = p.db.
		Preload(clause.Associations).
		Where("parameters_id!=?", parameter.IdParameter).
		Find(&product)
	if result.Error != nil {
		return nil, result.Error
	}
	if len(product) == 0 {
		return nil, fmt.Errorf("or you set incorrect parameter or we cannot find data with this parameter")
	}
	return product, nil
}

func (p ProductData) ShowProductWithSpecificProductGroups(action string) ([]Products, error) {
	var product []Products
	var parameter GroupOfProducts
	result := p.db.
		Preload(clause.Associations).
		Where("name=?", action).
		Find(&parameter)
	result = p.db.
		Preload(clause.Associations).
		Where("group_of_products=?", parameter.IdGroupProduct).
		Find(&product)
	if result.Error != nil {
		return nil, result.Error
	}
	if len(product) == 0 {
		return nil, fmt.Errorf("or you set incorrect parameter or we cannot find data with this parameter")
	}
	return product, nil
}

func (p ProductData) ShowProduct() ([]Products, error) {
	var product []Products
	var groupParameters []GroupOfParameters
	result := p.db.
		Preload(clause.Associations).
		Find(&product)
	if result.Error != nil {
		return nil, result.Error
	}
	result = p.db.
		Preload(clause.Associations).
		Find(&groupParameters)
	if result.Error != nil {
		return nil, result.Error
	}
	for i := 0; i < len(product); i++ {
		for j := 0; j < len(groupParameters); j++ {
			if product[i].GroupOfProductsStruct.GroupParameters == groupParameters[j].IdGroupParameter {
				product[i].GroupOfProductsStruct.GroupParametersStruct = groupParameters[j]
			}
		}
	}
	return product, nil
}

func (p ProductData) DeleteDataWithSpecialParameters(param string) error {
	var product []Products
	var value Parameters
	var groupParameters []GroupOfParameters
	result := p.db.
		Where("name=?", param).
		Find(&value).
		Exec(`DELETE FROM "products" WHERE parameters_id=?`, value.IdParameter)
	result = p.db.
		Preload(clause.Associations).
		Find(&product)
	if result.Error != nil {
		return result.Error
	}
	result = p.db.
		Preload(clause.Associations).
		Find(&groupParameters)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (p ProductData) ChangeDataIntoDb(paramOfGroupParam, paramOfGroupProduct2 string)  error {
	var idGroupOfParameters GroupOfParameters
	result := p.db.
		Where("name=?", paramOfGroupParam).
		Find(&idGroupOfParameters)
	if result.Error != nil {
		return result.Error
	}
	result = p.db.
		Model(&GroupOfProducts{}).
		Where("group_parameters=?", idGroupOfParameters.IdGroupParameter).
		Update("group_parameters", 4)
	if result.Error != nil {
		return result.Error
	}
	result = p.db.
		Model(&GroupOfProducts{}).
		Where("name=?", paramOfGroupProduct2).
		Update("group_parameters", idGroupOfParameters.IdGroupParameter)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (p ProductData) ShowGroupOfParamAndGroupOfProduct() ([]Products, error) {
	var product []Products
	var groupParameters []GroupOfParameters
	result := p.db.
		Preload(clause.Associations).
		Find(&product)
	if result.Error != nil {
		return nil, result.Error
	}
	result = p.db.
		Preload(clause.Associations).
		Find(&groupParameters)
	if result.Error != nil {
		return nil, result.Error
	}
	for i := 0; i < len(product); i++ {
		for j := 0; j < len(groupParameters); j++ {
			if product[i].GroupOfProductsStruct.GroupParameters == groupParameters[j].IdGroupParameter {
				product[i].GroupOfProductsStruct.GroupParametersStruct = groupParameters[j]
			}
		}
	}
	return product, nil
}