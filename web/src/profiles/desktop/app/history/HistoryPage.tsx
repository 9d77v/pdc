import React, { useEffect } from "react"
import { Timeline, message } from "antd"
import { useQuery } from "@apollo/react-hooks"
import { DesktopOutlined, MobileOutlined } from '@ant-design/icons'
import { useHistory } from "react-router-dom"
import Img from "src/components/img"
import { formatRelativeTime } from "src/utils/util"
import { AppPath } from "src/consts/path"
import { LIST_HISTORY } from "src/gqls/history/query"

export default function HistoryPage() {

    const history = useHistory()
    const { error, data } = useQuery(LIST_HISTORY,
        {
            variables: {
                sourceType: 1,
                searchParam: {
                    page: 1,
                    pageSize: 50,
                }
            }
        })

    useEffect(() => {
        if (error) {
            message.error("接口请求失败！")
        }
    }, [error])

    const dataArr = data?.histories?.edges || []
    const timelineOptions = dataArr?.map((value: any, index: number) => {
        const link = AppPath.VIDEO_DETAIL + "?episode_id=" + value.subSourceID
        return <Timeline.Item key={index}  >
            <div
                style={{ display: "flex", height: 123 }}>
                <div style={{ display: "flex", alignItems: "center", width: 120 }}>
                    {formatRelativeTime(value.updatedAt)}
                </div>
                <div style={{ margin: 10, marginLeft: 20, marginRight: 20, padding: 10, cursor: "pointer" }}
                    onClick={() => history.push(link)}>
                    <Img src={value.cover} width={160} height={100}
                        currentTime={value.currentTime} remainingTime={value.remainingTime} />
                </div>
                <div style={{ borderBottom: "1px solid #e5e9ef", width: "100%" }}>
                    <div style={{ fontSize: 14, color: "#222", marginTop: 20, cursor: "pointer" }}
                        onClick={() => history.push(link)}
                    >{value.title}</div>
                    <div style={{ position: "absolute", bottom: 0 }}>
                        {value.platform === "desktop" ? <DesktopOutlined /> : <MobileOutlined />} {value.remainingTime === 0 ? "第" + value.num + "话" +
                            value.subTitle + " 已看完" : "看到第" + value.num + "话 " +
                            value.subTitle}
                    </div>
                </div>
            </div>
        </Timeline.Item>

    })
    return (<div style={{ backgroundColor: "white", padding: 50, textAlign: "left" }}>
        <Timeline mode={"left"}>
            {timelineOptions}
        </Timeline>
    </div>)
}
