import { Modal, Form, Input, InputNumber, Button, Typography } from 'antd'
import React, { useState, useEffect } from 'react'
import { Uploader } from '../../components/Uploader'
import { SubtitleForm } from './SubtitleForm'
import {
    DeleteOutlined
} from '@ant-design/icons';


interface EpisodeCreateFormProps {
    visible: boolean
    onCreate: (values: any) => void
    onCancel: () => void
    num: number,
}

export const EpisodeCreateForm: React.FC<EpisodeCreateFormProps> = ({
    visible,
    onCreate,
    onCancel,
    num
}) => {
    const [form] = Form.useForm()
    const [url, setUrl] = useState('')
    const [coverUrl, setCoverUrl] = useState('')
    const [subtitleVisible, setSubtitleVisible] = useState(false)

    const showSubtitleModal = () => {
        setSubtitleVisible(true)
    }

    const hideSubtitleModal = () => {
        setSubtitleVisible(false)
    }

    const onFinish = (values: any) => {
        console.log('Finish:', values)
    }

    useEffect(() => {
        if (num !== form.getFieldValue("num")) {
            form.setFieldsValue({
                "num": num
            })
        }
    }, [num, form])

    const layout = {
        labelCol: { span: 4 },
        wrapperCol: { span: 16 },
    }

    const removeSubtitle = (name: string) => {
        const subtitles = form.getFieldValue('subtitles') || []
        for (const i in subtitles) {
            if (subtitles[i].name === name) {
                subtitles.splice(i, 1)
            }
        }
        form.setFieldsValue({ subtitles: [...subtitles] })
    }
    return (
        <Modal
            visible={visible}
            title="新增分集"
            okText="确定"
            cancelText="取消"
            onCancel={
                () => {
                    onCancel()
                    form.resetFields()
                    setUrl('')
                    setCoverUrl('')
                }
            }
            getContainer={false}
            onOk={() => {
                form.setFieldsValue({
                    "url": url,
                    'cover': coverUrl
                })
                form
                    .validateFields()
                    .then((values: any) => {
                        const subtitles = form.getFieldValue('subtitles') || []
                        values.subtitles = subtitles
                        onCreate(values)
                        form.resetFields()
                    })
                    .catch(info => {
                        console.log('Validate Failed:', info)
                    })
                setUrl('')
                setCoverUrl('')
            }}
        >
            <Form.Provider
                onFormFinish={(name, { values, forms }) => {
                    if (name === 'subtitleForm') {
                        const { episodeCreateForm } = forms
                        const subtitles = episodeCreateForm.getFieldValue('subtitles') || []
                        episodeCreateForm.setFieldsValue({ subtitles: [...subtitles, values] })
                        setSubtitleVisible(false)
                    }
                }}
            >
                <Form
                    {...layout}
                    form={form}
                    layout="horizontal"
                    name="episodeCreateForm"
                    onFinish={onFinish}
                    initialValues={{ num: num }}
                >
                    <Form.Item
                        name="title"
                        label="标题"
                    >
                        <Input />
                    </Form.Item>
                    <Form.Item name="num" label="话"
                        rules={[{ required: true, message: '请输入话数!' }]}
                    >
                        <InputNumber />
                    </Form.Item>
                    <Form.Item name="desc" label="简介">
                        <Input type="textarea" />
                    </Form.Item>
                    <Form.Item name="cover" label="封面">
                        <Uploader
                            fileLimit={1}
                            bucketName="image"
                            validFileTypes={["image/jpeg", "image/png", "image/webp"]}
                            setURL={setCoverUrl}
                        />
                    </Form.Item>
                    <Form.Item
                        name="url"
                        label="上传视频"
                        rules={[{ required: true, message: '请上传视频!' }]}
                    >
                        <Uploader
                            fileLimit={1}
                            bucketName="video"
                            validFileTypes={["video/mp4"]}
                            setURL={setUrl}
                        />
                    </Form.Item>
                    <Form.Item
                        label="字幕列表"
                        shouldUpdate={(prevValues, curValues) => prevValues.subtitles !== curValues.subtitles}
                    >
                        {({ getFieldValue }) => {
                            const subtitles = getFieldValue('subtitles') || []
                            return subtitles.length ? (
                                <ul>
                                    {subtitles.map((subtitle: any, index: number) => (
                                        <li key={index}>
                                            {subtitle.name} - {subtitle.url}<Button onClick={() => removeSubtitle(subtitle.name)} icon={<DeleteOutlined />} type="link" danger></Button>
                                        </li>
                                    ))}
                                </ul>
                            ) : (
                                    <Typography.Text className="ant-form-text" type="secondary">
                                        暂无字幕
                                    </Typography.Text>
                                )
                        }}
                    </Form.Item>

                    <Button htmlType="button" style={{ margin: '0 8px' }} onClick={showSubtitleModal}>
                        添加字幕
                </Button>

                </Form>
                <SubtitleForm visible={subtitleVisible} onCancel={hideSubtitleModal} />
            </Form.Provider>
        </Modal >
    )
}