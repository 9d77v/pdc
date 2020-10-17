import { TabBar } from 'antd-mobile';
import React, { useState, useEffect } from 'react';
import { UserOutlined, HomeOutlined, PlaySquareOutlined } from '@ant-design/icons';

import "./index.less"
import { useHistory, Route, useLocation } from 'react-router-dom';
import { useQuery } from '@apollo/react-hooks';
import { GET_USER } from '../../consts/user.gpl';
import { NewUser } from '../desktop/settings/user/UserCreateForm';
import { UpdateProfileForm } from './me/UpdateFrofileForm';
import UpdatePasswordForm from './me/UpdatePasswordForm';
import HomeNavBar from './home/HomeNavBar';
import { Scanner } from './home/scanner';
import { QRCodePage } from './me/QRCodePage';
import { AddFriendPage } from './contact/AddFriendPage';

const MeIndex = React.lazy(() => import('./me'))
const HomeIndex = React.lazy(() => import('./home'))
const VideoDetail = React.lazy(() => import('./media/video/VideoDetail'))
const VideoList = React.lazy(() => import('./media/video/VideoList'))
const VideoNavBar = React.lazy(() => import('./media/video/VideoNavBar'))
const HistoryPage = React.lazy(() => import('./media/history/HistoryPage'))

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
            case "/app/media/videos":
                setSelectedTab("mediaTab")
                break
            case "/app/user":
                setSelectedTab("meTab")
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
            <Route exact path="/app/media/videos/:id"  >
                <VideoDetail />
            </Route>
            <Route exact path="/app/media/history"  >
                <HistoryPage />
            </Route>
            <Route exact path="/app/user/profile"  >
                <UpdateProfileForm user={user} />
            </Route>
            <Route exact path="/app/user/account"  >
                <UpdatePasswordForm />
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
                    <HomeIndex />
                </TabBar.Item>
                <TabBar.Item
                    title="多媒体"
                    key="media"
                    icon={<PlaySquareOutlined />}
                    selectedIcon={<PlaySquareOutlined style={{ color: "#85dbf5" }} />}
                    selected={selectedTab === 'mediaTab'}
                    onPress={() => {
                        setSelectedTab('mediaTab')
                        history.push('/app/media/videos')
                    }}
                    data-seed="logId"
                >
                    <Route exact path="/app/media/videos">
                        <VideoNavBar />
                        <VideoList />
                    </Route>
                </TabBar.Item>
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
