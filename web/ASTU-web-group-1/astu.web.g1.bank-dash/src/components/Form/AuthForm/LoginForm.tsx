import React, { use, useState } from "react";
import InputGroup from "../InputGroup";

const LoginForm = () => {
  return (
    <form className="flex flex-col items-center w-full justify-center   md:px-16">
      {/* <p className="text-center text-red-600 text-sm mb-6">Credential Error</p> */}
      <div className="w-full flex flex-col">
        <InputGroup
          id="email"
          label="Email Address"
          inputType="email"
          registerName="email"
          register=""
          placeholder="Enter email address"
        />
        <InputGroup
          id="password"
          label="Password"
          inputType="password"
          registerName="password"
          register=""
          placeholder="Enter password"
        />
      </div>

      <button
        type="submit"
        className="bg-[#1814f3] text-white px-10 py-3 font-Lato font-bold rounded-lg w-full mt-4"
      >
        Login
      </button>
    </form>
  );
};

export default LoginForm;
