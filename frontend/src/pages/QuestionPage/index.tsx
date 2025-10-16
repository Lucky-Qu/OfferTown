import {useCallback, useEffect, useMemo, useState} from "react";
import NaviBar from "../../components/NaviBar";
import style from "./questionpage.module.css";
import {getAllCategory} from "../../api/category.ts";
import showErrorMessageAlert from "../../utils/ErrorMessageAlert";
import type {Category, Question} from "../../types";
import {getAllQuestions, getQuestionsByCategory} from "../../api/question.ts";
import QuestionItem from "./QuestionItem";
import PageNaviBar from "../../components/PageNaviBar";

/**
 * 后端题目接口返回的基础结构。
 */
type FetchResult = {
    questions: Question[]
};

/**
 * 根据当前窗口高度估算每页显示的题目数量。
 */
const computePageSize = () => {
    const estimated = Math.floor(window.innerHeight * 0.7 / 50) - 2;
    return Math.max(1, estimated);
};

/**
 * 题库页面：支持分类筛选、关键词过滤以及分页展示。
 */
const QuestionPage = () => {
    const [categories, setCategories] = useState<Category[]>([]);
    const [selectedCategory, setSelectedCategory] = useState<string>("all");
    const [questions, setQuestions] = useState<Question[]>([]);
    const [isLoading, setIsLoading] = useState(false);
    const [pageSize, setPageSize] = useState<number>(computePageSize());
    const [page, setPage] = useState(1);
    const [searchKeyword, setSearchKeyword] = useState("");

    // 加载分类列表
    useEffect(() => {
        getAllCategory()
            .then(res => {
                if (res.code !== 1001) {
                    showErrorMessageAlert(res.message);
                    return;
                }
                setCategories(res.data?.categories ?? []);
            })
            .catch(err => showErrorMessageAlert(err.toString()));
    }, []);

    // 监听窗口大小变化，动态调整每页展示数量
    useEffect(() => {
        const handleResize = () => {
            setPageSize(computePageSize());
        };
        window.addEventListener("resize", handleResize);
        return () => {
            window.removeEventListener("resize", handleResize);
        };
    }, []);

    /**
     * 按当前分类从后端同步题目列表。
     */
    const loadQuestions = useCallback(async (categoryName: string) => {
        setIsLoading(true);
        try {
            let data: FetchResult | null = null;
            if (categoryName === "all") {
                const res = await getAllQuestions();
                if (res.code !== 1001) {
                    showErrorMessageAlert(res.message);
                } else {
                    data = res.data;
                }
            } else {
                const res = await getQuestionsByCategory(categoryName);
                if (res.code !== 1001) {
                    showErrorMessageAlert(res.message);
                } else {
                    data = res.data;
                }
            }
            setQuestions(data?.questions ?? []);
        } catch (e) {
            showErrorMessageAlert((e as Error).message);
            setQuestions([]);
        } finally {
            setIsLoading(false);
        }
    }, []);

    useEffect(() => {
        setPage(1);
        loadQuestions(selectedCategory);
    }, [selectedCategory, loadQuestions]);

    // 当搜索词变化时重置页码
    useEffect(() => {
        setPage(1);
    }, [searchKeyword]);

    const filteredQuestions = useMemo(() => {
        if (!searchKeyword.trim()) {
            return questions;
        }
        const keyword = searchKeyword.trim().toLowerCase();
        return questions.filter(question => question.title.toLowerCase().includes(keyword));
    }, [questions, searchKeyword]);

    const maxPage = useMemo(() => {
        if (!filteredQuestions.length) {
            return 1;
        }
        return Math.max(1, Math.ceil(filteredQuestions.length / pageSize));
    }, [filteredQuestions.length, pageSize]);

    useEffect(() => {
        if (page > maxPage) {
            setPage(maxPage);
        }
    }, [page, maxPage]);

    const questionsToDisplay = useMemo(() => {
        const startIndex = (page - 1) * pageSize;
        return filteredQuestions.slice(startIndex, startIndex + pageSize);
    }, [filteredQuestions, page, pageSize]);

    return (
        <>
            <NaviBar/>
            <div className={style.container}>
                <input
                    className={style.search_questions}
                    placeholder={"搜索题目..."}
                    value={searchKeyword}
                    onChange={event => setSearchKeyword(event.target.value)}
                />
                <div className={style.main_content}>
                    <div className={style.aside_category}>
                        <div
                            className={`${style.aside_category_item} ${
                                selectedCategory === "all" ? style.aside_category_item_active : ""
                            }`}
                            onClick={() => setSelectedCategory("all")}
                        >
                            全部
                        </div>
                        {categories.map(category => (
                            <div
                                key={category.ID}
                                className={`${style.aside_category_item} ${
                                    selectedCategory === category.name ? style.aside_category_item_active : ""
                                }`}
                                onClick={() => setSelectedCategory(category.name)}
                            >
                                {category.name}
                            </div>
                        ))}
                    </div>
                    <div className={style.question_part}>
                        <div className={style.questions}>
                            <div className={style.question_title}>
                                <span className={style.question_title_no}>编号</span>
                                <span className={style.question_title_name}>题目</span>
                            </div>
                            <div className={style.question_list}>
                                {isLoading ? (
                                    <div className={style.placeholder}>题目加载中...</div>
                                ) : questionsToDisplay.length > 0 ? (
                                    questionsToDisplay.map(question => (
                                        <QuestionItem
                                            key={question.ID}
                                            questionNum={question.ID}
                                            questionTitle={question.title}
                                        />
                                    ))
                                ) : (
                                    <div className={style.placeholder}>暂无符合条件的题目</div>
                                )}
                            </div>
                        </div>
                        <div className={style.question_page_navi}>
                            <PageNaviBar page={page} maxPages={maxPage} onChangePage={nextPage => setPage(nextPage)}/>
                        </div>
                    </div>
                </div>
            </div>
        </>
    );
};

export default QuestionPage;
