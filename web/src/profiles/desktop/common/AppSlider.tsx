import React, { useState } from 'react'
import {
    PlaySquareOutlined, ShoppingOutlined, UserOutlined, LockOutlined,
    ProfileOutlined, HomeOutlined, DashboardOutlined, ToolOutlined,
    CalculatorOutlined
} from '@ant-design/icons'
import { Layout, Menu } from 'antd'
import { Link, useLocation } from 'react-router-dom'
import { AdminPath, AppPath, PathDict } from 'src/consts/path'


const { Sider } = Layout
const { SubMenu } = Menu

const locationMap = new Map<string, any>([
    [AppPath.HOME, {
        "defaultOpenKeys": ["home"],
        "defaultSelectedKeys": ['home']
    }],
    [AppPath.DEVICE, {
        "defaultOpenKeys": ["device"],
        "defaultSelectedKeys": ['device-telemetry']
    }],
    [AppPath.DEVICE_TELEMETRY, {
        "defaultOpenKeys": ["device"],
        "defaultSelectedKeys": ['device-telemetry']
    }],
    [AppPath.DEVICE_CAMERA, {
        "defaultOpenKeys": ["device"],
        "defaultSelectedKeys": ['device-camera']
    }],
    [AppPath.UTIL, {
        "defaultOpenKeys": ["util"],
        "defaultSelectedKeys": ['util']
    }],
    [AppPath.UTIL_CALCULATOR, {
        "defaultOpenKeys": ["util"],
        "defaultSelectedKeys": ['util-calculator']
    }],
    [AppPath.USER_PROFILE, {
        "defaultOpenKeys": ["user"],
        "defaultSelectedKeys": ['user-profile']
    }],
    [AppPath.USER_ACCOUNT, {
        "defaultOpenKeys": ["user"],
        "defaultSelectedKeys": ['user-account']
    }],
    [AppPath.VIDEO, {
        "defaultOpenKeys": ["video"],
        "defaultSelectedKeys": ['video-suggest']
    }],
    [AppPath.VIDEO_SUGGEST, {
        "defaultOpenKeys": ["video"],
        "defaultSelectedKeys": ['video-suggest']
    }],
    [AppPath.VIDEO_SEARCH, {
        "defaultOpenKeys": ["video"],
        "defaultSelectedKeys": ['video-search']
    }],
    [AppPath.HISTORY, {
        "defaultOpenKeys": ["video"],
        "defaultSelectedKeys": ['video-history']
    }],
    [AppPath.THING_DASHBOARD, {
        "defaultOpenKeys": ["thing"],
        "defaultSelectedKeys": ['thing-dashboard']
    }],
    [AppPath.THING_LIST, {
        "defaultOpenKeys": ["thing"],
        "defaultSelectedKeys": ['thing-things']
    }],
    [AppPath.THING_ANALYSIS, {
        "defaultOpenKeys": ["thing"],
        "defaultSelectedKeys": ['thing-analysis']
    }],
    [AdminPath.VIDEO, {
        "defaultOpenKeys": ["settings-video"],
        "defaultSelectedKeys": ['video-suggest']
    }],
    [AdminPath.VIDEO_SEREIS_LIST, {
        "defaultOpenKeys": ["settings-video"],
        "defaultSelectedKeys": ['video-suggest']
    }],
    [AdminPath.VIDEO_LIST, {
        "defaultOpenKeys": ["settings-video"],
        "defaultSelectedKeys": ['video-list']
    }],
    [AdminPath.VIDEO_SEREIS_LIST, {
        "defaultOpenKeys": ["settings-video"],
        "defaultSelectedKeys": ['video-series-list']
    }],
    [AdminPath.VIDEO_DATA_ANALYSIS, {
        "defaultOpenKeys": ["settings-video"],
        "defaultSelectedKeys": ['video-data_analysis']
    }],
    [AdminPath.DEVICE, {
        "defaultOpenKeys": ["settings-device"],
        "defaultSelectedKeys": ['device-list']
    }],
    [AdminPath.DEVICE_LIST, {
        "defaultOpenKeys": ["settings-device"],
        "defaultSelectedKeys": ['device-list']
    }],
    [AdminPath.DEVICE_MODEL_LIST, {
        "defaultOpenKeys": ["settings-device"],
        "defaultSelectedKeys": ['device-model-list']
    }],
    [AdminPath.DEVICE_DASHBOARD_LIST, {
        "defaultOpenKeys": ["settings-device"],
        "defaultSelectedKeys": ['device-dashboard-list']
    }],
    [AdminPath.USER, {
        "defaultOpenKeys": ["settings-user"],
        "defaultSelectedKeys": ['user-list']
    }],
    [AdminPath.USER_LIST, {
        "defaultOpenKeys": ["settings-user"],
        "defaultSelectedKeys": ['user-list']
    }]
])

interface IAppHeaderProps {
    roleID: number
    config?: any
}

const AppMenu = (props: IAppHeaderProps) => {
    const { config, roleID } = props
    return (
        <Menu
            mode="inline"
            theme="dark"
            defaultSelectedKeys={config["defaultSelectedKeys"]}
            defaultOpenKeys={config["defaultOpenKeys"]}
            style={{ height: '100%', borderRight: 0 }}
        >
            <Menu.Item key="home">
                <Link to={AppPath.HOME}>  <span>
                    <HomeOutlined />
                    <span>{PathDict.get(AppPath.HOME)}</span>
                </span></Link>
            </Menu.Item>
            <SubMenu
                key="video"
                style={{ display: roleID > 0 ? "block" : "none" }}
                title={
                    <span>
                        <PlaySquareOutlined />
                        <span>{PathDict.get(AppPath.VIDEO)}</span>
                    </span>
                }
            >
                <Menu.Item key="video-suggest">
                    <Link to={AppPath.VIDEO_SUGGEST}>{PathDict.get(AppPath.VIDEO_SUGGEST)}</Link>
                </Menu.Item>
                <Menu.Item key="video-search">
                    <Link to={AppPath.VIDEO_SEARCH}>{PathDict.get(AppPath.VIDEO_SEARCH)}</Link>
                </Menu.Item>
                <Menu.Item key="video-history">
                    <Link to={AppPath.HISTORY}>{PathDict.get(AppPath.HISTORY)}</Link>
                </Menu.Item>
            </SubMenu>
            <SubMenu
                key="device"
                style={{ display: roleID > 0 ? "block" : "none" }}
                title={
                    <span>
                        <DashboardOutlined />
                        <span>{PathDict.get(AppPath.DEVICE)}</span>
                    </span>
                }
            >
                <Menu.Item key="device-telemetry">
                    <Link to={AppPath.DEVICE_TELEMETRY}>{PathDict.get(AppPath.DEVICE_TELEMETRY)}</Link>
                </Menu.Item>
                <Menu.Item key="device-camera">
                    <Link to={AppPath.DEVICE_CAMERA}>{PathDict.get(AppPath.DEVICE_CAMERA)}</Link>
                </Menu.Item>
            </SubMenu>
            {/* <SubMenu
                key="thing"
                style={{ display: roleID >= 1 && roleID <= 3 ? "block" : "none" }}
                title={
                    <span>
                        <ShoppingOutlined />
                        <span>{PathDict.get(AppPath.THING)}</span>
                    </span>
                }
            >
                <Menu.Item key="thing-dashboard">
                    <Link to={AppPath.THING_DASHBOARD}>{PathDict.get(AppPath.THING_DASHBOARD)}</Link>
                </Menu.Item>
                <Menu.Item key="thing-things">
                    <Link to={AppPath.THING_LIST}>{PathDict.get(AppPath.THING_LIST)}</Link>
                </Menu.Item>
                <Menu.Item key="thing-analysis">
                    <Link to={AppPath.THING_ANALYSIS}>{PathDict.get(AppPath.THING_ANALYSIS)}</Link>
                </Menu.Item>
            </SubMenu> */}
            <SubMenu
                key="util"
                style={{ display: roleID > 0 ? "block" : "none" }}
                title={
                    <span>
                        <ToolOutlined />
                        <span>{PathDict.get(AppPath.UTIL)}</span>
                    </span>
                }
            >
                <Menu.Item key="util-calculator">
                    <Link to={AppPath.UTIL_CALCULATOR}>   <span>
                        <CalculatorOutlined />
                        <span>{PathDict.get(AppPath.UTIL_CALCULATOR)}</span>
                    </span></Link>
                </Menu.Item>
            </SubMenu>
            <SubMenu
                key="user"
                title={
                    <span>
                        <UserOutlined />
                        <span>{PathDict.get(AppPath.USER)}</span>
                    </span>
                }
            >
                <Menu.Item key="user-profile">
                    <Link to={AppPath.USER_PROFILE}> <span>
                        <ProfileOutlined />
                        <span>{PathDict.get(AppPath.USER_PROFILE)}</span>
                    </span></Link>
                </Menu.Item>
                <Menu.Item key="user-account">
                    <Link to={AppPath.USER_ACCOUNT}>
                        <span>
                            <LockOutlined />
                            <span>{PathDict.get(AppPath.USER_ACCOUNT)}</span>
                        </span>
                    </Link>
                </Menu.Item>
            </SubMenu>
        </Menu>
    )
}

const AdminMenu = (props: IAppHeaderProps) => {
    const { config, roleID } = props
    return (
        <Menu
            mode="inline"
            theme="dark"
            defaultSelectedKeys={config["defaultSelectedKeys"]}
            defaultOpenKeys={config["defaultOpenKeys"]}
            style={{ height: '100%', borderRight: 0 }}
        >
            <Menu.Item key="home">
                <Link to={AdminPath.HOME}>  <span>
                    <HomeOutlined />
                    <span>{PathDict.get(AdminPath.HOME)}</span>
                </span></Link>
            </Menu.Item>
            <SubMenu
                key="settings-user"
                style={{ display: (roleID === 1) ? "block" : "none" }}
                title={
                    <span>
                        <UserOutlined />
                        <span>{PathDict.get(AdminPath.USER)}</span>
                    </span>
                }
            >
                <Menu.Item key="user-list" >
                    <Link to={AdminPath.USER_LIST}>{PathDict.get(AdminPath.USER_LIST)}</Link>
                </Menu.Item>
            </SubMenu>
            <SubMenu
                key="settings-video"
                style={{ display: (roleID === 1 || roleID === 2) ? "block" : "none" }}
                title={
                    <span>
                        <PlaySquareOutlined />
                        <span>{PathDict.get(AdminPath.VIDEO)}</span>
                    </span>
                }
            >
                <Menu.Item key="video-list" >
                    <Link to={AdminPath.VIDEO_LIST}>{PathDict.get(AdminPath.VIDEO_LIST)}</Link>
                </Menu.Item>
                <Menu.Item key="video-series-list" >
                    <Link to={AdminPath.VIDEO_SEREIS_LIST}>{PathDict.get(AdminPath.VIDEO_SEREIS_LIST)}</Link>
                </Menu.Item>
                <Menu.Item key="video-data_analysis" >
                    <Link to={AdminPath.VIDEO_DATA_ANALYSIS}>{PathDict.get(AdminPath.VIDEO_DATA_ANALYSIS)}</Link>
                </Menu.Item>
            </SubMenu>
            <SubMenu
                key="settings-device"
                style={{ display: (roleID === 1 || roleID === 2) ? "block" : "none" }}
                title={
                    <span>
                        <DashboardOutlined />
                        <span>{PathDict.get(AdminPath.DEVICE)}</span>
                    </span>
                }
            >
                <Menu.Item key="device-dashboard-list" >
                    <Link to={AdminPath.DEVICE_DASHBOARD_LIST}>{PathDict.get(AdminPath.DEVICE_DASHBOARD_LIST)}</Link>
                </Menu.Item>
                <Menu.Item key="device-list" >
                    <Link to={AdminPath.DEVICE_LIST}>{PathDict.get(AdminPath.DEVICE_LIST)}</Link>
                </Menu.Item>
                <Menu.Item key="device-model-list" >
                    <Link to={AdminPath.DEVICE_MODEL_LIST}>{PathDict.get(AdminPath.DEVICE_MODEL_LIST)}</Link>
                </Menu.Item>

            </SubMenu>
        </Menu>
    )
}

export const AppSlider = (props: IAppHeaderProps) => {
    const [collapsed, setCollapsed] = useState(false)
    const location = useLocation()
    let config = locationMap.get(location.pathname) || []
    const roleID = props.roleID
    const isApp = location.pathname.indexOf("/app") >= 0

    return (
        <Sider width={200} className="site-layout-background" collapsible collapsed={collapsed} onCollapse={() => setCollapsed(!collapsed)}>
            {isApp ? <AppMenu roleID={roleID} config={config} /> : <AdminMenu roleID={roleID} config={config} />}
        </Sider>)
}
