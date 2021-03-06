import { Table, Button, message } from 'antd'
import { useState, useEffect } from 'react'
import { useQuery } from '@apollo/react-hooks'
import { UserCreateForm } from './UserCreateForm'
import { useMutation } from '@apollo/react-hooks'
import dayjs from 'dayjs'
import { UserUpdateForm } from './UserUpdateForm'
import { Img } from 'src/components'
import { GenderMap, FullRoleMap } from 'src/consts/consts'
import { TablePaginationConfig } from 'antd/lib/table'
import Search from 'antd/lib/input/Search'
import { ADD_USER, UPDATE_USER } from 'src/gqls/user/mutation'
import { LIST_USER } from 'src/gqls/user/query'
import { IUpdateUser } from 'src/module/user/user.model'


export default function UserTable() {
    const [visible, setVisible] = useState(false)
    const [updateUserVisible, setUpdateUserVisible] = useState(false)
    const [pagination, setPagination] = useState({
        current: 1,
        pageSize: 10,
        showQuickJumper: true,
        showSizeChanger: true,
        total: 0
    })
    const [updateUserData, setUpdateUserData] = useState<IUpdateUser>({
        id: 0,
        name: "",
        avatar: "",
        password: "",
        roleID: 0,
        gender: 0,
        color: "",
        birthDate: undefined,
        ip: ""
    })
    const [keyword, setKeyword] = useState("")
    const [addUser] = useMutation(ADD_USER)
    const [updateUser] = useMutation(UPDATE_USER)
    const { loading, error, data, refetch, fetchMore } = useQuery(LIST_USER,
        {
            variables: {
                searchParam: {
                    page: pagination.current,
                    pageSize: pagination.pageSize,
                    keyword: keyword,
                    sorts: [{
                        field: 'id',
                        isAsc: false
                    }]
                },
            },
            fetchPolicy: "cache-and-network"
        })


    const onUserCreate = async (values: any) => {
        await addUser({
            variables: {
                "input": {
                    "name": values.name,
                    "avatar": values.avatar,
                    "password": values.password,
                    "roleID": values.roleID,
                    "gender": values.gender,
                    "birthDate": values.birthDate ? values.birthDate.unix() : 0,
                    "ip": values.ip,
                }
            }
        })
        setVisible(false)
        await refetch()
    }

    const onUserUpdate = async (values: any) => {
        await updateUser({
            variables: {
                "input": {
                    "id": values.id,
                    "avatar": values.avatar === "" ? undefined : values.avatar,
                    "password": values.password,
                    "roleID": values.roleID,
                    "gender": values.gender,
                    "birthDate": values.birthDate ? values.birthDate.unix() : 0,
                    "ip": values.ip,
                }
            }
        })
        setUpdateUserVisible(false)
        await refetch()
    }

    const onChange = async (pageConfig: TablePaginationConfig) => {
        fetchMore({
            variables: {
                searchParam: {
                    page: pageConfig.current || 1,
                    pageSize: pageConfig.pageSize || 10,
                    keyword: keyword,
                    sorts: [{
                        field: 'id',
                        isAsc: false
                    }]
                }
            },
            updateQuery: (previousResult, { fetchMoreResult }) => {
                const newEdges = fetchMoreResult ? fetchMoreResult.users.edges : []
                const totalCount = fetchMoreResult ? fetchMoreResult.users.totalCount : 0
                setPagination({
                    ...pagination,
                    current: pageConfig.current || 1,
                    pageSize: pageConfig.pageSize || 10
                })
                return newEdges.length
                    ? {
                        users: {
                            __typename: previousResult.users.__typename,
                            edges: newEdges,
                            totalCount
                        }
                    }
                    : previousResult
            }
        })
    }
    useEffect(() => {
        if (error) {
            message.error("接口请求失败！")
        }
    }, [error])
    const columns = [
        { title: 'ID', dataIndex: 'id', key: 'id' },
        { title: '名称', dataIndex: 'name', key: 'name', width: 200 },
        {
            title: '头像', dataIndex: 'avatar', key: 'avatar', width: 80,
            render: (value: string) => <Img src={value ? value : ''} width={40} height={53.5} />
        },
        {
            title: '角色', dataIndex: 'roleID', key: 'roleID', width: 200,
            render: (value: number) => FullRoleMap.get(value)
        },
        {
            title: '性别', dataIndex: 'gender', key: 'gender', width: 100,
            render: (value: number) => GenderMap.get(value)
        },
        {
            title: '出生日期', dataIndex: 'birthDate', key: 'birthDate', width: 150,
            render: (value: number) => value ? dayjs(value * 1000).format("YYYY年MM月DD日") : ""
        },
        {
            title: 'ip', dataIndex: 'ip', key: 'ip', width: 100,
        },
        {
            title: '创建时间', dataIndex: 'createdAt', key: 'createdAt', width: 160,
            render: (value: number) => dayjs(value * 1000).format("YYYY-MM-DD HH:mm:ss")
        },
        {
            title: '更新时间', dataIndex: 'updatedAt', key: 'updatedAt', width: 160,
            render: (value: number) => dayjs(value * 1000).format("YYYY-MM-DD HH:mm:ss")
        },
        {
            title: '操作', dataIndex: 'operation', key: 'operation', fixed: "right" as const, width: 120,
            render: (value: any, record: any) =>
                <span><Button
                    onClick={() => {
                        setUpdateUserData({
                            "id": record.id,
                            "name": record.name,
                            "avatar": record.avatar,
                            "password": "",
                            "roleID": record.roleID,
                            "gender": record.gender,
                            "color": record.color,
                            "birthDate": record.birthDate ? dayjs(record.birthDate * 1000) : undefined,
                            "ip": record.ip,
                        })
                        setUpdateUserVisible(true)
                    }}>编辑用户</Button>
                </span>
        },
    ]
    return (
        <div style={{ display: "flex", flexDirection: "column" }}>
            <Button
                type="primary"
                onClick={() => {
                    setVisible(true)
                }}
                style={{ float: 'left', marginBottom: 6, marginTop: 5, zIndex: 1, width: 100 }}
            >
                新增用户
            </Button>
            <Search
                placeholder="搜索"
                onSearch={value => setKeyword(value)}
                style={{ width: 200, marginBottom: 12 }}
            />
            <UserCreateForm
                visible={visible}
                onCreate={onUserCreate}
                onCancel={() => {
                    setVisible(false)
                }}
            />
            <UserUpdateForm
                visible={updateUserVisible}
                data={updateUserData}
                onUpdate={onUserUpdate}
                onCancel={() => {
                    setUpdateUserVisible(false)
                }}
            />
            <Table
                loading={loading}
                rowKey={record => record.id}
                columns={columns}
                scroll={{ x: 1300 }}
                bordered
                onChange={onChange}
                pagination={{
                    ...pagination,
                    total: data ? data.users.totalCount : 0
                }}
                dataSource={data ? data.users.edges : []}
            />
        </div>

    )
}
