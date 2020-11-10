import React from "react"
import DeviceCards from "src/profiles/common/device/DeviceCard"

export default function DeviceIndex() {
    return (
        <>
            <div style={{ fontSize: 26, textAlign: 'left', padding: 10 }}>
                设备
            </div>
            <DeviceCards width={200} />
        </>)
}