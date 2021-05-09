[[define "form"]][[range .Columns]]
                <Form.Item
                    name="[[.Name]]"
                    label="[[.Comment]]"
                    rules={[{ required: true, message: '请输入[[.Comment]]!' }]}
                > [[if eq .TSType "dayjs.Dayjs"]]
                    <DatePicker />[[else if eq .TSType "string[]"]]
                    <Select
                        mode="tags"
                        size={"large"}
                        style={{ width: '100%' }}
                    >
                    </Select>[[else if eq .TSType "number"]]
                    <InputNumber />[[else]]
                    <Input />[[end]]    
                </Form.Item>[[end]][[end]]
