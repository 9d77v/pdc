import React, { useEffect, useState } from "react";
import Buttons from "./Buttons";
import { calcOptions, CalculatorKey } from "./const";
import DisplayToolbar from "./DisplayToolbar";

import "./index.less"

interface ICalculatorProps {
    marginTop?: number
}
const Calculator: React.FC<ICalculatorProps> = ({
    marginTop = 0
}) => {
    const [formula, setFormula] = useState("")
    const [input, setInput] = useState("")
    const [historyOperations, setHistoryOperations] = useState<string[]>([])
    const [isShowHistory, setIsShowHistory] = useState(false)

    useEffect(() => {
        const hoStr = localStorage.getItem(CalculatorKey.HistoryOperations) || ''
        const historyOperations: string[] = hoStr === '' ? [] : JSON.parse(hoStr)
        const formula = localStorage.getItem(CalculatorKey.Formula) || ''
        const input = localStorage.getItem(CalculatorKey.Input) || ''
        setHistoryOperations(historyOperations)
        setFormula(formula)
        setInput(input)
    }, [])

    useEffect(() => {
        const dom = document.getElementById('calculator')
        if (dom !== null) {
            dom.scrollTop = calcOptions.historyHeight * historyOperations.length * 3
        }
    }, [historyOperations])

    useEffect(() => {
        const obj = document.getElementById("display-formula")
        if (obj) {
            obj.scrollTop = obj.scrollHeight
        }
    }, [formula])
    return (
        <div className={"calculator"}>
            <DisplayToolbar
                marginTop={marginTop}
                setFormula={setFormula}
                formula={formula}
                input={input}
                setIsShowHistory={() => setIsShowHistory(!isShowHistory)}
                isShowHistory={isShowHistory}
            />
            <Buttons
                formula={formula}
                setFormula={setFormula}
                input={input}
                setInput={setInput}
                historyOperations={historyOperations}
                setHistoryOperations={setHistoryOperations}
            />
        </div>
    );
}

export default Calculator