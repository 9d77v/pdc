import React from "react"
import { Breadcrumb } from 'antd';
import { useLocation, Link } from "react-router-dom";


const breadcrumbNameMap = new Map<string, string>([
    ['/settings', '系统设置'],
    ['/settings/videos', '视频管理'],
    ['/media', '媒体库'],
    ['/media/videos', '视频']
])

export const AppNavigator = () => {
    const location = useLocation();
    const pathSnippets = location.pathname.split('/').filter(i => i);
    const extraBreadcrumbItems = pathSnippets.map((_, index) => {
        const url = `/${pathSnippets.slice(0, index + 1).join('/')}`;
        return (
            <Breadcrumb.Item key={url}>
                <Link to={url}>{breadcrumbNameMap.get(url)}</Link>
            </Breadcrumb.Item>
        );
    });
    return (<Breadcrumb style={{ textAlign: 'left' }}>{extraBreadcrumbItems}</Breadcrumb>)
}