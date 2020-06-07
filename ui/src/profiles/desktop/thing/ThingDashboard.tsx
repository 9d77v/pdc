import React, { useEffect } from "react"

import { message } from "antd"
import { useQuery } from "@apollo/react-hooks";
import { Pie } from "../../../components/Pie";
import { THING_SERIES } from "../../../consts/thing.gql";
import moment from "moment";
import { CategoryMap } from "../../../consts/category_data";
import { SerieData } from "../../../consts/chart";

const mapFunc = (value: SerieData) => {
    return {
        name: CategoryMap.get(value.name),
        value: value.value
    }
}
export const ThingDashboard = () => {

    const start = moment().startOf("year").unix()
    const { error, data } = useQuery(THING_SERIES,
        {
            variables: {
                "dimension": "category",
                "index1": "price",
                "index2": "num",
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
    const chartStyle = { width: 600, height: 400, padding: 10 }
    return (
        <div style={{ display: "grid", gridTemplateColumns: "1fr 1fr" }}>
            <Pie title="2020年物品金额" name="物品类别" data={data ? data.Series1.map(mapFunc) : []} style={chartStyle} />
            <Pie title="2020年物品数量" name="物品类别" data={data ? data.Series2.map(mapFunc) : []} style={chartStyle} />
            <Pie title="现存物品金额" name="物品类别" data={data ? data.Series3.map(mapFunc) : []} style={chartStyle} />
            <Pie title="现存物品数量" name="物品类别" data={data ? data.Series4.map(mapFunc) : []} style={chartStyle} />

        </div>
    )
}