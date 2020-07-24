import { TabBar } from 'antd-mobile';
import React, { useState } from 'react';
import { UserOutlined, HomeOutlined } from '@ant-design/icons';

import "./index.less"
import { useHistory, Route } from 'react-router-dom';
import { useQuery } from '@apollo/react-hooks';
import { GET_USER } from '../../consts/user.gpl';
import { NewUser } from '../desktop/settings/user/UserCreateForm';
import { UpdateProfileForm } from './me/UpdateFrofileForm';
import UpdatePasswordForm from './me/UpdatePasswordForm';

const MeIndex = React.lazy(() => import('./me'))
const HomeIndex = React.lazy(() => import('./home'))
const VideoDetail = React.lazy(() => import('./home/media/VideoDetail'))


export default function MobileIndex() {
    const [selectedTab, setSelectedTab] = useState("homeTab")
    const history = useHistory();
    const token = localStorage.getItem('accessToken');
    if (!token) {
        history.push('/login')
    }
    const { data } = useQuery(GET_USER);
    const user: NewUser = data?.userInfo

    return (
        <div style={{ position: 'fixed', height: '100%', width: '100%', top: 0 }}>
            <Route exact path="/app/media/videos/:id"  >
                <VideoDetail />
            </Route>
            <Route exact path="/app/user/profile"  >
                <UpdateProfileForm user={user} />
            </Route>
            <Route exact path="/app/user/account"  >
                <UpdatePasswordForm />
            </Route>
            <TabBar
                unselectedTintColor="#949494"
                tintColor="#33A3F4"
                barTintColor="white"
            >
                <TabBar.Item
                    title="首页"
                    key="home"
                    icon={<HomeOutlined />}
                    selectedIcon={<HomeOutlined style={{ color: "#85dbf5" }} />}
                    selected={selectedTab === 'homeTab'}
                    onPress={() => {
                        setSelectedTab('homeTab')
                    }}
                    data-seed="logId"
                >
                    <HomeIndex />
                </TabBar.Item>

                <TabBar.Item
                    icon={<UserOutlined />}
                    selectedIcon={<UserOutlined style={{ color: "#85dbf5" }} />}
                    title="我的"
                    key="my"
                    selected={selectedTab === 'meTab'}
                    onPress={() => {
                        setSelectedTab('meTab')
                    }}
                >
                    <MeIndex name={user ? user.name.toString() : ""} avatar={user ? user.avatar.toString() : ""} />
                </TabBar.Item>
            </TabBar>
        </div>
    );
}
