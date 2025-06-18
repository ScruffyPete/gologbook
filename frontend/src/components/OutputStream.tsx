import { useEffect, useState } from "react"
import { fetchEventSource } from '@microsoft/fetch-event-source'

interface OutputStreamProps {
    projectID: string
    streamKey: number
}

export function OutputStream({ projectID, streamKey }: OutputStreamProps) {
    const [content, setContent] = useState<string>("")
    const [loading, setLoading] = useState<boolean>(false)

    const getToken = () => localStorage.getItem("token")

    const fetchInitialOutput = async () => {
        setLoading(true)
        setContent("")

        try {
            const res = await fetch(`/api/documents/${projectID}/output/`, {
                headers: { Authorization: `Bearer ${getToken()}` },
            })
            const data = await res.json()
            setContent(data.body)
        } catch (err) {
            console.error("Initial fetch failed:", err)
        } finally {
            setLoading(false)
        }
    }

    const startStream = () => {
        const controller = new AbortController()

        setLoading(true)
        setContent("")

        fetchEventSource(`/api/documents/${projectID}/stream/`, {
            signal: controller.signal,
            headers: {
                Authorization: `Bearer ${getToken()}`,
            },
            async onopen(response) {
                if (!response.ok) {
                    throw new Error(`Stream error: ${response.statusText}`)
                }
            },
            onmessage(ev) {
                if (!ev.data || !ev.data.trim()) return
                setLoading(false)
                console.log("Incoming:", ev.data, "data")
                setContent((prev) => (prev ? prev + "\n\n" : "") + ev.data)
            },
            onerror(err) {
                console.error("Stream error:", err)
                controller.abort()
            },
        })
        return controller
    }

    useEffect(() => {
        fetchInitialOutput()
    }, [projectID])

    useEffect(() => {
        const controller = startStream()
        return () => controller?.abort()
    }, [streamKey, projectID])

    return (
        <div className="p-4 h-full text-sm whitespace-pre-wrap font-mono">
            {loading && (
                <div className="text-muted-foreground italic">Streaming output...</div>
            )}
            {content ? (
                <div>{content}</div>
            ) : !loading ? (
                <div className="text-muted-foreground">No output.</div>
            ) : null}
        </div>
    )
}
