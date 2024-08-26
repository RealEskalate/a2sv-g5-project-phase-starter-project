"use client";
import React from "react";
import Image from "next/image";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import {
  faCreditCard,
  faCreditCardAlt,
  faSpinner,
} from "@fortawesome/free-solid-svg-icons";

const NotFound = () => {
  return (
    <div className="w-full flex flex-col px-6 gap-4 min-h-screen items-center justify-center">
      <Image src={"/assets/pageNotFound.svg"} height={300} width={300} alt="" />
      <h2 className="text-3xl text-gray-500 dark:text-white">
        404 (page not found)
      </h2>
    </div>
  );
};

export default NotFound;
