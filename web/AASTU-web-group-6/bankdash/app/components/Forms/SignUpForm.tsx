"use client";
import React from 'react';
import { useForm } from 'react-hook-form';
import UserValue from '@/types/UserValue';
import { z } from 'zod';
import { zodResolver } from '@hookform/resolvers/zod';
import SignupService from '@/app/Services/SignupService';
import ResponseValue from '@/types/ResponseValue';

// Define the schema using Zod
const schema = z.object({
  name: z.string().min(1, "Name is required"),
  email: z.string().email("Invalid email address"),
  dateOfBirth: z.string().min(1, "Date of Birth is required"),
  permanentAddress: z.string(),
  postalCode: z.string(),
  username: z.string().min(1, "Username is required"),
  password: z.string().min(6, "Password must be at least 6 characters"),
  confirmPassword: z.string().min(6, "Confirm Password is required"),
  presentAddress: z.string(),
  city: z.string(),
  country: z.string(),
  profilePicture: z.string().url("Invalid URL"),
  preference: z.object({
    currency: z.string(),
    sentOrReceiveDigitalCurrency: z.boolean().optional(),
    receiveMerchantOrder: z.boolean().optional(),
    accountRecommendations: z.boolean().optional(),
    timeZone: z.string(),
    twoFactorAuthentication: z.boolean().optional(),
  }),
}).refine((data) => data.password === data.confirmPassword, {
  path: ["confirmPassword"],
  message: "Passwords do not match",
});

const SignUpForm = () => {
  const { register, handleSubmit, formState: { errors } } = useForm<UserValue>({
    resolver: zodResolver(schema),
  });

  const onSubmit = async (data: UserValue) => {
    try {
      const responseData: ResponseValue = await SignupService(data);
      if (responseData.success) {
        console.log("Signup successful:");
      } else {
        console.error("Signup failed:");
      }
    } catch (error) {
      console.error("An error occurred:", error);
    }
  };

  return (
    <form onSubmit={handleSubmit(onSubmit)} className="space-y-4 p-6 bg-white rounded shadow-md max-w-md mx-auto">
      <h2 className="text-2xl font-bold mb-4">Signup</h2>

      <div>
        <label className="block">Name</label>
        <input {...register("name")} className="input" />
        {errors.name && <p className="text-red-500">{errors.name.message}</p>}
      </div>

      <div>
        <label className="block">Email</label>
        <input {...register("email")} type="email" className="input" />
        {errors.email && <p className="text-red-500">{errors.email.message}</p>}
      </div>

      <div>
        <label className="block">Date of Birth</label>
        <input {...register("dateOfBirth")} type="date" className="input" />
        {errors.dateOfBirth && <p className="text-red-500">{errors.dateOfBirth.message}</p>}
      </div>

      <div>
        <label className="block">Permanent Address</label>
        <input {...register("permanentAddress")} className="input" />
        {errors.permanentAddress && <p className="text-red-500">{errors.permanentAddress.message}</p>}
      </div>

      <div>
        <label className="block">Postal Code</label>
        <input {...register("postalCode")} className="input" />
        {errors.postalCode && <p className="text-red-500">{errors.postalCode.message}</p>}
      </div>

      <div>
        <label className="block">Username</label>
        <input {...register("username")} className="input" />
        {errors.username && <p className="text-red-500">{errors.username.message}</p>}
      </div>

      <div>
        <label className="block">Password</label>
        <input {...register("password")} type="password" className="input" />
        {errors.password && <p className="text-red-500">{errors.password.message}</p>}
      </div>

      <div>
        <label className="block">Confirm Password</label>
        <input {...register("confirmPassword")} type="password" className="input" />
        {errors.confirmPassword && <p className="text-red-500">{errors.confirmPassword.message}</p>}
      </div>

      <div>
        <label className="block">Present Address</label>
        <input {...register("presentAddress")} className="input" />
        {errors.presentAddress && <p className="text-red-500">{errors.presentAddress.message}</p>}
      </div>

      <div>
        <label className="block">City</label>
        <input {...register("city")} className="input" />
        {errors.city && <p className="text-red-500">{errors.city.message}</p>}
      </div>

      <div>
        <label className="block">Country</label>
        <input {...register("country")} className="input" />
        {errors.country && <p className="text-red-500">{errors.country.message}</p>}
      </div>

      <div>
        <label className="block">Profile Picture URL</label>
        <input {...register("profilePicture")} className="input" />
        {errors.profilePicture && <p className="text-red-500">{errors.profilePicture.message}</p>}
      </div>

      <h3 className="text-xl font-semibold mt-6">Preferences</h3>

      <div>
        <label className="block">Currency</label>
        <input {...register("preference.currency")} className="input" />
        {errors.preference?.currency && <p className="text-red-500">{errors.preference.currency.message}</p>}
      </div>

      <div>
        <label className="block">Sent or Receive Digital Currency</label>
        <input {...register("preference.sentOrReceiveDigitalCurrency")} type="checkbox" />
        {errors.preference?.sentOrReceiveDigitalCurrency && (
          <p className="text-red-500">{errors.preference.sentOrReceiveDigitalCurrency.message}</p>
        )}
      </div>

      <div>
        <label className="block">Receive Merchant Order</label>
        <input {...register("preference.receiveMerchantOrder")} type="checkbox" />
        {errors.preference?.receiveMerchantOrder && (
          <p className="text-red-500">{errors.preference.receiveMerchantOrder.message}</p>
        )}
      </div>

      <div>
        <label className="block">Account Recommendations</label>
        <input {...register("preference.accountRecommendations")} type="checkbox" />
        {errors.preference?.accountRecommendations && (
          <p className="text-red-500">{errors.preference.accountRecommendations.message}</p>
        )}
      </div>

      <div>
        <label className="block">Time Zone</label>
        <input {...register("preference.timeZone")} className="input" />
        {errors.preference?.timeZone && <p className="text-red-500">{errors.preference.timeZone.message}</p>}
      </div>

      <div>
        <label className="block">Two-Factor Authentication</label>
        <input {...register("preference.twoFactorAuthentication")} type="checkbox" />
        {errors.preference?.twoFactorAuthentication && (
          <p className="text-red-500">{errors.preference.twoFactorAuthentication.message}</p>
        )}
      </div>

      <button type="submit" className="w-full py-2 mt-6 bg-blue-500 text-white rounded">
        Signup
      </button>
    </form>
  );
}

export default SignUpForm;
