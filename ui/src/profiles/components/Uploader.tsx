import { message, Upload } from 'antd';
import React, { useState } from 'react'
import { getUploadURL } from "../../gqls/upload.gql";
import { UploadFile } from "antd/lib/upload/interface";
import axios from 'axios'
import crypto from 'crypto'
import { getVttFromFile, getType } from '../../utils/subtitle';

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
            return new Promise<void>(async (resolve) => {
                const fileType = getType(file)
                let fileName = file.name
                if (fileType === "ass" || fileType === "srt") {
                    const vttText = await getVttFromFile(file);
                    const hash = crypto.createHash('sha256');
                    hash.update(vttText);
                    fileName = `${hash.digest('hex')}.vtt`
                }
                const data = await getUploadURL(bucketName, fileName);
                setAction(data.data.presignedUrl)
                resolve();
            });
        },
        transformFile(file: File) {
            return new Promise<File>(async (resolve) => {
                const fileType = getType(file)
                if (fileType === "ass" || fileType === "srt") {
                    const vttText = await getVttFromFile(file);
                    const blob = new Blob([vttText], {
                        type: 'text/vtt',
                    })
                    file = new File([blob], file.name, { type: 'text/vtt', lastModified: Date.now() });
                }
                resolve(file)
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