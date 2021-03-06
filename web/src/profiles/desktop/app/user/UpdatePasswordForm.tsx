import { useHistory } from 'react-router-dom'
import { Form, Input, Button, Select } from 'antd'
import React, { useState } from 'react'
import { GenderMap } from "src/consts/consts"
import { useMutation } from '@apollo/react-hooks'
import { apolloClient } from 'src/utils/apollo_client'
import { AppPath } from 'src/consts/path'
import { UPDATE_PASSWORD } from 'src/gqls/user/mutation'

export default function UpdatePasswordForm() {
    const history = useHistory()

    const [loading, setLoading] = useState(false)
    const [updatePassword] = useMutation(UPDATE_PASSWORD)

    const [form] = Form.useForm()
    const layout = {
        labelCol: { span: 8 },
        wrapperCol: { span: 16 },
    }
    const tailLayout = {
        wrapperCol: { offset: 8, span: 16 },
    }
    let genderOptions: any[] = []
    GenderMap.forEach((value: string, key: number) => {
        genderOptions.push(<Select.Option
            value={key}
            key={'user_gender_options_' + key}>{value}</Select.Option>)
    })

    const onUpdate = async (values: any) => {
        setLoading(true)
        const data = await updatePassword({
            variables: {
                "oldPassword": values.oldPassword,
                "newPassword": values.newPassword
            }
        })
        setLoading(false)
        if (!data.errors) {
            apolloClient.resetStore()
            localStorage.clear()
            history.push(AppPath.LOGIN)
        }
    }

    const onFinish = (values: any) => {
        form
            .validateFields()
            .then((values: any) => {
                onUpdate(values)
            })
            .catch(info => {
                console.log('Validate Failed:', info)
            })
    }

    const onFinishFailed = (errorInfo: any) => {
        console.log('Failed:', errorInfo)
    }

    return (
        <div style={{ display: "flex", justifyContent: "center", padding: 20, height: "100%" }}>
            <Form
                {...layout}
                form={form}
                layout="horizontal"
                name="updatePasswordForm"
                initialValues={{

                }}
                onFinish={onFinish}
                onFinishFailed={onFinishFailed}
            >
                <Form.Item
                    name="oldPassword"
                    label="旧密码"
                    rules={[{ required: true, message: '请输入旧密码!' }]}
                >
                    <Input />
                </Form.Item>
                <Form.Item
                    name="newPassword"
                    label="新密码"
                    rules={[
                        {
                            required: true,
                            message: '请输入新密码!'
                        },
                        {
                            message: '密码长度需要在10-32之间!',
                            min: 10,
                            max: 32
                        },
                        ({ getFieldValue }) => ({
                            validator(rule, value) {
                                if (!value || getFieldValue('oldPassword') !== value) {
                                    return Promise.resolve()
                                }
                                return Promise.reject('新旧密码不能相同!')
                            },
                        }),
                    ]}
                    hasFeedback
                >
                    <Input.Password />
                </Form.Item>
                <Form.Item
                    name="confirmNewPassword"
                    label="确认密码"
                    dependencies={['newPassword']}
                    rules={[
                        {
                            required: true,
                            message: '请确认密码!',
                        },
                        {
                            message: '密码长度需要在10-32之间!',
                            min: 10,
                            max: 32
                        },
                        ({ getFieldValue }) => ({
                            validator(rule, value) {
                                if (!value || getFieldValue('newPassword') === value) {
                                    return Promise.resolve()
                                }
                                return Promise.reject('两次密码不一致!')
                            },
                        }),
                    ]}
                    hasFeedback
                >
                    <Input.Password />
                </Form.Item>
                <Form.Item {...tailLayout}>
                    <Button type="primary" htmlType="submit" loading={loading}>
                        更新密码
                </Button>
                </Form.Item>
            </Form>
        </div >
    )
}
