import { Layout } from 'antd';
import React from 'react';
import "./index.less"
import VideoTable from './settings/video/VideoTable';
import {
    Route, useHistory,
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
import { useQuery } from '@apollo/react-hooks';
import { GET_USER } from '../../consts/user.gpl';
import { NewUser } from './settings/user/UserCreateForm';

export default function App() {

    const history = useHistory();
    const token = localStorage.getItem('accessToken');
    if (!token) {
        history.push('/login')
    }

    const { data } = useQuery(GET_USER);
    const user: NewUser = data?.userInfo
    return (
        <Layout style={{ textAlign: "center" }}>
            <AppHeader name={user ? user.name.toString() : ""} avatar={user ? user.avatar.toString() : ""} />
            <Layout style={{
                overflow: 'auto',
                height: 'calc(100vh - 64px)',
            }}>
                <AppSlider roleID={user ? user.roleID : 0} />
                <Layout style={{ padding: '10px' }}>
                    <AppNavigator />
                    <div className={"wall"}>
                        <Route exact path="/app/home">
                            欢迎使用{document.title}
                        </Route>
                        <Route exact path="/app/settings/videos">
                            <VideoTable />
                        </Route>
                        <Route exact path="/app/settings/users">
                            <UserTable />
                        </Route>
                        <Route exact path="/app/media/videos/:id"  >
                            <VideoDetail />
                        </Route>
                        <Route exact path="/app/media/videos">
                            <VideoPage />
                        </Route>
                        <Route exact path="/app/thing/dashboard">
                            <ThingDashboard />
                        </Route>
                        <Route exact path="/app/thing/things">
                            <ThingTable />
                        </Route>
                        <Route exact path="/app/thing/analysis">
                            <ThingAnalysis />
                        </Route>
                    </div>
                </Layout>
            </Layout>
        </Layout>
    )
}
