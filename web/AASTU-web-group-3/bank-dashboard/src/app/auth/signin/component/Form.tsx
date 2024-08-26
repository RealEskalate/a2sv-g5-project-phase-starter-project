"use client";
import { useForm } from "react-hook-form";
import { signIn } from "next-auth/react";
import { useSearchParams } from "next/navigation";
import Link from 'next/link'
// import {transferingPeoples } from '@/../../public/Icons'
import transferImage from '@/../../public/Images/transferr.png'
import Image from "next/image";

export default function Form() {
    const { register, handleSubmit, formState, reset } = useForm();
    const { errors } = formState;
    const searchParams = useSearchParams();
    const callbackUrl = searchParams.get('callbackUrl') || "/";

    const onSubmit = async (data: any) => {
        try {
            await signIn("credentials", {
                redirect: true,
                callbackUrl:'/dashboard',
                userName: data.userName,
                password: data.password,
            });
            reset();
        } catch (error) {
            console.error("Network response was not ok");
        }
    };



    return (
        <div className="container flex justify-center w-11/12 lg:grid grid-cols-2">
            <div className="hidden lg:block">
                <Image src={transferImage} alt="login image" width={500} height={600} />
            </div>
            <div className="md:w-[60%] bg-gray-100 dark:bg-darkComponent shadow-lg rounded-lg py-16 px-6">
                <h1 className="text-3xl font-bold text-center mb-4 text-[#25324B] dark:text-darkText">Welcome Back</h1>
                <form onSubmit={handleSubmit(onSubmit)} noValidate>
                    <div className="flex flex-col gap-2">
                        <label htmlFor="userName" className="dark:text-darkText">UserName</label>
                        <input
                            type="text"
                            id="userName"
                            className="border rounded-md w-full h-12 border-[#CCCCF5] dark:border-darkComponent p-2 bg-white dark:bg-darkPage dark:text-darkText"
                            {...register("userName", {
                                required: {
                                    value: true,
                                    message: "Username is required",
                                },
                            })}
                        />
                        <p className="text-red-500">{errors.userName?.message?.toString()}</p>
                    </div>
                    <div className="flex flex-col gap-2">
                        <label htmlFor="password" className="dark:text-darkText">Password</label>
                        <input
                            type="password"
                            id="password"
                            className="border rounded-md w-full h-12 border-[#CCCCF5] dark:border-darkComponent p-2 bg-white dark:bg-darkPage dark:text-darkText"
                            {...register("password", {
                                required: {
                                    value: true,
                                    message: "Password is required",
                                },
                            })}
                        />
                        <p className="text-red-500">{errors.password?.message?.toString()}</p>
                    </div>
                    <button
                        type="submit"
                        className="border w-full rounded-xl h-12 bg-[#4640DE] dark:bg-darkAccent text-white dark:text-darkText font-semibold my-4"
                    >
                        Login 
                    </button>
                    <div className="dark:text-darkText">
                        Don’t have an account? <Link href="/auth/signup"><span className="text-[#4640DE] dark:text-darkAccent font-semibold ml-4 hover:underline">Sign Up</span></Link>
                    </div>
                </form>
            </div>
        </div>
    );
    
}
