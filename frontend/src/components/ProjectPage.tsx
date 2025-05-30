import { Entry } from "@/types/Entry";
import { Project } from "@/types/Project";
import { LoaderFunctionArgs, useLoaderData } from "react-router-dom";
import { ScrollArea } from "./ui/scroll-area";
import { Textarea } from "./ui/textarea";
import { useEffect, useRef, useState } from "react";
import { Button } from "./ui/button";
import { Card, CardContent, CardFooter } from "./ui/card";
import { ArrowUpCircle } from "lucide-react";
import { Insight } from "@/types/Insight";

export async function loader({ params }: LoaderFunctionArgs) {
    try {
        const [projectRes, entriesRes, insightsRes] = await Promise.all([
            fetch(`/api/projects/${params.projectId}`, {
                headers: {
                    'Authorization': `Bearer ${localStorage.getItem('token')}`
                }
            }),
            fetch(`/api/entries/?project_id=${params.projectId}`, {
                headers: {
                    'Authorization': `Bearer ${localStorage.getItem('token')}`
                }
            }),
            fetch(`/api/insights/?project_id=${params.projectId}`, {
                headers: {
                    'Authorization': `Bearer ${localStorage.getItem('token')}`
                }
            })
        ]);

        if (!projectRes.ok || !entriesRes.ok || !insightsRes.ok) {
            throw new Error('Failed to load data');
        }

        const [project, entries, insights] = await Promise.all([
            projectRes.json(),
            entriesRes.json(),
            insightsRes.json()
        ]);

        return { project, entries, insights };
    } catch (error) {
        console.error(error);
        throw new Response('Failed to load project', { status: 500 });
    }
}

export function ProjectPage() {
    const { project, entries, insights } = useLoaderData<{ project: Project, entries: Entry[], insights: Insight[] }>()

    const [entriesState, setEntriesState] = useState<Entry[]>(entries)
    const [insightsByEntryId, setInsightsByEntryId] = useState<Record<string, Insight>>({})
    const [newEntryBody, setNewEntryBody] = useState("")

    const scrollAnchorRef = useRef<HTMLDivElement | null>(null)
    const textareaRef = useRef<HTMLTextAreaElement>(null)

    useEffect(() => {
        setEntriesState(entries)
    }, [project.id, entries])

    useEffect(() => {
        setInsightsByEntryId(insights.reduce((acc, insight) => {
            insight.entry_ids.forEach(entryID => {
                acc[entryID] = insight
            })
            return acc
        }, {} as Record<string, Insight>))
    }, [insights])
    
    useEffect(() => {
        scrollAnchorRef.current?.scrollIntoView({ behavior: "smooth" })
    }, [entriesState])

    const handleSubmit = async () => {
        const entry = newEntryBody.trim()
        try {
            const res = await fetch(`/api/entries/`, {
                method: 'POST',
                headers: {
                    'Authorization': `Bearer ${localStorage.getItem('token')}`,
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({ body: entry, project_id: project.id })
            })
            const newEntry = await res.json()
            setEntriesState((prev) => [...prev, newEntry])
            setNewEntryBody("")
            if (textareaRef.current) {
                textareaRef.current.style.height = "auto"
            }
        } catch (err) {
            console.error(err)
            alert('Error submitting an entry')
        }
    }

    return (
        // <div className="flex flex-col h-screen w-full">
        <div className="flex flex-col flex-1 w-full">
            <ScrollArea className="flex-1 overflow-y-auto px-4 py-2 space-y-2">
                {entriesState.map(entry => (
                    <div key={entry.id} className="flex gap-4">
                        <div className="w-2/5">
                            {insightsByEntryId[entry.id] ? (
                                <Card className="bg-muted-foreground/10 p-2 text-sm">
                                    <CardContent>{insightsByEntryId[entry.id].body}</CardContent>
                                    <CardFooter className="px-4 pb-4 text-xs text-muted-foreground justify-end">
                                        {new Date(insightsByEntryId[entry.id].created_at).toLocaleString()}
                                    </CardFooter>
                                </Card>
                                
                            ) : (
                                <div className="text-xs text-muted-foreground italic">No insight</div>
                            )}
                        </div>
                        <div className="w-3/5">
                            <Card className="bg-muted shadow-sm rounded-lg px-4 py-1 my-2">
                                <CardContent className="px-4 pt-4 pb-2 text-sm whitespace-pre-line">
                                    {entry.body}
                                </CardContent>
                                <CardFooter className="px-4 pb-4 text-xs text-muted-foreground justify-end">
                                    {new Date(entry.created_at).toLocaleString()}
                                </CardFooter>
                            </Card>
                        </div>
                    </div>
                ))}
                <div ref={scrollAnchorRef} />
            </ScrollArea>
            <div className="flex gap-2 border-t p-4">
                <div className="relative w-full">
                    <Textarea
                        ref={textareaRef}
                        className="w-full pr-12 text-sm bg-background border rounded-md p-3 leading-5 max-h-[33vh] overflow-y-auto resize-none"
                        value={newEntryBody}
                        onChange={e => setNewEntryBody(e.target.value)}
                        onInput={(e) => {
                            const target = e.currentTarget
                            target.style.height = "auto" // Reset height
                            target.style.height = `${Math.min(target.scrollHeight, window.innerHeight / 3)}px` // Limit to ⅓ screen
                        }}
                        onKeyDown={(e) => {
                            if (e.key === "Enter" && !e.shiftKey) {
                                e.preventDefault()
                                handleSubmit()
                            }
                        }}
                        placeholder="Type here..."
                        rows={1}
                    />
                    <Button
                        onClick={handleSubmit}
                        className="absolute bottom-2 right-2 h-8 w-8 p-0 rounded-full"
                        variant="ghost"
                    >
                        <ArrowUpCircle className="h-6 w-6 text-muted-foreground" />
                    </Button>
                </div>
            </div>
        </div >
    )
}