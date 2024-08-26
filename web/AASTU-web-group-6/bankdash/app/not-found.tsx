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
    <div className="w-full flex flex-col px-6 gap-2 min-h-screen items-center justify-center">
      <div className="b flex items-center justify-center p-2 bg-blue-100 rounded-full">
        <div className="card-loader p-6 bg-blue-200 rounded-full">
          <FontAwesomeIcon
            icon={faCreditCard}
            className="atm-card text-[42px] text-[#1814f6]"
          />
        </div>
      </div>
    </div>
  );
};

export default NotFound;
