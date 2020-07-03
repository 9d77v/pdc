import React, { useState } from "react"
import { PieLineSerieData } from "../consts/chart"
import ReactEchartsCore from 'echarts-for-react/lib/core';
import moment from "moment";
import echarts from 'echarts/lib/echarts';
import 'echarts/lib/chart/pie';
import 'echarts/lib/chart/line';
import 'echarts/lib/component/tooltip';
import 'echarts/lib/component/title';
import 'echarts/lib/component/legend';
interface PieProps {
    title: string
    start: moment.Moment
    group: string
    data: PieLineSerieData
    unit: string
    style?: React.CSSProperties
}

export const PieLine: React.FC<PieProps> = ({
    title,
    start,
    group,
    data,
    unit,
    style
}) => {
    const [node, setNode] = useState<any>()
    const showData = FormatData(data, start, unit, group)
    const option = {
        legend: {},
        tooltip: {
            trigger: 'axis' as const,
            showContent: false
        },
        dataset: {
            source: showData.source
        },
        xAxis: { type: 'category' as const },
        yAxis: { gridIndex: 0 },
        grid: { top: '55%' },
        series: showData.series
    };
    const onEvents = {
        'updateAxisPointer': (event: any) => {
            const xAxisInfo = event.axesInfo[0];
            if (xAxisInfo) {
                const dimension = xAxisInfo.value + 1;
                if (node) {
                    const echarts_instance = node.getEchartsInstance();
                    echarts_instance.setOption({
                        series: {
                            id: 'pie',
                            label: {
                                formatter: unit === '￥' ? '{b}:' + unit + ' {@[' + dimension + ']} ({d}%)' : '{b} :  {@[' + dimension + ']}' + unit + ' ({d}%)'
                            },
                            encode: {
                                value: dimension,
                                tooltip: dimension
                            }
                        }
                    });
                }
            }
        }
    }
    return (<ReactEchartsCore
        echarts={echarts}
        ref={(node: any) => setNode(node)}
        option={option}
        style={style}
        onEvents={onEvents}
    />)
}

const FormatData = (data: PieLineSerieData, start: moment.Moment, unit: string, group: string | undefined) => {
    if (!data || (data.x1.length === 0 && group === "")) {
        return {
            sereis: [],
            source: []
        }
    }
    const dataMap = new Map<string, number>()
    const x2Set = new Set<string>()
    for (let i = 0; i < data.x1.length; i++) {
        dataMap.set(data.x1[i] + data.x2[i], data.y[i])
        x2Set.add(data.x2[i])

    }
    const series: any[] = []
    for (let i = 0; i < x2Set.size; i++) {
        series.push({ type: 'line', smooth: true, seriesLayoutBy: 'row' })
    }
    let dimension = ""
    if (data.x1.length > 0) {
        dimension = data.x1[data.x1.length - 1]
    }
    switch (group) {
        case "month":
            dimension = dimension.substr(8) + "日"
            break
        case "year":
            dimension = dimension.substr(5) + "月"
            break
        default:
            dimension = dimension + "年"
    }
    series.push({
        type: 'pie' as const,
        id: 'pie',
        radius: '30%',
        center: ['50%', '25%'],
        label: {
            formatter: unit === '￥' ? '{b}:' + unit + ' {@' + dimension + '} ({d}%)' : '{b} :  {@' + dimension + '}' + unit + ' ({d}%)'

        },
        encode: {
            itemName: 'product',
            value: dimension,
            tooltip: dimension
        }
    })
    const source: any[] = []
    const row: string[] = ['product']
    const rawRow: string[] = []
    switch (group) {
        case "month":
            const days = Number(moment(start.format("YYYY-MM-DD")).endOf("month").format("DD"))
            for (let i = 1; i <= days; i++) {
                const day = i < 10 ? "0" + i.toString() : i.toString()
                row.push(day + "日")
                rawRow.push(start.format("YYYY-MM-" + day))
            }
            break;
        case "year":
            for (let i = 1; i <= 12; i++) {
                const month: string = i < 10 ? "0" + i.toString() : i.toString()
                row.push(month + "月")
                rawRow.push(start.format("YYYY-") + month)
            }
            break;
        default:
            const begin: number = Number(data.x1[0])
            const end: number = Number(data.x1[data.x1.length - 1])
            for (let i = begin; i <= end; i++) {
                row.push(i.toString() + "年")
                rawRow.push(i.toString())
            }
    }
    source.push(row)
    x2Set.forEach((value: string) => {
        const row: (string | number)[] = [value === "" ? "未知" : value]
        for (let i = 0; i < rawRow.length; i++) {
            row.push(dataMap.get(rawRow[i] + value) || 0)
        }
        source.push(row)
    })
    return {
        series: series,
        source: source
    }
}
