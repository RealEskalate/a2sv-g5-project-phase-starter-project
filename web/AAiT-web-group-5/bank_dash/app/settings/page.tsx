"use client";
import React from "react";
import Input from "../../components/input";
import { useForm } from "react-hook-form";
import Top from "@/components/Top";

interface contactForm {
  name: string;
  email: string;
  date: string;
  PermanentAddress: string;
  postal: string;
  username: string;
  password: string;
  CurrentAddress: string;
  city: string;
  country: string;
}

export default function Setting() {
  const { register, handleSubmit, formState } = useForm<contactForm>();
  const { errors } = formState;

  return (
    <div className="flex flex-col items-center">
      <Top topicName="Setting" />
      <div className="px-8 py-6 bg-slate-100 w-full">
        <div className="rounded-xl bg-white px-9 py-6 flex flex-col gap-5">
          <div className="flex gap-16 text-[#718EBF]">
            <div className="text-[#1814F3] border-b-[#1814F3] border-b-[3px] h-8">
              Edit Profile
            </div>
            <div className=" hover:text-[#5b5bce]  h-8">Preferences</div>
            <div className=" hover:text-[#5b5bce]  h-8">Security</div>
          </div>
          <div className="flex gap-8">
            <div className="">
              <img
                className="rounded-full mt-5"
                src="https://placehold.co/130"
                alt=""
              />
            </div>
            <div className="flex gap-7 w-full">
              <div className="w-full">
                <Input
                  field="Your Name"
                  namee="name"
                  placeholder="Charlene Reed "
                  regex={/^[A-Za-z0-9\s]+$/}
                  regexMsg="Incorrect name format"
                  minLength={2}
                  requiredMsg="Name field is required"
                  errors={errors}
                  register={register}
                  mode="input"
                />
                <Input
                  field="Email"
                  namee="email"
                  placeholder="charlenereed@gmail.com "
                  regex={
                    /^[a-z0-9!#$%&'*+/=?^_`{|}~-]+(?:\.[a-z0-9!#$%&'*+/=?^_`{|}~-]+)*@(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?$/
                  }
                  regexMsg="Incorrect email format"
                  minLength={0}
                  requiredMsg="Email field is required"
                  errors={errors}
                  register={register}
                  mode="input"
                />
                <Input
                  field="Date of Birth "
                  namee="date"
                  placeholder="25 January 1990"
                  regex={/^[A-Za-z0-9\s]+$/}
                  regexMsg="Incorrect name format"
                  minLength={2}
                  requiredMsg="Name field is required"
                  errors={errors}
                  register={register}
                  mode="input"
                />
                <Input
                  field="Permanent Address "
                  namee="PermanentAddress"
                  placeholder="San Jose, California, USA"
                  regex={/^[A-Za-z0-9\s]+$/}
                  regexMsg="Incorrect name format"
                  minLength={2}
                  requiredMsg="Name field is required"
                  errors={errors}
                  register={register}
                  mode="input"
                />
                <Input
                  field="Postal Code"
                  namee="postal"
                  placeholder="45962"
                  regex={/^[A-Za-z0-9\s]+$/}
                  regexMsg="Incorrect name format"
                  minLength={2}
                  requiredMsg="Name field is required"
                  errors={errors}
                  register={register}
                  mode="input"
                />
              </div>
              <div className="w-full">
                <Input
                  field="User Name"
                  namee="username"
                  placeholder="Charlene Reed "
                  regex={/^[A-Za-z0-9\s]+$/}
                  regexMsg="Incorrect name format"
                  minLength={2}
                  requiredMsg="Name field is required"
                  errors={errors}
                  register={register}
                  mode="input"
                />
                <Input
                  field="Password"
                  namee="password"
                  placeholder="**********"
                  regex={/^[A-Za-z0-9\s]+$/}
                  regexMsg="Incorrect name format"
                  minLength={2}
                  requiredMsg="Name field is required"
                  errors={errors}
                  register={register}
                  mode="input"
                />
                <Input
                  field="Present Address"
                  namee="CurrentAddress"
                  placeholder="San Jose, California, USA"
                  regex={/^[A-Za-z0-9\s]+$/}
                  regexMsg="Incorrect name format"
                  minLength={2}
                  requiredMsg="Name field is required"
                  errors={errors}
                  register={register}
                  mode="input"
                />
                <Input
                  field="City"
                  namee="city"
                  placeholder="San Jose"
                  regex={/^[A-Za-z0-9\s]+$/}
                  regexMsg="Incorrect name format"
                  minLength={2}
                  requiredMsg="Name field is required"
                  errors={errors}
                  register={register}
                  mode="input"
                />
                <Input
                  field="Country"
                  namee="country"
                  placeholder="USA"
                  regex={/^[A-Za-z0-9\s]+$/}
                  regexMsg="Incorrect name format"
                  minLength={2}
                  requiredMsg="Name field is required"
                  errors={errors}
                  register={register}
                  mode="input"
                />
              </div>
            </div>
          </div>
          <div className="flex justify-end text-white">
            <div className="bg-blue-800 py-3 px-16 rounded-lg">Save</div>
          </div>
        </div>
      </div>
    </div>
  );
}
