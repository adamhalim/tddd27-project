import axios from "axios"

const instance = axios.create({
    baseURL: 'http://localhost:8080/api/',
});

interface videoStats {
    url: string
}
export const fetchVideoURL = async (id: string) => {
    const res = await instance.get('video/', {
        params: {
            chunkName: id,
        }
    })
    if (res.status === 200) {
        const data: videoStats = res.data
        return data.url
    } else {
        // TODO: Error handling here
    }
}

export const fetchPreviewUrl = async (chunkName: string) => {
    const res = await instance.get('preview/', {
        params: {
            chunkName: chunkName,
        }
    })
    if (res.status === 200) {
        const data: videoStats = res.data
        return data.url
    } else {
        // TODO: Error handling here
    }
}