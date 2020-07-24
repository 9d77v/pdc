import React from "react"
import { Tabs } from "antd-mobile";
import { Route, useHistory, useLocation } from "react-router-dom";

const tabs = [
    { title: '首页', sub: '1' },
    { title: '视频', sub: '2' },
];

const VideoList = React.lazy(() => import('./media/VideoList'))

export default function HomeIndex() {
    let page = 0
    const location = useLocation();

    switch (location.pathname) {
        case "/app/home":
            page = 0
            break
        case "/app/media/videos":
            page = 1
    }
    const history = useHistory()
    return (
        <Tabs tabs={tabs}
            initialPage={page}
            renderTab={tab => <span>{tab.title}</span>}
            onChange={(tab: any, index: number) => {
                switch (index) {
                    case 1:
                        history.push("/app/media/videos")
                        break
                    case 0:
                        history.push("/app/home")
                        break
                }
            }}
        >
            <Route exact path="/app/home">
                <div style={{ display: 'flex', alignItems: 'center', height: "100%", justifyContent: 'center', backgroundColor: '#eee' }}>
                    欢迎使用个人数据 中心
        </div>
            </Route>
            <Route exact path="/app/media/videos">
                <VideoList />
            </Route>
        </Tabs >)
}