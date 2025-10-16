import {createRoot} from "react-dom/client";
import ErrorMessageAlert from "../../components/ErrorMessageAlert";

/**
 * 以浮层形式展示错误提示。
 * @param message 要展示的错误信息
 * @param duration 浮层持续时间，默认 2000ms
 */
const showErrorMessageAlert = (message: string, duration?: number) => {
    const lifespan = duration ?? 2000;
    const container = document.createElement("div");
    document.body.appendChild(container);
    const root = createRoot(container);

    const handleClose = () => {
        root.unmount();
        container.remove();
    };

    root.render(<ErrorMessageAlert message={message} duration={lifespan} onClose={handleClose}/>);
};

export default showErrorMessageAlert;
