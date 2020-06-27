import { Layout } from 'antd';
import React from 'react';
import "./index.less"
import VideoTable from './settings/video/VideoTable';
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
import ThingTable from './thing/ThingTable';
import { ThingDashboard } from './thing/ThingDashboard';
import { ThingAnalysis } from './thing/ThingAnalysis';
import UserTable from './settings/user/UserTable';

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
                                <Route path="/app/settings/users">
                                    <UserTable />
                                </Route>
                                <Route path="/app/media/videos/:id"  >
                                    <VideoDetail />
                                </Route>
                                <Route path="/app/media/videos">
                                    <VideoPage />
                                </Route>
                                <Route path="/app/thing/dashboard">
                                    <ThingDashboard />
                                </Route>
                                <Route path="/app/thing/things">
                                    <ThingTable />
                                </Route>
                                <Route path="/app/thing/analysis">
                                    <ThingAnalysis />
                                </Route>
                            </Switch>
                        </div>
                    </Layout>
                </Layout>
            </Layout>
        </Router >
    )
}
