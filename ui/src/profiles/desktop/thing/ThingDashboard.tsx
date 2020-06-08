import React, { useEffect, useState } from "react"

import { message, Select } from "antd"
import { useQuery } from "@apollo/react-hooks";
import { Pie } from "../../../components/Pie";
import { THING_SERIES } from "../../../consts/thing.gql";
import moment from "moment";
import { CategoryMap, ThingStatusMap } from "../../../consts/category_data";
import { SerieData } from "../../../consts/chart";

const mapFunc = (value: SerieData) => {
    return {
        name: CategoryMap.get(value.name) || ThingStatusMap.get(parseInt(value.name))?.text || (value.name === "" ? "未知" : value.name),
        value: value.value
    }
}
export const ThingDashboard = () => {

    const start = moment().startOf("year").unix()
    const year = new Date().getFullYear()
    const [dynamicDimension, setDynamicDimension] = useState("category")
    const [dynamicIndex, setDynamicIndex] = useState("num")
    const { error, data } = useQuery(THING_SERIES,
        {
            variables: {
                "dimension": dynamicDimension,
                "index": dynamicIndex,
                "start": start,
                "end": 0,
                "status1": [1, 2, 3, 4, 5, 6],
                "status2": [1, 2, 3, 4, 5]
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
            <span>
                <Select defaultValue={dynamicDimension} onChange={(value: string) => setDynamicDimension(value)} style={{ width: 200 }}>
                    <Select.Option value="category">物品类别</Select.Option>
                    <Select.Option value="status">状态</Select.Option>
                    <Select.Option value="location">位置</Select.Option>
                    <Select.Option value="purchase_platform">购买平台</Select.Option>
                </Select>
                <Select defaultValue={dynamicIndex} onChange={(value: string) => setDynamicIndex(value)} style={{ width: 200 }}>
                    <Select.Option value="num">数量</Select.Option>
                    <Select.Option value="price">金额</Select.Option>
                </Select>
            </span>
            <div style={{ display: "grid", gridTemplateColumns: "1fr 1fr", padding: 10 }}>
                <Pie title={year + "年物品分析"} name="物品类别" data={data ? data.Series1.map(mapFunc) : []} style={chartStyle} unit={dynamicIndex === "price" ? '￥' : "件"} />
                <Pie title="现存物品分析" name="物品类别" data={data ? data.Series3.map(mapFunc) : []} style={chartStyle} unit={dynamicIndex === "price" ? '￥' : "件"} />
            </div>
        </div>

    )
}