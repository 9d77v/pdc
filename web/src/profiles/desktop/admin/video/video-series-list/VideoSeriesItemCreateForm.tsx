import { Modal, Form, Input, Select } from 'antd'
import React, { useState, useEffect, FC } from 'react'
import { useQuery } from '@apollo/react-hooks'
import { VIDEO_COMBO } from 'src/gqls/video/query'

interface IVideoSeriesItemCreateFormProps {
    visible: boolean
    onCreate: (values: any) => void
    onCancel: () => void
    video_series_id: number,
}
const { Option } = Select
export const VideoSeriesItemCreateForm: FC<IVideoSeriesItemCreateFormProps> = ({
    visible,
    onCreate,
    onCancel,
    video_series_id
}) => {
    const [form] = Form.useForm()
    const [value, setValue] = useState(0)
    const [keyword, setKeyword] = useState("")
    const { data } = useQuery(VIDEO_COMBO,
        {
            variables: {
                searchParam: {
                    page: 1,
                    pageSize: 10,
                    keyword: keyword,
                    sorts: [{
                        field: 'id',
                        isAsc: false
                    }]
                }
            },
            fetchPolicy: "cache-and-network"
        })
    const onFinish = (values: any) => {
        console.log('Finish:', values)
    }
    useEffect(() => {
        form.setFieldsValue({
            "videoSeriesID": video_series_id,
        })
    }, [form, video_series_id])

    const layout = {
        labelCol: { span: 4 },
        wrapperCol: { span: 16 },
    }

    let timer: any
    const handleSearch = (value: string) => {
        clearTimeout(timer)
        timer = setTimeout(() => {
            setKeyword(value)
        }, 1000)
    }

    const handleChange = (value: number) => {
        setValue(value)
    }
    const options = data === undefined ? null : data.videos.edges.map((d: any) =>
        <Option key={d.value} value={d.value}>{d.text}</Option>)
    return (
        <Modal
            visible={visible}
            title="新增视频"
            okText="确定"
            cancelText="取消"
            onCancel={
                () => {
                    onCancel()
                    form.resetFields(["videoID", "alias"])
                }
            }
            getContainer={false}
            onOk={() => {
                form
                    .validateFields()
                    .then((values: any) => {
                        onCreate(values)
                        form.resetFields(["videoID", "alias"])
                    })
                    .catch(info => {
                        console.log('Validate Failed:', info)
                    })
            }}
            maskClosable={false}
        >
            <Form
                {...layout}
                form={form}
                layout="horizontal"
                name="videoSeriesItemCreateForm"
                onFinish={onFinish}
            >
                <Form.Item
                    name="videoSeriesID"
                    label="视频系列"
                    noStyle
                >
                    <Input hidden />
                </Form.Item>
                <Form.Item
                    name="videoID"
                    label="视频"
                    rules={[{ required: true, message: '请选择视频!' }]}
                >
                    <Select
                        showSearch
                        value={value}
                        defaultActiveFirstOption={false}
                        showArrow={true}
                        filterOption={false}
                        onSearch={handleSearch}
                        onChange={handleChange}
                        notFoundContent={null}
                    >
                        {options}
                    </Select>
                </Form.Item>
                <Form.Item
                    name="alias"
                    label="别名"
                    rules={[{ required: true, message: '请设置视频别名!' }]}
                >
                    <Input />
                </Form.Item>
            </Form>
        </Modal >
    )
}
