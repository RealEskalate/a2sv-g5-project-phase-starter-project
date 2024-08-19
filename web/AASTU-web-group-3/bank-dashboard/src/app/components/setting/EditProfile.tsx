import React, { useRef } from "react";
import { useForm } from "react-hook-form";
import Image from "next/image";
import {circleWithPen,profilepic} from "@/../../public/Icons"

interface FormInput {
  name: string;
  username: string;
  email: string;
  password: string;
  dob: string;
  presentAddress: string;
  permanentAddress: string;
  city: string;
  postalCode: string;
  country: string;
  profileImage: FileList;
}

const EditProfile = () => {
  const { register, handleSubmit, setValue } = useForm<FormInput>({
    defaultValues: {
      name: "",
      username: "",
      email: "",
      password: "",
      dob: "",
      presentAddress: "",
      permanentAddress: "",
      city: "",
      postalCode: "",
      country: "",
    },
  });

  const fileInputRef = useRef<HTMLInputElement>(null);

  const onSubmit = async (data: FormInput) => {
    console.log("Form submitted:", data);
    const formData = new FormData();
    Object.keys(data).forEach((key) => {
      if (key !== "profileImage" || data.profileImage.length > 0) {
        formData.append(key, (data as any)[key]);
      }
    });

    if (data.profileImage && data.profileImage.length > 0) {
      formData.append("profileImage", data.profileImage[0]);
    }

    // Submit the formData object to your server or API
  };

  const handleImageClick = () => {
    if (fileInputRef.current) {
      fileInputRef.current.click();
    }
  };

  const handleFileChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    const file = event.target.files;
    if (file && file.length > 0) {
      setValue("profileImage", file);
      console.log("Selected file:", file[0]);
    }
  };

  const formFields = [
    {
      id: "name",
      label: "Your Name",
      placeholder: "Charlene Reed",
      type: "text",
    },
    {
      id: "username",
      label: "User Name",
      placeholder: "Charlene Reed",
      type: "text",
    },
    {
      id: "email",
      label: "Email",
      placeholder: "charlenereed@gmail.com",
      type: "email",
    },
    {
      id: "password",
      label: "Password",
      placeholder: "**********",
      type: "password",
    },
    {
      id: "dob",
      label: "Date of Birth",
      placeholder: "Enter Date of Birth",
      type: "date",
    },
    {
      id: "presentAddress",
      label: "Present Address",
      placeholder: "San Jose, California, USA",
      type: "text",
    },
    {
      id: "permanentAddress",
      label: "Permanent Address",
      placeholder: "San Jose, California, USA",
      type: "text",
    },
    { id: "city", label: "City", placeholder: "San Jose", type: "text" },
    {
      id: "postalCode",
      label: "Postal Code",
      placeholder: "45962",
      type: "text",
    },
    { id: "country", label: "Country", placeholder: "USA", type: "text" },
  ];

  return (
    <div className="p-4 flex flex-col md:flex-row gap-8">
      <div className="relative rounded-full w-64 h-64 mb-5 md:w-40 md:h-40">
        <Image
          src={profilepic}
          width={256}
          height={256}
          alt="profilepic"
          className="w-64 h-64 md:w-40 md:h-40 object-cover rounded-full"
        />
        <Image
          src={circleWithPen}
          alt="edit icon"
          width={64}
          height={64}
          className="absolute z-30 right-1 bottom-10 object-cover md:w-10 md:h-10 md:bottom-5 cursor-pointer"
          onClick={handleImageClick}
        />
        <input
          type="file"
          ref={fileInputRef}
          className="hidden"
          accept="image/*"
          onChange={handleFileChange}
        />
      </div>

      <form
        onSubmit={handleSubmit(onSubmit)}
        className="flex flex-wrap items-end justify-between w-full md:w-4/5"
      >
        {formFields.map((field) => (
          <div key={field.id} className="mb-3 w-full md:w-[45%]">
            <label
              className="block text-black text-sm  mb-2"
              htmlFor={field.id}
            >
              {field.label}
            </label>
            <input
              className="w-full p-3 md:p-2 text-[#718EBF] border-2 text-sm border-[#DFEAF2] rounded-lg focus:outline-none"
              type={field.type}
              id={field.id}
              placeholder={field.placeholder}
              {...register(field.id as keyof FormInput, {
                required: {
                  value: true,
                  message: `${field.label} is required`,
                },
              })}
            />
          </div>
        ))}
        <div className="flex justify-end w-full">
          <button
            className=" w-full md:w-1/5 bg-[#1814F3] text-white font-semibold py-2 px-4 rounded-lg focus:outline-none"
            type="submit"
          >
            Save
          </button>
        </div>
      </form>
    </div>
  );
};

export default EditProfile;
