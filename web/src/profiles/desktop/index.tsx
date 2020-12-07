import React from 'react'
import {
    Route,
} from "react-router-dom"
import { Layout } from 'antd'
import "./index.less"
import { ConfigProvider } from 'antd'
import zhCN from 'antd/es/locale/zh_CN'
import { AppHeader } from './common/AppHeader'
import { AppSlider } from './common/AppSlider'
import { AppNavigator } from './common/AppNavigator'
import { useQuery } from '@apollo/react-hooks'
import { NewUser } from 'src/models/user'
import { AdminPath, AppPath } from 'src/consts/path'
import UpdateProfileForm from './app/user/UpdateFrofileForm'
import UpdatePasswordForm from './app/user/UpdatePasswordForm'
import Calculator from 'src/components/calculator'
import VideoIndex from './app/video'
import VideoSearch from './app/video/VideoSearch'
import DeviceIndex from './admin/device/device-list'
import DeviceTelemetry from './app/device/DeviceTelemetry'
import DeviceCamera from './app/device/DeviceCamera'
import { GET_USER } from 'src/gqls/user/query'
import { EpisodePage } from './app/video/EpisodePage'

const VideoTable = React.lazy(() => import('./admin/video/video-list'))
const VideoSeriesTable = React.lazy(() => import('./admin/video/video-series-list'))
const VideoCreateIndex = React.lazy(() => import('./admin/video/video-list/video-create'))

const UserTable = React.lazy(() => import('./admin/user/UserTable'))

const ThingTable = React.lazy(() => import('./app/thing/ThingTable'))
const ThingDashboard = React.lazy(() => import('./app/thing/ThingDashboard'))
const ThingAnalysis = React.lazy(() => import('./app/thing/ThingAnalysis'))

const HistoryPage = React.lazy(() => import("./app/history/HistoryPage"))

const DeviceModelIndex = React.lazy(() => import('./admin/device/device-model-list/index'))
const DeviceDashboardList = React.lazy(() => import("./admin/device/device-dashboard-list/index"))

const DesktopIndex = () => {
    const { data } = useQuery(GET_USER)
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
                            <Route exact path={AppPath.HOME}>
                                欢迎使用个人数据中心
                            </Route>
                            <Route exact path={AppPath.DEVICE_TELEMETRY}>
                                <DeviceTelemetry />
                            </Route>
                            <Route exact path={AppPath.DEVICE_CAMERA}>
                                <DeviceCamera />
                            </Route>
                            <Route exact path={AppPath.VIDEO_DETAIL}  >
                                <EpisodePage />
                            </Route>
                            <Route exact path={AppPath.VIDEO_SUGGEST}>
                                <VideoIndex />
                            </Route>
                            <Route exact path={AppPath.VIDEO_SEARCH}>
                                <VideoSearch />
                            </Route>
                            <Route exact path={AppPath.HISTORY}>
                                <HistoryPage />
                            </Route>
                            <Route exact path={AppPath.THING_DASHBOARD}>
                                <ThingDashboard />
                            </Route>
                            <Route exact path={AppPath.THING_ANALYSIS}>
                                <ThingAnalysis />
                            </Route>
                            <Route exact path={AppPath.THING_LIST}>
                                <ThingTable />
                            </Route>
                            <Route exact path={AppPath.USER_PROFILE}>
                                <UpdateProfileForm user={user} />
                            </Route>
                            <Route exact path={AppPath.USER_ACCOUNT}>
                                <UpdatePasswordForm />
                            </Route>
                            <Route exact path={AppPath.UTIL_CALCULATOR}>
                                <Calculator />
                            </Route>
                            <Route exact path={AdminPath.HOME}>
                                欢迎使用个人数据中心管理后台
                            </Route>
                            <Route exact path={AdminPath.VIDEO_LIST}>
                                <VideoTable />
                            </Route>
                            <Route exact path={AdminPath.VIDEO_CREATE}>
                                <VideoCreateIndex />
                            </Route>
                            <Route exact path={AdminPath.VIDEO_SEREIS_LIST}>
                                <VideoSeriesTable />
                            </Route>
                            <Route exact path={AdminPath.DEVICE_MODEL_LIST}>
                                <DeviceModelIndex />
                            </Route>
                            <Route exact path={AdminPath.DEVICE_DASHBOARD_LIST}>
                                <DeviceDashboardList />
                            </Route>
                            <Route exact path={AdminPath.DEVICE_LIST}>
                                <DeviceIndex />
                            </Route>
                            <Route exact path={AdminPath.USER_LIST}>
                                <UserTable />
                            </Route>
                        </div>
                    </Layout>
                </Layout>
            </Layout>
        </ConfigProvider>
    )
}

export default DesktopIndex
