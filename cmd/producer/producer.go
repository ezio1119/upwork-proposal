package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"upwork-proposal/config"

	"github.com/segmentio/kafka-go"
)

func main() {

	w := &kafka.Writer{
		Addr:     kafka.TCP(config.KafkaEndpoint),
		Topic:    config.KafkaTopic,
		Balancer: &kafka.LeastBytes{},
	}
	defer w.Close()

	texts, err := readTextsFromCSV("./text.csv")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("will produce: %s \n", texts)

	msgs := make([]kafka.Message, len(texts))
	for i := range texts {
		msgs[i].Key = []byte(strconv.Itoa(i))
		msgs[i].Value = []byte(texts[i])
	}

	if err := w.WriteMessages(context.Background(), msgs...); err != nil {
		log.Fatal("failed to write messages:", err)
	}

}

func readTextsFromCSV(path string) ([]string, error) {
	csvFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		return nil, err
	}

	texts := make([]string, 0, len(csvLines))
	for _, l := range csvLines {
		texts = append(texts, l...)
	}

	return texts, nil
}
