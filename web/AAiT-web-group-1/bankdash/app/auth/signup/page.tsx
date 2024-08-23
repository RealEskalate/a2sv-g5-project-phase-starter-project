"use client"
import Link from 'next/link'
import React, { useEffect } from 'react'
import { useForm } from 'react-hook-form'
// import { useSignUpMutation } from '../service/getApi'
import { useRouter } from 'next/navigation'
import google from '../../Images/IcongoogleIcon.svg'
import Image from 'next/image'

interface formtype {
    name: string,
    email: string,
    password: string, 
    confirmPassword: string,
    role: string
}

const Page = () => {
    const router = useRouter()
    const form = useForm<formtype>()
    const {control, register, formState, handleSubmit} = form
    const { errors } = formState

    // rtk query hook
    // const [signUp, { data, isLoading, isError}] = useSignUpMutation()

    // if (isError){
    //     alert('Invalid information')
    //     router.push(`/signup`)
    // }
    // if (isLoading){
    //       return  <h1 className='text-center text-lg mt-72'>Loading ....</h1>
    // }



    const onSubmit = async (user: formtype) => {
        try{
            const res = await signUp({...user, "role": "user"})
            const { data } = res;

            if (data.message === "Successfully sent OTP"){
                router.push(`/verify-email?email=${user.email}`)
            }
            else{
                alert('Invalid information')
                router.push(`/signup`)
            }

        } catch (err) {
            alert('Email already exists')
            console.error(err)
        }}


    return (
    <div className='flex justify-center'>
        <div className='mt-10 w-5/12 bg-slate-40'>

            <div className='flex w-full flex-col space-y-6 items-center'>
                <h1 className='text-4xl text-[#25324B] font-poppins font-black'>Sign up today!</h1>

                <Link className='w-full flex justify-center' href="/api/auth/signin">
                <div className='flex w-9/12 justify-center space-x-3 border border-[#CCCCF5] py-2'>
                    <Image src={google} alt='google'></Image>
                    <h3 className=' text-[#4640DE] epi font-bold text-lg'>sign up with google</h3>
                </div>
                </Link>

                <div className='flex'>
                    <h3 className='font-light epi text-[#202430]'>Or Sign Up with Email</h3>
                </div>
            </div>

            <div className='w-full flex justify-center'>

                        {/* using react-hook-form to handle signup form */}

                <form onSubmit={(e) => e.preventDefault()} className='mb-3 mt-5 w-3/4 flex flex-col space-y-5' action="">
                    
                    <div className='flex space-y-2 flex-col'>
                        <label className='font-semibold' htmlFor="name">Full Name</label>
                        <input className='bg-white border p-2' type="text" id='name' {...register("name", 
                            {required: "Name is required"}
                        )} placeholder='Full name' />
                    </div>
                    { errors.name ? <p className='text-red-500 text-sm'>{errors.name.message}</p> : null}

                    <div className='flex space-y-2 flex-col'>
                        <label className='font-semibold' htmlFor="email">Email Address</label>
                        <input className='bg-white border p-2' type="text" placeholder='email' {...register("email", {required: "Email is required"})} />
                    </div>
                    { errors.email ? <p className='text-red-500 text-sm'>{errors.email.message}</p> : null}

                    <div className='flex space-y-2 flex-col'>
                        <label className='font-semibold' htmlFor="password">Password</label>
                        <input className='bg-white border p-2' type="password" placeholder='password' {...register("password", {required: "Password is required"})} />
                    </div>
                    { errors.password ? <p className='text-red-500 text-sm'>{errors.password.message}</p> : null}

                    <div className='flex space-y-2 flex-col'>
                        <label className='font-semibold' htmlFor="confirm-password">Confirm Password</label>
                        <input className='bg-white border p-2' type="password" placeholder='password' {...register("confirmPassword", {required: "You have to confirm your password"})}  />
                    </div>
                    { errors.confirmPassword ? <p className='text-red-500 text-sm'>{errors.confirmPassword.message}</p> : null}

                    <div>
                        <button onClick={handleSubmit(onSubmit)} className='w-full py-4 rounded-full bg-indigo-900 text-lg text-white epi'>continue</button>
                    </div>

                    <div className='space-y-6 mt-10'>

                        <h2 className='font-poppins  text-[#7C8493]'>Already have an account?<Link href={`/login`}> <span className='font-bold text-indigo-800'>Login</span></Link></h2>  

                        <p className='text-[#7C8493] epi'>
                            By clicking Continue, you acknowledge that you have read and accepted our <span className='text-indigo-800'>Terms of Service</span > and <span className='text-indigo-800'>Privacy Policy.</span>
                        </p>
                    </div>
                </form>
            </div>
        </div>

    </div>
  )
}

export default Page;