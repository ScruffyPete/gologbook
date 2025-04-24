import { Entry } from "@/types/Entry";
import { Project } from "@/types/Project";
import { LoaderFunctionArgs, useLoaderData } from "react-router-dom";
import { ScrollArea } from "./ui/scroll-area";
import { Textarea } from "./ui/textarea";
import { useEffect, useRef, useState } from "react";
import { Button } from "./ui/button";
import { Card, CardContent, CardFooter } from "./ui/card";

export async function loader({ params }: LoaderFunctionArgs) {
    try {
        const [projectRes, entriesRes] = await Promise.all([
            fetch(`/api/projects/${params.projectId}`),
            fetch(`/api/projects/${params.projectId}/entries`)
        ]);

        if (!projectRes.ok || !entriesRes.ok) {
            throw new Error('Failed to load data');
        }

        const [project, entries] = await Promise.all([
            projectRes.json(),
            entriesRes.json()
        ]);

        return { project, entries };
    } catch (error) {
        console.error(error);
        throw new Response('Failed to load project', { status: 500 });
    }
}

export function ProjectPage() {
    const { project, entries } = useLoaderData<{ project: Project, entries: Entry[] }>()

    const [entriesState, setEntriesState] = useState<Entry[]>(entries)
    const [newEntryBody, setNewEntryBody] = useState("")

    const scrollAnchorRef = useRef<HTMLDivElement | null>(null)

    useEffect(() => {
        setEntriesState(entries)
    }, [project.id, entries])

    useEffect(() => {
        scrollAnchorRef.current?.scrollIntoView({ behavior: "smooth" })
    }, [entriesState])

    const handleSubmit = async () => {
        const entry = newEntryBody.trim()
        try {
            const res = await fetch(`/api/projects/${project.id}/entries`, {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({ body: entry })
            })
            const newEntry = await res.json()
            setEntriesState((prev) => [...prev, newEntry])
            setNewEntryBody("")
        } catch (err) {
            console.error(err)
            alert('Error submitting an entry')
        }
    }


    return (
        <div className="flex flex-col h-screen">
            <ScrollArea className="flex-1 overflow-y-auto px-4 py-2 space-y-2">
                {entriesState.map(entry => (
                    <Card className="bg-muted shadow-sm rounded-lg px-4 py-1 my-2">
                        <CardContent className="px-4 pt-4 pb-2 text-sm whitespace-pre-line">
                            {entry.body}
                        </CardContent>
                        <CardFooter className="px-4 pb-4 text-xs text-muted-foreground justify-end">
                            {new Date(entry.createdAt).toLocaleString()}
                        </CardFooter>
                    </Card>
                ))}
                <div ref={scrollAnchorRef} />
            </ScrollArea>
            <div className="flex gap-2 border-t p-4">
                <Textarea
                    className="flex-1 resize-none"
                    value={newEntryBody}
                    onChange={e => setNewEntryBody(e.target.value)}
                    placeholder="Type a message..."
                    rows={1}
                />
                <Button onClick={handleSubmit}>Send</Button>
            </div>
        </div >
    )
}