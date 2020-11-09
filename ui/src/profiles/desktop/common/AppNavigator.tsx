import React from "react"
import { Breadcrumb } from 'antd'
import { useLocation, Link } from "react-router-dom"
import { PathDict } from "src/consts/path"


export const AppNavigator = () => {
    const location = useLocation()
    const pathSnippets = location.pathname.split('/').filter(i => i)
    const extraBreadcrumbItems = pathSnippets.map((_, index) => {
        const url = `/${pathSnippets.slice(0, index + 1).join('/')}`
        const name = PathDict.get(url)
        return (
            <Breadcrumb.Item key={url}>
                <Link to={url}>{name}</Link>
            </Breadcrumb.Item>
        )
    })
    return (<Breadcrumb style={{ textAlign: 'left', paddingBottom: 10 }}>{extraBreadcrumbItems}</Breadcrumb>)
}
