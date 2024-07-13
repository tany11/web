package usecase

import (
	"back2/internal/domain/entity"
	"back2/internal/domain/errors"
	"back2/internal/domain/repository"
	"fmt"
	"log"
	"time"
)

type OrderUseCase struct {
	orderRepo    repository.OrderRepository
	customerRepo repository.CustomerRepository
}

func NewOrderUseCase(orderRepo repository.OrderRepository, customerRepo repository.CustomerRepository) *OrderUseCase {
	return &OrderUseCase{
		orderRepo:    orderRepo,
		customerRepo: customerRepo,
	}
}

func (uc *OrderUseCase) Create(order *entity.Orders) error {
	// 顧客を電話番号で検索
	customer, err := uc.customerRepo.GetByPhoneNumber(string(order.PhoneNumber))
	if err != nil && err != errors.ErrCustomerNotFound {
		return fmt.Errorf("顧客検索中にエラーが発生しました: %w", err)
	}

	// 顧客が見つからない場合、新規作成
	if customer == nil {
		log.Printf("電話番号が見つからなかったため、新規顧客を作成します")
		newCustomer := &entity.Customer{
			PhoneNumber:  order.PhoneNumber,
			CustomerName: order.CustomerName,
			Address:      order.Address,
			GroupID:      1,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
			// 他の必要な顧客情報をここで設定
		}
		err = uc.customerRepo.Create(newCustomer)
		log.Printf("顧客を作成しました")
		if err != nil {
			return fmt.Errorf("新規顧客の作成に失敗しました: %w", err)
		}
		customer = newCustomer
	} else {
		log.Printf("既存の顧客が見つかりました")
	}

	// 顧客IDを注文データに設定
	order.CustomerID = customer.ID
	order.DriverID = "1"
	order.CreatedAt = time.Now()
	order.UpdatedAt = time.Now()

	// 注文を作成
	err = uc.orderRepo.Create(order)
	if err != nil {
		log.Printf("注文の作成に失敗しました: %v", err)
		return fmt.Errorf("注文の作成に失敗しました: %w", err)
	}

	log.Printf("注文を作成しました: ID=%d", order.ID)
	return nil
}

func (uc *OrderUseCase) List(groupID, page, pageSize int) ([]*entity.Orders, error) {
	offset := (page - 1) * pageSize
	orders, err := uc.orderRepo.List(groupID, offset, pageSize)
	if err != nil {
		log.Printf("注文リストの取得中にエラーが発生しました: %v", err)
		return nil, fmt.Errorf("注文リストの取得に失敗しました: %w", err)
	}
	return orders, nil
}

func (uc *OrderUseCase) GetByID(id int64) (*entity.Orders, error) {
	return uc.orderRepo.GetByID(id)
}

// func (uc *OrderUseCase) List(groupID int, page, pageSize int) ([]*entity.Order, error) {
// 	offset := (page - 1) * pageSize
// 	return uc.orderRepo.List(groupID, offset, pageSize)
// }

func (uc *OrderUseCase) Update(order *entity.Orders) error {
	return uc.orderRepo.Update(order)
}

func (uc *OrderUseCase) Delete(id int64) error {
	return uc.orderRepo.Delete(id)
}

func (uc *OrderUseCase) GetByCustomerID(customerID int) ([]*entity.Orders, error) {
	return uc.orderRepo.ListByCustomerID(customerID)
}
