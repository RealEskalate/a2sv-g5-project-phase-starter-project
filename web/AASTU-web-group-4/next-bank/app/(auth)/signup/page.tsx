"use client";
import { useState } from "react";
import { useForm } from "react-hook-form";
import { User } from "@/types/index";
import Image from "next/image";
import Cookie from "js-cookie";
import { registerUser } from "@/services/authentication";
import { uploadImage } from "@/components/Imageupload"; // Import your upload function
import { logo } from "../../../constants/index";
import Link from "next/link";

const SignupForm = () => {
  const [step, setStep] = useState(1); // Manage form steps
  const [imageUrl, setImageUrl] = useState(""); // Store image URL after upload

  const {
    register,
    handleSubmit,
    formState: { errors, isValid },
  } = useForm<User>({
    mode: "onChange", // Enable form validation on change
  });

  // Handle file change and upload image
  const handleFileChange = async (event: React.ChangeEvent<HTMLInputElement>) => {
    const file = event.target.files?.[0];
    if (file) {
      try {
        const url = await uploadImage(file);
        setImageUrl(url); // Set the uploaded image URL
        console.log("Image uploaded successfully:", url);
      } catch (error) {
        console.error("Error uploading image:", error);
      }
    }
  };

  // Handle form submission
  const onSubmit = async (data: User) => {
    try {
      const updatedData = {
        ...data,
        profilePicture: imageUrl || "", // Add profile picture if available
      };

      const registeredUser = await registerUser(updatedData); // Register the user
      console.log("Registered User:", registeredUser);

      // Store access and refresh tokens in cookies
      Cookie.set("accessToken", registeredUser.data.access_token);
      Cookie.set("refreshToken", registeredUser.data.refresh_token);

      window.location.href = "/"; // Redirect to homepage on success
    } catch (error) {
      console.error("Registration Error:", error);
    }
  };

  // Step navigation
  const handleNextStep = () => {
    if (isValid) setStep(step + 1); // Move to the next step if form is valid
  };
  const handlePreviousStep = () => setStep(step - 1); // Move to the previous step

  return (
    <form onSubmit={handleSubmit(onSubmit)} className="max-w-3xl mx-auto p-8 bg-white shadow-md rounded-lg">
      <div className="flex justify-center items-center mb-8">
        <Image src={logo.icon} alt="Logo" height={60} width={60} />
        <h1 className="font-bold text-3xl text-gray-700 font-serif p-2">
          <p className="text-gray-600">NEXT BANK</p>
        </h1>
      </div>

      <h2 className="text-blue-600 font-semibold text-2xl text-center mb-4">Sign Up</h2>
      <p className="text-center text-gray-600 mb-8">Please enter your details</p>

      {step === 1 && (
        <div className="grid gap-6 sm:grid-cols-1 md:grid-cols-2">
          {/* Step 1 Fields */}
          <div>
            <label htmlFor="name" className="block font-medium text-gray-700">Full Name</label>
            <input
              type="text"
              id="name"
              placeholder="ex: John"
              {...register("name", { required: "Name is required" })}
              className="border border-gray-300 rounded-lg p-3 mt-1 w-full focus:outline-none focus:border-blue-500"
            />
            {errors.name && <p className="text-red-500 mt-1">{errors.name.message}</p>}
          </div>

          <div>
            <label htmlFor="email" className="block font-medium text-gray-700">Email</label>
            <input
              type="email"
              id="email"
              placeholder="ex: example@gmail.com"
              {...register("email", {
                required: "Email is required",
                pattern: /^\S+@\S+\.\S+$/,
              })}
              className="border border-gray-300 rounded-lg p-3 mt-1 w-full focus:outline-none focus:border-blue-500"
            />
            {errors.email && <p className="text-red-500 mt-1">{errors.email.message}</p>}
          </div>

          <div>
            <label htmlFor="dateOfBirth" className="block font-medium text-gray-700">Date of Birth</label>
            <input
              type="date"
              id="dateOfBirth"
              {...register("dateOfBirth", { required: "Date of Birth is required" })}
              className="border border-gray-300 rounded-lg p-3 mt-1 w-full focus:outline-none focus:border-blue-500"
            />
            {errors.dateOfBirth && <p className="text-red-500 mt-1">{errors.dateOfBirth.message}</p>}
          </div>

          <div>
            <label htmlFor="username" className="block font-medium text-gray-700">Username</label>
            <input
              type="text"
              id="username"
              placeholder="ex: Star"
              {...register("username", { required: "Username is required" })}
              className="border border-gray-300 rounded-lg p-3 mt-1 w-full focus:outline-none focus:border-blue-500"
            />
            {errors.username && <p className="text-red-500 mt-1">{errors.username.message}</p>}
          </div>

          <div>
            <label htmlFor="password" className="block font-medium text-gray-700">Password</label>
            <input
              type="password"
              id="password"
              placeholder="ex: password"
              {...register("password", { required: "Password is required", minLength: 8 })}
              className="border border-gray-300 rounded-lg p-3 mt-1 w-full focus:outline-none focus:border-blue-500"
            />
            {errors.password && <p className="text-red-500 mt-1">{errors.password.message}</p>}
          </div>

          <div>
            <label htmlFor="permanentAddress" className="block font-medium text-gray-700">Permanent Address</label>
            <textarea
              id="permanentAddress"
              placeholder="ex: A.A"
              {...register("permanentAddress", { required: "Permanent Address is required" })}
              className="border border-gray-300 rounded-lg p-3 mt-1 w-full focus:outline-none focus:border-blue-500"
            />
            {errors.permanentAddress && <p className="text-red-500 mt-1">{errors.permanentAddress.message}</p>}
          </div>

          <div className="col-span-2 flex justify-center">
            <button type="button" onClick={handleNextStep} className="bg-blue-500 hover:bg-blue-600 text-white font-bold py-3 px-6 rounded-lg mt-8 transition duration-300">
              Continue
            </button>
          </div>

          <div className="col-span-2 flex justify-center">
            <p className="mx-2">Already have an account?</p>
            <Link className="text-blue-800" href="/signin">
              Login
            </Link>
          </div>
        </div>
      )}

      {step === 2 && (
        <div className="grid gap-6 sm:grid-cols-1 md:grid-cols-2">
          {/* Step 2 Fields */}
          <div>
            <label htmlFor="presentAddress" className="block font-medium text-gray-700">Present Address</label>
            <input
              type="text"
              id="presentAddress"
              placeholder="ex: Bola"
              {...register("presentAddress", { required: "Present Address is required" })}
              className="border border-gray-300 rounded-lg p-3 mt-1 w-full focus:outline-none focus:border-blue-500"
            />
            {errors.presentAddress && <p className="text-red-500 mt-1">{errors.presentAddress.message}</p>}
          </div>

          <div>
            <label htmlFor="city" className="block font-medium text-gray-700">City</label>
            <input
              type="text"
              id="city"
              placeholder="ex: A.A"
              {...register("city", { required: "City is required" })}
              className="border border-gray-300 rounded-lg p-3 mt-1 w-full focus:outline-none focus:border-blue-500"
            />
            {errors.city && <p className="text-red-500 mt-1">{errors.city.message}</p>}
          </div>

          <div>
            <label htmlFor="country" className="block font-medium text-gray-700">Country</label>
            <input
              type="text"
              id="country"
              placeholder="ex: Ethiopia"
              {...register("country", { required: "Country is required" })}
              className="border border-gray-300 rounded-lg p-3 mt-1 w-full focus:outline-none focus:border-blue-500"
            />
            {errors.country && <p className="text-red-500 mt-1">{errors.country.message}</p>}
          </div>

         

          <div>
            <label htmlFor="file" className="block font-medium text-gray-700">Profile Picture</label>
            <input
              type="file"
              id="file"
              accept="image/*"
              onChange={handleFileChange}
              className="border border-gray-300 rounded-lg p-3 mt-1 w-full focus:outline-none focus:border-blue-500"
            />
          </div>

          <div className="col-span-2 flex justify-between">
            <button type="button" onClick={handlePreviousStep} className="bg-gray-500 hover:bg-gray-600 text-white font-bold py-3 px-6 rounded-lg mt-8 transition duration-300">
              Back
            </button>
            <button type="submit" className="bg-blue-500 hover:bg-blue-600 text-white font-bold py-3 px-6 rounded-lg mt-8 transition duration-300">
              Submit
            </button>
          </div>
        </div>
      )}
    </form>
  );
};

export default SignupForm;
