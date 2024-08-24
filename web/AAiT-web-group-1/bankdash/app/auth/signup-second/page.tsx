"use client";
import Link from "next/link";
import React, { useEffect } from "react";
import { useForm } from "react-hook-form";
import { useRouter } from "next/navigation";
import { useSelector, useDispatch } from "react-redux";
import { formType } from "@/types/formType";
import { setform } from "@/lib/redux/slices/formSlice";
import Image from "next/image";
import logo from '../../../public/images/Logo.svg'



const Page = () => {
  const router = useRouter();
  const dispatch = useDispatch();
  const form = useForm<formType>();
  const { control, register, formState, handleSubmit } = form;
  const { errors } = formState;


  const onSubmit = (form: formType) => {
    dispatch(setform(form));
    router.push(`/auth/signup-third`);
  };

  return (
    <div className="flex justify-center">
      <div className="mt-6 w-5/12 bg-slate-40">
        <div className="flex mx-2 w-full flex-col space-y-6 items-center">
          <div className=" md:block hidden">
            <Image src={logo} className="ml-1" alt="LOGO" />
          </div>
        </div>

        <div className="flex mt-4 justify-center">
          <ul className="flex items-center">
            <li className=" border py-2 px-4 bg-blue-600  rounded-full">1</li>
            <div className="w-16 mx-2 h-1 bg-blue-700"></div>
            <li className=" border py-2 px-4 bg-blue-600  rounded-full">2</li>
            <div className="w-16 mx-2 h-1 border"></div>
            <li className=" border py-2 px-4 rounded-full">3</li>
          </ul>
        </div>

        <div className="w-full flex justify-center">
          {/* using react-hook-form to handle signup form */}

          <form
            onSubmit={(e) => e.preventDefault()}
            className="mb-3 mt-5 w-3/4 flex flex-col space-y-3"
            action=""
          >
            <div className="flex space-y-2 flex-col">
              <label className="font-semibold" htmlFor="name">
                Username
              </label>
              <input
                className="bg-white border p-2"
                type="text"
                id="username"
                {...register("username", { required: "username is required" })}
                placeholder="Username"
              />
            </div>
            {errors.username ? (
              <p className="text-red-500 text-sm">{errors.username.message}</p>
            ) : null}

            <div className="flex space-y-2 flex-col">
              <label className="font-semibold" htmlFor="email">
                Password
              </label>
              <input
                className="bg-white border p-2"
                type="text"
                placeholder="password"
                {...register("password", { required: "password is required" })}
              />
            </div>
            {errors.password ? (
              <p className="text-red-500 text-sm">{errors.password.message}</p>
            ) : null}

            <div className="flex space-y-2 flex-col">
              <label className="font-semibold" htmlFor="password">
                Present Address
              </label>
              <input
                className="bg-white border p-2"
                type="text"
                placeholder="Present Address"
                {...register("presentAddress", { required: "Present Address is required" })}
              />
            </div>
            {errors.presentAddress ? (
              <p className="text-red-500 text-sm">{errors.presentAddress.message}</p>
            ) : null}


            <div className="flex space-y-2 flex-col">
              <label className="font-semibold" htmlFor="confirm-password">
                City
              </label>
              <input
                className="bg-white border p-2"
                type="text"
                placeholder="city"
                {...register("city", {
                  required: "city is required",
                })}
              />
            </div>
            {errors.city ? (
              <p className="text-red-500 text-sm">
                {errors.city.message}
              </p>
            ) : null}

            <div className="flex space-y-2 flex-col">
              <label className="font-semibold" htmlFor="confirm-password">
                Country
              </label>
              <input
                className="bg-white border p-2"
                type="text"
                placeholder="Country"
                {...register("country", {
                  required: "Please enter your country",
                })}
              />
            </div>

            {errors.country ? (
              <p className="text-red-500 text-sm">
                {errors.country.message}
              </p>
            ) : null}

            <div className="flex justify-between px-1">
              <button
                onClick={() => router.push(`/auth/signup-first`)}
                className="w-1/5 py-3 rounded-2xl bg-[#1814F3]  text-white "
              >
                back
              </button>
              <button
                onClick={handleSubmit(onSubmit)}
                className="w-1/5 py-3 rounded-2xl bg-[#1814F3]  text-white "
              >
                next
              </button>
            </div>

            <div className="space-y-6 mt-10">
              <h2 className="font-poppins  text-[#7C8493]">
                Already have an account?
                <Link href={`/login`}>
                  <span className="font-bold text-indigo-800"> Login</span>
                </Link>
              </h2>
            </div>
          </form>
        </div>
      </div>
    </div>
  );
};

export default Page;
