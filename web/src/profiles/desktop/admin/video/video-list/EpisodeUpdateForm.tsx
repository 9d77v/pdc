import { Modal, Form, Input, InputNumber, Button, Typography } from 'antd'
import { useState, useEffect, FC } from 'react'
import { Episode } from 'src/models/video'
import { SubtitleForm } from './SubtitleForm'
import {
    DeleteOutlined
} from '@ant-design/icons'
import { Uploader } from 'src/components'


interface IEpisodeUpdateFormProps {
    visible: boolean
    onUpdate: (values: Episode) => void
    onCancel: () => void
    data: Episode,
}

export const EpisodeUpdateForm: FC<IEpisodeUpdateFormProps> = ({
    visible,
    onUpdate,
    onCancel,
    data
}) => {
    const [form] = Form.useForm()
    const [url, setUrl] = useState<string[]>([])
    const [coverUrl, setCoverUrl] = useState<string[]>([])
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
        const subtitles = []
        for (const item of data.subtitles) {
            subtitles.push({
                "name": item.name,
                "url": item.url
            })
        }
        form.setFieldsValue({
            "id": data.id,
            "title": data.title,
            "desc": data.desc,
            "num": data.num,
            "subtitles": subtitles
        })
    }, [data, form])

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
            title="修改分集"
            okText="确定"
            cancelText="取消"
            onCancel={
                () => {
                    onCancel()
                    form.resetFields()
                    setUrl([])
                    setCoverUrl([])
                }
            }
            getContainer={false}
            onOk={() => {
                form.setFieldsValue({
                    "url": url[0],
                    'cover': coverUrl[0]
                })
                form
                    .validateFields()
                    .then((values: any) => {
                        const subtitles = form.getFieldValue('subtitles') || []
                        console.log(subtitles)
                        values.subtitles = subtitles
                        onUpdate(values)
                        form.resetFields()
                    })
                    .catch(info => {
                        console.log('Validate Failed:', info)
                    })
                setUrl([])
                setCoverUrl([])
            }}
            maskClosable={false}
        >
            <Form.Provider
                onFormFinish={(name, { values, forms }) => {
                    if (name === 'subtitleForm') {
                        const { episodeUpdateForm } = forms
                        const subtitles = episodeUpdateForm.getFieldValue('subtitles') || []
                        episodeUpdateForm.setFieldsValue({ subtitles: [...subtitles, values] })
                        setSubtitleVisible(false)
                    }
                }}
            >
                <Form
                    {...layout}
                    form={form}
                    layout="horizontal"
                    name="episodeUpdateForm"
                    onFinish={onFinish}
                >
                    <Form.Item
                        name="id"
                        noStyle
                    >
                        <Input hidden />
                    </Form.Item>
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
