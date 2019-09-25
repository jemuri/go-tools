package estool

import (
	"context"
	"fmt"
	"github.com/jemuri/go-tools/config"
	"os"
	"reflect"
	"strconv"
	"testing"

	"github.com/olivere/elastic"
)

func TestMain(m *testing.M) {
	config.Init("../../conf/config.toml", "CONF_ENV")
	os.Exit(m.Run())
}
func TestEsClient_CreateIndex(t *testing.T) {
	type fields struct {
		client *elastic.Client
	}
	type args struct {
		ctx   context.Context
		index string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "case1",
			fields:  fields{},
			args:    args{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &EsClient{
				client: tt.fields.client,
			}
			if err := e.CreateIndex(tt.args.ctx, tt.args.index); (err != nil) != tt.wantErr {
				t.Errorf("EsClient.CreateIndex() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEsClient_Create(t *testing.T) {
	type fields struct {
		client *elastic.Client
	}
	type args struct {
		ctx      context.Context
		index    string
		typ      string
		id       string
		refresh  string
		bodyJson interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "case1",
			fields:  fields{},
			args:    args{},
			want:    "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &EsClient{
				client: tt.fields.client,
			}
			got, err := e.Create(tt.args.ctx, tt.args.index, tt.args.typ, tt.args.id, tt.args.refresh, tt.args.bodyJson)
			if (err != nil) != tt.wantErr {
				t.Errorf("EsClient.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("EsClient.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

/*
如果批量upsert，那么需要将代码

req := elastic.NewBulkIndexRequest().Index("twitter").Type("tweet").Id(strconv.Itoa(n)).Doc(tweet)
改成

req := elastic.NewBulkUpdateRequest().Index("twitter").Type("tweet").Id(strconv.Itoa(n)).Doc(tweet).DocAsUpsert(true)
*/
func Test_BULK(t *testing.T) {
	client, err := elastic.NewClient()
	if err != nil {
		fmt.Println("出错了", err)
	}

	n := 0
	for i := 0; i < 1000; i++ {
		bulkRequest := client.Bulk()
		for j := 0; j < 10000; j++ {
			n++
			tweet := struct{ User string }{User: "olivere"}
			req := elastic.NewBulkIndexRequest().Index("twitter").Type("tweet").Id(strconv.Itoa(n)).Doc(tweet)
			bulkRequest = bulkRequest.Add(req)
		}
		bulkResponse, err := bulkRequest.Do(context.Background())
		if err != nil {
			fmt.Println(err)
		}
		if bulkResponse != nil {

		}
		fmt.Println(i)
	}
}

func TestEsClient_DeleteType(t *testing.T) {
	type fields struct {
		client *elastic.Client
	}
	type args struct {
		ctx   context.Context
		index string
		typ   string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "case1",
			fields: fields{
				client: &elastic.Client{},
			},
			args: args{
				ctx:   context.Background(),
				index: "spider",
				typ:   "price",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, err := elastic.NewClient()
			if err != nil {
				fmt.Println("粗粗哦", err)
			}
			e := &EsClient{
				client: client,
			}

			if err := e.DeleteType(tt.args.ctx, tt.args.index, tt.args.typ); (err != nil) != tt.wantErr {
				t.Logf("EsClient.DeleteType() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEsClient_PageSearch(t *testing.T) {
	type fields struct {
		client *elastic.Client
	}
	type args struct {
		ctx    context.Context
		index  string
		typ    string
		name   string
		q      string
		size   int
		pageNo int
		pretty bool
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			name:   "case1",
			fields: fields{},
			args: args{
				ctx:    context.Background(),
				index:  "price",
				typ:    "tag",
				name:   "name",
				q:      "内存",
				size:   20,
				pageNo: 2,
				pretty: false,
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := GetEsClient()
			got, err := client.PageSearch(tt.args.ctx, tt.args.index, tt.args.typ, tt.args.name, tt.args.q, tt.args.size, tt.args.pageNo, tt.args.pretty)
			if (err != nil) != tt.wantErr {
				t.Logf("EsClient.PageSearch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Logf("EsClient.PageSearch() = %v, want %v", got, tt.want)
			}
		})
	}
}



func TestEsClient_NewBulk(t *testing.T) {



}