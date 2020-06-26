import React, { useState } from 'react';
import { SettingOutlined, PlaySquareOutlined, ShoppingOutlined } from '@ant-design/icons';
import { Layout, Menu } from 'antd';
import { Link, useLocation } from 'react-router-dom';

const { Sider } = Layout;
const { SubMenu } = Menu;

const locationMap = new Map<string, any>([
    ["/app/settings/videos", {
        "defaultOpenKeys": ["settings"],
        "defaultSelectedKeys": ['settings-videos']
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
export const AppSlider = () => {
    const [collapsed, setCollapsed] = useState(false);
    const location = useLocation();
    const config = locationMap.get(location.pathname) || []
    return (
        <Sider width={200} className="site-layout-background" collapsible collapsed={collapsed} onCollapse={() => setCollapsed(!collapsed)}>
            <Menu
                mode="inline"
                theme="dark"
                defaultSelectedKeys={config["defaultSelectedKeys"]}
                defaultOpenKeys={config["defaultOpenKeys"]}
                style={{ height: '100%', borderRight: 0 }}
            >
                <SubMenu
                    key="media"
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
                    style={{ display: "none" }}
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
                    key="settings"
                    title={
                        <span>
                            <SettingOutlined />
                            <span>系统配置</span>
                        </span>
                    }
                >
                    <Menu.Item key="settings-videos">
                        <Link to="/app/settings/videos">视频管理</Link>
                    </Menu.Item>
                </SubMenu>
            </Menu>
        </Sider>)
}
