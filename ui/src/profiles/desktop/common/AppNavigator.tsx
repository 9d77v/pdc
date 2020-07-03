import React from "react"
import { Breadcrumb } from 'antd';
import { useLocation, Link, useRouteMatch } from "react-router-dom";


const breadcrumbNameMap = new Map<string, string>([
    ['/app/settings', '系统设置'],
    ['/app/settings/videos', '视频管理'],
    ['/app/settings/users', '用户管理'],
    ['/app/media', '媒体库'],
    ['/app/media/videos', '视频'],
    ['/app/media/videos/:id', '播放页'],
    ['/app/thing', '物品'],
    ['/app/thing/dashboard', '物品概览'],
    ['/app/thing/things', '物品列表'],
    ['/app/thing/analysis', '物品分析']
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
