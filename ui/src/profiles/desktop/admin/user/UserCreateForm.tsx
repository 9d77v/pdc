import { Modal, Form, Input, Select, DatePicker } from 'antd';
import React, { useState } from 'react'
import { GenderMap, RoleMap } from 'src/consts/consts';
import { Uploader } from 'src/components/Uploader';
import { NewUser } from 'src/models/user';

interface UserCreateFormProps {
    visible: boolean;
    onCreate: (values: NewUser) => void;
    onCancel: () => void;
}

export const UserCreateForm: React.FC<UserCreateFormProps> = ({
    visible,
    onCreate,
    onCancel,
}) => {
    const [form] = Form.useForm();
    const [url, setUrl] = useState('')
    const layout = {
        labelCol: { span: 5 },
        wrapperCol: { span: 15 },
    }
    let genderOptions: any[] = []
    GenderMap.forEach((value: string, key: number) => {
        genderOptions.push(<Select.Option
            value={key}
            key={'user_gender_options_' + key}>{value}</Select.Option>)
    })
    let roleOptions: any[] = []
    RoleMap.forEach((value: string, key: number) => {
        roleOptions.push(<Select.Option
            value={key}
            key={'user_role_options_' + key}>{value}</Select.Option>)
    })
    return (
        <Modal
            visible={visible}
            title="新增用户"
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
                        onCreate(values);
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
                name="userCreateForm"
                style={{ maxHeight: 600, overflowY: 'auto' }}
                initialValues={{ roleID: 3, gender: 0 }}
            >
                <Form.Item
                    name="name"
                    label="名称"
                    rules={[{ required: true, message: '请输入名称!' }]}
                >
                    <Input />
                </Form.Item>
                <Form.Item
                    name="password"
                    label="密码"
                    rules={[{ required: true, message: '请输入密码!' }]}
                >
                    <Input type="password" />
                </Form.Item>
                <Form.Item
                    name="roleID"
                    label="角色"
                    hasFeedback
                    rules={[{ required: true, message: '请选择角色!' }]}
                >
                    <Select placeholder="请选择角色!">
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