"use client";
import React from "react";
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

export const SignUpForm = () => {
  const form = useForm({
    resolver: zodResolver(signUpSchema),
    defaultValues: {
      name: "",
      email: "",
      dateOfBirth: "",
      permanentAddress: "",
      postalCode: "",
      userName: "",
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

  const onSubmit = (values: z.infer<typeof signUpSchema>) => {
    console.log("123");
    console.log(values);
  };

  return (
    <div className="flex min-h-screen w-full max-w-[420px] flex-col justify-center gap-5 py-10 md:gap-8">
      <div className="flex items-center gap-2 ">
        <Image src="/icons/logo.png" width={25} height={25} alt="logo" />
        <h1 className="text-primaryBlack font-[900] text-[1.5rem]">
          BankDash.
        </h1>
      </div>
      <h1 className="gap-2 text-2xl text-primaryBlack font-bold">Register</h1>
      <Form {...form}>
        <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-4">
          {/* Name */}
          <FormField
            control={form.control}
            name="name"
            render={({ field }: { field: object }) => (
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
            render={({ field }: { field: object }) => (
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
          <div className="md:flex gap-5">
            {/* Postal Code */}
            <FormField
              control={form.control}
              name="postalCode"
              render={({ field }: { field: object }) => (
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
              render={({ field }: { field: object }) => (
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
          </div>

          {/* Postal Code */}
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
                      selected={field.value ? new Date(field.value) : undefined}
                      onSelect={(date) => {
                        field.onChange(date ? format(date, "yyyy-MM-dd") : "");
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
            name="userName"
            render={({ field }: { field: object }) => (
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
            render={({ field }: { field: object }) => (
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
          {/* City */}
          <div className="md:flex gap-5">
            <FormField
              control={form.control}
              name="city"
              render={({ field }: { field: object }) => (
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
              render={({ field }: { field: object }) => (
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
          </div>
          {/* Country */}
          <FormField
            control={form.control}
            name="country"
            render={({ field }: { field: object }) => (
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
          <FormField
            control={form.control}
            name="profilePicture"
            render={({ field }: { field: object }) => (
              <FormItem className="mb-1">
                <FormLabel>Profile Picture</FormLabel>
                <FormControl>
                  <Input
                    className="rounded-2xl md:w-[400px]"
                    placeholder="Profile URL"
                    {...field}
                  />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
          <div className="flex gap-5">
            {/* Preferences - Currency */}
            <FormField
              control={form.control}
              name="preference.currency"
              render={({ field }: { field: object }) => (
                <FormItem className="mb-1">
                  <FormLabel>Preferred Currency</FormLabel>
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
            {/* Preferences - Time Zone */}
            <FormField
              control={form.control}
              name="preference.timeZone"
              render={({ field }: { field: object }) => (
                <FormItem className="mb-1">
                  <FormLabel>Time Zone</FormLabel>
                  <FormControl>
                    <Input
                      className="rounded-2xl md:w-[400px]"
                      placeholder="GMT-5"
                      {...field}
                    />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
          </div>
          {/* Preferences - Sent or Receive Digital Currency */}
          <FormField
            control={form.control}
            name="preference.sentOrReceiveDigitalCurrency"
            render={({ field }) => (
              <FormItem className="flex items-end gap-2">
                <Switch
                  className="data-[state=checked]:bg-[#16DBCC]"
                  checked={field.value}
                  onCheckedChange={field.onChange}
                />
                <div className="space-y-0.5">
                  Send or Receive Digital Currency
                </div>
              </FormItem>
            )}
          />
          {/* Preferences - Receive Merchant Order */}
          <FormField
            control={form.control}
            name="preference.receiveMerchantOrder"
            render={({ field }) => (
              <FormItem className="flex items-end gap-2">
                <Switch
                  className="data-[state=checked]:bg-[#16DBCC]"
                  checked={field.value}
                  onCheckedChange={field.onChange}
                />
                <div className="space-y-0.5">Receive merchant order</div>
              </FormItem>
            )}
          />
          {/* Preferences - Account Recommendations */}
          <FormField
            control={form.control}
            name="preference.accountRecommendations"
            render={({ field }) => (
              <FormItem className="flex items-center md:items-end gap-2 ">
                <Switch
                  className="data-[state=checked]:bg-[#16DBCC]"
                  checked={field.value}
                  onCheckedChange={field.onChange}
                />
                <div className="space-y-0.5">Account Recommendations</div>
              </FormItem>
            )}
          />
          {/* Preferences - Two Factor Authentication */}
          <FormField
            control={form.control}
            name="preference.twoFactorAuthentication"
            render={({ field }) => (
              <FormItem className="flex items-center md:items-end gap-2">
                <Switch
                  className="data-[state=checked]:bg-[#16DBCC]"
                  checked={field.value}
                  onCheckedChange={field.onChange}
                />
                <div className="space-y-0.5">
                  Enable or disable two factor authentication
                </div>
              </FormItem>
            )}
          />
          {/* Submit Button */}
          <Button
            type="submit"
            className="bg-primaryBlue rounded-2xl mt-12 w-full"
          >
            Register
          </Button>
        </form>
      </Form>
      <Link href="/auth/sign-in" className="text-sm text-center text-blue-400">
        Already have an account? Sign in
      </Link>
    </div>
  );
};

export default SignUpForm;
