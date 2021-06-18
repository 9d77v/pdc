import { Form, Input, Button, Select, message } from 'antd'
import { DatePicker, Uploader } from 'src/components'
import { useState, useEffect, FC } from 'react'
import { GenderMap } from "src/consts/consts"
import { useMutation } from '@apollo/react-hooks'
import { UPDATE_PROFILE } from 'src/gqls/user/mutation'
import userStore from 'src/module/user/user.store'
import {
    useRecoilValue,
} from 'recoil'

interface IUpdateProfileFormProps {
    refetch: () => void
}
const UpdateProfileForm: FC<IUpdateProfileFormProps> = (props: IUpdateProfileFormProps) => {
    const currentUserInfo = useRecoilValue(userStore.currentUserInfo)
    const [url, setUrl] = useState<string[]>([])
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
        if (!data.errors) {
            message.success("更新个人资料成功")
        }
        setLoading(false)
    }

    const onFinish = (values: any) => {
        form.setFieldsValue({
            "avatar": url[0]
        })
        form
            .validateFields()
            .then(async (values: any) => {
                await onUpdate(values)
                props.refetch()
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
            name: currentUserInfo.name,
            gender: currentUserInfo.gender,
            birthDate: currentUserInfo.birthDate,
            ip: currentUserInfo.ip
        })
    }, [form, currentUserInfo])
    return (
        <div style={{ display: "flex", justifyContent: "center", padding: 20, height: "100%" }}>
            <Form
                {...layout}
                form={form}
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

export default UpdateProfileForm
