"use client";

import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import { z } from "zod";

import { Button } from "@/components/ui/button";
import {
  Form,
  FormControl,
  FormDescription,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import { Switch } from "@/components/ui/switch";
import ky from "ky";
import { getSession, useSession } from "next-auth/react";

const Security = () => {
  const formSchema = z.object({
    twoFactor: z.boolean().default(true).optional(),
    currentPassword: z
      .string()
      .min(6, { message: "Password must be at least 6 characters." }),
    newPassword: z
      .string()
      .min(6, { message: "Password must be at least 6 characters." }),
  });

  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      twoFactor: true,
      currentPassword: "",
      newPassword: "",
    },
  });

  async function onSubmit(values: z.infer<typeof formSchema>) {
    const formData = {
      password: values.currentPassword,
      newPassword: values.newPassword,
    };

    const { data: session } = await useSession();
    const accessToken = session?.user?.accessToken;

    console.log(accessToken);

    if (!accessToken) {
      throw new Error("No access token found");
    }

    const res = await ky
      .post("https://bank-dashboard-6acc.onrender.com/auth/change_password", {
        json: formData,
        headers: {
          Authorization: `Bearer ${accessToken}`,
        },
      })
      .json();

    return res;
  }
  return (
    <div className="md:px-16">
      <Form {...form}>
        <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-8">
          <div className="mt-5 ">
            <h1 className="mb-2 text-primaryBlack font-bold">
              Two-factor Authentication
            </h1>
            <div className="flex flex-col gap-2">
              <FormField
                control={form.control}
                name="twoFactor"
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
            </div>

            <h1 className="mt-5 mb-2 text-primaryBlack font-bold">
              Change Password
            </h1>
            <FormField
              control={form.control}
              name="currentPassword"
              render={({ field }) => (
                <FormItem className="mb-1">
                  <FormLabel>Current Password</FormLabel>
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

            <FormField
              control={form.control}
              name="newPassword"
              render={({ field }) => (
                <FormItem className="mb-1">
                  <FormLabel>New Password</FormLabel>
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
          </div>
          <Button
            type="submit"
            className="mt-5 md:w-auto w-full px-8 float-end bg-primaryBlue text-white"
          >
            Save
          </Button>
        </form>
      </Form>
    </div>
  );
};

export default Security;
