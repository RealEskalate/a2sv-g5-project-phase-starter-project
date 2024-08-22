'use client'
import InputGroup from "./InputGroup";

import { useForm } from "react-hook-form";
import { z } from "zod";
import { zodResolver } from "@hookform/resolvers/zod";

const editProfileSchema = z.object({
  name: z.string().min(1, "Name is required"),
  username: z.string().min(1, "Username is required"),
  email: z.string().email("Invalid email address"),
  password: z.string().min(6, "Password must be at least 6 characters long"),
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
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm({
    resolver: zodResolver(editProfileSchema),
  });

  const onSubmit = (data: any) => {
    console.log(data);
  };  

  return (
    <div className="w-full">
      <form action="" onSubmit={handleSubmit(onSubmit)}>
        <div className="flex flex-col md:flex-row md:space-x-5">
          <InputGroup
            id="name"
            label="Your Name"
            inputType="text"
            registerName="name"
            register={register}
            placeholder="Charlene Reed"
          />
          <InputGroup
            id="username"
            label="User Name"
            inputType="text"
            registerName="username"
            register={register}
            placeholder="Charlene Reed"
          />
        </div>
        <div className="flex flex-col md:flex-row md:space-x-5">
          <InputGroup
            id="email"
            label="Email"
            inputType="text"
            registerName="email"
            register={register}
            placeholder="charlene.reed@gmail.com"
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
            placeholder="25 January 1990"
          />
          <InputGroup
            id="presentAddress"
            label="Present Address"
            inputType="text"
            registerName="presentAddress"
            register={register}
            placeholder="San Jose, California, USA"
          />
        </div>
        <div className="flex flex-col md:flex-row md:space-x-5">
          <InputGroup
            id="permanentAddress"
            label="Permanent Address"
            inputType="text"
            registerName="permanentAddress"
            register={register}
            placeholder="San Jose, California, USA"
          />
          <InputGroup
            id="city"
            label="City"
            inputType="text"
            registerName="city"
            register={register}
            placeholder="San Jose"
          />
        </div>
        <div className="flex flex-col md:flex-row md:space-x-5">
          <InputGroup
            id="postalCode"
            label="Postal Code"
            inputType="text"
            registerName="postalCode"
            register={register}
            placeholder="45322"
          />
          <InputGroup
            id="country"
            label="Country"
            inputType="text"
            registerName="country"
            register={register}
            placeholder="USA"
          />
        </div>
        <div className="flex justify-end">
          <button
            type="submit"
            className="bg-[#1814f3] text-white px-10 py-2 rounded-lg w-full md:w-auto mt-4"
          >
            Save
          </button>
        </div>
      </form>
    </div>
  );
};

export default EditProfileForm;
