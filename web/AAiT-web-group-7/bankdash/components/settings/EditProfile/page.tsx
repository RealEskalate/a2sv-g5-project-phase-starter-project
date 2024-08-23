"use client";
import React from "react";
import { useForm } from "react-hook-form";
type Form = {
  Name: string;
  Email: string;
  DOT: string;
  PA: string;
  PC: string;
  UN: string;
  password: string;
  PresentAddress: string;
  City: string;
  Country: string;
};
const EditProfile = () => {
  const form = useForm<Form>();
  const { register, handleSubmit, formState } = form;
  const { errors } = formState;
  const onSubmit = (data: Form) => {
    console.log(data);
  };
  return (
    <div className="flex flex-col text-sm">
      <div className="flex  gap-8 py-10">
        <div className="relative">
          <img src="pubimg/pp.png" className="rounded-full" />
          <button className="bg-[#1814F3] h-8 w-8 flex justify-center items-center rounded-full absolute right-0 top-20">
            <img src="pubimg/pencil.svg" />
          </button>
        </div>
        <form
          onSubmit={handleSubmit(onSubmit)}
          className="flex flex-col items-center gap-5"
        >
          <div className="flex gap-8">
          <div className="flex flex-col gap-3">
            <div className="flex flex-col items-start justify-center gap-2 ">
              <label className="text-[#232323] ">Your Name</label>
              <input
                id="name"
                {...register("Name", {
                  required: {
                    value: true,
                    message: "Full Name is required",
                  },
                })}
                placeholder="Full Name"
                className="w-[400px] rounded-xl p-3 border border-[#DFEAF2]  text-[#718EBF]"
              />
              <p
                className="text-red-600 flex text-xs font-semibold gap-1
            "
              >
                {errors.Name && (
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    fill="none"
                    viewBox="0 0 24 24"
                    stroke-width="1.5"
                    stroke="currentColor"
                    className="size-4"
                  >
                    <path
                      stroke-linecap="round"
                      stroke-linejoin="round"
                      d="M12 9v3.75m9-.75a9 9 0 1 1-18 0 9 9 0 0 1 18 0Zm-9 3.75h.008v.008H12v-.008Z"
                    />
                  </svg>
                )}
                {errors.Name?.message}{" "}
              </p>
            </div>
            <div className="flex flex-col items-start justify-center gap-2">
              <label className="text-[#232323] ">Email</label>
              <input
                placeholder="Email"
                type="email"
                id="email"
                {...register("Email", {
                  required: {
                    value: true,
                    message: "Email is required",
                  },
                  pattern: {
                    value: /^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/,
                    message: "Invalid Email",
                  },
                })}
                className="w-[400px] rounded-xl p-3 border border-[#DFEAF2]  text-[#718EBF]"
              />
              <p
                className="text-red-600 flex text-xs font-semibold gap-1
            "
              >
                {errors.Email && (
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    fill="none"
                    viewBox="0 0 24 24"
                    stroke-width="1.5"
                    stroke="currentColor"
                    className="size-4"
                  >
                    <path
                      stroke-linecap="round"
                      stroke-linejoin="round"
                      d="M12 9v3.75m9-.75a9 9 0 1 1-18 0 9 9 0 0 1 18 0Zm-9 3.75h.008v.008H12v-.008Z"
                    />
                  </svg>
                )}
                {errors.Email?.message}
              </p>
            </div>

            <div className="flex flex-col items-start justify-center gap-2">
              <label className="text-[#232323] ">Date of Birth</label>
              <input
                type="date"
                {...register("DOT", {
                  required: {
                    value: true,
                    message: "Date of Birth is required",
                  },
                })}
                className="w-[400px] rounded-xl p-3 border border-[#DFEAF2]  text-[#718EBF]"
              />
              <p
                className="text-red-600 flex  text-xs font-semibold gap-1
            "
              >
                {errors.DOT && (
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    fill="none"
                    viewBox="0 0 24 24"
                    stroke-width="1.5"
                    stroke="currentColor"
                    className="size-4"
                  >
                    <path
                      stroke-linecap="round"
                      stroke-linejoin="round"
                      d="M12 9v3.75m9-.75a9 9 0 1 1-18 0 9 9 0 0 1 18 0Zm-9 3.75h.008v.008H12v-.008Z"
                    />
                  </svg>
                )}
                {errors.DOT?.message}{" "}
              </p>
            </div>
            <div className="flex flex-col items-start justify-center gap-2">
              <label className="text-[#232323] ">Permanent Address</label>
              <input
                type="text"
                placeholder="Address"
                {...register("PA", {
                  required: {
                    value: true,
                    message: "Address is required",
                  },
                })}
                className="w-[400px] rounded-xl p-3 border border-[#DFEAF2]  text-[#718EBF]"
              />
              <p
                className="text-red-600 flex text-xs font-semibold gap-1
            "
              >
                {errors.PA && (
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    fill="none"
                    viewBox="0 0 24 24"
                    stroke-width="1.5"
                    stroke="currentColor"
                    className="size-4"
                  >
                    <path
                      stroke-linecap="round"
                      stroke-linejoin="round"
                      d="M12 9v3.75m9-.75a9 9 0 1 1-18 0 9 9 0 0 1 18 0Zm-9 3.75h.008v.008H12v-.008Z"
                    />
                  </svg>
                )}
                {errors.PA?.message}{" "}
              </p>
            </div>

            <div className="flex flex-col items-start justify-center gap-2">
              <label className="text-[#232323] ">Postal Code</label>
              <input
                type="text"
                placeholder="Code"
                {...register("PC", {
                  required: {
                    value: true,
                    message: "Postal Code is required",
                  },
                })}
                className="w-[400px] rounded-xl p-3 border border-[#DFEAF2]  text-[#718EBF]"
              />
              <p
                className="text-red-600 flex text-xs font-semibold gap-1
            "
              >
                {errors.PC && (
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    fill="none"
                    viewBox="0 0 24 24"
                    stroke-width="1.5"
                    stroke="currentColor"
                    className="size-4"
                  >
                    <path
                      stroke-linecap="round"
                      stroke-linejoin="round"
                      d="M12 9v3.75m9-.75a9 9 0 1 1-18 0 9 9 0 0 1 18 0Zm-9 3.75h.008v.008H12v-.008Z"
                    />
                  </svg>
                )}
                {errors.PC?.message}{" "}
              </p>
            </div>
          </div>

          <div className="flex flex-col gap-3">
            <div className="flex flex-col items-start justify-center gap-2">
              <label className="text-[#232323] ">User Name</label>
              <input
                type="text"
                {...register("UN", {
                  required: {
                    value: true,
                    message: "Username is required",
                  },
                })}
                placeholder="Name"
                className="w-[400px] rounded-xl p-3 border border-[#DFEAF2]  text-[#718EBF]"
              />
              <p
                className="text-red-600 flex text-xs  font-semibold gap-1
            "
              >
                {errors.UN && (
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    fill="none"
                    viewBox="0 0 24 24"
                    stroke-width="1.5"
                    stroke="currentColor"
                    className="size-4"
                  >
                    <path
                      stroke-linecap="round"
                      stroke-linejoin="round"
                      d="M12 9v3.75m9-.75a9 9 0 1 1-18 0 9 9 0 0 1 18 0Zm-9 3.75h.008v.008H12v-.008Z"
                    />
                  </svg>
                )}
                {errors.UN?.message}{" "}
              </p>
            </div>
            <div className="flex flex-col items-start justify-center gap-2">
              <label className="text-[#232323] ">Password</label>
              <input
                type="password"
                {...register("password", {
                  required: {
                    value: true,
                    message: "Password is required",
                  },
                  minLength: {
                    value: 8,
                    message: "Password must be at least 8 characters",
                  },
                })}
                placeholder="password"
                className="w-[400px] rounded-xl p-3 border border-[#DFEAF2]  text-[#718EBF]"
              />
              <p
                className="text-red-600 flex text-xs font-semibold gap-1
            "
              >
                {errors.password && (
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    fill="none"
                    viewBox="0 0 24 24"
                    stroke-width="1.5"
                    stroke="currentColor"
                    className="size-4"
                  >
                    <path
                      stroke-linecap="round"
                      stroke-linejoin="round"
                      d="M12 9v3.75m9-.75a9 9 0 1 1-18 0 9 9 0 0 1 18 0Zm-9 3.75h.008v.008H12v-.008Z"
                    />
                  </svg>
                )}
                {errors.password?.message}
              </p>
            </div>

            <div className="flex flex-col items-start justify-center gap-2">
              <label className="text-[#232323] ">Present Address</label>
              <input
                type="text"
                placeholder="Address"
                {...register("PresentAddress", {
                  required: {
                    value: true,
                    message: "Address is required",
                  },
                })}
                className="w-[400px] rounded-xl p-3 border border-[#DFEAF2]  text-[#718EBF]"
              />
              <p
                className="text-red-600 flex text-xs font-semibold gap-1
            "
              >
                {errors.PresentAddress && (
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    fill="none"
                    viewBox="0 0 24 24"
                    stroke-width="1.5"
                    stroke="currentColor"
                    className="size-4"
                  >
                    <path
                      stroke-linecap="round"
                      stroke-linejoin="round"
                      d="M12 9v3.75m9-.75a9 9 0 1 1-18 0 9 9 0 0 1 18 0Zm-9 3.75h.008v.008H12v-.008Z"
                    />
                  </svg>
                )}
                {errors.PresentAddress?.message}{" "}
              </p>
            </div>
            <div className="flex flex-col items-start justify-center gap-2">
              <label className="text-[#232323] ">City</label>
              <input
                type="text"
                {...register("City", {
                  required: {
                    value: true,
                    message: "City is required",
                  },
                })}
                placeholder="City"
                className="w-[400px] rounded-xl p-3 border border-[#DFEAF2]  text-[#718EBF]"
              />
              <p
                className="text-red-600 flex text-xs font-semibold gap-1
            "
              >
                {errors.City && (
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    fill="none"
                    viewBox="0 0 24 24"
                    stroke-width="1.5"
                    stroke="currentColor"
                    className="size-4"
                  >
                    <path
                      stroke-linecap="round"
                      stroke-linejoin="round"
                      d="M12 9v3.75m9-.75a9 9 0 1 1-18 0 9 9 0 0 1 18 0Zm-9 3.75h.008v.008H12v-.008Z"
                    />
                  </svg>
                )}
                {errors.City?.message}{" "}
              </p>
            </div>

            <div className="flex flex-col items-start justify-center gap-2">
              <label className="text-[#232323] ">Country</label>
              <input
                type="text"
                {...register("Country", {
                  required: {
                    value: true,
                    message: "Country is required",
                  },
                })}
                placeholder="Country"
                className="w-[400px] rounded-xl p-3 border border-[#DFEAF2]  text-[#718EBF]"
              />
              <p
                className="text-red-600 flex text-xs font-semibold gap-1
            "
              >
                {errors.Country && (
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    fill="none"
                    viewBox="0 0 24 24"
                    stroke-width="1.5"
                    stroke="currentColor"
                    className="size-4"
                  >
                    <path
                      stroke-linecap="round"
                      stroke-linejoin="round"
                      d="M12 9v3.75m9-.75a9 9 0 1 1-18 0 9 9 0 0 1 18 0Zm-9 3.75h.008v.008H12v-.008Z"
                    />
                  </svg>
                )}
                {errors.Country?.message}{" "}
              </p>
            </div>
          </div>
          </div>

          <div className="flex w-full justify-end mt-5 px-[30px] ">
              <button
                type="submit"
                className="px-10 py-3 text-white rounded-xl bg-[#1814F3]"
              >
                Save
              </button>
            </div>

        </form>
      </div>
    </div>
  );
};

export default EditProfile;
