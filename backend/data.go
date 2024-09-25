package main

import "time"

var orders = []Order{
	{
		OrderUID:          "1",
		TrackNumber:       "123456",
		Entry:             "entry1",
		Locale:            "ru",
		InternalSignature: "",
		CustomerID:        "ITMO_1",
		DeliveryService:   "service1",
		ShardKey:          "shard1",
		SmID:              1,
		DateCreated:       time.Now(),
		OofShard:          "",
		Delivery:          Delivery{},
		Payment:           Payment{},
		Items:             []Item{},
	},
	{
		OrderUID:          "2",
		TrackNumber:       "654321",
		Entry:             "entry2",
		Locale:            "ru",
		InternalSignature: "",
		CustomerID:        "ITMO_2",
		DeliveryService:   "service2",
		ShardKey:          "shard2",
		SmID:              2,
		DateCreated:       time.Now(),
		OofShard:          "",
		Delivery:          Delivery{},
		Payment:           Payment{},
		Items:             []Item{},
	},
}
