import { useHistory } from 'react-router-dom'
import { Form, Input, Button, Select, DatePicker, message } from 'antd'
import React, { useState, useEffect, FC } from 'react'
import { GenderMap } from "src/consts/consts"
import { Uploader } from "src/components/Uploader"
import dayjs from 'dayjs'
import { useMutation } from '@apollo/react-hooks'
import { NavBar, Icon } from 'antd-mobile'
import { NewUser } from 'src/models/user'
import { UPDATE_PROFILE } from 'src/gqls/user/mutation'


interface IUpdateProfileFormProps {
    user: NewUser
}

export const UpdateProfileForm: FC<IUpdateProfileFormProps> = ({
    user
}) => {
    const history = useHistory()

    const [url, setUrl] = useState("")
    const [loading, setLoading] = useState(false)
    const [updateProfile] = useMutation(UPDATE_PROFILE)

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
        const data = await updateProfile({
            variables: {
                "input": {
                    "avatar": values.avatar === "" ? undefined : values.avatar,
                    "gender": values.gender,
                    "birthDate": values.birthDate ? values.birthDate.unix() : 0,
                    "ip": values.ip,
                }
            }
        })
        setLoading(false)
        if (!data.errors) {
            message.success("更新个人资料成功")
            history.goBack()
        }
    }

    const onFinish = (values: any) => {
        form.setFieldsValue({
            "avatar": url
        })
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
    useEffect(() => {
        form.setFieldsValue({
            name: user?.name,
            gender: user?.gender,
            birthDate: dayjs(user?.birthDate * 1000),
            ip: user?.ip
        })
    }, [form, user])
    return (
        <div style={{ height: "100%" }}>
            <NavBar
                mode="light"
                icon={<Icon type="left" />}
                onLeftClick={() => history.goBack()}
            >修改个人资料</NavBar>
            <Form
                {...layout}
                form={form}
                style={{ padding: 20 }}
                layout="horizontal"
                name="updateProfileForm"
                initialValues={{

                }}
                onFinish={onFinish}
                onFinishFailed={onFinishFailed}
            >
                <Form.Item
                    name="name"
                    label="名称"
                >
                    <Input disabled={true} />
                </Form.Item>
                <Form.Item
                    name="gender"
                    label="性别"
                    hasFeedback
                >
                    <Select placeholder="请选择性别!">
                        {genderOptions}
                    </Select>
                </Form.Item>
                <Form.Item name="avatar" label="头像">
                    <Uploader
                        fileLimit={1}
                        bucketName="image"
                        validFileTypes={["image/jpeg", "image/png", "image/webp"]}
                        setURL={setUrl}
                    />
                </Form.Item>
                <Form.Item
                    name="birthDate"
                    label="出生日期"
                >
                    <DatePicker />
                </Form.Item>
                <Form.Item
                    name="ip"
                    label="ip"
                >
                    <Input />
                </Form.Item>
                <Form.Item {...tailLayout}>
                    <Button type="primary" htmlType="submit" loading={loading}>
                        更新资料
                </Button>
                </Form.Item>
            </Form>
        </div >
    )
}
