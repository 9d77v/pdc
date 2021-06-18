import SparkMD5 from "spark-md5"

export async function getTextFromFile(file: File) {
    return new Promise<string>(resolve => {
        const reader = new FileReader()
        reader.onload = () => {
            resolve(reader.result?.toString() || "")
        }
        reader.readAsText(file)
    })
}

export async function getFileMD5(file: File) {
    return new Promise<string>(resolve => {
        const reader = new FileReader()
        var spark = new SparkMD5(); //创建md5对象（基于SparkMD5）
        if (file.size > 1024 * 1024 * 10) {
            var data1 = file.slice(0, 1024 * 1024 * 10); //将文件进行分块 file.slice(start,length)
            reader.readAsBinaryString(data1); //将文件读取为二进制码
        } else {
            reader.readAsBinaryString(file);
        }

        reader.onload = (e: any) => {
            spark.appendBinary(e.target.result);
            var md5 = spark.end()
            resolve(md5)
        }
    })
}

export function replaceURL(url: string): string {
    if (url !== "") {
        let newURL = ""
        const arr = url.split("/")
        for (let i = 3; i < arr.length; i++) {
            newURL += "/" + arr[i]
        }
        return newURL
    }
    return url
}

export async function blobToArrayBuffer(blob: Blob) {
    return new Promise(function (resolve): any {
        var reader = new FileReader()

        reader.onloadend = function () {
            resolve(reader.result)
        }
        reader.readAsArrayBuffer(blob)
    })
}
