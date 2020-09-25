import React from 'react';
import {
    Route,
} from "react-router-dom";
import { Layout } from 'antd';
import "./index.less"
import { ConfigProvider } from 'antd';
import zhCN from 'antd/es/locale/zh_CN';
import { AppHeader } from './common/AppHeader';
import { AppSlider } from './common/AppSlider';
import { AppNavigator } from './common/AppNavigator';
import { useQuery } from '@apollo/react-hooks';
import { GET_USER } from '../../consts/user.gpl';
import { NewUser } from './settings/user/UserCreateForm';
import { UpdateProfileForm } from './user/UpdateFrofileForm';

const VideoTable = React.lazy(() => import('./settings/video/video-list/VideoTable'))
const VideoSeriesTable = React.lazy(() => import('./settings/video/video-series-list/VideoSeriesTable'))

const UserTable = React.lazy(() => import('./settings/user/UserTable'))
const UpdatePasswordForm = React.lazy(() => import("./user/UpdatePasswordForm"))


const ThingTable = React.lazy(() => import('./thing/ThingTable'))
const ThingDashboard = React.lazy(() => import('./thing/ThingDashboard'))
const ThingAnalysis = React.lazy(() => import('./thing/ThingAnalysis'))

const VideoList = React.lazy(() => import('./media/video/VideoList'))
const VideoDetail = React.lazy(() => import('./media/video/VideoDetail'))


const HistoryPage = React.lazy(() => import("./media/history/HistoryPage"))


const DeviceModelIndex = React.lazy(() => import('./settings/device/device-model-list/index'))
const DeviceIndex = React.lazy(() => import('./settings/device/device-list/index'))
const DeviceDashboardList = React.lazy(() => import("./settings/device/device-dashboard-list/index"))
export default function DesktopIndex() {
    const { data } = useQuery(GET_USER);
    const user: NewUser = data?.userInfo
    return (
        <ConfigProvider locale={zhCN}>
            <Layout style={{ textAlign: "center" }}>
                <AppHeader
                    name={user ? user.name.toString() : ""}
                    avatar={user ? user.avatar.toString() : ""}
                    roleID={user ? user.roleID : 0} />
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
                            <Route exact path="/app/media/videos/:id"  >
                                <VideoDetail />
                            </Route>
                            <Route exact path="/app/media/videos">
                                <VideoList />
                            </Route>
                            <Route exact path="/app/media/history">
                                <HistoryPage />
                            </Route>
                            <Route exact path="/app/thing/dashboard">
                                <ThingDashboard />
                            </Route>
                            <Route exact path="/app/thing/things">
                                <ThingTable />
                            </Route>
                            <Route exact path="/app/user/profile">
                                <UpdateProfileForm user={user} />
                            </Route>
                            <Route exact path="/app/user/account">
                                <UpdatePasswordForm />
                            </Route>
                            <Route exact path="/app/thing/analysis">
                                <ThingAnalysis />
                            </Route>
                            <Route exact path="/admin/video/video-list">
                                <VideoTable />
                            </Route>
                            <Route exact path="/admin/video/video-series-list">
                                <VideoSeriesTable />
                            </Route>
                            <Route exact path="/admin/device/device-model-list">
                                <DeviceModelIndex />
                            </Route>
                            <Route exact path="/admin/device/device-dashboard-list">
                                <DeviceDashboardList />
                            </Route>
                            <Route exact path="/admin/device/device-list">
                                <DeviceIndex />
                            </Route>
                            <Route exact path="/admin/user/user-list">
                                <UserTable />
                            </Route>
                        </div>
                    </Layout>
                </Layout>
            </Layout>
        </ConfigProvider>
    )
}
