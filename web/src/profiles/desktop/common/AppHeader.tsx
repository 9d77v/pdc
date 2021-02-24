import React from "react"
import { Layout, Dropdown, Avatar, Button } from 'antd'
import { Link, useHistory, useLocation } from "react-router-dom"
import { apolloClient } from "src/utils/apollo_client"
import { AdminPath, AppPath } from "src/consts/path"
import {
    useRecoilValue,
} from 'recoil';
import userStore from "src/module/user/user.store"

const { Header } = Layout

export const AppHeader = () => {
    const currentUserInfo = useRecoilValue(userStore.currentUserInfo)

    const history = useHistory()
    const logout = () => {
        localStorage.clear()
        apolloClient.resetStore()
        history.push(AppPath.LOGIN)
    }

    const gotoAdmin = () => {
        history.push(AdminPath.HOME)
    }

    const gotoApp = () => {
        history.push(AppPath.HOME)
    }

    const location = useLocation()
    const isApp = location.pathname.indexOf("/app") >= 0
    if (isApp) {
        document.title = "个人数据中心"
    } else {
        document.title = "个人数据中心管理后台"
    }
    return (
        <Header style={{ position: 'fixed', zIndex: 1, width: '100%' }}>
            <Link to={AppPath.HOME} style={{ fontSize: 26, color: "white", textAlign: "left", float: 'left' }}>{document.title}</Link>
            <div ></div>
            <div style={{ float: 'right', height: 56, alignItems: 'center', display: 'flex' }}>
                <Dropdown
                    overlay={
                        <div style={{
                            width: 200, height: 280, display: "flex",
                            justifyItems: "center", alignItems: "center",
                            flexDirection: "column", paddingTop: 20,
                            backgroundColor: "#fff", border: "0.5px solid"
                        }}>
                            <Avatar style={{
                                backgroundColor: "#00a2ae",
                                marginBottom: 20
                            }} size={80} gap={1} src={currentUserInfo.avatar} >{currentUserInfo.name}</Avatar>
                            <div className="title">{currentUserInfo.name}</div>
                            <Button
                                style={(currentUserInfo.roleID === 1 || currentUserInfo.roleID === 2) ? { display: "flex", marginBottom: 15 } : { display: "none" }}
                                onClick={() => isApp ? gotoAdmin() : gotoApp()} > {isApp ? "系统设置" : "退出设置"}</Button>
                            <Button onClick={() => logout()} danger >退出登录</Button>
                        </div>} >
                    <Avatar style={{ backgroundColor: "#00a2ae", verticalAlign: 'middle', float: 'left' }} size={"large"} gap={1} src={currentUserInfo.avatar} >{currentUserInfo.name}</Avatar>
                </Dropdown>
            </div >
        </Header>)
}
