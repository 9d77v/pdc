import React from "react"
import { useHistory } from "react-router-dom"
import "src/styles/card.less"
import { Grid } from "antd-mobile"
import {
    PlayCircleOutlined, DashboardOutlined,
    ContactsOutlined, LockOutlined, BookOutlined, CalculatorOutlined
} from '@ant-design/icons'
import { AppPath } from "src/consts/path"

interface IApp {
    text: string
    icon: JSX.Element
    url: string
}

const publicData: IApp[] = [
    {
        text: "视频",
        icon: <PlayCircleOutlined style={{ fontSize: 26 }} />,
        url: AppPath.VIDEO_SUGGEST
    }, {
        text: "设备",
        icon: <DashboardOutlined style={{ fontSize: 26 }} />,
        url: AppPath.DEVICE
    }
]

const privateData: IApp[] = [
    // {
    //     "text": "通讯录",
    //     "icon": <ContactsOutlined style={{ fontSize: 26, color: "brown" }} />,
    //     "url": "/app/contacts"
    // }, 
    // {
    //     text: "密码箱",
    //     icon: <LockOutlined style={{ fontSize: 26, color: "red" }} />,
    //     url: "/app/util/password"
    // },
    // {
    //     "text": "记事本",
    //     "icon": <BookOutlined style={{ fontSize: 26, color: "green" }} />,
    //     "url": "/app/note"
    // },
    {
        text: "计算器",
        icon: <CalculatorOutlined style={{ fontSize: 26, color: "blue" }} />,
        url: AppPath.UTIL_CALCULATOR
    }
]
export default function HomeIndex() {
    const history = useHistory()
    return (
        <>
            <Grid data={publicData} columnNum={4} onClick={(item: any) => history.push(item.url)} />
            <br />
            <Grid data={privateData} columnNum={4} onClick={(item: any) => history.push(item.url)} />
        </>)
}
