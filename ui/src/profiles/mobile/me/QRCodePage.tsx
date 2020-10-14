import { Icon, NavBar } from 'antd-mobile'
import React from 'react'
import { useHistory } from 'react-router-dom'
var QRCode = require('qrcode.react')
interface IQRCodeProps {
    text: string
}
export const QRCodePage = (props: IQRCodeProps) => {
    const history = useHistory()
    return <div style={{ height: "100%" }}>
        <NavBar
            mode="light"
            icon={<Icon type="left" />}
            onLeftClick={() => history.goBack()}
        >我的二维码</NavBar>
        <div style={{ marginTop: 200, display: "flex", justifyContent: "center", alignItems: "center" }}>
            <QRCode value={"pdc://" + document.location.host + "/card/" + props.text} size={256} />
        </div>
    </div>
}