import { Modal, Form, Input[[.InputComponents]] } from 'antd'[[.CustomComponents]]
import { FC, useEffect } from 'react'
import { IUpdate[[.Name]] } from 'src/module/[[.Module]]/[[.LowerName]].model'

interface I[[.Name]]UpdateFormProps {
    visible: boolean
    id: number,
    onUpdate: (values: IUpdate[[.Name]]) => void
    onCancel: () => void
}

export const [[.Name]]UpdateForm: FC<I[[.Name]]UpdateFormProps> = ({
    visible,
    id,
    onUpdate,
    onCancel,
}) => {
    const [form] = Form.useForm()
    const layout = {
        labelCol: { span: 5 },
        wrapperCol: { span: 15 },
    }
    const { error, data } = useQuery([[.TitleName]]_DETAIL,
        {
            variables: {
                searchParam: {
                    ids: [id]
                },
            },
            fetchPolicy: "cache-and-network"
        })

    useEffect(() => {
        if (error) {
            message.error("接口请求失败！")
        }
    }, [error])

    useEffect(() => {
        const obj = data?.[[.LowerName]]s?.edges[0]
        form.setFieldsValue({
            "id": obj.id,[[range .Columns]][[if eq .TSType "dayjs.Dayjs"]]
            "[[.Name]]": obj.[[.Name]] ? dayjs(obj.[[.Name]] * 1000) : undefined,[[else]]
            "[[.Name]]": obj.[[.Name]],[[end]][[end]]
        })
    }, [form, data])

    return (
        <Modal
            visible={visible}
            title="编辑[[.ModuleName]]"
            okText="确定"
            cancelText="取消"
            onCancel={
                () => {
                    onCancel()
                    form.resetFields()
                }
            }
            getContainer={false}
            onOk={() => {
                form
                    .validateFields()
                    .then((values: any) => {
                        form.resetFields()
                        onUpdate(values)
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
                name="[[.LowerName]]UpdateForm"
                style={{ maxHeight: 600, overflowY: "auto" }}
            >
                <Form.Item
                    name="id"
                    noStyle
                >
                    <Input hidden />
                </Form.Item>[[template "form" .]]
            </Form>
        </Modal>
    )
}
