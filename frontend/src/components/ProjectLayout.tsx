import { ScrollArea } from "@radix-ui/react-scroll-area";
import { ReactNode } from "react";

interface ProjectLayoutProps {
    children: ReactNode
    outputStream: ReactNode
}

export function ProjectLayout({ children, outputStream }: ProjectLayoutProps) {
    return (
        <div className="flex w-full h-full overflow-hidden">
            <ScrollArea className="w-1/2 h-full border-r overflow-auto">
                {children}
            </ScrollArea>

            <ScrollArea className="w-1/2 h-full overflow-auto">
                {outputStream}
            </ScrollArea>
        </div>
    )
}