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
import formSchema from "@/schema";
import { format } from "date-fns";
import { CalendarIcon } from "@radix-ui/react-icons";
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from "@/components/ui/popover";
import { Calendar } from "@/components/ui/calendar";

export function ProfileForm() {
  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      username: "",
      email: "",
      dateOfBirth: "",
      permanentAddress: "",
      postalCode: "",
      presentAddress: "",
      city: "",
      country: "",
    },
  });

  function onSubmit(values: z.infer<typeof formSchema>) {
    console.log(values);
  }

  return (
    <Form {...form}>
      <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-8">
        <div className="gap-5 md:flex">
          <div>
            <FormField
              control={form.control}
              name="name"
              render={({ field }) => (
                <FormItem className="mb-1">
                  <FormLabel>Name</FormLabel>
                  <FormControl>
                    <Input
                      className="rounded-2xl md:w-[300px]"
                      placeholder="Your Name"
                      {...field}
                    />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
            <FormField
              control={form.control}
              name="email"
              render={({ field }) => (
                <FormItem className="mb-1">
                  <FormLabel>Email</FormLabel>
                  <FormControl>
                    <Input
                      className="rounded-2xl"
                      placeholder="you@example.com"
                      {...field}
                    />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
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
                          className="md:w-[300px] w-full rounded-2xl text-left text-[#77809c] font-[400] flex justify-start pl-4"
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
            <FormField
              control={form.control}
              name="permanentAddress"
              render={({ field }) => (
                <FormItem className="mb-1">
                  <FormLabel>Permanent Address</FormLabel>
                  <FormControl>
                    <Input
                      className="rounded-2xl"
                      placeholder="Your Permanent Address"
                      {...field}
                    />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
            <FormField
              control={form.control}
              name="postalCode"
              render={({ field }) => (
                <FormItem className="mb-1">
                  <FormLabel>Postal Code</FormLabel>
                  <FormControl>
                    <Input
                      className="rounded-2xl"
                      placeholder="Postal Code"
                      {...field}
                    />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
          </div>
          <div>
            <FormField
              control={form.control}
              name="presentAddress"
              render={({ field }) => (
                <FormItem className="mb-1">
                  <FormLabel>Present Address</FormLabel>
                  <FormControl>
                    <Input
                      className="rounded-2xl md:w-[300px]"
                      placeholder="Your Present Address"
                      {...field}
                    />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
            <FormField
              control={form.control}
              name="city"
              render={({ field }) => (
                <FormItem className="mb-1">
                  <FormLabel>City</FormLabel>
                  <FormControl>
                    <Input
                      className="rounded-2xl"
                      placeholder="City"
                      {...field}
                    />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
            <FormField
              control={form.control}
              name="country"
              render={({ field }) => (
                <FormItem className="mb-1">
                  <FormLabel>Country</FormLabel>
                  <FormControl>
                    <Input
                      className="rounded-2xl"
                      placeholder="Country"
                      {...field}
                    />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
            <FormField
              control={form.control}
              name="username"
              render={({ field }) => (
                <FormItem className="mb-1">
                  <FormLabel>Username</FormLabel>
                  <FormControl>
                    <Input
                      type="text"
                      className="rounded-2xl"
                      placeholder="Username"
                      {...field}
                    />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
            <FormField
              control={form.control}
              name="password"
              render={({ field }) => (
                <FormItem className="mb-1">
                  <FormLabel>Password</FormLabel>
                  <FormControl>
                    <Input
                      className="rounded-2xl"
                      type="password"
                      placeholder="Password"
                      {...field}
                    />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
          </div>
        </div>
        <Button
          type="submit"
          className="md:w-auto w-full px-8 float-end bg-primaryBlue text-white"
        >
          Save
        </Button>
      </form>
    </Form>
  );
}
