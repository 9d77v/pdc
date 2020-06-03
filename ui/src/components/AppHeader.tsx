import React from "react"
import { Layout } from 'antd';
import { Link } from "react-router-dom";
//, Menu, Avatar, Dropdown 
const { Header } = Layout;

export const AppHeader = () => {
    return (
        <Header className="header">
            <Link to="/" style={{ fontSize: 32, color: "white", textAlign: "left", float: 'left' }}>{document.title}</Link>
            <div ></div>
            {/* <div style={{ float: 'right', height: 56, alignItems: 'center', display: 'flex' }}>
                <Dropdown
                    overlay={
                        <Menu>
                            <Menu.Item key="1">退出登录</Menu.Item>
                        </Menu>} >
                    <Avatar style={{ backgroundColor: "#00a2ae", verticalAlign: 'middle', float: 'left' }} size={"large"} gap={1}>admin</Avatar>
                </Dropdown>
            </div > */}
        </Header>)
}