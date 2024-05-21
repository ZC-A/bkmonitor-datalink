package doris

import (
	"context"
	"database/sql"

	"github.com/TencentBlueKing/bkmonitor-datalink/pkg/unify-query/log"
	"github.com/abulo/ratel/v2/stores/doris"
)

var (
	db *sql.DB
)

type Service struct {
	ctx        context.Context
	cancelFunc context.CancelFunc
}

func (s *Service) Type() string {
	return "doris"
}

func (s *Service) Reload(ctx context.Context) {
	log.Infof(context.TODO(), "Closing current doris service...")
	s.Close()

	s.ctx, s.cancelFunc = context.WithCancel(ctx)

	defer func() {
		if r := recover(); r != nil {
			log.Errorf(context.TODO(), "failed to create new doris client") // 这个 doris pkg 在创建连接时如果出现错误会 panic 而非返回 err msg
		}
	}()
	// 更新上下文方法
	log.Infof(context.TODO(), "doris service context update success.")
	config := &doris.Config{
		Username:      Username,
		Password:      Password,
		Host:          Host,
		Port:          Port,
		DriverName:    DriverName,
		Charset:       CharSet,
		Local:         Local,
		MaxOpenConns:  MaxIdleConns,
		MaxIdleConns:  MaxIdleConns,
		MaxLifetime:   MaxLifetime,
		MaxIdleTime:   MaxIdleTime,
		DisableMetric: DisableMetric,
		DisableTrace:  DisableTrace,
	}
	client := doris.NewClient(config)
	db = client.DB
	if err := db.PingContext(s.ctx); err != nil {
		log.Errorf(context.TODO(), "fail to ping doris db, %s", err.Error())
		return
	}
	log.Infof(context.TODO(), "doris service start or reloaded sucessfully")

}

// Wait
func (s *Service) Wait() {
}

func (s *Service) Start(ctx context.Context) {
	s.Reload(ctx)
}

// Close
func (s *Service) Close() {
	if s.cancelFunc != nil {
		s.cancelFunc()
	}
	log.Infof(context.TODO(), "doris service context canceled")
	if db != nil {
		err := db.Close()
		if err != nil {
			log.Errorf(context.TODO(), "doris service close failed: "+err.Error())
		}
		log.Infof(context.TODO(), "doris service close success")
	}

}
