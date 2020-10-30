import React from "react"
import { Breadcrumb } from 'antd';
import { useLocation, Link, useRouteMatch } from "react-router-dom";


const breadcrumbNameMap = new Map<string, string>([
    ['/app/home', '首页'],
    ['/app/device', '设备'],
    ['/app/user', '个人设置'],
    ['/app/user/profile', '个人资料'],
    ['/app/user/account', '账户安全'],
    ['/app/media', '媒体库'],
    ['/app/media/history', '最近播放'],
    ['/app/media/videos', '视频'],
    ['/app/media/videos/:id', '播放页'],
    ['/app/thing', '物品'],
    ['/app/thing/dashboard', '物品概览'],
    ['/app/thing/things', '物品列表'],
    ['/app/thing/analysis', '物品分析'],
    ['/admin/home', '首页'],
    ['/admin/video', '视频管理'],
    ['/admin/video/video-list', '视频列表'],
    ['/admin/video/video-list/video-create', '新增视频'],
    ['/admin/video/video-series-list', '视频系列列表'],
    ['/admin/device', '设备管理'],
    ['/admin/device/device-list', '设备列表'],
    ['/admin/device/device-model-list', '设备模板列表'],
    ['/admin/device/device-dashboard-list', '设备仪表盘'],
    ['/admin/user', '用户管理'],
    ['/admin/user/user-list', '用户列表'],
])

const matchRotes = '/app/media/videos/:id'
export const AppNavigator = () => {
    const location = useLocation();
    const match = useRouteMatch(matchRotes);
    const pathSnippets = location.pathname.split('/').filter(i => i);
    const extraBreadcrumbItems = pathSnippets.map((_, index) => {
        const url = `/${pathSnippets.slice(0, index + 1).join('/')}`;
        let name = breadcrumbNameMap.get(url)
        if (index === pathSnippets.length - 1 && match) {
            name = breadcrumbNameMap.get(match.path)
        }
        return (
            <Breadcrumb.Item key={url}>
                <Link to={url}>{name}</Link>
            </Breadcrumb.Item>
        );
    });
    return (<Breadcrumb style={{ textAlign: 'left', paddingBottom: 10 }}>{extraBreadcrumbItems}</Breadcrumb>)
}
