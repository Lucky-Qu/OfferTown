import {createRoot} from "react-dom/client";
import ErrorMessageAlert from "../../components/ErrorMessageAlert";

const createErrorMessageAlert = (message:string, duration?:number) => {
    duration = duration || 5000;
    const container = document.createElement("div")
    document.body.appendChild(container)
    const root = createRoot(container)
    root.render(<ErrorMessageAlert message={message}, duration={duration}></>)
}