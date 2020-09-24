package main

import (
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	"github.com/cloudfoundry/bytefmt"
	"github.com/soichisumi/go-util/logger"
	"go.uber.org/zap"
	"google.golang.org/api/iterator"
	"sort"
	"strings"
)

func main() {
	ctx := context.Background()
	c, err := storage.NewClient(ctx) // use default application service account
	if err != nil {
		logger.Fatal("", zap.Error(err))
	}
	b := c.Bucket("BUCKET_NAME")
	it := b.Objects(ctx, &storage.Query{
		Prefix: "FOLDER_NAME", // objectName = {folder1}/{folder2}/.../{fileName}
	})

	sizeMap := make(map[string]int64)

	logger.Info("start listing objects")
	for {
		o, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			logger.Error("", zap.Error(err))
			break
		}
		splitted := strings.Split(o.Name, "/")
		name := splitted[len(splitted)-1]

		var tableName string
		switch {
		case strings.Contains(name, "-manifest"):
			tableName = strings.Split(name, "-manifest")[0]
		case strings.Contains(name, ".avro"):
			tableName = strings.Split(name, ".avro")[0]
		}

		sizeMap[tableName] = sizeMap[tableName] + o.Size

		logger.Info("obj", zap.String("name", o.Name), zap.Int64("size", o.Size))
	}
	logger.Info("object listing done.")

	logger.Info("result:")

	keys := make([]string, 0, len(sizeMap))
	for k := range sizeMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var total int64
	for _, k := range keys {
		total += sizeMap[k]
	}

	for _, k := range keys {
		fmt.Printf("|%s|%s Byte|%f%%|\n", k, bytefmt.ByteSize(uint64(sizeMap[k])), float64(sizeMap[k])/float64(total) * 100)
	}
	logger.Info("done", zap.String("totalSize", bytefmt.ByteSize(uint64(total))))
}