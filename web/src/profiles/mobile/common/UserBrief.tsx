import { Avatar } from 'antd'
import React from 'react'
import { ManOutlined, WomanOutlined } from '@ant-design/icons'
import userStore from 'src/module/user/user.store'
import {
    useRecoilValue,
} from 'recoil'
interface IUserBriefProps {
    host: string
}

export const UserBrief = (props: IUserBriefProps) => {
    const currentUserInfo = useRecoilValue(userStore.currentUserInfo)
    return (
        <div style={{
            display: "flex",
            flexDirection: "row",
            height: 80,
            margin: 10
        }}
        >
            <Avatar style={{
                backgroundColor: "#00a2ae",
            }} size={80} gap={1} src={currentUserInfo.avatar} >{currentUserInfo.name}</Avatar>
            <div style={{ flex: 1, flexDirection: "column", paddingLeft: 20 }}>
                <div style={{ display: "flex" }}>
                    <div style={{ fontSize: 22 }}> {currentUserInfo.name}</div>
                    <div style={{ display: "flex", justifyContent: "center", alignItems: "center", paddingLeft: 10 }}>
                        {currentUserInfo.gender === 0 ? <ManOutlined style={{ color: "blue" }} /> : <WomanOutlined style={{ color: "pink" }} />}
                    </div>
                </div>
                <div style={{ fontSize: 16 }}>主机: {props.host}</div>
                <div style={{ fontSize: 16 }}>UID: {currentUserInfo.uid}</div>
            </div>
        </div>
    )
}
