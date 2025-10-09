// Package service category.go
//
// 功能:
// - 新增分类
// - 更新分类中题目
// - 删除分类
//
// 作者: LuckyQu
// 创建日期: 2025-10-05
// 修改日期: 2025-10-09
package service

import (
	"backend/internal/code"
	"backend/internal/dto"
	"backend/internal/model"
	"backend/internal/repository"
	"backend/internal/validator"
	"errors"
	"gorm.io/gorm"
)

// AddNewCategory 新增分类
func AddNewCategory(categoryDTO *dto.CreateCategoryDTO) code.Code {
	// 获取分类和其中的题目名
	newCategoryName := categoryDTO.Name
	questionNames := categoryDTO.QuestionName
	// 开启事务
	tx := repository.NewTransaction()
	if tx.Error != nil {
		tx.Rollback()
		return code.DatabaseError
	}
	// 查询分类名是否合法
	isValid, eCode := validator.IsCategoryNameValid(tx, newCategoryName)
	if eCode != code.Success {
		tx.Rollback()
		return eCode
	}
	if !isValid {
		tx.Rollback()
		return code.InvalidCategoryName
	}
	// 新建分类
	err := repository.CreateCategory(tx, &model.Category{Name: newCategoryName})
	if err != nil {
		tx.Rollback()
		return code.DatabaseError
	}
	// 检查是否传入了题目名
	if questionNames != nil {
		// 获取分类ID
		category, err := repository.GetCategoryByName(tx, newCategoryName)
		if err != nil {
			tx.Rollback()
			return code.DatabaseError
		}
		categoryId := category.ID
		// 将题目循环写入关系表
		for _, questionName := range questionNames {
			question, err := repository.GetQuestionByName(tx, questionName)
			if err != nil {
				tx.Rollback()
				return code.DatabaseError
			}
			err = repository.AddQuestionToCategoryById(tx, question.ID, categoryId)
			if err != nil {
				tx.Rollback()
				return code.DatabaseError
			}
		}
	}
	// 提交事务
	tx.Commit()
	return code.Success
}

// UpdateCategory 更新分类
func UpdateCategory(categoryDTO *dto.UpdateCategoryDTO) code.Code {
	//开启事务
	tx := repository.NewTransaction()
	if tx.Error != nil {
		tx.Rollback()
		return code.DatabaseError
	}
	// 取到需要更新的分类
	needUpdateCategory, err := repository.GetCategoryByName(tx, categoryDTO.OldName)
	if err != nil {
		tx.Rollback()
		return code.DatabaseError
	}
	// 拿到ID便于操作
	needUpdateCategoryId := needUpdateCategory.ID
	// 新建更新项
	updates := make(map[string]interface{})
	// 检查是否需要更新分类名
	if categoryDTO.Name != "" {
		// 检查分类名合法性
		isValid, eCode := validator.IsCategoryNameValid(tx, categoryDTO.Name)
		if eCode != code.Success {
			tx.Rollback()
			return eCode
		}
		if !isValid {
			tx.Rollback()
			return code.InvalidCategoryName
		}
		// 新增分类名进更新项
		updates["name"] = categoryDTO.Name
	}
	// 提交更新
	if len(updates) > 0 {
		err = repository.UpdateCategoryById(tx, updates, needUpdateCategoryId)
		if err != nil {
			tx.Rollback()
			return code.DatabaseError
		}
	}
	// 检查是否需要更新关系表
	if categoryDTO.QuestionName != nil {
		// 清空原有关系表
		err = repository.DeleteCategoryRelationWithQuestionById(tx, needUpdateCategoryId)
		if err != nil {
			tx.Rollback()
			return code.DatabaseError
		}
		// 检查是否传入为空数组
		if len(categoryDTO.QuestionName) == 0 {
			// 提交事务
			tx.Commit()
			return code.Success
		}
		// 循环新建关系表
		for _, questionName := range categoryDTO.QuestionName {
			question, err := repository.GetQuestionByName(tx, questionName)
			if err != nil {
				tx.Rollback()
				return code.DatabaseError
			}
			err = repository.AddQuestionToCategoryById(tx, question.ID, needUpdateCategoryId)
			if err != nil {
				tx.Rollback()
				return code.DatabaseError
			}
		}
	}
	// 提交事务
	tx.Commit()
	return code.Success
}

// DeleteCategory 删除分类
func DeleteCategory(categoryName string) code.Code {
	// 开启事务
	tx := repository.NewTransaction()
	if tx.Error != nil {
		tx.Rollback()
		return code.DatabaseError
	}
	// 通过分类名拿到分类ID
	category, err := repository.GetCategoryByName(tx, categoryName)
	if err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.CategoryNotExist
		}
		return code.DatabaseError
	}
	categoryId := category.ID
	// 通过分类ID删除关系表
	err = repository.DeleteCategoryRelationWithQuestionById(tx, categoryId)
	if err != nil {
		tx.Rollback()
		return code.DatabaseError
	}
	// 通过分类ID删除分类
	err = repository.DeleteCategoryById(tx, categoryId)
	if err != nil {
		tx.Rollback()
		return code.DatabaseError
	}
	// 提交事务
	tx.Commit()
	return code.Success
}
