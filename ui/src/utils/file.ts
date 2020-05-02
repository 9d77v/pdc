export async function getTextFromFile(file: File) {
    return new Promise<string>(resolve => {
        const reader = new FileReader();
        reader.onload = () => {
            resolve(reader.result?.toString())
        };
        reader.readAsText(file);
    });
}