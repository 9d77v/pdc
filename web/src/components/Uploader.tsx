import { message, Upload } from 'antd';
import React, { useState, useEffect } from 'react'
import { getUploadURL } from "src/consts/http";
import { UploadFile } from "antd/lib/upload/interface";
import axios from 'axios'
import crypto from 'crypto'
import { getVttFromFile, getType } from 'src/utils/subtitle';
import { getTextFromFile, replaceURL } from 'src/utils/file';
import { supportedSubtitleSuffix } from 'src/consts/consts';

interface UploaderProps {
    fileLimit: number
    bucketName: string
    filePathPrefix?: string
    validFileTypes: string[]
    setURL: (e: any) => any
}


export const Uploader: React.FC<UploaderProps> = ({ fileLimit, bucketName, filePathPrefix, validFileTypes, setURL }) => {
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

    const isSubtitleType = (fileType: string): Boolean => {
        if (fileType === "vtt") {
            return false
        }
        for (let t of supportedSubtitleSuffix) {
            if (fileType === t) {
                return true
            }
        }
        return false
    }
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
            return new Promise<File>(async (resolve) => {
                let fileType = getType(file)
                if (isSubtitleType(fileType)) {
                    const vttText = await getVttFromFile(file);
                    const blob = new Blob([vttText], {
                        type: 'text/vtt',
                    })
                    file = new File([blob], file.name, { type: 'text/vtt', lastModified: Date.now() });
                }
                let fileName = file.name
                let fileString = ""
                if (isSubtitleType(fileType)) {
                    fileString = await getVttFromFile(file)
                } else if (validFileTypes[0].indexOf("image") !== -1) {
                    fileString = await getTextFromFile(file)
                }
                if (fileString !== "") {
                    const hash = crypto.createHash('sha256');
                    hash.update(fileString);
                    if (isSubtitleType(fileType)) {
                        fileType = "vtt"
                    }
                    fileName = `${hash.digest('hex')}.${fileType}`
                }
                fileName = filePathPrefix === undefined ? fileName : filePathPrefix + fileName
                const data = await getUploadURL(bucketName, fileName);
                setAction(data.data.presignedUrl)
                resolve(file);
            });
        },
        onSuccess: (response: any, file: UploadFile) => {
            message.success(`${file.name} 文件上传成功.`);
            const index = action.indexOf("?")
            const url = replaceURL(action.substring(0, index))
            if (fileLimit === 1) {
                setURL(url)
                setFileList([file])
            } else {
                file.status = 'done';
                file.response = response;
                file.url = url
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
            <p className="ant-upload-hint">点击或拖拽上传文件</p>
        </Upload.Dragger>
    )
}
