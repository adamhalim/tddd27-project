import axios from 'axios';
import { v4 } from 'uuid';
import React, { useState } from 'react';
import './style.css'
import UploadProgress from '../UploadProgress';
import { useAuth0 } from '@auth0/auth0-react';


const ALLOWED_FILETYES = [
    "video/mp4",        // MPEG-4 .mp4
    "video/x-msvideo",  // A/V .avi
    "video/x-ms-wmv",   // Windows Media.wmv
    "video/quicktime",   // Quicktime .mov
    "video/avi",
    "video/x-flv"
]

const generatevideoURL = (chunkName: string): string => {
    return `/video/${chunkName}`
}


const UploadButton = () => {

    const { getAccessTokenSilently, getAccessTokenWithPopup, user } = useAuth0();

    const instance = axios.create({
        baseURL: 'http://localhost:8080/api/auth/',
    });
    interface ChunkConstants {
        maxFileSize: string,
        chunkSize: string
    }


    const [progress, setProgress] = useState(0);
    const [uploadInProgress, setUploadInProgress] = useState(false);
    const [errorOccured, setErrorOccured] = useState(false);
    const [fileName, setFileName] = useState("")
    const [statusText, setStatusText] = useState("");
    const [loading, setLoading] = useState(false);
    const [videoURL, setVideoURL] = useState("");

    const submit = async (e: React.ChangeEvent<HTMLInputElement>) => {
        setErrorOccured(false);
        setProgress(0);
        setStatusText("")
        setVideoURL("");
        const file = e.target.files?.item(0);

        if (file) {
            setUploadInProgress(true);
            setStatusText("Uploading file...")
            setLoading(true);
            setFileName(file.name)

            if (!ALLOWED_FILETYES.includes(file.type)) {
                setLoading(false)
                setErrorOccured(true)
                setStatusText(`filetype ${file.type} not allowed!`)
                return
            }

            const accessToken = await getAccessTokenSilently({ audience: 'http://localhost:3000/' })
                .then((res) => {
                    return res;
                }).catch((err) => {
                    console.log(err);
                    // getAccessTokenSilently() with audience won't work on localhost,
                    // but will work with a popup. Ghetto workaround, but it works for now..
                    return getAccessTokenWithPopup({ audience: 'http://localhost:3000/' })
                })


            const { maxFileSize, chunkSize } = await getChunkConstants(accessToken)
            const MAX_FILESIZE = parseInt(maxFileSize)
            const CHUNK_SIZE = parseInt(chunkSize)

            if (file.size > MAX_FILESIZE) {
                setErrorOccured(true)
                setLoading(false)
                setStatusText(`File size of ${(file.size / 1e6).toFixed(1)} MB greaer than maximum ${(MAX_FILESIZE / 1e6).toFixed(1)}MB`)
                return;
            }

            const chunkCount = getChunkCount(file, CHUNK_SIZE);
            const chunkName = v4()

            for (let chunk = 0; chunk < chunkCount; chunk++) {
                const blob = file.slice(chunk * CHUNK_SIZE, (chunk + 1) * CHUNK_SIZE);
                const success = await uploadChunk(blob, chunk, file.name, chunkName, accessToken);
                if (success) {
                    setProgress(((chunk + 1) / chunkCount) * 100)
                } else {
                    // TODO: set error response from uploadChuck
                    setErrorOccured(true)
                    setLoading(false)
                    setStatusText("error occured during upload");
                    return;
                }
            }
            setLoading(false)
            setStatusText("Upload complete!")
            setTimeout(() => {
                setStatusText("Transcoding file...")
                setLoading(true)
            },3000)

            const transcodeStatus = await allChunksUploaded(chunkName, accessToken)
            setLoading(false)
            if (transcodeStatus) {
                setStatusText("Transcoding complete!")
                setVideoURL(generatevideoURL(chunkName))
            } else {
                setStatusText("Transcoding failed...")
            }
        }
    }

    const uploadChunk = async (chunk: Blob, count: number, fileName: string, chunkName: string, accessToken: string): Promise<boolean> => {
        const res = await instance.post('videos/', chunk, {
            params: {
                id: count,
                fileName: fileName,
                chunkName: chunkName,
            },
            headers: {
                'Content-Type': 'application/json',
                Authorization: `Bearer ${accessToken}`
            },
            withCredentials: true,
        });
        if (res.status === 200) {
            return true;
        }
        return false;
    }

    const allChunksUploaded = async (chunkName: string, accessToken: string): Promise<boolean> => {
        const res = await instance.post('videos/combine/', {}, {
            params: {
                chunkName: chunkName,
            },
            headers: {
                'Content-Type': 'application/json',
                Authorization: `Bearer ${accessToken}`
            },
            withCredentials: true,
        });
        if (res.status === 200) {
            return true
        }
        return false
    }

    const getChunkConstants = async (accessToken: string): Promise<ChunkConstants> => {
        const res = await instance.get('videos/chunks/', {
            headers: {
                Authorization: `Bearer ${accessToken}`
            }
        })
        const data: ChunkConstants = res.data
        console.log(res)
        return data
    }


    // Very ugly code.
    // <input> is very difficult to style, so we just hide it
    // and place a button instead
    const handleButtonClick = () => {
        document.getElementById('upload-button')?.click()
    };

    const getChunkCount = (file: File, chunkSize: number): number => {
        const chunkCount = file.size % chunkSize === 0 ? 1 :
            Math.floor(file.size / chunkSize) + 1;
        return chunkCount;
    }


    return (
        <div className='upload-container'>
            <button onClick={handleButtonClick} >Upload file</button>
            <input
                type='file'
                onChange={submit}
                style={{ display: 'none' }}
                id='upload-button'
            />
            {
                uploadInProgress && <UploadProgress fileName={fileName} progress={progress} statusText={statusText} loading={loading} errorOccured={errorOccured} videoURL={videoURL} />
            }
        </div>
    )
}

export default UploadButton