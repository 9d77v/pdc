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
            variables: {
                ids: [1, 2]
            },
            fetchPolicy: "cache-and-network"
        })

    useEffect(() => {
        if (data) {
            const devices = data?.devices.edges
            const newDevices: any[] = []
            for (let device of devices) {
                let newTelemetries: any[] = []
                for (let element of device.telemetries) {
                    let t: any = {
                        id: element.id,
                        factor: element.factor,
                        scale: element.scale,
                        value: element.value,
                        unit: element.unit,
                        name: element.name
                    }
                    newTelemetries.push(t)
                }
                let d: any = {
                    id: device.id,
                    telemetries: newTelemetries
                }
                newDevices.push(d)
            }
            setDataResource(newDevices)
        }
    }, [data])

    const {
        sendMessage,
        lastMessage,
    } = useWebSocket(iotSocketURL, {
        onOpen: () => () => { console.log('opened') },
        shouldReconnect: (closeEvent) => true,
        share: true,
    });
    useEffect(() => {
        sendMessage(deviceTelemetryPrefix + ".1.*;" + deviceTelemetryPrefix + ".2.*")
    }, [sendMessage])

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
                const msg = telemetryMap.get(Number(t.id))
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

    const cards = dataResource?.map((v: any) => {
        const cardItems = v.telemetries.map((item: any) => {
            const value = item.value ? (item.factor * item.value).toFixed(item.scale) : "-"
            return <div key={item.id}>{item.name}: {value}{item.unit}</div>
        })
        return <div key={v.id}
            className="pdc-card-default"
            style={{
                width: "50%",
                display: "flex",
                alignItems: "center",
                justifyContent: "center",
                flexDirection: "column"
            }}>
            {v.id === 1 ? "客厅" : "卧室"}
            {cardItems}
        </div>
    })
    return (
        <Route exact path="/app/home">
            <div style={{
                display: 'flex',
                alignItems: 'center',
                height: "100%",
                justifyContent: 'center',
                backgroundColor: '#eee'
            }}>
                {cards}
            </div>
        </Route>)
}