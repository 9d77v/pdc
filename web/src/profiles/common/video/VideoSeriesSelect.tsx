import React, { useMemo } from "react"
import "src/styles/button.less"
import { AppPath } from "src/consts/path"
interface IVideoSeriesSelectProps {
    data: any
    videoID: number
}

const VideoSeriesSelect = (props: IVideoSeriesSelectProps) => {
    const { data, videoID } = props
    const series = useMemo(() => {
        let seriesName: string = ""
        let seriesButtons: any[] = []
        if (data && data.length > 0 && data[0].items) {
            const items = data[0].items
            seriesName = data[0].name
            seriesButtons = items.map((value: any, index: number) => {
                const link = AppPath.VIDEO_DETAIL + "?epid=" + value.episodeID + "&autoJump=true"
                if (videoID === Number(value.videoID)) {
                    return <div
                        key={"pdc-button-" + value.videoID}
                        className={"pdc-button-selected"} >
                        {value.alias}
                    </div>
                }
                return <div
                    key={"pdc-button-" + value.videoID}
                    className={"pdc-button"}
                    onClick={() => { window.location.replace(link) }}
                >
                    {value.alias}
                </div>
            })
        }
        return {
            seriesButtons: seriesButtons,
            seriesName: seriesName
        }
    }, [data, videoID])

    return (
        <div>
            <div style={{ textAlign: "left", paddingLeft: 10 }}>
                {series.seriesName === "" ? "" : series.seriesName + "系列"}
            </div>
            <div>{series.seriesButtons}</div>
        </div>
    )
}

export default VideoSeriesSelect
