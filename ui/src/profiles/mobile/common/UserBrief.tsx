import { Avatar } from 'antd'
import React from 'react'
import { NewUser } from '../../desktop/settings/user/UserCreateForm'
import { ManOutlined, WomanOutlined } from '@ant-design/icons';

interface IUserBriefProps {
    user: NewUser
}

export const UserBrief = (props: IUserBriefProps) => {
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
            }} size={80} gap={1} src={props.user?.avatar} >{props.user?.name}</Avatar>
            <div style={{ flex: 1, flexDirection: "column", paddingLeft: 20 }}>
                <div style={{ display: "flex" }}>
                    <div style={{ fontSize: 36 }}> {props.user?.name}</div>
                    <div style={{ display: "flex", justifyContent: "center", alignItems: "center", paddingLeft: 10 }}>
                        {props.user?.gender === 0 ? <ManOutlined style={{ color: "blue" }} /> : <WomanOutlined style={{ color: "pink" }} />}
                    </div>
                </div>
                <div style={{ fontSize: 16 }}>UID: {props.user?.id}</div>
            </div>
        </div>
    )
}










