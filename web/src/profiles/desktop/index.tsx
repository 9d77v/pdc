import React, { useEffect } from 'react'
import {
    Route,
} from "react-router-dom"
import {
    useRecoilValue,
    useSetRecoilState,
} from 'recoil';
import { Layout } from 'antd'
import "./index.less"
import { ConfigProvider } from 'antd'
import zhCN from 'antd/es/locale/zh_CN'
import { AppHeader } from './common/AppHeader'
import { AppSlider } from './common/AppSlider'
import { AppNavigator } from './common/AppNavigator'
import { useQuery } from '@apollo/react-hooks'
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
import VideoDataAnalysisIndex from './admin/video/video-data-analysis'
import DataAnalysisIndex from './app/user/DataAnalysisIndex'
import NoteIndex from './app/note'
import userStore from 'src/module/user/user.store'
import globalStore from 'src/module/global/global.store';
import BookTable from './admin/book/book-list/BookTable';

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
    const setCurrentUserInfo = useSetRecoilState(userStore.currentUserInfo)
    const collapsed = useRecoilValue(globalStore.menuCollapsed)

    const { data, refetch } = useQuery(GET_USER)
    useEffect(() => {
        if (data) {
            setCurrentUserInfo(data.userInfo)
        }
    }, [data, setCurrentUserInfo])
    return (
        <ConfigProvider locale={zhCN}>
            <Layout style={{ textAlign: "center" }}>
                <AppHeader />
                <Layout style={{
                    height: 'calc(100vh - 64px)',
                }}>
                    <AppSlider />
                    <Layout style={{
                        padding: '10px', paddingLeft: collapsed ? 100 : 220, marginTop: 64,
                        minHeight: 'calc(100vh - 64px)', height: "100%", backgroundColor: "#fff"
                    }}>
                        <AppNavigator />
                        <Route exact path={AppPath.HOME}>
                            欢迎使用个人数据中心
                            </Route>
                        <Route exact path={AppPath.DEVICE_TELEMETRY}>
                            <DeviceTelemetry />
                        </Route>
                        <Route exact path={AppPath.DEVICE_CAMERA}>
                            <DeviceCamera />
                        </Route>
                        <Route exact path={AppPath.VIDEO_DETAIL}>
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
                            <UpdateProfileForm refetch={refetch} />
                        </Route>
                        <Route exact path={AppPath.USER_ACCOUNT}>
                            <UpdatePasswordForm />
                        </Route>
                        <Route exact path={AppPath.USER_DATA_ANALYSIS}>
                            <DataAnalysisIndex />
                        </Route>
                        <Route path={AppPath.UTIL_NOTE}>
                            <NoteIndex />
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
                        <Route exact path={AdminPath.VIDEO_DATA_ANALYSIS}>
                            <VideoDataAnalysisIndex />
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
                        <Route exact path={AdminPath.BOOK_LIST}>
                            <BookTable />
                        </Route>
                        <Route exact path={AdminPath.USER_LIST}>
                            <UserTable />
                        </Route>
                    </Layout>
                </Layout>
            </Layout>
        </ConfigProvider>
    )
}

export default DesktopIndex
