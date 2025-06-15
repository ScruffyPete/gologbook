import { redirect } from "react-router-dom"

export function isValidToken(token: string | null): boolean {
    if (!token) return false

    try {
        const payload = JSON.parse(atob(token.split('.')[1]))
        const now = Math.floor(Date.now() / 1000)
        return payload.exp && payload.exp > now
    } catch {
        return false
    }
}

export function rootRedirectLoader() {
    const token = localStorage.getItem("token")
    return isValidToken(token) ? redirect("/projects") : redirect("/login")
}

export function redirectIfAuthedLoader() {
    const token = localStorage.getItem("token")
    if (isValidToken(token)) {
        return redirect("/projects")
    }
    return null
}

export function authLoader() {
    const token = localStorage.getItem("token")
    if (!token || !isValidToken(token)) {
        return redirect("/login")
    }
    return null
}