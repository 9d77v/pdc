import { Row } from "antd"

import { HISTORY_STATISTIC } from "src/gqls/history/query"
import { useQuery } from "@apollo/react-hooks"
import VideoStatisticCards from "src/profiles/common/video/VideoStatisticCard"

const VideoDataAnalysisIndex = () => {
    const { data } = useQuery(HISTORY_STATISTIC,
        {
            variables: {
                sourceType: 1
            },
            fetchPolicy: "cache-and-network"
        })

    return (
        <div>
            <Row gutter={16}>
                <VideoStatisticCards
                    data={data?.historyStatistic.data}
                    cardTitles={["观看人数", "观看动画数", "观看视频数", "观看时长"]}
                />
            </Row>
        </div>)
}

export default VideoDataAnalysisIndex
