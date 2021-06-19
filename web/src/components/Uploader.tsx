import { message, Upload } from 'antd';
import React, { useState } from 'react'
import { getUploadURL } from "src/consts/http";
import { UploadFile } from "antd/lib/upload/interface";
import axios from 'axios'
import SparkMD5 from 'spark-md5'
import { getVttFromFile, getType } from 'src/utils/subtitle';
import { getFileMD5, getTextFromFile, replaceURL } from 'src/utils/file';
import { supportedSubtitleSuffix } from 'src/consts/consts';
import { uuid } from '@nano-sql/core/lib/utilities';

interface UploaderProps {
    fileLimit: number
    bucketName: string
    validFileTypes: string[]
    setURL: (url: string[]) => void
}

const Uploader: React.FC<UploaderProps> = ({ fileLimit, bucketName, validFileTypes, setURL }) => {
    const [action, setAction] = useState('');
    const emptyFileList: UploadFile<any>[] = []
    const [fileList, setFileList] = useState(emptyFileList)
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
            return new Promise<File | boolean>(async (resolve) => {
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
                } else if (validFileTypes[0].indexOf("video") !== -1) {
                    fileName = `${await getFileMD5(file)}.mp4`
                }
                if (fileString !== "") {
                    if (isSubtitleType(fileType)) {
                        fileType = "vtt"
                    }
                    fileName = `${SparkMD5.hash(fileString)}.${fileType}`
                }
                let checkFileName = fileName
                if (bucketName === "image") {
                    checkFileName = checkFileName.split(".")[0] + ".webp"
                }
                const data = await getUploadURL(bucketName, checkFileName);
                if (data.data.presignedUrl.ok) {
                    const url = data.data.presignedUrl.url
                    const newFile: UploadFile = {
                        uid: uuid(),
                        size: file.size,
                        name: file.name,
                        type: file.type
                    }
                    newFile.status = 'done';
                    newFile.url = url
                    const newFileList = [...fileList, newFile]
                    setFileList(newFileList)
                    let fileURLs: string[] = []
                    for (let tmpFile of newFileList) {
                        fileURLs.push(tmpFile.url || "")
                    }
                    setURL(fileURLs)
                    resolve(false)
                } else {
                    if (fileName !== checkFileName) {
                        const data = await getUploadURL(bucketName, fileName);
                        setAction(data.data.presignedUrl.url)
                    } else {
                        setAction(data.data.presignedUrl.url)
                    }
                    resolve(file)
                }
            });
        },
        onSuccess: (response: any, file: UploadFile) => {
            message.success(`${file.name} 文件上传成功.`);
            const index = action.indexOf("?")
            const url = replaceURL(action.substring(0, index))
            file.url = url
            let all_done = true
            for (let tmpFile of fileList) {
                if (tmpFile.uid === file.uid && tmpFile.status === 'uploading') {
                    tmpFile.status = 'done'
                }
                if (tmpFile.status !== 'done') {
                    all_done = false
                }
            }
            if (all_done) {
                let fileURLs: string[] = []
                for (let tmpFile of fileList) {
                    let obj: any = tmpFile.originFileObj
                    if (obj.url.indexOf("/image/") >= 0) {
                        fileURLs.push(obj.url.split(".")[0] + ".webp")
                    } else {
                        fileURLs.push(obj.url)
                    }
                }
                setURL(fileURLs)
            }
        },
        onError: (error: any) => {
            message.error(`文件上传失败.`);
            setURL([])
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
                        'Content-Type': file.type
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

export default Uploader
