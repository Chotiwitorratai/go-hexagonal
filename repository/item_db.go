package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"go-hexagonal/model"
	"strconv"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type itemRepositoryDB struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewItemRepositoryDB(db *gorm.DB, redisClient *redis.Client) itemRepositoryDB {
	return itemRepositoryDB{db, redisClient}
}

func (r itemRepositoryDB) CreateItem(ci *model.CreateItem) (*model.Item, error) {
	item := &model.Item{}
	item.Name = ci.Name
	item.Price = ci.Price
	item.Stock = ci.Stock
	err := r.db.Create(&item).Error
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (r itemRepositoryDB) GetAllItem() (items []model.Item, err error) {
	key := "repository::items"

	//Redis GET
	if resp, err := r.redis.Get(context.Background(), key).Result(); err == nil {
		if json.Unmarshal([]byte(resp), &items) == nil {
			fmt.Println("redis")
			return items, nil
		}
	}

	//Repository
	err = r.db.Find(&items).Error
	if err != nil {
		return nil, err
	}

	//Redis SET
	data, err := json.Marshal(items)
	if err != nil {
		return nil, err
	}
	err = r.redis.Set(context.Background(), key, string(data), 0).Err()
	if err != nil {
		return nil, err
	}

	fmt.Println("database")

	return items, nil
}

func (r itemRepositoryDB) GetItemByID(item_id uint) (*model.Item, error) {
	item := &model.Item{}
	key := "repository::item"

	//Redis GET
	if resp, err := r.redis.Get(context.Background(), key).Result(); err == nil {
		if json.Unmarshal([]byte(resp), &item) == nil {
			fmt.Println("redis")
			return item, nil
		}
	}

	//Repository
	err := r.db.Find(item, item_id).Error
	if err != nil {
		return nil, err
	}

	//Redis SET
	if data, err := json.Marshal(item); err == nil {
		r.redis.Set(context.Background(), key, string(data), 0)
	}

	return item, nil
}

func (r itemRepositoryDB) UpdateItem(item_id uint, qty int) (*model.Item, error) {
	item := &model.Item{}
	err := r.db.Model(item).Where("item_id = ?", item_id).Update("Stock", qty).Error
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (r itemRepositoryDB) UpdateItemDetail(update *model.UpdateItem, item_id uint) (*model.Item, error) {
	item := &model.Item{
		ItemID: item_id,
	}
	item.Name = update.Name
	item.Price = update.Price
	item.Stock = update.Stock
	tx := r.db.Model(&model.Item{ItemID: item_id}).Updates(item)
	if tx.Error != nil {
		return nil, tx.Error
	}
	// json, err := json.Marshal(tx)
	// if err != nil {
	// 	return nil, err
	// }
	// fmt.Println("tx :", string(json))
	return item, nil
}

func (r itemRepositoryDB) GetStockByID(item_id uint) (int, error) {

	item := &model.Item{ItemID: item_id}
	key := "repository::item_id"

	//Redis GET
	if resp, err := r.redis.Get(context.Background(), key).Result(); err == nil {
		intVar, err := strconv.Atoi(resp)
		if err != nil {
			return 0, err
		}
		return intVar, nil
	}

	//Repository
	err := r.db.First(item, item_id).Error
	if err != nil {
		return 0, err
	}

	//Redis SET
	r.redis.Set(context.Background(), key, item.Stock, 0)

	return item.Stock, nil
}
