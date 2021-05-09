import { Modal, Form, Input[[.InputComponents]] } from 'antd'[[.CustomComponents]]
import { FC, useEffect } from 'react'
import { IUpdate[[.Name]] } from 'src/module/[[.LowerName]]/[[.LowerName]].model'

interface I[[.Name]]UpdateFormProps {
    visible: boolean
    data: IUpdate[[.Name]],
    onUpdate: (values: IUpdate[[.Name]]) => void
    onCancel: () => void
}

export const [[.Name]]UpdateForm: FC<I[[.Name]]UpdateFormProps> = ({
    visible,
    data,
    onUpdate,
    onCancel,
}) => {
    const [form] = Form.useForm()
    const layout = {
        labelCol: { span: 5 },
        wrapperCol: { span: 15 },
    }
    useEffect(() => {
        form.setFieldsValue({
            "id": data.id,[[range .Columns]]
            "[[.Name]]": data.[[.Name]],[[end]]
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
