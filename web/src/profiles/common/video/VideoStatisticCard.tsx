import { Card, Col, Statistic } from "antd"
import { useMemo } from "react"
import {
    ArrowUpOutlined, ArrowDownOutlined
} from '@ant-design/icons'

interface IVideoStatisticCardsProps {
    data: number[][]
    cardTitles: string[]
    nums?: number
}

const VideoStatisticCards = (props: IVideoStatisticCardsProps) => {
    const { data, cardTitles, nums } = props

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
        return data?.map((value: any, index: number) => {
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
                if (value[2] === 0) {
                    percent = 1
                } else {
                    percent = (value[1] / value[2]) - 1
                }
            } else if (value[1] < value[2]) {
                arrow = <ArrowDownOutlined />
                color = '#cf1322'
                percent = 1 - (value[1] / value[2])
            }
            if (cardTitles[index] === "观看时长") {
                today = formatDuration(value[0])
                yesterday = formatDuration(value[1])
            }
            return (<Col span={nums ? 24 / nums : 24 / data.length} key={index}>
                <Card title={cardTitles[index]}>
                    <Statistic title="今日" value={today.value} suffix={today.suffix}
                        precision={today.precision} />
                    <Statistic title="昨日" value={yesterday.value} suffix={yesterday.suffix}
                        precision={today.precision} />
                    <Statistic title="较前日" value={percent * 100}
                        precision={2}
                        valueStyle={{ color: color }}
                        prefix={arrow}
                        suffix="%" />
                </Card>
            </Col>)
        })
    }, [data, cardTitles, nums])
    return (
        <>
            {cards}
        </>
    )
}

export default VideoStatisticCards


