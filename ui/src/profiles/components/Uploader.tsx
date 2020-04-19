import { message, Upload } from 'antd';
import React, { useState } from 'react'
import { getUploadURL } from "../../gqls/upload.gql";
import { UploadFile } from "antd/lib/upload/interface";
import axios from 'axios'
import crypto from 'crypto'
import { getVttFromFile, getType } from '../../utils/subtitle';

interface UploaderProps {
    fileLimit: number
    bucketName: string
    validFileTypes: string[]
    setURL: (e: any) => any
}

export const Uploader: React.FC<UploaderProps> = ({ fileLimit, bucketName, validFileTypes, setURL }) => {
    const [action, setAction] = useState('');
    const emptyFileList: UploadFile<any>[] = []
    const [fileList, setFileList] = useState(emptyFileList)
    const defultFileURLs: string[] = []
    const [fileURLs, setFileURLs] = useState(defultFileURLs)
    let isMulti = false
    if (fileLimit !== 1) {
        isMulti = true
    }
    const props = {
        name: 'file',
        multiple: isMulti,
        method: "PUT" as const,
        action: action,
        accept: validFileTypes.join(","),
        fileList: fileList,
        showUploadList: {
            showDownloadIcon: true,
            downloadIcon: 'download ',
            showRemoveIcon: true,
        },
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
        onSuccess: async (response: any, file: UploadFile) => {
            message.success(`${file.name} 文件上传成功.`);
            const index = action.indexOf("?")
            if (fileLimit === 1) {
                setURL(action.substring(0, index))
                setFileList([file])
            } else {
                const tempFileURLs = [...fileURLs, action.substring(0, index)]
                setURL(tempFileURLs)
                setFileURLs(tempFileURLs)
                file.status = 'done';
                file.response = response;
                const tmpFileList = [...fileList, file]
                setFileList(tmpFileList)
            }
        },
        onError: (error: any) => {
            message.error(`文件上传失败.`);
            if (fileLimit === 1) {
                setURL('')
            } else {
                setURL([])
            }
            setFileList([])
        },
        onChange(info: any) {
            const { status } = info.file;
            let tmpFileList: UploadFile<any>[] = [...info.fileList];
            if (fileLimit > 0) {
                tmpFileList = tmpFileList.slice(-fileLimit);
            }
            setFileList(tmpFileList)
            let fileURLS: string[] = []
            tmpFileList = tmpFileList.map(file => {
                if (file.xhr) {
                    const url = file.xhr.responseURL
                    const index = url.indexOf("?")
                    file.url = url.substring(0, index)
                    fileURLS.push(file.url ? file.url.toString() : "")
                }
                return file;
            });
            if (status !== 'uploading') {
            }
            if (status === 'done') {
                message.success(`${info.file.name} 文件上传成功s.`);
            } else if (status === 'error') {
                message.error(`${info.file.name} 文件上传失败s.`);
            }
        },
        async customRequest({
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
            await axios
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