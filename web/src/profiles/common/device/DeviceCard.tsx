import React, { useEffect, useState, useRef, useMemo } from "react"
import { useQuery } from "@apollo/react-hooks";
import useWebSocket from "react-use-websocket";
import { iotTelemetrySocketURL } from "src/utils/ws_client";
import { pb } from "src/pb/compiled";
import "src/styles/card.less"
import { blobToArrayBuffer } from "src/utils/file";
import { Card } from "src/components/Card";
import { GET_MOBILE_HOME_DEVICES } from "src/gqls/device/query";

interface IDeviceCardsProps {
    width: string | number
}

const DeviceCards = (props: IDeviceCardsProps) => {
    const [dataResource, setDataResource] = useState<any[]>([])
    const [telemetryMap, setTelemetryMap] = useState<Map<number, pb.Telemetry>>(new Map<number, pb.Telemetry>())
    const updateTelemetryCallback: any = useRef()

    const { data } = useQuery(GET_MOBILE_HOME_DEVICES,
        {
            variables: {
                deviceType: 0
            }
        })

    const height = useMemo(() => {
        let h: number = 0
        dataResource.map((dataItem: any) => {
            const currentHeight = 32 + dataItem.telemetries.length * 20 + 40
            if (currentHeight > h) {
                h = currentHeight
            }
            return dataItem
        })
        return h
    }, [dataResource])

    useEffect(() => {
        if (data) {
            const deviceDashboards = data?.appDeviceDashboards.edges
            const newDeviceDashboards: any[] = []
            for (let deviceDashboard of deviceDashboards) {
                let newTelemetries: any[] = []
                for (let element of deviceDashboard.telemetries) {
                    let t: any = {
                        telemetryID: element.telemetryID,
                        deviceID: element.deviceID,
                        factor: element.factor,
                        scale: element.scale,
                        value: element.value,
                        unit: element.unit,
                        name: element.name
                    }
                    newTelemetries.push(t)
                }
                let d: any = {
                    id: deviceDashboard.id,
                    name: deviceDashboard.name,
                    telemetries: newTelemetries
                }
                newDeviceDashboards.push(d)
            }
            setDataResource(newDeviceDashboards)
        }
    }, [data])

    const {
        sendMessage,
        lastMessage,
    } = useWebSocket(iotTelemetrySocketURL, {
        onOpen: () => () => { console.log('opened') },
        shouldReconnect: (closeEvent) => true,
        queryParams: {
            'token': localStorage.getItem('accessToken') || "",
        },
        share: true,
        reconnectAttempts: 720
    })

    useEffect(() => {
        let telemetries: string[] = []
        for (let element of data ? data.appDeviceDashboards.edges : []) {
            for (let t of element.telemetries) {
                telemetries.push(t.deviceID + "." + t.telemetryID)
            }
        }
        if (telemetries.length > 0) {
            sendMessage(telemetries.join(";"))
        }
    }, [sendMessage, data])

    useEffect(() => {
        if (lastMessage) {
            blobToArrayBuffer(lastMessage.data).then((d: any) => {
                const msg = pb.Telemetry.decode(new Uint8Array(d))
                setTelemetryMap(t => t.set(msg.ID, msg))
            })
        }
    }, [lastMessage])

    const callBack = () => {
        for (let element of dataResource) {
            for (let t of element.telemetries) {
                const msg = telemetryMap.get(Number(t.telemetryID))
                if (msg) {
                    t.value = msg.Value
                    t.updatedAt = msg.ActionTime?.seconds
                }
            }
        }
        setTelemetryMap(new Map<number, pb.Telemetry>())
    }

    useEffect(() => {
        updateTelemetryCallback.current = callBack;
        return () => { };
    })

    useEffect(() => {
        const tick = () => {
            updateTelemetryCallback.current()
        }
        const timer: NodeJS.Timeout = setInterval(tick, 1000)
        return () => {
            clearInterval(timer);
        }
    }, [])

    const cards = dataResource.map((dataItem: any) => {
        const cardItems = dataItem.telemetries.map((t: any) => {
            const value = t.value === null ? "-" : (t.factor * (t.value || 0)).toFixed(t.scale)
            return <div key={t.telemetryID}>{t.name}: {value}{t.unit}</div>
        })
        return (<Card
            key={dataItem.id}
            title={dataItem.name}
            height={height}
            width={props.width}
            cardItems={cardItems}
        />
        )
    })
    return (
        <>
            {cards}
        </>
    )
}

export default DeviceCards