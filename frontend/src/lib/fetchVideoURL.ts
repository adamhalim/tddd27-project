import axios from "axios"

const instance = axios.create({
    baseURL: 'http://localhost:8080/api/',
});

interface fetchVideoData {
    url: string,
    viewcount: number,
    videotitle: string,

}

interface fetchPreviewData {
    url: string
}
export const fetchVideoURL = async (id: string): Promise<fetchVideoData | undefined> => {
    const res = await instance.get('video/', {
        params: {
            chunkName: id,
        }
    })
    if (res.status === 200) {
        const data: fetchVideoData = res.data
        return data
    } else {
        // TODO: Error handling here
        return undefined
    }
}

export const fetchPreviewUrl = async (chunkName: string) => {
    const res = await instance.get('preview/', {
        params: {
            chunkName: chunkName,
        }
    })
    if (res.status === 200) {
        const data: fetchPreviewData = res.data
        return data.url
    } else {
        // TODO: Error handling here
    }
}