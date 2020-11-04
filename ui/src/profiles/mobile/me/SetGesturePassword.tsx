import { Icon, NavBar } from 'antd-mobile'
import React, { FC } from 'react'
import { useHistory } from 'react-router-dom'
import GesturePassword from '@alitajs/gesture-password-react';
import { message } from 'antd';
import { GesturePasswordKey } from 'src/consts/consts';
import bcrypt from 'bcryptjs'

const SetGesturePassword: FC = () => {
    const history = useHistory()
    const handleChange = (data: number[]) => {
        if (data.length < 6) {
            message.warning("请至少连接6个点")
        } else {
            const password = data.join("")
            const hashPassword = bcrypt.hashSync(password, 10);
            localStorage.setItem(GesturePasswordKey, hashPassword)
            history.push("/gesture_login")
        }
    }

    return <div style={{ height: "100%" }}>
        <NavBar
            mode="light"
            icon={<Icon type="left" />}
            onLeftClick={() => history.goBack()}
        >设置手势密码</NavBar>
        <div style={{
            height: "100%",
            display: "flex",
            justifyContent: "center",
            flexDirection: "column",
            alignItems: "center",
            backgroundColor: "#fff"
        }}>
            <GesturePassword width={375} height={300} onChange={handleChange} />
        </div>
    </div>
}

export default SetGesturePassword