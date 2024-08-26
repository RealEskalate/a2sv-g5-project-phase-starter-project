"use client";
import React, { useEffect } from "react";
import { useForm } from "react-hook-form";
import {
  userUpdatePreference,
  getCurrentUser,
} from "../../lib/api/userControl";
import { getSession } from "next-auth/react";
import { changePassword } from "../../lib/api/authenticationController";
type Data = {
  access_token: string;
  data: string;
  refresh_token: string;
};

type SessionDataType = {
  user: Data;
};

const HeadingLabel = ({ label }: { label: string }) => {
  return (
    <h1 className="text-sm font-medium lg:text-lg text-[#333B69] dark:text-[#9faaeb]">
      {label}
    </h1>
  );
};

const InputLabel = ({ label, htmlFor }: { label: string; htmlFor: string }) => {
  return (
    <label
      htmlFor={htmlFor}
      className="text-xs text-[#232323] lg:text-base dark:text-[#feffffc7]"
    >
      {label}
    </label>
  );
};

import { useState } from "react";
import { useRouter } from "next/navigation";

const ToggleSwitch = ({
  handleToogle,
  enabled,
}: {
  enabled: boolean;
  handleToogle: Function;
}) => {
  const toggleSwitch = () => {
    handleToogle(!enabled);
  };

  return (
    <div className="flex items-center gap-3">
      <input
        type="checkbox"
        id="two-factor-toggle"
        className="peer hidden"
        checked={enabled}
        onChange={toggleSwitch}
      />
      <label
        htmlFor="two-factor-toggle"
        className={`cursor-pointer rounded-full w-14 h-7 flex items-center relative transition-colors duration-300 ${
          enabled ? "bg-[#16DBCC]" : "bg-gray-200"
        }`}
      >
        <span
          className={`bg-white w-6 h-6 rounded-full transition-transform duration-300 ${
            enabled ? "translate-x-7" : ""
          }`}
        ></span>
      </label>

      <InputLabel
        label="Enable or disable two-factor authentication"
        htmlFor="two-factor-toggle"
      />
    </div>
  );
};

interface SecurityFormInputs {
  old_password: string;
  new_password: string;
}

const SecuritySetting = () => {
  const [session, setSession] = useState<Data | null>(null);
  const [loading, setLoading] = useState(true);

  const {
    register,
    handleSubmit,
    formState: { errors },
    reset,
  } = useForm<SecurityFormInputs>();
  const router = useRouter();

  const [enabled, setEnabled] = useState(false);

  useEffect(() => {
    const fetchSession = async () => {
      const sessionData = (await getSession()) as SessionDataType | null;
      if (sessionData && sessionData.user) {
        setSession(sessionData.user);
        setLoading(false);
      } else {
        router.push(
          `./api/auth/signin?callbackUrl=${encodeURIComponent("/accounts")}`
        );
      }
    };

    fetchSession();
  }, [router]);

  const handleToogle = (value: boolean) => {
    setEnabled(value);
  };

  const onSubmit = async (data: SecurityFormInputs) => {
    const user = await getCurrentUser(session?.access_token!);
    user.preference.twoFactorAuthentication = enabled;
    await userUpdatePreference(user.preference, session?.access_token!);
    await changePassword(
      { password: data.old_password, newPassword: data.new_password },
      session?.access_token!
    );
    reset();
  };

  if (loading) {
    return null;
  }

  return (
    <div className="">
      <form
        action=""
        className="flex flex-col gap-4"
        onSubmit={handleSubmit(onSubmit)}
      >
        <HeadingLabel label="Two-factor Authentication" />
        <div className="flex items-center">
          <ToggleSwitch handleToogle={handleToogle} enabled={enabled} />
        </div>

        <HeadingLabel label="Change Password" />
        <div className="space-y-2">
          <InputLabel label="Current Password" htmlFor="old_password" />
          <input
            type="password"
            id="old_password"
            placeholder="**********"
            className="border rounded-xl px-3 py-2 text-lg flex items-center w-full md:w-64 dark:bg-[#050914] dark:border-[#333B69] "
            {...register("old_password", {
              required: "Current Password is required",
              minLength: {
                value: 6,
                message: "Password must be at least 8 characters",
              },
            })}
          />
          {errors.old_password && (
            <span className="text-red-600 text-xs">
              {errors.old_password.message}
            </span>
          )}
        </div>
        <div className="space-y-2">
          <InputLabel label="New Password" htmlFor="new_password" />
          <input
            type="password"
            id="new_password"
            placeholder="**********"
            className="border rounded-xl px-3 py-2 text-lg flex items-center w-full md:w-64 dark:bg-[#050914] dark:border-[#333B69] "
            {...register("new_password", {
              required: "New Password is required",
              minLength: {
                value: 6,
                message: "Password must be at least 8 characters",
              },
              validate: (value) =>
                value !== "" ||
                "New Password cannot be the same as the old password",
            })}
          />
          {errors.new_password && (
            <span className="text-red-600 text-xs">
              {errors.new_password.message}
            </span>
          )}
        </div>
        <div className="w-full flex justify-end">
          <button
            type="submit"
            className="bg-[#1814F3] hover:bg-[#423fef] text-white px-5 py-3 rounded-xl md:w-2/12 text-sm w-full self-end"
          >
            Save
          </button>
        </div>
      </form>
    </div>
  );
};

export default SecuritySetting;
