import { Icon, NavBar } from 'antd-mobile'
import React, { useEffect, useState } from 'react'
import QrReader from 'react-qr-reader'
import { useHistory } from 'react-router-dom'
import "../../../../style/button.less"
export const Scanner = () => {
    const history = useHistory()
    const [result, setResult] = useState("")
    const [resultDiv, setResultDiv] = useState(<div />)
    const handleScan = (data: any) => {
        if (data) {
            setResult(data)
        }
    }
    const handleError = (err: any) => {
        console.error(err)
    }

    useEffect(() => {
        if (result !== "") {
            if (result.indexOf("pdc://") !== -1) {
                const url = result.replace("pdc:", document.location.protocol)
                const path = "/app/contact/addContact/" + btoa(url)
                history.replace(path)
            } else if (result.indexOf("http://") !== -1 || result.indexOf("https://") !== -1) {
                setResultDiv(
                    <div style={{
                        display: "flex", flexDirection: "column",
                        justifyContent: "center", alignItems: "center"
                    }}>
                        <div>{result}</div>
                        <div className={"pdc-button"} onClick={() => {
                            window.open(result, "_blank")
                            history.goBack()
                        }}>点击跳转</div>
                    </div>
                )
            } else {
                setResultDiv(
                    <div>{result}</div>
                )
            }
        }
    }, [result, history])
    return (
        <div style={{ height: "100%", textAlign: "center" }}>
            <NavBar
                mode="light"
                icon={<Icon type="left" />}
                onLeftClick={() => history.goBack()}
            >扫一扫</NavBar>
            <QrReader
                delay={300}
                onError={handleError}
                onScan={handleScan}
                style={{ width: '100%' }}
            />
            <div>将二维码放入框内</div>
            {resultDiv}
        </div>
    )
}