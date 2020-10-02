import React from "react"
import { Avatar } from 'antd';
import { useHistory } from "react-router-dom";
import { apolloClient } from "../../../utils/apollo_client";
import { QrcodeOutlined, LogoutOutlined, LockOutlined } from '@ant-design/icons';
import { List } from "antd-mobile";

interface AppHeaderProps {
    name: string;
    avatar: string;
}
const Item = List.Item

export default function MeIndex(props: AppHeaderProps) {
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
            <div style={{
                display: "flex",
                flexDirection: "row",
                height: 80
            }}
                onClick={() => history.push("/app/user/profile")}
            >
                <Avatar style={{
                    backgroundColor: "#00a2ae",
                    marginLeft: 20,
                    marginRight: 20
                }} size={80} gap={1} src={props.avatar} >{props.name}</Avatar>

                <div style={{ flex: 1, fontSize: 36 }}>{props.name}</div>

            </div>
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