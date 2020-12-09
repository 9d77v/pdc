
export function checkIsFile(source: any) {
    return source instanceof File
}

export function getExt(url: string): string {
    if (url.includes('?')) {
        return getExt(url.split('?')[0])
    }

    if (url.includes('#')) {
        return getExt(url.split('#')[0])
    }

    return url
        .trim()
        .toLowerCase()
        .split('.')
        .pop() || ''
}

export function getType(source: any) {
    return checkIsFile(source) ? getExt(source.name) : getExt(source)
}

export async function getVttFromFile(file: File) {
    return new Promise<string>(resolve => {
        const reader = new FileReader()
        reader.onload = () => {
            var subsrt = require('subsrt')
            var vtt = subsrt.convert(reader.result?.toString(), { format: "vtt" })
            resolve(vtt)
        }
        reader.readAsText(file)
    })
}
