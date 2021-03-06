import { Form, Input, Switch, Select } from 'antd'
import { DatePicker } from 'src/components'
import { forwardRef, Ref, useImperativeHandle, useState } from 'react'
import { Uploader } from 'src/components'


const { TextArea } = Input

interface Values {
    title: string
    description: string
}

interface IVideoCreateStepOneFormProps {
    onCreate?: (values: Values) => void
}

const VideoCreateStepOneForm = (props: IVideoCreateStepOneFormProps, ref: Ref<any>) => {
    const [form] = Form.useForm()
    const getForm = () => {
        return form
    }
    const getURL = () => {
        return url[0]
    }
    useImperativeHandle(ref, () => ({
        getForm,
        getURL
    }))

    const [url, setUrl] = useState<string[]>([])
    const layout = {
        labelCol: { span: 4 },
        wrapperCol: { span: 16 },
    }

    return (
        <Form
            {...layout}
            form={form}
            layout="horizontal"
            name="videoCreateForm"
            initialValues={{
                isShow: true,
                isHideOnMobile: false,
                subtitle_lang: "简体中文",
                theme: ""
            }}
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
            <Form.Item name="isHideOnMobile" label="是否手机隐藏" valuePropName='checked'>
                <Switch />
            </Form.Item>
            <Form.Item name="theme" label="主题" >
                <Select onChange={() => { }}>
                    <Select.Option value="">默认</Select.Option>
                    <Select.Option value="vjs-theme-lemon">柠檬</Select.Option>
                </Select>
            </Form.Item>
        </Form>
    )
}

export default forwardRef(VideoCreateStepOneForm)
