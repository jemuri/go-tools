package estool

import (
	"context"
	"github.com/jemuri/go-tools/config"
	"sync"

	"github.com/pkg/errors"

	"github.com/jemuri/go-tools/strtools"

	"github.com/jemuri/go-tools/log"
	"github.com/olivere/elastic"
)

var esClient *elastic.Client

//var host = "http://127.0.0.1:9200/"
var host string
var esLock sync.Mutex
var onceDo sync.Once
var hostDo sync.Once

// EsClient EsClient
type EsClient struct {
	client *elastic.Client
}


// GetEsClient GetEsClient
func GetEsClient() *EsClient {
	ctx := log.LogInit(context.Background(), map[string]interface{}{})

	if host == "" {
		hostDo.Do(func() {
			host = config.CertainString("elasticsearch/host")
		})
	}

	if esClient == nil {
		esLock.Lock()
		if esClient == nil {
			var errNew error
			esClient, errNew = elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(host))
			if errNew != nil {
				log.Errorf(ctx, "es NewEsClient err: %v", errNew)
				panic(errNew)
			}

			info, code, err := esClient.Ping(host).Do(ctx)
			if err != nil {
				log.Errorf(ctx, "es Ping err: %v", err)
				panic(err)
			}

			log.Infof(ctx, "es code: %d, info: %+v", code, info)

			version, err := esClient.ElasticsearchVersion(host)
			if err != nil {
				log.Errorf(ctx, "es esVersion err: %v", err)
				panic(err)
			}
			log.Infof(ctx, "es version %s", version)
		}
		esLock.Unlock()
	}

	return &EsClient{esClient}
}

// GetClient GetClient
func GetClient() *EsClient {
	ctx := log.LogInit(context.Background(), map[string]interface{}{})

	if host == "" {
		hostDo.Do(func() {
			host = config.CertainString("elasticsearch/host")
		})
	}

	client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(host))
	if err != nil {
		log.Errorf(ctx, "es GetClient err: %v", err)
		panic(err)
	}
	onceDo.Do(func() {
		info, code, err := esClient.Ping(host).Do(ctx)
		if err != nil {
			log.Errorf(ctx, "es Ping err: %v", err)
			panic(err)
		}

		log.Infof(ctx, "es code: %d, info: %+v", code, info)

		version, err := esClient.ElasticsearchVersion(host)
		if err != nil {
			log.Errorf(ctx, "es esVersion err: %v", err)
			panic(err)
		}
		log.Infof(ctx, "es version %s", version)
	})

	return &EsClient{client}
}

// CreateIndex CreateIndex
func (e *EsClient) CreateIndex(ctx context.Context, index string) error {
	// Create an index
	_, err := e.client.CreateIndex(index).Do(ctx)
	if err != nil {
		// Handle error
		return err
	}
	return nil
}

// Create 创建
func (e *EsClient) Create(ctx context.Context, index, typ, id, refresh string, bodyJson interface{}) (string, error) {

	//使用结构体
	res, err := e.client.Index().
		Index(index).
		Type(typ).
		Id(id).
		BodyJson(bodyJson).
		Refresh(refresh).
		Do(ctx)
	if err != nil {
		log.Errorf(ctx, "es create err: %v", err)
		return "", err
	}
	log.Infof(ctx, "es Create successfully with id: %s", res.Id)

	return res.Id, nil
}

// Delete 删除
func (e *EsClient) Delete(ctx context.Context, index, typ, id string) error {

	res, err := e.client.Delete().Index(index).
		Type(typ).
		Id(id).
		Do(ctx)
	log.Infof(ctx, "es Delete result: %+v, err: %v", res.Result, err)
	return nil
}

// Delete 删除
func (e *EsClient) DeleteType(ctx context.Context, index, typ string) error {

	res, err := e.client.Delete().Index(index).
		//Type(typ).
		Do(ctx)

	log.Infof(ctx,"[删除index: %s,type: %s执行结果]: %v",index,typ,res)
	return err
}

// Update 修改
func (e *EsClient) Update(ctx context.Context, index, typ, id string, m map[string]interface{}) error {
	res, err := e.client.Update().
		Index(index).
		Type(typ).
		Id(id).
		Doc(m).
		Do(ctx)
	log.Infof(ctx, "es Update res: %+v err: %v", res.Result, err)
	return err
}

// Get 查找
func (e *EsClient) Get(ctx context.Context, index, typ, id string) (string, error) {
	//通过id查找
	res, err := e.client.Get().
		Index(index).
		Type(typ).
		Id(id).
		Do(ctx)
	if err != nil {
		log.Errorf(ctx, "es Get err: %v", err)
		return "", err
	}
	if !res.Found {
		return "", nil
	}

	log.Infof(ctx, "es Get document %s ", strtools.ToString(res))
	return strtools.ToString(res), nil
}

// List 简单分页
func (e *EsClient) PageSearch(ctx context.Context, index, typ, name, q string, size, pageNo int, pretty bool) (interface{}, error) {
	if size < 0 || pageNo < 1 {
		log.Warnf(ctx, "es list invalid parameters!")
		return nil, errors.New("es list invalid parameters!")
	}
	/**
	/条件查询
	    //年龄大于30岁的
	    boolQ := elastic.NewBoolQuery()
	    boolQ.Must(elastic.NewMatchQuery("last_name", "smith"))
	    boolQ.Filter(elastic.NewRangeQuery("age").Gt(30))
	 */
	// NewMatchPhraseQuery 短语搜索 搜索name字段中有 q
	mpQuery := elastic.NewMatchPhraseQuery(name, q)
	res, err := e.client.Search(). // search in index "tweets"
		Index(index).
		Type(typ).
		Query(mpQuery). // specify the query
		From((pageNo - 1) * size).Size(size). // take documents 0-9
		Pretty(pretty). // pretty print request and response JSON
		Do(ctx)

	return res, err
}

// List 简单分页
func (e *EsClient) PageSortSearch(ctx context.Context, index, typ, name, q, sortField string, size, pageNo int, sortAsc, pretty bool) (interface{}, error) {
	if size < 0 || pageNo < 1 {
		log.Warnf(ctx, "es list invalid parameters!")
		return nil, errors.New("es list invalid parameters!")
	}
	mpQuery := elastic.NewMatchPhraseQuery(name, q)
	res, err := e.client.Search(). // search in index "tweets"
		Index(index).
		Type(typ).
		Query(mpQuery). // specify the query
		Sort(sortField, sortAsc). // sort by "user" field, ascending
		From((pageNo - 1) * size).Size(size). // take documents 0-9
		Pretty(pretty). // pretty print request and response JSON
		Do(ctx)

	return res, err
}


func(e *EsClient) NewBulk() *elastic.BulkService {
	return e.client.Bulk()
}


// Create 批量创建
func (e *EsClient) BulkCreateReq(ctx context.Context, index, typ, id string, bodyJson interface{}) *elastic.BulkIndexRequest {

	res := elastic.NewBulkIndexRequest().
		Index(index).
		Type(typ).
		Id(id).
		Doc(bodyJson)

	return res
}
