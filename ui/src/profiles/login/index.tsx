import "./index.less"
import { useHistory } from 'react-router-dom'
import { Form, Input, Button, message } from 'antd'
import React, { useEffect } from 'react'
import { LOGIN } from 'src/consts/user.gpl'
import { useMutation } from '@apollo/react-hooks'
import { GesturePasswordKey } from "src/consts/consts"
import { AppPath } from "src/consts/path"

const layout = {
    labelCol: { span: 8 },
    wrapperCol: { span: 16 },
}
const tailLayout = {
    wrapperCol: { offset: 8, span: 16 },
}

const Login = () => {
    const history = useHistory()
    const [login, { data }] = useMutation(LOGIN)

    useEffect(() => {
        if (data) {
            localStorage.setItem("accessToken", data.login.accessToken)
            localStorage.setItem("refreshToken", data.login.refreshToken)
            history.push(AppPath.HOME)
        }
    }, [data, history])

    useEffect(() => {
        const token = localStorage.getItem('accessToken')
        if (token) {
            if (localStorage.getItem(GesturePasswordKey)) {
                history.push(AppPath.GESTURE_LOGIN)
            } else {
                history.push(AppPath.HOME)
            }
        }
    }, [history])

    const onFinish = async (values: any) => {
        await login({
            variables: {
                "username": values.username,
                "password": values.password
            }
        })
    }

    const onFinishFailed = (errorInfo: any) => {
        console.log('Failed:', errorInfo)
        message.error(errorInfo)
    }
    return (
        <div className="login-background">
            <div className="login-form">
                <div className={'title'}>个人数据中心</div>
                <Form
                    {...layout}
                    name="basic"
                    initialValues={{ remember: true }}
                    onFinish={onFinish}
                    onFinishFailed={onFinishFailed}
                >
                    <Form.Item
                        label="用户名"
                        name="username"
                        rules={[{ required: true, message: '请输入用户名!' }]}
                    >
                        <Input />
                    </Form.Item>

                    <Form.Item
                        label="密码"
                        name="password"
                        rules={[{ required: true, message: '请输入密码!' }]}
                    >
                        <Input.Password />
                    </Form.Item>
                    <Form.Item {...tailLayout}>
                        <Button htmlType="submit" className="login-submit">
                            登录
                        </Button>
                    </Form.Item>
                </Form>
            </div>
        </div>
    )
}

export default Login
