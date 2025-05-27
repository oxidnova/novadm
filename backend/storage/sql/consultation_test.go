package sql

import (
	"testing"
	"time"

	"github.com/oxidnova/go-kit/x/errorx"
	"github.com/oxidnova/novadm/backend/driver/schema/code"
	"github.com/oxidnova/novadm/backend/internal/config"
	"github.com/oxidnova/novadm/backend/storage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCrossConsultation(t *testing.T) {
	// 初始化测试数据库连接
	dsn := getDsn()
	mockDeps := new(mockDependencies)

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
	mockDeps.On("Config").Return(dbConfig)

	// 创建存储实例
	stg, err := NewStorage(mockDeps)
	require.NoError(t, err)

	// 测试创建咨询
	t.Run("CreateCrossConsultation", func(t *testing.T) {
		consultation := &storage.CrossConsultation{
			Prompt:  "test prompt",
			Content: "test content",
			Status:  1,
		}

		err := stg.CreateCrossConsultation(consultation)
		assert.NoError(t, err)
		assert.NotEmpty(t, consultation.ID)
	})

	// 测试根据ID获取咨询
	t.Run("GetCrossConsultationByID", func(t *testing.T) {
		// 先创建一条记录
		consultation := &storage.CrossConsultation{
			Prompt:  "test prompt",
			Content: "test content",
			Status:  1,
		}
		err := stg.CreateCrossConsultation(consultation)
		require.NoError(t, err)

		// 获取记录
		result, err := stg.GetCrossConsultationByID(consultation.ID)
		assert.NoError(t, err)
		assert.Equal(t, consultation.Prompt, result.Prompt)
		assert.Equal(t, consultation.Content, result.Content)
		assert.Equal(t, consultation.Status, result.Status)
	})

	// 测试更新咨询
	t.Run("UpdateCrossConsultation", func(t *testing.T) {
		// 先创建一条记录
		consultation := &storage.CrossConsultation{
			Prompt:  "test prompt",
			Content: "test content",
			Status:  1,
		}
		err := stg.CreateCrossConsultation(consultation)
		require.NoError(t, err)

		// 更新记录
		consultation.Content = "updated content"
		consultation.Status = 2
		err = stg.UpdateCrossConsultation(consultation)
		assert.NoError(t, err)

		// 验证更新
		result, err := stg.GetCrossConsultationByID(consultation.ID)
		assert.NoError(t, err)
		assert.Equal(t, "updated content", result.Content)
		assert.Equal(t, 2, result.Status)
	})

	// 测试根据状态获取咨询列表
	t.Run("GetCrossConsultationsByStatus", func(t *testing.T) {
		// 创建测试数据
		consultation1 := &storage.CrossConsultation{Prompt: "p1", Content: "c1", Status: 1}
		consultation2 := &storage.CrossConsultation{Prompt: "p2", Content: "c2", Status: 1}
		require.NoError(t, stg.CreateCrossConsultation(consultation1))
		require.NoError(t, stg.CreateCrossConsultation(consultation2))

		// 获取状态为1的记录
		results, err := stg.GetCrossConsultationsByStatus(1)
		assert.NoError(t, err)
		assert.GreaterOrEqual(t, len(results), 2)
	})

	// 测试根据时间范围获取咨询列表
	t.Run("GetCrossConsultationsByTimeRange", func(t *testing.T) {
		startTime := time.Now().Add(-time.Hour)
		endTime := time.Now().Add(time.Hour)

		// 创建测试数据
		consultation := &storage.CrossConsultation{Prompt: "p1", Content: "c1", Status: 1}
		require.NoError(t, stg.CreateCrossConsultation(consultation))

		// 获取时间范围内的记录
		results, err := stg.GetCrossConsultationsByTimeRange(startTime, endTime)
		assert.NoError(t, err)
		assert.NotEmpty(t, results)
	})

	// 测试删除咨询
	t.Run("DeleteCrossConsultation", func(t *testing.T) {
		// 创建测试数据
		consultation := &storage.CrossConsultation{Prompt: "p1", Content: "c1", Status: 1}
		require.NoError(t, stg.CreateCrossConsultation(consultation))

		// 删除记录
		err := stg.DeleteCrossConsultation(consultation.ID)
		assert.NoError(t, err)

		// 验证删除
		_, err = stg.GetCrossConsultationByID(consultation.ID)
		assert.Error(t, err)
		xerr := errorx.ConvertError(err)
		assert.Equal(t, xerr.Code, code.NotFound)
	})

	// 测试使用事务更新咨询
	t.Run("UpdateCrossConsultationById", func(t *testing.T) {
		// 创建测试数据
		consultation := &storage.CrossConsultation{Prompt: "p1", Content: "c1", Status: 1}
		require.NoError(t, stg.CreateCrossConsultation(consultation))

		// 使用更新函数更新记录
		err := stg.UpdateCrossConsultationById(consultation.ID, func(old storage.CrossConsultation) (storage.CrossConsultation, error) {
			old.Content = "updated in transaction"
			old.Status = 3
			return old, nil
		})
		assert.NoError(t, err)

		// 验证更新
		result, err := stg.GetCrossConsultationByID(consultation.ID)
		assert.NoError(t, err)
		assert.Equal(t, "updated in transaction", result.Content)
		assert.Equal(t, 3, result.Status)
	})

	// 清理测试数据
	t.Cleanup(func() {
		for _, status := range []int{1, 2, 3} {
			// 获取所有咨询记录
			results, err := stg.GetCrossConsultationsByStatus(status)
			if err != nil {
				t.Logf("Failed to obtain test data: %v", err)
				return
			}

			// 删除所有测试记录
			for _, consultation := range results {
				err := stg.DeleteCrossConsultation(consultation.ID)
				if err != nil {
					t.Logf("Failed to delete test data: %v", err)
				}
			}
		}

		stg.Close()
	})
}
