import React from "react"
import "src/styles/button.less"
import { useHistory } from "react-router-dom"

interface IVideoSeriesSelectProps {
    data?: any[]
    videoID: number
}

const VideoSeriesSelect = (props: IVideoSeriesSelectProps) => {
    const { data, videoID } = props
    const history = useHistory()
    let seriesButtons: any[] = []
    let seriesName: string = ""
    if (data && data.length > 0 && data[0].items) {
        const items = data[0].items
        seriesName = data[0].name
        seriesButtons = items.map((value: any, index: number) => {
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
                onClick={() => { history.replace('/app/media/videos/' + value.videoID) }}
            >
                {value.alias}
            </div>
        })
    }
    return (
        <div>
            <div style={{ textAlign: "left", paddingLeft: 10 }}>
                {seriesName === "" ? "" : seriesName + "系列"}
            </div>
            <div>{seriesButtons}</div>
        </div>
    )
}

export default VideoSeriesSelect
