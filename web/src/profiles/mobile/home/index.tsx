import { useHistory } from "react-router-dom"
import "src/styles/card.less"
import "./index.less"
import { Grid } from "antd-mobile"
import {
    PlayCircleOutlined, DashboardOutlined,
    ContactsOutlined, LockOutlined, BookOutlined, CalculatorOutlined
} from '@ant-design/icons'
import { AppPath } from "src/consts/path"
import { IApp } from "src/models/app"
import { Book } from 'react-bootstrap-icons';


const publicData: IApp[] = [
    {
        text: "视频",
        icon: <PlayCircleOutlined style={{ fontSize: 26 }} />,
        url: AppPath.VIDEO_SUGGEST
    }, {
        text: "设备",
        icon: <DashboardOutlined style={{ fontSize: 26 }} />,
        url: AppPath.DEVICE_TELEMETRY
    }, {
        text: "图书",
        icon: <Book size={26} />,
        url: AppPath.BOOK_INDEX
    }
]

const utilData: IApp[] = [
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
    {
        text: "记事本",
        icon: <BookOutlined style={{ fontSize: 26, color: "green" }} />,
        url: AppPath.UTIL_NOTE
    },
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
            <div className="sub-title">应用</div>
            <Grid data={publicData} columnNum={4} onClick={(item: any) => history.push(item.url)} />
            <div className="sub-title">工具</div>
            <Grid data={utilData} columnNum={4} onClick={(item: any) => history.push(item.url)} />
        </>)
}
