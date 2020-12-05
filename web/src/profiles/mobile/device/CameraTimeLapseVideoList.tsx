import { useQuery } from "@apollo/react-hooks"
import { Steps } from "antd-mobile"
import React, { FC, useMemo } from "react"
import { CommonPlayer } from "src/components/videoplayer/CommonPlayer"
import { GET_CAMERA_TIME_LAPSE_VIDEOS } from "src/gqls/device/query"


interface ICameraTimeLapseVideoListProps {
    deviceID: number
}

const CameraTimeLapseVideoList: FC<ICameraTimeLapseVideoListProps> = ({
    deviceID
}) => {
    const { data } = useQuery(GET_CAMERA_TIME_LAPSE_VIDEOS,
        {
            variables: {
                deviceID: deviceID
            }
        })

    const steps = useMemo(() => {
        if (data) {
            const videos = data.cameraTimeLapseVideos.edges
            return videos.map((value: any, index: number) => {
                return <Steps.Step
                    key={index}
                    icon={<span
                        style={{
                            fontSize: 12,
                            color: "grey",
                            paddingTop: 2,
                            paddingLeft: 2
                        }}
                    >{value.date.substring(5)}</span>}
                    description={
                        <div style={{ height: 180 }}><CommonPlayer
                            id={index}
                            url={value.videoURL}
                            height={180}
                            width={"calc(100% - 2px)"}
                            autoDestroy={true}
                        />
                        </div>}
                />
            })
        }
        return []
    }, [data])
    return (
        <Steps>
            {steps}
        </Steps>
    )
}

export default CameraTimeLapseVideoList