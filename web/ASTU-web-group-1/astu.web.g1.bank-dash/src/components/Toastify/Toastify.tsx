import React from "react";
import "react-toastify/dist/ReactToastify.css";
import { toast } from "react-toastify";

const Toastify = () => {
  return <div></div>;
};

export default Toastify;

export const toastError = (message: string) =>
  toast.error(message, { theme: "light" });

export const toastSuccess = (message: string) =>
  toast.success(message, { theme: "light" });
