import React from "react"
import { AppPath } from "src/consts/path"
import "src/styles/button.less"

interface IVideoSelectProps {
    data?: any[]
    num: number
}

const VideoSelect = (props: IVideoSelectProps) => {
    const { data, num } = props
    let buttons: any[] = []
    if (data && data.length > 0) {
        buttons = data.map((value: any, index: number) => {
            const link = AppPath.VIDEO_DETAIL + "?episode_id=" + value.id
            if (index === num) {
                return <div key={"pdc-button-" + value.id} className={"pdc-button-selected"}  >
                    {value.num}
                </div>
            }
            return <div key={"pdc-button-" + value.id}
                className={"pdc-button"}
                onClick={() => { window.location.replace(link) }} >
                {value.num}
            </div>
        })
    }
    return (
        <div>
            <div style={{ textAlign: 'left', paddingLeft: 10 }}>选集</div>
            <div>{buttons}</div>
        </div>
    )
}

export default VideoSelect
