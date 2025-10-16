import {useEffect, useState} from "react";
import style from "./ErrorMessageAlert.module.css";

/**
 * 错误提示浮层组件的 Props。
 * - `duration` 表示提示展示的时间，单位毫秒。
 * - `onClose` 在浮层销毁后触发，用于清理容器节点。
 */
interface ErrorMessageAlertProps {
    message: string
    duration: number
    onClose: () => void
}

/**
 * 短暂展示后自动淡出的错误提示浮层。
 */
const ErrorMessageAlert = (props: ErrorMessageAlertProps) => {
    const durationTime = props.duration;
    const [show, setShow] = useState(true);

    useEffect(() => {
        const timer = setTimeout(() => {
            setShow(false);
        }, durationTime);
        const closeTimer = setTimeout(props.onClose, durationTime + 500);
        return () => {
            clearTimeout(timer);
            clearTimeout(closeTimer);
        };
    }, [durationTime, props.onClose]);

    return (
        <div className={`${style.card} ${show ? "" : style.inactive}`}>
            <div className={style.message}>{props.message}</div>
        </div>
    );
};

export default ErrorMessageAlert;
