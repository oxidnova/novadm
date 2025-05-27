package sql

import (
	"net/url"
	"os"
	"testing"
	"time"

	"github.com/oxidnova/go-kit/logx"
	"github.com/oxidnova/novadm/backend/internal/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockDependencies struct {
	mock.Mock
}

func (m *mockDependencies) Logger() logx.Logger {
	args := m.Called()
	return args.Get(0).(logx.Logger)
}

func (m *mockDependencies) Config() *config.Config {
	args := m.Called()
	return args.Get(0).(*config.Config)
}

func TestNewStorage(t *testing.T) {
	dsn := getDsn()

	t.Run("success", func(t *testing.T) {
		// 创建模拟依赖
		mockDeps := new(mockDependencies)
		mockLogger := logx.Default()

		// 设置配置
		dbConfig := &config.Config{
			DB: config.DB{
				Dsn:             dsn,
				MaxOpenConns:    10,
				MaxIdleConns:    5,
				ConnMaxLifeTime: time.Hour,
				ConnMaxIdleTime: time.Minute * 30,
			},
		}

		// 设置模拟期望
		mockDeps.On("Logger").Return(mockLogger)
		mockDeps.On("Config").Return(dbConfig)

		// 创建存储实例
		storage, err := NewStorage(mockDeps)
		assert.NoError(t, err)
		assert.NotNil(t, storage)
		assert.NotNil(t, storage.db)

		// 验证模拟调用
		mockDeps.AssertExpectations(t)

		// 清理
		defer storage.Close()
	})

	t.Run("invalid_dsn", func(t *testing.T) {
		// 创建模拟依赖
		mockDeps := new(mockDependencies)
		mockLogger := logx.Default()

		// 设置无效的配置
		dbConfig := &config.Config{
			DB: config.DB{
				Dsn: "invalid_dsn",
			},
		}

		// 设置模拟期望
		mockDeps.On("Logger").Return(mockLogger)
		mockDeps.On("Config").Return(dbConfig)

		// 尝试创建存储实例
		storage, err := NewStorage(mockDeps)
		assert.Error(t, err)
		assert.Nil(t, storage)

		// 验证模拟调用
		mockDeps.AssertExpectations(t)
	})

	t.Run("with_migration_path", func(t *testing.T) {
		// 创建模拟依赖
		mockDeps := new(mockDependencies)
		mockLogger := logx.Default()

		// 设置带迁移路径的配置
		dbConfig := &config.Config{
			DB: config.DB{
				Dsn:           dsn,
				MigrationPath: "./testdata/migrations",
			},
		}

		// 设置模拟期望
		mockDeps.On("Logger").Return(mockLogger)
		mockDeps.On("Config").Return(dbConfig)

		// 创建存储实例
		storage, err := NewStorage(mockDeps)
		assert.NoError(t, err)
		assert.NotNil(t, storage)
		assert.NotNil(t, storage.db)

		// 验证模拟调用
		mockDeps.AssertExpectations(t)

		// 清理
		defer storage.Close()
	})
}

func TestStorage_Close(t *testing.T) {
	dsn := getDsn()
	// 创建模拟依赖
	mockDeps := new(mockDependencies)
	mockLogger := logx.Default()

	// 设置配置
	dbConfig := &config.Config{
		DB: config.DB{
			Dsn: dsn,
		},
	}

	// 设置模拟期望
	mockDeps.On("Logger").Return(mockLogger)
	mockDeps.On("Config").Return(dbConfig)

	// 创建存储实例
	storage, err := NewStorage(mockDeps)
	assert.NoError(t, err)
	assert.NotNil(t, storage)

	// 测试关闭
	err = storage.Close()
	assert.NoError(t, err)

	// 验证模拟调用
	mockDeps.AssertExpectations(t)
}

func getDsn() string {
	dsn := os.Getenv("TEST_DSN")
	if dsn == "" {
		return "postgres://postgres:password@localhost:5432/novadm_mock?sslmode=disable"
	}

	// modify database name
	u, err := url.Parse(dsn)
	if err != nil {
		panic(err)
	}
	u.Path = "novadm_mock"

	return u.String()
}
