import { Icon, NavBar } from "antd-mobile"
import React from "react"
import { useHistory, useLocation } from "react-router-dom"
import { LivePlayer } from "src/components/videoplayer/LivePlayer"


const DeviceCameraDetail = () => {
    const location = useLocation()
    const query = new URLSearchParams(location.search)
    const deviceID = query.get("device_id")
    const deviceName = query.get("device_name")
    const history = useHistory()
    return (
        <div style={{ height: "100%", textAlign: "center" }}>
            <NavBar
                mode="light"
                style={{ position: "fixed", width: "100%", zIndex: 10, top: 0 }}
                icon={<Icon type="left" />}
                onLeftClick={() => history.goBack()}
            >{deviceName}</NavBar>
            <div style={{ marginTop: 45 }}>
                <LivePlayer
                    url={`/hls/stream${deviceID}.m3u8`}
                    height={231}
                    width={"100%"}
                    autoDestroy={true}
                />
            </div>
        </div >
    )
}

export default DeviceCameraDetail