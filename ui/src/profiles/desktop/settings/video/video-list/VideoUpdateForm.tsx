import { Modal, Form, Input, Switch, DatePicker, Select } from 'antd';
import React, { useState, useEffect } from 'react'
import { Uploader } from '../../../../../components/Uploader';
import moment from 'moment';

const { TextArea } = Input;

interface Values {
    title: string;
    description: string;
}

interface UpdateVideoProps {
    title: string,
    desc: string,
    cover: string,
    pubDate: number,
    tags: string[],
    isShow: boolean,
}
interface VideoUpdateFormProps {
    visible: boolean;
    data: UpdateVideoProps,
    onUpdate: (values: Values) => void;
    onCancel: () => void;
}

export const VideoUpdateForm: React.FC<VideoUpdateFormProps> = ({
    visible,
    data,
    onUpdate,
    onCancel,
}) => {
    const [form] = Form.useForm();
    const [url, setUrl] = useState("")
    const layout = {
        labelCol: { span: 4 },
        wrapperCol: { span: 16 },
    }
    useEffect(() => {
        form.setFieldsValue({
            "title": data.title,
            "desc": data.desc,
            "tags": data.tags,
            "pubDate": moment(data.pubDate * 1000),
            "isShow": data.isShow
        })
    }, [form, data]);

    return (
        <Modal
            visible={visible}
            title="编辑视频"
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
                    "cover": url,
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
                name="videoUpdateForm"
            >
                <Form.Item
                    name="title"
                    label="标题"
                    rules={[{ required: true, message: '请输入标题!' }]}
                >
                    <Input />
                </Form.Item>
                <Form.Item name="desc" label="简介">
                    <TextArea rows={4} />
                </Form.Item>
                <Form.Item name="cover" label="封面">
                    <Uploader
                        fileLimit={1}
                        bucketName="image"
                        validFileTypes={["image/jpeg", "image/png", "image/webp"]}
                        setURL={setUrl}
                    />
                </Form.Item>
                <Form.Item name="pubDate" label="上映时间">
                    <DatePicker />
                </Form.Item>
                <Form.Item name="tags" label="标签">
                    <Select
                        mode="tags"
                        size={"large"}
                        style={{ width: '100%' }}
                    >
                    </Select>
                </Form.Item>
                <Form.Item name="isShow" label="是否显示" valuePropName='checked'>
                    <Switch />
                </Form.Item>
            </Form>
        </Modal>
    );
};