import React from "react"
import { useHistory } from "react-router-dom";
import "../../../style/card.less"
import { Grid } from "antd-mobile";
import {
    PlayCircleOutlined, DashboardOutlined,
    ContactsOutlined, LockOutlined, BookOutlined, CalculatorOutlined
} from '@ant-design/icons';

const publicData = [
    {
        "text": "视频",
        "icon": <PlayCircleOutlined style={{ fontSize: 26 }} />,
        "url": "/app/media/videos"
    }, {
        "text": "设备",
        "icon": <DashboardOutlined style={{ fontSize: 26 }} />,
        "url": "/app/device"
    }
]

const privateData = [
    {
        "text": "通讯录",
        "icon": <ContactsOutlined style={{ fontSize: 26, color: "brown" }} />,
        "url": "/app/contacts"
    }, {
        "text": "密码箱",
        "icon": <LockOutlined style={{ fontSize: 26, color: "red" }} />,
        "url": "/app/password"
    }, {
        "text": "记事本",
        "icon": <BookOutlined style={{ fontSize: 26, color: "green" }} />,
        "url": "/app/note"
    }, {
        "text": "计算器",
        "icon": <CalculatorOutlined style={{ fontSize: 26, color: "blue" }} />,
        "url": "/app/calculator"
    }
]
export default function HomeIndex() {
    const history = useHistory()
    return (
        <>
            <Grid data={publicData} columnNum={2} onClick={(item: any) => history.push(item.url)} />
            <br />
            {/* <Grid data={privateData} columnNum={4} onClick={(item: any) => history.push(item.url)} /> */}
        </>)
}