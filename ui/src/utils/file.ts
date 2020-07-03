export async function getTextFromFile(file: File) {
    return new Promise<string>(resolve => {
        const reader = new FileReader();
        reader.onload = () => {
            resolve(reader.result?.toString())
        };
        reader.readAsText(file);
    });
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