import { message, Upload } from 'antd';
import React, { useState } from 'react'
import { getUploadURL } from "../../stores/videostore";
import { UploadFile } from "antd/lib/upload/interface";
import axios from 'axios'
interface UploaderProps {
    bucketName: string
    validFileTypes: string[]
    setURL: (e: any) => any
}


const SingleUploader: React.FC<UploaderProps> = ({ bucketName, validFileTypes, setURL }) => {
    const [action, setAction] = useState('');
    const emptyFileList: UploadFile<any>[] = []
    const [fileList, setFileList] = useState(emptyFileList)
    const props = {
        name: 'file',
        multiple: false,
        method: "PUT" as const,
        action: action,
        accept: validFileTypes.join(","),
        fileList: fileList,
        beforeUpload: (file: File) => {
            return new Promise<void>(async (resolve, reject) => {
                const data = await getUploadURL(bucketName, file.name);
                setAction(data.data.presignedUrl)
                resolve();
            });
        },
        onSuccess: (response: any, file: any) => {
            message.success(`${file.name} 文件上传成功.`);
            const index = action.indexOf("?")
            setURL(action.substring(0, index))
            setFileList([file])
        },
        onError: (error: any) => {
            message.error(`文件上传失败.`);
            setURL('')
            setFileList([])
        },
        onChange(info: any) {
            const { status } = info.file;
            let tmpFileList: UploadFile<any>[] = [...info.fileList];
            tmpFileList = tmpFileList.slice(-1);

            tmpFileList = tmpFileList.map(file => {
                if (file.xhr) {
                    const url = file.xhr.responseURL
                    const index = url.indexOf("?")
                    file.url = url.substring(0, index)
                }
                return file;
            });

            if (status !== 'uploading') {
                setURL('')
            }
            if (status === 'done') {
                message.success(`${info.file.name} 文件上传成功.`);
                setURL(tmpFileList[0].url)

            } else if (status === 'error') {
                message.error(`${info.file.name} 文件上传失败.`);
                setURL('')
            }
            setFileList(tmpFileList)
        },
        customRequest({
            action,
            data,
            file,
            filename,
            headers,
            onError,
            onProgress,
            onSuccess,
            withCredentials,
        }: any) {
            axios
                .put(action, file, {
                    withCredentials, headers: {
                        'Content-Type': validFileTypes.join(",")
                    },
                    onUploadProgress: ({ total, loaded }) => {
                        onProgress({ percent: Math.round(loaded / total * 100).toFixed(2) }, file);
                    },
                })
                .then(({ data: respones }) => {
                    onSuccess(respones, file);
                })
                .catch(onError);
            return {
                abort() {
                    console.log('upload progress is aborted.');
                },
            };
        },
    };
    return (
        <Upload.Dragger {...props}>
            <p className="ant-upload-hint">点击或拖拽上传{validFileTypes.join(",")}文件</p>
        </Upload.Dragger>
    )
}

export { SingleUploader }