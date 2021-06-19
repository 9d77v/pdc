import { TabBar } from 'antd-mobile'
import { lazy, useState, useEffect } from 'react'
import { UserOutlined, HomeOutlined, MessageOutlined } from '@ant-design/icons'
import "./index.less"
import { useHistory, Route, useLocation } from 'react-router-dom'
import { useQuery } from '@apollo/react-hooks'
import { AppPath } from 'src/consts/path'
import { GET_USER } from 'src/gqls/user/query'
import {
    useSetRecoilState,
} from 'recoil';
import userStore from 'src/module/user/user.store'

import NoteIndex from './note'
import UpdateProfileForm from './me/UpdateFrofileForm'
import DeviceIndex from './device'

// const DeviceCameraDetail = lazy(() => import('./device/DeviceCameraDetail'))
const UpdatePasswordForm = lazy(() => import('./me/UpdatePasswordForm'))
const HomeNavBar = lazy(() => import('./home/HomeNavBar'))
const Scanner = lazy(() => import('./home/scanner'))
const AddFriendPage = lazy(() => import('./contact/AddFriendPage'))
const VideoIndex = lazy(() => import('./video'))
const VideoList = lazy(() => import('./video/VideoList'))
const EpisodePage = lazy(() => import('./video/EpisodePage'))
const CalculatorMobile = lazy(() => import('./home/calculator'))
const MeIndex = lazy(() => import('./me'))
const HomeIndex = lazy(() => import('./home'))
const DataAnalysisIndex = lazy(() => import('./me/DataAnalysisIndex'))
const QRCodePage = lazy(() => import('./me/QRCodePage'))
const MessageIndex = lazy(() => import('./message'))
const HistoryPage = lazy(() => import('./history/HistoryPage'))
const GesturePassword = lazy(() => import('./me/SetGesturePassword'))
const BookSearch = lazy(() => import('./book/BookSearch'))
const BookDetail = lazy(() => import('./book/BookDetail'))
const BookshelfDetail = lazy(() => import('./book/BookshelfDetail'))

const MobileIndex = () => {
    const [selectedTab, setSelectedTab] = useState("homeTab")
    const [visible, setVisible] = useState(false)
    const history = useHistory()
    const setCurrentUserInfo = useSetRecoilState(userStore.currentUserInfo)
    const { data, refetch } = useQuery(GET_USER)
    useEffect(() => {
        if (data) {
            setCurrentUserInfo(data.userInfo)
        }
    }, [data, setCurrentUserInfo])
    const location = useLocation()
    useEffect(() => {
        let isHome = true
        switch (location.pathname) {
            case AppPath.USER:
                setSelectedTab("meTab")
                break
            case AppPath.MSG:
                setSelectedTab("msgTab")
                break
            case AppPath.HOME:
                setSelectedTab("homeTab")
                break
            default:
                isHome = false
        }
        setVisible(isHome)
    }, [location])

    return (
        <div style={{ position: 'fixed', height: '100%', width: '100%', top: 0 }}>
            <Route exact path={AppPath.VIDEO_SUGGEST}>
                <VideoIndex />
            </Route>
            <Route exact path={AppPath.VIDEO_SEARCH}  >
                <VideoList />
            </Route>
            <Route exact path={AppPath.VIDEO_DETAIL}  >
                <EpisodePage />
            </Route>
            <Route exact path={AppPath.HISTORY} >
                <HistoryPage />
            </Route>
            {/* <Route exact path={AppPath.DEVICE_CAMERA_DETAIL}  >
                <DeviceCameraDetail />
            </Route> */}
            <Route path={AppPath.DEVICE}  >
                <DeviceIndex />
            </Route>
            <Route path={AppPath.BOOK_INDEX}  >
                <BookSearch />
            </Route>
            <Route path={AppPath.BOOK_BOOK_DETAIL}  >
                <BookDetail />
            </Route>
            <Route path={AppPath.BOOK_BOOKSHELF_DETAIL}  >
                <BookshelfDetail />
            </Route>
            <Route exact path={AppPath.UTIL_CALCULATOR}  >
                <CalculatorMobile />
            </Route>
            <Route exact path={AppPath.UTIL_NOTE}  >
                <NoteIndex />
            </Route>
            <Route exact path={AppPath.USER_PROFILE}  >
                <UpdateProfileForm refetch={refetch} />
            </Route>
            <Route exact path={AppPath.USER_ACCOUNT}  >
                <UpdatePasswordForm />
            </Route>
            <Route exact path={AppPath.USER_GESTURE_PASSWORD}  >
                <GesturePassword />
            </Route>
            <Route exact path={AppPath.USER_DATA_ANALYSIS}>
                <DataAnalysisIndex />
            </Route>
            <Route exact path={AppPath.UESR_QECODE}  >
                <QRCodePage />
            </Route>
            <Route exact path={AppPath.SCANNER}  >
                <Scanner />
            </Route>
            <Route exact path={AppPath.CONTACT_ADD}  >
                <AddFriendPage />
            </Route>
            <TabBar
                unselectedTintColor="#949494"
                tintColor="#33A3F4"
                barTintColor="white"
                hidden={!visible}
            >
                <TabBar.Item
                    title="首页"
                    key="home"
                    icon={<HomeOutlined />}
                    selectedIcon={<HomeOutlined style={{ color: "#85dbf5" }} />}
                    selected={selectedTab === 'homeTab'}
                    onPress={() => {
                        setSelectedTab('homeTab')
                        history.push(AppPath.HOME)
                    }}
                    data-seed="logId"
                >
                    <HomeNavBar hidden={!visible} />
                    <Route exact path={AppPath.HOME}>
                        <HomeIndex />
                    </Route>
                </TabBar.Item>
                {/* <TabBar.Item
                    icon={<MessageOutlined />}
                    selectedIcon={<MessageOutlined style={{ color: "#85dbf5" }} />}
                    title="消息"
                    key="msg"
                    selected={selectedTab === 'msgTab'}
                    onPress={() => {
                        setSelectedTab('msgTab')
                        history.push(AppPath.MSG)
                    }}
                >
                    <Route exact path={AppPath.MSG}  >
                        <MessageIndex />
                    </Route>
                </TabBar.Item> */}
                <TabBar.Item
                    icon={<UserOutlined />}
                    selectedIcon={<UserOutlined style={{ color: "#85dbf5" }} />}
                    title="我的"
                    key="my"
                    selected={selectedTab === 'meTab'}
                    onPress={() => {
                        setSelectedTab('meTab')
                        history.push(AppPath.USER)
                    }}
                >
                    <Route exact path={AppPath.USER}  >
                        <MeIndex />
                    </Route>
                </TabBar.Item>
            </TabBar>
        </div>
    )
}

export default MobileIndex
