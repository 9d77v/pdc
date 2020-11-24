import React, { useEffect } from "react"
import { message } from "antd"
import { useQuery } from "@apollo/react-hooks"
import { LIST_HISTORY } from "src/consts/history.gpl"
import { DesktopOutlined, MobileOutlined } from '@ant-design/icons'
import { useHistory } from "react-router-dom"
import Img from "src/components/img"
import { formatDetailTime } from "src/utils/util"
import { NavBar, Icon } from "antd-mobile"
import { AppPath } from "src/consts/path"

export default function HistoryPage() {

    const history = useHistory()
    const { error, data } = useQuery(LIST_HISTORY,
        {
            variables: {
                sourceType: 1,
                page: 1,
                pageSize: 50,
            },
            fetchPolicy: "cache-and-network"
        })

    useEffect(() => {
        if (error) {
            message.error("接口请求失败！")
        }
    }, [error])

    const dataArr = data?.histories?.edges || []
    const options = dataArr?.map((value: any, index: number) => {
        return <div key={index} style={{ display: "flex", height: 123 }} >
            <div style={{ marginLeft: 20, marginRight: 20, padding: 10, cursor: "pointer" }}
                onClick={() => history.push(AppPath.VIDEO_DETAIL + "?video_id=" + value.sourceID)}>
                <Img src={value.cover} width={160} height={100}
                    currentTime={value.currentTime} remainingTime={value.remainingTime} />
            </div>
            <div style={{ borderBottom: "1px solid #e5e9ef", position: "relative", width: "100%" }}>
                <div style={{ fontSize: 14, color: "#222", marginTop: 20, cursor: "pointer" }}
                    onClick={() => history.push(AppPath.VIDEO_DETAIL + "?video_id=" + value.sourceID)}
                >{value.title}</div>
                <div style={{ position: "absolute", bottom: 0 }}>
                    <div>  {value.deviceType === "desktop" ? <DesktopOutlined /> : <MobileOutlined />}  {formatDetailTime(value.updatedAt)}</div>
                    <div> {value.remainingTime === 0 ? "第" + value.num + "话" +
                        value.subTitle + " 已看完" : "看到第" + value.num + "话 " +
                        value.subTitle}</div>
                </div>
            </div>
        </div>
    })
    return (<div style={{ height: "100%", overflowY: "auto" }}>
        <NavBar
            mode="light"
            icon={<Icon type="left" />}
            style={{ position: "fixed", width: "100%", zIndex: 10 }}
            onLeftClick={() => history.goBack()}
        >最近播放</NavBar>
        <div style={{ marginTop: 45 }}>
            {options}
        </div>
    </div >)
}
