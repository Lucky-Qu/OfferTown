import api from "./index.ts";

/**
 * 获取全部分类列表。
 * 后端约定 `page=0`、`page_size=1` 代表返回所有数据。
 */
const getAllCategory = async () => {
    const response = await api.post("/category", {page: 0, page_size: 1});
    return response.data;
};

export {getAllCategory};
