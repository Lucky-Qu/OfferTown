// Package service category_question.go
//
// 功能:
// - 根据请求获取指定问题或分类
//
// 作者: LuckyQu
// 创建日期: 2025-10-10
// 修改日期: 2025-10-10
package service

import (
	"backend/internal/code"
	"backend/internal/dto"
	"backend/internal/model"
	"backend/internal/repository"
	"errors"
	"gorm.io/gorm"
)

// GetCategoryQuestion 根据请求获取指定问题或分类
func GetCategoryQuestion(requestDTO *dto.GetCategoryQuestionRequestDTO) (*dto.GetCategoryQuestionResponseDTO, code.Code) {
	switch requestDTO.Target {
	case "category":
		// 获取一个题目的全部分类
		question, err := repository.GetQuestionByName(nil, requestDTO.Name)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, code.QuestionNotExists
			}
			return nil, code.DatabaseError
		}
		// 从关系表中获取分类的Id
		categoryQuestionRelation, err := repository.GetCategoriesByQuestionId(nil, question.ID)
		if err != nil {
			// 题目没有分类
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return &dto.GetCategoryQuestionResponseDTO{
					TotalCount: 0,
					Categories: []model.Category{},
				}, code.Success
			}
			return nil, code.DatabaseError
		}
		// 拿到Id
		var categoryIds []uint
		for _, relation := range categoryQuestionRelation {
			categoryIds = append(categoryIds, relation.CategoryId)
		}
		// 将查询到的分类Id查询分类
		categories, err := repository.GetCategoriesByIds(nil, categoryIds)
		if err != nil {
			return nil, code.DatabaseError
		}
		return &dto.GetCategoryQuestionResponseDTO{
			TotalCount: len(categories),
			Categories: categories,
		}, code.Success

	case "question":
		// 获取一个分类的全部题目
		// 拿到分类id
		category, err := repository.GetCategoryByName(nil, requestDTO.Name)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, code.CategoryNotExists
			}
			return nil, code.DatabaseError
		}
		// page为0表示获取数所有数据，pageSize为0表示获取数据总量
		if requestDTO.PageSize == 0 {
			// 获取总数
			categoryQuestionRelations, err := repository.GetQuestionsByCategoryId(nil, category.ID)
			if err != nil {
				return nil, code.DatabaseError
			}
			return &dto.GetCategoryQuestionResponseDTO{
				TotalCount: len(categoryQuestionRelations),
			}, code.Success
		}
		if requestDTO.Page == 0 {
			// 获取全部数据
			// 获取关系表
			categoryQuestionRelations, err := repository.GetQuestionsByCategoryId(nil, category.ID)
			if err != nil {
				return nil, code.DatabaseError
			}
			// 获取questionIds
			questionIds := make([]uint, 0, len(categoryQuestionRelations))
			for _, relation := range categoryQuestionRelations {
				questionIds = append(questionIds, relation.QuestionId)
			}
			// 批量查询题目
			questions, err := repository.GetQuestionsByIds(nil, questionIds)
			if err != nil {
				return nil, code.DatabaseError
			}
			// 返回结果
			return &dto.GetCategoryQuestionResponseDTO{
				TotalCount: len(questions),
				Questions:  questions,
			}, code.Success
		}
		// 获取指定数据
		// 解析Offset和Limit
		offset := (requestDTO.Page - 1) * requestDTO.PageSize
		limit := requestDTO.PageSize
		// 拿到关系表
		relations, err := repository.GetQuestionByCategoryIdWithOffsetAndLimit(nil, category.ID, offset, limit)
		if err != nil {
			return nil, code.DatabaseError
		}
		// 拿到questionIds
		questionIds := make([]uint, 0, len(relations))
		for _, relation := range relations {
			questionIds = append(questionIds, relation.QuestionId)
		}
		// 批量查询
		questions, err := repository.GetQuestionsByIds(nil, questionIds)
		if err != nil {
			return nil, code.DatabaseError
		}
		return &dto.GetCategoryQuestionResponseDTO{
			TotalCount: len(questions),
			Questions:  questions,
		}, code.Success
	}
	// 未知target
	return nil, code.UnknownTarget
}
