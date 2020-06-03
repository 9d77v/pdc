import React, { useState } from 'react';
import { SettingOutlined } from '@ant-design/icons';
import { Layout, Menu } from 'antd';
import { Link, useLocation } from 'react-router-dom';

const { Sider } = Layout;
const { SubMenu } = Menu;

const locationMap = new Map<string, any>([
    ["/settings/videos", {
        "defaultOpenKeys": ["settings"],
        "defaultSelectedKeys": ['settings-videos']
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
                    key="settings"
                    title={
                        <span>
                            <SettingOutlined />
                            <span>系统配置</span>
                        </span>
                    }
                >
                    <Menu.Item key="settings-videos">
                        <Link to="/settings/videos">视频管理</Link>
                    </Menu.Item>
                </SubMenu>
            </Menu>
        </Sider>)
}