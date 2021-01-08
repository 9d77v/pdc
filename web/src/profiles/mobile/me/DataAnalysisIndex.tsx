import { APP_HISTORY_STATISTIC } from "src/gqls/history/query"
import { useQuery } from "@apollo/react-hooks"
import VideoStatisticCards from "src/profiles/common/video/VideoStatisticCard"
import React from "react"
import { useHistory } from "react-router-dom"
import { Icon, NavBar } from "antd-mobile"

const DataAnalysisIndex = () => {
    const { data } = useQuery(APP_HISTORY_STATISTIC,
        {
            variables: {
                sourceType: 1
            },
            fetchPolicy: "cache-and-network"
        })

    const history = useHistory()
    return <div style={{ height: "100%", overflowY: "scroll" }}>
        <NavBar
            mode="light"
            icon={<Icon type="left" />}
            onLeftClick={() => history.goBack()}
        >数据统计</NavBar>
        <VideoStatisticCards
            data={data?.appHistoryStatistic.data}
            cardTitles={["观看动画数", "观看视频数", "观看时长"]}
            nums={1}
        />
    </div>
}

export default DataAnalysisIndex
