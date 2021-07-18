import { Icon, NavBar } from 'antd-mobile'
import { useHistory } from 'react-router-dom'
import { UserBrief } from 'src/profiles/mobile/common/UserBrief'
import {
    useRecoilValue,
} from 'recoil'
import userStore from 'src/module/user/user.store'
import { getServerURL } from 'src/consts/http'

var QRCode = require('qrcode.react')

const QRCodePage = () => {
    const currentUserInfo = useRecoilValue(userStore.currentUserInfo)
    const history = useHistory()
    let value=document.location.origin + "/pdc/" + currentUserInfo.uid
    if (document.location.protocol === "https:") { 
        value=getServerURL()+"/pdc/" + currentUserInfo.uid
    }
        return <div style={{ height: "100%" }}>
        <NavBar
            mode="light"
            icon={<Icon type="left" />}
            onLeftClick={() => history.goBack()}
        >我的二维码名片</NavBar>
        <div style={{ marginLeft: 20, marginTop: 50, marginBottom: 20 }}> <UserBrief host={document.location.host} /></div>
        <div style={{ display: "flex", flexDirection: "column", justifyContent: "center", alignItems: "center" }}>
            <QRCode
                value={value}
                size={256}
                imageSettings={{
                    src: currentUserInfo.avatar,
                    width: 50,
                    height: 50,
                }} />
        </div>
    </div>
}

export default QRCodePage
