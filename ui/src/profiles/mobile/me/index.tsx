import React from "react"
import { useHistory } from "react-router-dom"
import { apolloClient } from "src/utils/apollo_client"
import { QrcodeOutlined, LogoutOutlined, LockOutlined } from '@ant-design/icons'
import { List } from "antd-mobile"
import { UserBrief } from "src/profiles/mobile/common/UserBrief"
import { NewUser } from "src/models/user"
import { AppPath } from "src/consts/path"

interface IMeIndexProps {
    user: NewUser
}
const Item = List.Item

export default function MeIndex(props: IMeIndexProps) {
    const history = useHistory()
    const logout = () => {
        localStorage.clear()
        apolloClient.resetStore()
        history.push(AppPath.LOGIN)
    }
    return (
        <div style={{
            height: "100%",
            display: "flex",
            flexDirection: "column", paddingTop: 20,
            backgroundColor: "#eee"
        }}>
            <UserBrief user={props?.user} host={document.location.host} />
            <List renderHeader={() => ''}>
                <Item
                    thumb={<QrcodeOutlined />}
                    arrow="horizontal"
                    onClick={() => history.push(AppPath.UESR_QECODE)}
                >我的二维码名片</Item>
                <Item
                    thumb={<LockOutlined />}
                    onClick={() => history.push(AppPath.USER_ACCOUNT)}
                    arrow="horizontal"
                >
                    修改密码
                </Item>
                <Item
                    thumb={<LockOutlined />}
                    onClick={() => history.push(AppPath.USER_GESTURE_PASSWORD)}
                    arrow="horizontal"
                >
                    手势密码
                </Item>
                <Item
                    thumb={<LogoutOutlined />}
                    onClick={() => logout()}
                    arrow="horizontal"
                >
                    退出登录
                </Item>
            </List>
        </div>)
}
