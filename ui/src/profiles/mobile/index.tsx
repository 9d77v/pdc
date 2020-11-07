import { TabBar } from 'antd-mobile';
import React, { useState, useEffect } from 'react';
import { UserOutlined, HomeOutlined, MessageOutlined } from '@ant-design/icons';

import "./index.less"
import { useHistory, Route, useLocation } from 'react-router-dom';
import { useQuery } from '@apollo/react-hooks';
import { GET_USER } from 'src/consts/user.gpl';
import { UpdateProfileForm } from './me/UpdateFrofileForm';
import UpdatePasswordForm from './me/UpdatePasswordForm';
import HomeNavBar from './home/HomeNavBar';
import { Scanner } from './home/scanner';
import { QRCodePage } from './me/QRCodePage';
import { AddFriendPage } from './contact/AddFriendPage';
import { NewUser } from 'src/models/user';

const CalculatorMobile = React.lazy(() => import('./home/calculator'))
const MeIndex = React.lazy(() => import('./me'))
const HomeIndex = React.lazy(() => import('./home'))
const MessageIndex = React.lazy(() => import('./message'))
const VideoDetail = React.lazy(() => import('./media/video/VideoDetail'))
const VideoList = React.lazy(() => import('./media/video/VideoList'))
const VideoNavBar = React.lazy(() => import('./media/video/VideoNavBar'))
const HistoryPage = React.lazy(() => import('./media/history/HistoryPage'))
const DeviceIndex = React.lazy(() => import('./device'))
const GesturePassword = React.lazy(() => import('./me/SetGesturePassword'))

export default function MobileIndex() {
    const [selectedTab, setSelectedTab] = useState("homeTab")
    const [visible, setVisible] = useState(false)
    const history = useHistory();
    const { data } = useQuery(GET_USER);
    const user: NewUser = data?.userInfo

    const location = useLocation();

    useEffect(() => {
        let isHome = true
        switch (location.pathname) {
            case "/app/user":
                setSelectedTab("meTab")
                break
            case "/app/msg":
                setSelectedTab("msgTab")
                break
            case "/app/home":
                setSelectedTab("homeTab")
                break
            default:
                isHome = false
        }
        setVisible(isHome)
    }, [location])

    return (
        <div style={{ position: 'fixed', height: '100%', width: '100%', top: 0 }}>
            <Route exact path="/app/media/videos">
                <div style={{ height: "100%", overflowY: "auto" }}>
                    <VideoNavBar />
                    <VideoList />
                </div>
            </Route>
            <Route exact path="/app/media/videos/:id"  >
                <VideoDetail />
            </Route>
            <Route exact path="/app/media/history"  >
                <HistoryPage />
            </Route>
            <Route exact path="/app/device"  >
                <DeviceIndex />
            </Route>
            <Route exact path="/app/util/calculator"  >
                <CalculatorMobile />
            </Route>
            <Route exact path="/app/user/profile"  >
                <UpdateProfileForm user={user} />
            </Route>
            <Route exact path="/app/user/account"  >
                <UpdatePasswordForm />
            </Route>
            <Route exact path="/app/user/gesture_password"  >
                <GesturePassword />
            </Route>
            <Route exact path="/app/user/qrcode"  >
                <QRCodePage user={user} />
            </Route>
            <Route exact path="/app/scanner"  >
                <Scanner user={user} />
            </Route>
            <Route exact path="/app/contact/addContact/:url"  >
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
                        history.push('/app/home')
                    }}
                    data-seed="logId"
                >  <HomeNavBar />
                    <Route exact path="/app/home">
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
                        history.push('/app/msg')
                    }}
                >
                    <Route exact path="/app/msg"  >
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
                        history.push('/app/user')
                    }}
                >
                    <Route exact path="/app/user"  >
                        <MeIndex user={user} />
                    </Route>
                </TabBar.Item>
            </TabBar>
        </div>
    );
}
