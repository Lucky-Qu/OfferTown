import api from "./index.ts";

/**
 * 获取题目总数量。
 * 后端约定 `page_size=0` 返回计数信息。
 */
const getQuestionsCount = async () => {
    const response = await api.post("/question", {page: 0, page_size: 0});
    return response.data;
};

/**
 * 获取全部题目列表。
 * @returns 后端响应数据，包含题目数组与总数
 */
const getAllQuestions = async () => {
    const response = await api.post("/question", {page: 0, page_size: 1});
    return response.data;
};

/**
 * 分页获取题目列表。
 * @param page 当前页码（从 1 开始）
 * @param pageSize 每页数量
 */
const getQuestions = async (page: number, pageSize: number) => {
    const response = await api.post("/question", {page: page, page_size: pageSize});
    return response.data;
};

/**
 * 获取指定分类下的全部题目。
 * @param categoryName 分类名称
 */
const getQuestionsByCategory = async (categoryName: string) => {
    const response = await api.post("/question-category", {
        target: "question",
        name: categoryName,
        page: 0,
        page_size: 1,
    });
    return response.data;
};

export {getQuestionsCount, getAllQuestions, getQuestions, getQuestionsByCategory};
