import { Icon, NavBar } from 'antd-mobile'
import React from 'react'
import { useHistory } from 'react-router-dom'
import { NewUser } from '../../desktop/settings/user/UserCreateForm'
import { UserBrief } from '../common/UserBrief'
var QRCode = require('qrcode.react')
interface IQRCodeProps {
    user: NewUser
}

export const QRCodePage = (props: IQRCodeProps) => {
    const history = useHistory()
    return <div style={{ height: "100%" }}>
        <NavBar
            mode="light"
            icon={<Icon type="left" />}
            onLeftClick={() => history.goBack()}
        >我的二维码名片</NavBar>
        <div style={{ marginLeft: 20, marginTop: 50, marginBottom: 20 }}> <UserBrief user={props?.user} host={document.location.host} /></div>
        <div style={{ display: "flex", flexDirection: "column", justifyContent: "center", alignItems: "center" }}>
            <QRCode
                value={document.location.origin + "/pdc/" + props.user?.uid}
                size={256}
                imageSettings={{
                    src: props.user?.avatar,
                    width: 50,
                    height: 50,
                }} />
        </div>
    </div>
}