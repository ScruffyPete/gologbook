import { Entry } from "@/types/Entry"
import { useEffect, useRef } from "react"
import { Card, CardContent, CardFooter } from "./ui/card"
import { Textarea } from "./ui/textarea"
import { Button } from "./ui/button"
import { ArrowUpCircle } from "lucide-react"

interface EntryFeedProps {
    entries: Entry[]
    newEntryBody: string
    setNewEntryBody: (value: string) => void
    onSubmit: () => void
}

export function EntryFeed({
    entries,
    newEntryBody,
    setNewEntryBody,
    onSubmit,
}: EntryFeedProps) {
    const scrollAnchorRef = useRef<HTMLDivElement | null>(null)
    const textareaRef = useRef<HTMLTextAreaElement>(null)

    useEffect(() => {
        scrollAnchorRef.current?.scrollIntoView({ behavior: "smooth" })
    }, [entries])

    return (
        <div className="flex flex-col flex-1 w-full h-full">
            <div className="flex-1 overflow-y-auto px-4 py-2 space-y-2">
                {entries.map(entry => (
                    <div key={entry.id} className="flex gap-4">
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
            </div>

            <div className="flex gap-2 border-t p-4">
                <div className="relative w-full">
                    <Textarea
                        ref={textareaRef}
                        className="w-full pr-12 text-sm bg-background border rounded-md p-3 leading-5 max-h-[33vh] overflow-y-auto resize-none"
                        value={newEntryBody}
                        onChange={e => setNewEntryBody(e.target.value)}
                        onInput={(e) => {
                            const target = e.currentTarget
                            target.style.height = "auto"
                            target.style.height = `${Math.min(target.scrollHeight, window.innerHeight / 3)}px`
                        }}
                        onKeyDown={(e) => {
                            if (e.key === "Enter" && !e.shiftKey) {
                                e.preventDefault()
                                onSubmit()
                            }
                        }}
                        placeholder="Type here..."
                        rows={1}
                    />
                    <Button
                        onClick={onSubmit}
                        className="absolute bottom-2 right-2 h-8 w-8 p-0 rounded-full"
                        variant="ghost"
                    >
                        <ArrowUpCircle className="h-6 w-6 text-muted-foreground" />
                    </Button>
                </div>
            </div>
        </div>
    )
}