import { useQuery } from "@apollo/react-hooks"
import React, { useMemo } from "react"
import { useHistory } from "react-router-dom"
import { Card } from "src/components/Card"
import { AppPath } from "src/consts/path"
import { GET_MOBILE_HOME_DEVICES } from "src/gqls/device/query"
import CameraPicture from "src/profiles/common/device/CameraPicture"

const DeviceCameraList = () => {
    const history = useHistory()
    const { data } = useQuery(GET_MOBILE_HOME_DEVICES,
        {
            variables: {
                deviceType: 1
            }
        })

    const cards = useMemo(() => {
        if (data) {
            return data.appDeviceDashboards.edges.map((dataItem: any) => {
                const cardItems = dataItem.cameras.map((t: any, index: number) => {
                    return (<div style={{ width: "100%" }} key={index}
                        onClick={() => {
                            history.push(AppPath.DEVICE_CAMERA_DETAIL +
                                `?device_id=${t.deviceID}&device_name=${t.deviceName}`)
                        }} >
                        <div style={{ textAlign: "left", margin: 10 }}>{t.deviceName}</div>
                        <CameraPicture
                            border={"1px solid grey"}
                            minHeight={150}
                            deviceID={t.deviceID} />
                    </div>)
                })
                return (<Card
                    key={dataItem.id}
                    title={dataItem.name}
                    width={"100%"}
                    cardItems={cardItems}
                />
                )
            })
        }
        return (<div />)
    }, [data, history])

    return (
        <div style={{
            display: "flex",
            flexDirection: "column",
            justifyContent: 'center',
            alignItems: 'center',
            padding: 5
        }}>
            {cards}
        </div >
    )
}

export default DeviceCameraList
