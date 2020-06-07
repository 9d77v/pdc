import React from "react"
import { SerieData } from "../consts/chart"
import ReactEcharts from 'echarts-for-react';

interface PieProps {
    title: string
    name: string
    data: SerieData[]
    style?: React.CSSProperties
}

export const Pie: React.FC<PieProps> = ({
    title,
    name,
    data,
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
            formatter: '{b} : {c} ({d}%)'
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
    return (<ReactEcharts option={option} style={style} />)
}
