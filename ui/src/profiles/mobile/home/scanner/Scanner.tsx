import { Icon, NavBar, SegmentedControl } from 'antd-mobile'
import React, { useEffect, useState } from 'react'
import QrReader from 'react-qr-reader'
import { useHistory } from 'react-router-dom'

export const Scanner = () => {
    const history = useHistory()
    const [result, setResult] = useState("")
    const [selectIndex, setSelectIndex] = useState(0)
    const handleScan = (data: any) => {
        if (data) {
            setResult(data)
        }
    }
    const handleError = (err: any) => {
        console.error(err)
    }
    const onChange = (e: any) => {
        setSelectIndex(e.nativeEvent.selectedSegmentIndex)
    }


    useEffect(() => {
        if (selectIndex === 0 && result !== "") {
            history.push("/app/scanner/result?url=" + result)
        }
    }, [selectIndex, result, history])
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
            <SegmentedControl values={['默认', '显示文本']} onChange={onChange} />
            {selectIndex === 1 ? <p>{result}</p> : ""}
        </div>
    )
}