"use client";
import React from "react";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { z } from "zod";
import {
  Form,
  FormField,
  FormItem,
  FormLabel,
  FormControl,
  FormMessage,
} from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import Image from "next/image";
import Link from "next/link";
import { signIn } from "next-auth/react";

const formSchema = z.object({
  username: z.string().min(2).max(50),
  password: z.string().min(8),
});

export const SignInForm = () => {
  const form = useForm({
    resolver: zodResolver(formSchema),
    defaultValues: {
      username: "",
      password: "",
    },
  });

  const onSubmit = (values: z.infer<typeof formSchema>) => {
    signIn("credentials", {
      username: values.username,
      password: values.password,
      redirect: false,
    });
  };

  return (
    <div className="flex min-h-screen w-full max-w-[420px] flex-col justify-center gap-5 py-10 md:gap-8">
      <div className="flex items-center gap-2 ">
        <Image src="/icons/logo.png" width={25} height={25} alt="logo" />
        <h1 className="text-primaryBlack font-[900] text-[1.5rem]">
          BankDash.
        </h1>
      </div>
      <h1 className="gap-2 text-2xl text-primaryBlack font-bold">
        Welcome Back
      </h1>
      <Form {...form}>
        <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-3">
          {/* Username */}
          <FormField
            control={form.control}
            name="username"
            render={({ field }: { field: object }) => (
              <FormItem className="mb-1">
                <FormLabel>Username</FormLabel>
                <FormControl>
                  <Input
                    className="rounded-2xl md:w-[400px]"
                    placeholder="john_doe"
                    {...field}
                  />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
          {/* Password */}
          <FormField
            control={form.control}
            name="password"
            render={({ field }: { field: object }) => (
              <FormItem className="mb-1">
                <FormLabel>Password</FormLabel>
                <FormControl>
                  <Input
                    className="rounded-2xl md:w-[400px]"
                    placeholder="********"
                    type="password"
                    {...field}
                  />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />

          {/* Submit Button */}
          <Button
            type="submit"
            className="bg-primaryBlue rounded-2xl mt-12 w-full"
          >
            Sign In
          </Button>
        </form>
      </Form>
      <Link href="/auth/sign-up" className="text-sm text-center text-blue-400">
        Don&apos;t have an account? Sign in
      </Link>
    </div>
  );
};

export default SignInForm;
