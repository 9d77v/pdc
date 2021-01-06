import { Card, Col, Row, Statistic } from "antd"
import React, { useMemo } from "react"

import {
    ArrowUpOutlined, ArrowDownOutlined
} from '@ant-design/icons'
import { HISTORY_STATISTIC } from "src/gqls/history/query"
import { useQuery } from "@apollo/react-hooks"

const VideoDataAnalysisIndex = () => {
    const { data } = useQuery(HISTORY_STATISTIC,
        {
            variables: {
                source_type: 1
            },
            fetchPolicy: "cache-and-network"
        })
    const formatDuration = (time: number) => {
        let suffix = ""
        if (time < 60) {
            suffix = "秒"
        } else if (time < 3600) {
            suffix = "分钟"
            time /= 60
        } else {
            suffix = "小时"
            time /= 3600
        }
        return {
            value: time,
            suffix: suffix,
            precision: 2
        }
    }
    const cards = useMemo(() => {
        if (data) {
            const showData = data.historyStatistic.data
            const cardTitles: string[] = ["观看人数", "观看动画数", "观看视频数", "观看时长"]
            return showData.map((value: any, index: number) => {
                let arrow = null
                let percent = 0
                let color = "black"
                let today = {
                    value: value[0],
                    suffix: "",
                    precision: 0
                }
                let yesterday = {
                    value: value[1],
                    suffix: "",
                    precision: 0
                }
                if (value[1] > value[2]) {
                    arrow = <ArrowUpOutlined />
                    color = '#3f8600'
                    if (value[1] === 0) {
                        percent = 100
                    } else {
                        percent = value[2] / value[1]
                    }
                } else if (value[1] < value[2]) {
                    arrow = <ArrowDownOutlined />
                    color = '#3f8600'
                    if (value[2] === 0) {
                        percent = 100
                    } else {
                        percent = value[1] / value[2]
                    }
                }
                if (index === 3) {
                    today = formatDuration(value[0])
                    yesterday = formatDuration(value[1])
                }
                return (<Col span={6} key={index}>
                    <Card title={cardTitles[index]}>
                        <Statistic title="今日" value={today.value} suffix={today.suffix} precision={today.precision} />
                        <Statistic title="昨日" value={yesterday.value} suffix={yesterday.suffix} precision={today.precision} />
                        <Statistic title="较前日" value={percent}
                            precision={2}
                            valueStyle={{ color: color }}
                            prefix={arrow}
                            suffix="%" />
                    </Card>
                </Col>)
            })
        }
        return null
    }, [data])
    return (
        <div>
            <Row gutter={16}>
                {cards}
            </Row>
        </div>)
}

export default VideoDataAnalysisIndex
