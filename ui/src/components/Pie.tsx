import React from "react"
import { SerieData } from "../consts/chart"
import ReactEchartsCore from 'echarts-for-react/lib/core';
import echarts from 'echarts/lib/echarts';
import 'echarts/lib/chart/pie';
import 'echarts/lib/component/tooltip';
import 'echarts/lib/component/title';
import 'echarts/lib/component/legend';
interface PieProps {
    title: string
    name: string
    data: SerieData[]
    unit: string
    style?: React.CSSProperties
}

export const Pie: React.FC<PieProps> = ({
    title,
    name,
    data,
    unit,
    style
}) => {
    let legendData: string[] = []
    data.forEach((value: SerieData) => {
        legendData.push(value.name)
    })
    const option = {
        title: {
            text: title,
            left: 'center'
        },
        label: {
            trigger: 'item' as const,
            formatter: unit === '￥' ? '{b} : ' + unit + '{c} ({d}%)' : '{b} : {c}' + unit + ' ({d}%)'
        },
        legend: {
            orient: 'vertical' as const,
            left: 'left',
            data: legendData
        },
        series: [
            {
                name: name,
                type: 'pie',
                radius: '55%',
                center: ['50%', '60%'],
                data: data,
                emphasis: {
                    itemStyle: {
                        shadowBlur: 10,
                        shadowOffsetX: 0,
                        shadowColor: 'rgba(0, 0, 0, 0.5)'
                    }
                }
            }
        ]
    }
    return (<ReactEchartsCore
        echarts={echarts}
        option={option}
        style={style} />)
}
