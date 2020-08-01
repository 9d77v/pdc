import React from "react"
import { Layout, Dropdown, Avatar } from 'antd';
import { Link, useHistory, useLocation } from "react-router-dom";
import { client } from "../../../utils/client";
const { Header } = Layout;

interface AppHeaderProps {
    name: string;
    avatar: string;
    roleID: number;
}

export const AppHeader = (props: AppHeaderProps) => {
    const history = useHistory();
    const logout = () => {
        localStorage.clear()
        client.resetStore()
        history.push("/login")
    }

    const gotoAdmin = () => {
        history.push("/admin/videos/video-list")
    }

    const gotoApp = () => {
        history.push("/app/home")
    }

    const location = useLocation();
    const roleID = props.roleID
    const isApp = location.pathname.indexOf("/app") >= 0
    if (isApp) {
        document.title = "个人数据中心"
    } else {
        document.title = "个人数据中心管理后台"
    }
    return (
        <Header className="header">
            <Link to="/app/home" style={{ fontSize: 26, color: "white", textAlign: "left", float: 'left' }}>{document.title}</Link>
            <div ></div>
            <div style={{ float: 'right', height: 56, alignItems: 'center', display: 'flex' }}>
                <Dropdown
                    overlay={
                        <div style={{
                            width: 200, height: 270, display: "flex",
                            justifyItems: "center", alignItems: "center",
                            flexDirection: "column", paddingTop: 20,
                            backgroundColor: "#fff", border: "0.5px solid"
                        }}>
                            <Avatar style={{
                                backgroundColor: "#00a2ae",
                                marginBottom: 20
                            }} size={80} gap={1} src={props.avatar} >{props.name}</Avatar>
                            <div className="title">{props.name}</div>
                            <button
                                style={(roleID === 1 || roleID === 2) ? { display: "flex", marginBottom: 5 } : { display: "none" }}
                                onClick={() => isApp ? gotoAdmin() : gotoApp()} > {isApp ? "系统设置" : "退出设置"}</button>
                            <button onClick={() => logout()} >退出登录</button>
                        </div>} >
                    <Avatar style={{ backgroundColor: "#00a2ae", verticalAlign: 'middle', float: 'left' }} size={"large"} gap={1} src={props.avatar} >{props.name}</Avatar>
                </Dropdown>
            </div >
        </Header>)
}