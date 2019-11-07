package images

import (
	"context"

	"github.com/aikizoku/rabbitgo/appengine/default/src/lib/cloudpubsub"
	"github.com/aikizoku/rabbitgo/appengine/default/src/lib/log"
)

// Client ... クライアント
type Client struct {
	psCli     *cloudpubsub.Client
	topicName string
}

// SendConvertRequest ... 画像変換リクエストを送信する
func (c *Client) SendConvertRequest(
	ctx context.Context,
	sourceID string,
	sources []*Object,
	dstIsArray bool,
	dstFilePath string,
	dstDocPath string,
	dstFieldName string) error {
	srcURLs := []string{}
	for _, source := range sources {
		if source == nil || source.URL == "" {
			continue
		}
		srcURLs = append(srcURLs, source.URL)
	}
	if len(srcURLs) == 0 {
		return nil
	}
	src := &ConvRequest{
		SourceID:     sourceID,
		SourceURLs:   srcURLs,
		DstIsArray:   dstIsArray,
		DstFilePath:  dstFilePath,
		DstDocPath:   dstDocPath,
		DstFieldName: dstFieldName,
	}
	err := c.psCli.Publish(ctx, c.topicName, src)
	if err != nil {
		log.Errorm(ctx, "c.psCli.Publish", err)
		return err
	}
	return nil
}

// NewClient ... クライアントを作成する
func NewClient(psCli *cloudpubsub.Client, topicName string) *Client {
	return &Client{
		psCli:     psCli,
		topicName: topicName,
	}
}
