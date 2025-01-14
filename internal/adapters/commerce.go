package adapters

import (
	"fmt"
	"log"

	"github.com/Edmartt/loyalty-system/internal/core/domain"
	ports "github.com/Edmartt/loyalty-system/internal/core/ports/database"
)

type CommerceRepository struct {
	dbConnection ports.IConnection
}

func NewCommerceRepository(dbConnection ports.IConnection) *CommerceRepository {
	return &CommerceRepository{
		dbConnection: dbConnection,
	}
}

func (c *CommerceRepository) ReadCommerceCampaign(id string) ([]domain.Campaign, error) {

	dbConnection, err := c.dbConnection.GetConnection()

	if err != nil {
		return nil, fmt.Errorf("error connecting to db: %v", err)
	}

	var result []domain.Campaign

	err = dbConnection.Select(&result, "SELECT c.id AS campaign_id, c.campaign_name, c.campaign_multiplier, c.campaign_percent_bonus, c.start_date, c.end_date, co.id AS commerce_id, co.name AS commerce_name, co.points_x_buy, co.value_x_point FROM campaign c JOIN commerce co ON c.commerce_id = co.id WHERE co.id = $1", id)

	if err != nil {
		log.Println(err.Error())
		return nil, fmt.Errorf("error fetching campaign data: %v", err)
	}

	return result, nil
}

func (c *CommerceRepository) SaveCommerce(commerce domain.Commerce) (*domain.Commerce, error) {
	dbConnection, err := c.dbConnection.GetConnection()

	if err != nil {
		return nil, fmt.Errorf("error connecting to DB: %v", err)
	}

	_, err = dbConnection.NamedExec("INSERT INTO commerce (id, name, points_x_buy, value_x_point) VALUES(:id, :name, :points_x_buy, :value_x_point)", commerce)

	if err != nil {
		return nil, fmt.Errorf("error saving data: %v", err)
	}

	return &commerce, nil
}

func (c *CommerceRepository) CreateCommerceCampaign(campaign domain.Campaign) (*domain.Campaign, error) {

	dbConnection, err := c.dbConnection.GetConnection()

	if err != nil {
		return nil, fmt.Errorf("error connecting to DB: %v", err)
	}

	_, err = dbConnection.NamedExec("INSERT INTO campaign(commerce_id, branch_id, campaign_name, campaign_type, campaign_multiplier, campaign_percent_bonus, start_date, end_date) VALUES(:commerce_id, :branch_id, :campaign_name, :campaign_type, :campaign_multiplier,:campaign_percent_bonus, :start_date, :end_date)", &campaign)

	if err != nil {
		return nil, fmt.Errorf("error saving campaign data: %v", err)
	}

	return &campaign, nil
}

func (c *CommerceRepository) ReadPointsXBuy(id string) (*domain.Commerce, error) {
	dbConnection, err := c.dbConnection.GetConnection()

	if err != nil {
		return nil, fmt.Errorf("error connecting to db: %v", err)
	}

	var commerce domain.Commerce

	err = dbConnection.Get(&commerce, "SELECT name, points_x_buy FROM commerce WHERE id=?", id)

	if err != nil {
		return nil, fmt.Errorf("error fetching commerce points: %v", err)
	}

	return &commerce, nil
}

func (c *CommerceRepository) SavePointsXBuy(id string, points int32) (string, error) {
	dbConnection, err := c.dbConnection.GetConnection()

	if err != nil {
		return "", fmt.Errorf("error connecting to db: %v", err)
	}

	_, err = dbConnection.Exec("UPDATE commerce SET points_x_buy = $1 WHERE id = $2", points, id)

	if err != nil {
		return "", fmt.Errorf("error updating commerce points: %v", err)
	}

	message := fmt.Sprintf("current commerce with ID %s given points updated to %d", id, points)

	return message, nil
}

func (c *CommerceRepository) ReadValueXPoint(id string) (*domain.Commerce, error) {
	dbConnection, err := c.dbConnection.GetConnection()

	if err != nil {
		return nil, fmt.Errorf("error connecting to db: %v", err)
	}

	var commerce domain.Commerce

	err = dbConnection.Get(&commerce, "SELECT value_x_point FROM commerce WHERE id=?", id)

	if err != nil {
		return nil, fmt.Errorf("error fetching commerce value per point: %v", err)
	}

	return &commerce, nil
}

func (c *CommerceRepository) SaveValueXPoint(id string, value int64) (string, error) {

	dbConnection, err := c.dbConnection.GetConnection()

	if err != nil {
		return "", fmt.Errorf("error connecting to db: %v", err)
	}

	_, err = dbConnection.Exec("UPDATE commerce SET value_x_point = $1 WHERE id = $2", value, id)

	if err != nil {
		return "", fmt.Errorf("error updating commerce points: %v", err)
	}

	message := fmt.Sprintf("current commerce with ID %s given value updated to %d", id, value)

	return message, nil
}
