import { useQuery } from "@apollo/react-hooks"
import { Modal } from "antd"
import { useMemo, useState } from "react"
import { Card } from "src/components"
import { LivePlayer } from "src/components/videoplayer/LivePlayer"
import { GET_MOBILE_HOME_DEVICES } from "src/gqls/device/query"

const DeviceCamera = () => {
    const [currentCamera, setCurrentCamera] = useState({
        url: "",
        title: "",
        visible: false
    })

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
                    return (<div key={index}>
                        <div style={{ textAlign: "left", margin: 10 }}>{t.deviceName}</div>
                        <LivePlayer
                            url={`/hls/stream${t.deviceID}.m3u8`}
                            height={200}
                            width={300}
                            autoplay={true}
                        />
                    </div>)
                })
                return (<Card
                    key={dataItem.id}
                    title={dataItem.name}
                    width={340}
                    cardItems={cardItems}
                />
                )
            })
        }
        return (<div />)
    }, [data])

    return (
        <>
            <Modal
                visible={currentCamera.visible}
                title={currentCamera.title}
                footer={null}
                destroyOnClose={true}
                getContainer={false}
                width={1020}
                onCancel={
                    () => {
                        setCurrentCamera({
                            title: "",
                            url: "",
                            visible: false,
                        })
                    }
                }
            > </Modal>
            {cards}
        </>
    )
}

export default DeviceCamera
