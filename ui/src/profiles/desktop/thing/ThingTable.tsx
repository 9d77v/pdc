import { Table, Button, message, Tag } from 'antd';
import React, { useState, useEffect } from 'react'

import { LIST_THING, ADD_THING, UPDATE_THING } from '../../../consts/thing.gql';
import { useQuery } from '@apollo/react-hooks';
import { ThingCreateForm } from './ThingCreateForm';
import { useMutation } from '@apollo/react-hooks';
import moment from 'moment';
import { ThingUpdateForm } from './ThingUpdateForm';
import { Img } from '../../../components/Img';
import { RubbishCategoryMap, ConsumerExpenditureMap, ThingStatusMap } from '../../../consts/consts';
import { TablePaginationConfig } from 'antd/lib/table';


export default function ThingTable() {
    const [visible, setVisible] = useState(false);
    const [pagination, setPagination] = useState({
        current: 1,
        pageSize: 10,
        showQuickJumper: true,
        showSizeChanger: true,
        total: 0
    })
    const [updateThingVisible, setUpdateThingVisible] = useState(false)
    const [updateThingData, setUpdateThingData] = useState({
        id: 0,
        name: "",
        num: 1,
        brandName: "",
        pics: [],
        unitPrice: 0,
        unit: "",
        specifications: "",
        category: 0,
        consumerExpenditure: "",
        location: "",
        status: 1,
        purchaseDate: 0,
        purchasePlatform: "",
        refOrderID: "",
        rubbishCategory: []
    })
    const [addThing] = useMutation(ADD_THING);
    const [updateThing] = useMutation(UPDATE_THING)
    const { loading, error, data, refetch, fetchMore } = useQuery(LIST_THING,
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

    useEffect(() => {
        if (error) {
            message.error("接口请求失败！")
        }
    }, [error])

    const onThingCreate = async (values: any) => {
        await addThing({
            variables: {
                "input": {
                    "name": values.name,
                    "num": values.num,
                    "brandName": values.brandName,
                    "pics": values.pics,
                    "unitPrice": values.unitPrice,
                    "unit": values.unit,
                    "specifications": values.specifications,
                    "category": values.category,
                    "consumerExpenditure": values.consumerExpenditure,
                    "location": values.location,
                    "status": values.status,
                    "purchaseDate": values.purchaseDate ? values.purchaseDate.unix() : 0,
                    "purchasePlatform": values.purchasePlatform,
                    "refOrderID": values.refOrderID,
                    "rubbishCategory": values.rubbishCategory
                }
            }
        });
        setVisible(false);
        await refetch()
    };

    const onThingUpdate = async (values: any) => {
        await updateThing({
            variables: {
                "input": {
                    "id": values.id,
                    "name": values.name,
                    "num": values.num,
                    "brandName": values.brandName,
                    "pics": values.pics,
                    "unitPrice": values.unitPrice,
                    "unit": values.unit,
                    "specifications": values.specifications,
                    "category": values.category,
                    "consumerExpenditure": values.consumerExpenditure,
                    "location": values.location,
                    "status": values.status,
                    "purchaseDate": values.purchaseDate ? values.purchaseDate.unix() : 0,
                    "purchasePlatform": values.purchasePlatform,
                    "refOrderID": values.refOrderID,
                    "rubbishCategory": values.rubbishCategory
                }
            }
        });
        setUpdateThingVisible(false);
        await refetch()
    };

    const onChange = async (pageConfig: TablePaginationConfig) => {
        fetchMore({
            variables: {
                page: pageConfig.current || 1,
                pageSize: pageConfig.pageSize || 10
            },
            updateQuery: (previousResult, { fetchMoreResult }) => {
                const newEdges = fetchMoreResult ? fetchMoreResult.Things.edges : [];
                const totalCount = fetchMoreResult ? fetchMoreResult.Things.totalCount : 0;
                setPagination({
                    ...pagination,
                    current: pageConfig.current || 1,
                    pageSize: pageConfig.pageSize || 10
                })
                return newEdges.length
                    ? {
                        Things: {
                            __typename: previousResult.Things.__typename,
                            edges: newEdges,
                            totalCount
                        }
                    }
                    : previousResult;
            }
        })
    }

    const columns = [
        { title: '名称', dataIndex: 'name', key: 'name', fixed: 'left' as const, width: 250 },
        { title: '数量', dataIndex: 'num', key: 'num', fixed: 'left' as const, width: 80, },
        {
            title: '图片', dataIndex: 'pics', key: 'pics', fixed: 'left' as const, width: 80,
            render: (value: string) => <Img src={value ? value[0] : ''} width={40} height={53.5} />
        },
        {
            title: '单价', dataIndex: 'unitPrice', key: 'unitPrice', fixed: 'left' as const,
            width: 80,
            render: (value: number, record: any) => {
                return <Tag color={"gold"} >
                    {`￥` + value}
                </Tag>
            }
        },
        {
            title: '消费支出', dataIndex: 'consumerExpenditure', key: 'consumerExpenditure',
            width: 120,
            render: (value: string) => {
                return <Tag color={"cyan"} >
                    {ConsumerExpenditureMap.get(value)}
                </Tag>
            }
        },
        {
            title: '位置', dataIndex: 'location', key: 'location', width: 200
        },

        { title: '品牌', dataIndex: 'brandName', key: 'brandName', width: 100 },
        {
            title: '单位', dataIndex: 'unit', key: 'unit', width: 100
        },
        {
            title: '规格', dataIndex: 'specifications', key: 'specifications', width: 100
        },
        {
            title: '购买日期', dataIndex: 'purchaseDate', key: 'purchaseDate', width: 150,
            render: (value: number) => moment(value * 1000).format("YYYY年MM月DD日")
        },
        {
            title: '购买平台', dataIndex: 'purchasePlatform', key: 'purchasePlatform', width: 100,
        },
        {
            title: '关联订单号', dataIndex: 'refOrderID', key: 'refOrderID', width: 220,
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
            title: '垃圾分类', dataIndex: 'rubbishCategory', key: 'rubbishCategory',
            width: 100,
            render: (values: number[]) => {
                if (values) {
                    const tagNodes = values.map(value => {
                        return (
                            <Tag color={RubbishCategoryMap.get(value)?.color || 'default'} key={"rubbish_category_tag" + value}>
                                {RubbishCategoryMap.get(value)?.text}
                            </Tag>
                        );
                    })
                    return <div>{tagNodes}</div>
                }
                return <div />
            }
        },
        {
            title: '状态', dataIndex: 'status', key: 'status', fixed: 'right' as const, width: 100,
            render: (value: number) => {
                return <Tag color={ThingStatusMap.get(value)?.color || "default"}>
                    {ThingStatusMap.get(value)?.text}
                </Tag>
            }
        },
        {
            title: '操作', dataIndex: 'operation', key: 'operation', fixed: "right" as const, width: 120, render: (value: any, record: any) =>
                <span><Button
                    onClick={() => {
                        setUpdateThingData({
                            "id": record.id,
                            "name": record.name,
                            "num": record.num,
                            "brandName": record.brandName,
                            "pics": record.pics,
                            "unitPrice": record.unitPrice,
                            "unit": record.unit,
                            "specifications": record.specifications,
                            "category": record.category,
                            "consumerExpenditure": record.consumerExpenditure,
                            "location": record.location,
                            "status": record.status,
                            "purchaseDate": record.purchaseDate,
                            "purchasePlatform": record.purchasePlatform,
                            "refOrderID": record.refOrderID,
                            "rubbishCategory": record.rubbishCategory
                        })
                        setUpdateThingVisible(true)
                    }}>编辑物品</Button>
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
                新增物品
            </Button>
            <ThingCreateForm
                visible={visible}
                onCreate={onThingCreate}
                onCancel={() => {
                    setVisible(false);
                }}
            />
            <ThingUpdateForm
                visible={updateThingVisible}
                data={updateThingData}
                onUpdate={onThingUpdate}
                onCancel={() => {
                    setUpdateThingVisible(false);
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
                    total: data ? data.Things.totalCount : 0
                }}
                dataSource={data ? data.Things.edges : []}
            />
        </div>

    );
}
