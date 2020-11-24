import React from "react"
import { calcOptions, CalculatorKey } from "./const"


interface IButton {
    text: string
    color: string
    backgroundColor: string
}

interface IButtonsProps {
    formula: string
    setFormula: any
    input: string
    setInput: any
    historyOperations: string[]
    setHistoryOperations: any
}

const Buttons: React.FC<IButtonsProps> = ({
    formula,
    setFormula,
    input,
    setInput,
    historyOperations,
    setHistoryOperations
}) => {
    const items: IButton[] = [
        { text: "AC", color: "#444", backgroundColor: "#d4d4d4" },
        { text: "", color: "black", backgroundColor: "black" },
        { text: "%", color: "#444", backgroundColor: "#d4d4d4" },
        { text: "退格", color: "#444", backgroundColor: "#f5f5f5" },

        { text: "7", color: "#fff", backgroundColor: "#444" },
        { text: "8", color: "#fff", backgroundColor: "#444" },
        { text: "9", color: "#fff", backgroundColor: "#444" },
        { text: "÷", color: "#fff", backgroundColor: "#f2a23c" },

        { text: "4", color: "#fff", backgroundColor: "#444" },
        { text: "5", color: "#fff", backgroundColor: "#444" },
        { text: "6", color: "#fff", backgroundColor: "#444" },
        { text: "x", color: "#fff", backgroundColor: "#f2a23c" },

        { text: "1", color: "#fff", backgroundColor: "#444" },
        { text: "2", color: "#fff", backgroundColor: "#444" },
        { text: "3", color: "#fff", backgroundColor: "#444" },
        { text: "-", color: "#fff", backgroundColor: "#f2a23c" },

        { text: ".", color: "#fff", backgroundColor: "#444" },
        { text: "0", color: "#fff", backgroundColor: "#444" },
        { text: "=", color: "#fff", backgroundColor: "#3da4ab" },
        { text: "+", color: "#fff", backgroundColor: "#f2a23c" }
    ]

    const onClick = (el: Element, index: number) => {
        const newValue = items[index].text
        const opt = formula
        const newOperation = formula + newValue
        switch (newValue) {
            case '0':
                if (opt === '0') {
                    break
                }
            // eslint-disable-next-line
            case '1':
            case '2':
            case '3':
            case '4':
            case '5':
            case '6':
            case '7':
            case '8':
            case '9':
                if (opt.length === calcOptions.maxOperationLength) {
                    return
                }
                if (opt === '0') {
                    setFormula(newValue)
                    setInput('')
                } else {
                    setFormula(newOperation)
                    if (/\+|-|x|÷/i.test(newOperation)) {
                        // eslint-disable-next-line
                        setInput(Number(eval(newOperation.replace(/x/g, '*').replace(/÷/g, '/'))).toFixed(calcOptions.floatFixed))
                    }
                }
                break
            case '.':
                if (opt.length === calcOptions.maxOperationLength) {
                    return
                }
                const a = opt.lastIndexOf('+')
                const b = opt.lastIndexOf('-')
                const c = opt.lastIndexOf('x')
                const d = opt.lastIndexOf('÷')
                let t = a
                if (b > t) {
                    t = b
                }
                if (c > t) {
                    t = c
                }
                if (d > t) {
                    t = d
                }
                if (t === -1) {
                    setFormula(newOperation)
                    break
                }
                const lastNumber = opt.substring(t, opt.length)
                if (lastNumber.indexOf('.') === -1) {
                    setFormula(newOperation)
                }
                break
            case 'C':
                setHistoryOperations([])
                setFormula('0')
                setInput('')
                break
            case 'AC':
                setFormula('0')
                setInput('')
                break
            case '退格':
                if (opt.length > 1) {
                    let o = opt.substring(0, opt.length - 1)
                    setFormula(o)
                    if (/\+|-|x|÷/i.test(o.substring(o.length - 1, o.length))) {
                        o = o.substring(0, o.length - 1)
                    }
                    if (/\+|-|x|÷/i.test(o)) {
                        // eslint-disable-next-line
                        setInput(Number(eval(o.replace(/x/g, '*').replace(/÷/g, '/'))).toFixed(calcOptions.floatFixed))
                    } else {
                        setInput('')
                    }
                } else if (opt !== '0') {
                    setFormula('0')
                    setInput('')
                }
                break
            case '+':
            case '-':
            case 'x':
            case '÷':
            case '%':
                if (opt.length === calcOptions.maxOperationLength) {
                    return
                }
                if (opt !== '') {
                    if (/\+|-|x|÷/i.test(opt.substring(opt.length - 1, opt.length))) {
                        setFormula(opt.substring(0, opt.length - 1) + newValue)
                    } else {
                        setFormula(newOperation)
                    }
                }
                break
            case '=':
                let formula = opt
                if (/\+|-|x|÷/i.test(opt.substring(opt.length - 1, opt.length))) {
                    formula = opt.substring(0, opt.length - 1)
                }
                // eslint-disable-next-line
                const input = Number(eval(formula.replace(/x/g, '*').replace(/÷/g, '/'))).toFixed(calcOptions.floatFixed)
                const ho = historyOperations
                if (ho.length === calcOptions.maxHistoryLength) {
                    ho.shift()
                }
                ho.push(formula + '=' + input)
                setHistoryOperations(ho)
                setFormula('0')
                setInput(input)
                break
            default:
                break
        }
        setTimeout(() => {
            localStorage.setItem(CalculatorKey.HistoryOperations, JSON.stringify(historyOperations))
            localStorage.setItem(CalculatorKey.Formula, formula)
            localStorage.setItem(CalculatorKey.Input, input)
        }, 1000)
    }

    const buttons = items.map((value: IButton, index: number) => {
        return <button style={{
            color: value.color,
            backgroundColor: value.backgroundColor,
            width: 80,
            height: 80,
            float: 'left',
            fontSize: "1.75rem",
            display: 'flex',
            cursor: "pointer",
            alignItems: 'center',
            justifyContent: 'center',
            margin: "0.25rem",
            border: "none",
            borderRadius: "50%"
        }}
            key={index}
            onClick={(el: any) => { onClick(el, index) }}
        >
            {value.text}
        </button>
    })
    return (
        <div className={"buttons"}>
            {buttons}
        </div >
    )
}

export default Buttons