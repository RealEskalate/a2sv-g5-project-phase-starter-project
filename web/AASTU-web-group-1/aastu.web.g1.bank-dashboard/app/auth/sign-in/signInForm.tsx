"use client";
import React, { useState } from "react";
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
import { signInSchema } from "@/schema";
import { useRouter } from "next/navigation";

export const SignInForm = () => {
  const router = useRouter();
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState("");
  const form = useForm({
    resolver: zodResolver(signInSchema),
    defaultValues: {
      userName: "",
      password: "",
    },
  });

  const onSubmit = async (values: z.infer<typeof signInSchema>) => {
    setLoading(true);
    try {
      const result = await signIn("credentials", {
        userName: values.userName,
        password: values.password,
        redirect: false,
      });
      console.log(result);
      if (result?.error) {
        console.error("Sign-in error:", result.error);
        setError(result.error);
      } else {
        console.log("Sign-in successful");
        router.push("/dashboard");
      }
    } catch (error) {
      console.error("Unexpected error:", error);
    } finally {
      setLoading(false);
    }
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
            name="userName"
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
            disabled={loading}
            type="submit"
            className="bg-primaryBlue rounded-2xl mt-12 w-full"
          >
            {loading ? "Signing In..." : "Sign In"}
          </Button>
          <p className="text-red-500 text-sm text-center">
            {error && "Something went wrong. Please try again"}
          </p>
        </form>
      </Form>
      <Link href="/auth/sign-up" className="text-sm text-center text-blue-400">
        Don&apos;t have an account? Sign Up
      </Link>
    </div>
  );
};

export default SignInForm;
