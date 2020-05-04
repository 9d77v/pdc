import { Layout, Menu, Breadcrumb } from 'antd';
import { SettingOutlined } from '@ant-design/icons';
import React, { useState } from 'react';
import "./index.less"
import VideoTable from './settings/VideoTable';

const { SubMenu } = Menu;
const { Header, Content, Sider } = Layout;

export default function DesktopIndex() {
    const [collapsed, setCollapsed] = useState(false);

    return (
        <Layout>
            <Header className="header">
                <div style={{ fontSize: 32, color: "white", textAlign: "left" }}>个人数据中心</div>
            </Header>
            <Layout style={{
                overflow: 'auto',
                height: 'calc(100vh - 64px)',
            }}>
                <Sider width={200} className="site-layout-background" collapsible collapsed={collapsed} onCollapse={() => setCollapsed(!collapsed)}>
                    <Menu
                        mode="inline"
                        theme="dark"
                        defaultSelectedKeys={['20']}
                        defaultOpenKeys={['sub1']}
                        style={{ height: '100%', borderRight: 0 }}
                    >
                        <SubMenu
                            key="sub1"
                            title={
                                <span>
                                    <SettingOutlined />
                                    <span>系统配置</span>
                                </span>
                            }
                        >
                            <Menu.Item key="20">视频管理</Menu.Item>
                        </SubMenu>
                    </Menu>
                </Sider>
                <Layout style={{ padding: '24px' }}>
                    <Breadcrumb style={{ textAlign: "left" }}>
                        <Breadcrumb.Item>系统配置</Breadcrumb.Item>
                        <Breadcrumb.Item>
                            <a href="/">视频管理</a>
                        </Breadcrumb.Item>
                    </Breadcrumb>
                    <Content
                        className="site-layout-background"
                        style={{
                            padding: 24,
                            margin: 0,
                            minHeight: 280,
                        }}
                    >
                        <VideoTable />
                        <div className="clear"></div>
                    </Content>
                </Layout>
            </Layout>
        </Layout>)
}