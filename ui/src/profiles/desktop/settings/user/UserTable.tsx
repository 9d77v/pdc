import { Table, Button, message } from 'antd';
import React, { useState, useEffect } from 'react'

import { LIST_USER, ADD_USER, UPDATE_USER } from '../../../../consts/user.gpl';
import { useQuery } from '@apollo/react-hooks';
import { UserCreateForm } from './UserCreateForm';
import { useMutation } from '@apollo/react-hooks';
import moment from 'moment';
import { UserUpdateForm } from './UserUpdateForm';
import { Img } from '../../../../components/Img';
import { GenderMap, FullRoleMap } from '../../../../consts/consts';
import { TablePaginationConfig } from 'antd/lib/table';


export default function UserTable() {
    const [visible, setVisible] = useState(false);
    const [updateUserVisible, setUpdateUserVisible] = useState(false)
    const [pagination, setPagination] = useState({
        current: 1,
        pageSize: 10,
        showQuickJumper: true,
        showSizeChanger: true,
        total: 0
    })
    const [updateUserData, setUpdateUserData] = useState({
        id: 0,
        name: "",
        avatar: "",
        password: "",
        roleID: 0,
        gender: 0,
        birthDate: 0,
        ip: ""
    })
    const [addUser] = useMutation(ADD_USER);
    const [updateUser] = useMutation(UPDATE_USER)
    const { loading, error, data, refetch, fetchMore } = useQuery(LIST_USER,
        {
            variables: {
                page: pagination.current,
                pageSize: pagination.pageSize,
                sorts: [{
                    field: 'id',
                    isAsc: false
                }]
            }
        });


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
        });
        setVisible(false);
        await refetch()
    };

    const onUserUpdate = async (values: any) => {
        await updateUser({
            variables: {
                "input": {
                    "id": values.id,
                    "avatar": values.avatar,
                    "password": values.pasword,
                    "roleID": values.roleID,
                    "gender": values.gender,
                    "birthDate": values.birthDate ? values.birthDate.unix() : 0,
                    "ip": values.ip,
                }
            }
        });
        setUpdateUserVisible(false);
        await refetch()
    };

    const onChange = async (pageConfig: TablePaginationConfig) => {
        fetchMore({
            variables: {
                page: pageConfig.current || 1,
                pageSize: pageConfig.pageSize || 10
            },
            updateQuery: (previousResult, { fetchMoreResult }) => {
                const newEdges = fetchMoreResult ? fetchMoreResult.users.edges : [];
                const totalCount = fetchMoreResult ? fetchMoreResult.users.totalCount : 0;
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
                    : previousResult;
            }
        })
    }
    useEffect(() => {
        if (error) {
            message.error("接口请求失败！")
        }
    }, [error])
    const columns = [
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
            render: (value: number) => moment(value * 1000).format("YYYY年MM月DD日")
        },
        {
            title: 'ip', dataIndex: 'ip', key: 'ip', width: 100,
        },
        {
            title: '创建时间', dataIndex: 'createdAt', key: 'createdAt', width: 160,
            render: (value: number) => moment(value * 1000).format("YYYY-MM-DD HH:mm:ss")
        },
        {
            title: '更新时间', dataIndex: 'updatedAt', key: 'updatedAt', width: 160,
            render: (value: number) => moment(value * 1000).format("YYYY-MM-DD HH:mm:ss")
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
                            "birthDate": record.birthDate,
                            "ip": record.ip,
                        })
                        setUpdateUserVisible(true)
                    }}>编辑用户</Button>
                </span>
        },
    ];
    return (
        <div>
            <Button
                type="primary"
                onClick={() => {
                    setVisible(true);
                }}
                style={{ float: 'left', marginBottom: 12, zIndex: 1 }}
            >
                新增用户
            </Button>
            <UserCreateForm
                visible={visible}
                onCreate={onUserCreate}
                onCancel={() => {
                    setVisible(false);
                }}
            />
            <UserUpdateForm
                visible={updateUserVisible}
                data={updateUserData}
                onUpdate={onUserUpdate}
                onCancel={() => {
                    setUpdateUserVisible(false);
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

    );
}
