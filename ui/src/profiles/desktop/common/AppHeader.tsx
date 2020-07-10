import React from "react"
import { Layout, Dropdown, Avatar } from 'antd';
import { Link, useHistory } from "react-router-dom";
import { client } from "../../../utils/client";
const { Header } = Layout;

interface AppHeaderProps {
    name: string;
    avatar: string;
}

export const AppHeader = (props: AppHeaderProps) => {
    const history = useHistory();
    const logout = () => {
        localStorage.clear()
        client.resetStore()
        history.push("/login")
    }
    return (
        <Header className="header">
            <Link to="/app/home" style={{ fontSize: 32, color: "white", textAlign: "left", float: 'left' }}>{document.title}</Link>
            <div ></div>
            <div style={{ float: 'right', height: 56, alignItems: 'center', display: 'flex' }}>
                <Dropdown
                    overlay={
                        <div style={{
                            width: 200, height: 250, display: "flex",
                            justifyItems: "center", alignItems: "center",
                            flexDirection: "column", paddingTop: 20,
                            backgroundColor: "#fff", border: "0.5px solid"
                        }}>
                            <Avatar style={{
                                backgroundColor: "#00a2ae",
                                marginBottom: 20
                            }} size={80} gap={1} src={props.avatar} >{props.name}</Avatar>
                            <div className="title">{props.name}</div>
                            <button onClick={() => logout()} >退出登录</button>
                        </div>} >
                    <Avatar style={{ backgroundColor: "#00a2ae", verticalAlign: 'middle', float: 'left' }} size={"large"} gap={1} src={props.avatar} >{props.name}</Avatar>
                </Dropdown>
            </div >
        </Header>)
}