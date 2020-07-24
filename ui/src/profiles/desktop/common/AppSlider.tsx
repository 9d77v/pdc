import React, { useState } from 'react';
import { SettingOutlined, PlaySquareOutlined, ShoppingOutlined, UserOutlined } from '@ant-design/icons';
import { Layout, Menu } from 'antd';
import { Link, useLocation } from 'react-router-dom';

const { Sider } = Layout;
const { SubMenu } = Menu;

const locationMap = new Map<string, any>([
    ["/app/home", {
        "defaultOpenKeys": ["home"],
        "defaultSelectedKeys": ['home']
    }],
    ["/app/settings/videos", {
        "defaultOpenKeys": ["settings"],
        "defaultSelectedKeys": ['settings-videos']
    }],
    ["/app/settings/users", {
        "defaultOpenKeys": ["settings"],
        "defaultSelectedKeys": ['settings-users']
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
    }]
])

interface AppHeaderProps {
    roleID: number;
}
export const AppSlider = (props: AppHeaderProps) => {
    const [collapsed, setCollapsed] = useState(false);
    const location = useLocation();
    const config = locationMap.get(location.pathname) || []
    const roleID = props.roleID
    return (
        <Sider width={200} className="site-layout-background" collapsible collapsed={collapsed} onCollapse={() => setCollapsed(!collapsed)}>
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
                <SubMenu
                    key="settings"
                    style={{ display: (roleID === 1 || roleID === 2) ? "block" : "none" }}
                    title={
                        <span>
                            <SettingOutlined />
                            <span>系统设置</span>
                        </span>
                    }
                >
                    <Menu.Item key="settings-videos" >
                        <Link to="/app/settings/videos">视频管理</Link>
                    </Menu.Item>
                    <Menu.Item key="settings-users"
                        style={{ display: roleID === 1 ? "block" : "none" }}
                    >
                        <Link to="/app/settings/users">用户管理</Link>
                    </Menu.Item>
                </SubMenu>
            </Menu>
        </Sider>)
}
