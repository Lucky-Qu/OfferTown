// Package service question.go
//
// 功能:
// - 新增题目
// - 删除题目
// - 更新题目
//
// 作者: LuckyQu
// 创建日期: 2025-10-09
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

// AddNewQuestion 新增题目
func AddNewQuestion(questionDTO *dto.CreateQuestionDTO) code.Code {
	// 开启事务
	tx := repository.NewTransaction()
	if tx.Error != nil {
		return code.DatabaseError
	}
	// 获取题目作者ID
	author, err := repository.GetUserByUsername(tx, questionDTO.AuthorName)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return code.UserNotExists
		}
		return code.DatabaseError
	}
	authorId := author.ID
	// 构建题目对象
	question := &model.Question{
		AuthorId: authorId,
		Title:    questionDTO.Title,
		Content:  questionDTO.Content,
		ImageUrl: questionDTO.ImageUrl,
		KeyPoint: questionDTO.KeyPoint,
	}
	// 检验题目合法性
	isValid, eCode := validator.IsQuestionValid(tx, question)
	if eCode != code.Success {
		tx.Rollback()
		return eCode
	}
	if !isValid {
		tx.Rollback()
		return code.InvalidQuestion
	}
	// 新建题目
	err = repository.AddNewQuestion(tx, question)
	if err != nil {
		tx.Rollback()
		return code.DatabaseError
	}
	// 检查是否有指定的分类
	if len(questionDTO.CategoryName) > 0 {
		// 获取题目ID
		addedQuestion, err := repository.GetQuestionByName(tx, question.Title)
		if err != nil {
			tx.Rollback()
			return code.DatabaseError
		}
		addedQuestionId := addedQuestion.ID
		// 循环添加分类
		for _, categoryName := range questionDTO.CategoryName {
			category, err := repository.GetCategoryByName(tx, categoryName)
			if err != nil {
				tx.Rollback()
				return code.DatabaseError
			}
			err = repository.AddQuestionToCategoryById(tx, addedQuestionId, category.ID)
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

// DeleteQuestion 删除题目
func DeleteQuestion(questionDTO *dto.DeleteQuestionDTO) code.Code {
	// 开启事务
	tx := repository.NewTransaction()
	if tx.Error != nil {
		return code.DatabaseError
	}
	// 拿到要删除的题目
	question, err := repository.GetQuestionByName(tx, questionDTO.QuestionName)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			tx.Rollback()
			return code.QuestionNotExists
		}
		tx.Rollback()
		return code.DatabaseError
	}
	// 删除题目的分类关系
	err = repository.DeleteQuestionRelationWithCategoryById(tx, question.ID)
	if err != nil {
		tx.Rollback()
		return code.DatabaseError
	}
	// 删除题目的用户关系
	err = repository.DeleteQuestionRelationWithUserById(tx, question.ID)
	if err != nil {
		tx.Rollback()
		return code.DatabaseError
	}
	// 删除题目
	err = repository.DeleteQuestionById(tx, question.ID)
	if err != nil {
		tx.Rollback()
		return code.DatabaseError
	}
	// 提交事务
	tx.Commit()
	return code.Success
}

// UpdateQuestion 更新题目
func UpdateQuestion(questionDTO *dto.UpdateQuestionDTO) code.Code {
	// 开启事务
	tx := repository.NewTransaction()
	if tx.Error != nil {
		return code.DatabaseError
	}
	// 新建更新项
	updates := make(map[string]interface{})
	// 检查传入的更新项是否需要更新
	// 获取旧题目
	question, err := repository.GetQuestionByName(tx, questionDTO.OldQuestionTitle)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			tx.Rollback()
			return code.QuestionNotExists
		}
		tx.Rollback()
		return code.DatabaseError
	}
	// 题目标题
	if questionDTO.Title != question.Title {
		isValid, eCode := validator.IsQuestionTitleValid(tx, questionDTO.Title)
		if eCode != code.Success {
			tx.Rollback()
			return eCode
		}
		if !isValid {
			tx.Rollback()
			return code.InvalidQuestion
		}
		updates["title"] = questionDTO.Title
	}
	// 题目内容
	if questionDTO.Content != question.Content {
		isValid, eCode := validator.IsQuestionContentValid(tx, questionDTO.Content)
		if eCode != code.Success {
			tx.Rollback()
			return eCode
		}
		if !isValid {
			tx.Rollback()
			return code.InvalidQuestion
		}
		updates["content"] = questionDTO.Content
	}
	// 题目图片
	if questionDTO.ImageUrl != question.ImageUrl {
		isValid, eCode := validator.IsQuestionImageValid(tx, questionDTO.ImageUrl)
		if eCode != code.Success {
			tx.Rollback()
			return eCode
		}
		if !isValid {
			tx.Rollback()
			return code.InvalidQuestion
		}
		updates["image_url"] = questionDTO.ImageUrl
	}
	// 题目关键点
	if questionDTO.KeyPoint != question.KeyPoint {
		isValid, eCode := validator.IsQuestionKeyPointValid(tx, questionDTO.KeyPoint)
		if eCode != code.Success {
			tx.Rollback()
			return eCode
		}
		if !isValid {
			tx.Rollback()
			return code.InvalidQuestion
		}
		updates["key_point"] = questionDTO.KeyPoint
	}
	// 提交数据库操作
	if len(updates) > 0 {
		err = repository.UpdateQuestion(tx, updates, question.ID)
		if err != nil {
			tx.Rollback()
			return code.DatabaseError
		}
	}
	// 检查是否更改分类-题目关系表
	if questionDTO.CategoryName != nil {
		// 清空分类-题目关系
		err = repository.DeleteQuestionRelationWithCategoryById(tx, question.ID)
		if err != nil {
			tx.Rollback()
			return code.DatabaseError
		}
		// 检测是否有新关系
		if len(questionDTO.CategoryName) > 0 {
			// 循环写入新关系
			for _, categoryName := range questionDTO.CategoryName {
				// 获取分类ID
				category, err := repository.GetCategoryByName(tx, categoryName)
				if err != nil {
					tx.Rollback()
					return code.DatabaseError
				}
				err = repository.AddQuestionToCategoryById(tx, question.ID, category.ID)
				if err != nil {
					tx.Rollback()
					return code.DatabaseError
				}
			}
		}
	}
	// 提交事务
	tx.Commit()
	return code.Success
}
