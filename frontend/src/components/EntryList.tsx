import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import { Entry } from "../types/Entry";

export default function EntryListWrapper() {
    const [entries, setEntries] = useState<Entry[]>([])
    const [body, setBody] = useState('')
    const { projectId } = useParams()

    useEffect(() => {
        fetch(`/api/projects/${projectId}/entries`)
            .then(res => res.json())
            .then(setEntries)
            .catch(err => console.error('Failed to fetch entries:', err))
    }, [projectId])

    async function handleSubmit(e: React.FormEvent) {
        e.preventDefault()
        if (!body.trim()) return

        try {
            const res = await fetch(`/api/projects/${projectId}/entries`, {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({ body })
            })
            const newEntry = await res.json()
            setEntries(prev => [...prev, newEntry])
            setBody('')
        } catch (err) {
            console.error('Failed create entry:', err)
            alert('Could not create entry')
        }
    }

    return (
        <div style={{ display: 'flex', flexDirection: 'column', gap: '1rem' }}>
            <EntryList entries={entries} />

            <form onSubmit={handleSubmit} style={{ display: 'flex', gap: '0.5rem' }}>
                <input
                    type="text"
                    placeholder="Type a message..."
                    value={body}
                    onChange={(e) => setBody(e.target.value)}
                    style={{ flex: 1 }}
                />
                <button type="submit">Send</button>
            </form>
        </div>
    )
}

export function EntryList({ entries }: { entries: Entry[] }) {
    return (
        <ul style={{ flex: 1 }}>
            {entries.map(entry => (
                <li key={entry.id}>{entry.body}</li>
            ))}
        </ul>
    )
}
