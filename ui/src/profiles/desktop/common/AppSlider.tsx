import React, { useState } from 'react';
import { PlaySquareOutlined, ShoppingOutlined, UserOutlined } from '@ant-design/icons';
import { Layout, Menu } from 'antd';
import { Link, useLocation } from 'react-router-dom';

const { Sider } = Layout;
const { SubMenu } = Menu;

const locationMap = new Map<string, any>([
    ["/app/home", {
        "defaultOpenKeys": ["home"],
        "defaultSelectedKeys": ['home']
    }],
    ["/app/user/profile", {
        "defaultOpenKeys": ["user"],
        "defaultSelectedKeys": ['user-profile']
    }],
    ["/app/user/account", {
        "defaultOpenKeys": ["user"],
        "defaultSelectedKeys": ['user-account']
    }],
    ["/app/media/videos", {
        "defaultOpenKeys": ["media"],
        "defaultSelectedKeys": ['media-videos']
    }],
    ["/app/media/history", {
        "defaultOpenKeys": ["media"],
        "defaultSelectedKeys": ['media-history']
    }],
    ["/app/thing/dashboard", {
        "defaultOpenKeys": ["thing"],
        "defaultSelectedKeys": ['thing-dashboard']
    }],
    ["/app/thing/things", {
        "defaultOpenKeys": ["thing"],
        "defaultSelectedKeys": ['thing-things']
    }],
    ["/app/thing/analysis", {
        "defaultOpenKeys": ["thing"],
        "defaultSelectedKeys": ['thing-analysis']
    }],
    ["/admin/video", {
        "defaultOpenKeys": ["settings-video"],
        "defaultSelectedKeys": ['video-list']
    }],
    ["/admin/video/video-list", {
        "defaultOpenKeys": ["settings-video"],
        "defaultSelectedKeys": ['video-list']
    }],
    ["/admin/video/video-series-list", {
        "defaultOpenKeys": ["settings-video"],
        "defaultSelectedKeys": ['video-series-list']
    }],
    ["/admin/device", {
        "defaultOpenKeys": ["settings-device"],
        "defaultSelectedKeys": ['device-list']
    }],
    ["/admin/device/device-list", {
        "defaultOpenKeys": ["settings-device"],
        "defaultSelectedKeys": ['device-list']
    }],
    ["/admin/device/device-model-list", {
        "defaultOpenKeys": ["settings-device"],
        "defaultSelectedKeys": ['device-model-list']
    }],
    ["/admin/device/device-dashboard-list", {
        "defaultOpenKeys": ["settings-device"],
        "defaultSelectedKeys": ['device-dashboard-list']
    }],
    ["/admin/user", {
        "defaultOpenKeys": ["settings-user"],
        "defaultSelectedKeys": ['user-list']
    }], ["/admin/user/user-list", {
        "defaultOpenKeys": ["settings-user"],
        "defaultSelectedKeys": ['user-list']
    }]
])

interface AppHeaderProps {
    roleID: number;
    config?: any
}

const AppMenu = (props: AppHeaderProps) => {
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
                <Link to="/app/home">首页</Link>
            </Menu.Item>
            <SubMenu
                key="media"
                style={{ display: roleID > 0 ? "block" : "none" }}
                title={
                    <span>
                        <PlaySquareOutlined />
                        <span>媒体库</span>
                    </span>
                }
            >
                <Menu.Item key="media-videos">
                    <Link to="/app/media/videos">视频</Link>
                </Menu.Item>
                <Menu.Item key="media-history">
                    <Link to="/app/media/history">最近播放</Link>
                </Menu.Item>
            </SubMenu>
            <SubMenu
                key="thing"
                style={{ display: roleID >= 1 && roleID <= 3 ? "block" : "none" }}
                title={
                    <span>
                        <ShoppingOutlined />
                        <span>物品</span>
                    </span>
                }
            >
                <Menu.Item key="thing-dashboard">
                    <Link to="/app/thing/dashboard">物品概览</Link>
                </Menu.Item>
                <Menu.Item key="thing-things">
                    <Link to="/app/thing/things">物品列表</Link>
                </Menu.Item>
                <Menu.Item key="thing-analysis">
                    <Link to="/app/thing/analysis">物品分析</Link>
                </Menu.Item>
            </SubMenu>
            <SubMenu
                key="user"
                title={
                    <span>
                        <UserOutlined />
                        <span>个人设置</span>
                    </span>
                }
            >
                <Menu.Item key="user-profile"                    >
                    <Link to="/app/user/profile">个人资料</Link>
                </Menu.Item>
                <Menu.Item key="user-account"                    >
                    <Link to="/app/user/account">账户安全</Link>
                </Menu.Item>
            </SubMenu>
        </Menu>
    )
}

const AdminMenu = (props: AppHeaderProps) => {
    const { config, roleID } = props
    return (
        <Menu
            mode="inline"
            theme="dark"
            defaultSelectedKeys={config["defaultSelectedKeys"]}
            defaultOpenKeys={config["defaultOpenKeys"]}
            style={{ height: '100%', borderRight: 0 }}
        >
            <SubMenu
                key="settings-user"
                style={{ display: (roleID === 1) ? "block" : "none" }}
                title={
                    <span>
                        <UserOutlined />
                        <span>用户管理</span>
                    </span>
                }
            >
                <Menu.Item key="user-list" >
                    <Link to="/admin/user/user-list">用户列表</Link>
                </Menu.Item>
            </SubMenu>
            <SubMenu
                key="settings-video"
                style={{ display: (roleID === 1 || roleID === 2) ? "block" : "none" }}
                title={
                    <span>
                        <PlaySquareOutlined />
                        <span>视频管理</span>
                    </span>
                }
            >
                <Menu.Item key="video-list" >
                    <Link to="/admin/video/video-list">视频列表</Link>
                </Menu.Item>
                <Menu.Item key="video-series-list" >
                    <Link to="/admin/video/video-series-list">视频系列列表</Link>
                </Menu.Item>
            </SubMenu>
            <SubMenu
                key="settings-device"
                style={{ display: (roleID === 1 || roleID === 2) ? "block" : "none" }}
                title={
                    <span>
                        <span>设备管理</span>
                    </span>
                }
            >
                <Menu.Item key="device-list" >
                    <Link to="/admin/device/device-list">设备列表</Link>
                </Menu.Item>
                <Menu.Item key="device-model-list" >
                    <Link to="/admin/device/device-model-list">设备模板列表</Link>
                </Menu.Item>
                <Menu.Item key="device-dashboard-list" >
                    <Link to="/admin/device/device-dashboard-list">设备仪表盘</Link>
                </Menu.Item>
            </SubMenu>
        </Menu>
    )
}

export const AppSlider = (props: AppHeaderProps) => {
    const [collapsed, setCollapsed] = useState(false);
    const location = useLocation();
    let config = locationMap.get(location.pathname) || []
    const roleID = props.roleID
    const isApp = location.pathname.indexOf("/app") >= 0

    return (
        <Sider width={200} className="site-layout-background" collapsible collapsed={collapsed} onCollapse={() => setCollapsed(!collapsed)}>
            {isApp ? <AppMenu roleID={roleID} config={config} /> : <AdminMenu roleID={roleID} config={config} />}
        </Sider>)
}
