"use client";

import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import { z } from "zod";
import { Button } from "@/components/ui/button";
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import { Switch } from "@/components/ui/switch";
import ky, { HTTPError } from "ky";
import { useState } from "react";
import { toast } from "sonner";
import { getSession } from "next-auth/react";
import { Currency } from "lucide-react";

const Preferences = () => {
  const [loading, setLoading] = useState(false);
  const formSchema = z.object({
    sentOrReceiveDigitalCurrency: z.boolean().default(true).optional(),
    receiveMerchantOrder: z.boolean().default(false).optional(),
    accountRecommendations: z.boolean().default(true).optional(),
    currency: z.string().min(1, "Currency is required"),
    timeZone: z.string().min(1, "Time Zone is required"),
  });

  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      currency: "",
      timeZone: "",
      sentOrReceiveDigitalCurrency: true,
      receiveMerchantOrder: false,
      accountRecommendations: true,
    },
  });

  const onSubmit = async (values: z.infer<typeof formSchema>) => {
    console.log(values);
    setLoading(true);
    try {
      const session = await getSession();
      const accessToken = session?.user.accessToken;
      const res = await ky.put(
        "https://bank-dashboard-o9tl.onrender.com/user/update-preference",
        {
          json: { ...values, twoFactorAuthentication: false },
          headers: {
            Authorization: `Bearer ${accessToken}`,
          },
        }
      );
      const data = await res.json();
      console.log("Response", data);
      toast("Update Successful");
    } catch (err) {
      if (err instanceof HTTPError && err.response) {
        const errorResponse = await err.response.json();
        console.error("Error Response", errorResponse);
      }
      console.error("Console Error", err);
      toast("Update Unsuccessful");
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="md:px-16">
      <Form {...form}>
        <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-8">
          <div className="md:flex gap-5">
            <FormField
              control={form.control}
              name="currency"
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
            <FormField
              control={form.control}
              name="timeZone"
              render={({ field }) => (
                <FormItem className="mb-1">
                  <FormLabel>Time Zone</FormLabel>
                  <FormControl>
                    <Input
                      className="rounded-2xl md:w-[400px]"
                      placeholder="(GMT-12:00) International Date Line West"
                      {...field}
                    />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
          </div>

          <div className="mt-5">
            <h1 className="mb-2 text-primaryBlack font-bold">Notification</h1>
            <div className="flex flex-col gap-2">
              <FormField
                control={form.control}
                name="sentOrReceiveDigitalCurrency"
                render={({ field }) => (
                  <FormItem className="flex items-end gap-2">
                    <Switch
                      className="data-[state=checked]:bg-[#16DBCC]"
                      checked={field.value}
                      onCheckedChange={field.onChange}
                    />
                    <label htmlFor={field.name} className="space-y-0.5">
                      I send or receive digital currency
                    </label>
                  </FormItem>
                )}
              />

              <FormField
                control={form.control}
                name="receiveMerchantOrder"
                render={({ field }) => (
                  <FormItem className="flex items-end gap-2">
                    <Switch
                      className="data-[state=checked]:bg-[#16DBCC]"
                      checked={field.value}
                      onCheckedChange={field.onChange}
                    />
                    <label htmlFor={field.name} className="space-y-0.5">
                      I receive merchant order
                    </label>
                  </FormItem>
                )}
              />

              <FormField
                control={form.control}
                name="accountRecommendations"
                render={({ field }) => (
                  <FormItem className="flex items-center md:items-end gap-2">
                    <Switch
                      className="data-[state=checked]:bg-[#16DBCC]"
                      checked={field.value}
                      onCheckedChange={field.onChange}
                    />
                    <label htmlFor={field.name} className="space-y-0.5">
                      There are recommendations for my account
                    </label>
                  </FormItem>
                )}
              />
            </div>
          </div>
          <Button
            disabled={loading}
            type="submit"
            className="mt-5 md:w-auto w-full px-8 float-end bg-primaryBlue text-white"
          >
            {loading ? "Loading..." : "Save"}
          </Button>
        </form>
      </Form>
    </div>
  );
};

export default Preferences;
