import "./index.less"
import { useHistory } from 'react-router-dom';
import { Form, Input, Button, message } from 'antd';
import React, { useEffect } from 'react';
import { LOGIN } from '../../consts/user.gpl';
import { useMutation } from '@apollo/react-hooks';

const layout = {
    labelCol: { span: 8 },
    wrapperCol: { span: 16 },
};
const tailLayout = {
    wrapperCol: { offset: 8, span: 16 },
};

export default function Login() {

    const history = useHistory();
    const [login, { data }] = useMutation(LOGIN);

    useEffect(() => {
        if (data) {
            localStorage.setItem("accessToken", data.login.accessToken)
            localStorage.setItem("refreshToken", data.login.refreshToken)
            history.push('/app/home')
        }
    }, [data, history])

    useEffect(() => {
        const token = localStorage.getItem('accessToken');
        if (token) {
            history.push('/app/home')
        }
    }, [history])

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
    );
};