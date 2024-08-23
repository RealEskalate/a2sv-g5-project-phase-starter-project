"use client";
import React, { useState } from "react";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { z } from "zod";
import {
  Form,
  FormField,
  FormItem,
  FormLabel,
  FormControl,
  FormMessage,
} from "@/components/ui/form";
import { format } from "date-fns";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from "@/components/ui/popover";
import { CalendarIcon } from "@radix-ui/react-icons";
import { Calendar } from "@/components/ui/calendar";
import { Switch } from "@/components/ui/switch";
import Image from "next/image";
import Link from "next/link";
import { signUpSchema } from "@/schema/index";
import ky from "ky";
import { useRouter } from "next/navigation";
import ProfileUpload from "@/app/dashboard/_components/ProfileUpload";
import Stepper from "./Stepper";

const steps = [
  "Personal Information",
  "Address & Profile Picture",
  "Preferences",
  "Submit",
];

export const SignUpForm = () => {
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState("");
  const [profilePictureUrl, setProfilePictureUrl] = useState("");
  const [currentStep, setCurrentStep] = useState(0);
  const router = useRouter();

  const form = useForm({
    resolver: zodResolver(signUpSchema),
    defaultValues: {
      name: "",
      email: "",
      dateOfBirth: "",
      permanentAddress: "",
      postalCode: "",
      username: "",
      password: "",
      presentAddress: "",
      city: "",
      country: "",
      profilePicture: "",
      preference: {
        currency: "",
        sentOrReceiveDigitalCurrency: false,
        receiveMerchantOrder: false,
        accountRecommendations: false,
        timeZone: "",
        twoFactorAuthentication: false,
      },
    },
  });

  const onSubmit = async (values: z.infer<typeof signUpSchema>) => {
    setLoading(true);
    try {
      const payload = {
        ...values,
        profilePicture: profilePictureUrl,
      };

      console.log(payload);
      const res: { success: boolean; message: string } = await ky
        .post("https://bank-dashboard-o9tl.onrender.com/auth/register", {
          json: payload,
        })
        .json();
      if (res.success) {
        router.push("/auth/sign-in");
      } else {
        setError(res.message);
      }
    } catch (error) {
      if (error) {
        const errorResponse = await (error as any)?.response.json();
        setError(errorResponse.message);
        console.error("Error response:", errorResponse);
      } else {
        console.error("Unexpected error:", error);
      }
    } finally {
      setLoading(false);
    }
  };

  const handleNext = async () => {
    let isValid = true;

    console.log("Current Step:", currentStep);

    if (currentStep === 0) {
      console.log("Triggering validation for Step 1");
      isValid = await form.trigger([
        "name",
        "email",
        "dateOfBirth",
        "username",
        "password",
      ]);
    } else if (currentStep === 1) {
      console.log("Triggering validation for Step 2");
      isValid = false;
      isValid = await form.trigger([
        "postalCode",
        "permanentAddress",
        "city",
        "presentAddress",
        "country",
      ]);
    } else if (currentStep === 2) {
      console.log("Triggering validation for Step 3");
      isValid = false;
      isValid = await form.trigger([
        "preference.timeZone",
        "preference.currency",
        "preference.sentOrReceiveDigitalCurrency",
        "preference.receiveMerchantOrder",
        "preference.accountRecommendations",
        "preference.twoFactorAuthentication",
      ]);
    }

    console.log("Is Valid:", isValid);

    if (isValid) {
      setCurrentStep((prevStep) => prevStep + 1);
    } else {
      console.log("Validation failed. Please check the fields.");
    }
  };

  const handleBack = () => {
    setCurrentStep((prevStep) => prevStep - 1);
  };

  return (
    <div className="flex min-h-screen w-full max-w-[420px] flex-col justify-center gap-5 py-10 md:gap-8">
      <div className="flex items-center gap-2">
        <Image src="/icons/logo.png" width={25} height={25} alt="logo" />
        <h1 className="text-primaryBlack font-[900] text-[1.5rem]">
          BankDash.
        </h1>
      </div>

      <h1 className="gap-2 text-2xl text-primaryBlack font-bold">
        {steps[currentStep]}
      </h1>

      <Stepper current={steps[currentStep]} />

      <Form {...form}>
        <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-4">
          {currentStep === 0 && (
            <>
              {/* Name */}
              <FormField
                control={form.control}
                name="name"
                render={({ field }) => (
                  <FormItem className="mb-1">
                    <FormLabel>Name</FormLabel>
                    <FormControl>
                      <Input
                        className="rounded-2xl md:w-[400px]"
                        placeholder="John Doe"
                        {...field}
                      />
                    </FormControl>
                    <FormMessage />
                  </FormItem>
                )}
              />
              {/* Email */}
              <FormField
                control={form.control}
                name="email"
                render={({ field }) => (
                  <FormItem className="mb-1">
                    <FormLabel>Email</FormLabel>
                    <FormControl>
                      <Input
                        className="rounded-2xl md:w-[400px]"
                        placeholder="johndoe@example.com"
                        type="email"
                        {...field}
                      />
                    </FormControl>
                    <FormMessage />
                  </FormItem>
                )}
              />
              {/* Date of Birth */}
              <FormField
                control={form.control}
                name="dateOfBirth"
                render={({ field }) => (
                  <FormItem className="mb-1">
                    <FormLabel className="block my-2">Date of Birth</FormLabel>
                    <Popover>
                      <PopoverTrigger asChild>
                        <FormControl>
                          <Button
                            variant={"outline"}
                            className="w-full rounded-2xl text-left text-[#77809c] font-[400] flex justify-start"
                          >
                            {field.value
                              ? format(new Date(field.value), "PPP")
                              : "25 January 1990"}
                            <CalendarIcon className="ml-auto h-4 w-4" />
                          </Button>
                        </FormControl>
                      </PopoverTrigger>
                      <PopoverContent align="start">
                        <Calendar
                          mode="single"
                          selected={
                            field.value ? new Date(field.value) : undefined
                          }
                          onSelect={(date) => {
                            field.onChange(
                              date ? format(date, "yyyy-MM-dd") : ""
                            );
                          }}
                          disabled={(date) =>
                            date > new Date() || date < new Date("1900-01-01")
                          }
                          initialFocus
                        />
                      </PopoverContent>
                    </Popover>
                    <FormMessage />
                  </FormItem>
                )}
              />
              {/* Username */}
              <FormField
                control={form.control}
                name="username"
                render={({ field }) => (
                  <FormItem className="mb-1">
                    <FormLabel>Username</FormLabel>
                    <FormControl>
                      <Input
                        className="rounded-2xl md:w-[400px]"
                        placeholder="john_doe"
                        {...field}
                      />
                    </FormControl>
                    <FormMessage />
                  </FormItem>
                )}
              />
              {/* Password */}
              <FormField
                control={form.control}
                name="password"
                render={({ field }) => (
                  <FormItem className="mb-1">
                    <FormLabel>Password</FormLabel>
                    <FormControl>
                      <Input
                        className="rounded-2xl md:w-[400px]"
                        placeholder="********"
                        type="password"
                        {...field}
                      />
                    </FormControl>
                    <FormMessage />
                  </FormItem>
                )}
              />
            </>
          )}
          {currentStep === 1 && (
            <>
              {/* Postal Code */}
              <FormField
                control={form.control}
                name="postalCode"
                render={({ field }) => (
                  <FormItem className="mb-1">
                    <FormLabel>Postal Code</FormLabel>
                    <FormControl>
                      <Input
                        className="rounded-2xl md:w-[400px]"
                        placeholder="12345"
                        {...field}
                      />
                    </FormControl>
                    <FormMessage />
                  </FormItem>
                )}
              />
              {/* Permanent Address */}
              <FormField
                control={form.control}
                name="permanentAddress"
                render={({ field }) => (
                  <FormItem className="mb-1">
                    <FormLabel>Permanent Address</FormLabel>
                    <FormControl>
                      <Input
                        className="rounded-2xl md:w-[400px]"
                        placeholder="123 Main St"
                        {...field}
                      />
                    </FormControl>
                    <FormMessage />
                  </FormItem>
                )}
              />
              {/* City */}
              <FormField
                control={form.control}
                name="city"
                render={({ field }) => (
                  <FormItem className="mb-1">
                    <FormLabel>City</FormLabel>
                    <FormControl>
                      <Input
                        className="rounded-2xl md:w-[400px]"
                        placeholder="New York"
                        {...field}
                      />
                    </FormControl>
                    <FormMessage />
                  </FormItem>
                )}
              />
              {/* Present Address */}
              <FormField
                control={form.control}
                name="presentAddress"
                render={({ field }) => (
                  <FormItem className="mb-1">
                    <FormLabel>Present Address</FormLabel>
                    <FormControl>
                      <Input
                        className="rounded-2xl md:w-[400px]"
                        placeholder="456 Elm St"
                        {...field}
                      />
                    </FormControl>
                    <FormMessage />
                  </FormItem>
                )}
              />
              {/* Country */}
              <FormField
                control={form.control}
                name="country"
                render={({ field }) => (
                  <FormItem className="mb-1">
                    <FormLabel>Country</FormLabel>
                    <FormControl>
                      <Input
                        className="rounded-2xl md:w-[400px]"
                        placeholder="USA"
                        {...field}
                      />
                    </FormControl>
                    <FormMessage />
                  </FormItem>
                )}
              />
              {/* Profile Picture */}
              <ProfileUpload setProfilePictureUrl={setProfilePictureUrl} />
            </>
          )}
          {currentStep === 2 && (
            <>
              {/* Time Zone */}
              <FormField
                control={form.control}
                name="preference.timeZone"
                render={({ field }) => (
                  <FormItem className="mb-1">
                    <FormLabel>Time Zone</FormLabel>
                    <FormControl>
                      <Input
                        className="rounded-2xl md:w-[400px]"
                        placeholder="GMT+0"
                        {...field}
                      />
                    </FormControl>
                    <FormMessage />
                  </FormItem>
                )}
              />
              {/* Currency */}
              <FormField
                control={form.control}
                name="preference.currency"
                render={({ field }) => (
                  <FormItem className="mb-1">
                    <FormLabel>Currency</FormLabel>
                    <FormControl>
                      <Input
                        className="rounded-2xl md:w-[400px]"
                        placeholder="USD"
                        {...field}
                      />
                    </FormControl>
                    <FormMessage />
                  </FormItem>
                )}
              />
              {/* Notifications */}
              <FormField
                control={form.control}
                name="preference.sentOrReceiveDigitalCurrency"
                render={({ field }) => (
                  <FormItem className="mb-1 flex items-center gap-2">
                    <FormControl>
                      <Switch {...field} value={field.value.toString()} />
                    </FormControl>
                    <FormLabel>Send or Receive Digital Currency</FormLabel>
                    <FormMessage />
                  </FormItem>
                )}
              />
              <FormField
                control={form.control}
                name="preference.receiveMerchantOrder"
                render={({ field }) => (
                  <FormItem className="mb-1 flex items-center gap-2">
                    <FormControl>
                      <Switch {...field} value={field.value.toString()} />
                    </FormControl>
                    <FormLabel>Receive Merchant Order</FormLabel>
                    <FormMessage />
                  </FormItem>
                )}
              />
              <FormField
                control={form.control}
                name="preference.accountRecommendations"
                render={({ field }) => (
                  <FormItem className="mb-1 flex items-center gap-2">
                    <FormControl>
                      <Switch {...field} value={field.value.toString()} />
                    </FormControl>
                    <FormLabel>Account Recommendations</FormLabel>
                    <FormMessage />
                  </FormItem>
                )}
              />
              <FormField
                control={form.control}
                name="preference.twoFactorAuthentication"
                render={({ field }) => (
                  <FormItem className="mb-1 flex items-center gap-2">
                    <FormControl>
                      <Switch {...field} value={field.value.toString()} />
                    </FormControl>
                    <FormLabel>Two-Factor Authentication</FormLabel>
                    <FormMessage />
                  </FormItem>
                )}
              />
            </>
          )}
          {currentStep === 3 && (
            <div className="flex flex-col gap-2">
              <p className="text-red-600">{error}</p>
              <p className="text-green-600">
                Review all details before submitting
              </p>
            </div>
          )}
          <div className="flex justify-between mt-4">
            {currentStep > 0 && (
              <Button
                type="button"
                variant="outline"
                onClick={handleBack}
                className="bg-[#343C6A]"
              >
                Back
              </Button>
            )}
            {currentStep < steps.length - 1 ? (
              <Button
                type="button"
                onClick={handleNext}
                disabled={loading}
                className="bg-[#343C6A]"
              >
                Next
              </Button>
            ) : (
              <Button type="submit" disabled={loading} className="bg-[#343C6A]">
                Submit
              </Button>
            )}
          </div>
        </form>
      </Form>
    </div>
  );
};
