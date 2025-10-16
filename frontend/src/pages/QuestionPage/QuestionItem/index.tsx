import style from "./QuestionItem.module.css";

/**
 * 题目列表项组件的属性。
 */
type QuestionItemProps = {
    questionNum: number
    questionTitle: string
}

/**
 * 展示题目编号、标题以及跳转按钮的列表项。
 */
const QuestionItem = (props: QuestionItemProps) => {
    return (
        <div className={style.question_items}>
            <div className={style.question_item}>
                <span className={style.question_item_no}>{props.questionNum}</span>
                <span className={style.question_item_title}>{props.questionTitle}</span>
                <button className={style.question_item_button}>前往</button>
            </div>
        </div>
    );
};

export default QuestionItem;
