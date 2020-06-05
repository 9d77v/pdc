import { Layout } from 'antd';
import React from 'react';
import "./index.less"
import VideoTable from './settings/VideoTable';
import {
    BrowserRouter as Router,
    Switch,
    Route
} from "react-router-dom";
import { AppHeader } from './common/AppHeader';
import { AppSlider } from './common/AppSlider';
import { AppNavigator } from './common/AppNavigator';
import { VideoPage } from './media/VideoPage';
import { VideoDetail } from './media/VideoDetail';

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
                    <Layout style={{ padding: '10px' }}>
                        <AppNavigator />
                        <div className={"wall"}>
                            <Switch>
                                <Route exact path="/">
                                    欢迎使用{document.title}
                                </Route>
                                <Route path="/app/settings/videos">
                                    <VideoTable />
                                </Route>
                                <Route path="/app/media/videos/:id"  >
                                    <VideoDetail />
                                </Route>
                                <Route path="/app/media/videos">
                                    <VideoPage />
                                </Route>
                            </Switch>
                        </div>
                    </Layout>
                </Layout>
            </Layout>
        </Router >
    )
}
