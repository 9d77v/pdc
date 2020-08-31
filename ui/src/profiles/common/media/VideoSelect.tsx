import React from "react"
import "../../../style/button.less"

interface IVideoSelectProps {
    data?: any[]
    num: number
    setNum: (num: number) => void
}

const VideoSelect = (props: IVideoSelectProps) => {
    const { data, num, setNum } = props
    let buttons: any[] = []
    if (data && data.length > 0) {
        buttons = data.map((value: any, index: number) => {
            if (index === num) {
                return <div key={"pdc-button-" + value.id} className={"pdc-button-selected"}  >
                    {value.num}
                </div>
            }
            return <div key={"pdc-button-" + value.id}
                className={"pdc-button"}
                onClick={() => { setNum(index) }} >
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
