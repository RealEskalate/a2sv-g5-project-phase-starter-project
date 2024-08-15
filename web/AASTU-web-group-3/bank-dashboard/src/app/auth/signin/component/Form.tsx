"use client";
import { useForm } from "react-hook-form";
import { signIn } from "next-auth/react";
import { useSearchParams } from "next/navigation";
import Link from 'next/link'

export default function Form() {
    const { register, handleSubmit, formState, reset } = useForm();
    const { errors } = formState;
    const searchParams = useSearchParams();
    const callbackUrl = searchParams.get('callbackUrl') || "/";

    const onSubmit = async (data: any) => {
        try {
            await signIn("credentials", {
                redirect: true,
                callbackUrl,
                userName: data.userName,
                password: data.password,
            });
            reset();
        } catch (error) {
            console.error("Network response was not ok");
        }
    };



    return (
        <div className="flex justify-center mx-6 my-12 mt-36 lg:grid grid-cols-2">
            <div className="hidden lg:block ">
                <h1></h1>
            </div>
            <div className="container  md:w-[70%]">
                <h1 className="text-3xl font-bold text-center mb-4 text-[#25324B]">Welcome Back</h1>
                <form onSubmit={handleSubmit(onSubmit)} noValidate>
                    <div className="flex flex-col gap-2">
                        <label htmlFor="userName">UserName</label>
                        <input
                            type="text"
                            id="userName"
                            className="border rounded-md w-full h-12 border-[#CCCCF5] p-2"
                            {...register("userName", {
                
                                required: {
                                    value: true,
                                    message: "username is required",
                                },
                            })}
                        />
                        <p style={{ color: "red" }}>{errors.email?.message?.toString()}</p>
                    </div>
                    <div className="flex flex-col gap-2">
                        <label htmlFor="password">Password</label>
                        <input
                            type="password"
                            id="password"
                            className="border rounded-md w-full h-12 border-[#CCCCF5] p-2"
                            {...register("password", {
                                required: {
                                    value: true,
                                    message: "Password is required",
                                },
                            })}
                        />
                        <p style={{ color: "red" }}>
                            {errors.password?.message?.toString()}
                        </p>
                    </div>
                    <button
                        type="submit"
                        className="border w-full rounded-xl h-12 bg-[#4640DE] text-white font-semibold my-4"
                    >
                        Continue
                    </button>
                    <div>
                        Don’t have an account? <Link href="/auth/signup"><span className="text-[#4640DE] font-semibold">Sign Up</span></Link>
                    </div>
                </form>
            </div>
        </div>
    );
}
