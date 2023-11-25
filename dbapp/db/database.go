package db

import (
	"dbapp/ent"
	"fmt"
	"log"
)

type IDatabase interface {
	// CreateCountry(code string, name string, currency string, sanctioned bool, armySize int) (*ent.Country, error)
	// SaveCountry(country *ent.Country) (*ent.Country, error)
	// UpdateCountry(country *ent.Country) (*ent.Country, error)
	// DeleteCountry(country *ent.Country) error
	// CreateCustomer(firstName string, lastName string, email string, phone string, address string, favouriteColour *string, country ent.Country) (*ent.Customer, error)
	// SaveCustomer(customer *ent.Customer) (*ent.Customer, error)
	// UpdateCustomer(customer *ent.Customer) (*ent.Customer, error)
	// DeleteCustomer(customer *ent.Customer) error
	// CreateProductLine(productCode int, itemName string, itemDescription string, itemPrice float64, itemKg float64, manufacturer ent.Country, assembler ent.Country, bannedCountries []*ent.Country) (*ent.ProductLine, error)
	// SaveProductLine(productLine ent.ProductLine) (*ent.ProductLine, error)
	// UpdateProductLine(productLine ent.ProductLine) (*ent.ProductLine, error)
	// DeleteProductLine(productLine *ent.ProductLine) error
	// GetAllCountries() []*ent.Country
	// GetAllProductLines() []*ent.ProductLine
	// GetAllCustomers() []*ent.Customer
	// GetAllOrders() []*ent.Order
	// GetAllOrderItems() []*ent.OrderItems
}

type Database struct {
	client *ent.Client
}

func Init() *Database {
	dbuser := "appuser"
	dbpass := "pass"
	dbhost := "localhost"
	dbport := "3306"
	dbname := "poc_ent"
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True", dbuser, dbpass, dbhost, dbport, dbname)
	client, err := ent.Open("mysql", connectionString)
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	return &Database{client: client}

}

// func (d *Database) CreateCountry(code string, name string, currency string, sanctioned bool, armySize int) (*ent.Country, error) {
// 	createdCountry, err := d.client.Country.Create().
// 		SetCode(code).
// 		SetName(name).
// 		SetCurrency(currency).
// 		SetSanctioned(sanctioned).
// 		SetArmySize(armySize).
// 		Save(context.Background())
// 	if err != nil {
// 		fmt.Printf("failed creating country: %v\n", err)
// 		return nil, fmt.Errorf("failed creating country: %v", err)
// 	}
// 	fmt.Printf("Country was created: %+v\n", spew.Sdump(*createdCountry))
// 	return createdCountry, nil
// }
// func (d *Database) SaveCountry(country *ent.Country) (*ent.Country, error) {
// 	savedCountry, err := d.client.Country.Create().
// 		SetCode(country.Code).
// 		SetName(country.Name).
// 		SetCurrency(country.Currency).
// 		SetSanctioned(country.Sanctioned).
// 		SetArmySize(country.ArmySize).
// 		Save(context.Background())
// 	if err != nil {
// 		fmt.Printf("failed saving country: %v\n", err)
// 		return nil, fmt.Errorf("failed saving country: %v", err)
// 	}
// 	fmt.Printf("Country was saved: %+v\n", spew.Sdump(*savedCountry))
// 	return savedCountry, nil
// }

// func (d *Database) UpdateCountry(country *ent.Country) (*ent.Country, error) {
// 	foundCountry, err := d.client.Country.Get(context.Background(), country.ID)
// 	if err != nil {
// 		fmt.Printf("failed updating country: %v\n", err)
// 		return nil, fmt.Errorf("failed updating country: %v", err)
// 	}
// 	updatedCountry, err := foundCountry.Update().
// 		SetCode(country.Code).
// 		SetName(country.Name).
// 		SetCurrency(country.Currency).
// 		SetSanctioned(country.Sanctioned).
// 		SetArmySize(country.ArmySize).
// 		Save(context.Background())
// 	if err != nil {
// 		fmt.Printf("failed updating country: %v\n", err)
// 		return nil, fmt.Errorf("failed updating country: %v", err)
// 	}
// 	fmt.Printf("Country was updated from: %+v\n to: %+v\n", spew.Sdump(*foundCountry), spew.Sdump(*updatedCountry))
// 	return updatedCountry, nil
// }
// func (d *Database) DeleteCountry(country *ent.Country) error {
// 	foundCountry, err := d.client.Country.Get(context.Background(), country.ID)
// 	if err != nil {
// 		fmt.Printf("failed deleting country: %v\n", err)
// 		return fmt.Errorf("failed updating country: %v", err)
// 	}
// 	err = d.client.Country.DeleteOne(foundCountry).Exec(context.Background())
// 	if err != nil {
// 		fmt.Printf("failed deleting country: %v\n", err)
// 		return fmt.Errorf("failed deleting country: %v", err)
// 	}
// 	fmt.Printf("Country was deleted: %+v\n", spew.Sdump(*foundCountry))
// 	return nil
// }

// func (d *Database) CreateCustomer(firstName string, lastName string, email string, phone string, address string, favouriteColour *string, country ent.Country) (*ent.Customer, error) {
// 	createdCustomer, err := d.client.Customer.Create().
// 		SetFirstName(firstName).
// 		SetLastName(lastName).
// 		SetEmail(email).
// 		SetPhone(phone).
// 		SetFavouriteColour(*favouriteColour).
// 		SetAddress(address).SetCountry(&country).Save(context.Background())
// 	if err != nil {
// 		fmt.Printf("failed creating customer: %v\n", err)
// 		return nil, fmt.Errorf("failed creating customer: %v", err)
// 	}
// 	fmt.Printf("Customer was created: %+v\n", spew.Sdump(*createdCustomer))
// 	return createdCustomer, nil
// }

// func (d *Database) SaveCustomer(customer *ent.Customer) (*ent.Customer, error) {
// 	savedCustomer, err := d.client.Customer.Create().
// 		SetFirstName(customer.FirstName).
// 		SetLastName(customer.LastName).
// 		SetEmail(customer.Email).
// 		SetPhone(customer.Phone).
// 		SetFavouriteColour(*customer.FavouriteColour).
// 		SetAddress(customer.Address).SetCountry(customer.Edges.Country).Save(context.Background())
// 	if err != nil {
// 		fmt.Printf("failed saving customer: %v\n", err)
// 		return nil, fmt.Errorf("failed saving customer: %v", err)
// 	}
// 	fmt.Printf("Customer was saved: %+v\n", spew.Sdump(*savedCustomer))
// 	return savedCustomer, nil
// }

// func (d *Database) UpdateCustomer(customer *ent.Customer) (*ent.Customer, error) {
// 	foundCustomer, err := d.client.Customer.Get(context.Background(), customer.ID)
// 	if err != nil {
// 		fmt.Printf("failed updating customer: %v\n", err)
// 		return nil, fmt.Errorf("failed updating customer: %v", err)
// 	}
// 	updatedCustomer, err := d.client.Customer.Create().
// 		SetFirstName(customer.FirstName).
// 		SetLastName(customer.LastName).
// 		SetEmail(customer.Email).
// 		SetPhone(customer.Phone).
// 		SetFavouriteColour(*customer.FavouriteColour).
// 		SetAddress(customer.Address).SetCountry(customer.Edges.Country).Save(context.Background())
// 	if err != nil {
// 		fmt.Printf("failed creating customer: %v\n", err)
// 		return nil, fmt.Errorf("failed creating customer: %v", err)
// 	}
// 	fmt.Printf("Customer was updated from : %+v\n to : %+v\n", spew.Sdump(*foundCustomer), spew.Sdump(*updatedCustomer))
// 	return updatedCustomer, nil
// }

// func (d *Database) DeleteCustomer(customer *ent.Customer) error {
// 	foundCustomer, err := d.client.Customer.Get(context.Background(), customer.ID)
// 	if err != nil {
// 		fmt.Printf("failed deleting customer: %v\n", err)
// 		return fmt.Errorf("failed updating customer: %v", err)
// 	}
// 	err = d.client.Customer.DeleteOne(foundCustomer).Exec(context.Background())
// 	if err != nil {
// 		fmt.Printf("failed deleting customer: %v\n", err)
// 		return fmt.Errorf("failed deleting customer: %v", err)
// 	}
// 	fmt.Printf("Customer was deleted: %+v\n", spew.Sdump(*foundCustomer))
// 	return nil
// }

// func (d *Database) CreateProductLine(productCode int, itemName string, itemDescription string, itemPrice float64, itemKg float64, manufacturer ent.Country, assembler ent.Country, bannedCountries []*ent.Country) (*ent.ProductLine, error) {
// 	productLine, err := d.client.ProductLine.Create().
// 		SetProductCode(productCode).
// 		SetItemName(itemName).
// 		SetItemDescription(itemDescription).
// 		SetItemPrice(itemPrice).
// 		SetItemKg(itemKg).
// 		SetManufactureCountry(&manufacturer).
// 		SetAssemblerCountry(&assembler).
// 		AddBannedInCountries(bannedCountries...).
// 		Save(context.Background())
// 	if err != nil {
// 		fmt.Printf("failed creating product line: %v\n", err)
// 		return nil, fmt.Errorf("failed creating product line: %v", err)
// 	}
// 	fmt.Printf("Product line was created: %+v\n", spew.Sdump(*productLine))
// 	return productLine, nil
// }
// func (d *Database) SaveProductLine(productLine ent.ProductLine) (*ent.ProductLine, error) {
// 	savedProductLine, err := d.client.ProductLine.Create().
// 		SetProductCode(productLine.ProductCode).
// 		SetItemName(productLine.ItemName).
// 		SetItemDescription(productLine.ItemDescription).
// 		SetItemPrice(productLine.ItemPrice).
// 		SetItemKg(productLine.ItemKg).
// 		SetManufactureCountry(productLine.Edges.ManufactureCountry).
// 		SetAssemblerCountry(productLine.Edges.AssemblerCountry).
// 		AddBannedInCountries(productLine.Edges.BannedInCountries...).
// 		Save(context.Background())
// 	if err != nil {
// 		fmt.Printf("failed saving product line: %v\n", err)
// 		return nil, fmt.Errorf("failed saving product line: %v", err)
// 	}
// 	fmt.Printf("Product line was saved: %+v\n", spew.Sdump(*savedProductLine))
// 	return savedProductLine, nil
// }

// func (d *Database) UpdateProductLine(productLine ent.ProductLine) (*ent.ProductLine, error) {
// 	foundProductLine, err := d.client.ProductLine.Get(context.Background(), productLine.ID)
// 	if err != nil {
// 		fmt.Printf("failed updating product line: %v\n", err)
// 		return nil, fmt.Errorf("failed updating product line: %v", err)
// 	}

// 	updatedProductLine, err := d.client.ProductLine.Create().
// 		SetProductCode(productLine.ProductCode).
// 		SetItemName(productLine.ItemName).
// 		SetItemDescription(productLine.ItemDescription).
// 		SetItemPrice(productLine.ItemPrice).
// 		SetItemKg(productLine.ItemKg).
// 		SetManufactureCountry(productLine.Edges.ManufactureCountry).
// 		SetAssemblerCountry(productLine.Edges.AssemblerCountry).
// 		AddBannedInCountries(productLine.Edges.BannedInCountries...).
// 		Save(context.Background())
// 	if err != nil {
// 		fmt.Printf("failed updating product line: %v\n", err)
// 		return nil, fmt.Errorf("failed updating product line: %v", err)
// 	}
// 	fmt.Printf("Product line was updated from: %+v\n to: %+v\n", spew.Sdump(*foundProductLine), spew.Sdump(*updatedProductLine))
// 	return updatedProductLine, nil
// }

// func (d *Database) DeleteProductLine(productLine *ent.ProductLine) error {
// 	foundProductLine, err := d.client.ProductLine.Get(context.Background(), productLine.ID)
// 	if err != nil {
// 		fmt.Printf("failed deleting productLine: %v\n", err)
// 		return fmt.Errorf("failed updating productLine: %v", err)
// 	}
// 	err = d.client.ProductLine.DeleteOne(foundProductLine).Exec(context.Background())
// 	if err != nil {
// 		fmt.Printf("failed deleting productLine: %v\n", err)
// 		return fmt.Errorf("failed deleting productLine: %v", err)
// 	}
// 	fmt.Printf("ProductLine was deleted: %+v\n", spew.Sdump(*foundProductLine))
// 	return nil
// }

// func (d *Database) GetAllCountries() []*ent.Country {
// 	countries, err := d.client.Country.Query().All(context.Background())
// 	if err != nil {
// 		fmt.Printf("failed getting all countries: %v\n", err)
// 	}
// 	fmt.Printf("Countries: %+v\n", spew.Sdump(countries))
// 	return countries
// }
// func (d *Database) GetAllProductLines() []*ent.ProductLine {
// 	productLine, err := d.client.ProductLine.Query().All(context.Background())
// 	if err != nil {
// 		fmt.Printf("failed getting all productLine: %v\n", err)
// 	}
// 	fmt.Printf("ProductLine: %+v\n", spew.Sdump(productLine))
// 	return productLine
// }
// func (d *Database) GetAllCustomers() []*ent.Customer {
// 	customers, err := d.client.Customer.Query().All(context.Background())
// 	if err != nil {
// 		fmt.Printf("failed getting all customers: %v\n", err)
// 	}
// 	fmt.Printf("Customers: %+v\n", spew.Sdump(customers))
// 	return customers
// }
// func (d *Database) GetAllOrders() []*ent.Order {
// 	orders, err := d.client.Order.Query().All(context.Background())
// 	if err != nil {
// 		fmt.Printf("failed getting all orders: %v\n", err)
// 	}
// 	fmt.Printf("Orders: %+v\n", spew.Sdump(orders))
// 	return orders
// }
// func (d *Database) GetAllOrderItems() []*ent.OrderItems {
// 	orderItems, err := d.client.OrderItems.Query().All(context.Background())
// 	if err != nil {
// 		fmt.Printf("failed getting all orderItems: %v\n", err)
// 	}
// 	fmt.Printf("OrderItems: %+v\n", spew.Sdump(orderItems))
// 	return orderItems
// }
