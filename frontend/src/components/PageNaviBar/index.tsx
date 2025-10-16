import {useEffect, useMemo, useState} from "react";
import style from "./PageNaviBar.module.css";

/**
 * 分页导航栏组件 Props。
 * - `page`：外部当前页；
 * - `maxPages`：总页数；
 * - `onChangePage`：页码变动回调。
 */
type PageNaviBarProps = {
    page: number
    maxPages: number
    onChangePage: (page: number) => void
}

/**
 * 标准分页导航，支持省略号、上一页/下一页快捷操作。
 */
const PageNaviBar = ({page, maxPages, onChangePage}: PageNaviBarProps) => {
    const safeMaxPages = Math.max(1, maxPages);
    const [currentPage, setCurrentPage] = useState(Math.min(Math.max(1, page), safeMaxPages));

    useEffect(() => {
        setCurrentPage(prev => {
            const normalizedNext = Math.min(Math.max(1, page), safeMaxPages);
            if (prev === normalizedNext) {
                return prev;
            }
            return normalizedNext;
        });
    }, [page, safeMaxPages]);

    useEffect(() => {
        if (currentPage > safeMaxPages) {
            setCurrentPage(safeMaxPages);
        }
    }, [currentPage, safeMaxPages]);

    useEffect(() => {
        onChangePage(currentPage);
    }, [onChangePage, currentPage]);

    const pageItems = useMemo(() => {
        if (safeMaxPages <= 7) {
            return Array.from({length: safeMaxPages}, (_, index) => index + 1);
        }

        const items: Array<number | string> = [1];
        const leftBoundary = Math.max(2, currentPage - 1);
        const rightBoundary = Math.min(safeMaxPages - 1, currentPage + 1);

        if (leftBoundary > 2) {
            items.push("left-ellipsis");
        }

        for (let i = leftBoundary; i <= rightBoundary; i += 1) {
            items.push(i);
        }

        if (rightBoundary < safeMaxPages - 1) {
            items.push("right-ellipsis");
        }

        items.push(safeMaxPages);
        return items;
    }, [currentPage, safeMaxPages]);

    const handlePageChange = (nextPage: number) => {
        if (nextPage < 1 || nextPage > safeMaxPages || nextPage === currentPage) {
            return;
        }
        setCurrentPage(nextPage);
    };

    return (
        <div className={style.container}>
            <button
                type="button"
                className={`${style.navButton} ${currentPage === 1 ? style.disabled : ""}`}
                onClick={() => handlePageChange(currentPage - 1)}
                disabled={currentPage === 1}
            >
                上一页
            </button>
            <div className={style.pageList}>
                {pageItems.map((item, index) => {
                    if (typeof item === "string") {
                        return (
                            <span key={`${item}-${index}`} className={style.ellipsis}>
                                ...
                            </span>
                        );
                    }
                    return (
                        <button
                            type="button"
                            key={item}
                            className={`${style.pageButton} ${currentPage === item ? style.active : ""}`}
                            onClick={() => handlePageChange(item)}
                        >
                            {item}
                        </button>
                    );
                })}
            </div>
            <button
                type="button"
                className={`${style.navButton} ${currentPage === safeMaxPages ? style.disabled : ""}`}
                onClick={() => handlePageChange(currentPage + 1)}
                disabled={currentPage === safeMaxPages}
            >
                下一页
            </button>
        </div>
    );
};

export default PageNaviBar;
