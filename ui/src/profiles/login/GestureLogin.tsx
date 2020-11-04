import React, { FC } from 'react'
import { useHistory } from 'react-router-dom'
import GesturePassword from '@alitajs/gesture-password-react';
import { message } from 'antd';
import { GesturePasswordKey } from 'src/consts/consts';
import bcrypt from 'bcryptjs'

const GestureLogin: FC = () => {
    const history = useHistory()
    const handleChange = (data: number[]) => {
        if (data.length < 6) {
            message.warning("请至少连接6个点")
        } else {
            const password = data.join("")
            const hashPassword = localStorage.getItem(GesturePasswordKey) || ""
            if (bcrypt.compareSync(password, hashPassword)) {
                history.push("/app/home")
            } else {
                message.error("密码错误")
            }
        }
    }

    return <div style={{
        height: "100%",
        display: "flex",
        justifyContent: "center",
        flexDirection: "column",
        alignItems: "center"
    }}>
        <span style={{ fontSize: 26, marginTop: 50, marginBottom: 20 }}>请输入手势密码</span>
        <GesturePassword width={375} height={300} onChange={handleChange} />
    </div>
}

export default GestureLogin