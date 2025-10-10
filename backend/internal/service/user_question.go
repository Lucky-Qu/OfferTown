// Package service user_question.go
//
// 功能:
// - 根据用户获取题目
// - 根据题目获取用户
//
// 作者: LuckyQu
// 创建日期: 2025-10-10
// 修改日期: 2025-10-10
package service

import (
	"backend/internal/code"
	"backend/internal/dto"
	"backend/internal/repository"
)

// GetQuestionsByUser 根据用户获取题目
func GetQuestionsByUser(requestDTO *dto.GetQuestionsByUserRequestDTO, userId uint) (*dto.GetQuestionsByUserResponseDTO, code.Code) {
	// 根据用户Id获取题目关系表
	relations, err := repository.GetQuestionsByUserId(nil, userId)
	if err != nil {
		return nil, code.DatabaseError
	}
	// 页容纳数为0，返回数量
	if requestDTO.PageSize == 0 {
		// 长度即为关系表长度
		// 返回数据
		return &dto.GetQuestionsByUserResponseDTO{
			Questions:  nil,
			TotalCount: len(relations),
		}, code.Success
	}
	// 页数为0，返回全部数据
	if requestDTO.Page == 0 {
		// 从关系表中获取题目Ids
		questionIds := make([]uint, 0, len(relations))
		for _, relation := range relations {
			questionIds = append(questionIds, relation.QuestionId)
		}
		// 用题目Ids批量查询
		questions, err := repository.GetQuestionsByIds(nil, questionIds)
		if err != nil {
			return nil, code.DatabaseError
		}
		// 返回题目数据
		return &dto.GetQuestionsByUserResponseDTO{
			Questions:  questions,
			TotalCount: len(questions),
		}, code.Success
	}
	// 查询指定数据
	// 提取Offset，Limit
	offset := (requestDTO.Page - 1) * requestDTO.PageSize
	limit := requestDTO.PageSize
	// 结合用户Id，Offset，Limit获取题目关系表
	relationsWithOffsetAndLimit, err := repository.GetQuestionsByUserIdWithOffsetAndLimit(nil, userId, offset, limit)
	if err != nil {
		return nil, code.DatabaseError
	}
	// 从添加Offset和Limit关系表中获取题目Ids
	questionIds := make([]uint, 0, len(relationsWithOffsetAndLimit))
	for _, relation := range relationsWithOffsetAndLimit {
		questionIds = append(questionIds, relation.QuestionId)
	}
	// 结合Offset，Limit，Ids查询
	questions, err := repository.GetQuestionsByIds(nil, questionIds)
	if err != nil {
		return nil, code.DatabaseError
	}
	// 返回题目数据
	return &dto.GetQuestionsByUserResponseDTO{
		Questions:  questions,
		TotalCount: len(relations),
	}, code.Success
}

// GetUsersByQuestion 根据题目获取用户
func GetUsersByQuestion(requestDTO *dto.GetUsersByQuestionRequestDTO) (*dto.GetUsersByQuestionResponseDTO, code.Code) {
	// 获取题目结构体
	question, err := repository.GetQuestionByName(nil, requestDTO.QuestionName)
	if err != nil {
		return nil, code.DatabaseError
	}
	// 获取总数
	relationsWithoutOffsetAndLimit, err := repository.GetUsersByQuestionId(nil, question.ID)
	// 页容纳数为0,返回数量
	if requestDTO.PageSize == 0 {
		return &dto.GetUsersByQuestionResponseDTO{
			Usernames:  nil,
			TotalCount: len(relationsWithoutOffsetAndLimit),
		}, code.Success
	}
	// 页数为0，返回全部数据
	if requestDTO.Page == 0 {
		// 提取关系表中的UsersIds
		userIds := make([]uint, 0, len(relationsWithoutOffsetAndLimit))
		for _, relation := range relationsWithoutOffsetAndLimit {
			userIds = append(userIds, relation.UserId)
		}
		// 批量根据UserIds查询Users
		users, err := repository.GetUsersByIds(nil, userIds)
		if err != nil {
			return nil, code.DatabaseError
		}
		// 将Users中的Username提取出
		usernames := make([]string, 0, len(users))
		for _, user := range users {
			usernames = append(usernames, user.Username)
		}
		// 返回结果
		return &dto.GetUsersByQuestionResponseDTO{
			Usernames:  usernames,
			TotalCount: len(relationsWithoutOffsetAndLimit),
		}, code.Success
	}
	// 查询指定数据
	// 提取Offset，Limit
	offset := (requestDTO.Page - 1) * requestDTO.PageSize
	limit := requestDTO.PageSize
	// 结合QuestionId，Offset，Limit查询关系表
	relations, err := repository.GetUsersByQuestionIdWithOffsetAndLimit(nil, question.ID, offset, limit)
	if err != nil {
		return nil, code.DatabaseError
	}
	// 提取关系表中的UsersIds
	userIds := make([]uint, 0, len(relations))
	for _, relation := range relations {
		userIds = append(userIds, relation.UserId)
	}
	// 批量根据UserIds查询Users
	users, err := repository.GetUsersByIds(nil, userIds)
	if err != nil {
		return nil, code.DatabaseError
	}
	// 将Users中的Username提取出
	usernames := make([]string, 0, len(users))
	for _, user := range users {
		usernames = append(usernames, user.Username)
	}
	// 返回结果
	return &dto.GetUsersByQuestionResponseDTO{
		Usernames:  usernames,
		TotalCount: len(relationsWithoutOffsetAndLimit),
	}, code.Success
}
