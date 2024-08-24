"use client";
import InputGroup from "./InputGroup";

import { z } from "zod";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";

import {
  useGetProfileQuery,
  useUpdateUserMutation,
} from "@/lib/redux/api/profileAPI";

import { setProfile } from "@/lib/redux/slices/profileSlice";
import { useEffect } from "react";
import { UpdatedUser } from "@/types/user.types";
import { useAppDispatch, useAppSelector } from "@/hooks/hoooks";
import { Loader } from "lucide-react";

const editProfileSchema = z.object({
  name: z.string().min(1, "Name is required"),
  username: z.string().min(1, "Username is required"),
  email: z.string().email("Invalid email address"),
  dateOfBirth: z.string().refine((date) => !isNaN(Date.parse(date)), {
    message: "Invalid date format",
  }),
  presentAddress: z.string().min(1, "Present address is required"),
  permanentAddress: z.string().min(1, "Permanent address is required"),
  city: z.string().min(1, "City is required"),
  postalCode: z.string().min(1, "Postal code is required"),
  country: z.string().min(1, "Country is required"),
});

const EditProfileForm = () => {
  const dispatch = useAppDispatch();
  const getData = useAppSelector((state) => state.profile);

  const { refetch, data, error, isSuccess } = useGetProfileQuery();

  useEffect(() => {
    if (isSuccess && data) {
      dispatch(setProfile(data?.data));
      console.log(data.data);
    }
  }, [data, dispatch]);

  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm({
    resolver: zodResolver(editProfileSchema),
    defaultValues: {
      name: getData?.name,
      username: getData?.username,
      email: getData?.email,
      dateOfBirth: getData?.dateOfBirth.split("T")[0],
      presentAddress: getData?.presentAddress,
      permanentAddress: getData?.permanentAddress,
      city: getData?.city,
      postalCode: getData?.postalCode,
      country: getData?.country,
    },
  });

  if (error) {
    return <h1>An Error Occured..</h1>;
  }

  const [updateUser, { isLoading }] = useUpdateUserMutation();

  const onSubmit = (data: UpdatedUser) => {
    updateUser(data).then((res: any) => {
      dispatch(setProfile(res?.data?.data));
      alert("Profile updated successfully");
      refetch();
    });
  };

  return (
    <div className="w-full">
      {data ? (
        <form action="" onSubmit={handleSubmit(onSubmit)}>
          <div className="flex flex-col md:flex-row md:space-x-5">
            <InputGroup
              id="name"
              label="Your Name"
              inputType="text"
              registerName="name"
              register={register}
              placeholder={data?.data?.name}
            />
            <InputGroup
              id="username"
              label="User Name"
              inputType="text"
              registerName="username"
              register={register}
              placeholder={data?.data?.username}
            />
          </div>
          <div className="flex flex-col md:flex-row md:space-x-5">
            <InputGroup
              id="email"
              label="Email"
              inputType="text"
              registerName="email"
              register={register}
              placeholder={data?.data?.email}
            />
            <InputGroup
              id="password"
              label="Password"
              inputType="password"
              registerName="password"
              register={register}
              placeholder="*********************"
            />
          </div>
          <div className="flex flex-col md:flex-row md:space-x-5">
            <InputGroup
              id="dateOfBirth"
              label="Date Of Birth"
              inputType="date"
              registerName="dateOfBirth"
              register={register}
              placeholder={data?.data?.dateOfBirth}
            />
            <InputGroup
              id="presentAddress"
              label="Present Address"
              inputType="text"
              registerName="presentAddress"
              register={register}
              placeholder={data?.data?.presentAddress}
            />
          </div>
          <div className="flex flex-col md:flex-row md:space-x-5">
            <InputGroup
              id="permanentAddress"
              label="Permanent Address"
              inputType="text"
              registerName="permanentAddress"
              register={register}
              placeholder={data?.data?.permanentAddress}
            />
            <InputGroup
              id="city"
              label="City"
              inputType="text"
              registerName="city"
              register={register}
              placeholder={data?.data?.city}
            />
          </div>
          <div className="flex flex-col md:flex-row md:space-x-5">
            <InputGroup
              id="postalCode"
              label="Postal Code"
              inputType="text"
              registerName="postalCode"
              register={register}
              placeholder={data?.data?.postalCode}
            />
            <InputGroup
              id="country"
              label="Country"
              inputType="text"
              registerName="country"
              register={register}
              placeholder={data?.data?.country}
            />
          </div>
          <div className="flex justify-end">
            <button
              type="submit"
              className="bg-[#1814f3] text-white px-10 py-2 rounded-lg w-full md:w-auto mt-4"
            >
              {isLoading ? <Loader className="animate-spin" /> : "Save"}
            </button>
          </div>
        </form>
      ) : (
        <h1>Loading...</h1>
      )}
    </div>
  );
};

export default EditProfileForm;
