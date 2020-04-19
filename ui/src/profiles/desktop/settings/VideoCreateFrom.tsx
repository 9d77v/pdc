import { Modal, Form, Input, Switch, DatePicker } from 'antd';
import React, { useState } from 'react'
import { Uploader } from '../../components/Uploader';

const { TextArea } = Input;

interface Values {
    title: string;
    description: string;
}

interface VideoCreateFormProps {
    visible: boolean;
    onCreate: (values: Values) => void;
    onCancel: () => void;
}

export const VideoCreateForm: React.FC<VideoCreateFormProps> = ({
    visible,
    onCreate,
    onCancel,
}) => {
    const [form] = Form.useForm();
    const [url, setUrl] = useState("")
    const [videoURLs, setVideoURLs] = useState([])
    const [subtitles, setSubtitles] = useState([])
    const layout = {
        labelCol: { span: 4 },
        wrapperCol: { span: 16 },
    }
    return (
        <Modal
            visible={visible}
            title="新增视频"
            okText="确定"
            cancelText="取消"
            onCancel={
                () => {
                    onCancel()
                    form.resetFields()
                    setUrl('')
                    setVideoURLs([])
                    setSubtitles([])
                }
            }
            getContainer={false}
            onOk={() => {
                form.setFieldsValue({
                    "cover": url,
                    "videoURLs": videoURLs,
                    "subtitles": subtitles
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
                setVideoURLs([])
                setSubtitles([])
            }}
        >
            <Form
                {...layout}
                form={form}
                layout="horizontal"
                name="form_in_modal"
                initialValues={{ isShow: true }}
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
                <Form.Item name="isShow" label="是否显示" valuePropName='checked'>
                    <Switch />
                </Form.Item>
                <Form.Item name="videoURLs" label="视频列表">
                    <Uploader
                        fileLimit={0}
                        bucketName="video"
                        validFileTypes={["video/mp4"]}
                        setURL={setVideoURLs}
                    />
                </Form.Item>
                <Form.Item name="subtitles" label="字幕列表">
                    <Uploader
                        fileLimit={0}
                        bucketName="vtt"
                        validFileTypes={["text/vtt", "text/ass", "text/srt"]}
                        setURL={setSubtitles}
                    />
                </Form.Item>
            </Form>
        </Modal>
    );
};