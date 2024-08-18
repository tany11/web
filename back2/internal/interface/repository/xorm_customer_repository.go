package repository

import (
	"back2/internal/domain/entity"
	"back2/internal/domain/repository"
	"log"

	"github.com/go-xorm/xorm"
)

type XormCustomerRepository struct {
	engine *xorm.Engine
}

func NewXormCustomerOrderRepository(engine *xorm.Engine) repository.CustomerRepository {
	return &XormCustomerRepository{engine: engine}
}

func (r *XormCustomerRepository) Create(customer *entity.Customer) error {
	_, err := r.engine.Insert(customer)
	return err
}

func (r *XormCustomerRepository) GetByID(id int64) (*entity.Customer, error) {
	order := new(entity.Customer)
	has, err := r.engine.ID(id).Get(order)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return order, nil
}

func (r *XormCustomerRepository) List(groupID, offset, limit int) ([]*entity.Customer, error) {
	orders := make([]*entity.Customer, 0)
	err := r.engine.Where("group_id = ?", groupID).Limit(limit, offset).Find(&orders)
	return orders, err
}

func (r *XormCustomerRepository) ListByCustomerID(customerID int) ([]*entity.Customer, error) {
	orders := make([]*entity.Customer, 0)
	err := r.engine.Where("customer_id = ?", customerID).Find(&orders)
	return orders, err
}

func (r *XormCustomerRepository) Update(order *entity.Customer) error {
	_, err := r.engine.ID(order.ID).Update(order)
	return err
}

func (r *XormCustomerRepository) Delete(id int64) error {
	_, err := r.engine.ID(id).Delete(new(entity.Customer))
	return err
}

func (r *XormCustomerRepository) GetByPhoneNumber(phoneNumber string) (*entity.Customer, error) {
	order := new(entity.Customer)
	has, err := r.engine.Where("phone_number = ?", phoneNumber).Get(order)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return order, nil
}

func (r *XormCustomerRepository) GetSearchList(customerSearch entity.CustomerSearch) ([]*entity.CustomerOrder, error) {
	customers := make([]*entity.CustomerOrder, 0)

	query := `
		SELECT DISTINCT c.* 
		FROM customer c
		JOIN orders o ON c.i_d = o.customer_i_d
		WHERE 1=1
	`
	var params []interface{}

	log.Printf("検索クエリの構築開始")

	if customerSearch.PhoneLast4 != "" {
		query += " AND c.phone_number LIKE ?"
		params = append(params, "%"+customerSearch.PhoneLast4)
		log.Printf("電話番号の条件追加: %s", customerSearch.PhoneLast4)
	}

	if customerSearch.CastID != "" {
		query += " AND o.actual_model = ?"
		params = append(params, customerSearch.CastID)
		log.Printf("キャストIDの条件追加: %s", customerSearch.CastID)
	}

	if customerSearch.StoreID != "" {
		query += " AND o.store_i_d = ?"
		params = append(params, customerSearch.StoreID)
		log.Printf("店舗IDの条件追加: %s", customerSearch.StoreID)
	}

	if customerSearch.CreatedFrom != "" {
		query += " AND o.created_at >= ?"
		params = append(params, customerSearch.CreatedFrom)
		log.Printf("作成日From条件追加: %s", customerSearch.CreatedFrom)
	}

	if customerSearch.CreatedTo != "" {
		query += " AND o.created_at <= ?"
		params = append(params, customerSearch.CreatedTo)
		log.Printf("作成日To条件追加: %s", customerSearch.CreatedTo)
	}

	err := r.engine.SQL(query, params...).Find(&customers)
	if err != nil {
		log.Printf("クエリ実行エラー: %v", err)
		return nil, err
	}

	log.Printf("検索結果: %v", customers)

	log.Printf("検索結果: %d件の顧客が見つかりました", len(customers))
	return customers, err
}
