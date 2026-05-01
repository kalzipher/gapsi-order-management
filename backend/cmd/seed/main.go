package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/github/gapsi-order-management-dashboard/backend/internal/config"
	"github.com/github/gapsi-order-management-dashboard/backend/internal/database"
	"github.com/github/gapsi-order-management-dashboard/backend/internal/orders"
)

const batchSize = 1000

func main() {
	cfg := config.Load()

	db, err := database.NewPostgresDB(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("could not connect to database: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("could not get sql db: %v", err)
	}
	defer sqlDB.Close()

	file, err := os.Open(cfg.CSVPath)
	if err != nil {
		log.Fatalf("could not open csv file %q: %v", cfg.CSVPath, err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	reader.TrimLeadingSpace = true

	headers, err := reader.Read()
	if err != nil {
		log.Fatalf("could not read csv headers: %v", err)
	}

	headerIndex := buildHeaderIndex(headers)

	batch := make([]orders.OrderEntity, 0, batchSize)

	var totalRead int
	var totalSkipped int

	for {
		record, err := reader.Read()
		if errors.Is(err, io.EOF) {
			break
		}

		if err != nil {
			log.Printf("skipping invalid csv row: %v", err)
			totalSkipped++
			continue
		}

		totalRead++

		entity, err := recordToOrderEntity(record, headerIndex)
		if err != nil {
			log.Printf("skipping row %d: %v", totalRead, err)
			totalSkipped++
			continue
		}

		batch = append(batch, entity)

		if len(batch) >= batchSize {
			if err := insertBatch(db, batch); err != nil {
				log.Fatalf("could not insert batch: %v", err)
			}

			log.Printf("processed rows: %d", totalRead)
			batch = batch[:0]
		}
	}

	if len(batch) > 0 {
		if err := insertBatch(db, batch); err != nil {
			log.Fatalf("could not insert final batch: %v", err)
		}
	}

	log.Printf("seed finished")
	log.Printf("rows read: %d", totalRead)
	log.Printf("rows skipped: %d", totalSkipped)
}

func buildHeaderIndex(headers []string) map[string]int {
	index := make(map[string]int, len(headers))

	for i, header := range headers {
		index[strings.TrimSpace(header)] = i
	}

	return index
}

func insertBatch(db *gorm.DB, batch []orders.OrderEntity) error {
	return db.
		Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "id"}},
			DoNothing: true,
		}).
		Create(&batch).
		Error
}

func recordToOrderEntity(record []string, headerIndex map[string]int) (orders.OrderEntity, error) {
	id := get(record, headerIndex, "__id__")
	if id == "" {
		return orders.OrderEntity{}, fmt.Errorf("missing __id__")
	}

	return orders.OrderEntity{
		ID:              id,
		Canal:           get(record, headerIndex, "canal"),
		Cantidad:        parseOptionalInt(get(record, headerIndex, "cantidad")),
		Company:         get(record, headerIndex, "company"),
		CP:              get(record, headerIndex, "cp"),
		CreatedAt:       parseOptionalTimestamp(get(record, headerIndex, "createdAt"), false),
		DaysToDelivery:  parseOptionalInt(get(record, headerIndex, "daysToDelivery")),
		ErrorCode:       get(record, headerIndex, "error"),
		ErrorMessage:    get(record, headerIndex, "errorMessage"),
		FechaCompra:     parseOptionalTimestamp(get(record, headerIndex, "fechaCompra"), true),
		FechaEstimada:   get(record, headerIndex, "fechaEstimada"),
		FulfillmentType: get(record, headerIndex, "fulfillmentType"),
		IsFlash:         parseOptionalBool(get(record, headerIndex, "isFlash")),
		IsMarketplace:   parseOptionalBool(get(record, headerIndex, "isMarketplace")),
		NoPedido:        get(record, headerIndex, "noPedido"),
		Plan:            get(record, headerIndex, "plan"),
		ProductType:     get(record, headerIndex, "productType"),
		SKU:             get(record, headerIndex, "sku"),
		StoreSelected:   get(record, headerIndex, "storeSelected"),
		TipoPago:        get(record, headerIndex, "tipoPago"),
		EDD1:            get(record, headerIndex, "edd1"),
		EDD2:            get(record, headerIndex, "edd2"),
	}, nil
}

func get(record []string, headerIndex map[string]int, column string) string {
	index, ok := headerIndex[column]
	if !ok || index < 0 || index >= len(record) {
		return ""
	}

	return normalizeString(record[index])
}

func normalizeString(value string) string {
	value = strings.TrimSpace(value)

	if value == "" {
		return ""
	}

	lower := strings.ToLower(value)

	switch lower {
	case "null", "n/a", "na", "undefined":
		return ""
	default:
		return value
	}
}

func parseOptionalInt(value string) *int {
	value = normalizeString(value)
	if value == "" {
		return nil
	}

	parsed, err := strconv.Atoi(value)
	if err != nil {
		return nil
	}

	return &parsed
}

func parseOptionalBool(value string) *bool {
	value = normalizeString(value)
	if value == "" {
		return nil
	}

	switch strings.ToLower(value) {
	case "true", "1", "yes", "y":
		parsed := true
		return &parsed
	case "false", "0", "no", "n":
		parsed := false
		return &parsed
	default:
		return nil
	}
}

func parseOptionalTimestamp(value string, nullIfUnixEpoch bool) *time.Time {
	value = normalizeString(value)
	if value == "" {
		return nil
	}

	value = strings.TrimPrefix(value, "__Timestamp__")

	parsed, err := time.Parse(time.RFC3339Nano, value)
	if err != nil {
		return nil
	}

	if nullIfUnixEpoch && parsed.Year() == 1970 && parsed.Month() == time.January && parsed.Day() == 1 {
		return nil
	}

	return &parsed
}
