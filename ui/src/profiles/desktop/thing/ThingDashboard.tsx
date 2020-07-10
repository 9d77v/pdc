import React, { useEffect, useState } from "react"

import { message, Select } from "antd"
import { useQuery } from "@apollo/react-hooks";
import { Pie } from "../../../components/Pie";
import { THING_SERIES } from "../../../consts/thing.gql";
import { ConsumerExpenditureMap, ThingStatusMap } from "../../../consts/consts";
import { SerieData } from "../../../consts/chart";

const mapFunc = (value: SerieData) => {
    return {
        name: ConsumerExpenditureMap.get(value.name) || ThingStatusMap.get(parseInt(value.name))?.text || (value.name === "" ? "未知" : value.name),
        value: value.value
    }
}
export const ThingDashboard = () => {
    const [dynamicDimension, setDynamicDimension] = useState("consumer_expenditure")
    const { error, data } = useQuery(THING_SERIES,
        {
            variables: {
                "dimension": dynamicDimension,
                "index1": "price",
                "index2": "num",
                "status": [1, 2, 3, 4, 5]
            }
        });

    useEffect(() => {
        if (error) {
            message.error("接口请求失败！")
        }
    }, [error])

    const chartStyle = { width: 800, height: 400, padding: 10 }
    return (
        <div style={{ display: "flex", flexDirection: "column" }}>
            <Select defaultValue={dynamicDimension} onChange={(value: string) => setDynamicDimension(value)} style={{ width: 200 }}>
                <Select.Option value="consumer_expenditure">消费支出</Select.Option>
                <Select.Option value="status">状态</Select.Option>
                <Select.Option value="location">位置</Select.Option>
                <Select.Option value="purchase_platform">购买平台</Select.Option>
            </Select>
            <div style={{ display: "grid", gridTemplateColumns: "1fr 1fr", padding: 10, overflowX: "scroll" }}>
                <Pie title="现存物品金额" name="物品类别" data={data ? data.series3.map(mapFunc) : []} style={chartStyle} unit={'￥'} />
                <Pie title="现存物品数量" name="物品类别" data={data ? data.series4.map(mapFunc) : []} style={chartStyle} unit={'件'} />
            </div>
        </div>

    )
}