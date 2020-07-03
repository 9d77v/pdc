import React, { useEffect, useState } from "react"

import { message, Select, Radio, DatePicker } from "antd"
import { useQuery } from "@apollo/react-hooks";
import { THING_ANALYZE } from "../../../consts/thing.gql";
import moment from "moment";
import { ConsumerExpenditureMap, ThingStatusMap } from "../../../consts/consts";
import { RadioChangeEvent } from "antd/lib/radio";
import { PieLine } from "../../../components/PieLine";


export const ThingAnalysis = () => {
    const [dynamicDimension, setDynamicDimension] = useState("consumer_expenditure")
    const [group, setGroup] = useState<"year" | "month" | "week" | "quarter" | "date" | "time" | undefined>("month")
    const [start, setStart] = useState(moment().startOf("month"))
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
    if (data) {
        data.Series1.x2 = data.Series1.x2.map((value: string) => {
            return ConsumerExpenditureMap.get(value) || ThingStatusMap.get(parseInt(value))?.text || value
        })
        data.Series2.x2 = data.Series2.x2.map((value: string) => {
            return ConsumerExpenditureMap.get(value) || ThingStatusMap.get(parseInt(value))?.text || value
        })
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
                onChange={(value: moment.Moment | null) => { setStart(value === null ? moment() : value) }}
                picker={group}
                defaultValue={start}
                style={{ width: 136, justifyContent: 'center' }} />}
            </span>
            <div style={{ display: "grid", gridTemplateColumns: "1fr 1fr", padding: 10, overflowX: "scroll" }}>
                <PieLine title={"物品金额"} start={start} data={data ? data.Series1 : undefined} style={chartStyle} unit={'￥'} group={group || ""} />
                <PieLine title={"物品数量"} start={start} data={data ? data.Series2 : undefined} style={chartStyle} unit={'件'} group={group || ""} />
            </div>
        </div>
    )
}
