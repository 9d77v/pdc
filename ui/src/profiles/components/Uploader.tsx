import { message, Upload } from 'antd';
import React, { useState, useEffect } from 'react'
import { getUploadURL } from "../../gqls/upload.gql";
import { UploadFile } from "antd/lib/upload/interface";
import axios from 'axios'
import crypto from 'crypto'
import { getVttFromFile, getType } from '../../utils/subtitle';
import { getTextFromFile } from '../../utils/file';

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
    const emptyFile: UploadFile = { uid: "", size: 0, name: "", type: "" }
    const [succeedFile, setSucceedFile] = useState(emptyFile)
    let isMulti = false
    if (fileLimit !== 1) {
        isMulti = true
    }
    const getAccept = (validFileTypes: string[]) => {
        let acceptFileTypes: string[] = []
        for (let validFileType of validFileTypes) {
            if (validFileType) {
                const t = validFileType.split("/")[1]
                acceptFileTypes.push(validFileType)
                acceptFileTypes.push("." + t)
            }
        }
        return acceptFileTypes.join(',')
    }
    const accept = getAccept(validFileTypes)

    const sortFile = (a: UploadFile, b: UploadFile) => {
        return parseInt(a.uid.split("-").pop() || '') - parseInt(b.uid.split("-").pop() || '')
    }
    useEffect(() => {
        let all_done = true
        for (let file of fileList) {
            if (file.uid === succeedFile.uid && file.status === 'uploading') {
                file.status = 'done'
            }
            if (file.status !== 'done') {
                all_done = false
            }
        }
        if (all_done) {
            let fileURLs: string[] = []
            for (let tmpFile of fileList) {
                let obj: any = tmpFile.originFileObj
                fileURLs.push(obj.url)
            }
            if (fileLimit === 1) {
                if (fileURLs.length === 1) {
                    setURL(fileURLs[0])
                } else {
                    setURL("")
                }
            } else {
                setURL(fileURLs)
            }
        }
    }, [fileList, succeedFile, setURL, fileLimit]);
    const props = {
        name: 'file',
        multiple: isMulti,
        method: "PUT" as const,
        action: action,
        accept: accept,
        fileList: fileList,
        showUploadList: {
            showPreviewIcon: true,
            showDownloadIcon: true,
            // downloadIcon: 'download ',
            showRemoveIcon: true,
        },
        beforeUpload: (file: File) => {
            return new Promise<void>(async (resolve) => {
                let fileType = getType(file)
                let fileName = file.name
                let fileString = ""
                if (fileType === "ass" || fileType === "srt") {
                    fileString = await getVttFromFile(file)
                } else if (validFileTypes[0].indexOf("image") !== -1) {
                    fileString = await getTextFromFile(file)
                }
                if (fileString !== "") {
                    const hash = crypto.createHash('sha256');
                    hash.update(fileString);
                    if (fileType === "ass") {
                        fileType = "vtt"
                    }
                    fileName = `${hash.digest('hex')}.${fileType}`
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
        onSuccess: (response: any, file: UploadFile) => {
            message.success(`${file.name} 文件上传成功.`);
            const index = action.indexOf("?")
            if (fileLimit === 1) {
                setURL(action.substring(0, index))
                setFileList([file])
            } else {
                file.status = 'done';
                file.response = response;
                file.url = action.substring(0, index)
                setSucceedFile(file)
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
            tmpFileList.sort(sortFile)
            setFileList(tmpFileList)
            if (status === 'done') {
                message.success(`${info.file.name} 文件上传成功s.`);
            } else if (status === 'error') {
                message.error(`${info.file.name} 文件上传失败s.`);
            }
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