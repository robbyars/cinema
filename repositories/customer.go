package repositories

import (
	"cinema/structs"
	"errors"

	"gorm.io/gorm"
)

type CustRepository interface {
	Login(user structs.LoginRequest) (result structs.Customer, err error)
	SignUp(user structs.Customer) (err error)
	Update(user structs.Customer) (err error)
	Delete(user structs.Customer) (err error)
	GetList() (users []structs.Customer, err error)
}

type customerRepository struct {
	db *gorm.DB
}

func CustNewRepository(database *gorm.DB) CustRepository {
	return &customerRepository{
		db: database,
	}
}

func (r *customerRepository) Login(customer structs.LoginRequest) (result structs.Customer, err error) {
	err = r.db.Where("username = ?", customer.Username).First(&result).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return result, err
	}

	return result, nil
}

func (r *customerRepository) SignUp(customer structs.Customer) (err error) {
	err = r.db.Create(&customer).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *customerRepository) Delete(customer structs.Customer) (err error) {
	result := r.db.Where("username = ?", customer.Username).Delete(&structs.Customer{})
	if result.RowsAffected == 0 {
		return errors.New("customer with the specified username not found")
	}

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *customerRepository) Update(customer structs.Customer) (err error) {
	result := r.db.Model(&structs.Customer{}).Where("username = ?", customer.Username).
		Updates(map[string]interface{}{
			"fullname": customer.Fullname,
			"email":    customer.Email,
			"phone":    customer.Phone,
		})
	if result.RowsAffected == 0 {
		return errors.New("customer with the specified username not found")
	}

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *customerRepository) GetList() (customer []structs.Customer, err error) {
	err = r.db.Find(&customer).Error
	if err != nil {
		return nil, err
	}

	return customer, nil
}
