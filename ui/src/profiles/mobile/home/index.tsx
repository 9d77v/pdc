import React, { useEffect, useState, useRef } from "react"
import { Route, useLocation } from "react-router-dom";
import { GET_MOBILE_HOME_DEVICES } from "../../../consts/device.gql";
import { useQuery } from "@apollo/react-hooks";
import useWebSocket from "react-use-websocket";
import { deviceTelemetryPrefix, iotSocketURL } from "../../../utils/ws_client";
import { pb } from "../../../pb/compiled";
import "../../../style/card.less"
import { blobToArrayBuffer } from "../../../utils/file";

export default function HomeIndex() {
    const location = useLocation();
    const [dataResource, setDataResource] = useState<any[]>([])
    const [telemetryMap, setTelemetryMap] = useState<Map<number, pb.Telemetry>>(new Map<number, pb.Telemetry>())
    const updateTelemetryCallback: any = useRef();

    switch (location.pathname) {
        case "/app/home":
            break
    }

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
    } = useWebSocket(iotSocketURL, {
        onOpen: () => () => { console.log('opened') },
        shouldReconnect: (closeEvent) => true,
        share: false,
    })


    useEffect(() => {
        let telemetries: string[] = []
        for (let element of data ? data.deviceDashboards.edges : []) {
            for (let t of element.telemetries) {
                telemetries.push(deviceTelemetryPrefix + "." + t.deviceID + "." + t.telemetryID)
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

    const cards = dataResource?.map((v: any, index: number) => {
        const cardItems = v.telemetries.map((t: any) => {
            const value = t.value === null ? "-" : (t.factor * (t.value || 0)).toFixed(t.scale)
            return <div key={t.telemetryID}>{t.name}: {value}{t.unit}</div>
        })
        let width: string = "50%"
        if (index === dataResource.length - 1 && dataResource.length % 2 === 1) {
            width = "100%"
        }
        return <div key={v.id}
            className="pdc-card-home"
            style={{
                width: width,
                display: "flex",
                justifyContent: "center",
                alignItems: "center",
                padding: 10,
                flexDirection: "column"
            }}>
            {v.name}
            <div style={{
                textAlign: "left",
                paddingTop: 10
            }}>
                {cardItems}
            </div>

        </div>
    })
    return (
        <Route exact path="/app/home">
            <div style={{
                display: 'flex',
                alignItems: 'center',
                height: "100%",
                opacity: 0.7,
                background: "#fff",
                border: 1,
                justifyContent: 'center',
                backgroundColor: '#eee'
            }}>
                {cards}
            </div>
        </Route>)
}