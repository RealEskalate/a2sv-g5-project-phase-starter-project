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

const Preferences = () => {
  const formSchema = z.object({
    digitalCurrency: z.boolean().default(true).optional(),
    merchantOrder: z.boolean().default(false).optional(),
    recommendation: z.boolean().default(true).optional(),
    currency: z.string(),
    timeZone: z.string(),
  });

  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      currency: "",
      timeZone: "",
      digitalCurrency: true,
      merchantOrder: false,
      recommendation: true,
    },
  });

  function onSubmit(values: z.infer<typeof formSchema>) {
    console.log(values);
  }
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

          <div className="mt-5 ">
            <h1 className="mb-2 text-primaryBlack font-bold">Notification</h1>
            <div className="flex flex-col gap-2">
              <FormField
                control={form.control}
                name="digitalCurrency"
                render={({ field }) => (
                  <FormItem className="flex items-end gap-2">
                    <Switch
                      className="data-[state=checked]:bg-[#16DBCC]"
                      checked={field.value}
                      onCheckedChange={field.onChange}
                    />
                    <div className="space-y-0.5">
                      I send or receive digital currency
                    </div>
                  </FormItem>
                )}
              />

              <FormField
                control={form.control}
                name="merchantOrder"
                render={({ field }) => (
                  <FormItem className="flex items-end gap-2">
                    <Switch
                      className="data-[state=checked]:bg-[#16DBCC]"
                      checked={field.value}
                      onCheckedChange={field.onChange}
                    />
                    <div className="space-y-0.5">I receive merchant order</div>
                  </FormItem>
                )}
              />

              <FormField
                control={form.control}
                name="recommendation"
                render={({ field }) => (
                  <FormItem className="flex items-center md:items-end gap-2 ">
                    <Switch
                      className="data-[state=checked]:bg-[#16DBCC]"
                      checked={field.value}
                      onCheckedChange={field.onChange}
                    />
                    <div className="space-y-0.5">
                      There are recommendation for my account
                    </div>
                  </FormItem>
                )}
              />
            </div>
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

export default Preferences;
