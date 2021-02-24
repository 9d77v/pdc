import { Icon, NavBar } from 'antd-mobile'
import React from 'react'
import { useHistory } from 'react-router-dom'
import { UserBrief } from 'src/profiles/mobile/common/UserBrief'
import {
    useRecoilValue,
} from 'recoil'
import userStore from 'src/module/user/user.store'

var QRCode = require('qrcode.react')

export const QRCodePage = () => {
    const currentUserInfo = useRecoilValue(userStore.currentUserInfo)
    const history = useHistory()
    return <div style={{ height: "100%" }}>
        <NavBar
            mode="light"
            icon={<Icon type="left" />}
            onLeftClick={() => history.goBack()}
        >我的二维码名片</NavBar>
        <div style={{ marginLeft: 20, marginTop: 50, marginBottom: 20 }}> <UserBrief host={document.location.host} /></div>
        <div style={{ display: "flex", flexDirection: "column", justifyContent: "center", alignItems: "center" }}>
            <QRCode
                value={document.location.origin + "/pdc/" + currentUserInfo.uid}
                size={256}
                imageSettings={{
                    src: currentUserInfo.avatar,
                    width: 50,
                    height: 50,
                }} />
        </div>
    </div>
}
