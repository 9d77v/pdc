import { TabBar } from 'antd-mobile'
import React, { useState, useEffect } from 'react'
import { UserOutlined, HomeOutlined, MessageOutlined } from '@ant-design/icons'

import "./index.less"
import { useHistory, Route, useLocation } from 'react-router-dom'
import { useQuery } from '@apollo/react-hooks'
import { GET_USER } from 'src/consts/user.gpl'
import { UpdateProfileForm } from './me/UpdateFrofileForm'
import UpdatePasswordForm from './me/UpdatePasswordForm'
import HomeNavBar from './home/HomeNavBar'
import { Scanner } from './home/scanner'
import { QRCodePage } from './me/QRCodePage'
import { AddFriendPage } from './contact/AddFriendPage'
import { NewUser } from 'src/models/user'
import { AppPath } from 'src/consts/path'
import CalculatorMobile from './home/calculator'
import VideoIndex from './video'
import VideoList from './video/VideoList'
import VideoDetail from './video/VideoDetail'
import DeviceIndex from './device'


const MeIndex = React.lazy(() => import('./me'))
const HomeIndex = React.lazy(() => import('./home'))
const MessageIndex = React.lazy(() => import('./message'))
const HistoryPage = React.lazy(() => import('./history/HistoryPage'))
const GesturePassword = React.lazy(() => import('./me/SetGesturePassword'))

const MobileIndex = () => {
    const [selectedTab, setSelectedTab] = useState("homeTab")
    const [visible, setVisible] = useState(false)
    const history = useHistory()
    const { data } = useQuery(GET_USER)
    const user: NewUser = data?.userInfo
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
                <VideoDetail />
            </Route>
            <Route exact path={AppPath.HISTORY} >
                <HistoryPage />
            </Route>
            <Route exact path={AppPath.DEVICE}  >
                <DeviceIndex />
            </Route>
            <Route exact path={AppPath.UTIL_CALCULATOR}  >
                <CalculatorMobile />
            </Route>
            <Route exact path={AppPath.USER_PROFILE}  >
                <UpdateProfileForm user={user} />
            </Route>
            <Route exact path={AppPath.USER_ACCOUNT}  >
                <UpdatePasswordForm />
            </Route>
            <Route exact path={AppPath.USER_GESTURE_PASSWORD}  >
                <GesturePassword />
            </Route>
            <Route exact path={AppPath.UESR_QECODE}  >
                <QRCodePage user={user} />
            </Route>
            <Route exact path={AppPath.SCANNER}  >
                <Scanner user={user} />
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
                >  <HomeNavBar />
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
                        <MeIndex user={user} />
                    </Route>
                </TabBar.Item>
            </TabBar>
        </div>
    )
}

export default MobileIndex
