import { Layout } from 'antd';
import React from 'react';
import "./index.less"
import VideoTable from './settings/VideoTable';
import {
    BrowserRouter as Router,
    Switch,
    Route
} from "react-router-dom";
import { AppHeader } from '../../components/AppHeader';
import { AppSlider } from '../../components/AppSlider';
import { AppNavigator } from '../../components/AppNavigator';
const { Content } = Layout;

export default function DesktopIndex() {

    return (
        <Router>
            <Layout>
                <AppHeader />
                <Layout style={{
                    overflow: 'auto',
                    height: 'calc(100vh - 64px)',
                }}>
                    <AppSlider />
                    <Layout style={{ padding: '24px' }}>
                        <AppNavigator />
                        <Content
                            className="site-layout-background"
                            style={{
                                padding: 24,
                                margin: 0,
                                minHeight: 280,
                            }}
                        >
                            <Switch>
                                <Route exact path="/">
                                    欢迎使用{document.title}
                                </Route>
                                <Route path="/settings/videos">
                                    <VideoTable />
                                </Route>
                            </Switch>
                        </Content>
                        <div className="clear"></div>
                    </Layout>
                </Layout>
            </Layout>
        </Router >
    )
}
