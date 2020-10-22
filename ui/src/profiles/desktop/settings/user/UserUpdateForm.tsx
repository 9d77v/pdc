import { Modal, Form, Input, Select, DatePicker } from 'antd';
import React, { useEffect, useState } from 'react'
import { GenderMap, RoleMap, FullRoleMap } from '../../../../consts/consts';
import { Uploader } from '../../../../components/Uploader';
import dayjs from 'dayjs';

interface UpdateUserProps {
    id: number
    name: string
    avatar: string
    password: string
    roleID: number
    gender: number
    birthDate: number
    ip: string
}

interface UserUpdateFormProps {
    visible: boolean;
    data: UpdateUserProps,
    onUpdate: (values: UpdateUserProps) => void;
    onCancel: () => void;
}

export const UserUpdateForm: React.FC<UserUpdateFormProps> = ({
    visible,
    data,
    onUpdate,
    onCancel,
}) => {
    const [form] = Form.useForm();
    const [url, setUrl] = useState('')
    const layout = {
        labelCol: { span: 5 },
        wrapperCol: { span: 15 },
    }
    useEffect(() => {
        form.setFieldsValue({
            "id": data.id,
            "name": data.name,
            "password": data.password,
            "roleID": data.roleID,
            "gender": data.gender,
            "birthDate": dayjs(data.birthDate * 1000),
            "ip": data.ip,
        })
    }, [form, data]);

    let genderOptions: any[] = []
    GenderMap.forEach((value: string, key: number) => {
        genderOptions.push(<Select.Option
            value={key}
            key={'user_gender_options_' + key}>{value}</Select.Option>)
    })
    let roleOptions: any[] = []
    if (data.roleID === 1) {
        FullRoleMap.forEach((value: string, key: number) => {
            roleOptions.push(<Select.Option
                value={key}
                key={'user_role_options_' + key}>{value}</Select.Option>)
        })
    } else {
        RoleMap.forEach((value: string, key: number) => {
            roleOptions.push(<Select.Option
                value={key}
                key={'user_role_options_' + key}>{value}</Select.Option>)
        })
    }
    return (
        <Modal
            visible={visible}
            title="编辑用户"
            okText="确定"
            cancelText="取消"
            onCancel={
                () => {
                    onCancel()
                    form.resetFields()
                    setUrl('')
                }
            }
            getContainer={false}
            onOk={() => {
                form.setFieldsValue({
                    "avatar": url
                })
                form
                    .validateFields()
                    .then((values: any) => {
                        form.resetFields();
                        onUpdate(values);
                    })
                    .catch(info => {
                        console.log('Validate Failed:', info);
                    });
                setUrl('')
            }}
            maskClosable={false}
        >
            <Form
                {...layout}
                form={form}
                layout="horizontal"
                name="userUpdateForm"
                style={{ maxHeight: 600, overflowY: 'auto' }}
            >
                <Form.Item
                    name="id"
                    noStyle
                >
                    <Input hidden />
                </Form.Item>
                <Form.Item
                    name="name"
                    label="名称"
                    rules={[{ required: true, message: '请输入名称!' }]}
                >
                    <Input disabled={true} />
                </Form.Item>
                <Form.Item
                    name="password"
                    label="密码"
                >
                    <Input type="password" />
                </Form.Item>
                <Form.Item
                    name="roleID"
                    label="角色"
                    hasFeedback
                    rules={[{ required: true, message: '请选择角色!' }]}
                >
                    <Select placeholder="请选择角色!" disabled={data.roleID === 1}>
                        {roleOptions}
                    </Select>
                </Form.Item>
                <Form.Item
                    name="gender"
                    label="性别"
                    hasFeedback
                    rules={[{ required: true, message: '请选择性别!' }]}
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
            </Form>
        </Modal>
    );
};