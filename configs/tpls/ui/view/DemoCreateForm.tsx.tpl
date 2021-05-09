import { Modal, Form, Input[[.InputComponents]] } from 'antd'[[.CustomComponents]]
import { FC } from 'react'
import { I[[.Name]] } from 'src/module/[[.LowerName]]/[[.LowerName]].model'

interface I[[.Name]]CreateFormProps {
    visible: boolean
    onCreate: (values: I[[.Name]]) => void
    onCancel: () => void
}

export const [[.Name]]CreateForm: FC<I[[.Name]]CreateFormProps> = ({
    visible,
    onCreate,
    onCancel,
}) => {
    const [form] = Form.useForm()
    const layout = {
        labelCol: { span: 5 },
        wrapperCol: { span: 15 },
    }
    return (
        <Modal
            visible={visible}
            title="新增[[.ModuleName]]"
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
                        onCreate(values)
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
                name="[[.LowerName]]CreateForm"
                style={{ maxHeight: 600, overflowY: "auto" }}
                initialValues={{ }}
            >[[template "form" .]]
            </Form>
        </Modal>
    )
}
