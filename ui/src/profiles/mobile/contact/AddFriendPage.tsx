import { Avatar } from 'antd'
import { Button, Card, Icon, List, NavBar, WhiteSpace, WingBlank } from 'antd-mobile'
import axios from 'axios'
import React, { useEffect, useState } from 'react'
import { useHistory, useLocation } from 'react-router-dom'
import "src/styles/list.less"
import { ManOutlined, WomanOutlined } from '@ant-design/icons';
export const AddFriendPage = () => {
    const history = useHistory()
    const [data, setData] = useState({
        "uid": "",
        "host": "",
        "name": "",
        "avatar": "",
        "gender": undefined
    })
    const location = useLocation()
    const query = new URLSearchParams(location.search)
    const url = atob(query.get("url") || "")
    useEffect(() => {
        let uid: string = ""
        let host: string = ""
        const arr = url.split("/", -1)
        if (arr.length === 5) {
            uid = arr[4]
            host = arr[2]
        }
        axios.get(url)
            .then(function (response) {
                setData({
                    "uid": uid,
                    "host": host,
                    "name": response.data.name,
                    "gender": response.data.gender,
                    "avatar": response.data.avatar
                })
            })
            .catch(function (error) {
                console.log(error);
            })
    }, [url])
    return (<div style={{ height: "100%", textAlign: "center" }}>
        <NavBar
            mode="light"
            icon={<Icon type="left" />}
            onLeftClick={() => history.goBack()}
        >添加好友</NavBar>
        <WingBlank size="lg">
            <WhiteSpace size="lg" />
            <Card>
                <Card.Header
                    thumb={
                        <div style={{ display: "flex", flexDirection: "row" }}>
                            <Avatar style={{
                                backgroundColor: "#00a2ae",
                                marginLeft: 20,
                                marginRight: 20
                            }} size={80} gap={1} src={data.name} >{data.name}</Avatar>
                            <div style={{ display: "flex", padding: 20 }}>
                                <div style={{ fontSize: 36 }}> {data.name}</div>
                                <div style={{ display: "flex", justifyContent: "center", alignItems: "center", paddingLeft: 10 }}>
                                    {data.gender === 0 ? <ManOutlined style={{ color: "blue" }} /> : <WomanOutlined style={{ color: "pink" }} />}
                                </div>
                            </div>
                        </div>}
                />
                <Card.Body>
                    <List >
                        <List.Item multipleLine extra={data.host} style={{ flexBasis: "80%" }}>
                            主机
                        </List.Item>
                        <List.Item multipleLine extra={data.uid} style={{ flexBasis: "80%" }}>
                            UID
                        </List.Item>
                    </List>
                    <Button type="primary">加好友</Button>
                </Card.Body>
            </Card>
            <WhiteSpace size="lg" />
        </WingBlank>
    </div>)
}