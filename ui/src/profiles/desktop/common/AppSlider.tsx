import React, { useState } from 'react';
import { SettingOutlined, PlaySquareOutlined } from '@ant-design/icons';
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