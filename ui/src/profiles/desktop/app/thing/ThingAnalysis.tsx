import React, { useEffect, useState } from "react"

import { message, Select, Radio } from "antd"
import { useQuery } from "@apollo/react-hooks";
import { THING_ANALYZE } from "src/consts/thing.gql";
import dayjs from "dayjs";
import { ConsumerExpenditureMap, ThingStatusMap } from "src/consts/consts";
import { RadioChangeEvent } from "antd/lib/radio";
import { PieLine } from "src/components/PieLine";
import DatePicker from "src/components/DatePicker";


export default function ThingAnalysis() {
    const [dynamicDimension, setDynamicDimension] = useState("consumer_expenditure")
    const [group, setGroup] = useState<"year" | "month" | "week" | "quarter" | "date" | "time" | undefined>("month")
    const [start, setStart] = useState(dayjs().startOf("month"))
    const { error, data } = useQuery(THING_ANALYZE,
        {
            variables: {
                "dimension": dynamicDimension,
                "index1": "price",
                "index2": "num",
                "start": start.unix(),
                'group': group || "",
            }
        });

    useEffect(() => {
        if (error) {
            message.error("接口请求失败！")
        }
    }, [error])
    let series1: any = {
        x1: [],
        x2: [],
        y: []
    }
    let series2: any = {
        x1: [],
        x2: [],
        y: []
    }
    if (data) {
        series1.x1 = data.series1.x1
        series1.x2 = data.series1.x2.map((value: string) => {
            return ConsumerExpenditureMap.get(value) || ThingStatusMap.get(parseInt(value))?.text || value
        })
        series1.y = data.series1.y

        series2.x1 = data.series2.x1
        series2.x2 = data.series2.x2.map((value: string) => {
            return ConsumerExpenditureMap.get(value) || ThingStatusMap.get(parseInt(value))?.text || value
        })
        series2.y = data.series2.y
    }
    const chartStyle = { width: 800, height: 700, padding: 10 }
    return (
        <div style={{ display: "flex", flexDirection: "column" }}>
            <Select defaultValue={dynamicDimension} onChange={(value: string) => setDynamicDimension(value)} style={{ width: 200 }}>
                <Select.Option value="consumer_expenditure">消费支出</Select.Option>
                <Select.Option value="status">状态</Select.Option>
                <Select.Option value="location">位置</Select.Option>
                <Select.Option value="purchase_platform">购买平台</Select.Option>
            </Select>
            <Radio.Group defaultValue={"month"} buttonStyle="solid" onChange={(e: RadioChangeEvent) => setGroup(e.target.value)} >
                <Radio.Button value="month">月</Radio.Button>
                <Radio.Button value="year">年</Radio.Button>
                <Radio.Button value={undefined}>全
                </Radio.Button>
            </Radio.Group>
            <span>{group === undefined ? "" : <DatePicker
                onChange={(value: dayjs.Dayjs | null) => { setStart(value === null ? dayjs() : value) }}
                picker={group}
                defaultValue={start}
                style={{ width: 136, justifyContent: 'center' }} />}
            </span>
            <div style={{ display: "grid", gridTemplateColumns: "1fr 1fr", padding: 10, overflowX: "scroll" }}>
                <PieLine title={"物品金额"} start={start} data={series1} style={chartStyle} unit={'￥'} group={group || ""} />
                <PieLine title={"物品数量"} start={start} data={series2} style={chartStyle} unit={'件'} group={group || ""} />
            </div>
        </div>
    )
}
