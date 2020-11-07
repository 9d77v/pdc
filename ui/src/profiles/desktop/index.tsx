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
import { GET_USER } from 'src/consts/user.gpl';
import { UpdateProfileForm } from './app/user/UpdateFrofileForm';
import AppDeviceIndex from './app/device';
import { NewUser } from 'src/models/user';
import Calculator from 'src/components/calculator';
const VideoTable = React.lazy(() => import('./admin/video/video-list'))
const VideoSeriesTable = React.lazy(() => import('./admin/video/video-series-list'))
const VideoCreateIndex = React.lazy(() => import('./admin/video/video-list/video-create'))

const UserTable = React.lazy(() => import('./admin/user/UserTable'))
const UpdatePasswordForm = React.lazy(() => import("./app/user/UpdatePasswordForm"))


const ThingTable = React.lazy(() => import('./app/thing/ThingTable'))
const ThingDashboard = React.lazy(() => import('./app/thing/ThingDashboard'))
const ThingAnalysis = React.lazy(() => import('./app/thing/ThingAnalysis'))

const VideoList = React.lazy(() => import('./app/media/video/VideoList'))
const VideoDetail = React.lazy(() => import('./app/media/video/VideoDetail'))


const HistoryPage = React.lazy(() => import("./app/media/history/HistoryPage"))


const DeviceModelIndex = React.lazy(() => import('./admin/device/device-model-list/index'))
const DeviceIndex = React.lazy(() => import('./admin/device/device-list/index'))
const DeviceDashboardList = React.lazy(() => import("./admin/device/device-dashboard-list/index"))
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
                                欢迎使用个人数据中心
                            </Route>
                            <Route exact path="/app/device">
                                <AppDeviceIndex />
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
                            <Route exact path="/app/util/calculator">
                                <Calculator />
                            </Route>
                            <Route exact path="/admin/home">
                                欢迎使用个人数据中心管理后台
                            </Route>
                            <Route exact path="/admin/video/video-list">
                                <VideoTable />
                            </Route>
                            <Route exact path="/admin/video/video-list/video-create">
                                <VideoCreateIndex />
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
