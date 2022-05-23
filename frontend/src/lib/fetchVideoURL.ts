import axios from "axios"

const instance = axios.create({
    baseURL: 'http://localhost:8080/api/video/',
});

interface videoStats {
    url: string
}
export const fetchVideoURL = async (id: string) => {
    const res = await instance.get('', {
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