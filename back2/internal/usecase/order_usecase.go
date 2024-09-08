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
	storeRepo    repository.StoreRepository // 追加
}

func NewOrderUseCase(orderRepo repository.OrderRepository, customerRepo repository.CustomerRepository, storeRepo repository.StoreRepository) *OrderUseCase {
	return &OrderUseCase{
		orderRepo:    orderRepo,
		customerRepo: customerRepo,
		storeRepo:    storeRepo, // 追加
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
			City1:        order.City,
			Address1:     order.Address,
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
		// 既存の顧客のUpdatedAtを更新
		customer.UpdatedAt = time.Now()
		err = uc.customerRepo.Update(customer)
		if err != nil {
			return fmt.Errorf("顧客情報の更新に失敗しました: %w", err)
		}
	}

	// ストアをIDで検索し、DummyStoreFlgを取得
	store, err := uc.storeRepo.GetByID(int64(order.StoreID))
	if err != nil {
		return fmt.Errorf("ストア検索中にエラーが発生しました: %w", err)
	}
	order.DummyStoreFlg = store.DummyStoreFlg

	// 顧客IDを注文データに設定
	order.CustomerID = customer.ID
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

func (uc *OrderUseCase) ListReserved(groupID, page, pageSize int) ([]*entity.Orders, error) {
	offset := (page - 1) * pageSize
	orders, err := uc.orderRepo.ListReserved(groupID, offset, pageSize)
	if err != nil {
		log.Printf("予約リストの取得中にエラーが発生しました: %v", err)
		return nil, fmt.Errorf("予約リストの取得に失敗しました: %w", err)
	}
	return orders, nil
}

func (uc *OrderUseCase) GetByID(id int64) (*entity.Orders, error) {
	return uc.orderRepo.GetByID(id)
}

func (uc *OrderUseCase) Update(order *entity.Orders) error {
	// 顧客IDで顧客を取得
	customer, err := uc.customerRepo.GetByID(order.CustomerID)
	if err != nil {
		return fmt.Errorf("顧客取得中にエラーが発生しました: %w", err)
	}

	// 電話番号が異なる場合、上書き
	if customer.PhoneNumber != order.PhoneNumber {
		customer.PhoneNumber = order.PhoneNumber
		customer.UpdatedAt = time.Now()
		err = uc.customerRepo.Update(customer)
		if err != nil {
			return fmt.Errorf("顧客情報の更新に失敗しました: %w", err)
		}
	}

	// ストアをIDで検索し、DummyStoreFlgを取得
	store, err := uc.storeRepo.GetByID(int64(order.StoreID))
	if err != nil {
		return fmt.Errorf("ストア検索中にエラーが発生しました: %w", err)
	}
	order.DummyStoreFlg = store.DummyStoreFlg

	// 注文を更新
	return uc.orderRepo.Update(order)
}

func (uc *OrderUseCase) UpdateSchedule(order *entity.Orders) error {
	return uc.orderRepo.UpdateSchedule(order)
}

func (uc *OrderUseCase) Delete(id int64) error {
	return uc.orderRepo.Delete(id)
}

func (uc *OrderUseCase) GetByCustomerID(customerID int) ([]*entity.Orders, error) {
	return uc.orderRepo.ListByCustomerID(customerID)
}

func (uc *OrderUseCase) UpdateCompletionFlg(id int64) error {
	return uc.orderRepo.UpdateCompletionFlg(id)
}

func (uc *OrderUseCase) UpdateIsDeleted(id int64) error {
	return uc.orderRepo.UpdateIsDeleted(id)
}

func (uc *OrderUseCase) GetTotalPriceAndUseTime(customerID int) (int, int, error) {
	return uc.orderRepo.GetTotalPriceAndUseTime(customerID)
}

func (uc *OrderUseCase) ListSchedule(startDate, endDate string) ([]*entity.Orders, error) {
	return uc.orderRepo.ListSchedule(startDate, endDate)
}
