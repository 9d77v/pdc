import { Form, Input, Button, message } from 'antd';
import React, { useEffect } from 'react';
import "./LoginForm.less"
import { LOGIN } from '../../consts/user.gpl';
import { useMutation } from '@apollo/react-hooks';
import { useHistory } from 'react-router-dom';
const layout = {
    labelCol: { span: 8 },
    wrapperCol: { span: 16 },
};
const tailLayout = {
    wrapperCol: { offset: 8, span: 16 },
};

export const LoginForm = () => {

    const history = useHistory();

    const [login, { data }] = useMutation(LOGIN);

    useEffect(() => {
        if (data) {
            localStorage.setItem("accessToken", data.login.accessToken)
            localStorage.setItem("refreshToken", data.login.refreshToken)
            history.push('/app/home')
        }
    }, [data, history])
    const onFinish = async (values: any) => {
        await login({
            variables: {
                "username": values.username,
                "password": values.password
            }
        });
    };

    const onFinishFailed = (errorInfo: any) => {
        console.log('Failed:', errorInfo);
        message.error(errorInfo);
    };

    return (
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
    );
};