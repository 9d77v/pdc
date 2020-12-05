import { useMutation } from '@apollo/react-hooks'
import React, { FC, useEffect, useState } from 'react'
import { CAMERA_CAPTURE } from 'src/gqls/device/mutation'

interface ICameraPictureProps {
    deviceID: number
    width?: number
    height?: number
    minWidth?: number
    minHeight?: number
    border?: string
}
const CameraPicture: FC<ICameraPictureProps> = ({
    deviceID,
    width,
    height,
    minWidth,
    minHeight,
    border
}) => {
    const [cameraCapture] = useMutation(CAMERA_CAPTURE)
    const [picture, setPicture] = useState(<div />)

    useEffect(() => {
        let isMounted = true;
        cameraCapture({
            variables: {
                "deviceID": deviceID
            }
        }).then(result => {
            if (isMounted) {
                setPicture(
                    < img src={result.data?.cameraCapture} alt="暂无图像" />
                )
            }
        })
        return (() => {
            isMounted = false
        })
    }, [cameraCapture, deviceID])

    return (
        <div style={{
            width: width, height: height,
            minWidth: minWidth, minHeight: minHeight,
            border: border
        }}>
            { picture}
        </div>
    )
}

export default CameraPicture
