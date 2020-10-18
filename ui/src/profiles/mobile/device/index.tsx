import React, { useEffect, useState, useRef } from "react"
import { useHistory } from "react-router-dom";
import { GET_MOBILE_HOME_DEVICES } from "../../../consts/device.gql";
import { useQuery } from "@apollo/react-hooks";
import useWebSocket from "react-use-websocket";
import { iotTelemetrySocketURL } from "../../../utils/ws_client";
import { pb } from "../../../pb/compiled";
import "../../../style/card.less"
import { blobToArrayBuffer } from "../../../utils/file";
import { Grid, Icon, NavBar } from "antd-mobile";

export default function DeviceIndex() {
    const history = useHistory()
    const [dataResource, setDataResource] = useState<any[]>([])
    const [telemetryMap, setTelemetryMap] = useState<Map<number, pb.Telemetry>>(new Map<number, pb.Telemetry>())
    const updateTelemetryCallback: any = useRef();

    const { data } = useQuery(GET_MOBILE_HOME_DEVICES,
        {
            fetchPolicy: "cache-and-network"
        })

    useEffect(() => {
        if (data) {
            const deviceDashboards = data?.deviceDashboards.edges
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
        for (let element of data ? data.deviceDashboards.edges : []) {
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
    return (
        <div style={{ height: "100%", textAlign: "center" }}>
            <NavBar
                mode="light"
                style={{ position: "fixed", width: "100%", zIndex: 10, top: 0 }}
                icon={<Icon type="left" />}
                onLeftClick={() => history.goBack()}
            >设备</NavBar>
            <div style={{ marginTop: 45 }}>
                <Grid data={dataResource}
                    columnNum={2}
                    renderItem={(dataItem: any) => {
                        const cardItems = dataItem.telemetries.map((t: any) => {
                            const value = t.value === null ? "-" : (t.factor * (t.value || 0)).toFixed(t.scale)
                            return <div key={t.telemetryID}>{t.name}: {value}{t.unit}</div>
                        })
                        return (<div key={dataItem.id}
                            className="pdc-card-home"
                            style={{
                                display: "flex",
                                alignItems: "center",
                                padding: 10,
                                opacity: 0.7,
                                height: "100%",
                                flexDirection: "column"
                            }}>
                            {dataItem.name}
                            <div style={{
                                textAlign: "left",
                                paddingTop: 10
                            }}>
                                {cardItems}
                            </div>
                        </div>)
                    }}
                /></div>
        </div>)
}