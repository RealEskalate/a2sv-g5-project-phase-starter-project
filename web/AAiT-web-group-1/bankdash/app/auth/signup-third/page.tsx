"use client";
import Link from "next/link";
import React, { useEffect, useState } from "react";
import { useForm } from "react-hook-form";
import { useRouter } from "next/navigation";
import { useSelector, useDispatch } from "react-redux";
import { formType } from "@/types/formType";
import { setform } from "@/lib/redux/slices/formSlice";
import Image from "next/image";
import logo from "../../../public/images/Logo.svg";
import { useSignUpMutation } from "@/lib/redux/api/authApi";

const Page = () => {
  const router = useRouter();
  const dispatch = useDispatch();
  const form = useForm<formType>();
  const { control, register, formState, handleSubmit } = form;
  let curr_form = useSelector((state: any) => state.form.value);
  const { errors } = formState;
  const [image, setImage] = useState<string>("");

  // rtk query hook
  const [signUp, { data, isLoading, isError }] = useSignUpMutation();

  if (isError) {
    // alert("Invalid information");
    // router.push(`/signup-first`)
  }
  if (isLoading) {
    return <h1 className="text-center text-lg mt-72">Loading ....</h1>;
  }

  const signUP = async () => {
    const user = curr_form;
    console.log("u:", user)

    try {
      const res = await signUp(user);
      const { data } = res;
      if (data.success) {
        alert("success");
        // router.push(`/`)
      }
      console.log(data.message);
    } catch (err) {
      console.error(err);
    }
  };

  const onSubmit = (form: formType) => {
    console.log(image);
    // curr_form = { ...curr_form, profilePicture: image };
    // dispatch(setLastForm(form));
    const decoy2 = { preference: form };
    // curr_form = { ...curr_form, ...decoy2 };
    console.log("c", curr_form);
    signUP();
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
            <div className="w-16 mx-2 h-1 bg-blue-700"></div>
            <li className=" border py-2 px-4 bg-blue-600 rounded-full">3</li>
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
              <label className="font-semibold" htmlFor="email">
                Add a profile picture
              </label>
              <input
                className="bg-white border p-2"
                type="text"
                placeholder="link for your profile picture"
                onChange={(e) => setImage(e.target.value)}
              />
            </div>

            <div className="flex space-y-2 flex-col">
              <label className="font-semibold" htmlFor="email">
                Currency
              </label>
              <input
                className="bg-white border p-2"
                type="text"
                placeholder="Currency"
                {...register("currency", { required: "currency is required" })}
              />
            </div>

            <div className="flex space-y-2 flex-col">
              <label className="font-semibold" htmlFor="email">
                Time Zone
              </label>
              <input
                className="bg-white border p-2"
                type="text"
                placeholder="timeZone"
                {...register("timeZone", { required: "currency is required" })}
              />
            </div>

            <div>
              <div className="flex mt-4 mb-4 space-x-6">
                <label htmlFor="">Send Or Receive Digital Currency</label>
                <input
                  className=""
                  type="checkbox"
                  {...register("sentOrReceiveDigitalCurrency")}
                />
              </div>

              <div className="flex mt-4 mb-4 space-x-6">
                <label htmlFor="">Receive Merchant Order</label>
                <input
                  className=""
                  type="checkbox"
                  {...register("receiveMerchantOrder", {})}
                />
              </div>

              <div className="flex mt-4 mb-4 space-x-6">
                <label htmlFor="">Account Recommendations</label>
                <input
                  className=""
                  type="checkbox"
                  {...register("accountRecommendations", {})}
                />
              </div>

              <div className="flex mt-4 mb-4 space-x-6">
                <label htmlFor="">Two FactorAuthentication</label>
                <input
                  className=""
                  type="checkbox"
                  {...register("twoFactorAuthentication", {})}
                />
              </div>
              {errors.twoFactorAuthentication ? (
                <p className="text-red-500 text-sm">
                  {errors.twoFactorAuthentication.message}
                </p>
              ) : null}
            </div>

            <div className="flex justify-between px-1">
              <button
                onClick={() => router.push(`/auth/signup-second`)}
                className="w-1/5 py-3 rounded-2xl bg-[#1814F3]  text-white "
              >
                BACK
              </button>
              <button
                onClick={handleSubmit(onSubmit)}
                className="w-1/5 py-3 rounded-2xl bg-[#1814F3]  text-white "
              >
                SUBMIT
              </button>
            </div>

            <div className="space-y-6 mt-10">
              <h2 className="font-poppins  text-[#7C8493]">
                Already have an account?
                <Link href={`/auth/login`}>
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
