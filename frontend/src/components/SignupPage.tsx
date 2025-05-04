'use client'

import { useForm } from 'react-hook-form'

import {
    Form,
    FormField,
    FormItem,
    FormLabel,
    FormControl,
} from './ui/form'
import { Input } from './ui/input'
import { Button } from './ui/button'
import { z } from 'zod'
import { zodResolver } from '@hookform/resolvers/zod'
import { Link, useNavigate } from 'react-router-dom'


const signupSchema = z.object({
    email: z.string().email(),
    password: z.string().min(1),
    confirmPassword: z.string().min(1),
})

export function SignupPage() {
    const navigate = useNavigate()
    const form = useForm<z.infer<typeof signupSchema>>({
        resolver: zodResolver(signupSchema),
        defaultValues: {
            email: '',
            password: '',
            confirmPassword: '',
        },
    })

    const onSubmit = async (values: z.infer<typeof signupSchema>) => {
        try {
            const res = await fetch('/api/signup', {
                method: 'POST',
                body: JSON.stringify(values),
                headers: {
                    'Content-Type': 'application/json'
                }
            })
            if (!res.ok) {
                alert('Error signing up')
            }
            navigate('/login')
        } catch (err) {
            console.error(err)
            alert('Error signing up')
        }
    }

    return (
        <div className="flex flex-col items-center justify-center h-screen">
            <h1 className="text-2xl font-bold">Signup</h1>
            <Form {...form}>
                <form
                    onSubmit={form.handleSubmit(onSubmit)}
                    className="space-y-4 w-80"
                >
                    <FormField
                        control={form.control}
                        name="email"
                        render={({ field }) => (
                            <FormItem>
                                <FormLabel>Email</FormLabel>
                                <FormControl>
                                    <Input
                                        id="email"
                                        type="email"
                                        placeholder="email@example.com"
                                        required
                                        {...field}
                                    />
                                </FormControl>
                            </FormItem>
                        )}
                    />
                    <FormField
                        control={form.control}
                        name="password"
                        render={({ field }) => (
                            <FormItem>
                                <FormLabel>Password</FormLabel>
                                <FormControl>
                                    <Input
                                        id="password"
                                        type="password"
                                        placeholder="Password"
                                        required
                                        {...field}
                                    />
                                </FormControl>
                            </FormItem>
                        )}
                    />
                    <FormField
                        control={form.control}
                        name="confirmPassword"
                        render={({ field }) => (
                            <FormItem>
                                <FormLabel>Confirm Password</FormLabel>
                                <FormControl>
                                    <Input
                                        id="confirmPassword"
                                        type="password"
                                        placeholder="Confirm Password"
                                        required
                                        {...field}
                                    />
                                </FormControl>
                            </FormItem>
                        )}
                    />
                    <Button type="submit" className="w-full">Signup</Button>
                    <div className="mt-4 text-center text-sm">
                        Already have an account?{" "}
                        <Link to="/login" className="underline underline-offset-4">
                            Login
                        </Link>
                    </div>
                </form>
            </Form>
        </div>
    )
}