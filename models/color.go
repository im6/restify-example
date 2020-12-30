package models

import (
	"time"
	"strings"
	"context"
	"github.com/im6/vp3/store"
	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

// Color is definition
type Color struct {
	Id         string     `json:"id"`
	Star       int64      `json:"like"`
	Colors     string     `json:"color"`
	Author     string     `json:"username"`
	Hidden     bool       `json:"hidden"`
	Created    time.Time  `json:"createdate"`
}

// Get color collection from DB
func GetColors(ctx context.Context, queryType string) ([]Color, error) {
	client := store.CreateSqlClient(ctx)
	var iter *firestore.DocumentIterator
	
	switch queryType {
	case "latest":
		iter = client.Collection("colors").Documents(ctx)
	case "popular":
		iter = client.Collection("colors").OrderBy("star", firestore.Desc).Documents(ctx)
	}

	var colors []Color
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		colorDict := doc.Data()
		colorInterface := colorDict["v"].([]interface{})
		colorStrArr := make([]string, len(colorInterface))
		for i, v := range colorInterface {
			colorStrArr[i] = v.(string)
		}
		colorStr := strings.Join(colorStrArr, "#")
    colorIns := Color{
			Id: doc.Ref.ID,
			Star: colorDict["star"].(int64),
			Author: colorDict["author"].(string),
			Hidden: colorDict["hidden"].(bool),
			Created: colorDict["created"].(time.Time),
			Colors: colorStr,
		}
		colors = append(colors, colorIns)
	}

	return colors, nil
}


func GetOneColor(ctx context.Context, docId string) (*Color, error) {
	client := store.CreateSqlClient(ctx)
	doc, err := client.Collection("colors").Doc(docId).Get(ctx)
	if err != nil {
		return nil, err
	}
	colorDict := doc.Data()
	colorInterface := colorDict["v"].([]interface{})
	colorStrArr := make([]string, len(colorInterface))
	for i, v := range colorInterface {
		colorStrArr[i] = v.(string)
	}
	colorStr := strings.Join(colorStrArr, "#")
	colorIns :=  &Color{
		Id: docId,
		Star: colorDict["star"].(int64),
		Author: colorDict["author"].(string),
		Hidden: colorDict["hidden"].(bool),
		Created: colorDict["created"].(time.Time),
		Colors: colorStr,
	}

	return colorIns, nil
}
