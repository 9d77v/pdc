import React from "react"
import { Avatar } from 'antd';
import { useHistory } from "react-router-dom";
import { apolloClient } from "../../../utils/apollo_client";
import "../../../style/button.less"

interface AppHeaderProps {
    name: string;
    avatar: string;
}

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
            <div className="pdc-button" style={{
                width: "100%",
                marginTop: 20,
                marginLeft: 0,
                marginRight: 0,
                borderRadius: 5
            }} onClick={() => history.push("/app/user/account")} >修改密码</div>
            <div className="pdc-button" style={{
                width: "100%",
                marginLeft: 0,
                marginTop: 10,
                marginRight: 0,
                borderRadius: 5
            }} onClick={() => logout()} >退出登录</div>
        </div>)
}