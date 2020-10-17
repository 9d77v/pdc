import React from "react"
import { useHistory } from "react-router-dom";
import { apolloClient } from "../../../utils/apollo_client";
import { QrcodeOutlined, LogoutOutlined, LockOutlined } from '@ant-design/icons';
import { List } from "antd-mobile";
import { UserBrief } from "../common/UserBrief";
import { NewUser } from "../../desktop/settings/user/UserCreateForm";

interface IMeIndexProps {
    user: NewUser
}
const Item = List.Item

export default function MeIndex(props: IMeIndexProps) {
    const history = useHistory();
    const logout = () => {
        localStorage.clear()
        apolloClient.resetStore()
        history.push("/login")
    }
    return (
        <div style={{
            height: "100%",
            display: "flex",
            flexDirection: "column", paddingTop: 20,
            backgroundColor: "#eee"
        }}>
            <UserBrief user={props?.user} />

            <List renderHeader={() => ''}>
                <Item
                    thumb={<QrcodeOutlined />}
                    arrow="horizontal"
                    onClick={() => history.push("/app/user/qrcode")}
                >我的二维码</Item>
                <Item
                    thumb={<LockOutlined />}
                    onClick={() => history.push("/app/user/account")}
                    arrow="horizontal"
                >
                    修改密码
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