import { NewUser } from 'src/models/user';
import { Button, Icon, List, NavBar } from 'antd-mobile'
import React, { useEffect, useState } from 'react'
import QrReader from 'react-qr-reader'
import { useHistory } from 'react-router-dom'
import { AppPath } from 'src/consts/path';

const Item = List.Item;

interface IScannerProps {
    user: NewUser
}

export const Scanner = (props: IScannerProps) => {
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
        if (props.user && result !== "" && (result.indexOf("http://") !== -1 || result.indexOf("https://") !== -1)) {
            if (result.indexOf("/pdc/") !== -1) {
                const arr = result.split("/")
                const id = arr[arr.length - 1]
                if (id !== props.user.uid) {
                    setResultDiv(
                        <Button type="primary" onClick={() => {
                            const path = AppPath.CONTACT_ADD + "?url=" + btoa(result)
                            history.replace(path)
                        }}>加好友</Button>
                    )
                } else {
                    setResultDiv(
                        <Button type="primary" onClick={() => {
                            history.replace(AppPath.USER)
                        }}>我的</Button>
                    )
                }
            } else {
                setResultDiv(
                    <Button type="primary" onClick={() => {
                        window.open(result, "_blank")
                        history.goBack()
                    }}>点击跳转</Button>
                )
            }
        }
    }, [result, history, props])
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
            <List renderHeader={() => '将二维码放入框内'}>
                {result ? <Item wrap style={{ wordBreak: "break-all" }}>{result}</Item> : null}
                {resultDiv}
            </List>
        </div>
    )
}