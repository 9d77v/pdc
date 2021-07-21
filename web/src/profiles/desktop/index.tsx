import { lazy, useEffect } from 'react'
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
import Calculator from 'src/components/calculator'
// import DeviceCamera from './app/device/DeviceCamera'
import { GET_USER } from 'src/gqls/user/query'
import userStore from 'src/module/user/user.store'
import globalStore from 'src/module/global/global.store'
import VideoDataAnalysisIndex from './admin/video/video-data-analysis'
import DataAnalysisIndex from './app/user/DataAnalysisIndex'
import DeviceCards from 'src/profiles/common/device/DeviceCard'
import dayjs from 'dayjs';

const UpdateProfileForm = lazy(() => import('./app/user/UpdateProfileForm'))
const UpdatePasswordForm = lazy(() => import('./app/user/UpdatePasswordForm'))

const NoteIndex = lazy(() => import('./app/note'))

const EpisodePage = lazy(() => import('./app/video/EpisodePage'))
const VideoSearch = lazy(() => import('./app/video/VideoSearch'))
const VideoIndex = lazy(() => import('./app/video'))

// const ThingTable = lazy(() => import('./app/thing/ThingTable'))
// const ThingDashboard = lazy(() => import('./app/thing/ThingDashboard'))
// const ThingAnalysis = lazy(() => import('./app/thing/ThingAnalysis'))

const HistoryPage = lazy(() => import("./app/history/HistoryPage"))

const BookIndex = lazy(() => import('./app/book'))
const AppBookDetail = lazy(() => import('./app/book/AppBookDetail'))
const AppBookshelfDetail = lazy(() => import('./app/book/AppBookshelfDetail'))

//admin compnents
const VideoTable = lazy(() => import('./admin/video/video-list'))
const VideoSeriesTable = lazy(() => import('./admin/video/video-series-list'))
const VideoCreateIndex = lazy(() => import('./admin/video/video-list/video-create'))

const UserTable = lazy(() => import('./admin/user/UserTable'))

const DeviceIndex = lazy(() => import('./admin/device/device-list'))
const DeviceModelIndex = lazy(() => import('./admin/device/device-model-list/index'))
const DeviceDashboardList = lazy(() => import("./admin/device/device-dashboard-list/index"))

const BookTable = lazy(() => import('./admin/book/book-list/BookTable'))
const BookshelfTable = lazy(() => import('./admin/book/bookshelf-list/BookshelfTable'))
const BookshelfDetail = lazy(() => import('./admin/book/bookshelf-list/BookshelfDetail'))
const BookDetail = lazy(() => import('./admin/book/book-list/BookDetail'))


const DesktopIndex = () => {
    const setCurrentUserInfo = useSetRecoilState(userStore.currentUserInfo)
    const collapsed = useRecoilValue(globalStore.menuCollapsed)

    const { data, refetch } = useQuery(GET_USER)
    useEffect(() => {
        if (data &&data.userInfo) {
            setCurrentUserInfo({
                avatar:data.userInfo.avatar,
                birthDate: dayjs((data.userInfo.birthDate |0)*1000),
                color: data.userInfo.color,
                gender: data.userInfo.gender,
                ip: data.userInfo.ip,
                name: data.userInfo.name,
                roleID: data.userInfo.roleID,
                uid: data.userInfo.uid,
            })
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
                            <DeviceCards width={200} />
                        </Route>
                        {/* <Route exact path={AppPath.DEVICE_CAMERA}>
                            <DeviceCamera />
                        </Route> */}
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
                        {/* <Route exact path={AppPath.THING_DASHBOARD}>
                            <ThingDashboard />
                        </Route>
                        <Route exact path={AppPath.THING_ANALYSIS}>
                            <ThingAnalysis />
                        </Route>
                        <Route exact path={AppPath.THING_LIST}>
                            <ThingTable />
                        </Route> */}
                        <Route exact path={AppPath.BOOK_INDEX}>
                            <BookIndex />
                        </Route>
                        <Route exact path={AppPath.BOOK_BOOK_DETAIL}>
                            <AppBookDetail />
                        </Route>
                        <Route exact path={AppPath.BOOK_BOOKSHELF_DETAIL}>
                            <AppBookshelfDetail />
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
                        <Route exact path={AdminPath.BOOKSHELF_LIST}>
                            <BookshelfTable />
                        </Route>
                        <Route exact path={AdminPath.BOOKSHELF_DETAIL}>
                            <BookshelfDetail />
                        </Route>
                        <Route exact path={AdminPath.BOOK_DETAIL}>
                            <BookDetail />
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
