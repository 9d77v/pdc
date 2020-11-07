import React from 'react';
import './DisplayToolbar.less';

interface IDisplayToolbarProps {
    marginTop: number
    formula: string
    input: string
    setIsShowHistory: any,
    isShowHistory: boolean
}
const DisplayToolbar: React.FC<IDisplayToolbarProps> = ({
    marginTop = 0,
    formula,
    input,
    setIsShowHistory,
    isShowHistory
}) => {
    return (
        <div className="display-toolbar">
            <form className="display">
                <textarea
                    readOnly
                    className="display-formula"
                    value={formula}
                    style={{ marginTop: marginTop }}
                ></textarea>
                <textarea
                    readOnly
                    className="display-input"
                    rows={1}
                    value={input}></textarea>
            </form>
            {/* <div className="toolbar">
                <div className="toolbar-item" id="view-history"
                    onClick={setIsShowHistory}>
                    {isShowHistory ? "返回" : "历史"}</div>
            </div> */}
        </div>
    )
}

export default DisplayToolbar